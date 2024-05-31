package parser

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/o-richard/fit/pkg/db"
)

const (
	samsungDirectory      = "./data/parsers/samsung"
	samsungSourceName     = "samsung"
	samsungDateTimeLayout = "2006-01-02 15:04:05.000"
)

type samsung struct{}

func (samsung) getEntries() (entries []db.HealthEntry, mostRecentTimestamp time.Time, source string, err error) {
	pedometerDaySummary := ".tracker.pedometer_day_summary."

	filepaths, err := filePathWalk(samsungDirectory, []string{pedometerDaySummary})
	if err != nil {
		return
	}
	timezoneLocation, err := parseSamsungTimezone()
	if err != nil {
		return
	}

	appdb, _ := db.NewDB()
	lastUpdatedAt, err := appdb.GetRecentFitnessSync(samsungSourceName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("another process is currently parsing samsung files. if this is a mistake, manually update the locked status of the 'samsung' record in the fitness sync table")
			return
		}
		err = fmt.Errorf("database error while checking the timestamp of the most recent fitness record from Samsung Health, %w", err)
		return
	}

	mostRecentTimestamp = lastUpdatedAt
	for filename, filepath := range filepaths {
		// any absent file is ignored
		if filepath == "" {
			continue
		}

		fullFilepath := fmt.Sprintf("%v/%v", samsungDirectory, filepath)
		var specificEntries []db.HealthEntry
		var specificRecentTimestamp time.Time
		switch filename {
		case pedometerDaySummary:
			specificEntries, specificRecentTimestamp, err = parseSamsungPedometerDaySummary(fullFilepath, lastUpdatedAt, timezoneLocation)
		}

		if err != nil {
			_ = appdb.UpdateFitnessSync(samsungSourceName, lastUpdatedAt)
			return
		}
		if specificRecentTimestamp.Compare(mostRecentTimestamp) > 0 {
			mostRecentTimestamp = specificRecentTimestamp
		}
		entries = append(entries, specificEntries...)
	}
	source = samsungSourceName
	err = nil
	return
}

/*
Attempts to find the timezone of the timestamps used in the files. Samsung Health does not record the timezone information together with the datetime.
The timezone is found in files whose names contains the substrings (.badge. .exercise. .report. .pedometer_step_count.). At least one of the files should be present.

Assumption: The first timezone found in any of the listed file corresponds to all the datetimes present in the files.
*/
func parseSamsungTimezone() (*time.Location, error) {
	badge := ".badge."
	report := ".report."
	exercise := ".exercise."
	pedometerStepCount := ".pedometer_step_count."

	filepaths, err := filePathWalk(samsungDirectory, []string{badge, report, exercise, pedometerStepCount})
	if err != nil {
		return nil, err
	}

	var location *time.Location
	for filename, filepath := range filepaths {
		// any absent file is ignored
		if filepath == "" {
			continue
		}

		fullFilepath := fmt.Sprintf("%v/%v", samsungDirectory, filepath)
		var records []map[string]string
		switch filename {
		case badge:
			records, err = parseCSV(fullFilepath, []string{"time_offset"}, 1, 1)
		case report:
			records, err = parseCSV(fullFilepath, []string{"timezone"}, 1, 1)
		case exercise:
			records, err = parseCSV(fullFilepath, []string{"com.samsung.health.exercise.time_offset"}, 1, 1)
		case pedometerStepCount:
			records, err = parseCSV(fullFilepath, []string{"com.samsung.health.step_count.time_offset"}, 1, 1)
		}
		if err != nil {
			return nil, fmt.Errorf("unable to obtain the timezone of the records, %w", err)
		}
		if len(records) == 0 {
			continue
		}

		// UTC+0300 - when retrieved from badge, exercise, pedometerStepCount
		// Africa/Nairobi - when retrieved from report
		var timezone string
		switch filename {
		case badge:
			timezone, _ = strings.CutPrefix(records[0]["time_offset"], "UTC")
		case report:
			timezone = records[0]["timezone"]
		case exercise:
			timezone, _ = strings.CutPrefix(records[0]["com.samsung.health.exercise.time_offset"], "UTC")
		case pedometerStepCount:
			timezone, _ = strings.CutPrefix(records[0]["com.samsung.health.step_count.time_offset"], "UTC")
		}
		if timezone == "" {
			continue
		}

		var t time.Time
		var loc *time.Location
		switch filename {
		case badge, exercise, pedometerStepCount:
			t, err = time.Parse("-0700", timezone)
			loc = t.Location()
		case report:
			loc, err = time.LoadLocation(timezone)
		}
		if err != nil {
			return nil, fmt.Errorf("unable to obtain the timezone of the records, %w", err)
		}

		location = loc
		break
	}

	if location == nil {
		return nil, fmt.Errorf("unable to obtain a timezone to use. the files used to extract the timezone may be missing or the timezone information in the file might be missing")
	}
	return location, nil
}

/*
Returns the records of health entries & the most recent timestamp found.

Assumption: all parsed columns are required hence no need of assigning defaults.
*/
func parseSamsungPedometerDaySummary(fullFilepath string, lastUpdatedAt time.Time, timezoneLocation *time.Location) ([]db.HealthEntry, time.Time, error) {
	records, err := parseCSV(fullFilepath, []string{"step_count", "update_time", "create_time", "distance", "calorie"}, 1, 0)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("unable to read %v, %w", fullFilepath, err)
	}

	var entries []db.HealthEntry
	mostRecentTimestamp := lastUpdatedAt
	for _, record := range records {
		startedAt, err := time.ParseInLocation(samsungDateTimeLayout, record["create_time"], timezoneLocation)
		if err != nil {
			return nil, time.Time{}, fmt.Errorf("unable to parse a timestamp, %w", err)
		}
		endedAt, err := time.ParseInLocation(samsungDateTimeLayout, record["update_time"], timezoneLocation)
		if err != nil {
			return nil, time.Time{}, fmt.Errorf("unable to parse a timestamp, %w", err)
		}

		/*
			In my files, some records are duplicated with the different time ranges (create_time TO update_time).
			Where the two datetimes match, it means it is when the sync took place.
			Where the two datetimes are not matching, it means the period when the record was taken.
		*/
		if endedAt.Compare(startedAt) <= 0 || endedAt.Compare(lastUpdatedAt) <= 0 {
			continue
		}
		if endedAt.Compare(mostRecentTimestamp) > 0 {
			mostRecentTimestamp = endedAt
		}

		title := "Pedometer Summary"
		content := fmt.Sprintf("I covered %v metres in %v steps. I burnt %v kilocalories (kcal).", record["distance"], record["step_count"], record["calorie"])
		entries = append(entries, db.HealthEntry{Type: db.Activity, Title: title, Content: content, StartedAt: startedAt, EndedAt: endedAt})
	}
	return entries, mostRecentTimestamp, nil
}

/*
Returns the records of health entries & the most recent timestamp found.

TODO: Determine the extra JSON data to parse.

Assumption: all parsed columns are required hence no need of assigning defaults.
*/
func parseSamsungReport(fullFilepath string, lastUpdatedAt time.Time, timezoneLocation *time.Location) ([]db.HealthEntry, time.Time, error) {
	records, err := parseCSV(fullFilepath, []string{"update_time", "compressed_content"}, 1, 0)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("unable to read %v, %w", fullFilepath, err)
	}

	var entries []db.HealthEntry
	mostRecentTimestamp := lastUpdatedAt
	for _, record := range records {
		updatedAt, err := time.ParseInLocation(samsungDateTimeLayout, record["update_time"], timezoneLocation)
		if err != nil {
			return nil, time.Time{}, fmt.Errorf("unable to parse a timestamp, %w", err)
		}
		var firstCharacter string
		if record["compressed_content"] != "" {
			firstCharacter = string(record["compressed_content"][0])
		}

		filePath := fmt.Sprintf("%v/jsons/com.samsung.shealth.report/%v/%v", samsungDirectory, firstCharacter, record["compressed_content"])
		file, err := os.Open(filePath)
		if err != nil {
			return nil, time.Time{}, fmt.Errorf("unable to open a json file '%v', %w", filePath, err)
		}

		data, _ := io.ReadAll(file)
		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, time.Time{}, fmt.Errorf("unable to decode the json file '%v', %w", filePath, err)
		}

		_ = file.Close()

		if updatedAt.Compare(mostRecentTimestamp) > 0 {
			mostRecentTimestamp = updatedAt
		}
		// TODO: parse result.
	}
	return entries, mostRecentTimestamp, nil
}

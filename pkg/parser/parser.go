package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/o-richard/fit/pkg/db"
)

type Parser interface {
	getEntries() (entries []db.HealthEntry, mostRecentTimestamp time.Time, source string, err error)
}

func ParseFitnessAppRecords(parserName string) error {
	parsers := map[string]Parser{"samsung": samsung{}}
	if _, ok := parsers[parserName]; !ok {
		return fmt.Errorf("unregistered parser, ensure the parser '%v' exists", parserName)
	}

	entries, mostRecentTimestamp, source, err := parsers[parserName].getEntries()
	if err != nil {
		return err
	}

	appdb, _ := db.NewDB()
	if len(entries) != 0 {
		if err := appdb.InsertHealthEntries(false, entries); err != nil {
			return fmt.Errorf("the entries from the fitness app could not be inserted to the databse, %w", err)
		}
	}

	if err := appdb.UpdateFitnessSync(source, mostRecentTimestamp); err != nil {
		if errors.Is(err, db.ErrNoAffectedRows) {
			return fmt.Errorf("the fitness sync record for the source '%v' does not exist. SOMETHING CRAZY HAPPENED!", source)
		}
		return fmt.Errorf("unable to update the fitness sync record for the source '%v'. YOU NEED TO GET YOUR HANDS DIRTY & UPDATE THE LOCKED STATUS IN THE DATABASE.", source)
	}
	return nil
}

// returns csv records.
// allowed headers consists of a valid list of header names matching the first row of the csv file. if it is an empty slice, all headers are included in the result.
// offset dictates how many rows to skip before reading the first valid row. the first valid row contains the headers of the csv records.
// rowsToReturn dictates how many rows to return after the row containing the header. zero dictates that all rows are returned. definitely the returned rows may be less than the provided value.
func parseCSV(filepath string, allowedHeaders []string, offset, rowsToReturn uint) ([]map[string]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error while opening csv file, %w", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	// allow for variable number of fields. the skipped rows likely contain metadata.
	csvReader.FieldsPerRecord = -1
	// accomodate for the provided offset
	for i := 0; i < int(offset); i++ {
		if _, err := csvReader.Read(); err != nil {
			return nil, fmt.Errorf("error while skipping the first %v rows in the file %v, %w", offset, filepath, err)
		}
	}
	csvReader.FieldsPerRecord = 0
	// Obtain csv headers
	headers, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("error while parsing csv headers, %w", err)
	}
	// Obtain indices of concern
	indices := make([]int, 0, len(headers))
	for i, header := range allowedHeaders {
		index := slices.Index(headers, header)
		if index < 0 {
			return nil, fmt.Errorf("header %v does not exist in the file %v", header, filepath)
		}
		indices = append(indices, i)
	}
	if len(indices) == 0 {
		for i := 0; i < len(headers); i++ {
			indices = append(indices, i)
		}
	}

	var parsedRows int
	var records []map[string]string
	for {
		if rowsToReturn != 0 && parsedRows == int(rowsToReturn) {
			break
		}

		row, err := csvReader.Read()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, fmt.Errorf("error while parsing csv data, %w", err)
			}
			break
		}
		data := make(map[string]string, len(indices))
		for _, index := range indices {
			data[headers[index]] = row[index]
		}
		records = append(records, data)
		parsedRows++
	}
	return records, nil
}

// returns filenames in the specified directory with the provided filenames as substrings (should be unique). if it is empty, no files are returned.
// for the returned filenames, the first substring match is considered so the provided filenames should be as distinguishable for correct results.
func filePathWalk(directorypath string, filenames []string) (map[string]string, error) {
	entries, err := os.ReadDir(directorypath)
	if err != nil {
		return nil, fmt.Errorf("error while opening directory, %w", err)
	}

	files := make(map[string]string, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			for _, filename := range filenames {
				if strings.Contains(entry.Name(), filename) {
					files[filename] = entry.Name()
					break
				}
			}
		}
	}
	return files, nil
}

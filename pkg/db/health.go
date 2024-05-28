package db

import (
	"fmt"
	"strings"
	"time"
)

type HealthEntry struct {
	Title   string
	Content string
	// comma-separated image paths (relative paths)
	Images    string
	StartedAt time.Time
	EndedAt   time.Time
}

// called during insertion to the database
func (h *HealthEntry) Validate() bool {
	h.Title = strings.TrimSpace(h.Title)
	h.Content = strings.TrimSpace(h.Content)
	return h.Content != "" && h.EndedAt.Compare(h.StartedAt) >= 0
}

type HealthEntryOfDay struct {
	ID      int
	ByUser  bool
	Title   string
	Content string
	Images  string
	// number of minutes from midnight
	StartedAtMinutes uint
	// number of minutes from midnight
	EndedAtMinutes uint
}

/*
Inserts multiple health entries at once.

Considered fields: title, content, images, startedAt, endedAt

Expected errors (apart from nil & unexpected errors):
  - ErrValidation
*/
func (db *DB) InsertHealthEntries(byUser bool, entries []HealthEntry) error {
	if len(entries) == 0 {
		return ErrValidation
	}

	query := strings.Builder{}
	_, _ = query.WriteString(` INSERT INTO entry (by_user, title, content, images, started_at, ended_at) VALUES `)
	maxIndex := len(entries) - 1
	args := make([]any, 0, len(entries)*5)
	for i := range entries {
		if !entries[i].Validate() {
			return ErrValidation
		}

		_, _ = query.WriteString(` (?,?,?,?,?,?)`)
		if i != maxIndex {
			_, _ = query.WriteString(`,`)
		} else {
			_, _ = query.WriteString(`;`)
		}
		args = append(args, byUser, entries[i].Title, entries[i].Content, entries[i].Images, entries[i].StartedAt, entries[i].EndedAt)
	}

	_, err := db.db.Exec(query.String(), args...)
	return err
}

/*
Retrieve distinct years of the health entries. (4-digit string).

Only nil & unexpected errors.
*/
func (db *DB) GetUniqueYearsOfHealthEntries() ([]string, error) {
	query := `
		SELECT strftime('%Y', started_at, 'localtime') AS year FROM entry
		UNION
		SELECT strftime('%Y', ended_at, 'localtime') AS year FROM entry`
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var years []string
	for rows.Next() {
		var year string
		if err := rows.Scan(&year); err != nil {
			return nil, err
		}
		years = append(years, year)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return years, nil
}

/*
Retrieve distinct months of the health entries of the provided year. (2-digit string).

Only nil & unexpected errors.
*/
func (db *DB) GetUniqueMonthsOfHealthEntries(year int) ([]string, error) {
	sYear := fmt.Sprintf("%.4d", year)
	query := `
		SELECT strftime('%m', started_at, 'localtime') AS month FROM entry WHERE strftime('%Y', started_at, 'localtime') = ?
		UNION
		SELECT strftime('%m', ended_at, 'localtime') AS month FROM entry WHERE strftime('%Y', ended_at, 'localtime') = ?`
	args := []any{sYear, sYear}
	rows, err := db.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var months []string
	for rows.Next() {
		var month string
		if err := rows.Scan(&month); err != nil {
			return nil, err
		}
		months = append(months, month)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return months, nil
}

/*
Retrieve distinct days of the health entries of the provided year and month. (2-digit string).

Only nil & unexpected errors.
*/
func (db *DB) GetUniqueDaysOfHealthEntries(year, month int) ([]string, error) {
	sYear := fmt.Sprintf("%.4d", year)
	sMonth := fmt.Sprintf("%.2d", month)
	query := `
		SELECT strftime('%d', started_at, 'localtime') AS day FROM entry WHERE strftime('%Y', started_at, 'localtime') = ? AND strftime('%m', started_at, 'localtime') = ?
		UNION
		SELECT strftime('%d', ended_at, 'localtime') AS day FROM entry WHERE strftime('%Y', ended_at, 'localtime') = ? AND strftime('%m', ended_at, 'localtime') = ?`
	args := []any{sYear, sMonth, sYear, sMonth}
	rows, err := db.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var days []string
	for rows.Next() {
		var day string
		if err := rows.Scan(&day); err != nil {
			return nil, err
		}
		days = append(days, day)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return days, nil
}

/*
Retrieve health entries of a particular date ordered by the startedAt field.

In case the startedAt field's date is not of the provided date but the endedAt's field date is, the start time becomes 0.00.
In case the endedAt field's date is not of the provided date but the startedAt's field date is, the end time becomes 23.59.

Only nil & unexpected errors.
*/
func (db *DB) GetHealthEntries(year, month, day int) ([]HealthEntryOfDay, error) {
	date := fmt.Sprintf("%.4d-%.2d-%.2d", year, month, day)
	query := `
		SELECT
			id, by_user, title, content, images,
			CASE 
				WHEN date(started_at, 'localtime') <> ? THEN 0
				ELSE ((CAST(strftime('%k', started_at, 'localtime') AS INT) * 60) + (CAST(strftime('%M', started_at, 'localtime') AS INT)))
			END AS started_at_minutes,
			CASE 
				WHEN date(ended_at, 'localtime') <> ? THEN ((23 * 60) + 59)
				ELSE ((CAST(strftime('%k', ended_at, 'localtime') AS INT) * 60) + (CAST(strftime('%M', ended_at, 'localtime') AS INT)))
			END AS ended_at_minutes
		FROM entry WHERE date(started_at, 'localtime') = ? OR date(ended_at, 'localtime') = ? ORDER BY started_at_minutes ASC`
	args := []any{date, date, date, date}
	rows, err := db.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []HealthEntryOfDay
	for rows.Next() {
		var entry HealthEntryOfDay
		if err := rows.Scan(&entry.ID, &entry.ByUser, &entry.Title, &entry.Content, &entry.Images, &entry.StartedAtMinutes, &entry.EndedAtMinutes); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

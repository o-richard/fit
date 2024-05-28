package db

import (
	"database/sql"
	"errors"
	"time"
)

/*
Retrieves the most recent fitness sync of the provided source. Locks the source from modification by another process.

Expected errors (apart from nil & unexpected errors):
  - sql.ErrNoRows (the source is already locked from modification)
*/
func (db *DB) GetRecentFitnessSync(source string) (time.Time, error) {
	query := `
	INSERT INTO fitness_sync (source, last_updated_at, is_locked) VALUES (?, ?, TRUE) ON CONFLICT (source) 
	DO UPDATE SET is_locked=excluded.is_locked WHERE NOT fitness_sync.is_locked RETURNING last_updated_at`
	args := []any{source, time.Time{}}

	var lastUpdatedAt time.Time
	err := db.db.QueryRow(query, args...).Scan(&lastUpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return time.Time{}, sql.ErrNoRows
	}
	return lastUpdatedAt, nil
}

/*
Updates the last updated timestamp of the source to match the provided time. Unlocks the source for modification by another process.

NOTE: ideally, the provided timestamp should be greater/equal to the database value & the source should have a locked status but neither is considered during the update.

Expected errors (apart from nil & unexpected errors):
  - ErrNoAffectedRows (the source record does not exist)
*/
func (db *DB) UpdateFitnessSync(source string, lastUpdatedAt time.Time) error {
	query := `UPDATE fitness_sync SET last_updated_at = ?, is_locked=FALSE WHERE source = ?`
	args := []any{lastUpdatedAt, source}
	result, err := db.db.Exec(query, args...)
	if err != nil {
		return err
	}
	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return ErrNoAffectedRows
	}
	return nil
}

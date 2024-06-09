package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

const dbTimestampLayout = "2006-01-02 15:04:05.999999999-07:00"

var (
	dbInstance *DB
	dbOnce     sync.Once
	errDB      error

	// invalid user input
	ErrValidation = errors.New("invalid data provided")
	// no rows affected by the operation
	ErrNoAffectedRows = errors.New("no rows were affected by the operation")
)

// creates a new database connection if none existed
func NewDB() (*DB, error) {
	dbOnce.Do(func() {
		db, err := setupDB()
		if err != nil {
			errDB = err
			return
		}
		dbInstance = &DB{db: db}
	})
	return dbInstance, errDB
}

func setupDB() (*sql.DB, error) {
	var isCreated bool
	databaseName := "db.sqlite3"
	if _, err := os.Stat(databaseName); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, err := os.Create(databaseName)
			if err != nil {
				return nil, fmt.Errorf("error while creating database: %w", err)
			}
			_ = file.Close()
			isCreated = true
		} else {
			return nil, fmt.Errorf("error while checking database existence: %w", err)
		}
	}

	db, err := sql.Open("sqlite3", databaseName)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to the database: %w", err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error while verifying connection to the database: %w", err)
	}
	if isCreated {
		if err = performDatabaseMigrations(db); err != nil {
			db.Close()
			if errPath := os.Remove(databaseName); errPath != nil {
				return nil, fmt.Errorf("error while deleting newly created database, remove the file '%v' manually if it exists: %w", databaseName, errPath)
			}
			return nil, fmt.Errorf("error while performing migrations on newly created database: %w", err)
		}
	}
	return db, nil
}

func performDatabaseMigrations(db *sql.DB) error {
	// timestamps (suffixed with '_at') are stored as texts in UTC timezone.
	migrations := `
		/* Tracks the most recent timestamp in the parsed files (already parsed data does not get re-parsed). */
		CREATE TABLE fitness_sync (
			id INTEGER PRIMARY KEY,
			source TEXT NOT NULL UNIQUE,
			/* the timestamp of the most recent record from the fitness app */
			last_updated_at TEXT NOT NULL,
			/* whether the source is currently being updated */
			is_locked BOOLEAN NOT NULL DEFAULT FALSE
		);
		/* Tracks health entries from fitness apps or user. */
		CREATE TABLE entry (
			id INTEGER PRIMARY KEY,
			type TEXT CHECK (type IN ('sleep', 'nutrition', 'activity', 'health')) NOT NULL,
			by_user BOOLEAN NOT NULL,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			images TEXT NOT NULL,
			started_at TEXT NOT NULL,
			ended_at TEXT NOT NULL,
			CHECK (unixepoch(started_at) > 0 AND unixepoch(ended_at) > 0 AND (unixepoch(ended_at) - unixepoch(started_at)) >= 0)
		);
	`
	_, err := db.Exec(migrations)
	return err
}

func (db *DB) Close() error {
	return db.db.Close()
}

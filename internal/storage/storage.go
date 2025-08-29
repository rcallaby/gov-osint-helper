package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	db *sql.DB
}

func NewSQLite(path string) (*SQLite, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS results (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		results TEXT
	)`)
	if err != nil {
		return nil, err
	}

	return &SQLite{db: db}, nil
}

func (s *SQLite) SaveResult(name string, results map[string]string) error {
	stmt, err := s.db.Prepare("INSERT INTO results(name, results) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Simple serialization
	return stmt.QueryRow(name, fmt.Sprintf("%v", results)).Err()
}

func (s *SQLite) Close() {
	s.db.Close()
}

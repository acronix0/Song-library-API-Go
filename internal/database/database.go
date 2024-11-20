package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func New(port int, host, userName, password, dbName string) (*Database, error) {
	const op = "storage.postgres.New"
	conStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, userName, password, dbName,
	)
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (s *Database) GetDB() *sql.DB {
	return s.db
}
func (s *Database) Close() error {
	return s.db.Close()
}

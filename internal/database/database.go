package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(port int, host, userName, password, dbName string) (*Database, error) {
	const op = "database.New"

	conStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, userName, password, dbName,
	)
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err == nil {
		return &Database{db: db}, nil
	}

	err = createDatabase(port , host, userName, password, dbName )
		if err != nil {
		return nil, fmt.Errorf("%s: failed to create: %w", op, err)
	}


	db, err = sql.Open("postgres", conStr)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to reconnect to database: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: failed to ping database: %w", op, err)
	}

	return &Database{db: db}, nil
}

func createDatabase(port int, host, userName, password, dbName string)  error{
const op = "database.createDatabase"
	serverConStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, userName, password,
	)
	serverDB, err := sql.Open("postgres", serverConStr)
	if err != nil {
		return err
	}
	defer serverDB.Close()

	createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
	if _, err = serverDB.Exec(createDBQuery); err != nil {
		return err
	}
	return nil
}
func (s *Database) GetDB() *sql.DB {
	return s.db
}
func (s *Database) Close() error {
	return s.db.Close()
}

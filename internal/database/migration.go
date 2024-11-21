package database

import (
	"database/sql"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitMigrations(
	migrationsPath,
	host,
	userName,
	dbName,
	password string,
	port int,
) error {


	serverConStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, userName, password,
	)

	err := ensureDatabaseExists(serverConStr, dbName)
	if err != nil {
		return fmt.Errorf("failed to ensure database exists: %w", err)
	}
encodedPassword := url.QueryEscape(password)
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		userName,
		encodedPassword,
		host,
		port,
		dbName,
	)
	resolvedPath, err := resolvePath(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to resolve migrations path: %w", err)
	}
	m, err := migrate.New(resolvedPath, dbURL)
	if err != nil {
		return fmt.Errorf("failed to initialize migrations: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}

func resolvePath(migrationsPath string) (string, error) {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	normalizedPath := strings.ReplaceAll(absPath, "\\", "/")

	return "file://" + normalizedPath + "/", nil
}


func ensureDatabaseExists(serverConStr, dbName string) error {
	serverDB, err := sql.Open("postgres", serverConStr)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}
	defer serverDB.Close()

	if err := serverDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping server: %w", err)
	}

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
	err = serverDB.QueryRow(query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check database existence: %w", err)
	}

	if exists {
		return nil
	}

	createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
	_, err = serverDB.Exec(createDBQuery)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	return nil
}

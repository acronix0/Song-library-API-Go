package song

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/acronix0/song-libary-api/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestSongRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening a stub database connection: %s", err)
	}
	defer db.Close()

	repo := NewSongRepository(db)

	tests := []struct {
		name          string
		mockBehavior  func(mock sqlmock.Sqlmock)
		input         dto.CreateSongDTO
		expectedID    int
		expectedError error
	}{
		{
			name: "Valid input with minimal fields",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT id FROM groups WHERE name = \\$1").
					WithArgs("Test Group").
					WillReturnError(sql.ErrNoRows)
				mock.ExpectQuery("INSERT INTO groups \\(name, created_at, updated_at\\)").
					WithArgs("Test Group", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectQuery("INSERT INTO songs \\(group_id").
					WithArgs(1, "Test Song", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(42))
				mock.ExpectCommit()
			},
			input: dto.CreateSongDTO{
				Song:  "Test Song",
				Group: "Test Group",
			},
			expectedID:    42,
			expectedError: nil,
		},
		{
			name: "Group not found and creation fails",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT id FROM groups WHERE name = \\$1").
					WithArgs("Test Group").WillReturnError(sql.ErrNoRows)
				mock.ExpectQuery("INSERT INTO groups").
					WithArgs("Test Group", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			},
			input: dto.CreateSongDTO{
				Song:  "Test Song",
				Group: "Test Group",
			},
			expectedError: fmt.Errorf("failed to create group: %w", errors.New("database error")),
			expectedID:    0,
		},
		{
			name: "Song insertion fails",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT id FROM groups WHERE name = \\$1").
					WithArgs("Test Group").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectQuery("INSERT INTO songs").
					WithArgs(1, "Test Song", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			},
			input: dto.CreateSongDTO{
				Song:  "Test Song",
				Group: "Test Group",
			},
			expectedError: fmt.Errorf("failed to create song: %w", errors.New("database error")),
			expectedID:    0,
		},
		{
			name: "Group already exists",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT id FROM groups WHERE name = \\$1").
					WithArgs("Existing Group").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
				mock.ExpectQuery("INSERT INTO songs \\(group_id").
					WithArgs(2, "Existing Song", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(99))
				mock.ExpectCommit()
			},
			input: dto.CreateSongDTO{
				Song:  "Existing Song",
				Group: "Existing Group",
			},
			expectedID:    99,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(mock)

			id, err := repo.Create(context.Background(), tt.input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error(),
					"expected error containing: %v, got: %v", tt.expectedError.Error(), err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedID, id)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}

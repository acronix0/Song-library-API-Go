package song

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSongRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening a stub database connection:: %s", err)
	}
	defer db.Close()

	repo := NewSongRepository(db)

	tests := []struct {
		name          string
		mockBehavior  func(mock sqlmock.Sqlmock)
		input         int
		expectedError error
	}{
		{
			name: "Valid deletion",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM songs WHERE id = \\$1").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input:         1,
			expectedError: nil,
		},
		{
			name: "Song ID is zero",
			mockBehavior: func(mock sqlmock.Sqlmock) {

			},
			input:         0,
			expectedError: fmt.Errorf("song ID is required"),
		},
		{
			name: "Failed to delete song due to database error",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM songs WHERE id = \\$1").
					WithArgs(1).
					WillReturnError(fmt.Errorf("database error"))
			},
			input:         1,
			expectedError: fmt.Errorf("failed to delete song: database error"),
		},
		{
			name: "No rows affected",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM songs WHERE id = \\$1").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			input:         1,
			expectedError: sql.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(mock)

			err := repo.Delete(context.Background(), tt.input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error(),
					"expected error containing: %v, got: %v", tt.expectedError.Error(), err)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

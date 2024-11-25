package lyrics

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestLyricsRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening a stub database connection: %s", err)
	}
	defer db.Close()

	repo := NewLyricsRepository(db)

	tests := []struct {
		name          string
		mockBehavior  func(mock sqlmock.Sqlmock)
		songID        int
		text          string
		expectedError error
	}{
		{
			name: "Valid input",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO song_lyrics").
					WithArgs(1, 1, "Verse 1", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("INSERT INTO song_lyrics").
					WithArgs(1, 2, "Verse 2", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			songID:        1,
			text:          "Verse 1\nVerse 2",
			expectedError: nil,
		},
		{
			name: "Empty text",
			mockBehavior: func(mock sqlmock.Sqlmock) {

			},
			songID:        1,
			text:          "",
			expectedError: fmt.Errorf("song ID and text are required"),
		},
		{
			name: "Error during insert",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO song_lyrics").
					WithArgs(1, 1, "Verse 1", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(fmt.Errorf("database error"))

				mock.ExpectRollback()
			},
			songID:        1,
			text:          "Verse 1\nVerse 2",
			expectedError: fmt.Errorf("failed to insert verse 1: database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(mock)

			err := repo.Create(context.Background(), tt.songID, tt.text)

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

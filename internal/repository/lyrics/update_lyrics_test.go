package lyrics

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestLyricsRepo_Update(t *testing.T) {
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
			name: "Valid update",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM song_lyrics WHERE song_id = \\$1").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectExec("INSERT INTO song_lyrics \\(song_id, verse_number, text, created_at, updated_at\\)").
					WithArgs(1, 1, "Verse 1", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec("INSERT INTO song_lyrics \\(song_id, verse_number, text, created_at, updated_at\\)").
					WithArgs(1, 2, "Verse 2", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
			songID:        1,
			text:          "Verse 1\nVerse 2",
			expectedError: nil,
		},
		{
			name: "Failed to start transaction",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(fmt.Errorf("transaction error"))
			},
			songID:        1,
			text:          "Verse 1\nVerse 2",
			expectedError: fmt.Errorf("failed to start transaction: transaction error"),
		},
		{
			name: "Failed to delete old lyrics",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM song_lyrics WHERE song_id = \\$1").
					WithArgs(1).
					WillReturnError(fmt.Errorf("delete error"))

				mock.ExpectRollback()
			},
			songID:        1,
			text:          "Verse 1\nVerse 2",
			expectedError: fmt.Errorf("failed to delete old lyrics: delete error"),
		},
		{
			name: "Failed to insert new verse",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM song_lyrics WHERE song_id = \\$1").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectExec("INSERT INTO song_lyrics \\(song_id, verse_number, text, created_at, updated_at\\)").
					WithArgs(1, 1, "Verse 1", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(fmt.Errorf("insert error"))

				mock.ExpectRollback()
			},
			songID:        1,
			text:          "Verse 1\nVerse 2",
			expectedError: fmt.Errorf("failed to insert new verse: insert error"),
		},
		{
			name: "Invalid input: empty song ID",
			mockBehavior: func(mock sqlmock.Sqlmock) {
			},
			songID:        0,
			text:          "Verse 1\nVerse 2",
			expectedError: fmt.Errorf("song ID and text are required"),
		},
		{
			name: "Invalid input: empty text",
			mockBehavior: func(mock sqlmock.Sqlmock) {
			},
			songID:        1,
			text:          "",
			expectedError: fmt.Errorf("song ID and text are required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(mock)

			err := repo.Update(context.Background(), tt.songID, tt.text)

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

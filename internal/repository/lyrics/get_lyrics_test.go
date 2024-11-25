package lyrics

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestLyricsRepo_Get(t *testing.T) {
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
		skip, take    int
		expectedText  string
		expectedError error
	}{
		{
			name: "Valid input, lyrics found",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"text"}).
					AddRow("Verse 1").
					AddRow("Verse 2").
					AddRow("Verse 3")

				mock.ExpectQuery("SELECT text FROM song_lyrics").
					WithArgs(1, 3, 0).
					WillReturnRows(rows)
			},
			songID:        1,
			skip:          0,
			take:          3,
			expectedText:  "Verse 1\nVerse 2\nVerse 3",
			expectedError: nil,
		},
		{
			name: "No lyrics found",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"text"})
				mock.ExpectQuery("SELECT text FROM song_lyrics").
					WithArgs(1, 3, 0).
					WillReturnRows(rows)
			},
			songID:        1,
			skip:          0,
			take:          3,
			expectedText:  "",
			expectedError: sql.ErrNoRows,
		},
		{
			name: "Database query error",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT text FROM song_lyrics").
					WithArgs(1, 3, 0).
					WillReturnError(fmt.Errorf("database error"))
			},
			songID:        1,
			skip:          0,
			take:          3,
			expectedText:  "",
			expectedError: fmt.Errorf("failed to fetch lyrics: database error"),
		},
		{
			name: "Row iteration error",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"text"}).
					AddRow("Verse 1").
					AddRow("Verse 2").
					CloseError(fmt.Errorf("iteration error"))

				mock.ExpectQuery("SELECT text FROM song_lyrics").
					WithArgs(1, 3, 0).
					WillReturnRows(rows)
			},
			songID:        1,
			skip:          0,
			take:          3,
			expectedText:  "",
			expectedError: fmt.Errorf("error iterating over rows: iteration error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(mock)

			result, err := repo.Get(context.Background(), tt.songID, tt.skip, tt.take)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error(),
					"expected error containing: %v, got: %v", tt.expectedError.Error(), err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedText, result)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

package song

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/acronix0/song-libary-api/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestSongRepo_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening a stub database connection: %s", err)
	}
	defer db.Close()

	repo := NewSongRepository(db)

	tests := []struct {
		name           string
		mockBehavior   func(mock sqlmock.Sqlmock)
		skip, take     int
		expectedResult []dto.ResponseSongDTO
		expectedError  error
	}{
		{
			name: "Valid input, songs found",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"song_id", "title", "group_name", "link", "release_date", "created_at", "updated_at", "text",
				}).AddRow(
					1, "Song 1", "Group 1", "http://example.com",
					time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC), "Lyrics 1",
				).AddRow(
					2, "Song 2", "Group 2", nil,
					nil,
					time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
					time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC), "Lyrics 2",
				)

				mock.ExpectQuery("SELECT s.id AS song_id, s.title, g.name AS group_name").
					WithArgs(10, 0).
					WillReturnRows(rows)
			},
			skip: 0,
			take: 10,
			expectedResult: []dto.ResponseSongDTO{
				{
					SongID:      1,
					Song:        "Song 1",
					Group:       "Group 1",
					Link:        toPtr("http://example.com"),
					ReleaseDate: timePtr(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)),
					CreatedAt:   timePtr(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)),
					UpdatedAt:   timePtr(time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)),
					Text:        toPtr("Lyrics 1"),
				},
				{
					SongID:      2,
					Song:        "Song 2",
					Group:       "Group 2",
					Link:        nil,
					ReleaseDate: nil,
					CreatedAt:   timePtr(time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)),
					UpdatedAt:   timePtr(time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC)),
					Text:        toPtr("Lyrics 2"),
				},
			},
			expectedError: nil,
		},
		{
			name: "No songs found",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"song_id", "title", "group_name", "link", "release_date", "created_at", "updated_at", "text",
				})
				mock.ExpectQuery("SELECT s.id AS song_id, s.title, g.name AS group_name").
					WithArgs(10, 0).
					WillReturnRows(rows)
			},
			skip:           0,
			take:           10,
			expectedResult: []dto.ResponseSongDTO{},
			expectedError:  sql.ErrNoRows,
		},
		{
			name: "Database query error",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT s.id AS song_id, s.title, g.name AS group_name").
					WithArgs(10, 0).
					WillReturnError(errors.New("database error"))
			},
			skip:           0,
			take:           10,
			expectedResult: nil,
			expectedError:  fmt.Errorf("failed to fetch songs: %w", errors.New("database error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(mock)

			result, err := repo.Get(context.Background(), tt.skip, tt.take)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error(),
					"expected error containing: %v, got: %v", tt.expectedError.Error(), err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func toPtr(value string) *string {
	return &value
}

func timePtr(t time.Time) *time.Time {
	return &t
}

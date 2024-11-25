package song

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/acronix0/song-libary-api/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestSongRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening a stub database connection: %s", err)
	}
	defer db.Close()

	repo := NewSongRepository(db)

	tests := []struct {
		name           string
		mockBehavior   func(mock sqlmock.Sqlmock)
		input          dto.UpdateSongDTO
		expectedResult dto.ResponseSongDTO
		expectedError  error
	}{
		{
			name: "Valid update with all fields (excluding lyrics)",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				fixedTime := time.Date(2024, time.November, 26, 1, 19, 37, 83834300, time.Local)

				mock.ExpectQuery("SELECT id FROM groups WHERE name = \\$1").
					WithArgs("New Group").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectQuery("UPDATE songs").
					WithArgs("New Song", 1, "http://example.com", time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), sqlmock.AnyArg(), 1).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "title", "group_name", "link", "release_date", "created_at", "updated_at",
					}).AddRow(
						1, "New Song", "New Group", "http://example.com", time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), fixedTime, fixedTime,
					))
			},
			input: dto.UpdateSongDTO{
				SongID:      1,
				Song:        toPtr("New Song"),
				Group:       toPtr("New Group"),
				Link:        toPtr("http://example.com"),
				ReleaseDate: timePtr(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)),
			},
			expectedResult: dto.ResponseSongDTO{
				SongID:      1,
				Song:        "New Song",
				Group:       "New Group",
				Link:        toPtr("http://example.com"),
				ReleaseDate: timePtr(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)),
				CreatedAt:   timePtr(time.Date(2024, time.November, 26, 1, 19, 37, 83834300, time.Local)),
				UpdatedAt:   timePtr(time.Date(2024, time.November, 26, 1, 19, 37, 83834300, time.Local)),
			},
			expectedError: nil,
		},

		{
			name: "Group not found",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT id FROM groups WHERE name = \\$1").
					WithArgs("Nonexistent Group").
					WillReturnError(sql.ErrNoRows)
			},
			input: dto.UpdateSongDTO{
				SongID: 1,
				Group:  toPtr("Nonexistent Group"),
			},
			expectedResult: dto.ResponseSongDTO{},
			expectedError:  fmt.Errorf("group not found: Nonexistent Group"),
		},
		{
			name:         "No fields to update",
			mockBehavior: func(mock sqlmock.Sqlmock) {},
			input: dto.UpdateSongDTO{
				SongID: 1,
			},
			expectedResult: dto.ResponseSongDTO{},
			expectedError:  fmt.Errorf("no fields to update"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(mock)

			result, err := repo.Update(context.Background(), tt.input)

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

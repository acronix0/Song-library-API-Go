package v1

import (
	"bytes"
	"errors"
	"log/slog"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/acronix0/song-libary-api/internal/dto"
	mock_service "github.com/acronix0/song-libary-api/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_updateSong(t *testing.T) {
	tests := []struct {
		name                 string
		songID               string
		inputBody            string
		inputSong            dto.UpdateSongDTO
		mockBehavior         func(r *mock_service.MockLibrary, input dto.UpdateSongDTO)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:   "Update Song: valid request",
			songID: "1",
			inputBody: `{
        "song": "Updated Song",
        "group": "Updated Group"
    }`,
			inputSong: dto.UpdateSongDTO{
				SongID: 1,
				Song:   ptr("Updated Song"),
				Group:  ptr("Updated Group"),
			},
			mockBehavior: func(r *mock_service.MockLibrary, input dto.UpdateSongDTO) {
				r.EXPECT().
					Update(gomock.Any(), input).
					Return(dto.ResponseSongDTO{
						SongID: 1,
						Song:   "Updated Song",
						Group:  "Updated Group",
					}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
        "song_id": 1,
        "song": "Updated Song",
        "group": "Updated Group"
    }`,
		},

		{
			name:      "Update Song: invalid input",
			songID:    "1",
			inputBody: `{"invalid_field": "Test"}`, // Неверное поле
			inputSong: dto.UpdateSongDTO{},         // Пустая структура
			mockBehavior: func(r *mock_service.MockLibrary, input dto.UpdateSongDTO) {
				// not called
				r.EXPECT().
					Update(gomock.Any(), gomock.Eq(input)).
					Times(0)
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{
        "message": "Invalid input"
    }`,
		},
		{
			name:   "Update Song: service error",
			songID: "1",
			inputBody: `{
				"song": "Updated Song",
				"group": "Updated Group"
			}`,
			inputSong: dto.UpdateSongDTO{
				SongID: 1,
				Song:   ptr("Updated Song"),
				Group:  ptr("Updated Group"),
			},
			mockBehavior: func(r *mock_service.MockLibrary, input dto.UpdateSongDTO) {
				r.EXPECT().
					Update(gomock.Any(), input).
					Return(dto.ResponseSongDTO{}, errors.New("service error"))
			},
			expectedStatusCode: 500,
			expectedResponseBody: `{
				"message": "Failed to update song"
			}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLibrary := mock_service.NewMockLibrary(ctrl)
			mockService := mock_service.NewMockServiceManager(ctrl)
			mockService.EXPECT().Library().Return(mockLibrary).AnyTimes()

			test.mockBehavior(mockLibrary, test.inputSong)

			log := slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
			)
			handler := NewV1Handler(mockService, log)

			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.PUT("/songs/:id", handler.updateSong)

			req := httptest.NewRequest("PUT", "/songs/"+test.songID, bytes.NewBufferString(test.inputBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.JSONEq(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}

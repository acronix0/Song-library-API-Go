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

func TestHandler_createSong(t *testing.T) {
	tests := []struct {
		name                 string
		inputBody            string
		inputSong            dto.CreateSongDTO
		mockBehavior         func(r *mock_service.MockLibrary, input dto.CreateSongDTO)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Create Song: valid request",
			inputBody: `{
				"song": "Test Song",
				"group": "Test Group"
			}`,
			inputSong: dto.CreateSongDTO{
				Song:  "Test Song",
				Group: "Test Group",
			},
			mockBehavior: func(r *mock_service.MockLibrary, input dto.CreateSongDTO) {
				r.EXPECT().
					CreateSong(gomock.Any(), input).
					Return(1, nil)
			},
			expectedStatusCode: 201,
			expectedResponseBody: `{
				"song_id": 1,
				"song": "Test Song",
				"group": "Test Group"
			}`,
		},
		{
			name: "Create Song: invalid input",
			inputBody: `{
		"invalid_field": "Test"
	}`,
			inputSong: dto.CreateSongDTO{},
			mockBehavior: func(r *mock_service.MockLibrary, input dto.CreateSongDTO) {
				//not be called
				r.EXPECT().
					CreateSong(gomock.Any(), gomock.Any()).Times(0)
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{
		"message": "Invalid input data"
	}`,
		},
		{
			name:      "Create Song: empty input",
			inputBody: `{}`,
			inputSong: dto.CreateSongDTO{},
			mockBehavior: func(r *mock_service.MockLibrary, input dto.CreateSongDTO) {
				//not be called
				r.EXPECT().
					CreateSong(gomock.Any(), gomock.Any()).Times(0)
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{
		"message": "Invalid input data"
	}`,
		},
		{
			name: "Create Song: only group provided",
			inputBody: `{
		"group": "Test Group"
	}`,
			inputSong: dto.CreateSongDTO{},
			mockBehavior: func(r *mock_service.MockLibrary, input dto.CreateSongDTO) {
				//not be called
				r.EXPECT().
					CreateSong(gomock.Any(), gomock.Any()).Times(0)
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{
		"message": "Invalid input data"
	}`,
		},
		{
			name: "Create Song: service error",
			inputBody: `{
		"song": "Test Song",
		"group": "Test Group"
	}`,
			inputSong: dto.CreateSongDTO{
				Song:  "Test Song",
				Group: "Test Group",
			},
			mockBehavior: func(r *mock_service.MockLibrary, input dto.CreateSongDTO) {
				r.EXPECT().
					CreateSong(gomock.Any(), gomock.Eq(input)).
					Return(0, errors.New("internal service error"))
			},
			expectedStatusCode: 500,
			expectedResponseBody: `{
		"message": "Failed to add song"
	}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("Starting test:", test.name)

			t.Log("Input Body:", test.inputBody)
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLibrary := mock_service.NewMockLibrary(ctrl)
			test.mockBehavior(mockLibrary, test.inputSong)

			mockService := mock_service.NewMockServiceManager(ctrl)
			mockService.EXPECT().Library().Return(mockLibrary).AnyTimes()
			log := slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
			)
			handler := NewV1Handler(mockService, log)

			gin.SetMode(gin.TestMode)
			r := gin.New()

			r.POST("/songs", handler.createSong)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/songs", bytes.NewBufferString(test.inputBody))
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)

			t.Log("Response Code:", w.Code)
			t.Log("Response Body:", w.Body.String())

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.JSONEq(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

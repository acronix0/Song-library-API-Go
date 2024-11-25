package v1

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http/httptest"
	"os"
	"testing"

	mock_service "github.com/acronix0/song-libary-api/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_deleteSong(t *testing.T) {
	tests := []struct {
		name                 string
		url                  string
		songId               int
		mockBehavior         func(r *mock_service.MockLibrary, id int)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:   "Delete Song: valid request",
			url:    "/songs/1",
			songId: 1,
			mockBehavior: func(r *mock_service.MockLibrary, id int) {
				r.EXPECT().
					Delete(gomock.Any(), id).
					Return(nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
    "message": "Song deleted successfully"
}`,
		},
		{
			name:   "Delete Song: invalid input",
			url:    "/songs/asd",
			songId: 0,
			mockBehavior: func(r *mock_service.MockLibrary, id int) {
				//not be called
				r.EXPECT().
					Delete(gomock.Any(), gomock.Any()).Times(0)
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{
		"message": "Invalid song_id parameter"
	}`,
		},
		{
			name:   "Delete Song: song not found",
			url:    "/songs/9999",
			songId: 9999,
			mockBehavior: func(r *mock_service.MockLibrary, id int) {
				r.EXPECT().
					Delete(gomock.Any(), id).
					Return(sql.ErrNoRows)
			},
			expectedStatusCode: 404,
			expectedResponseBody: `{
		"message": "Song not found"
	}`,
		},
		{
			name:   "Delete Song: service error",
			url:    "/songs/1",
			songId: 1,
			mockBehavior: func(r *mock_service.MockLibrary, id int) {
				r.EXPECT().
					Delete(gomock.Any(), id).
					Return(errors.New("internal service error"))
			},
			expectedStatusCode: 500,
			expectedResponseBody: `{
		"message": "Failed to delete song"
	}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("Starting test:", test.name)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLibrary := mock_service.NewMockLibrary(ctrl)
			test.mockBehavior(mockLibrary, test.songId)

			mockService := mock_service.NewMockServiceManager(ctrl)
			mockService.EXPECT().Library().Return(mockLibrary).AnyTimes()
			log := slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
			)
			handler := NewV1Handler(mockService, log)

			gin.SetMode(gin.TestMode)
			r := gin.New()

			r.DELETE("/songs/:id", handler.deleteSong)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", test.url, nil)
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)

			t.Log("Response Code:", w.Code)
			t.Log("Response Body:", w.Body.String())

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.JSONEq(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

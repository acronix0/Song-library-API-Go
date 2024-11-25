package v1

import (
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

func TestHandler_getSongText(t *testing.T) {
	tests := []struct {
		name                 string
		queryParams          string
		songID               int
		skip                 int
		take                 int
		mockBehavior         func(r *mock_service.MockLibrary, songID, skip, take int)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Get Song Text: valid request",
			queryParams: "song_id=1&skip=0&take=10",
			songID:      1,
			skip:        0,
			take:        10,
			mockBehavior: func(r *mock_service.MockLibrary, songID, skip, take int) {
				r.EXPECT().
					GetSongText(gomock.Any(), songID, skip, take).
					Return("This is a test song text", nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
                "song_text": "This is a test song text"
            }`,
		},
		{
			name:        "Get Song Text: invalid song_id",
			queryParams: "song_id=abc&skip=0&take=10",
			songID:      0,
			skip:        0,
			take:        10,
			mockBehavior: func(r *mock_service.MockLibrary, songID, skip, take int) {
				//not called
				r.EXPECT().
					GetSongText(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{
                "message": "Invalid song_id parameter"
            }`,
		},
		{
			name:        "Get Song Text: invalid skip parameter",
			queryParams: "song_id=1&skip=-1&take=10",
			songID:      1,
			skip:        -1,
			take:        10,
			mockBehavior: func(r *mock_service.MockLibrary, songID, skip, take int) {
				//not called
				r.EXPECT().
					GetSongText(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{
                "message": "Invalid skip parameter"
            }`,
		},
		{
			name:        "Get Song Text: service error",
			queryParams: "song_id=1&skip=0&take=10",
			songID:      1,
			skip:        0,
			take:        10,
			mockBehavior: func(r *mock_service.MockLibrary, songID, skip, take int) {
				r.EXPECT().
					GetSongText(gomock.Any(), songID, skip, take).
					Return("", errors.New("Failed to fetch song text"))
			},
			expectedStatusCode: 500,
			expectedResponseBody: `{
        "message": "Failed to fetch song text"
    }`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("Starting test:", test.name)
			t.Logf("id:%d skip:%d take:%d", test.songID, test.skip, test.take)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLibrary := mock_service.NewMockLibrary(ctrl)
			mockService := mock_service.NewMockServiceManager(ctrl)
			mockService.EXPECT().Library().Return(mockLibrary).AnyTimes()

			test.mockBehavior(mockLibrary, test.songID, test.skip, test.take)

			log := slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
			)
			handler := NewV1Handler(mockService, log)

			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.GET("/songs/text", handler.getSongText)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/songs/text?"+test.queryParams, nil)

			r.ServeHTTP(w, req)

			t.Log("Response Code:", w.Code)
			t.Log("Response Body:", w.Body.String())

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.JSONEq(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

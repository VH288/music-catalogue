package tracks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VH288/music-catalogue/internal/models/trackactivities"
	"github.com/VH288/music-catalogue/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_UpsertTrackActivities(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSvc := NewMockservice(mockCtrl)

	vtrue := true

	tests := []struct {
		name               string
		expectedStatusCode int
		mockFn             func()
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().UpsertTrackActivites(gomock.Any(), uint(1), trackactivities.TrackActivityRequest{
					SpotifyID: "spotifyID",
					IsLiked:   &vtrue,
				}).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "fail",
			mockFn: func() {
				mockSvc.EXPECT().UpsertTrackActivites(gomock.Any(), uint(1), trackactivities.TrackActivityRequest{
					SpotifyID: "spotifyID",
					IsLiked:   &vtrue,
				}).Return(assert.AnError)
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}

			h.RegisterRoutes()
			w := httptest.NewRecorder()

			endpoint := `/tracks/liking`

			reqBody := trackactivities.TrackActivityRequest{
				SpotifyID: "spotifyID",
				IsLiked:   &vtrue,
			}

			reqBodyBytes, err := json.Marshal(reqBody)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, endpoint, io.NopCloser(bytes.NewBuffer(reqBodyBytes)))
			assert.NoError(t, err)

			token, err := jwt.CreateToken(uint(1), "username", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

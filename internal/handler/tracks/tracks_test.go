package tracks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VH288/music-catalogue/internal/models/spotify"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_Seacrh(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSvc := NewMockservice(mockCtrl)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedBody       spotify.SearchResponse
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: 200,
			expectedBody: spotify.SearchResponse{
				Total:  1,
				Limit:  2,
				Offset: 0,
				Items: []spotify.SpotifyTrackObject{
					{
						ID:               "7yQnZkAxSiAavC0zFRB1NI",
						Name:             "I’m invincible",
						Explicit:         false,
						AlbumName:        "UTA'S SONGS ONE PIECE FILM RED",
						AlbumType:        "album",
						AlbumTotalTracks: 8,
						AlbumImagesURL: []string{
							"https://i.scdn.co/image/ab67616d0000b2730cbecafa929898c82adc519c",
							"https://i.scdn.co/image/ab67616d00001e020cbecafa929898c82adc519c",
							"https://i.scdn.co/image/ab67616d000048510cbecafa929898c82adc519c",
						},
						ArtistsName: []string{
							"Ado",
						},
					},
				},
			},
			wantErr: false,
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "im invisible", 2, 1).Return(&spotify.SearchResponse{
					Total:  1,
					Limit:  2,
					Offset: 0,
					Items: []spotify.SpotifyTrackObject{
						{
							ID:               "7yQnZkAxSiAavC0zFRB1NI",
							Name:             "I’m invincible",
							Explicit:         false,
							AlbumName:        "UTA'S SONGS ONE PIECE FILM RED",
							AlbumType:        "album",
							AlbumTotalTracks: 8,
							AlbumImagesURL: []string{
								"https://i.scdn.co/image/ab67616d0000b2730cbecafa929898c82adc519c",
								"https://i.scdn.co/image/ab67616d00001e020cbecafa929898c82adc519c",
								"https://i.scdn.co/image/ab67616d000048510cbecafa929898c82adc519c",
							},
							ArtistsName: []string{
								"Ado",
							},
						},
					},
				}, nil)
			},
		},
		{
			name:               "failed",
			expectedStatusCode: 400,
			expectedBody:       spotify.SearchResponse{},
			wantErr:            true,
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "im invisible", 2, 1).Return(nil, assert.AnError)
			},
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

			endpoint := `/tracks/search?query=im+invisible&pagesize=2&pageindex=1`

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := spotify.SearchResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}

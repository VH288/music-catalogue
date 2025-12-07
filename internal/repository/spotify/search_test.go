package spotify

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/VH288/music-catalogue/internal/configs"
	"github.com/VH288/music-catalogue/pkg/httpclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_outbound_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockHTTPClient := httpclient.NewMockHTTPClient(mockCtrl)
	type args struct {
		query  string
		limit  int
		offset int
	}
	next := "https://api.spotify.com/v1/search?offset=2&limit=2&query=i%27m%20invisible&type=track&market=jp&locale=en-US,en;q%3D0.5"
	tests := []struct {
		name    string
		args    args
		want    *SpotifySearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:  "i'm invisible",
				limit:  2,
				offset: 0,
			},
			want: &SpotifySearchResponse{
				Tracks: SpotifyTrack{
					Href:   "https://api.spotify.com/v1/search?offset=0&limit=2&query=i%27m%20invisible&type=track&market=jp&locale=en-US,en;q%3D0.5",
					Limit:  2,
					Next:   &next,
					Offset: 0,
					Total:  1,
					Items: []SpotifyTrackObject{
						{
							ID:       "7yQnZkAxSiAavC0zFRB1NI",
							Name:     "I’m invincible",
							Href:     "https://api.spotify.com/v1/tracks/7yQnZkAxSiAavC0zFRB1NI",
							Explicit: false,
							Artists: []SpotifyArtistsObject{
								{
									Name: "Ado",
									Href: "https://api.spotify.com/v1/artists/6mEQK9m2krja6X1cfsAjfl",
								},
							},
							Album: SpotifyAlbumObject{
								Name:        "UTA'S SONGS ONE PIECE FILM RED",
								AlbumType:   "album",
								TotalTracks: 8,
								Images: []SpotifyAlbumImage{
									{
										URL: "https://i.scdn.co/image/ab67616d0000b2730cbecafa929898c82adc519c",
									},
									{
										URL: "https://i.scdn.co/image/ab67616d00001e020cbecafa929898c82adc519c",
									},
									{
										URL: "https://i.scdn.co/image/ab67616d000048510cbecafa929898c82adc519c",
									},
								},
							},
						},
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := `https://api.spotify.com/v1/search`

				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessTokken")
				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(searchResponse)),
				}, nil)
			},
		},
		{
			name: "fail",
			args: args{
				query:  "i'm invisible",
				limit:  10,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := `https://api.spotify.com/v1/search`

				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessTokken")
				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 500,
					Body:       io.NopCloser(bytes.NewBufferString(`Internal Server Error`)),
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			o := &outbound{
				cfg:         &configs.Config{},
				client:      mockHTTPClient,
				AccessToken: "accessTokken",
				TokenType:   "Bearer",
				ExpiredAt:   time.Now().Add(1 * time.Hour),
			}
			got, err := o.Search(context.Background(), tt.args.query, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Fatalf("outbound.Search() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outbound.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

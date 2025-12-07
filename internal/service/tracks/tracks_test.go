package tracks

import (
	"context"
	"reflect"
	"testing"

	"github.com/VH288/music-catalogue/internal/models/spotify"
	spotifyRepo "github.com/VH288/music-catalogue/internal/repository/spotify"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_service_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSpotifyOutbound := NewMockspotifyOutbound(mockCtrl)
	type args struct {
		query     string
		pageSize  int
		pageIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    *spotify.SearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:     "i'm invisible",
				pageSize:  2,
				pageIndex: 1,
			},
			want: &spotify.SearchResponse{
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
			mockFn: func(args args) {
				next := "https://api.spotify.com/v1/search?offset=2&limit=2&query=i%27m%20invisible&type=track&market=jp&locale=en-US,en;q%3D0.5"
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 2, 0).Return(&spotifyRepo.SpotifySearchResponse{
					Tracks: spotifyRepo.SpotifyTrack{
						Href:   "https://api.spotify.com/v1/search?offset=0&limit=2&query=i%27m%20invisible&type=track&market=jp&locale=en-US,en;q%3D0.5",
						Limit:  2,
						Next:   &next,
						Offset: 0,
						Total:  1,
						Items: []spotifyRepo.SpotifyTrackObject{
							{
								ID:       "7yQnZkAxSiAavC0zFRB1NI",
								Name:     "I’m invincible",
								Href:     "https://api.spotify.com/v1/tracks/7yQnZkAxSiAavC0zFRB1NI",
								Explicit: false,
								Artists: []spotifyRepo.SpotifyArtistsObject{
									{
										Name: "Ado",
										Href: "https://api.spotify.com/v1/artists/6mEQK9m2krja6X1cfsAjfl",
									},
								},
								Album: spotifyRepo.SpotifyAlbumObject{
									Name:        "UTA'S SONGS ONE PIECE FILM RED",
									AlbumType:   "album",
									TotalTracks: 8,
									Images: []spotifyRepo.SpotifyAlbumImage{
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
				}, nil)
			},
		},
		{
			name: "fail",
			args: args{
				query:     "i'm invisible",
				pageSize:  2,
				pageIndex: 1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 2, 0).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				spotifyOutbound: mockSpotifyOutbound,
			}
			got, err := s.Search(context.Background(), tt.args.query, tt.args.pageSize, tt.args.pageIndex)
			if (err != nil) != tt.wantErr {
				t.Fatalf("service.Search() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

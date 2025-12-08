package tracks

import (
	"context"

	"github.com/VH288/music-catalogue/internal/models/trackactivities"
	"github.com/VH288/music-catalogue/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type spotifyOutbound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type trackActivitiesRepo interface {
	Create(ctx context.Context, model trackactivities.TrackActivity) error
	Update(ctx context.Context, model trackactivities.TrackActivity) error
	Get(ctx context.Context, userID uint, spotifyID string) (*trackactivities.TrackActivity, error)
	GetBulkSpotifyIDS(ctx context.Context, userID uint, spotifyIDs []string) (map[string]trackactivities.TrackActivity, error)
}

type service struct {
	spotifyOutbound     spotifyOutbound
	trackActivitiesRepo trackActivitiesRepo
}

func NewService(spotifyOutbound spotifyOutbound, trackActivitiesRepo trackActivitiesRepo) *service {
	return &service{spotifyOutbound: spotifyOutbound, trackActivitiesRepo: trackActivitiesRepo}
}

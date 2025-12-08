package tracks

import (
	"context"

	"github.com/VH288/music-catalogue/internal/models/spotify"
	"github.com/VH288/music-catalogue/internal/models/trackactivities"
	spotifyrepo "github.com/VH288/music-catalogue/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error search track to spotify")
		return nil, err
	}

	trackIDs := make([]string, len(trackDetails.Tracks.Items))
	for idx, item := range trackDetails.Tracks.Items {
		trackIDs[idx] = item.ID
	}

	trackActivities, err := s.trackActivitiesRepo.GetBulkSpotifyIDS(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get track activities from database")
		return nil, err
	}

	return modelToResponse(trackDetails, trackActivities), nil
}

func modelToResponse(data *spotifyrepo.SpotifySearchResponse, mapTrackActivities map[string]trackactivities.TrackActivity) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks.Items {
		if item.Name != "" {
			artistsName := make([]string, len(item.Artists))

			for idx, artist := range item.Artists {
				artistsName[idx] = artist.Name
			}

			imageURL := make([]string, len(item.Album.Images))

			for idx, image := range item.Album.Images {
				imageURL[idx] = image.URL
			}

			items = append(items, spotify.SpotifyTrackObject{
				ID:       item.ID,
				Name:     item.Name,
				Explicit: item.Explicit,
				IsLiked:  mapTrackActivities[item.ID].IsLiked,

				ArtistsName: artistsName,

				AlbumName:        item.Album.Name,
				AlbumType:        item.Album.AlbumType,
				AlbumTotalTracks: item.Album.TotalTracks,
				AlbumImagesURL:   imageURL,
			})
		}
	}
	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Total:  data.Tracks.Total,
		Items:  items,
	}
}

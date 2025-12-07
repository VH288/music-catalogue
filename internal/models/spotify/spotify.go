package spotify

type SearchResponse struct {
	Items  []SpotifyTrackObject `json:"items"`
	Total  int                  `json:"total"`
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
}

type SpotifyTrackObject struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Explicit bool   `json:"explicit"`
	// album fields
	AlbumName        string   `json:"album_name"`
	AlbumType        string   `json:"album_type"`
	AlbumTotalTracks int      `json:"album_total_tracks"`
	AlbumImagesURL   []string `json:"album_image_url"`
	// artist fields
	ArtistsName []string `json:"artists_name"`
}

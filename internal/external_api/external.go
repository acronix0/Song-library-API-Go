package externalapi

import "context"

type ExternalAPIClient interface {
	FetchSongDetails(ctx context.Context, group, song string) (*SongDetail, error)
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

package dto

import (
	"errors"
	"time"
)

type UpdateSongDTO struct {
	SongID      int        `json:"song_id"`
	Song        *string    `json:"song,omitempty"`
	Group       *string    `json:"group,omitempty"`
	Link        *string    `json:"link,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Text        *string    `json:"lyrics,omitempty"`
}
type CreateSongDTO struct {
	SongID      int        `json:"song_id"`
	Song        string     `json:"song" binding:"required"`
	Group       string     `json:"group" binding:"required"`
	Link        *string    `json:"link,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Text        *string    `json:"lyrics,omitempty"`
}

type ResponseSongDTO struct {
	SongID      int        `json:"song_id" binding:"required"`
	Song        string     `json:"song,omitempty" binding:"required"`
	Group       string     `json:"group,omitempty" binding:"required"`
	Link        *string    `json:"link,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Text        *string    `json:"lyrics,omitempty"`
}
type LiricsDTO struct {
	SongID int
	Text   string
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (u UpdateSongDTO) Validate() error {
	if u.Song == nil && u.Group == nil && u.Link == nil &&
		u.ReleaseDate == nil && u.Text == nil {
		return errors.New("no fields to update")
	}
	return nil
}

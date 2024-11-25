package repository

import (
	"context"
	"database/sql"

	"github.com/acronix0/song-libary-api/internal/dto"
	"github.com/acronix0/song-libary-api/internal/repository/lyrics"
	"github.com/acronix0/song-libary-api/internal/repository/song"
)

type Song interface {
	Create(ctx context.Context, song dto.CreateSongDTO) (int, error)
	Get(ctx context.Context, skip, take int) ([]dto.ResponseSongDTO, error)
	Update(ctx context.Context, song dto.UpdateSongDTO) (dto.ResponseSongDTO, error)
	Delete(ctx context.Context, songID int) error
}
type Lyrics interface {
	Create(ctx context.Context, songId int, text string) error
	Get(ctx context.Context, songId, skip, take int) (string, error)
	Update(ctx context.Context, songId int, text string) error
	Delete(ctx context.Context, songID int) error
}
type RepositoryManager interface {
	Song() Song
	Lyrics() Lyrics
}

var _ Song = (*song.SongRepo)(nil)
var _ Lyrics = (*lyrics.LyricsRepo)(nil)

type repositories struct {
	db     *sql.DB
	song   Song
	lyrics Lyrics
}

func NewRepositoryManager(db *sql.DB) *repositories {
	return &repositories{db: db}
}
func (r *repositories) DB() *sql.DB {
	return r.db
}
func (r *repositories) Song() Song {
	if r.song == nil {
		r.song = song.NewSongRepository(r.DB())
	}
	return r.song
}

func (r *repositories) Lyrics() Lyrics {
	if r.lyrics == nil {
		r.lyrics = lyrics.NewLyricsRepository(r.DB())
	}
	return r.lyrics
}

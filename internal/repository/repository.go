package repository

import (
	"context"
	"database/sql"

	"github.com/acronix0/song-libary-api/internal/dto"
	"github.com/acronix0/song-libary-api/internal/repository/liryc"
	"github.com/acronix0/song-libary-api/internal/repository/song"
)





type Song interface{
	Create(ctx context.Context, groupName, songName string) (int, error)
	Get(ctx context.Context, skip, take int) ([]dto.SongDTO, error)
	Update(ctx context.Context,song dto.SongDTO) (error)
	Delete(ctx context.Context, songID int) (error)
}
type Liryc interface{
	Create(ctx context.Context, songId int, text string) (error)
	Get(ctx context.Context,songId, skip, take int) (string, error)
	Update(ctx context.Context, songId int, text string) (error)
	Delete(ctx context.Context, songID int) (error)
}
type RepositoryManager interface {
	Song() Song
	Liryc() Liryc
}


var _ Song = (*song.SongRepo)(nil) 
//var _ Liryc = (*lir)


type repositories struct {
	db *sql.DB
	song Song
	liryc Liryc
}

func NewRepositoryManager(db *sql.DB) *repositories {
  return &repositories{db: db}
}
func (r *repositories) DB() *sql.DB{
	return r.db
}
func (r *repositories)Song() Song{
	if r.song == nil {
		r.song = song.NewSongRepository(r.DB())
	}
	return r.song
}

func (r *repositories) Liryc() Liryc{
	if r.liryc == nil {
    r.liryc = liryc.NewLyricsRepository(r.DB())
  }
  return r.liryc
}

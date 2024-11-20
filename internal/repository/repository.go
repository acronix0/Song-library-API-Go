package repository

import (
	"context"

	"github.com/acronix0/song-libary-api/internal/dto"
)





type Song interface{
	Create(ctx context.Context, groupName, songName string) (error, int)
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
}


type repositories struct {
	song Song
}

func (r *repositories)Song() Song{
	if r.song == nil {
		r.song = 
	}
}

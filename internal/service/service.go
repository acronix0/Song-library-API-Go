package service

import (
	"context"

	"github.com/acronix0/song-libary-api/internal/dto"
)

type ServiceManager interface {
	Library() Library
}

type Library interface {
	GetSongs(ctx context.Context, skip, take int) ([]dto.ResponseSongDTO, error)
	CreateSong(ctx context.Context, song dto.CreateSongDTO) (int, error)
	GetSongText(ctx context.Context, songId, skip, take int) (string, error)
	Update(ctx context.Context, song dto.UpdateSongDTO) (dto.ResponseSongDTO, error)
	Delete(ctx context.Context, songID int) error
}

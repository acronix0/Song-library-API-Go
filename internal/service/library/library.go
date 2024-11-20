package library

import (
	"context"

	"github.com/acronix0/song-libary-api/internal/dto"
	"github.com/acronix0/song-libary-api/internal/repository"
)

type LibraryService struct {
	repo repository.Song
}

func NewLibraryService(repo repository.Song) *LibraryService {
  return &LibraryService{repo: repo}
}
func (s *LibraryService)GetAllSongs(ctx context.Context, skip, take int) ([]dto.SongDTO, error) {
	return s.Repo().
}
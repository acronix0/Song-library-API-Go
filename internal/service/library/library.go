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
func (s *LibraryService)GetSongs(ctx context.Context, skip, take int) ([]dto.SongDTO, error) {
	return s.repo.Get(ctx, skip, take)
}
func (s *LibraryService) CreateSong(ctx context.Context, groupName, songName string) (int, error){
	return s.repo.Create(ctx, groupName, songName)
}

func (s *LibraryService) Update(ctx context.Context, song dto.SongDTO) (error){
  return s.repo.Update(ctx, song)
}

func (s *LibraryService) Delete(ctx context.Context, songID int) (error){
  return s.repo.Delete(ctx, songID)
}


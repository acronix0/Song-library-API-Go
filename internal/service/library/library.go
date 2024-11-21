package library

import (
	"context"

	"github.com/acronix0/song-libary-api/internal/dto"
	"github.com/acronix0/song-libary-api/internal/repository"
)

type LibraryService struct {
	SongRepo repository.Song
	LyricsRepo repository.Lyrics
}

func NewLibraryService(songRepo repository.Song, lyricsRepo repository.Lyrics) *LibraryService {
	return &LibraryService{SongRepo: songRepo, LyricsRepo: lyricsRepo}
}
func (s *LibraryService)GetSongs(ctx context.Context, skip, take int) ([]dto.SongDTO, error) {
	return s.SongRepo.Get(ctx, skip, take)
}
func (s *LibraryService) CreateSong(ctx context.Context,song dto.SongDTO) (int, error){
	songID, err := s.SongRepo.Create(ctx, song)
	if err!= nil {
    return 0, err
  }

	if song.Text != nil {
		err = s.LyricsRepo.Create(ctx, songID, *song.Text)
		if err != nil {
			return 0, nil
		}
	}
	
	return songID, nil
}

func (s *LibraryService)GetSongText(ctx context.Context,songId, skip, take int) (string, error) {
	return s.LyricsRepo.Get(ctx, songId, skip, take)
}

func (s *LibraryService) Update(ctx context.Context, song dto.SongDTO) (dto.SongDTO, error){
	if song.Text != nil{
		err:= s.LyricsRepo.Update(ctx, song.SongID, *song.Text)
		if err == nil {
			return dto.SongDTO{}, err
		}
	}
	
  return s.SongRepo.Update(ctx, song)
}

func (s *LibraryService) Delete(ctx context.Context, songID int) (error){
	err := s.LyricsRepo.Delete(ctx, songID)
	if err != nil {
		return err
	}
	err = s.SongRepo.Delete(ctx, songID)
	if err != nil {
		return err
	}
  return nil
}


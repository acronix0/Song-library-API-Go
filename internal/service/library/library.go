package library

import (
	"context"
	"fmt"
	"time"

	"github.com/acronix0/song-libary-api/internal/dto"
	externalapi "github.com/acronix0/song-libary-api/internal/external_api"
	"github.com/acronix0/song-libary-api/internal/repository"
)

type LibraryService struct {
	SongRepo    repository.Song
	LyricsRepo  repository.Lyrics
	ExternalAPI externalapi.ExternalAPIClient
}

func NewLibraryService(songRepo repository.Song, lyricsRepo repository.Lyrics, externalApi externalapi.ExternalAPIClient) *LibraryService {
	return &LibraryService{SongRepo: songRepo, LyricsRepo: lyricsRepo, ExternalAPI: externalApi}
}
func (s *LibraryService) GetSongs(ctx context.Context, skip, take int) ([]dto.ResponseSongDTO, error) {
	return s.SongRepo.Get(ctx, skip, take)
}
func (s *LibraryService) CreateSong(ctx context.Context, song dto.CreateSongDTO) (int, error) {
	songID, err := s.SongRepo.Create(ctx, song)
	if err != nil {
		return 0, err
	}

	if song.Text != nil {
		err = s.LyricsRepo.Create(ctx, songID, *song.Text)
		if err != nil {
			return 0, fmt.Errorf("failed to create lyrics: %w", err)
		}
	} else {
		apiResponse, err := s.ExternalAPI.FetchSongDetails(ctx, song.Group, song.Song)
		if err != nil {
			return 0, fmt.Errorf("failed to fetch song details from external API: %w", err)
		}

		if apiResponse.Text != "" {
			err = s.LyricsRepo.Create(ctx, songID, apiResponse.Text)
			if err != nil {
				return 0, fmt.Errorf("failed to create lyrics from external API: %w", err)
			}
		}

		updateData := dto.UpdateSongDTO{
			SongID:      songID,
			Link:        &apiResponse.Link,
			ReleaseDate: parseTimePtr(apiResponse.ReleaseDate),
		}
		_, err = s.SongRepo.Update(ctx, updateData)
		if err != nil {
			return 0, fmt.Errorf("failed to update song details from external API: %w", err)
		}
	}

	return songID, nil
}
func parseTimePtr(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}
	parsedTime, err := time.Parse("02.01.2006", dateStr) // Формат, указанный в Swagger
	if err != nil {
		return nil
	}
	return &parsedTime
}

func (s *LibraryService) GetSongText(ctx context.Context, songId, skip, take int) (string, error) {
	return s.LyricsRepo.Get(ctx, songId, skip, take)
}

func (s *LibraryService) Update(ctx context.Context, song dto.UpdateSongDTO) (dto.ResponseSongDTO, error) {
	if song.Text != nil {
		err := s.LyricsRepo.Update(ctx, song.SongID, *song.Text)
		if err == nil {
			return dto.ResponseSongDTO{}, err
		}
	}

	return s.SongRepo.Update(ctx, song)
}

func (s *LibraryService) Delete(ctx context.Context, songID int) error {
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

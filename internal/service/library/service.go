package library

import (
	externalapi "github.com/acronix0/song-libary-api/internal/external_api"
	"github.com/acronix0/song-libary-api/internal/repository"
	"github.com/acronix0/song-libary-api/internal/service"
)

type services struct {
	LibraryService service.Library
}

func NewService(
	repoManager repository.RepositoryManager, externalApi externalapi.ExternalAPIClient,
) *services {
	return &services{
		LibraryService: NewLibraryService(
			repoManager.Song(),
			repoManager.Lyrics(),
			externalApi,
		),
	}
}

func (s *services) Library() service.Library {
	return s.LibraryService
}

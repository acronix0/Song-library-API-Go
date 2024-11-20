package library

import (
	"context"

	"github.com/acronix0/song-libary-api/internal/repository"
	"github.com/acronix0/song-libary-api/internal/service"
	def "github.com/acronix0/song-libary-api/internal/service"
)

var _ def.Library = (*services)(nil)

type services struct {
	LibraryService service.Library
}


func NewService(
	repoManager repository.RepositoryManager,
) *services {
	return &services{
		LibraryService: NewLibraryService(repoManager.Song()),
	}
}
func (s *services) Repo() repository.Song
func (s *services) Library() service.Library {
	return s.LibraryService
}
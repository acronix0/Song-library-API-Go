package library

import (

	"github.com/acronix0/song-libary-api/internal/repository"
	"github.com/acronix0/song-libary-api/internal/service"
)



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
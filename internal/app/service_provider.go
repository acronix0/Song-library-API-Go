package app

import (
	"context"
	"github.com/acronix0/song-libary-api/internal/config"
	"github.com/acronix0/song-libary-api/internal/database"
)

type serviceProvider struct {
	config   *config.Config
	database *database.Database
}
func newServiceProvider(_ context.Context) *serviceProvider {
	return &serviceProvider{}
}
func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(),
		)
	}

	return s.userService
}
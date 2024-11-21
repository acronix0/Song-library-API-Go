package app

import (
	"log/slog"
	"os"

	"github.com/acronix0/song-libary-api/internal/config"
	"github.com/acronix0/song-libary-api/internal/database"
	"github.com/acronix0/song-libary-api/internal/repository"
	"github.com/acronix0/song-libary-api/internal/service"
	"github.com/acronix0/song-libary-api/internal/service/library"
)

type serviceProvider struct {
	config   *config.Config
	dataBase *database.Database
	logger    *slog.Logger
	serviceManager service.ServiceManager
	repositoryManager repository.RepositoryManager
}
func newServiceProvider(cfg *config.Config, logger *slog.Logger) *serviceProvider {
	return &serviceProvider{config: cfg, logger: logger}
}

func (s *serviceProvider) Config() *config.Config{
	if s.config == nil {
    config, err := config.Load() 
		if err!= nil {
      panic(err)
    }
		s.config = config
	}
	return s.config
}
func (s *serviceProvider) Database() *database.Database{
	if s.dataBase == nil {
		cfg := s.Config()
		db, err := database.NewDatabase(
			cfg.DatabaseConndection.Port, 
			cfg.DatabaseConndection.Host, 
			cfg.DatabaseConndection.User, 
			cfg.DatabaseConndection.Password, 
			cfg.DatabaseConndection.Name)
  	if err!= nil {
    	panic(err)
  	}
		s.dataBase = db
	}

	return s.dataBase
}
func (s *serviceProvider) Logger() *slog.Logger{
	if s.logger == nil {
    s.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
  }
  return s.logger
}
func (s *serviceProvider) ServiceManager() service.ServiceManager {
	if s.serviceManager == nil {
		s.serviceManager = library.NewService(s.RepositoryManager())
	
	}

	return s.serviceManager
}

func (s *serviceProvider) RepositoryManager() repository.RepositoryManager{
	if s.repositoryManager == nil {
    s.repositoryManager = repository.NewRepositoryManager(s.Database().GetDB())
  }
  return s.repositoryManager
}
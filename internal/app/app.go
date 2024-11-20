package app

import (
	"context"
	"log/slog"
	"github.com/acronix0/song-libary-api/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
	logger             *slog.Logger
	config *config.Config
}

func NewApp(ctx context.Context) (*App, error) {
  a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}
func (a *App) Run() error {
	return a.runGRPCServer()
}
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	cfg, err := config.MustLoad()
	if err != nil { 
		a.logger.Error(err.Error())
	}
	a.config = cfg
	return nil
}
func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider(ctx)
	return nil
}
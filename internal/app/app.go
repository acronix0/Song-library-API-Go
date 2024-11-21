package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/acronix0/song-libary-api/internal/config"
	"github.com/acronix0/song-libary-api/internal/database"
	delivery "github.com/acronix0/song-libary-api/internal/delivery/http"
	"github.com/acronix0/song-libary-api/internal/server"
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
// @title Song Library API
// @version 1.0
// @description API for managing songs and their texts.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host localhost:8082
// @BasePath /api/v1
func (a *App) Run() {
	handlers := delivery.NewHandler(a.serviceProvider.ServiceManager(), a.logger)
	srv := server.NewServer(a.config, handlers.Init(a.config))
		go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("error occurred while running http server: %s\n", err.Error())
		}
	}()

	a.logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		a.logger.Error("failed to stop server: %v", err)
	}
}
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServiceProvider,
		a.InitMigrations,
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
	cfg, err := config.Load()
	if err != nil { 
		a.logger.Error(err.Error())
	}
	
	a.config = cfg
	return nil
}

func (a *App)initLogger(_ context.Context) error {
	var log *slog.Logger
	switch a.config.AppEnv {
	case config.EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	 a.logger = log
	 return nil
}
func (a *App) InitMigrations(_ context.Context) error {
	return database.InitMigrations(
		a.config.MigrationsPath, 
		a.config.DatabaseConndection.Host, 
		a.config.DatabaseConndection.User, 
		a.config.DatabaseConndection.Name, 
		a.config.DatabaseConndection.Password, 
		a.config.DatabaseConndection.Port,
	)
}
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.config, a.logger)
	return nil
}
package http

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/acronix0/song-libary-api/docs"
	"github.com/acronix0/song-libary-api/internal/config"
	v1 "github.com/acronix0/song-libary-api/internal/delivery/http/v1"
	"github.com/acronix0/song-libary-api/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services service.ServiceManager
	logger   *slog.Logger
}

func NewHandler(service service.ServiceManager, logger *slog.Logger) *Handler {
	return &Handler{
		services: service,
		logger:   logger,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTPConfig.Host, cfg.HTTPConfig.Port)
	if cfg.AppEnv != config.EnvLocal {
		docs.SwaggerInfo.Host = cfg.HTTPConfig.Host
	}

	if cfg.AppEnv != config.EnvProd {
		router.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewV1Handler(h.services, h.logger)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}

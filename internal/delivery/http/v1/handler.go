package v1

import (
	"log/slog"

	"github.com/acronix0/song-libary-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.ServiceManager
	logger   *slog.Logger
}

func NewV1Handler(services service.ServiceManager, logger *slog.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initSongsRoutes(v1)
	}

}

func (h *Handler) initSongsRoutes(api *gin.RouterGroup) {
	songsGroup := api.Group("/songs")
	{
		songsGroup.GET("/", h.getSongs)
		songsGroup.PUT("/:id", h.updateSong)
		songsGroup.POST("/", h.createSong)
		songsGroup.GET("/text", h.getSongText)
		songsGroup.DELETE("/", h.deleteSong)
	}
}

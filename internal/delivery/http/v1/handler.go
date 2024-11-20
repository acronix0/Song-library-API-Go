package v1

import (
	"github.com/acronix0/song-libary-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     service.ServiceManager
}

func NewHandler(services service.ServiceManager) *Handler {
	return &Handler{
		services:     services,
	}
}
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initSongsRoutes(v1)

	}
}

func (h *Handler) initSongsRoutes(api *gin.RouterGroup){
	songsGroup := api.Group("/songs")
  {
      songsGroup.GET("/", h.getSongs)
  }
}
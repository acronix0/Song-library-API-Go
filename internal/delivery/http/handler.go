package http

import (
	"net/http"

	"github.com/acronix0/song-libary-api/internal/config"
	"github.com/acronix0/song-libary-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     service.ServiceManager
}

func NewHandler(service service.ServiceManager) *Handler {
	return &Handler{services: service}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

/* 	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTPConfig.Host, cfg.HTTPConfig.Port)
	if cfg.Env != config.EnvLocal {
		docs.SwaggerInfo.Host = cfg.HTTPConfig.Host
	}
 */
/* 	if cfg.Env != config.EnvProd {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	} */

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
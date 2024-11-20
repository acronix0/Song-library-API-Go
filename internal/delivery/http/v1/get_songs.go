package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get songs
// @Tags songs
// @Description Get songs with filtering and pagination
// @Produce json
// @Param group query string false "Group name"
// @Param song query string false "Song name"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} []domain.Song
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Router /songs [get]
func (h *Handler) getSongs(c *gin.Context) {
	// Пример извлечения параметров из запроса
	group := c.Query("group")
	song := c.Query("song")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	// Вызов бизнес-логики
	songs, err := h.services.Library().GetSongs(c.Request.Context(), skip, take)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Failed to get songs")
		return
	}

	c.JSON(http.StatusOK, songs)
}

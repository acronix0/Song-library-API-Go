package v1

import (
	"net/http"

	"github.com/acronix0/song-libary-api/internal/dto"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Update song
// @Tags songs
// @Description Update an existing song
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body dto.SongDTO true "Updated song data"
// @Success 200 {object} dto.SongDTO
// @Failure 400 {object} Response "Invalid input"
// @Failure 500 {object} Response "Internal server error"
// @Router /songs/{id} [put]
func (h *Handler) updateSong(c *gin.Context) {
	var request dto.SongDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		newResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	updatedSong, err := h.services.Library().Update(
		c.Request.Context(),
		request,
	)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Failed to update song")
		return
	}

	c.JSON(http.StatusOK, updatedSong)
}
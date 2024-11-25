package v1

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/acronix0/song-libary-api/internal/dto"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Update song
// @Tags Songs
// @Description Update an existing song
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body dto.UpdateSongDTO true "Updated song data"
// @Success 200 {object} dto.ResponseSongDTO
// @Failure 400 {object} Response "Invalid input"
// @Failure 500 {object} Response "Internal server error"
// @Router /songs/{id} [put]
func (h *Handler) updateSong(c *gin.Context) {
	const op = "handller.v1.updateSong"
	logger := h.logger.With(
		slog.String("op", op),
	)
	var request dto.UpdateSongDTO

	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil || songID <= 0 {
		newResponse(c, http.StatusBadRequest, "Invalid song_id parameter")
		logger.Error("Failed to update songs")
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		newResponse(c, http.StatusBadRequest, "Invalid input")
		logger.Error("Failed to update songs: Invalid input")
		return
	}
	if err := request.Validate(); err != nil {
		newResponse(c, http.StatusBadRequest, "Invalid input")
		logger.Error("Failed to update songs: Invalid input")
		return
	}
	request.SongID = songID

	updatedSong, err := h.services.Library().Update(c.Request.Context(), request)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Failed to update song")
		logger.Error("Failed to update songs", slog.String("error", err.Error()))
		return
	}
	logger.Debug("Song updated successfully")
	c.JSON(http.StatusOK, updatedSong)
}

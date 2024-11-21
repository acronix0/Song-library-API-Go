package v1

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// deleteSong handles the DELETE request to remove a song by its ID.
// @Summary Delete Song
// @Tags Songs
// @Description Delete a song by its ID.
// @Param song_id query int true "ID of the song to delete"
// @Produce json
// @Success 200 {object} Response "Song deleted successfully"
// @Failure 400 {object} Response "Invalid song_id parameter"
// @Failure 404 {object} Response "Song not found"
// @Failure 500 {object} Response "Internal server error"
// @Router /songs [delete]
func (h *Handler) deleteSong(c *gin.Context) {
	const op = "handler.v1.deleteSong"
	logger := h.logger.With(
		slog.String("op", op),
	)

	songID, err := strconv.Atoi(c.Query("song_id"))
	if err != nil || songID <= 0 {
		newResponse(c, http.StatusBadRequest, "Invalid song_id parameter")
		logger.Error("Invalid song_id parameter", slog.String("error", err.Error()))
		return
	}

	err = h.services.Library().Delete(c.Request.Context(), songID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			newResponse(c, http.StatusNotFound, "Song not found")
			logger.Error("Song not found", slog.String("error", err.Error()))
			return
		}

		newResponse(c, http.StatusInternalServerError, "Failed to delete song")
		logger.Error("Failed to delete song", slog.String("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully"})
}

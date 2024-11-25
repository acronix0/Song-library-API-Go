package v1

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// getSongText handles the GET request for fetching the text of a song.
// @Summary Get Song Text
// @Tags Songs
// @Description Get the text of a song by its ID.
// @Produce json
// @Param song_id query int true "ID of the song"
// @Param skip query int true "Number of verses to skip" default(0)
// @Param take query int true "Number of verses to fetch" default(10)
// @Success 200 {object} Response "string"
// @Failure 400 {object} Response "Invalid query parameters"
// @Failure 404 {object} Response "Song text not found"
// @Failure 500 {object} Response "Internal server error"
// @Router /songs/text [get]
func (h *Handler) getSongText(c *gin.Context) {
	const op = "handler.v1.getSongText"
	logger := h.logger.With(
		slog.String("op", op),
	)

	songID, err := strconv.Atoi(c.Query("song_id"))
	if err != nil || songID <= 0 {
		newResponse(c, http.StatusBadRequest, "Invalid song_id parameter")
		logger.Error("Invalid song_id parameter")
		return
	}

	skip, err := strconv.Atoi(c.DefaultQuery("skip", "0"))
	if err != nil || skip < 0 {
		newResponse(c, http.StatusBadRequest, "Invalid skip parameter")
		logger.Error("Invalid skip parameter")
		return
	}

	take, err := strconv.Atoi(c.DefaultQuery("take", "10"))
	if err != nil || take <= 0 {
		newResponse(c, http.StatusBadRequest, "Invalid take parameter")
		logger.Error("Invalid take parameter")
		return
	}

	songText, err := h.services.Library().GetSongText(c.Request.Context(), songID, skip, take)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newResponse(c, http.StatusNotFound, "Song text not found")
			logger.Error("Song text not found", slog.String("error", err.Error()))
			return
		}

		newResponse(c, http.StatusInternalServerError, "Failed to fetch song text")
		logger.Error("Failed to fetch song text", slog.String("error", err.Error()))
		return
	}
	logger.Debug("Song text fetch successfully")
	c.JSON(http.StatusOK, gin.H{"song_text": songText})
}

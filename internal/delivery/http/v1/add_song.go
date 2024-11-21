package v1

import (
	"log/slog"
	"net/http"

	"github.com/acronix0/song-libary-api/internal/dto"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Add Song
// @Tags Songs
// @Description Add a new song to the library.
// @Accept json
// @Produce json
// @Param song body dto.SongDTO true "Song object"
// @Success 201 {object} dto.SongDTO
// @Failure 400 {object} Response "Invalid input data"
// @Failure 500 {object} Response "Internal server error"
// @Router /songs [post]
func (h *Handler) createSong(c *gin.Context) {
	const op = "handler.v1.addSongs"
	logger := h.logger.With(
		slog.String("op", op),
	)

	var songDTO dto.SongDTO
	if err := c.ShouldBindJSON(&songDTO); err != nil {
		newResponse(c, http.StatusBadRequest, "Invalid input data")
		logger.Error("Invalid input data", slog.String("error", err.Error()))
		return
	}

	songID, err := h.services.Library().CreateSong(c.Request.Context(), songDTO)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Failed to add song")
		logger.Error("Failed to add song", slog.String("error", err.Error()))
		return
	}

	songDTO.SongID = songID
	c.JSON(http.StatusCreated, songDTO)
}

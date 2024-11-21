package v1

import (
	"log/slog"
	"net/http"
	"strconv"

	_"github.com/acronix0/song-libary-api/internal/dto"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// getSongs handles the GET request for fetching a list of songs.
// @Summary Get Songs
// @Tags Songs
// @Description Get a paginated list of songs.
// @Produce json
// @Param skip query int true "Number of records to skip" default(0)
// @Param take query int true "Number of records to fetch" default(10)
// @Success 200 {array} dto.SongDTO
// @Failure 400 {object} Response "Invalid query parameters"
// @Failure 500 {object} Response "Internal server error"
// @Router /songs [get]
func (h *Handler) getSongs(c *gin.Context) {
	const op = "handller.v1.getSongs"
	logger := h.logger.With(
		slog.String("op", op),
	)
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

  songs, err := h.services.Library().GetSongs(c.Request.Context(), skip, take)
  if err != nil {
      newResponse(c, http.StatusInternalServerError, "Failed to fetch songs")
			logger.Error("Failed to fetch songs")
      return
  }

  c.JSON(http.StatusOK, songs)
}

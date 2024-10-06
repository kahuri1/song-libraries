package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// @Summary Delete Song
// @Tags Song
// @Description Delete a song by ID
// @ID delete-song
// @Accept json
// @Produce json
// @Param id query int64 true "ID of the song to delete"
// @Success 200 {string} string "Song deleted successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid ID parameter"
// @Failure 404 {object} model.ErrorResponse "Group not found"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /Song [delete]
func (h *Handler) DeleteSong(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "song not found"})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}
	var id64 int64 = int64(id)
	err = h.service.DeleteGroup(id64)
	if err != nil {
		log.Printf("Failed delete song: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": model.ErrorResponse{err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

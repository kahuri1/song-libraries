package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// @Summary Get Song Info
// @Tags Info
// @Description Retrieve details of a song based on group and song name
// @ID get-song-info
// @Accept json
// @Produce json
// @Param request body model.RequestInfo true "Request Info"
// @Success 200 {object} model.SongDetailInfo "Song details"
// @Failure 400 {object} model.ErrorResponse "Bad request"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /info [post]
func (h *Handler) SongInfo(c *gin.Context) {
	var s model.SongDetailInfo
	var r model.RequestInfo

	if err := c.ShouldBindJSON(&r); err != nil {
		log.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	s, err := h.service.SongInfo(r.Group, r.Name)
	if err != nil {
		log.Printf("Failed delete song: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": model.ErrorResponse{err.Error()}})
		return
	}

	c.JSON(http.StatusOK, s)
}

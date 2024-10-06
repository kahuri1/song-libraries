package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// @Summary Update Song Details
// @Tags Song
// @Description Update the details of a song
// @ID update-song-details
// @Accept json
// @Produce json
// @Param input body model.SongDetail true "Song details"
// @Success 200 {object} string "Update confirmation"
// @Failure 400 {object} model.ErrorResponse "Invalid ID parameter"
// @Failure 404 {object} model.ErrorResponse "Group not found"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /Song/Details [put]
func (h *Handler) UpdateSongDetails(c *gin.Context) {
	var songDetails model.SongDetail
	d, err := c.GetRawData()

	err = json.Unmarshal(d, &songDetails)
	if err != nil {
		log.Errorf("unmarshal handlerError")
		return
	}
	_, err = h.service.UpdateSongDetails(songDetails)
	if err != nil {
		log.Printf("Failed to update song detail: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": fmt.Sprintf("updated")})
}

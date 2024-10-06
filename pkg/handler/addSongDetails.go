package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// @Summary AddSongDetails
// @Tags Song
// @Description create songDetails
// @ID create-song
// @Accept json
// @Produce json
// @Param input body model.SongDetail true "info songDetails"
// @Success 200 {int64} int64 "id songDetails"
// @Router /Song/Details [post]
func (h *Handler) SongDetails(c *gin.Context) {
	var songDetail model.SongDetail

	d, err := c.GetRawData()

	err = json.Unmarshal(d, &songDetail)
	if err != nil {
		log.Errorf("unmarshal handlerError")

		return
	}

	id, err := h.service.SongDetails(songDetail)
	if err != nil {
		log.Printf("Failed to create songDetail: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

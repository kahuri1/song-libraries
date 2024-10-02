package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) PutUpdateSongDetails(c *gin.Context) {
	var songDetails model.SongDetail
	d, err := c.GetRawData()

	err = json.Unmarshal(d, &songDetails)
	if err != nil {
		log.Errorf("unmarshal handlerError")
		return
	}
	_, err = h.service.PutUpdateSongDetails(songDetails)
	if err != nil {
		log.Printf("Failed to update song detail: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": fmt.Sprintf("updated")})
}

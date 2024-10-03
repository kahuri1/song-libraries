package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) AddSong(c *gin.Context) {

	var song model.Song

	d, err := c.GetRawData()

	err = json.Unmarshal(d, &song)
	if err != nil {
		log.Errorf("unmarshal handlerError")

		return
	}
	id, err := h.service.AddSong(song)
	if err != nil {
		log.Printf("Failed to create group: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})

}

package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) SongText(c *gin.Context) {
	var g model.GetSongText
	var b model.BodyPagination
	d, err := c.GetRawData()

	err = json.Unmarshal(d, &g)
	if err != nil {
		log.Errorf("unmarshal handlerError")
		return
	}
	err = json.Unmarshal(d, &b)
	if err != nil {
		log.Errorf("unmarshal handlerError")
		return
	}

	lines, err := h.service.SongText(g)
	if err != nil {
		log.Printf("Failed to create group: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	b.Lines = lines
	err = h.service.PaginationTextSong(&b)
	if err != nil {
		log.Printf("Failed to create group: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, b)
}

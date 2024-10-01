package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) PostAddGroup(c *gin.Context) {

	var group model.Group
	d, err := c.GetRawData()

	err = json.Unmarshal(d, &group)
	if err != nil {
		log.Errorf("unmarshal handlerError")

		return
	}

	id, err := h.service.PostAddGroup(group)
	if err != nil {
		log.Printf("Failed to create group: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

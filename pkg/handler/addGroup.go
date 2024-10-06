package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// @Summary AddGroup
// @Tags Group
// @Description create group
// @ID create-group
// @Accept json
// @Produce json
// @Param input body model.Group true "group name"
// @Success 200 {int64} int64 "id group"
// @Router /Group [post]
func (h *Handler) AddGroup(c *gin.Context) {

	var group model.Group
	d, err := c.GetRawData()

	err = json.Unmarshal(d, &group)
	if err != nil {
		log.Errorf("unmarshal handlerError")

		return
	}

	id, err := h.service.AddGroup(group)
	if err != nil {
		log.Printf("Failed to create group: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) Library(c *gin.Context) {
	var s model.LibraryConfig
	d, err := c.GetRawData()
	err = json.Unmarshal(d, &s)
	if err != nil {
		log.Errorf("unmarshal handlerError")
		return
	}

	library, err := h.service.Library(s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed GetLibrary"})
		return
	}
	c.JSON(http.StatusOK, library)
}

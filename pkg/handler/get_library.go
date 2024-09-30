package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetLibrary(c *gin.Context) {
	library, err := h.service.GetLibrary()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed GetLibrary"})
		return
	}
	c.JSON(http.StatusOK, library)
}

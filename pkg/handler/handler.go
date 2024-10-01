package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song-libraries/pkg/model"
)

type songLibsService interface {
	GetLibrary() (model.SongsResponse, error)
	PostAddGroup(g model.Group) (int64, error)
	PostAddSong(g model.Song) (int64, error)
	PostSongDetails(g model.SongDetail) (int64, error)
}

type Handler struct {
	service songLibsService
}

func Newhandler(service songLibsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/library", h.GetLibrary)
	r.POST("/Group", h.PostAddGroup)
	r.POST("/Song", h.PostAddSong)
	r.POST("/Song/Details", h.PostSongDetails)
	return r
}

package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kahuri1/song-libraries/docs"
	"github.com/kahuri1/song-libraries/pkg/model"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type songLibsService interface {
	Library(config model.LibraryConfig) (model.SongsResponse, error)
	AddGroup(g model.Group) (int64, error)
	AddSong(g model.Song) (int64, error)
	SongDetails(g model.SongDetail) (int64, error)
	UpdateSongDetails(g model.SongDetail) (bool, error)
	DeleteGroup(id int64) error
	DeleteSong(id int64) error
	SongText(g model.GetSongText) ([]string, error)
	PaginationTextSong(g *model.BodyPagination) error
	SongInfo(group string, song string) (model.SongDetailInfo, error)
}

type Handler struct {
	service songLibsService
}

func Newhandler(service songLibsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/library", h.Library)
	r.POST("/Group", h.AddGroup)
	r.DELETE("/Group/", h.DeleteGroup)
	r.POST("/Song", h.AddSong)
	r.DELETE("/Song/", h.DeleteSong)
	r.POST("/Song/Details", h.SongDetails)
	r.PUT("/Song/Details", h.UpdateSongDetails)
	r.GET("/Song/Details", h.SongText)
	r.POST("/info", h.SongInfo)
	return r
}

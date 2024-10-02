package service

import (
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
)

type repo interface {
	GetLibrary() (model.SongsResponse, error)
	PostAddGroup(g model.Group) (int64, error)
	CheckPostAddGroup(g model.Group) (int64, error)
	PostAddSong(g model.Song) (int64, error)
	CheckAddSong(g *model.Song)
	CheckDuplicateSong(g *model.Song) (int64, error)
	PostSongDetails(g model.SongDetail) (int64, error)
	CheckDuplicateSongDetail(g model.SongDetail) (int64, error)
	DeleteGroup(id int64) error
	DeleteSong(id int64) error
	PutUpdateSongDetails(g model.SongDetail) (bool, error)
}

type Service struct {
	repo repo
}

func NewService(repo repo) *Service {
	log.Info("service init")

	return &Service{
		repo: repo,
	}
}

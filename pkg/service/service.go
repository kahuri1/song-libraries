package service

import (
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
)

type repo interface {
	Library() (model.SongsResponse, error)
	AddGroup(g model.Group) (int64, error)
	CheckPostAddGroup(g model.Group) (int64, error)
	AddSong(g model.Song) (int64, error)
	CheckAddSong(g *model.Song)
	CheckDuplicateSong(g *model.Song) (int64, error)
	SongDetails(g model.SongDetail) (int64, error)
	CheckDuplicateSongDetail(g model.SongDetail) (int64, error)
	DeleteGroup(id int64) error
	DeleteSong(id int64) error
	UpdateSongDetails(g model.SongDetail) (bool, error)
	SongID(title string) (int64, error)
	SongText(g model.GetSongText) (string, error)
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

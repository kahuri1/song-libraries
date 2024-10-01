package service

import (
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
)

type repo interface {
	GetLibrary() (model.SongsResponse, error)
	PostAddGroup(g model.Group) (int64, error)
	CheckPostAddGroup(g model.Group) (int64, error)
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

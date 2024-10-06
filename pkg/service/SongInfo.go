package service

import (
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) SongInfo(group string, song string) (model.SongDetailInfo, error) {
	var g model.SongDetailInfo

	g, err := s.repo.SongInfo(group, song)
	if err != nil {
		log.Errorf("failed delete group: %w", err)
	}
	return g, nil
}

package service

import (
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) PutUpdateSongDetails(detail model.SongDetail) (bool, error) {

	val, err := s.repo.PutUpdateSongDetails(detail)
	if err != nil {
		log.Errorf("failed to done song detail: %w", err)
		return false, err
	}
	return val, nil
}

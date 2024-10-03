package service

import (
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) AddGroup(g model.Group) (int64, error) {
	idCheck, err := s.repo.CheckPostAddGroup(g)
	if err != nil {
		return idCheck, err
	}
	id, err := s.repo.AddGroup(g)
	if err != nil {
		log.Errorf("failed to create name group: %w", err)
	}
	return id, nil
}

package service

import log "github.com/sirupsen/logrus"

func (s *Service) DeleteSong(id int64) error {

	err := s.repo.DeleteGroup(id)
	if err != nil {
		log.Errorf("failed delete group: %w", err)
	}
	return nil
}

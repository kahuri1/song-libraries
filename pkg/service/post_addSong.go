package service

import (
	"fmt"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) PostAddSong(g model.Song) (int64, error) {
	s.repo.CheckAddSong(&g)

	idCheck, err := s.repo.CheckDuplicateSong(&g)
	if idCheck != 0 {
		return idCheck, fmt.Errorf("the song has already been created")
	}
	id, err := s.repo.PostAddSong(g)
	if err != nil {
		log.Errorf("failed to create song: %w", err)
	}
	return id, nil
}

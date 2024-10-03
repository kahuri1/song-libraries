package service

import (
	"fmt"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) SongDetails(g model.SongDetail) (int64, error) {
	checkSongID, err := s.repo.CheckDuplicateSongDetail(g)
	if checkSongID != 0 {
		return checkSongID, fmt.Errorf("the song_detail has already been created")
	}

	id, err := s.repo.SongDetails(g)
	if err != nil {
		log.Errorf("failed to create song_details: %w", err)
	}
	return id, nil
}

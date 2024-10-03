package service

import (
	"fmt"
	"github.com/kahuri1/song-libraries/pkg/model"
	log "github.com/sirupsen/logrus"
	"strings"
)

func (s *Service) SongText(g model.GetSongText) ([]string, error) {
	songID, err := s.repo.SongID(g.Title)
	if err != nil {
		log.Errorf("failed request search somgID: %w", err)
	}
	g.SongID = songID

	text, err := s.repo.SongText(g)
	if err != nil {
		log.Errorf("failed search song: %w", err)
	}
	textLines := strings.Split(text, "\\n")
	return textLines, nil
}

func (s *Service) PaginationTextSong(g *model.BodyPagination) error {
	g.Limit = 4
	page := g.Page
	start := (page - 1) * g.Limit
	end := start + g.Limit
	var filteredLines []string
	for _, line := range g.Lines {
		if line != "" {
			filteredLines = append(filteredLines, line)
		}
	}
	g.Lines = filteredLines
	if start >= len(g.Lines) {
		return fmt.Errorf("zero array")
	}
	if end > len(g.Lines) {
		end = len(g.Lines)
	}
	g.Lines = g.Lines[start:end]
	return nil
}

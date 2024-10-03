package service

import "github.com/kahuri1/song-libraries/pkg/model"

func (s *Service) Library() (model.SongsResponse, error) {

	library, err := s.repo.Library()
	if err != nil {
		return model.SongsResponse{}, err
	}

	return library, nil
}

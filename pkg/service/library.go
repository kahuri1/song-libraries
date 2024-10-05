package service

import "github.com/kahuri1/song-libraries/pkg/model"

func (s *Service) Library(c model.LibraryConfig) (model.SongsResponse, error) {

	library, err := s.repo.Library(c)
	if err != nil {
		return model.SongsResponse{}, err
	}

	return library, nil
}

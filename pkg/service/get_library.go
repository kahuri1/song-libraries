package service

import "github.com/kahuri1/song-libraries/pkg/model"

func (s *Service) GetLibrary() (model.SongsResponse, error) {

	library, err := s.repo.GetLibrary()
	if err != nil {
		return model.SongsResponse{}, err
	}

	return library, nil
}

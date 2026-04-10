package service

import (
	"server/internal/core/ports"
	"server/internal/dto"
	"server/internal/models"
	"time"

	"github.com/google/uuid"
)

type MovieService struct {
	repo      ports.MovieRepository
	jwtSecret string
}

func NewMovieService(r ports.MovieRepository, jwtSecret string) *MovieService {
	return &MovieService{
		repo:      r,
		jwtSecret: jwtSecret,
	}
}

func (s *MovieService) CreateMovie(data dto.CreateMovie) (models.Movie, error) {
	movie := models.Movie{
		ID:         uuid.New(),
		Name:       data.Name,
		Duration:   data.Duration,
		Release:    data.Release,
		Poster:     data.Poster,
		CreatedAt:  time.Now(),
	}

	return s.repo.CreateMovie(movie)
}

func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
	return s.repo.GetAllMovies()
}

func (s *MovieService) GetMovieByID(id uuid.UUID) (models.Movie, error) {
	return s.repo.GetMovieByID(id)
}

func (s *MovieService) UpdateMovie(id uuid.UUID, data dto.UpdateMovie) (models.Movie, error) {
	existingMovie, err := s.repo.GetMovieByID(id)
	if err != nil {
		return models.Movie{}, err
	}	

	existingMovie.Name = data.Name
	existingMovie.Duration = data.Duration
	existingMovie.Release = data.Release
	existingMovie.Poster = data.Poster	
	return s.repo.UpdateMovie(id,existingMovie)
}

func (s *MovieService) DeleteMovie(id uuid.UUID) error {
	_, err := s.repo.GetMovieByID(id)
	if err != nil {
		return err
	}	
	return s.repo.DeleteMovie(id)
}

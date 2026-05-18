package service

import (
	"server/internal/domain/entities"
	"server/internal/domain/ports"
	"server/internal/presentation/dto"
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

func (s *MovieService) CreateMovie(data dto.CreateMovie) (entities.Movie, error) {
	movie := entities.Movie{
		ID:        uuid.New(),
		Name:      data.Name,
		Duration:  data.Duration,
		Release:   data.Release,
		Poster:    data.Poster,
		CreatedAt: time.Now(),
	}

	return s.repo.CreateMovie(movie)
}

func (s *MovieService) GetAllMovies() ([]entities.Movie, error) {
	return s.repo.GetAllMovies()
}

func (s *MovieService) GetMovieByID(id uuid.UUID) (entities.Movie, error) {
	return s.repo.GetMovieByID(id)
}

func (s *MovieService) UpdateMovie(id uuid.UUID, data dto.UpdateMovie) (entities.Movie, error) {
	existingMovie, err := s.repo.GetMovieByID(id)
	if err != nil {
		return entities.Movie{}, err
	}

	existingMovie.Name = data.Name
	existingMovie.Duration = data.Duration
	existingMovie.Release = data.Release
	existingMovie.Poster = data.Poster
	return s.repo.UpdateMovie(id, existingMovie)
}

func (s *MovieService) DeleteMovie(id uuid.UUID) error {
	_, err := s.repo.GetMovieByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteMovie(id)
}

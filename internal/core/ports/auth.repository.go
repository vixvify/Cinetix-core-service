package ports

import (
	"server/internal/models"

	"github.com/google/uuid"
)

type MovieRepository interface {
	CreateMovie(movie models.Movie) (models.Movie, error)
	GetMovieByID(id uuid.UUID) (models.Movie, error)
	GetAllMovies() ([]models.Movie, error)
	UpdateMovie(id uuid.UUID, movie models.Movie) (models.Movie, error)
	DeleteMovie(id uuid.UUID) error
}
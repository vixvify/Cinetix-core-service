package ports

import (
	"server/internal/domain/entities"

	"github.com/google/uuid"
)

type MovieRepository interface {
	CreateMovie(movie entities.Movie) (entities.Movie, error)
	GetMovieByID(id uuid.UUID) (entities.Movie, error)
	GetAllMovies() ([]entities.Movie, error)
	UpdateMovie(id uuid.UUID, movie entities.Movie) (entities.Movie, error)
	DeleteMovie(id uuid.UUID) error
}

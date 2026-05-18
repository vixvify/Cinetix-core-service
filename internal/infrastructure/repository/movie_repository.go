package repository

import (
	"server/internal/domain/entities"
	"server/internal/domain/ports"
	appErr "server/internal/shared/errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieRepoGorm struct {
	db *gorm.DB
}

func NewMovieRepoGorm(db *gorm.DB) ports.MovieRepository {
	return &MovieRepoGorm{db: db}
}

func (r *MovieRepoGorm) CreateMovie(data entities.Movie) (entities.Movie, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return entities.Movie{}, appErr.Internal(err)
	}

	return data, nil
}

func (r *MovieRepoGorm) GetMovieByID(id uuid.UUID) (entities.Movie, error) {
	var movie entities.Movie
	if err := r.db.First(&movie, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.Movie{}, appErr.NotFound("movie not found", err)
		}
		return entities.Movie{}, appErr.Internal(err)
	}

	return movie, nil
}

func (r *MovieRepoGorm) GetAllMovies() ([]entities.Movie, error) {
	var movies []entities.Movie
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, appErr.Internal(err)
	}
	return movies, nil
}

func (r *MovieRepoGorm) UpdateMovie(id uuid.UUID, data entities.Movie) (entities.Movie, error) {
	var movie entities.Movie

	if err := r.db.Save(&movie).Error; err != nil {
		return entities.Movie{}, appErr.Internal(err)
	}

	return movie, nil
}

func (r *MovieRepoGorm) DeleteMovie(id uuid.UUID) error {
	var movie entities.Movie

	if err := r.db.Delete(&movie).Error; err != nil {
		return appErr.Internal(err)
	}
	return nil
}

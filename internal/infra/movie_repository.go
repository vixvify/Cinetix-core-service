package infra

import (
	"server/internal/core/ports"
	appErr "server/internal/errors"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieRepoGorm struct {
	db *gorm.DB
}

func NewMovieRepoGorm(db *gorm.DB) ports.MovieRepository {
	return &MovieRepoGorm{db: db}
}

func (r *MovieRepoGorm) CreateMovie(data models.Movie) (models.Movie, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return models.Movie{}, appErr.Internal(err)
	}

	return data, nil
}

func (r *MovieRepoGorm) GetMovieByID(id uuid.UUID) (models.Movie, error) {
	var movie models.Movie
	if err := r.db.First(&movie, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Movie{}, appErr.NotFound("movie not found", err)
		}	
		return models.Movie{}, appErr.Internal(err)
	}

	return movie, nil
}

func (r *MovieRepoGorm) GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, appErr.Internal(err)
	}
	return movies, nil
}

func (r *MovieRepoGorm) UpdateMovie(id uuid.UUID, data models.Movie) (models.Movie, error) {
	var movie models.Movie
	
	if err := r.db.Save(&movie).Error; err != nil {
		return models.Movie{}, appErr.Internal(err)
	}

	return movie, nil
}

func (r *MovieRepoGorm) DeleteMovie(id uuid.UUID) error {
	var movie models.Movie
	
	if err := r.db.Delete(&movie).Error; err != nil {
		return appErr.Internal(err)
	}
	return nil
}


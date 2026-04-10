package mapper

import (
	"server/internal/dto"
	"server/internal/models"
)

func ToMovieResponse(u models.Movie) dto.MovieResponse {
	return dto.MovieResponse{
		ID:         u.ID,
		Name:       u.Name,
		Duration:   u.Duration,
		Release:    u.Release,
		Poster:     u.Poster,
	}
}

func ToMovieResponseList(movies []models.Movie) []dto.MovieResponse {
	out := make([]dto.MovieResponse, 0, len(movies))
	for _, m := range movies {
		out = append(out, ToMovieResponse(m))
	}
	return out
}
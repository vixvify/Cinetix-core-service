package mapper

import (
	"server/internal/domain/entities"
	"server/internal/presentation/dto"
)

func ToMovieResponse(u entities.Movie) dto.MovieResponse {
	return dto.MovieResponse{
		ID:       u.ID,
		Name:     u.Name,
		Duration: u.Duration,
		Release:  u.Release,
		Poster:   u.Poster,
	}
}

func ToMovieResponseList(movies []entities.Movie) []dto.MovieResponse {
	out := make([]dto.MovieResponse, 0, len(movies))
	for _, m := range movies {
		out = append(out, ToMovieResponse(m))
	}
	return out
}

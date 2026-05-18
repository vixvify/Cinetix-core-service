package dto

import "github.com/google/uuid"

type CommmonMovie struct {
	Name     string `json:"name"`
	Duration string `json:"duration"`
	Release  string `json:"release"`
	Poster   string `json:"poster"`
}

type CreateMovie = CommmonMovie

type UpdateMovie = CommmonMovie

type MovieResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"`
	Release  string `json:"release"`
	Poster   string `json:"poster"`
}

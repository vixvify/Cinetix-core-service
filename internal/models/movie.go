package models

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string    `json:"name"`
	Duration  string    `json:"duration"`
	Release   string    `json:"release"`
	Poster    string    `json:"poster"`
	CreatedAt time.Time `json:"created_at"`
}


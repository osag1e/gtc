package model

import "github.com/google/uuid"

type Books struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
	Price  float64   `json:"price"`
}

func NewUUID() uuid.UUID {
	return uuid.New()
}

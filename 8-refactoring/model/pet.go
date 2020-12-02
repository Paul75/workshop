package model

import "github.com/google/uuid"

type Pet struct {
	ID          string
	Name        string
	Description string
}

func NewPet(name, desc string) *Pet {
	return &Pet{
		ID:          uuid.New().String(),
		Name:        name,
		Description: desc,
	}
}

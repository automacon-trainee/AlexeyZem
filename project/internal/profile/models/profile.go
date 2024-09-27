package models

import (
	"project/internal/library/models"
)

type Profile struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	Lastname  string        `json:"lastname"`
	TakenBook []models.Book `json:"taken_book"`
}

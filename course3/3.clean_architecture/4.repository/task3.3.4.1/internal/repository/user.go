package repository

import (
	"projectrepo/internal/models"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsExist  bool   `json:"is_exist"`
}

func (u User) ToModelsUser() *models.User {
	return &models.User{
		Username: u.Username,
		Password: u.Password,
	}
}

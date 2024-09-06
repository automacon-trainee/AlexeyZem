package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

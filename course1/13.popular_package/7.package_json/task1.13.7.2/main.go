package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	Text string `json:"text"`
}

func getUserFromJSON(data []byte) ([]User, error) {
	var users []User
	err := json.Unmarshal(data, &users)
	return users, err
}

func main() {
	data := []byte(`[{"name": "AlexeyZem", "age": 20, "comments": [{"text": "hello world"}]}]`)
	users, err := getUserFromJSON(data)
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}

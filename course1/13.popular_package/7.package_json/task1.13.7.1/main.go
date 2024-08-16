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

func getJSON(data []User) (string, error) {
	jsonData, err := json.Marshal(data)
	return string(jsonData), err
}

func main() {
	users := []User{
		{"John", 30, []Comment{{"some comment"}}},
		{"Gorn", 10, []Comment{{"some comment from Gorn"}}},
	}
	jsonData, err := getJSON(users)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(jsonData)
}

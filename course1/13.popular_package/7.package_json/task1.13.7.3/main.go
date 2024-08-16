package main

import (
	"encoding/json"
	"fmt"
	"os"
	path "path/filepath"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	Text string `json:"text"`
}

func WriteJSON(filepath string, data []User) error {
	err := os.MkdirAll(path.Dir(filepath), os.ModePerm)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, jsonData, os.ModePerm)
	return err
}

func main() {
	users := []User{
		{"Name", 9, []Comment{{"some text"}}},
	}
	err := WriteJSON("aa.txt", users)
	if err != nil {
		fmt.Println("error:", err)
	}
}

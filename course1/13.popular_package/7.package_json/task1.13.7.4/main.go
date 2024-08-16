package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteJSON(filePath string, data any) error {
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, jsonData, os.ModePerm)
}

func main() {
	data := []map[string]any{
		{"name": "AlexeyZem", "age": 24},
		{"name": "John", "age": 25},
	}
	err := WriteJSON("aa.txt", data)
	if err != nil {
		fmt.Println("error:", err)
	}
}

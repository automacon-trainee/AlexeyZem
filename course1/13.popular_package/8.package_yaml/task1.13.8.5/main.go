package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Person struct {
	Name string `yaml:"name" json:"name"`
	Age  int    `yaml:"age" json:"age"`
}

func unmarshal(data []byte, v any) error {
	errJ := json.Unmarshal(data, v)
	if errJ != nil {
		errM := yaml.Unmarshal(data, v)
		if errM != nil {
			return fmt.Errorf("unmarshal error. For json: %w, for yaml: %w", errJ, errM)
		}
	}

	return nil
}

func main() {
	data := []byte(`{"name":"John","age":30}`)
	var person Person
	err := unmarshal(data, &person)
	if err != nil {
		fmt.Println("unmarshal error:", err)
	} else {
		fmt.Println("Name:", person.Name, "Age:", person.Age)
	}
}

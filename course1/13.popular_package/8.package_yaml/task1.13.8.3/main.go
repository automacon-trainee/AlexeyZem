package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type User struct {
	Name     string    `yaml:"name"`
	Age      int       `yaml:"age"`
	Comments []Comment `yaml:"comments"`
}

type Comment struct {
	Text string `yaml:"text"`
}

func WriteYaml(filename string, data any) error {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		return err
	}
	bytes, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	err = os.WriteFile(filename, bytes, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

func main() {
	var user = User{Name: "John", Age: 25, Comments: []Comment{{"some comment"}, {"another comment"}}}
	err := WriteYaml("user.yaml", user)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
}

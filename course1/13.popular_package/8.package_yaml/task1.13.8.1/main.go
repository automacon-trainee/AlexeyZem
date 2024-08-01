package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
}
type Server struct {
	Port string `yaml:"port"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func getYAML(data []Config) string {
	out, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(out)
}

func main() {
	var config = Config{Server: Server{Port: ":8080"},
		DB: DB{Host: "localhost",
			Port:     ":8080",
			User:     "admin",
			Password: "admin123",
		},
	}
	fmt.Println(getYAML([]Config{config}))
}

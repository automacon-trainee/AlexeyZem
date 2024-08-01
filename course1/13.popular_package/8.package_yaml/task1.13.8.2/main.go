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

func getConfigFromYAML(data []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func main() {
	data := []byte(`
server:
 port: "8080"
db:
 host: "localhost"
 port: "5432"
 user: "admin"
 password: "password123"`)
	config, err := getConfigFromYAML(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}

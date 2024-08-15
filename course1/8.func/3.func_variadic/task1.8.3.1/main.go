package main

import (
	"fmt"
	"strings"
)

func userInfo(name string, age int, cities ...string) string {
	city := strings.Join(cities, ", ")
	res := fmt.Sprintf("Имя: %s, возраст: %d, города: %v", name, age, city)
	return res
}

func main() {
	name := "John"
	age := 10
	city := []string{"Moscow", "Saint Petersburg"}
	fmt.Println(userInfo(name, age, city...))
}

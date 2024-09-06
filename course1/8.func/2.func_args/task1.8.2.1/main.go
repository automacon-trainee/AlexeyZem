package main

import (
	"fmt"
)

func UserInfo(name, city, phone string, age, weight int) string {
	return fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s, Возраст: %d, Вес:  %d", name, city, phone, age, weight)
}

func main() {
	name := "Jane"
	city := "Paris"
	phone := "987-654-3210"
	age := 25
	weight := 150
	fmt.Println(UserInfo(name, city, phone, age, weight))
}

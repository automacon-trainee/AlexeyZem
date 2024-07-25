package main

import (
	"fmt"
)

func userInfo(name string, age int, cities ...string) string {
	res := fmt.Sprintf("Имя: %s, возраст: %d, города: ", name, age)

	for i, city := range cities {
		if i != len(cities)-1 {
			res += city + ", "
		} else {
			res += city
		}
	}

	return res
}

func main() {
	name := "John"
	age := 10
	city := []string{"Moscow", "Saint Petersburg"}
	fmt.Println(userInfo(name, age, city...))
}

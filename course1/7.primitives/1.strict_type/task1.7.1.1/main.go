package main

import (
	"fmt"
	"log"
)

func main() {
	var name string
	var age int
	var city string

	fmt.Print("Введите ваше имя: ")
	_, err := fmt.Scanln(&name)

	if err != nil {
		log.Println("Неверный формат данных:", err)
		return
	}

	fmt.Print("Введите ваш возраст: ")
	_, err = fmt.Scanln(&age)

	if err != nil {
		log.Println("Неверный формат данных:", err)
		return
	}

	fmt.Print("Введите ваш город: ")
	_, err = fmt.Scanln(&city)

	if err != nil {
		log.Println("Неверный формат данных: ", err)
		return
	}

	fmt.Println("Имя:", name, "\nВозраст:", age, "\nГород:", city)
}

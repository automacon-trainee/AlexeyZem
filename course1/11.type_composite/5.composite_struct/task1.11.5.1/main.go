package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/brianvoe/gofakeit"
)

type user struct {
	name string
	age  int
}

func getUsers() [10]user {
	maxAge := 60
	minAge := 18
	res := [10]user{}
	for i := 0; i < 10; i++ {
		res[i].name = gofakeit.Name()

		age, _ := rand.Int(rand.Reader, big.NewInt(int64(maxAge-minAge)))
		res[i].age = int(age.Int64()) + minAge
	}
	return res
}

func preparePrint(users []user) {
	for _, u := range users {
		fmt.Printf("Имя: %s, Возраст: %d\n", u.name, u.age)
	}
}

func main() {
	users := getUsers()
	preparePrint(users[:])
}

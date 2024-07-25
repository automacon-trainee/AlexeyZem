package main

import (
	"fmt"

	"github.com/icrowley/fake"
)

func GenerateFakeData() string {
	name := fake.FullName()
	address := fake.StreetAddress()
	phone := fake.Phone()
	email := fake.EmailAddress()

	res := "Name: " + name + "\nAddress: " + address + "\nPhone: " + phone + "\nEmail: " + email

	return res
}

func main() {
	fmt.Println(GenerateFakeData())
}

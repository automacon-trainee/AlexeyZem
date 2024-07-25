package main

import "fmt"

func changeInt(a *int) {
	*a = 20
}

func changeStr(a *string) {
	*a = "Goodbye, world!"
}

func changeFloat(a *float64) {
	*a = 6.28
}

func changeBool(a *bool) {
	*a = false
}

func main() {
	a := 0
	changeInt(&a)
	fmt.Println(a)

	b := ""
	changeStr(&b)
	fmt.Println(b)

	c := 2.0
	changeFloat(&c)
	fmt.Println(c)

	d := true
	changeBool(&d)
	fmt.Println(d)

}

package main

import "fmt"

func mutate(a *int) {
	*a = 42
}

func revereString(str *string) {
	sl := []byte(*str)
	for i := 0; i < len(sl)/2; i++ {
		sl[i], sl[len(sl)-1-i] = sl[len(sl)-1-i], sl[i]
	}
	*str = string(sl)
}

func main() {
	a := 0
	mutate(&a)
	fmt.Println(a)

	str := "Hello"
	revereString(&str)
	fmt.Println(str)
}

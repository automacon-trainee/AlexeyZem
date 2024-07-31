package main

import (
	"fmt"
	"reflect"
)

func getVariableType(x any) string {
	return fmt.Sprintf("%v", reflect.TypeOf(x))
}

func main() {}

package main

import (
	"fmt"
	"os/exec"
)

func ExecBin(filename string, args ...string) string {
	sl, err := exec.Command(filename, args...).Output()
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(sl)
}

func main() {
	fmt.Println(ExecBin("ls", "-la"))
	fmt.Println(ExecBin("wrong absolutely name and cal"))
}

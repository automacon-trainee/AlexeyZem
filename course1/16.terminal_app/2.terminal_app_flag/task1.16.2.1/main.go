package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printTree(path, prefix string, isLast bool, depth int) {
	if depth < 0 {
		return
	}
	sl := strings.Split(path, "/")
	fmt.Print(prefix, sl[0])
	sl = sl[1:]
	path = strings.Join(sl, "/")
	if depth <= 0 {
		isLast = !isLast
	}
	fmt.Println()
	if isLast {
		fmt.Println("|")
	}
	printTree(path, prefix+"----", isLast, depth-1)
}

func main() {
	var depth int
	flag.IntVar(&depth, "n", 0, "Целочисленное значение")
	flag.Parse()
	path := os.Args[len(os.Args)-1]
	if !strings.HasPrefix(path, "/") {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = filepath.Join(wd, path)
	}
	printTree(path, "", true, depth)
}

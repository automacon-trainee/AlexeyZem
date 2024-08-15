package main

func concatStrings(args ...string) string {
	res := ""
	for _, arg := range args {
		res += arg
	}
	return res
}

func main() {}

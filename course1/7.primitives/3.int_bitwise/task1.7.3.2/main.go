package main

import (
	"fmt"
	"log"
)

func getFilePermissions(flag int) string {
	other := flag % 10
	owner := flag / 100
	group := (flag / 10) % 10
	if owner > 7 || owner < 0 || group > 7 || group < 0 || other > 7 || other < 0 {
		return "Wrong permissions data"
	}

	res := "Owner:"
	res = updatePermission(res, owner)
	res += "Group:"
	res = updatePermission(res, group)
	res += "Other:"
	res = updatePermission(res, other)
	return res
}

func updatePermission(str string, name int) string {
	const (
		read    = 4
		write   = 2
		execute = 1
	)
	if name&read != 0 {
		str += "Read,"
	} else {
		str += "-,"
	}
	if name&write != 0 {
		str += "Write,"
	} else {
		str += "-,"
	}
	if name&execute != 0 {
		str += "Execute "
	} else {
		str += "- "
	}
	return str
}

func main() {
	var flag int
	_, err := fmt.Scanln(&flag)
	if err != nil {
		log.Println("Wrong data", err)
	}

	fmt.Println(getFilePermissions(flag))
}

package main

import (
	"sort"
)

func sortDesc(arr [8]int) [8]int {
	sort.Ints(arr[:])
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

func sortAsc(arr [8]int) [8]int {
	sort.Ints(arr[:])
	return arr
}

func main() {}

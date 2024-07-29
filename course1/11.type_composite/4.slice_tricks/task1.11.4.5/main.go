package main

func FilterDividers(xs []int, divider int) []int {
	var res []int
	for _, val := range xs {
		if val%divider == 0 {
			res = append(res, val)
		}
	}
	return res
}

func main() {}

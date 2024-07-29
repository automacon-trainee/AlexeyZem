package main

func sum(xs [8]int) int {
	res := 0
	for _, x := range xs {
		res += x
	}
	return res
}

func average(xs [8]int) float64 {
	return float64(sum(xs)) / float64(len(xs))
}

func averageFloat(xs [8]float64) float64 {
	var sumArr float64
	for _, x := range xs {
		sumArr += x
	}
	return sumArr / float64(len(xs))
}

func reverse(xs [8]int) [8]int {
	res := [8]int{}
	for i, x := range xs {
		res[len(xs)-i-1] = x
	}
	return res
}

func main() {

}

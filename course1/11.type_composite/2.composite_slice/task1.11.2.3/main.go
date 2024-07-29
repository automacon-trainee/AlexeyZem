package main

func findSingleNumber(number []int) int {
	for i := 0; i < len(number); i++ {
		for j := 0; j < len(number); j++ {
			if i != j && number[i]^number[j] == 0 {
				break
			}
			if j == len(number)-1 {
				return number[i]
			}
		}
	}
	return 0
}

func main() {

}

package main

func mergeMaps(map1, map2 map[string]int) map[string]int {
	for k, v := range map1 {
		map2[k] += v
	}
	return map2
}

func main() {}

package main

func InsertionSort(list []int) {
	n := len(list)
	for i := 1; i < n; i++ {
		temp := list[i]
		j := i - 1
		for j >= 0 && list[j] > temp {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = temp
	}
}

func SelectionSort(list []int) {
	n := len(list)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		list[i], list[minIndex] = list[minIndex], list[i]
	}
}

func MergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	mid := len(list) / 2
	left := MergeSort(list[:mid])
	right := MergeSort(list[mid:])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func QuickSort(list []int, low, high int) {
	if low < high {
		pi := partition(list, low, high)
		QuickSort(list, low, pi-1)
		QuickSort(list, pi+1, high)
	}
}

func partition(list []int, low, high int) int {
	pivot := list[high]
	i := low - 1
	for j := low; j < high; j++ {
		if list[j] < pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[high] = list[high], list[i+1]
	return i + 1
}

func GeneralSort(list []int) {
	const maxInsertionSort = 12
	if len(list) < maxInsertionSort {
		InsertionSort(list)
	}
	QuickSort(list, 0, len(list)-1)
}

func main() {}

package main

type User struct {
	ID   int
	Name string
	Age  int
}

func Merge(left, right []User) []User {
	if len(left) == 0 {
		return right
	}
	if len(right) == 0 {
		return left
	}
	result := make([]User, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0].ID <= right[0].ID {
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

func main() {
}

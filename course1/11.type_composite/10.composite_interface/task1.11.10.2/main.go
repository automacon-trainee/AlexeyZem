package main

//use any, because linter need it
//in this case any == interface{}

func Operate(f func(xs ...any) any, args ...any) any {
	return f(args...)
}

func Concat(args ...any) any {
	var res string
	for _, arg := range args {
		if s, ok := arg.(string); ok {
			res += s
		}
	}
	return res
}

func Sum(args ...any) any {
	sumF := 0.0
	sum := 0
	for _, arg := range args {
		if num, ok := arg.(float64); ok {
			sumF += num
		}
		if num, ok := arg.(int); ok {
			sum += num
		}
	}
	if sum == 0 {
		return sumF
	}
	return sum
}

func main() {}

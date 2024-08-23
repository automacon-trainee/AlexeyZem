package main

import (
	"fmt"
	"strings"
	"sync"
)

func waitGroupExample(goroutines ...func() string) string {
	res := ""
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, goroutine := range goroutines {
		wg.Add(1)
		go func(goroutine func() string) {
			defer wg.Done()
			mu.Lock()
			res += goroutine() + "\n"
			mu.Unlock()
		}(goroutine)
	}
	wg.Wait()
	return strings.TrimSpace(res)
}

func main() {
	count := 10
	goroutines := make([]func() string, count)
	for i := 0; i < count; i++ {
		j := i
		goroutines[i] = func() string {
			return fmt.Sprintf("goroutine %d", j)
		}
	}
	fmt.Println(waitGroupExample(goroutines...))
}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	res := contextWithTimeout(context.Background(), 1*time.Second, 2*time.Second)
	fmt.Println(res)
	res = contextWithTimeout(context.Background(), 2*time.Second, 1*time.Second)
	fmt.Println(res)
}
func contextWithTimeout(ctx context.Context, timeout, timeAfter time.Duration) string {
	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	select {
	case <-c.Done():
		return "Превышено время ожидания контекста"
	case <-time.After(timeAfter):
		return "Превышено время ождания"
	}
}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	res := contextWithDeadline(context.Background(), 1*time.Second, 2*time.Second)
	fmt.Println(res)
	res = contextWithDeadline(context.Background(), 2*time.Second, 1*time.Second)
	fmt.Println(res)
}

func contextWithDeadline(ctx context.Context, contextDeadline, timeAfter time.Duration) string {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(contextDeadline))
	defer cancel()
	select {
	case <-time.After(timeAfter):
		return "time after exceeded"
	case <-ctx.Done():
		return "context deadline exceeded"
	}
}

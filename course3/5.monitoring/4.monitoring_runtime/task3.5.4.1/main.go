// В приложении приведен код, который мониторит количество горутин, а также имитирует активную
// работу
// с созданием горутин.
// Реализуй  функцию  monitorGoroutines  так,  чтобы  она  корректно  отслеживала  изменения  в
// количестве
// горутин и уведомляла пользователя о значительных изменениях.

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

	"golang.org/x/sync/errgroup"
)

func CalculatePercentageChange(initalValue, finalValue int) (float64, error) {
	if finalValue == 0 {
		return 0.0, errors.New("cannot calculate percentage")
	}

	fromNumToPercent := 100.0

	return (float64(finalValue)/float64(initalValue) - 1) * fromNumToPercent, nil
}

func MonitorGoroutine(prevGoroutines int) {
	for {
		fmt.Println("Количество горутин: ", prevGoroutines)
		curGoroutines := runtime.NumGoroutine()
		percent, err := CalculatePercentageChange(prevGoroutines, curGoroutines)
		if err != nil {
			fmt.Printf("error in calculating percentage: %v\n", err)
		}
		if percent >= 20 {
			fmt.Println("Предупреждение: количество горутин увеличилось более чем на 20%")
		}
		if percent <= -20 {
			fmt.Println("Предупреждение: количество горутин уменьшилось более чем на 20%")
		}
		prevGoroutines = curGoroutines
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	g, _ := errgroup.WithContext(context.Background())
	go MonitorGoroutine(runtime.NumGoroutine())

	for i := 0; i < 100; i++ {
		g.Go(func() error {
			time.Sleep(5 * time.Second)
			return nil
		})
		time.Sleep(80 * time.Millisecond)
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

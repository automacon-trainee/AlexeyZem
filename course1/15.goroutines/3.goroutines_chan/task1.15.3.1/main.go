package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

type Order struct {
	ID       int
	Complete bool
}

var orders []Order
var completeOrders map[int]bool
var wg sync.WaitGroup
var processTime chan time.Duration
var sinceProgramStarted time.Duration
var count int
var limitCount int

func main() {
	count = 30
	limitCount = 5
	processTime = make(chan time.Duration, count)
	orders = GenerateOrders(count)
	completeOrders = GenerateCompleteOrders(count)
	programStart := time.Now()
	LimitSpawnOrderProcessing(limitCount)

	wg.Wait()
	sinceProgramStarted = time.Since(programStart)
	go func() {
		time.Sleep(time.Second * 1)
		close(processTime)
	}()
	checkTimeDifference(limitCount)
}

func checkTimeDifference(limitCount int) {
	var averageTime time.Duration
	var orderProcessTotalTime time.Duration
	var orderProcessedCount int
	for v := range processTime {
		orderProcessedCount++
		orderProcessTotalTime += v
	}
	if orderProcessedCount != count {
		panic("orderProcessedCount != count")
	}
	averageTime = orderProcessTotalTime / time.Duration(orderProcessedCount)
	fmt.Println("orderProcessTotalTime:", orderProcessTotalTime/time.Second)
	fmt.Println("averageTime:", averageTime/time.Second)
	fmt.Println("sinceProgramStarted:", sinceProgramStarted/time.Second)
	fmt.Println("sinceProgramStartedAverage:", sinceProgramStarted/(time.Duration(orderProcessedCount)*time.Second))
	fmt.Println("orderProcessTotalTime - sinceProgramStarted:", (orderProcessTotalTime-sinceProgramStarted)/time.Second)
	if (orderProcessTotalTime/time.Duration(limitCount)-sinceProgramStarted)/time.Second > 0 {
		panic("(orderProcessTotalTime/time.Duration(limitCount)-sinceProgramStarted)/time.Second > 0")
	}
}
func LimitSpawnOrderProcessing(limitCount int) {
	limit := make(chan struct{}, limitCount)
	var t time.Time
	for _, order := range orders {
		limit <- struct{}{}
		wg.Add(1)
		go func(order Order) {
			defer wg.Done()
			t = time.Now()
			OrderProcessing(order)
			time.Sleep(time.Second)
			processTime <- time.Since(t)
			<-limit
		}(order)
	}
}

func OrderProcessing(order Order) {
	if completeOrders[order.ID] {
		order.Complete = true
		fmt.Printf("order %v completed\n", order.ID)
	} else {
		fmt.Printf("order %v failed \n", order.ID)
	}
}

func GenerateCompleteOrders(limitCount int) map[int]bool {
	res := make(map[int]bool)
	for i := 0; i < limitCount; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(limitCount)))
		if int(num.Int64()) > limitCount/2 {
			res[i] = true
		}
	}
	return res
}
func GenerateOrders(count int) []Order {
	res := make([]Order, count)
	for i := 0; i < count; i++ {
		res[i] = Order{i, false}
	}
	return res
}

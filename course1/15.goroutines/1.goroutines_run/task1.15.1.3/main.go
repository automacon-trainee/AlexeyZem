package main

import (
	"fmt"
	"time"
)

func worker(jobs, results chan int, idx int) {
	for j := range jobs {
		fmt.Println("worker", idx, "job", j, "start")
		time.Sleep(time.Second)
		results <- j * j
		fmt.Println("worker", idx, "job", j, "stop")
	}
}

func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for i := 0; i < 3; i++ {
		go worker(jobs, results, i)
	}
	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < numJobs; i++ {
		fmt.Println(<-results)
	}
}

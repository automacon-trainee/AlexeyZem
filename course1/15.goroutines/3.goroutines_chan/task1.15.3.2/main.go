package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://httpbin.org/get"
	parallelRequest := 5
	requestCount := 50
	result := benchRequest(url, parallelRequest, requestCount)
	for i := 0; i < requestCount; i++ {
		statusCode := <-result
		if statusCode != 200 {
			panic(fmt.Sprintf("Ошибка при отправке запроса: %v", statusCode))
		}
	}
	close(result)
	fmt.Println("Все горутины завершили работу")
}

func benchRequest(url string, parallelRequest, requestCount int) chan int {
	ch := make(chan int, parallelRequest)
	res := make(chan int)
	for i := 0; i < requestCount; i++ {
		go func() {
			statusCode, err := httpRequest(url)
			if err != nil {
				res <- 0
			}
			res <- statusCode
			<-ch
		}()
	}
	return res
}
func httpRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

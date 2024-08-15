package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ValuteInfo struct {
	CharCode string
	Value    float64
}
type Course struct {
	Valute map[string]ValuteInfo
}

func currencyPairRate(firstCurrency, secondCurrency string, num float64) (float64, error) {
	url := "https://www.cbr-xml-daily.ru/daily_json.js"
	r, err := http.NewRequestWithContext(context.Background(), "GET", url, new(bytes.Buffer))
	if err != nil {
		return 0.0, err
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	client := http.Client{}
	req, err := client.Do(r)
	defer func() {
		err = req.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		return 0.0, err
	}

	course, err := parse(req)
	if err != nil {
		return 0.0, err
	}

	return calculate(course, firstCurrency, secondCurrency, num)
}

func parse(req *http.Response) (Course, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return Course{}, err
	}

	course := Course{}
	err = json.Unmarshal(body, &course)
	if err != nil {
		return Course{}, err
	}
	return course, nil
}

func calculate(course Course, firstCurrency, secondCurrency string, num float64) (float64, error) {
	var sum float64
	if val, ok := course.Valute[firstCurrency]; ok {
		sum = val.Value * num
	} else {
		return 0.0, fmt.Errorf("no such field")
	}

	if val, ok := course.Valute[secondCurrency]; ok {
		sum /= val.Value
	} else {
		return 0.0, fmt.Errorf("no such field")
	}
	return sum, nil
}

func main() {
	money := 100.0
	firstCurrency := "USD"
	SecondCurrency := "EUR"
	rate, err := currencyPairRate(firstCurrency, SecondCurrency, money)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rate)
}

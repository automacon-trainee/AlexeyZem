package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	res := CallService()
	fmt.Println(res)
}

func CallService() string {
	data := make(chan string, 2)
	serviceLocator := NewServiceLocator()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		str, err := serviceLocator.slowService(ctx)
		if err != nil {
			fmt.Println(err)
		}
		data <- str
	}()
	go func() {
		defer cancel()
		str, err := serviceLocator.fastService(ctx)
		if err != nil {
			fmt.Println(err)
		}
		data <- str
	}()
	select {
	case <-ctx.Done():
		if len(data) > 1 {
			panic("more than one result received")
		}
	}
	CheckService(serviceLocator)
	return <-data
}

func CheckService(s *ServiceLocator) {
	if s.slow {
		panic("slow")
	}
	if !s.fast {
		panic("fast")
	}
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{client: &http.Client{
		Timeout: 5 * time.Second,
	}}
}

type ServiceLocator struct {
	client *http.Client
	fast   bool
	slow   bool
}

func (s *ServiceLocator) slowService(ctx context.Context) (string, error) {
	defer func() {
		s.slow = true
	}()
	time.Sleep(2 * time.Second)
	return s.doRequest(ctx, "https://api.exmo.com/v1/ticker")
}

func (s *ServiceLocator) fastService(ctx context.Context) (string, error) {
	defer func() {
		s.fast = true
	}()
	time.Sleep(500 * time.Millisecond)
	return s.doRequest(ctx, "https://api.exmo.com/v1/ticker")
}

func (s *ServiceLocator) doRequest(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error doing request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}
	return string(body), nil
}

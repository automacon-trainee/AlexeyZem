package main

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestProductString(t *testing.T) {
	p := Product{Price: 10, Name: "apple", Count: 10, CreatedAt: time.Now()}
	s := p.String()
	want := fmt.Sprintf("Name: %s, Price: %.2f, Count: %v", p.Name, p.Price, p.Count)
	if s != want {
		t.Errorf("String() = %s, want %s", s, want)
	}
}

func TestByCount(t *testing.T) {
	products := generateProducts(100)
	sort.Sort(ByCount(products))
	for i := 1; i < len(products); i++ {
		if products[i].Count < products[i-1].Count {
			t.Errorf("ByCount() not sort by count")
		}
	}
}

func TestByPrice(t *testing.T) {
	products := generateProducts(100)
	sort.Sort(ByPrice(products))
	for i := 1; i < len(products); i++ {
		if products[i].Price < products[i-1].Price {
			t.Errorf("ByCount() not sort by count")
		}
	}
}

func TestByCreatedAt(t *testing.T) {
	products := generateProducts(100)
	sort.Sort(ByCreatedAt(products))
	for i := 1; i < len(products); i++ {
		if products[i].CreatedAt.Unix() < products[i-1].CreatedAt.Unix() {
			t.Errorf("ByCount() not sort by count")
		}
	}
}

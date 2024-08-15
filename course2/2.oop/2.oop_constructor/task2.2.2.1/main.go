package main

import (
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOptions func(*Order)

func NewOrder(id int, opts ...OrderOptions) *Order {
	order := &Order{
		ID: id,
	}
	for _, opt := range opts {
		opt(order)
	}
	return order
}

func WithCustomerID(id string) OrderOptions {
	return func(order *Order) {
		order.CustomerID = id
	}
}

func WithItems(items ...string) OrderOptions {
	return func(order *Order) {
		order.Items = items
	}
}

func WithOrderDate(orderDate time.Time) OrderOptions {
	return func(order *Order) {
		order.OrderDate = orderDate
	}
}

func main() {}

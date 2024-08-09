package main

import (
	"reflect"
	"testing"
	"time"
)

func TestNewOrder(t *testing.T) {
	timeNow := time.Now()
	items := []string{"item1", "item2", "item3"}
	customerId := "custom"
	order := NewOrder(1,
		WithOrderDate(timeNow),
		WithItems(items...),
		WithCustomerID(customerId),
	)
	if order.ID != 1 {
		t.Errorf("Order.ID = %d, want 1", order.ID)
	}
	if order.OrderDate != timeNow {
		t.Errorf("Order.OrderDate = %v, want %v", order.OrderDate, timeNow)
	}
	if !reflect.DeepEqual(order.Items, items) {
		t.Errorf("Order.Items = %v, want %v", order.Items, items)
	}
	if order.CustomerID != customerId {
		t.Errorf("Order.CustomerID = %v, want %v", order.CustomerID, customerId)
	}
}

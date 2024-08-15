package main

import (
	"testing"
)

func TestOrder(t *testing.T) {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	if len(order.Dishes) != 2 {
		t.Errorf("len(order.Dishes) = %d; want 2(problem in adding)", len(order.Dishes))
	}

	order.CalculateTotal()
	if order.Total != dish2.Price+dish1.Price {
		t.Errorf("expect %v, got %v", dish2.Price+dish1.Price, order.Total)
	}
	order.RemoveDish(dish1)
	if len(order.Dishes) != 1 {
		t.Errorf("len(order.Dishes) = %d; want 1 (problem in delete)", len(order.Dishes))
	}
	order.CalculateTotal()
	if order.Total != dish2.Price {
		t.Errorf("expect %v, got %v", dish2.Price, order.Total)
	}
}

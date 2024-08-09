package main

import (
	"fmt"
)

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

type DineInOrder struct {
	orderDetails map[string]int
}

func (d *DineInOrder) AddItem(item string, quantity int) error {
	if _, ok := d.orderDetails[item]; ok {
		return fmt.Errorf("item %s already exists", item)
	}
	if quantity <= 0 {
		return fmt.Errorf("quantity cannot be negative")
	}
	d.orderDetails[item] = quantity
	fmt.Println("dine in order:", d.orderDetails)
	return nil
}

func (d *DineInOrder) RemoveItem(item string) error {
	if _, ok := d.orderDetails[item]; !ok {
		return fmt.Errorf("item %s does not exist", item)
	}
	delete(d.orderDetails, item)
	fmt.Println("dine in order:", d.orderDetails)
	return nil
}

func (d *DineInOrder) GetOrderDetails() map[string]int {
	return d.orderDetails
}

type TakeAwayOrder struct {
	orderDetails map[string]int
}

func (t *TakeAwayOrder) AddItem(item string, quantity int) error {
	if _, ok := t.orderDetails[item]; ok {
		return fmt.Errorf("item %s already exists", item)
	}
	if quantity <= 0 {
		return fmt.Errorf("quantity cannot be negative")
	}
	t.orderDetails[item] = quantity
	fmt.Println("take away order:", t.orderDetails)
	return nil
}

func (t *TakeAwayOrder) RemoveItem(item string) error {
	if _, ok := t.orderDetails[item]; !ok {
		return fmt.Errorf("item %s does not exist", item)
	}
	delete(t.orderDetails, item)
	fmt.Println("take away order:", t.orderDetails)
	return nil
}

func (t *TakeAwayOrder) GetOrderDetails() map[string]int {
	return t.orderDetails
}

func ManageOrder(o Order) {
	err := o.AddItem("Pizza", 2)
	if err != nil {
		fmt.Println(err)
	}
	err = o.AddItem("Burger", 3)
	if err != nil {
		fmt.Println(err)
	}
	err = o.RemoveItem("Pizza")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}
	ManageOrder(dineIn)
	ManageOrder(takeAway)
}

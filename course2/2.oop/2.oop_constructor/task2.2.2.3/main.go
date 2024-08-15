package main

import (
	"errors"
)

type Customer struct {
	ID   int
	Name string
	Acc  Account
}

type Options func(*Customer)

func NewCustomer(id int, opts ...Options) *Customer {
	c := &Customer{
		ID: id,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithName(name string) Options {
	return func(c *Customer) {
		c.Name = name
	}
}

func WithAcc(acc Account) Options {
	return func(c *Customer) {
		c.Acc = acc
	}
}

type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type SavingsAccount struct {
	balance float64
	minSum  float64
}

func (c *SavingsAccount) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("cannot deposit negative amount")
	}
	c.balance += amount
	return nil
}

func (c *SavingsAccount) Withdraw(amount float64) error {
	if c.balance < c.minSum {
		return errors.New("cannot withdraw, because balance is smaller than min sum")
	}

	if amount > c.balance {
		return errors.New("cannot withdraw negative amount")
	}
	c.balance -= amount
	return nil
}

func (c *SavingsAccount) Balance() float64 {
	return c.balance
}

func main() {}

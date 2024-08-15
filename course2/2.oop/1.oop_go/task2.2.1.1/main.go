package main

import (
	"errors"
	"fmt"
	"log"
)

type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type CurrentAccount struct {
	balance float64
}

func (c *CurrentAccount) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("cannot deposit negative amount")
	}
	c.balance += amount
	return nil
}

func (c *CurrentAccount) Withdraw(amount float64) error {
	if amount > c.balance {
		return errors.New("cannot withdraw negative amount")
	}
	c.balance -= amount
	return nil
}

func (c *CurrentAccount) Balance() float64 {
	return c.balance
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
		return errors.New("cannot withdraw, because balance is smaller than 500")
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

func ProcessAccount(a Account) {
	err := a.Deposit(400)
	if err != nil {
		log.Println(err)
	}
	err = a.Withdraw(200)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Balance: %2f\n", a.Balance())
}

func main() {
	c := &CurrentAccount{}
	s := &SavingsAccount{minSum: 500}
	ProcessAccount(c)
	ProcessAccount(s)
}

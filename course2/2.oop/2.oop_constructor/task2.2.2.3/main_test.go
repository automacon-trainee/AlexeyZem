package main

import (
	"fmt"
	"log"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	savings := SavingsAccount{minSum: 1000}
	err := savings.Deposit(1000)
	if err != nil {
		log.Println(err)
	}
	customer := NewCustomer(1, WithName("John"), WithAcc(&savings))
	if customer.Name != "John" {
		t.Errorf("customer.Name should be John")
	}
	if customer.ID != 1 {
		t.Errorf("customer.ID should be 1")
	}
	if customer.Acc != &savings {
		t.Errorf("customer.Acc should be savings")
	}
}

type TestCase struct {
	amount  float64
	err     error
	balance float64
	minSum  float64
}

func TestAccount(t *testing.T) {
	balance := 300.0
	testsWithdraw := []TestCase{
		{amount: 200, err: nil, balance: 100, minSum: 200},
		{amount: 100, err: fmt.Errorf("cannot withdraw, because balance is smaller than min sum"), balance: 300, minSum: 500},
		{amount: 500, err: fmt.Errorf("cannot withdraw negative amount"), balance: 300, minSum: 100},
	}
	for _, tt := range testsWithdraw {
		s := SavingsAccount{balance: balance, minSum: tt.minSum}
		err := s.Withdraw(tt.amount)
		if s.Balance() != tt.balance {
			t.Errorf("s.balance should be %v, but is %v", tt.balance, s.balance)
		}
		if err != nil && tt.err != nil && err.Error() != tt.err.Error() {
			t.Errorf("err should be %v, but is %v", tt.err, err)
		}
	}
	testAmount := []TestCase{
		{amount: 200, err: nil, balance: 500, minSum: 200},
		{amount: -10, err: fmt.Errorf("cannot deposit negative amount"), balance: 300, minSum: 100},
	}
	for _, tt := range testAmount {
		s := SavingsAccount{balance: balance, minSum: tt.minSum}
		err := s.Deposit(tt.amount)
		if s.Balance() != tt.balance {
			t.Errorf("s.balance should be %v, but is %v", tt.balance, s.balance)
		}
		if err != nil && tt.err != nil && err.Error() != tt.err.Error() {
			t.Errorf("err should be %v, but is %v", tt.err, err)
		}
	}
}

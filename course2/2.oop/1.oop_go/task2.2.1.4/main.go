package main

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) error
}

func ProcessPayment(p PaymentMethod, amount float64) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println("Не удалось обработать платеж:", err)
	}
}

type CreditCard struct {
	balance float64
}

func (c *CreditCard) Pay(amount float64) error {
	if c.balance < amount {
		return errors.New("недостаточный баланс")
	}
	if amount < 0 {
		return errors.New("недопустимая сумма платежа")
	}
	c.balance -= amount
	fmt.Printf("Оплачено %v, с помощью кредитной карты\n", amount)
	return nil
}

type Bitcoin struct {
	balance float64
}

func (b *Bitcoin) Pay(amount float64) error {
	if b.balance < amount {
		return errors.New("недостаточный баланс")
	}
	if amount < 0 {
		return errors.New("недопустимая сумма платежа")
	}
	b.balance -= amount
	fmt.Printf("Оплачено %v, с помощью биткоина\n", amount)
	return nil
}

func main() {
	cc := &CreditCard{500.0}
	btc := &Bitcoin{2.0}
	ProcessPayment(cc, 200)
	ProcessPayment(btc, 1)
}

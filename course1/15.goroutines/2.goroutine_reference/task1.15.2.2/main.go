// Линтер не проходит из-зв глобальной переменной
// Задача на исправление, анонимной функции, поэтому не трогал переменную.

package main

import (
	"fmt"
	"sync"
)

type User struct {
	ID      int
	Name    string
	Bonuses int
}

type BonusOperation struct {
	UserID int
	Amount int
}

var users = []User{
	{ID: 1, Name: "Bob", Bonuses: 100},
	{ID: 2, Name: "Alice", Bonuses: 200},
	{ID: 3, Name: "Kate", Bonuses: 300},
	{ID: 4, Name: "Tom", Bonuses: 400},
	{ID: 5, Name: "John", Bonuses: 500},
}

func main() {
	bonusOperations := []BonusOperation{
		{UserID: 1, Amount: 10},
		{UserID: 2, Amount: 20},
		{UserID: 3, Amount: 30},
		{UserID: 4, Amount: 40},
		{UserID: 5, Amount: 50},
	}
	DeductBonuses(users, bonusOperations)
	for _, user := range users {
		fmt.Printf("User %s has %d bonuses\n", user.Name, user.Bonuses)
	}
}

func DeductBonuses(users []User, bonuses []BonusOperation) {
	wg := sync.WaitGroup{}
	for i, user := range users {
		wg.Add(1)
		go func(user User, i int) {
			defer wg.Done()
			for _, bonus := range bonuses {
				if bonus.UserID == user.ID {
					users[i].Bonuses -= bonus.Amount
				}
			}
		}(user, i)
	}
	wg.Wait()
}

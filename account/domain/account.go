package domain

import (
	"errors"
	"sync"
)

type Account struct {
	mu      sync.Mutex
	Name    string
	Balance int
}

func GetAccount(name string) (*Account, error) {
	return &Account{
		Name: name,
		// Balance: balance,
	}, nil
}

func AddAccount(name string, balance int) (*Account, error) {
	return &Account{
		Name:    name,
		Balance: balance,
	}, nil
}

func (a Account) Transfer(ToAccount Account, amount int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.Balance < amount {
		return errors.New("insufficient balance")
	}

	a.Balance -= amount
	return nil
}

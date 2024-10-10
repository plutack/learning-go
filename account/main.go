package account

import (
	"fmt"
	"sync"
)

var (
	lastRegisteredId int // this initializes to 0
	mu               sync.Mutex
)

type (
	money   float64
	Account struct {
		Id           int
		Name         string
		Balance      money
		InterestRate int
		accountType  string
	}
)

func New(name string, balance money, accountType string) (*Account, error) {
	mu.Lock()
	defer mu.Unlock()

	if balance < 2000 {
		return nil, fmt.Errorf("initial balance %.2f is lower than 2000", balance)
	}
	var interestRate int

	if accountType != "fixedDeposit" && accountType != "savings" {
		return nil, fmt.Errorf("%s is an invalid account type", accountType)
	}

	if accountType == "fixedDeposit" {
		interestRate = 5
	}
	if accountType == "savings" {
		interestRate = 2
	}
	lastRegisteredId++

	newAccount := Account{
		Id:           lastRegisteredId,
		Name:         name,
		Balance:      balance,
		InterestRate: interestRate,
		accountType:  accountType,
	}
	return &newAccount, nil
}

func (a *Account) GetBalance() money {
	return a.Balance
}

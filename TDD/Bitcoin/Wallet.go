package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface{
	String() string
}

type Wallet struct {
	balance Bitcoin
}

// in GO the args are copied
func (w *Wallet)Deposit(amount Bitcoin) {
	fmt.Printf("address of balance is", &w.balance)
	w.balance += amount
}

func (w *Wallet)Balance() Bitcoin {
	return (*w).balance
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance{
		return errors.New("Can't do that tnx")
	}
	w.balance -=amount
	return nil
}

package main

import (
	"fmt"
	"errors"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var operationErrors = map[string]error{
	"nofunds": errors.New("Insuficient Funds"),
}

func NewWallet(balance Bitcoin) Wallet {
	return Wallet{
		balanceValue: balance,
		operationErrors: operationErrors,
	}
}

type Wallet struct {
	balanceValue Bitcoin
	operationErrors map[string]error
}

func (w *Wallet) balance() Bitcoin {
	return w.balanceValue
}

func (w *Wallet) deposit(value Bitcoin) {
	w.balanceValue += value
}

func (w *Wallet) withdraw(value Bitcoin) error {
	if w.balanceValue < value {
		return w.operationErrors["nofunds"]
	}
	w.balanceValue -= value
	return nil
}

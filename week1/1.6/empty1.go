package main

import (
	"fmt"
	"strconv"
)

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough money bruh")
	}
	w.Cash -= amount
	return nil
}

func (w *Wallet) String() string {
	return "Wallet with " + strconv.Itoa(w.Cash) + " money"
}

func main() {
	wallet := &Wallet{Cash: 100}
	fmt.Printf("Raw payment: %#v\n", wallet)
	fmt.Printf("Sposob oplaty: %s\n", wallet)
}

package main

import (
	"fmt"
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

type Payer interface {
	Pay(int) error
}

func Buy(in interface{}) {
	var p Payer
	var ok bool
	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T is not a payment method\n\n", in)
		return
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Payment error with %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Thanks for purchase with %T\n\n", p)
}

func main() {
	wallet := &Wallet{Cash: 100}
	num := 140
	Buy(wallet)
	Buy(&num)
	// fmt.Printf("Raw payment: %#v\n", wallet)
	// fmt.Printf("Sposob oplaty: %s\n", wallet)
}

package main

import "fmt"

// Wallet

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

// Wallet ends

// Card

type Card struct {
	Balance    int
	ValidUntil string
	CardHolder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Not enough money")
	}
	c.Balance -= amount
	return nil
}

// Card ends

// ApplePay

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Not enough money on apple bruh")
	}
	a.Money -= amount
	return nil
}

// ApplePay ends

type Payer interface {
	Pay(int) error
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Payment error with %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Thanks for payment with %T\n\n", p)
}

func main() {
	wallet := &Wallet{Cash: 100}
	Buy(wallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100, CardHolder: "Bobins"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)
}

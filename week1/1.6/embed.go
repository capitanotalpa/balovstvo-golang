package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Ringer interface {
	Ring(string) error
}

type NFCPhone interface {
	Payer
	Ringer
}

type Phone struct {
	Money   int
	AppleID string
}

func (p *Phone) Pay(amount int) error {
	if p.Money < amount {
		return fmt.Errorf("Not enough moneeey")
	}
	p.Money -= amount
	return nil
}

func (p *Phone) Ring(number string) error {
	if number == "" {
		return fmt.Errorf("Enter number please")
	}
	return nil
}

func PayForMetroWithPhone(phone NFCPhone) {
	err := phone.Pay(1)
	if err != nil {
		fmt.Printf("Payment error %v\n\n", err)
		return
	}
	fmt.Printf("Opened metro daaaam %T\n", phone)
}

func main() {
	myPhone := &Phone{Money: 9}
	PayForMetroWithPhone(myPhone)
}

package main

import (
	"fmt"
	"visibility/person"
)

func main() {
	p := person.NewPerson(1, "Boba", "secret")

	fmt.Printf("main.PrintPerson: %+v\n", p)

	secret := person.GetSecret(p)
	fmt.Println("GetSecret", secret)
}

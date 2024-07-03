package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	var acc Account = Account{
		Id:     1,
		Name:   "Bobus",
		Person: Person{Address: "Moscow", Name: "Bibino"},
	}
	fmt.Printf("%#v\n", acc)

	acc.Owner = Person{2, "Bobino", "Yoshkar-Ola"}

	fmt.Println("%#v\n", acc)

	fmt.Println(acc.Address)
	fmt.Println(acc.Person.Name, acc.Name)
}

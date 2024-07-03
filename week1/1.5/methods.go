package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

// does not change anything (like me)
func (p Person) UpdateName(name string) {
	p.Name = name
}

// YEAH
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (a *Account) SetName(name string) {
	a.Name = name
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	person := Person{1, "Bibi"}
	fmt.Printf("%#v\n", person)
	person.SetName("Aoaoao")
	// (&person).SetName("Aoaoaoa")
	fmt.Printf("updated person: %#v\n", person)

	// return

	var acc Account = Account{
		Id:   1,
		Name: "Boba",
		Person: Person{
			Id:   2,
			Name: "Lololo",
		},
	}
	fmt.Printf("%#v\n", acc)
	acc.SetName("Hihich")
	fmt.Printf("%#v\n", acc)
	acc.Person.SetName("Hohich")
	fmt.Printf("%#v\n", acc)

	sl := MySlice([]int{1, 2})
	sl.Add(13)
	fmt.Println(sl.Count(), sl)
}

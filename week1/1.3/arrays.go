package main

import "fmt"

func main() {
	var a1 [3]int
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a 1 full %#v\n", a1)

	const size = 2
	var a2 [2 * size]bool
	fmt.Println("a2", a2)

	a3 := [...]int{1, 2, 3}
	fmt.Println("a3", a3)

	
	// a3[4] = 12

	var idx =  4
	a3[idx] = 4
}
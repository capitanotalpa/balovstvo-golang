package main

import "fmt"

func deferTest() {
	defer func() {
		fmt.Println("in panic mode")
		if err := recover(); err != nil {
			fmt.Println("panic at the disco 1: ", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic at the disco 2: ", err)
			panic("second panic")
		}
	}()
	fmt.Println("Some useful work")
	panic("something bad happened")
	return
}

func main() {
	deferTest()
	return
}

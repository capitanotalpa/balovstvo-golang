package main

import "fmt"

func doNothing() {
	fmt.Println("i'm regular func")
}

func main() {
	func(in string) {
		fmt.Println("anon func out:", in)
	}("nothing")

	printer := func(in string) {
		fmt.Println("prionter outs: ", in)
	}
	printer("as variable")

	type strFuncType func(string)

	worker := func(callback strFuncType) {
		callback("as callback")
	}
	worker(printer)

	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}
	successLogger := prefixer("SUCCESS")
	successLogger("expected behaviour")
	failLogger := prefixer("FAIL")
	failLogger("unexpected behaviour")
}

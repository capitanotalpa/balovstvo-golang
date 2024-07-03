package main

import "fmt"

func main() {
	boolVal := true
	if boolVal {
		fmt.Println("boolVal is true")
	}

	mapVal := map[string]string{"name": "hohohoho"}
	if keyValue, keyExist := mapVal["name"]; keyExist {
		fmt.Println("name =", keyValue)
	}
	if _, keyExist := mapVal["name"]; keyExist {
		fmt.Println("key 'name' exists")
	}

	cond := 1
	if cond == 1 {
		fmt.Println("cond is 1")
	} else if cond == 2 {
		fmt.Println("cond is 2")
	}

	strVal := "name"
	switch strVal {
	case "name":
		fallthrough
	case "test", "lastName":
		// some work
	default:
		// some work
	}

	var val1, val2 = 2, 2
	switch {
	case val1 > 1 || val2 < 11:
		fmt.Println("block 1")
	case val2 > 10:
		fmt.Println("block 2")
	}

Loop:
	for key, val := range mapVal {
		println("switch in loop", key, val)
		switch {
		// case key == "lastName":
		// break

		case key == "firstName" && val == "hohohoho":
			println("switch - break loop here")
			break Loop
		}
	} // end for
}

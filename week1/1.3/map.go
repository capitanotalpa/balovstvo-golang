package main

import "fmt"

func main() {
	var user map[string]string = map[string]string{
		"name":     "Vasya",
		"lastName": "Pupkin",
	}

	profile := make(map[string]string, 10)

	mapLength := len(user)
	fmt.Printf("%d %+v\n", mapLength, profile)

	mName := user["middleName"]
	fmt.Println("mName:", mName)

	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	delete(user, "lastName")
	fmt.Printf("%#v\n", user)
}

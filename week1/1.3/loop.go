package main

import "fmt"

func main() {
	for {
		fmt.Println("loop iteration")
		break
	}

	isRun := true
	for isRun {
		fmt.Println("loop iteration with condition")
		isRun = false
	}

	for i := 0; i < 2; i++ {
		fmt.Println("loop iteration", i)
		if i == 1 {
			fmt.Println("cont", i)
			continue
		}
	}

	sl := []int{1, 2, 3}
	idx := 0

	for idx < len(sl) {
		fmt.Println("while-type loop, idx:", idx, "value", sl[idx])
		idx++
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("c-style loop, ", i, sl[i])
	}

	for idx := range sl {
		fmt.Println("range by slice index", idx)
	}

	for idx, val := range sl {
		fmt.Println("range slice idx-val", idx, val)
	}

	profile := map[int]string{1: "biba", 2: "boba"}

	for key := range profile {
		fmt.Println("by map key", key)
	}

	for key, value := range profile {
		fmt.Println("by map key/val", key, value)
	}

	for _, val := range profile {
		fmt.Println("range by val", val)
	}

	str := "Привет мир"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}

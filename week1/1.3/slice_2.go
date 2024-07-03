package main

import "fmt"

func main() {
	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)

	sl1 := buf[1:4]
	sl2 := buf[:2]
	sl3 := buf[2:]
	fmt.Println(sl1, sl2, sl3)

	newBuf := buf[:]
	newBuf[0] = 9
	fmt.Println(buf, newBuf, sl2)

	newBuf = append(newBuf, 6)
	newBuf[0] = 1
	fmt.Println("buf", buf)
	fmt.Println("newBuf", newBuf)

	// my example
	testSlice1 := make([]int, 5, 5)
	testSlice2 := testSlice1[:]
	fmt.Println("Before: ", testSlice1, testSlice2)

	testSlice2 = append(testSlice2, 1)
	testSlice2[0] = 1
	fmt.Println("After: ", testSlice1, testSlice2)

	// end of my
	var emptyBuf []int            // l=0, c=0
	copied := copy(emptyBuf, buf) // copied = 0
	fmt.Println(copied, emptyBuf)

	newBuf = make([]int, len(buf), len(buf))
	copy(newBuf, buf)
	fmt.Println(newBuf)

	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6})
	fmt.Println(ints)
}

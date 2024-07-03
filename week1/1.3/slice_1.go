package main

import "fmt"

func main() {
	var buf0 []int // l=0, c=0
	buf1 := []int{} // l=0, c=0
	buf2 := []int{42} // l=1, c=1
	buf3 := make([]int, 0) // l=0, cap=0
	buf4 := make([]int, 5) // l=5, cap=5
	buf5 := make([]int, 5, 10) // len=5, cap=10

	fmt.Println(buf0, buf1, buf2, buf3, buf4, buf5)

	someInt := buf2[0]
	fmt.Println(someInt)

	var buf []int // len=0, cap=0
	buf = append(buf, 9, 10) // len=2, cap=2
	buf = append(buf, 12) // len=3, cap=4

	otherBuf := make([]int, 3)
	buf = append(buf, otherBuf...)
	fmt.Println(buf, otherBuf)

	var bufLen, bufCap int = len(buf), cap(buf)

	fmt.Println(bufLen, bufCap)
}

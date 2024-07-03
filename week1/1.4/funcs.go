package main

import "fmt"

func singleIn(in int) int {
	return in
}

func multIn(a, b int, c int) int {
	return a + b + c
}

func namedReturn() (out int) {
	out = 2
	return
}

func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some err")
	}
	return in, nil
}

func multipleNamedReturn(ok bool) (res int, err error) {
	res = 1
	if ok {
		err = fmt.Errorf("some err")
		return
	}
	res = 2
	return
}

func sum(in ...int) (result int) {
	fmt.Printf("in :=  %#v\n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	fmt.Println(multipleNamedReturn((false)))

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums, sum(nums...))
	return
}

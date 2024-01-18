package main

import "fmt"

func Pointer_sementics() {
	a := 1
	ptr := &a

	b := 100

	add(b)
	addPtr(ptr)
}

func add(x int) int {
	fmt.Println("addition :", x+x)
	return x + x
}
func addPtr(x *int) int {
	fmt.Println("addition :", *x+*x)
	return *x + *x
}

package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	x := 100
	y := &x
	p1 := person{"man", 20}
	ptr := &p1

	ptr.age = 30 // auto dereferancing with structs no need to use *ptr.age
	*y = 200     // must use *y only y can't be used same goes with slice,maps .

	inc(&x)

	// Pointer_sementics()
	Method_sets()
}

func inc(a *int) {
	*a = 200
	fmt.Println(*a)
}

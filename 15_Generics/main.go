package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type myNumber interface {
	constraints.Integer | constraints.Float
}

func add[T myNumber](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("heloo")
	fmt.Println(add(1, 3))
}

package main

import "fmt"

func main() {
	add(4, 4, "2", 2)
	Generics()

}

func add[A, B ~int, T any, F int](i A, j B, a T, b F) (T, F) {
	fmt.Printf("type of a is %T \n", a)
	fmt.Printf("type of b is %T \n", b)
	fmt.Printf("type of i is %T \n", i)
	fmt.Printf("type of j is %T \n", j)
	return a, b
}

// type equalizer[T any] interface {
// }

// func newEmptyList[T any]() []T {
// 	return make([]T, 0)
// }

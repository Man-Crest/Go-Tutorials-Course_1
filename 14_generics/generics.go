package main

import "fmt"

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type canWalk interface {
	walk()
}

type Person[T any] struct {
	name T
}

type Animal[T any] struct {
	name T
}

func (P Person[T]) walk() T {
	fmt.Printf("person %v can walk\n", P.name)
	return P.name
}

func (A Animal[T]) walk() T {
	fmt.Printf("Animal %v can walk\n", A.name)
	return A.name
}

func Add[T Num](a T, b T) T {
	return a + b
}

func Generics() {
	fmt.Println(Add(2, 3))

	P1 := Person{"man"}
	M1 := Animal{"monkey"}

	s := []canWalk{P1, M1}
	fmt.Println(s)
}

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (a Person) Age() {
	fmt.Println("age is ", a.age)
}
func (a *Person) AgePtr() {
	fmt.Println("age is ", a.age)
}

type Temp interface {
	Age()
	AgePtr()
}

func printPerson(t Temp) {
	t.Age()
	t.AgePtr()
}

func Method_sets() {

	p1 := Person{"man", 20}

	p1.Age()
	p1.AgePtr()
	// printPerson(p1) // if we have methods , one of them is using ptr as arguments then we can not pass value in such interface

	p2 := Person{"nil", 40}
	ptr := &p2

	p2.Age()
	p2.AgePtr()
	printPerson(ptr) //if methods having arguments as *int and int both then we have to use ptr to pass
}

package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Printf("person name is\t %v and age is \t %v \n", p.name, p.age)
}

func main() {
	///#59 variadic function

	defer variadic("Ancient magic", 1, 2, 3, 4, 5)

	s1 := []int{7, 8, 9}
	variadic("harry potter", s1...)

	//#61
	p1 := person{"man", 20}
	p1.speak()

}

func variadic(name string, chapters ...int) {
	fmt.Printf("name of book is \t %v \tand chapters are \t%v \n", name, chapters)
}

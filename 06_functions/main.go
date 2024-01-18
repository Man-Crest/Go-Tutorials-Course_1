package main

import "fmt"

func main() {
	// print("a", 1, 2, 3, 4, 5, 6)
	s1 := []int{1, 3, 4}
	//to pass a slice "s1..." is used when "...int "is used in function argument it is caled as variadic fuctions
	print("a", s1...)
	temp("b", s1)

	//anonymous functions
	// func(){}()
	func() {
		fmt.Println("hello")
	}()

	//anonymopus function with parameters
	func(name string) {
		fmt.Println(name)
	}("man")

	//functoin expression name := func(){}

	x := func() {
		fmt.Println("hiiii!!")
	}
	y := func(id int) {
		fmt.Printf("hiiii!! %v", id)
	}

	x()
	y(5)

	//function returning fuction

	fmt.Println("\ndoMath output is ", doMath(1, 3, add))
	fmt.Println("\ndoMath output is ", doMath(3, 1, subtract))
}

// ------ variadic functions ---------//

// to catch multiple arguments in a slice ...int is used
func print(sliceName string, slice ...int) {
	fmt.Printf("your slice is : %v \n %v", sliceName, slice)
}

func temp(sliceName string, slice []int) {
	fmt.Printf("your slice is : %v \n %v", sliceName, slice)
}

//function return function and takes arguments as function
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}

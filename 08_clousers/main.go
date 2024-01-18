package main

import "fmt"

func main() {
	temp := clouserFunc(1)
	fmt.Println(temp) // returns address

	temp()
	temp()
	temp()
	temp()

	fmt.Println("your factorial is", factorial(5))
}

func clouserFunc(x int) func() {
	return func() {
		x++
		fmt.Println(x)
	}
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

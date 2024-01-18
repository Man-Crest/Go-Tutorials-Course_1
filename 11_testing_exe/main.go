package main

import "fmt"

func main() {
	Add(2, 3)
}

func Add(a int, b int) int {
	fmt.Println("addition is :", a+b)
	return a + b
}

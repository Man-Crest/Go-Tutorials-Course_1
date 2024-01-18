package main

// func main() {
// 	arr := [5]int{1, 3, 6, 5, 2} ////composite array

// 	for i, val := range arr {
// 		fmt.Printf("index is %v \t and value is %v \n", i, val)
// 	}

// }
import (
	"fmt"
)

func main() {
	a := make([]int, 0, 10) //[]
	b := make([]int, 10)    // [0,0,0,0,0,0,0....]

	fmt.Println(a)
	fmt.Println(b)

	a = append(a, 11)
	b = append(b, 22)

	fmt.Println(a) //[11]
	fmt.Println(b) //[0,0,0,0,0,0....22]
}

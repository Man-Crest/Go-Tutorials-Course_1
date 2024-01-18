package main

import "fmt"

func main() {
	fmt.Println(Addfunc(1, 2, 3))
	fmt.Println(Addfunc(6, 3))
	fmt.Println(Addfunc(8, 0))
	fmt.Println(Addfunc(-1, 2, -33))

}

func Addfunc(s ...int) int {
	ans := 0
	for _, val := range s {
		ans = ans + val
	}
	return ans
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4, 21, 5, 3, 1}
	b := []int{4, 21, 5, 3, 1}
	fmt.Println("before sorting", a)
	fmt.Println("before sorting", b)
	sortSlice(a, b)
	fmt.Println("after sorting", a)
	fmt.Println("after sorting ", b)

}

// slice is always pointing to an array so slice can never be created without array , silece is pointing to the undelying array so when we pass slice to a function argument it sends the pointer not the copy of slice .

//  in general other types sends copy in function arguments so changes may seen only inside of function scope , not refllacted in main func , to change in actual value we have to pass a pointer of that type {"*"} , but slice behaviour is diff it always send pointe r.

// to avoid change in actual value we have to make a copy of slice using copy() method .

func sortSlice(x []int, y []int) {
	sort.Ints(x)

	i := make([]int, 5)
	copy(i, y)

	sort.Ints(i)

	fmt.Println("inside function ", x)
	fmt.Println("inside function ", i)
}

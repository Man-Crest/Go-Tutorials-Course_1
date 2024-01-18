package main

import "fmt"

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	random()
// 	random()
// 	random()
// 	random()
// 	random()
// }
// func random() {
// 	x := rand.Intn(500)

// 	if x < 100 {
// 		fmt.Println("values is less then 100")
// 	} else if x >= 100 && x <= 200 {
// 		fmt.Println("value is in between 100 and 200")
// 	} else if x > 200 && x <= 300 {
// 		fmt.Println("value is in between 201 and 300")
// 	} else {
// 		fmt.Println("value must lies between 301 and 500")
// 	}

// 	fmt.Printf("the value is %v \n", x)
// }

// package main

// import "fmt"

// func main() {

// 	for x := 0; x < 100; x++ {
// 		fmt.Printf("value is %v \n", x)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	z := 0
// 	for z <= 100 {
// 		fmt.Printf("value is %v \n", z)
// 		z++
// 	}
// }

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {

// 	for i := 0; i < 40; i++ {
// 		x := rand.Intn(5)
// 		fmt.Printf("value is %v and iteration is %v \n", x, i)
// 	}
// }

func main() {
	// x := make([]int, 20)
	// x := []int{5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 2}

	// for i, val := range x {
	// 	fmt.Printf("index is %v value is %v: \n", i, val)
	// }

	salaries := map[string]int{"man": 40000, "jemish": 50000, "ghno": 00}

	for i, salary := range salaries {
		fmt.Printf("salary of %v is : %v \n", i, salary)
	}
}

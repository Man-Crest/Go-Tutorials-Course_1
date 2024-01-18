// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	startTime := time.Now()

// 	x := factorial(1000)
// 	y := count(x)
// 	fmt.Println(<-y)
// 	elapsedTime := time.Since(startTime)
// 	fmt.Printf("Time taken with goroutines: %s\n", elapsedTime)
// }

// func factorial(n int) chan int {
// 	numbers := make(chan int)

// 	go func() {
// 		for i := 1; i <= n; i++ {
// 			numbers <- i
// 		}
// 		close(numbers)
// 	}()
// 	return numbers
// }

// func count(ch chan int) chan int {
// 	ans := make(chan int)
// 	go func() {
// 		var sum int = 1
// 		for c := range ch {
// 			sum = sum + c
// 			fmt.Println(sum)
// 		}
// 		ans <- sum
// 		close(ans)
// 	}()
// 	return ans
// }

package main

import (
	"fmt"
	"time"
)

func factorial(n int) int {
	fact := 1
	for i := 2; i <= n; i++ {
		fact += i
		fmt.Println(fact)
	}
	return fact
}

func main() {
	number := 1000
	startTime := time.Now()
	fact := factorial(number)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Factorial of %d is %d\n", number, fact)
	fmt.Printf("Time taken without goroutines: %s\n", elapsedTime)
}

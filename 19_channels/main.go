package main

import "fmt"

func main() {

	ch := make(chan int)
	q := make(chan int, 1)
	go input(ch, q)
	c(ch, q)
}

func c(ch <-chan int, q chan int) {
	// for val := range ch {
	// 	fmt.Println("here is your channel value", val)
	// }
	// _, ok := <-ch
	// if !ok {
	// 	fmt.Println("channel is closed now", ok)
	// }

	for {
		select {
		case v, ok := <-ch:
			fmt.Println("value got from a channel is :", v)
			if !ok {
				fmt.Println("channel is closed")
			}
		case v := <-q:
			fmt.Println("values are enough", v)
			close(q)
			return
		}
	}
}

func input(ch chan<- int, q chan<- int) {

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	q <- 0
}

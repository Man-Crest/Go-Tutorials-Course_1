package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quite := make(chan int)

	go send(eve, odd, quite)

	receve(eve, odd, quite)
}

func receve(e, o, q chan int) {
	for {
		select {
		// second parameter ok defines weather channel is closed or open
		case v, ok := <-e:
			if !ok {
				fmt.Println("channel is closed")
			} else {
				fmt.Println("data comming from even channel", v)
			}
		case v, ok := <-o:
			if !ok {
				fmt.Println("channel is closed")
			} else {
				fmt.Println("data comming from odd channel", v)
			}
		case v, ok := <-q:
			if !ok {
				fmt.Println("channel is closed")
			} else {
				fmt.Println("data comming from quite channel", v)
				close(q)
			}
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}

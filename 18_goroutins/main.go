package main

import (
	"fmt"
	"sync"
	"time"
)

var Wg = sync.WaitGroup{}

func main() {
	start := time.Now()
	respch := make(chan string)
	Wg.Add(2)
	a := userName()
	go userLikes(respch)
	user := <-respch
	fmt.Println(user)
	go userMatch(respch)
	user = <-respch
	fmt.Println(user)
	fmt.Println(a)
	fmt.Println(time.Since(start))
	Wg.Wait()
}
func userName() string {
	time.Sleep(time.Millisecond * 100)
	return "abc"
}
func userLikes(respch chan string) {
	// fmt.Println("500")
	respch <- "500"
	Wg.Done()
}
func userMatch(respch chan string) {
	// fmt.Println("xyz")
	respch <- "xyz"
	Wg.Done()
}

package main

import "fmt"

type temp struct {
	t int
}

func main() {
	a := 5
	ptr := &a

	fmt.Println("addres :", ptr, *ptr)
	t1 := temp{5}
	ptr1 := &t1

	fmt.Println("before value sementics", t1)
	editTemp(t1)
	fmt.Println("after value sementics", t1)
	editTempPtr(ptr1)
	fmt.Println("after ptr sementics", t1)
}

func editTemp(a temp) {
	a.t++
	fmt.Println("value sementics :", a.t)
}
func editTempPtr(a *temp) {
	a.t++
	fmt.Println("value sementics :", a.t)
}

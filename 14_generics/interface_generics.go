// package main

// import "fmt"

// type substractable interface {
// 	~int | any
// }

// type Person struct {
// 	name string
// }
// type Car struct {
// 	name string
// }

// type Movable[S substractable] interface {
// 	Move(S)
// }

// func (c Car) Move(meters substractable) {
// 	fmt.Printf("it moves %v meters", meters)
// }

// func Move[V Movable, S substractable](v V, meters S) {
// 	v.Move(meters)
// }

// func int_gen() {

// 	p1 := Person{"man"}
// 	c1 := Car{"g-vagon"}

// }

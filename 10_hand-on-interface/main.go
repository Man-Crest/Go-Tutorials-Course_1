package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type ractangle struct {
	l float64
	w float64
}

func (c circle) area() float64 {
	var area float64 = math.Pi * c.radius * c.radius
	fmt.Println(area)
	return area
}
func (r ractangle) area() float64 {
	var area float64 = r.l * r.w
	fmt.Println(area)
	return area
}

type shape interface {
	area() float64
}

func shapeInfo(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{
		2,
	}
	c2 := ractangle{
		2,
		3,
	}
	shapeInfo(c1)
	shapeInfo(c2)
}

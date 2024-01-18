package main

import "fmt"

func main() {
	p1 := person{
		"man",
		"patel",
		[]string{"a", "d", "c", "b"},
	}

	fmt.Println("struct p1 :", p1)
	fmt.Println("struct p1 :", p1.flavours[1])

	m1 := make(map[string]person)
	m1[p1.lname] = p1

	fmt.Println("map : ", m1)

	details := struct {
		first   string
		friends map[string]string
		drinks  []string
	}{
		"harsh",
		map[string]string{"a": "x", "b": "y", "c": "z"},
		[]string{"q", "r", "e", "w"},
	}

	fmt.Println("annonimus struct is :", details)
}

type person struct {
	fname    string
	lname    string
	flavours []string
}

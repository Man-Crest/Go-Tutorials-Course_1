package main

import (
	"fmt"
	"log"
	"os"
)

type person struct {
	name    string
	surname string
}

// stringer interface , it changes the way how the elem is printed
func (p person) String() string {
	return fmt.Sprintf("person is %s %s", p.name, p.surname)
}

func logInfo(s fmt.Stringer) {
	log.Println("log from 19", s.String())
}

func main() {
	p1 := person{
		"man",
		"patel",
	}
	p2 := person{
		"harsh",
		"patel",
	}

	// now i do println but it prints like this ("person is %s %s", p.name, p.surname") because we initialised stringer interface
	fmt.Println(p1)
	fmt.Println(p2)

	//to log the Println log is used instead of fmt
	log.Println(p1)
	log.Println(p2)

	logInfo(p1)
	logInfo(p2)

	//weriter interface

	file, err := os.Create("output.txt")

	defer file.Close()

	if err != nil {
		log.Fatalf("error %s", err)
	}

	s1 := []byte("hey there")
	_, err = file.Write(s1)

	if err != nil {
		log.Fatalf("error %s", err)
	}

}

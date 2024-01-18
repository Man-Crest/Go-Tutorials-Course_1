package main

import (
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	name string
	age  int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].age < a[j].age }

// by name sorting
// func (a byAge) Len() int           { return len(a) }
// func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a byAge) Less(i, j int) bool { return a[i].name < a[j].name }

func main() {

	p1 := person{"a", 1}
	p2 := person{"b", 2}
	p3 := person{"c", 3}
	p4 := person{"d", 4}
	p5 := person{"e", 5}

	people := []person{p1, p4, p5, p2, p3}
	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)

	pass := `asdf123`
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(hash))

	pass2 := `asdf123`
	err1 := bcrypt.CompareHashAndPassword(hash, []byte(pass2))
	if err1 != nil {
		fmt.Println("panic")
		panic(err1)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Number int    `json:"number"`
}

func main() {
	p1 := Person{"man", 20, 12121212}

	js, err := json.Marshal(p1)

	// encode through NewEncoder
	// err = json.NewEncoder(os.Stdout).Encode(p1)
	// 	if err != nil {
	// 		fmt.Println("error")
	// 	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(js))

	var people1 Person

	check := json.Valid(js)
	if check {
		json.Unmarshal(js, &people1)
		fmt.Println(people1)
	} else {
		fmt.Println("error in unmarshaling")
	}

}

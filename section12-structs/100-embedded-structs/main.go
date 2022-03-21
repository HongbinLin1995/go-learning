package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretagent struct {
	person
	ilk   bool
	first bool
}

func main() {
	p1 := secretagent{
		person: person{
			first: "Hongbin",
			last:  "Lin",
			age:   27,
		},
		ilk:   true,
		first: false,
	}
	fmt.Println(p1)
	fmt.Println(p1.person.first, p1.person.last, p1.person.age, p1.ilk, p1.first)
	p2 := secretagent{
		person: person{
			first: "Kaisi",
			last:  "Zhou",
			age:   26,
		},
		ilk:   false,
		first: true,
	}
	fmt.Println(p2)
	fmt.Println(p2.first, p2.last, p2.age, p2.ilk)
}

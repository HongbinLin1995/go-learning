package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, ", ", p.age)
}

func main() {
	p := person{
		first: "Hongbin",
		last:  "Lin",
		age:   27,
	}
	fmt.Println(p)
	p.speak()
}

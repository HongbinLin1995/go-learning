package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Hongbin",
		last:  "Lin",
		age:   27,
	}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
	p2 := person{
		first: "Kaisi",
		last:  "Zhou",
		age:   26,
	}
	fmt.Println(p2)
	fmt.Println(p2.first, p2.last, p2.age)
}

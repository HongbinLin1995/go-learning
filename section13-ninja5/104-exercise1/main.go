package main

import (
	"fmt"
)

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	p1 := person{
		first:    "Hongbin",
		last:     "Lin",
		icecream: []string{"chocolate", "martini", "rum and coke"},
	}
	p2 := person{
		first:    "Miss",
		last:     "Moneypenny",
		icecream: []string{"strawberry", "vanilla", "capuccino"},
	}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last)
	for i, v := range p1.icecream {
		fmt.Println(i, v)
	}
	fmt.Printf("\n")

	fmt.Println(p2)
	fmt.Println(p2.first, p2.last)
	for i, v := range p2.icecream {
		fmt.Println(i, v)
	}
}

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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, p := range m {
		fmt.Println("First name: ", p.first)
		fmt.Println("Last name: ", p.last)
		for i, p2 := range p.icecream {
			fmt.Println("\tFavorite icecream: ", i, p2)
		}
		fmt.Printf("\n")
	}
}

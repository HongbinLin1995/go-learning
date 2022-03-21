package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Hongbin", 64}
	p4 := person{"Kaisi", 56}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.SliceStable(people, func(i, j int) bool { return people[i].first < people[j].first })
	fmt.Println(people)
	sort.SliceStable(people, func(i, j int) bool { return people[i].age > people[j].age })
	fmt.Println(people)
}

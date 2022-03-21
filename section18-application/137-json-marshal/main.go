package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"json_Frist"`
	Last  string `json:"json_Last"`
	Age   int    `json:"json_Age"`
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	fmt.Println(p1.First, p1.Last, p1.Age)
	bs1, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs1))
	fmt.Printf("\n")

	fmt.Println(p2)
	bs2, err := json.Marshal(p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs2))
	fmt.Printf("\n")

	people := []person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", bs)
	fmt.Println(bs)
	fmt.Println(string(bs))
	fmt.Printf("\n")
}

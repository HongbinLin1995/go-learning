package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	s1 := `{"First":"James","Last":"Bond","Age":32}`
	bs1 := []byte(s1)
	s2 := `{"First":"Miss","Last":"Moneypenny","Age":27}`
	bs2 := []byte(s2)
	fmt.Printf("\n")
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Printf("\n")
	fmt.Println(s1)
	fmt.Println(bs1)
	fmt.Printf("\n")
	fmt.Println(s2)
	fmt.Println(bs2)
	fmt.Printf("\n")

	var p1 person
	var p2 person
	var p []person

	err := json.Unmarshal(bs, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Printf("\n")

	err = json.Unmarshal(bs1, &p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1)
	fmt.Printf("\n")

	err = json.Unmarshal(bs2, &p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
	fmt.Printf("\n")
}

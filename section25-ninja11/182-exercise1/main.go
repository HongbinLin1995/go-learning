package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string   `json:"json_Frist"`
	Last    string   `json:"json_Last"`
	Sayings []string `json:"json_Sayings"`
}

func main() {
	// start
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		fmt.Println("error")
	}
	fmt.Println(string(bs))

}

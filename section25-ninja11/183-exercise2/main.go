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
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("there was an error in toJSON: %v", err)
	}
	return bs, nil

}

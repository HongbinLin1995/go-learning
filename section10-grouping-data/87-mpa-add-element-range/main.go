package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Hongbin Lin": 27,
		"Kaisi Zhou":  26,
	}
	fmt.Println(m)
	fmt.Println(m["Hongbin Lin"])
	fmt.Println(m["Kaisi Zhou"])
	fmt.Println(m["Kaisi Lin"])
	fmt.Printf("\n")

	m["Kaisi Lin"] = 7
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}

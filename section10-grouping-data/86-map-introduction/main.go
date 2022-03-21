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
	fmt.Println(m["Unknown"])
	fmt.Printf("\n")

	v, ok := m["Hongbin Lin"]
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Printf("\n")

	v, ok = m["Unknown"]
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Printf("\n")

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("This is the IF print", v)
		fmt.Printf("\n")
	}

	if v, ok := m["Hongbin Lin"]; ok {
		fmt.Println("This is the IF print", v)
		fmt.Printf("\n")
	}

	n := map[int]string{
		27: "Hongbin Lin",
		26: "Kaisi Zhou",
	}
	fmt.Println(n)
	fmt.Printf("\n")

	a := []int{1, 2, 3}
	fmt.Println(a)
	u := append(a[0:1], a[1:2]...)
	fmt.Println(u)

}

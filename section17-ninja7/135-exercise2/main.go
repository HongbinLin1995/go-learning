package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	fmt.Println(p)
	p.name = "Zhoukai Si"

}

func main() {
	p := person{
		name: "Hongbin Lin",
	}
	fmt.Println(p)
	fmt.Println(&p)
	changeMe(&p)
	fmt.Println(p.name)
}

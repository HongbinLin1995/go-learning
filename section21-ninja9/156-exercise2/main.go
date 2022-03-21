package main

import "fmt"

type person struct {
	First string
	Last  string
}

func (p *person) speak() {
	fmt.Println("Hello")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		First: "Hongbin",
		Last:  "Lin",
	}
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Printf("\n")

	saySomething(&p) // work
	// saySomething(p)  // Does not work
	p.speak()
}

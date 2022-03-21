package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c := circle{5}
	// info(c)
	fmt.Println(c.area())
}

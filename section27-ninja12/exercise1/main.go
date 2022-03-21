package main

import (
	"fmt"

	"github.com/HongbinLin1995/go-learning/section27-ninja12/exercise1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}

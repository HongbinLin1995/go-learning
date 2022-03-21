package main

import (
	"fmt"
	"os"
)

var y int = 42
var name string = "Hongbin Lin"

func main() {
	fmt.Print(y)
	fmt.Println(y)
	fmt.Printf("%v\n", y)
	fmt.Printf("%#v\n", y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%s\t%b\t%x\n", name, y, y)
	fmt.Println(s)
	fmt.Printf("%v\n\n", y)

	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")
}

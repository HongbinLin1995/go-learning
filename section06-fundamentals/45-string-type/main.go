package main

import (
	"fmt"
)

func main() {
	s := "Hello, Hongbin Lin"
	fmt.Println(s)
	s = "Hello, Amazing Lin"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(len(s))

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(bs[0], bs[1])

	c := string(bs[0])
	fmt.Println(c)
	fmt.Printf("\n")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Printf("\n")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}
	fmt.Printf("\n")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}
	fmt.Printf("\n")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %q\n", i, v)
	}
	fmt.Printf("\n")

	p := "Hello, 世界"
	fmt.Printf("%s\n", p)
	fmt.Printf("%q\n", p)
	fmt.Printf("%x\n", p)
	fmt.Printf("%x\n", "Hello, 世")
	for i := 0; i < len(p); i++ {
		fmt.Printf("%x ", p[i])
	}
	fmt.Printf("\n")
}

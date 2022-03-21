package main

import (
	"fmt"
)

func main() {

	// Wrong code!!!
	// x := make([]string, 50, 52)
	// fmt.Println("Frist time(x)")
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	// fmt.Println("Second time(x)")
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// for i := 0; i < len(x); i++ {
	// 	fmt.Println(i, x[i])
	// }

	// Correct code!!!
	y := make([]string, 50, 52)
	fmt.Println("Frist time(y)")
	fmt.Println(len(y))
	fmt.Println(cap(y))

	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", states)

	var s = make([]int, 3)
	fmt.Println(s)
	_ = copy(y, states)
	fmt.Println("Second time(y)")
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
	for i := 0; i < len(y); i++ {
		fmt.Println(i, y[i])
	}
}

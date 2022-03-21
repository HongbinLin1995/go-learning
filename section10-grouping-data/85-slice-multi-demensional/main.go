package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	te := []string{"111", "222", "333"}
	fmt.Println(te)

	xp := [][]string{jb, mp, te}
	fmt.Println(xp)

	a := [][]int{{1, 2, 3}, {4, 5}}
	fmt.Println(a)

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	_ = append(x[:2], x[5:]...) // the same underlying array stores the value of the new slice
	fmt.Println(x)
}

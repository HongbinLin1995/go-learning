package main

import (
	"testing"
)

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	test1 := test{[]int{21, 21}, 42}
	test2 := test{[]int{3, 4, 5}, 12}
	test3 := test{[]int{1, 1}, 2}
	test4 := test{[]int{-1, 0, 1}, 0}
	tests := []test{test1, test2, test3, test4}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

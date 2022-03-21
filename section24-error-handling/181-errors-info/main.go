package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("Solution 5: a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	value, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(value)
}

func sqrt(f float64) (float64, error) {
	// var ErrNorgateMath = errors.New("solution 2: norgate math: square root of negative number")
	if f < 0 {
		// return 0, errors.New("solution 1: norgate math: square root of negative number")  // solution 1

		// var ErrNorgateMath = errors.New("solution 2: norgate math: square root of negative number") // solution 2
		// return 0, ErrNorgateMath

		// return 0, fmt.Errorf("solution 3: norgate math again: square root of negative number: %v", f) // solution 3

		// ErrNorgateMath := fmt.Errorf("solution 4: norgate math again: square root of negative number: %v", f) //solution 4
		// return 0, ErrNorgateMath
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f) //solution 5
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}

	} else {
		return f * f, nil
	}
}

package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("lat:%v\tlong:%v\t\terr:%v", se.lat, se.long, se.err)
}

func main() {
	area, err := sqrt(-10.23)
	fmt.Printf("%T\n", area)
	fmt.Printf("%T\n", err)
	fmt.Println(area)
	fmt.Println(err)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("Wrong")
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}

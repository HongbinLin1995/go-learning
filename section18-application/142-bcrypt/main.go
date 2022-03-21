package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs := []byte(s)
	hs, err := bcrypt.GenerateFromPassword(bs, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(hs)
	fmt.Printf("\n")

	loginPword1 := `password1234`
	bloginPword1 := []byte(loginPword1)
	fmt.Println(bloginPword1)

	err = bcrypt.CompareHashAndPassword(bs, []byte(bloginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")
}

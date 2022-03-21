package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	// 01
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Printf("\n")

	//02
	var answer1, answer2, answer3 string
	fmt.Print("Name: ")
	s1, err := fmt.Scan(&answer1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Fav Food: ")
	s2, err := fmt.Scan(&answer2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Fav Sport: ")
	s3, err := fmt.Scan(&answer3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s1, s2, s3)
	fmt.Println(answer1, answer2, answer3)
	fmt.Printf("\n")

	//03
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer f.Close()
	r := strings.NewReader("James Bond")
	fmt.Println(r)
	io.Copy(f, r)
	f.Close()
	fmt.Printf("\n")

	//04
	f, err = os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bs)
	fmt.Println(string(bs))
	f.Close()
}

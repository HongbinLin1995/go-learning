package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f1, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(f1)

	defer foo()

	f2, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened:-> ", err)
		log.Println("err happened:-> ", err)
		// log.Fatalln(err)
		log.Panicln(err)
		// panic(err)
	}

	f2.Close()
	f1.Close()
	fmt.Println("check the log.txt file in the directory")
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}

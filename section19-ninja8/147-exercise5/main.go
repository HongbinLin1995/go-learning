package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"B - Shaken, not stirred",
			"A - Youth is no guarantee of innovation",
			"C - In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"L - James, it is soo good to see you",
			"M - Would you like me to take care of that for you, James?",
			"A - I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"O - Oh, James. You didn't.",
			"D - Dear God, what has James done now?",
			"C - Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("-------")

	sort.SliceStable(users, func(i, j int) bool { return users[i].Age < users[j].Age })
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("-------")

	sort.SliceStable(users, func(i, j int) bool { return users[i].Last < users[j].Last })
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

}

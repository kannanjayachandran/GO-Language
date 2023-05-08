package main

import "fmt"

func main() {
	var age int
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("You can vote..")
	} else {
		fmt.Println("You cannot vote..")
	}
}

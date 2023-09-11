package main

import (
	"fmt"
)

func main() {

	// for loop
	fmt.Println("For loop")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// while loop
	fmt.Println("While loop")
	x := 10
	for x > 0 {
		fmt.Println(x)
		x--
	}

	// looping over a string
	fmt.Println("\nstring looping")
	str := "Hello GO Lang"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

	// for each
	sentence := "This is a sentence."
	for _, char := range sentence {
		fmt.Println(string(char))
	}
}

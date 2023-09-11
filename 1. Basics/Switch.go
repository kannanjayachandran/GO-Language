package main

import "fmt"

func main() {

	a := 10

	// normal switch statement
	switch a {
	case 1:
		fmt.Println("case1")

	case 10:
		fmt.Println("Case2")

	default:
		fmt.Println("Default switch")
	}

	// switch with passthrough
	switch {
	case 1 < 2:
		fmt.Println("Hello GO")
		fallthrough

	case 100 > 10:
		fmt.Println("Passed through")

	case 100 < 10:
		fmt.Println("null")

	}
}

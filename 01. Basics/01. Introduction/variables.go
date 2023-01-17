package main

import "fmt"

// Constants using iota (Incremental values)
const (
	CAR = iota
	BUS
	TRAIN
	FLIGHT
)

func main() {

	// integer
	var i int
	i = 100
	fmt.Println(i)
	fmt.Println()

	j := 200
	fmt.Println(j)
	fmt.Println()

	// boolean
	var flag bool = false

	// if else
	if flag {
		fmt.Print("_")
	} else {
		fmt.Println("This is boolean variable")
	}
	fmt.Println()

	// for loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// we can run while loop also like this
	k := 0
	for k < 10 {
		fmt.Println(k + 100)
		k++
	}
	fmt.Println()

	fmt.Println(CAR)
	fmt.Println(BUS)
	fmt.Println(TRAIN)
	fmt.Println(FLIGHT)
}

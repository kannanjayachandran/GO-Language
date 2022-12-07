package main

import "fmt"

func main() {

	// integer
	var i int
	i = 100
	fmt.Println(i)

	j := 200
	fmt.Println(j)

	// boolean
	var flag bool = false

	// if else
	if flag {
		fmt.Print("_")
	} else {
		fmt.Println("This is boolean variable")
	}

	// for loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// we can run while loop also like this
	k := 0
	for k < 10 {
		fmt.Println(k + 100)
		k++
	}

}

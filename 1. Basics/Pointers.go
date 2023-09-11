package main

import "fmt"

func main() {

	var x int = 12345
	y := &x

	fmt.Println(x, y)

	*y = 120
	fmt.Println(x, y)
}

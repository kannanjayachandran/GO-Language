package main

import (
	"fmt"
)

func main() {

	// output
	x := 10000000
	y := 50
	z := 10.2345
	fmt.Println("{ Using println } ", y, z, x)
	fmt.Printf("\nValue : %v --- Type : %T\n", z, z)
	fmt.Printf("\nValue : %v --- Binary_representation : %b\n", y, y)
	fmt.Printf("\nValue : %v --- Scientific_notation : %e\n", z, z)
	fmt.Printf("\nValue : %v --- Rounding_float : %.1f\n", z, z)
	str := fmt.Sprintf("\n%v - %T", x, x)
	fmt.Println(str)

	// input
	var name string
	fmt.Println("Enter your name :")
	fmt.Scan(&name)
	fmt.Println("Your name is :", name)
}

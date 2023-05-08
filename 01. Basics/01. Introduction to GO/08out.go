package main

import "fmt"

func main() {
	x := 100
	y := 50
	z := 10.234
	fmt.Println("Hello world")
	fmt.Printf("Type: %T \nValue: %v\nBinary rep: %b\nScientific Notation: %e\nRounding float: %.2f", x, y, x, z, z)
	fmt.Println("\n%50")

	// Sprintf
	str := fmt.Sprintf("%T", z)
	fmt.Println(str)
}

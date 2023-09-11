package main

import "fmt"

func main() {

	// numbers
	var a int = 120
	var b int8 = 10
	var c int16 = -45
	var d int32 = 53
	var e int64 = 100

	var f uint = 5
	var g uint32
	var h rune = 45

	var i float32 = 1.5
	var j float64 = 34.323

	fmt.Println("Integers Data types: \t", a, b, c, d, e, f, g, h)
	fmt.Println("Floats types: \t", i, j)

	// Complex
	var k complex128 = complex(5, 6)
	var l complex64 = complex64(1)
	fmt.Println(k, l)

	// boolean
	var isTrue bool = true
	var isFalse bool = false
	fmt.Println("Boolean values:\t", isTrue, isFalse)
	fmt.Printf("a = %v, b = %v, a == b = %v\n", isTrue, isFalse, isTrue == isFalse)

	// Char and string
	var letter byte = 'K'
	var sentence string = "This is a string in go lang"

	fmt.Println(letter, sentence)

	// implicit variable assignment
	x := 100
	y := "This is implicit assignment."

	fmt.Println(x, y)
}

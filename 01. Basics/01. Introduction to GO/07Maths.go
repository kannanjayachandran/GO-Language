package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	a := 1000000
	b := int8(1)

	c := a + int(b)

	fmt.Println("Sum: ", c)
	fmt.Println("Max function: ", math.Max(10, 100))
	fmt.Println("Square root function: ", math.Sqrt(900))
	fmt.Println("Power function: ", math.Pow(10, 3))

	// String to integer
	str := "12345"
	str1 := "Hello"
	num, err := strconv.Atoi(str)
	num1, err1 := strconv.Atoi(str1)

	fmt.Println(num, err)
	fmt.Println(num1, err1)

	str2 := "12345"
	num2, err2 := strconv.ParseInt(str2, 10, 32) // str, base, bytes
	fmt.Println(num2, err2)

}

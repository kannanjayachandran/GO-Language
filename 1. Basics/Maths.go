package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func main() {

	a := 1000000
	b := int8(1)
	c := a + int(b)
	var nonNegative = math.Dim(6, 3)

	fmt.Println("Sum: ", c)
	fmt.Println("Max function: ", math.Max(10, 100))
	fmt.Println("Square root function: ", math.Sqrt(900))
	fmt.Println("Power function: ", math.Pow(10, 3))
	fmt.Println(nonNegative)

	// parsing strings to integers
	str := "12345"
	str1 := "Hello"
	num, err := strconv.Atoi(str)
	num1, err1 := strconv.Atoi(str1)
	str2 := "-342"

	fmt.Println(num, err, "Type of num is", reflect.TypeOf(num))
	fmt.Println(num1, err1)

	num2, err2 := strconv.ParseInt(str2, 10, 16) // str, base, bytes
	fmt.Println(num2, err2)
}

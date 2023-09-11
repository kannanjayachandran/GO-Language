package main

import (
	"errors"
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Division by zero")
	}
	return x / y, nil
}

// args
func sumInts(nums ...int) int {
	sum := 0

	for _, val := range nums {
		sum += val
	}
	return sum
}

// HOF
func addOne(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func callFunc(callable func(string) string, arg string) {
	res := callable(arg)
	fmt.Println(res)
}

func main() {

	fmt.Println("Addition : ", add(100, 200))
	res, err := divide(7, 0)

	if err != nil {
		fmt.Println("Error !", err)
	} else {
		fmt.Println("Ans : ", res)
	}

	total := sumInts(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(total)

	innerFunc := addOne(65)
	val := innerFunc(65)
	fmt.Println(val)

	myFunc := func(str string) string {
		return str + "world!"
	}
	callFunc(myFunc, "Hello ")
}

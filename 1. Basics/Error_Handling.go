package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

func main() {

	ans, err := divide(10, 2)

	if err == nil {
		fmt.Println(ans)
	} else {
		fmt.Println(err)
	}
}

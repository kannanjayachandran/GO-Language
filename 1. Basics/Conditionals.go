package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	} else if num == 2 || num == 3 {
		return true
	} else if num%2 == 0 || num%3 == 0 {
		return false
	}
	count := 5
	for count*count <= num {
		if num%count == 0 || num+2%count == 0 {
			return false
		}
		count += 6
	}
	return true
}

func main() {
	fmt.Println("Enter the number to check: ")
	ans := isPrime(17)
	fmt.Println(ans)
}

package main

import "fmt"

func main() {

	// Initializing a map with string key and int values
	dict := make(map[string]int)

	fmt.Println(dict)
	dict["Go"] = 1
	dict["Python"] = 2
	dict["Java"] = 3
	dict["JS/TS"] = 4
	fmt.Println("The map is :", dict)

	// Accessing elements
	fmt.Println("Accessing elements :", dict["Go"])

	// Updating element
	dict["Rust"] = 5
	dict["Python"] = 1
	fmt.Println("The modified map is :", dict)

	// deleting element
	delete(dict, "Rust")
	fmt.Println("The modified map is :", dict)

	// checking existence
	val, isIn := dict["Rust"]
	if isIn {
		fmt.Println(val, "exists")
	} else {
		fmt.Println("Entered key does not exits!")
	}

	// iterating over map
	for key, val := range dict {
		fmt.Println(key, ":", val)
	}

	// another example
	dict2 := make(map[string][]int)
	dict2["First"] = []int{1, 2, 3, 4, 5}
	fmt.Println("New map : ", dict2)

	// between 1 and n find all the multiples of 1, 2, 3, 4, 5
	n := 100
	divisibleMap := make(map[int]uint)

	for num := 1; num <= n; num++ {
		for x := 1; x <= 5; x++ {
			if num%x == 0 {
				divisibleMap[x]++
			}
		}
	}
	fmt.Println("Divisibility Map :", divisibleMap)
}

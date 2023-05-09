package main

import "fmt"

func main() {

	// array
	var arr [10]int

	fmt.Println(arr)

	for i := 0; i < 10; i++ {
		arr[i] = i + 1
	}

	for _, value := range arr {
		fmt.Printf("%d ", value)
	}

	// another way to create array
	arr2 := [5]int{5, 4, 3, 2, 1}
	fmt.Println("\n", len(arr2))

	// automatic size
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	test_arr := [2][2]string{{"One", "Two"}, {"Three", "Four"}}

	// passing as reference/copy - won't change the original array
	test(test_arr)
	fmt.Println(test_arr)

	// Slices
	slice_arr := test_arr[:1]
	fmt.Println(slice_arr)
	fmt.Println(len(slice_arr), cap(slice_arr))

	// Slice would point at the underlying array itself. If we modify the slice it would also change the original array.
	// Slices are dynamic in nature
	arr4 := []string{}

	for i := 0; i < 10; i++ {
		arr4 = append(arr4, "Ben 10")
		fmt.Println(arr4, len(arr4), cap(arr4))
	}

	// Creating slice using make()
	arr5 := make([]int, 5, 10) // len=5, cap = 10
	fmt.Println(arr5)

	arr6 := []bool{}
	arr7 := []bool{true, false, false, true}
	arr6 = append(arr6, arr7...)
	fmt.Println(arr7)

}

func test(test_arr [2][2]string) {
	test_arr[0][0] = "New edit"
}

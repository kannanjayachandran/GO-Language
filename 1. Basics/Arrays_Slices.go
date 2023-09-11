package main

import "fmt"

func main() {

	// array
	var arr [10]int
	fmt.Println("Default array : ", arr)

	// Adding elements	// len(arr) | cap(arr)
	for i := 0; i < cap(arr); i++ {
		arr[i] = i + 1
	}
	fmt.Println("New array : ", arr)

	for _, value := range arr {
		fmt.Printf("%d ", value)
	}

	// another way to create array
	arr2 := [5]int{10, 23, 3, 5, 9}
	fmt.Println("\nThe new array is :", arr2, "\nIts length is : ", len(arr2))

	// a fixed size array, fixed size
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Size inferred array: ", arr3)

	// String array
	test_arr := [2][2]string{
		{"Go", "two"}, {"dimensional", "array"},
	}
	fmt.Println(test_arr)

	// passing as value / copy - won't change the original array
	changeArray(test_arr)
	fmt.Println(test_arr)

	// passing the original array { passing by reference }
	changeArray2(&test_arr)
	fmt.Println(test_arr)

	// Slices
	// Slice would point at the underlying array itself. If we modify the slice it would also change the original array.
	// Slices are dynamic in nature
	arr4 := []string{}
	fmt.Println("\n\nSlice ", arr4)

	for i := 0; i < 10; i++ {
		arr4 = append(arr4, "#")
		fmt.Println(arr4, " Len:", len(arr4), " Cap:", cap(arr4))
	}

	// slice from array
	slice_arr := test_arr[:1]
	fmt.Println("\nSlice from array:", slice_arr)
	fmt.Println("len of slice", len(slice_arr), "\ncapacity of slice", cap(slice_arr))

	// Creating slice using make()
	arr5 := make([]int, 5, 10) // len=5, cap = 10
	fmt.Println("\nSlice using make : ", arr5)

	arr6 := []bool{}
	arr7 := []bool{true, false, false, true}
	arr6 = append(arr6, arr7...)
	fmt.Println(arr7)

}

func changeArray(test_arr [2][2]string) {
	test_arr[0][0] = "Something"
}

func changeArray2(arr *[2][2]string) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			arr[i][j] = "*"
		}
	}
}

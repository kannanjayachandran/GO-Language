package main

import "fmt"

type Book struct {
	title         string
	numberOfPages int
}

func changeTitle(book *Book, title string) {
	book.title = title
}

func change(x *int) {
	*x = 0
}

func main() {

	var x int = 12345
	var y *int = &x
	z := &x

	fmt.Println("x is ", x)
	fmt.Println("The pointer is : ", y, z)

	*y = 120
	fmt.Println("Now x is ", x)

	change(y)
	fmt.Println("Now x is ", x)

	// pointer to a pointer
	lis := []int{1, 2, 3, 4, 5}
	a := &lis
	b := &a

	fmt.Println(lis, a, b)
	fmt.Println(*a, *b)

	// struct pointers
	var book Book = Book{"Joe in wonderland", 723}
	fmt.Println(book)
	changeTitle(&book, "King kong in wonderland")
	fmt.Println(book)

	// slice of pointers
	one := 1
	two := 2

	lis2 := []*int{&one, &two}
	fmt.Println(lis2)
}

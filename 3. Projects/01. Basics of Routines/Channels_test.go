package main

import "fmt"

func sum(c chan int) {

	// The value of x is whatever comes through the channel c
	x := <-c
	// Again the next thing that comes through the channel would go to y
	y := <-c

	c <- x + y
}

func main() {

	// making a channel
	c := make(chan int)

	go sum(c)

	// feeding data into channel
	c <- 100
	c <- 200

	r := <-c

	fmt.Println(r)
}

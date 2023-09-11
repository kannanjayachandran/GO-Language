package main

import (
	"fmt"
	"time"
)

func add(x, y int, returnVal chan int) {
	time.Sleep(1 * time.Second)
	returnVal <- x + y
}

func main() {

	returnChan := make(chan int)
	go add(10, 20, returnChan)
	ans := <-returnChan

	fmt.Println(ans)
}

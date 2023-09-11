package main

import (
	"fmt"
	"time"
)

func run1() {
	time.Sleep(1 * time.Second)
	fmt.Println("Run1 Completed")
}
func run2() {
	time.Sleep(2 * time.Second)
	fmt.Println("Run2 Completed")
}
func run3() {
	time.Sleep(3 * time.Second)
	fmt.Println("Run3 Completed")
}

func main() {

	go run1()
	go run2()
	go run3()

	time.Sleep(4 * time.Second)
}

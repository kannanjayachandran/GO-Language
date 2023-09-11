package main

import "fmt"

func doPanic() {
	panic("Fail ~ Panic occurred")
}

func handlePanic() {
	r := recover()
	fmt.Println(r)
}

func main() {

	fmt.Println("Start")
	defer handlePanic()
	doPanic()
	fmt.Println("End")
}

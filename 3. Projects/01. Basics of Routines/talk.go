package main

import (
	"fmt"
	"time"
)

func talk(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	go talk("Hello")
	talk("Go")
}

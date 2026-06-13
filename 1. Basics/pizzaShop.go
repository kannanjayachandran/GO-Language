package main

import (
	"fmt"
	"time"
)

// Worker 1: The Baker
func makePizza(pizzaChan chan string) {
	for i := 1; i <= 3; i++ {
		pizzaName := fmt.Sprintf("Pizza #%d", i)
		fmt.Printf("🍕 Baking %s...\n", pizzaName)
		time.Sleep(1 * time.Second) // Simulating baking time
		
		// Sending the pizza into the channel (conveyor belt)
		pizzaChan <- pizzaName 
	}
	
	// Crucial: Close the channel when done so the packer knows no more pizzas are coming
	close(pizzaChan) 
}

// Worker 2: The Main Chef / Packer
func main() {
	// Create a channel that moves strings (pizza names)
	pizzaChan := make(chan string)

	// Spin up the baker goroutine using the 'go' keyword
	go makePizza(pizzaChan)

	// The main function acts as the packer. 
	// It waits and receives pizzas from the channel.
	for pizza := range pizzaChan {
		fmt.Printf("📦 Packed and delivered %s!\n", pizza)
	}

	fmt.Println("🎉 All pizzas delivered! Shop is closing.")
}

// When the main function tries to pull a pizza from the channel, 
// but the baker hasn't finished baking it yet, main will automatically pause and wait. 
// It doesn't waste CPU power; it just waits patiently until a pizza shows up on the conveyor belt.
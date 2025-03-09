package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Generate a random number between 1 and 100
	randomNumber := rand.Intn(100) + 1

	// Print the result
	fmt.Println("The random number is:", randomNumber)
}

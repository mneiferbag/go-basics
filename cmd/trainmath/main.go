package main

import (
	"fmt"
	"math/rand"
)

func randomInts() (int, int) {
	return rand.Intn(100), rand.Intn(100)
}

func main() {
	for {
		a, b := randomInts()
		fmt.Printf("a = %d, b = %d\n", a, b)

		var sum int
		fmt.Print("Enter sum: ")
		fmt.Scanln(&sum)

		if sum == (a + b) {
			fmt.Println("Correct")
		} else {
			fmt.Printf("Incorrect (%d)\n", (a + b))
			break
		}
	}
}

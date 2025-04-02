package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	const maxGuessCount = 5
	goal := rand.IntN(100)
	guessCount := 0
	var guess int

	for guessCount < maxGuessCount {
		guessesLeft := maxGuessCount - guessCount
		fmt.Printf("You have %d guesses left\n", guessesLeft)
		fmt.Print("Your guess: ")

		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue // Don't count invalid inputs
		}

		guessCount++

		if guess > goal {
			fmt.Println("Too high")
		} else if guess < goal {
			fmt.Println("Too low")
		} else {
			fmt.Printf("You won! The number was %d\n", goal)
			return
		}

		fmt.Println() // Add a blank line between guesses
	}

	fmt.Printf("Sorry, you lost. The number was %d\n", goal)
}

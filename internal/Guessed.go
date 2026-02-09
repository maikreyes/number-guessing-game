package internal

import (
	"fmt"
	"math/rand"
)

func Guessed(attemps int) error {

	number := rand.Intn(100)

	var userNumber int

	switch attemps {
	case 10:
		fmt.Println(`Great! You have selected the Easy difficulty level.
Let's start the game!`)
	case 5:
		fmt.Println(`Great! You have selected the Medium difficulty level.
Let's start the game!`)
	case 3:
		fmt.Println(`Great! You have selected the Hard difficulty level.
Let's start the game!`)

	}

	for i := 0; i < attemps; i++ {

		fmt.Print("Enter your number: ")
		fmt.Scan(&userNumber)

		if userNumber != number {
			if userNumber > number {
				fmt.Printf("\nIncorrect! The number is less than %d\n", userNumber)
				continue
			}
			fmt.Printf("\nIncorrect! The number is greater than %d\n", userNumber)
			continue
		}

		fmt.Printf("\nCongratulations! You guessed the correct number in %d attempts\n", i+1)
		break
	}

	return nil
}

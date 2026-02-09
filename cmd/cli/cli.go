package cli

import (
	"fmt"
	"number-guessing-game/internal"
)

func Run() {

	var option int

	welcomeMessage()

	fmt.Print("\nEnter your choice: ")
	fmt.Scanln(&option)

	switch option {
	case 1:
		internal.Guessed(10)
	case 2:
		internal.Guessed(5)
	case 3:
		internal.Guessed(3)
	default:
		fmt.Println("Please select an option")
	}

}

func welcomeMessage() {

	fmt.Println(`Welcome to the Number Guessing Game!

I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)`)

}

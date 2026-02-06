package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ncruces/zenity"
)

var MAX_GUESSES int = 10

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)

	var guesses int = 0
	var matching = false
	var title = "Guess"

	for !matching && guesses < MAX_GUESSES {

		input, err := zenity.Entry("Enter your guess:",
			zenity.Title(title))

		if err != nil {
			break
		}
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input")
			break
		}

		guesses++

		if guess > r {
			title = fmt.Sprint("Last guess was too high (guesses ", guesses, ")")
		} else if guess < r {
			title = fmt.Sprint("Last guess was too low (guesses ", guesses, ")")
		} else if guess == r {
			matching = guess == r
		}
	}

	if matching {
		zenity.Info(fmt.Sprint("It took you ", guesses, " guesses!"), zenity.Title("Guess"))
	} else {
		zenity.Info(fmt.Sprint("You failed! The number was ", r, "!"), zenity.Title("Guess"))
	}
}

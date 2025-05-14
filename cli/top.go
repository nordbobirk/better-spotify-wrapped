package cli

import (
	"better-spotify-wrapped/internal"
	"fmt"
	"strconv"
)

// Top handles execution of the "top" CLI command.
func Top(args []string) {
	if len(args) < 2 {
		fmt.Println("Please specify 'track' or 'artist' and an amount as argument.")
		return
	}

	isValidAmount, parsedAmount := isValidAmount(args[1])

	if !isValidAmount {
		fmt.Println("Invalid amount argument, please specify an integer between 1 and 10.")
		return
	}

	switch args[0] {
	case "track":
		internal.TopTracks(parsedAmount)
	case "artist":
		internal.TopArtists(parsedAmount)
	default:
		fmt.Println("Please specify 'track' or 'artist' as argument.")
	}
}

func isValidAmount(amount string) (valid bool, parsedAmount int) {
	if amount == "" {
		return false, 0
	}

	parsedAmount, err := strconv.Atoi(amount)

	if err != nil {
		return false, 0
	}

	return parsedAmount >= 1 && parsedAmount <= 10, parsedAmount
}

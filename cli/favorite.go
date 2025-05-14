package cli

import (
	"better-spotify-wrapped/internal"
	"fmt"
)

// Favorite handles execution of the "favorite" CLI command.
func Favorite(args []string) {
	fmt.Println("Analyzing data to find your favorite track...")
	internal.Favorite()
}

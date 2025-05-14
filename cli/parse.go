package cli

import (
	"better-spotify-wrapped/internal"
	"fmt"
)

// Parse handles execution of the "parse" CLI command.
func Parse(args []string) {
	fmt.Println("Parsing Spotify data...")
	internal.ParseData()
}

package cli

import (
	"fmt"
)

// Help prints the help message to stdout.
func Help() {
	fmt.Println("Usage: \n bsw [command] [options]")
	fmt.Println("Available Commands:")
	fmt.Println("  parse - Parse raw Spotify listening data to an aggregated JSON file. Expects the Spotify data in the data folder (created automatically in the first run). This command must be run before any other commands.")
	fmt.Println("  favorite - Show information about your favorite track.")
	// TODO: Birk extend this
}

package cli

import (
	"fmt"
)

// Help prints the help message to stdout.
func Help() {
	fmt.Println("Usage: \n bsw [command] [options]")
	fmt.Println("Available Commands:")
	fmt.Println("  parse - Parse raw Spotify listening data to an aggregated JSON file. Expects the Spotify data in the data folder (created automatically in the first run). This command must be run before any other commands.")
	fmt.Println("  favorite <artist|track> - Show information about your favorite artist ortrack.")
	fmt.Println("  top <artist|track> <amount> - Show the top artists or tracks by listening time.")
	fmt.Println("  hate <artist|track> - Show information about your least favorite artist or track.")
	fmt.Println("  count <artist|track> - Count the number of different artists or tracks you have listened to.")
	// TODO: Birk extend this
}

package cli

import (
	"better-spotify-wrapped/internal"
	"fmt"
)

// Hate handles execution of the "hate" CLI command.
func Hate(args []string) {
	if len(args) == 0 {
		fmt.Println("Please specify 'artist' or 'track' as argument.")
		return
	}

	switch args[0] {
	case "artist":
		internal.HateArtist()
	case "track":
		internal.HateTrack()
	default:
		fmt.Println("Please specify 'artist' or 'track' as argument.")
	}
}

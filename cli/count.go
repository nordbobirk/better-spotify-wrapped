package cli

import (
	"better-spotify-wrapped/internal"
	"fmt"
	"strings"
)

func Count(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide 'artist' or 'track' as an argument.")
		return
	}

	switch strings.ToLower(args[0]) {
	case "artist":
		internal.CountArtists()
	case "track":
		internal.CountTracks()
	default:
		fmt.Println("Please provide 'artist' or 'track' as an argument.")
	}
}

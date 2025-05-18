package cli

import (
	"better-spotify-wrapped/internal"
	"fmt"
	"strings"
)

// Favorite handles execution of the "favorite" CLI command.
func Favorite(args []string) {
	if len(args) == 0 {
		fmt.Println("Please specify 'track' or 'artist' as argument.")
		return
	}

	switch strings.ToLower(args[0]) {
	case "track":
		internal.FavoriteTrack()
	case "artist":
		internal.FavoriteArtist()
	default:
		fmt.Println("Please specify 'track' or 'artist' as argument.")
	}
}

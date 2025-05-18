package internal

import (
	"fmt"
)

// CountArtists counts and prints the number of different artists you have listened to.
func CountArtists() {
	aggregatedArtists := LoadAggregatedArtistData()

	fmt.Println("You have listened to", len(aggregatedArtists), "different artists.")
}

// CountTracks counts and prints the number of different tracks you have listened to.
func CountTracks() {
	aggregatedTracks := LoadAggregatedTrackData()

	fmt.Println("You have listened to", len(aggregatedTracks), "different tracks.")
}

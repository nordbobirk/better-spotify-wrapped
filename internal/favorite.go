package internal

import (
	"fmt"
	"log"
)

// Favorite analyzes the aggregated listening data to find your favorite track by listening time.
func Favorite() {
	aggregatedEvents := LoadAggregatedTrackData()

	var favorite *AggregatedPlaybackEvent = nil // initialize null pointer

	for _, event := range aggregatedEvents {
		if favorite == nil {
			favorite = &event
		}

		if event.MsPlayed > favorite.MsPlayed {
			favorite = &event
		}
	}

	if favorite == nil {
		log.Fatal("No favorite track found")
	}

	fmt.Println("Your favorite track is:")
	fmt.Println("  Track name: ", *favorite.TrackName)
	fmt.Println("  Album name: ", *favorite.AlbumName)
	fmt.Println("  Album artist name: ", *favorite.AlbumArtistName)
	fmt.Println("  Minutes listened: ", favorite.MsPlayed/1000/60)
}

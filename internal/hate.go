package internal

import (
	"fmt"
	"log"
)

// HateArtist analyzed the aggregated listening data to find and print your least favorite artist by listening time.
func HateArtist() {
	aggregatedEvents := LoadAggregatedArtistData()

	var hate *ArtistAggregation = nil // initialize nil pointer

	for _, event := range aggregatedEvents {
		if hate == nil {
			hate = &event
		}

		if event.MsPlayed < hate.MsPlayed && event.MsPlayed > 0 {
			hate = &event
		}
	}

	if hate == nil {
		log.Fatal("No least favorite artist found")
	}

	fmt.Println("Your least favorite artist is:")
	fmt.Println("  Artist name: ", *hate.ArtistName)
	fmt.Println("  Milliseconds listened: ", hate.MsPlayed)
}

// HateTrack analyzed the aggregated listening data to find and print your least favorite track by listening time.
func HateTrack() {
	aggregatedEvents := LoadAggregatedTrackData()

	var hate *TrackAggregation = nil // initialize nil pointer

	for _, event := range aggregatedEvents {
		if hate == nil {
			hate = &event
		}

		if event.MsPlayed < hate.MsPlayed && event.MsPlayed > 0 {
			hate = &event
		}
	}

	if hate == nil {
		log.Fatal("No least favorite track found")
	}

	fmt.Println("Your least favorite track is:")
	fmt.Println("  Track name: ", *hate.TrackName)
	fmt.Println("  Album name: ", *hate.AlbumName)
	fmt.Println("  Album artist name: ", *hate.AlbumArtistName)
	fmt.Println("  Milliseconds listened: ", hate.MsPlayed)
}

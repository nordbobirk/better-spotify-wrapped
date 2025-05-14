package internal

import (
	"fmt"
	"sort"
)

func TopTracks(amount int) {
	aggregatedTracks := LoadAggregatedTrackData()

	var topTracks []*TrackAggregation

	for _, track := range aggregatedTracks {
		topTracks = append(topTracks, &track)
	}

	// sort by listening time
	sort.Slice(topTracks, func(i, j int) bool {
		return topTracks[i].MsPlayed > topTracks[j].MsPlayed
	})

	fmt.Println("Top tracks:")

	// print top tracks
	for i := range amount {
		fmt.Printf("%d. %s - %s - %s (%d min)\n", i+1, *(topTracks[i].TrackName), *(topTracks[i].AlbumName), *(topTracks[i].AlbumArtistName), topTracks[i].MsPlayed/1000/60)
	}
}

func TopArtists(amount int) {
	aggregatedArtists := LoadAggregatedArtistData()

	var topArtists []*ArtistAggregation

	for _, artist := range aggregatedArtists {
		topArtists = append(topArtists, &artist)
	}

	// sort by listening time
	sort.Slice(topArtists, func(i, j int) bool {
		return topArtists[i].MsPlayed > topArtists[j].MsPlayed
	})

	fmt.Println("Top artists:")

	// print top artists
	for i := range amount {
		fmt.Printf("%d. %s (%d min)\n", i+1, *(topArtists[i].ArtistName), topArtists[i].MsPlayed/1000/60)
	}
}

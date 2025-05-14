package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// ParseData parses raw Spotify listening data to an aggregated JSON file.
func ParseData() {
	entries, err := os.ReadDir(DataDir)

	if err != nil {
		log.Fatal(err)
	}

	if len(entries) == 0 {
		log.Fatal("No data found in data directory")
	}

	aggregatedTracks := make(map[string]TrackAggregation)
	aggregatedArtists := make(map[string]ArtistAggregation)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if entry.Name() == AggregatedTrackDataFile {
			continue
		}

		if entry.Name() == AggregatedArtistDataFile {
			continue
		}

		if !strings.Contains(entry.Name(), ".json") {
			continue
		}

		if !strings.Contains(entry.Name(), "Streaming_History_Audio") {
			continue
		}

		parseFile(entry, aggregatedTracks, aggregatedArtists)
	}

	// write results to files
	trackFile, trackFileErr := os.Create(DataDir + "/" + AggregatedTrackDataFile)
	artistFile, artistFileErr := os.Create(DataDir + "/" + AggregatedArtistDataFile)

	if trackFileErr != nil {
		log.Fatal(trackFileErr)
	}

	if artistFileErr != nil {
		log.Fatal(artistFileErr)
	}

	defer trackFile.Close()
	defer artistFile.Close()

	json.NewEncoder(artistFile).Encode(aggregatedArtists)
	json.NewEncoder(trackFile).Encode(aggregatedTracks)
}

func parseFile(entry os.DirEntry, aggregatedTracks map[string]TrackAggregation, aggregatedArtists map[string]ArtistAggregation) {
	fmt.Println("Parsing file: ", entry.Name())
	fileContent, err := os.ReadFile(DataDir + "/" + entry.Name())

	if err != nil {
		log.Fatal(err)
	}

	var events []PlaybackEvent
	err = json.Unmarshal(fileContent, &events)

	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		parseEvent(event, aggregatedTracks, aggregatedArtists)
	}
}

func parseEvent(event PlaybackEvent, aggregatedTracks map[string]TrackAggregation, aggregatedArtists map[string]ArtistAggregation) {
	if event.SpotifyTrackURI == nil {
		// if the event does not correspond to a track we ignore it
		return
	}

	if track, ok := aggregatedTracks[*event.SpotifyTrackURI]; ok {
		// if the event already exists in the map we update the aggregated event
		track.MsPlayed += event.MsPlayed
		aggregatedTracks[*event.SpotifyTrackURI] = track
	} else {
		// if the event does not exist in the map we create a new aggregated event
		aggregatedTracks[*event.SpotifyTrackURI] = TrackAggregation{
			MsPlayed:        event.MsPlayed,
			TrackName:       event.TrackName,
			AlbumArtistName: event.AlbumArtistName,
			AlbumName:       event.AlbumName,
		}
	}

	if artist, ok := aggregatedArtists[*event.AlbumArtistName]; ok {
		// if the event already exists in the map we update the aggregated event
		artist.MsPlayed += event.MsPlayed
		aggregatedArtists[*event.AlbumArtistName] = artist
	} else {
		// if the event does not exist in the map we create a new aggregated event
		aggregatedArtists[*event.AlbumArtistName] = ArtistAggregation{
			MsPlayed:   event.MsPlayed,
			ArtistName: event.AlbumArtistName,
		}
	}

}

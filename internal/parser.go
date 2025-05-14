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

	aggregatedEvents := make(map[string]AggregatedPlaybackEvent)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if entry.Name() == AggregatedTrackDataFile {
			continue
		}

		if !strings.Contains(entry.Name(), ".json") {
			continue
		}

		if !strings.Contains(entry.Name(), "Streaming_History_Audio") {
			continue
		}

		parseFile(entry, aggregatedEvents)
	}

	// write result to file
	file, err := os.Create(DataDir + "/" + AggregatedTrackDataFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	json.NewEncoder(file).Encode(aggregatedEvents)
}

func parseFile(entry os.DirEntry, aggregatedEvents map[string]AggregatedPlaybackEvent) {
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
		parseEvent(event, aggregatedEvents)
	}
}

func parseEvent(event PlaybackEvent, aggregatedEvents map[string]AggregatedPlaybackEvent) {
	if event.SpotifyTrackURI == nil {
		// if the event does not correspond to a track we ignore it
		return
	}

	if mapEvent, ok := aggregatedEvents[*event.SpotifyTrackURI]; ok {
		// if the event already exists in the map we update the aggregated event
		mapEvent.MsPlayed += event.MsPlayed
		aggregatedEvents[*event.SpotifyTrackURI] = mapEvent
	} else {
		// if the event does not exist in the map we create a new aggregated event
		aggregatedEvents[*event.SpotifyTrackURI] = AggregatedPlaybackEvent{
			MsPlayed:        event.MsPlayed,
			TrackName:       event.TrackName,
			AlbumArtistName: event.AlbumArtistName,
			AlbumName:       event.AlbumName,
		}
	}

}

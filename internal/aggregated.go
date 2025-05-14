package internal

import (
	"encoding/json"
	"log"
	"os"
)

// LoadAggregatedTrackData loads the aggregated listening data from the JSON file.
func LoadAggregatedTrackData() map[string]TrackAggregation {
	file, err := os.Open(DataDir + "/" + AggregatedTrackDataFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var aggregatedEvents map[string]TrackAggregation
	json.NewDecoder(file).Decode(&aggregatedEvents)

	return aggregatedEvents
}

func LoadAggregatedArtistData() map[string]ArtistAggregation {
	file, err := os.Open(DataDir + "/" + AggregatedArtistDataFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var aggregatedEvents map[string]ArtistAggregation
	json.NewDecoder(file).Decode(&aggregatedEvents)

	return aggregatedEvents
}

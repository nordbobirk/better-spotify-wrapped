package internal

import (
	"encoding/json"
	"log"
	"os"
)

// LoadAggregatedTrackData loads the aggregated listening data from the JSON file.
func LoadAggregatedTrackData() map[string]AggregatedPlaybackEvent {
	file, err := os.Open(DataDir + "/" + AggregatedTrackDataFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var aggregatedEvents map[string]AggregatedPlaybackEvent
	json.NewDecoder(file).Decode(&aggregatedEvents)

	return aggregatedEvents
}

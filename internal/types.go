package internal

import (
	"time"
)

type PlaybackEvent struct {
	Timestamp             time.Time `json:"ts"`
	Platform              string    `json:"platform"`
	MsPlayed              int       `json:"ms_played"`
	ConnectionCountry     string    `json:"conn_country"`
	IPAddress             string    `json:"ip_addr"`
	TrackName             *string   `json:"master_metadata_track_name"`
	AlbumArtistName       *string   `json:"master_metadata_album_artist_name"`
	AlbumName             *string   `json:"master_metadata_album_album_name"`
	SpotifyTrackURI       *string   `json:"spotify_track_uri"`
	EpisodeName           *string   `json:"episode_name"`
	EpisodeShowName       *string   `json:"episode_show_name"`
	SpotifyEpisodeURI     *string   `json:"spotify_episode_uri"`
	AudiobookTitle        *string   `json:"audiobook_title"`
	AudiobookURI          *string   `json:"audiobook_uri"`
	AudiobookChapterURI   *string   `json:"audiobook_chapter_uri"`
	AudiobookChapterTitle *string   `json:"audiobook_chapter_title"`
	ReasonStart           string    `json:"reason_start"`
	ReasonEnd             string    `json:"reason_end"`
	Shuffle               bool      `json:"shuffle"`
	Skipped               bool      `json:"skipped"`
	Offline               *bool     `json:"offline"`
	OfflineTimestamp      *int64    `json:"offline_timestamp"`
	IncognitoMode         bool      `json:"incognito_mode"`
}

type AggregatedPlaybackEvent struct {
	// SpotifyTrackURI       *string `json:"spotify_track_uri"` this is the key, so it's not needed here
	MsPlayed        int     `json:"ms_played"`
	TrackName       *string `json:"track_name"`
	AlbumArtistName *string `json:"album_artist_name"`
	AlbumName       *string `json:"album_name"`
}

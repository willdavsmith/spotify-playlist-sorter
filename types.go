package main

type Tracks struct {
	Items []Item `json:"items"`
}

type Item struct {
	Track ItemTrack `json:"track"`
}

type ItemTrack struct {
	ItemArtists []ItemArtist `json:"artists"`
}

type ItemArtist struct {
	Href string `json:"href"`
}

type Artist struct {
	Genres []string `json:"genres"`
}

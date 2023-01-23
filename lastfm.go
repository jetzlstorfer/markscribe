package main

import (
	"github.com/shkh/lastfm-go/lastfm"
)

func getFavouriteAlbums(year string) []lastfm.Album {

	api := lastfm.New(lastFmKey, lastFmSecret)

	params := lastfm.P{
		"user":  "jetinski1985",
		"limit": 5,
		"from":  year,
		"to":    year,
	}
	albums, err := api.User.GetTopAlbums(params)
	if err != nil {
		panic(err)
	}
	return albums
}

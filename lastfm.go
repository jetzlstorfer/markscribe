package main

import (
	"github.com/shkh/lastfm-go/lastfm"
)

func getFavouriteAlbums(count int) interface{} {
	api := lastfm.New(lastFMApiKey, lastFMSecret)

	params := lastfm.P{
		"user":  lastFMUser,
		"limit": count,
		"from":  2020,
		"to":    2022,
	}
	albums, err := api.User.GetTopAlbums(params)
	if err != nil {
		panic(err)
	}
	return albums.Albums
}

func getFavouriteTracks(count int) interface{} {
	api := lastfm.New(lastFMApiKey, lastFMSecret)

	params := lastfm.P{
		"user":  lastFMUser,
		"limit": count,
		"from":  2020,
		"to":    2022,
	}
	tracks, err := api.User.GetTopTracks(params)
	if err != nil {
		panic(err)
	}
	return tracks.Tracks
}

func getFavouriteArtists(count int) interface{} {
	api := lastfm.New(lastFMApiKey, lastFMSecret)

	params := lastfm.P{
		"user":  lastFMUser,
		"limit": count,
		"from":  2020,
		"to":    2022,
	}
	artists, err := api.User.GetTopArtists(params)
	if err != nil {
		panic(err)
	}
	return artists.Artists
}

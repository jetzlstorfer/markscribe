package main

import (
	"github.com/shkh/lastfm-go/lastfm"
)

func getFavouriteAlbums(count int) interface{} {

	params := lastfm.P{
		"user":  lastFMUser,
		"limit": count,
	}
	albums, err := lastfmapi.User.GetTopAlbums(params)
	if err != nil {
		panic(err)
	}
	return albums.Albums
}

func getFavouriteTracks(count int) interface{} {
	params := lastfm.P{
		"user":  lastFMUser,
		"limit": count,
	}
	tracks, err := lastfmapi.User.GetTopTracks(params)
	if err != nil {
		panic(err)
	}
	return tracks.Tracks
}

func getRecentTracks(count int) interface{} {
	params := lastfm.P{
		"user":  lastFMUser,
		"limit": count,
	}
	tracks, err := lastfmapi.User.GetRecentTracks(params)
	if err != nil {
		panic(err)
	}
	return tracks.Tracks
}

func getFavouriteArtists(count int) interface{} {
	params := lastfm.P{
		"user":  lastFMUser,
		"limit": count,
	}
	artists, err := lastfmapi.User.GetTopArtists(params)
	if err != nil {
		panic(err)
	}
	return artists.Artists
}

package main

import (
	"strings"

	graphql "github.com/graph-gophers/graphql-go"
)

func (_ *Resolver) Songs(album string) []Song {
	var filtered []Song
	filtered = Filter(songs, func(v *Song) bool {
		return strings.Contains(v.Album, album)
	})
	return filtered
}

func (_ *Resolver) Album(id graphql.ID) interface{} {
	for _, album := range albums {
		if album.ID == id {
			return album
		}
	}
	return nil
}

func (_ *Resolver) Artist(name string) interface{} {
	for _, artist := range artists {
		if artist.Name == name {
			return artist
		}
	}
	return nil
}

func (_ *Resolver) Genre(kind string) interface{} {
	for _, album := range albums {
		if album.Genre == kind {
			return album
		}
	}
	return nil
}

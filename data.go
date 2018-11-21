package main

import graphql "github.com/graph-gophers/graphql-go"

var albums = []*Album{
	{
		ID:     "ts-dark-as-night",
		Artist: "1",
		Title:  "Dark As Night",
		Year:   "2008",
		Genre:  "country",
		Type:   "album",
	},
}

var albumData = make(map[graphql.ID]*Album)

func init() {
	for _, album := range albums {
		albumData[album.ID] = album
	}
}

var artists = []*Artist{
	{
		ID:   "1",
		Name: "Nahko and Medicine for the People",
		Type: "artist",
	},
}

var artistData = make(map[graphql.ID]*Artist)

func init() {
	for _, artist := range artists {
		artistData[artist.ID] = artist
	}
}

var songs = []*Song{
	{
		ID:       "1",
		Album:    "Dark As Night",
		Title:    "Buddying Trees",
		Duration: "06:24",
		Type:     "song",
	},
	{
		ID:       "2",
		Album:    "Dark As Night",
		Title:    "Aloha Ke Akua",
		Duration: "05:56",
		Type:     "song",
	},
}

var songData = make(map[graphql.ID]*Song)

func init() {
	for _, song := range songs {
		songData[song.ID] = song
	}
}

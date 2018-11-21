package main

import graphql "github.com/graph-gophers/graphql-go"

// "Album"
type Album struct {
	ID     graphql.ID
	Artist string
	Title  string
	Year   string
	Genre  string
	Type   string
}

type albumResolver struct {
	album *Album
}

func (resolve *albumResolver) ID() graphql.ID {
	return resolve.album.ID
}

func (resolve *albumResolver) Artist() string {
	return resolve.album.Artist
}

func (resolve *albumResolver) Title() string {
	return resolve.album.Title
}

func (resolve *albumResolver) Year() string {
	return resolve.album.Year
}

func (resolve *albumResolver) Genre() *string {
	genre := &resolve.album.Genre
	if len(*genre) == 0 {
		return nil
	}
	return genre
}

func (resolve *albumResolver) Type() string {
	return resolve.album.Type
}

// "Artist"
type Artist struct {
	ID   graphql.ID
	Name string
	Type string
}

type artistResolver struct {
	artist *Artist
}

func (resolve *artistResolver) ID() graphql.ID {
	return resolve.artist.ID
}

func (resolve *artistResolver) Name() string {
	return resolve.artist.Name
}

func (resolve *artistResolver) Type() string {
	return resolve.artist.Type
}

type Song struct {
	ID       graphql.ID
	Album    string
	Title    string
	Duration string
	Type     string
}

type songResolver struct {
	song *Song
}

func (resolve *songResolver) ID() graphql.ID {
	return resolve.song.ID
}

func (resolve *songResolver) Album() string {
	return resolve.song.Album
}

func (resolve *songResolver) Title() string {
	return resolve.song.Title
}

func (resolve *songResolver) Duration() string {
	return resolve.song.Duration
}

func (resolve *songResolver) Type() string {
	return resolve.song.Type
}

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

type AlbumResolver struct {
	album *Album
}

func (resolve *AlbumResolver) ID() graphql.ID {
	return resolve.album.ID
}

func (resolve *AlbumResolver) Artist() string {
	return resolve.album.Artist
}

func (resolve *AlbumResolver) Title() string {
	return resolve.album.Title
}

func (resolve *AlbumResolver) Year() string {
	return resolve.album.Year
}

func (resolve *AlbumResolver) Genre() string {
	return resolve.album.Genre
}

func (resolve *AlbumResolver) Type() string {
	return resolve.album.Type
}

// "Artist"
type Artist struct {
	ID   graphql.ID
	Name string
	Type string
}

type ArtistResolver struct {
	artist *Artist
}

func (resolve *ArtistResolver) ID() graphql.ID {
	return resolve.artist.ID
}

func (resolve *ArtistResolver) Name() string {
	return resolve.artist.Name
}

func (resolve *ArtistResolver) Type() string {
	return resolve.artist.Type
}

type Song struct {
	ID       graphql.ID
	Album    string
	Title    string
	Duration string
	Type     string
}

type SongResolver struct {
	song *Song
}

func (resolve *SongResolver) ID() graphql.ID {
	return resolve.song.ID
}

func (resolve *SongResolver) Album() string {
	return resolve.song.Album
}

func (resolve *SongResolver) Title() string {
	return resolve.song.Title
}

func (resolve *SongResolver) Duration() string {
	return resolve.song.Duration
}

func (resolve *SongResolver) Type() string {
	return resolve.song.Type
}

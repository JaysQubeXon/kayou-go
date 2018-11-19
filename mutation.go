package main

import "github.com/graph-gophers/graphql-go"

func (_ *Resolver) CreateSong(args struct {
	id       graphql.ID
	album    string
	title    string
	duration string
}) Song {
	var song *Song
	song.ID = args.id
	song.Album = args.album
	song.Title = args.title
	song.Duration = args.duration
	song.Type = "song"
	songs = append(songs, song)
	return *song
}

func (_ *Resolver) UpdateSong(args struct {
	id       graphql.ID
	album    string
	title    string
	duration string
}) Song {
	song := &Song{}
	for _, selectedSong := range songs {
		if selectedSong.ID == args.id {
			if len(args.album) > 0 {
				selectedSong.Album = args.album
			}
			if len(args.title) > 0 {
				selectedSong.Title = args.title
			}
			if len(args.duration) > 0 {
				selectedSong.Duration = args.duration
			}
			song = selectedSong
			break
		}
	}
	return *song
}

func (_ *Resolver) DeleteSong(args struct{ id graphql.ID }) Song {
	song := &Song{}
	for i, selectedSong := range songs {
		if selectedSong.ID == args.id {
			song = songs[i]
			songs = append(songs[:i], songs[i+1:]...)
			break
		}
	}
	return *song
}

func (_ *Resolver) CreateArtist(args struct {
	id   graphql.ID
	name string
}) Artist {
	var artist *Artist
	artist.ID = args.id
	artist.Name = args.name
	artist.Type = "artist"
	artists = append(artists, artist)
	return *artist
}

func (_ *Resolver) UpdateArtist(args struct {
	id   graphql.ID
	name string
}) Artist {
	artist := &Artist{}
	for _, selectedArtist := range artists {
		if selectedArtist.ID == args.id {
			if len(args.name) > 0 {
				selectedArtist.Name = args.name
			}
			artist = selectedArtist
			break
		}
	}
	return *artist
}

func (_ *Resolver) DeleteArtist(args struct{ id graphql.ID }) Artist {
	artist := &Artist{}
	for i, selectedArtist := range artists {
		if selectedArtist.ID == args.id {
			artist = artists[i]
			artists = append(artists[:i], artists[i+1:]...)
			break
		}
	}
	return *artist
}

func (_ *Resolver) CreateAlbum(args struct {
	id     graphql.ID
	artist string
	title  string
	year   string
	genre  string
}) Album {
	var album *Album
	album.ID = args.id
	album.Artist = args.artist
	album.Title = args.title
	album.Year = args.year
	album.Genre = args.genre
	album.Type = "album"
	albums = append(albums, album)
	return *album
}

func (_ *Resolver) UpdateAlbum(args struct {
	id     graphql.ID
	artist string
	title  string
	year   string
	genre  string
}) Album {
	album := &Album{}
	for _, selectedAlbum := range albums {
		if selectedAlbum.ID == args.id {
			if len(args.artist) > 0 {
				selectedAlbum.Artist = args.artist
			}
			if len(args.title) > 0 {
				selectedAlbum.Title = args.title
			}
			if len(args.year) > 0 {
				selectedAlbum.Year = args.year
			}
			if len(args.genre) > 0 {
				selectedAlbum.Genre = args.genre
			}
			album = selectedAlbum
			break
		}
	}
	return *album
}

func (_ *Resolver) DeleteAlbum(args struct{ id graphql.ID }) Album {
	album := &Album{}
	for i, selectedAlbum := range albums {
		if selectedAlbum.ID == args.id {
			album = albums[i]
			albums = append(albums[:i], albums[i+1:]...)
			break
		}
	}
	return *album
}

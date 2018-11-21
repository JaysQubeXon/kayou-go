package main

import (
	"github.com/graph-gophers/graphql-go"
)

func (_ *Resolver) CreateSong(args *struct {
	ID       graphql.ID
	Album    string
	Title    string
	Duration string
}) *songResolver {
	song := &Song{
		ID:       args.ID,
		Album:    args.Album,
		Title:    args.Title,
		Duration: args.Duration,
		Type:     "song",
	}
	songData[song.ID] = song
	return &songResolver{song}
}

func (_ *Resolver) UpdateSong(args *struct {
	ID       graphql.ID
	Album    string
	Title    string
	Duration string
}) *songResolver {
	for _, selectedSong := range songData {
		if selectedSong.ID == args.ID {
			if len(args.Album) > 0 {
				selectedSong.Album = args.Album
			}
			if len(args.Title) > 0 {
				selectedSong.Title = args.Title
			}
			if len(args.Duration) > 0 {
				selectedSong.Duration = args.Duration
			}
			return &songResolver{selectedSong}

		}
	}
	return nil
}

func (_ *Resolver) DeleteSong(args struct{ ID graphql.ID }) *songResolver {

	for i, selectedSong := range songData {
		if selectedSong.ID == args.ID {
			songData[args.ID] = nil
			delete(songData, i)
			return &songResolver{selectedSong}
		}
	}
	return nil
}

func (_ *Resolver) CreateArtist(args *struct {
	ID   graphql.ID
	Name string
}) *artistResolver {
	artist := &Artist{
		ID:   args.ID,
		Name: args.Name,
		Type: "artist",
	}
	artistData[args.ID] = artist
	return &artistResolver{artist}
}

func (_ *Resolver) UpdateArtist(args *struct {
	ID   graphql.ID
	Name string
}) *artistResolver {
	for _, selectedArtist := range artists {
		if selectedArtist.ID == args.ID {
			if len(args.Name) > 0 {
				selectedArtist.Name = args.Name
			}
			return &artistResolver{selectedArtist}
		}
	}
	return nil
}

func (_ *Resolver) DeleteArtist(args *struct{ ID graphql.ID }) *artistResolver {
	for i, selectedArtist := range artistData {
		if selectedArtist.ID == args.ID {
			artistData[args.ID] = nil
			delete(artistData, i)
			return &artistResolver{selectedArtist}
		}
	}
	return nil
}

func (_ *Resolver) CreateAlbum(args *struct {
	ID     graphql.ID
	Artist string
	Title  string
	Year   string
	Genre  string
}) *albumResolver {
	album := &Album{
		ID:     args.ID,
		Artist: args.Artist,
		Title:  args.Title,
		Year:   args.Year,
		Genre:  args.Genre,
		Type:   "album",
	}
	albumData[args.ID] = album
	return &albumResolver{album}
}

func (_ *Resolver) UpdateAlbum(args *struct {
	ID     graphql.ID
	Artist string
	Title  string
	Year   string
	Genre  string
}) *albumResolver {
	for _, selectedAlbum := range albumData {
		if selectedAlbum.ID == args.ID {
			if len(args.Artist) > 0 {
				selectedAlbum.Artist = args.Artist
			}
			if len(args.Title) > 0 {
				selectedAlbum.Title = args.Title
			}
			if len(args.Year) > 0 {
				selectedAlbum.Year = args.Year
			}
			if len(args.Genre) > 0 {
				selectedAlbum.Genre = args.Genre
			}
			return &albumResolver{selectedAlbum}
		}
	}
	return nil
}

func (_ *Resolver) DeleteAlbum(args *struct{ ID graphql.ID }) *albumResolver {
	for i, selectedAlbum := range albumData {
		if selectedAlbum.ID == args.ID {
			albumData[args.ID] = nil
			delete(albumData, i)
			return &albumResolver{selectedAlbum}

		}
	}
	return nil
}

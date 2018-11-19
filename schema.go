package main

import (
	"errors"
)

func KayouSchema() (string, error) {
	s := `
		schema {
			query: Query
			mutation: Mutation 
		}
		type Query {
			songs(): [Song]
			album(id: String!): Album
			artist(name: String!): Artist
			genre(kind: String!): Album
		}
		type Mutation {
			# Create a new Song
			createSong(id: ID!, album: String!, title: String!, duration: String!): Song
			# Update a song with specific ID
			updateSong(id: ID!, album: String!, title: String!, duration: String!): Song
			# Delete a song with specified ID
			deleteSong(id: ID!): Boolean

			# Create a new Artist
			createArtist(id: ID!, name: String!): Artist
			# Update an Artist with specific ID
			updateArtist(id: ID!, name: String!): Artist
			# Delete an Artist with specified ID
			deleteArtist(id: ID!): Boolean

			# Create a new Album
			createAlbum(id: ID, artist: String!, title: String!, year: String!, genre: String): Album
			# Update an Album with specific ID
			updateAlbum(id: ID, artist: String!, title: String!, year: String!, genre: String): Album
			# Delete an Album with specidied ID
			deleteAlbum(id: ID!): Boolean
		}
		# Song in an album
		type Song {
			# ID of the song
			id: ID!
			# Album the song belongs to
			album: String!
			# Title of the song
			title: String!
			# Duration of the song
			duration: String!
		}
		# Album 
		type Album {
			id: ID!
			artist: String!
			title: String!
			year: String!
			genre: String
		}
		type Artist {
			id: ID!
			name: String!
		}
	`
	if len(s) > 0 {
		return s, nil
	}
	return s, errors.New("Schema is empty, please provide one")
}

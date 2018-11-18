package main

import "github.com/graphql-go/graphql"

// "Root Mutation"
func RootMutation(albumType graphql.Type, artistType graphql.Type, songType graphql.Type) (*graphql.Object, error) {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createSong": &graphql.Field{
				Type:        songType,
				Description: "Create a new song and add to a specified album",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"album": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var song Song
					song.ID = params.Args["id"].(string)
					song.Album = params.Args["album"].(string)
					song.Title = params.Args["title"].(string)
					song.Duration = params.Args["duration"].(string)
					songs = append(songs, song)
					return song, nil
				},
			},
			"updateSong": &graphql.Field{
				Type:        songType,
				Description: "Update a songs properties",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"album": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)
					album, albumOk := params.Args["album"].(string)
					title, titleOk := params.Args["title"].(string)
					duration, durationOk := params.Args["duration"].(string)
					song := Song{}
					for _, selectedSong := range songs {
						if selectedSong.ID == id {
							if albumOk {
								selectedSong.Album = album
							}
							if titleOk {
								selectedSong.Title = title
							}
							if durationOk {
								selectedSong.Duration = duration
							}
							song = selectedSong
							break
						}
					}
					return song, nil
				},
			},
			"deleteSong": &graphql.Field{
				Type:        songType,
				Description: "Delete a song from the songs list",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"].(string)
					var song = Song{}
					for i, selectedSong := range songs {
						if selectedSong.ID == id {
							song = songs[i]
							songs = append(songs[:i], songs[i+1:]...)
							break
						}
					}
					return song, nil
				},
			},
			"createArtist": &graphql.Field{
				Type:        artistType,
				Description: "Create a new artist",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var artist Artist
					artist.ID = params.Args["id"].(string)
					artist.Name = params.Args["name"].(string)
					artists = append(artists, artist)
					return artist, nil
				},
			},
		},
	}), nil
}

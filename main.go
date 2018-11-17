package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
)

func main() {
	albumType, artistType, songType := dataTypes()
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"songs": &graphql.Field{
				Type: graphql.NewList(songType),
				Args: graphql.FieldConfigArgument{
					"album": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					album := params.Args["album"].(string)
					filtered := Filter(songs, func(v Song) bool {
						return strings.Contains(v.Album, album)
					})
					return filtered, nil
				},
			},
			"album": &graphql.Field{
				Type: albumType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"].(string)
					for _, album := range albums {
						if album.ID == id {
							return album, nil
						}
					}
					return nil, nil
				},
			},
			"artist": &graphql.Field{
				Type: artistType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name := params.Args["name"].(string)
					for _, artist := range artists {
						if artist.Name == name {
							return artist, nil
						}
					}
					return nil, nil
				},
			},
			"genre": &graphql.Field{
				Type: albumType,
				Args: graphql.FieldConfigArgument{
					"kind": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					kind := params.Args["kind"].(string)
					for _, album := range albums {
						if album.Genre == kind {
							return album, nil
						}
					}
					return nil, nil
				},
			},
		},
	})
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createSong": &graphql.Field{
				Type: songType,
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
				Type: songType,
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
				Type: songType,
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
				Type: artistType,
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
	})
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})
	fmt.Printf("server listening on port 8080\n")
	http.ListenAndServe(":8080", nil)
}

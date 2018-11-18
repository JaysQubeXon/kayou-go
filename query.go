package main

import (
	"strings"

	"github.com/graphql-go/graphql"
)

// "Root query"
func RootQuery(albumType graphql.Type, artistType graphql.Type, songType graphql.Type) (*graphql.Object, error) {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Query",
		Description: "Root query object",
		Fields: graphql.Fields{
			"songs": &graphql.Field{
				Type:        graphql.NewList(songType),
				Description: "Get all songs in an album",
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
				Type:        albumType,
				Description: "Get album by ID",
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
				Type:        artistType,
				Description: "Get artist by name",
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
				Type:        albumType,
				Description: "Get albums by genre kind",
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
	}), nil
}

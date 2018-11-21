package main

import (
	graphql "github.com/graph-gophers/graphql-go"
)

func (_ *Resolver) Album(args struct{ ID graphql.ID }) *albumResolver {
	if a := albumData[args.ID]; a != nil {
		return &albumResolver{a}
	}
	return nil
}

func (_ *Resolver) Artist(args struct{ Name string }) *artistResolver {
	for _, artist := range artistData {
		if artist.Name == args.Name {
			return &artistResolver{artist}
		}
	}
	return nil
}

func (_ *Resolver) Genre(args struct{ Kind string }) *albumResolver {
	for _, album := range albumData {
		if album.Genre == args.Kind {
			return &albumResolver{album}
		}
	}
	return nil
}

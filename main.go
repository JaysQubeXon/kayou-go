package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!\n" }

func main() {
	s := `
		schema {
			query: Query
		}
		type Query {
			hello: String!
		}
	`
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	fmt.Printf("server listening on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

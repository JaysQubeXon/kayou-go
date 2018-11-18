package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	rootQuery, _ := RootQuery(dataTypes())
	rootMutation, _ := RootMutation(dataTypes())
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

package main

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

type Album struct {
	ID		string `json:"id:omitempty"`
	Artist	string `json:"artist"`
	Title	string `json:"title"`
	Year	string `json:"year"`
	Gender	string `json:"gender"`
	Type	string `json:"type"`
}

type Artist struct {
	ID		string `json:"id:omitempty"`
	Name	string `json:"name"`
	Type	string `json:"type"`
}

type Song struct {
	ID			string `json:"id:omitempty"`
	Album		string `json:"album"`
	Title		string `json:"title"`
	Duration	string `json:"duration"`
	Type		string `json:"type"`
}

func main() {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{})
	http.HandleFunc("graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:			schema,
			RequestString:	r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8080", nil)
}

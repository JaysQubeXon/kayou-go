package main

type Album struct {
	ID     string `json:"id:omitempty"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Year   string `json:"year"`
	Genre  string `json:"genre"`
	// Description string `json:"description"`
	Type string `json:"type"`
}

type Artist struct {
	ID   string `json:"id:omitempty"`
	Name string `json:"name"`
	// Description string `json:"description"`
	Type string `json:"type"`
}

type Song struct {
	ID       string `json:"id:omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	// Description string `json:"description"`
	Type string `json:"type"`
}

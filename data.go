package main

var albums = []Album{
	Album{
		ID:     "ts-dark-as-night",
		Artist: "1",
		Title:  "Dark As Night",
		Year:   "2008",
		Genre:  "country",
		Type:   "album",
	},
}

var artists = []Artist{
	Artist{
		ID:   "1",
		Name: "Nahko and Medicine for the People",
		Type: "artist",
	},
}

var songs = []Song{
	Song{
		ID:       "1",
		Album:    "Dark As Night",
		Title:    "Buddying Trees",
		Duration: "06:24",
		Type:     "song",
	},
	Song{
		ID:       "2",
		Album:    "Dark As Night",
		Title:    "Aloha Ke Akua",
		Duration: "05:56",
		Type:     "song",
	},
}

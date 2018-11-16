package main

// Filter receives songs as first param
// and a func as second param, that checks if
// value is a song in the Song array. If it is,
// it appends it to a new array of type Song
// and returns it.
func Filter(songs []Song, f func(Song) bool) []Song {
	vsf := make([]Song, 0)
	for _, v := range songs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

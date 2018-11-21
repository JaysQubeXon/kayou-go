package main

func filter(songs *[]*Song, f func(*Song) bool) *[]*Song {
	vsf := make([]*Song, 0)
	for _, v := range *songs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return &vsf
}

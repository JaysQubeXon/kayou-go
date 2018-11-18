# kayou-go

The Go implementation of the kayou server

### instructions:

BUild with `go build` command.

open browser at:

[`http://localhost:8080/graphql`](http://localhost:8080/graphql)

example queries:

create a new song with the `createSong` mutation:

`http://localhost:8080/graphql?query=mutation+_{createSong(id:"3",album:"Dark As Night",title:"Warrior People",duration:"04:57"){title,duration}}`

update a song with `updateSong` mutation:

`http://localhost:8080/graphql?query=mutation+_{updateSong(id:"1",album: "Dark As Night",title:"Warrior People!",duration:"06:55"){album,title,duration}}`

delete a song with `deleteSong` mutation:

`http://localhost:8080/graphql?query=mutation+_{deleteSong(id:"1"){id,album,title,duration}}`


create a new artist with `createArtist`:

`http://localhost:8080/graphql?query=mutation+_{createArtist(id:"2", name: "boaz barlia"){id,name}}`

other options:

query for all songs:

`http://localhost:8080/graphql?query={songs(album:"Dark As Night"){id,album,title,duration}}`

query for an artist:

`http://localhost:8080/graphql?query={artist(name:"Nahko and Medicine for the People"){id,name}}`

query for an album:

`http://localhost:8080/graphql?query={album(id:"ts-dark-as-night"){id,title,year}}`

query for a specific genre:

`http://localhost:8080/graphql?query={genre(kind:%22country%22){id,title,year}}`


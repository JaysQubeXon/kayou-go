# kayou-go

The Go implementation of the kayou server using the [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go) package.

### instructions:

Build with `go build` command and run with `./kayou-go`.

open browser at:

[`http://localhost:8080`](http://localhost:8080)

example queries:

create a new song with the `CreateSong` mutation:

```javascript
mutation createNewSong($id: ID!, $album: String!, $title: String!, $duration: String!) {
    createSong(id: $id, album: $album, title: $title, duration: $duration) {
    id
    album
    title
    duration
    }
}
```
add the following into the `Query Variables` section:
```JSON
{
  "id": "7",
  "album": "Dark As Night",
  "title": "Warrior People",
  "duration": "04:57"
}
```

Query options:

query for an album:

```javascript
{
  album(id: "ts-dark-as-night") {
    id
    title
    artist
    year
    genre
  }   
}
```
returns:
```JSON
{
  "data": {
    "album": {
      "id": "ts-dark-as-night",
      "title": "Dark As Night",
      "artist": "1",
      "year": "2008",
      "genre": "country"
    }
  }
}
```

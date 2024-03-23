package album

import (
	"errors"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Add the new album to the slice
func add(newAlbum *album) {
	albums = append(albums, *newAlbum)
}

func find(id *string) (album, error) {
	for _, a := range albums {
		if a.ID == *id {
			return a, nil
		}
	}
	return album{}, errors.New("id not found")
}

func update(updated *album) error {
	for index, a := range albums {
		if a.ID == *&updated.ID {
			albums[index] = *updated
			return nil
		}
	}
	return errors.New("id not found")
}

func delete(id *string) error {
	for index, a := range albums {
		if a.ID == *id {
			removeNoOrder(albums, index)
			return nil
		}
	}
	return errors.New("id not found")
}

func removeNoOrder(s []album, i int) []album {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeaWithOrder(slice []album, s int) []album {
	return append(slice[:s], slice[s+1:]...)
}

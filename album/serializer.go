package album

type albumResponse struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type newAlbumRequest struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func Response(a Album) albumResponse {
	return albumResponse{
		ID:     a.ID,
		Title:  a.Title,
		Artist: a.Artist,
		Price:  a.Price,
	}
}

func validateUpdate(nr *newAlbumRequest, id string) *Album {
	// i should validate sice and price
	return &Album{
		ID:     id,
		Title:  nr.Title,
		Artist: nr.Artist,
		Price:  nr.Price,
	}
}

func validateCreate(nr *newAlbumRequest) *Album {
	return &Album{
		Title:  nr.Title,
		Artist: nr.Artist,
		Price:  nr.Price,
	}
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

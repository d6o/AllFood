package main

type Food interface {
	Search(int, string, string) ([]Place, error)
}

type Place struct {
	PlaceID string
	Name    string
	Lat     float64
	Lng     float64
}

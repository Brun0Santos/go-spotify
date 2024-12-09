package handler

import (
	"net/http"
)

type ArtistsGet struct {
}

func NewArtistsGet() *ArtistsGet {
	return &ArtistsGet{}
}

func (handler *ArtistsGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

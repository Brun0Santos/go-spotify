package repositories

import "github.com/Brun0Santos/api-spotify/internal/models"

type Artist interface {
	GetByUID(uid string) (models.Artist, error)
}

package models

type Image struct {
	ID     int64  `json:"-"`
	UID    string `json:"-"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

package models

type Artist struct {
	ID         int64    `json:"id"`
	UID        string   `json:"uid"`
	Name       string   `json:"name"`
	Popularity int64    `json:"popularity"`
	Genres     []string `json:"genres"`

	URI          string   `json:"uri,omitempty"`
	Href         string   `json:"href,omitempty"`
	ExternalURLs []string `json:"external_urls,omitempty"`
	Followers    int      `json:"followers,omitempty"`

	Images []Image `json:"images"`
	Users  []User  `json:"users"`
}

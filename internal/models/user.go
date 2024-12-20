package models

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Address  string `json:"address"`
}

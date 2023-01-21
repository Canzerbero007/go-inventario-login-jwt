package models

type User struct {
	Id    int64  `json:"Id"`
	Email string `json:"Email"`
	Name  string `json:"Name"`
}

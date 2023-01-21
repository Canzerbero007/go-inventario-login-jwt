package models

type Product struct {
	Id          int64   `json:"Id"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	Price       float32 `json:"Price"`
}

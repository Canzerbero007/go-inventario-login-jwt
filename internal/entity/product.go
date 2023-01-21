package entity

type Product struct {
	Id          int64   `db:"Id"`
	Name        string  `db:"Name"`
	Description string  `db:"Description"`
	Price       float32 `db:"Price"`
	CreatedBy   int64   `db:"Created_by"`
}

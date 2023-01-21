package entity

type User struct {
	Id       int64  `db:"Id"`
	Email    string `db:"Email"`
	Name     string `db:"Name"`
	Password string `db:"Password"`
}

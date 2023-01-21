package entity

type UserRole struct {
	Id     int64 `db:"Id"`
	UserId int64 `db:"UserId"`
	RoleId int64 `db:"RoleId"`
}

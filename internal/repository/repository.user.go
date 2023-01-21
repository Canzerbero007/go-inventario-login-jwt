package repository

import (
	"context"

	"yanita_inventario/internal/entity"
)

const (
	qryUserInsert = `INSERT INTO USERS (Email, Name, Password) values (?, ?, ?);`

	qryUserFindByEmail = `SELECT Id, Email, Name, Password FROM USERS WHERE Email = ?;`

	qryInsertUserRole = `insert into USER_ROLES (UserId, RoleId) values (:userid, :roleid);`

	qryRemoveUserRole = `delete from USER_ROLES where UserId = :userid and RoleId = :roleid;`
)

func (r *repo) UserRegister(ctx context.Context, email, name, password string) error {

	_, err := r.db.ExecContext(ctx, qryUserInsert, email, name, password)
	return err

}

func (r *repo) UserFindByEmail(ctx context.Context, Email string) (*entity.User, error) {

	u := &entity.User{}

	err := r.db.GetContext(ctx, u, qryUserFindByEmail, Email)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repo) SaveUserRole(ctx context.Context, userId, roleId int64) error {
	data := entity.UserRole{
		UserId: userId,
		RoleId: roleId,
	}

	_, err := r.db.NamedExecContext(ctx, qryInsertUserRole, data)
	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userId, roleId int64) error {
	data := entity.UserRole{
		UserId: userId,
		RoleId: roleId,
	}

	_, err := r.db.NamedExecContext(ctx, qryRemoveUserRole, data)

	return err
}

func (r *repo) GetUserRoles(ctx context.Context, userID int64) ([]entity.UserRole, error) {
	roles := []entity.UserRole{}

	err := r.db.SelectContext(ctx, &roles, "select UserId, RoleId from USER_ROLES where Id = ?", userID)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

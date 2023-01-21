package service

import (
	"context"
	"errors"

	"yanita_inventario/encryption"
	models "yanita_inventario/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrRoleAlreadyAdded   = errors.New("role was already added for this user")
	ErrRoleNotFound       = errors.New("role not found")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {

	u, _ := s.repo.UserFindByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)
	return s.repo.UserRegister(ctx, email, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.UserFindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		Id:    u.Id,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

func (s *serv) AddUserRole(ctx context.Context, userId, roleId int64) error {

	roles, err := s.repo.GetUserRoles(ctx, userId)
	if err != nil {
		return err
	}

	for _, r := range roles {
		if r.RoleId == roleId {
			return ErrRoleAlreadyAdded
		}
	}

	return s.repo.SaveUserRole(ctx, userId, roleId)
}

func (s *serv) RemoveUserRole(ctx context.Context, userId, roleId int64) error {
	roles, err := s.repo.GetUserRoles(ctx, userId)
	if err != nil {
		return err
	}

	roleFound := false
	for _, r := range roles {
		if r.RoleId == roleId {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}

	return s.repo.RemoveUserRole(ctx, userId, roleId)
}

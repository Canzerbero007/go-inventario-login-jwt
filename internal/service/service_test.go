package service

import (
	"os"
	"testing"

	mock "github.com/stretchr/testify/mock"

	"yanita_inventario/encryption"
	"yanita_inventario/internal/entity"
	"yanita_inventario/internal/repository"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)
	u := &entity.User{Email: "test@exists.com", Password: encryptedPassword}
	adminUser := &entity.User{
		Id:       1,
		Email:    "admin@email.com",
		Name:     "",
		Password: encryptedPassword,
	}
	customerUser := &entity.User{
		Id:       2,
		Email:    "customer@email.com",
		Name:     "",
		Password: encryptedPassword,
	}

	repo = &repository.MockRepository{}
	repo.On("UserFindByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("UserFindByEmail", mock.Anything, "test@exists.com").Return(u, nil)
	repo.On("UserFindByEmail", mock.Anything, "admin@email.com").Return(adminUser, nil)
	repo.On("UserFindByEmail", mock.Anything, "customer@email.com").Return(customerUser, nil)

	repo.On("RegisterUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo.On("GetUserRoles", mock.Anything, int64(1)).Return([]entity.UserRole{{UserId: 1, RoleId: 1}}, nil)
	repo.On("GetUserRoles", mock.Anything, int64(2)).Return([]entity.UserRole{{UserId: 2, RoleId: 3}}, nil)

	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("RemoveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo.On("ProductRegister", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}

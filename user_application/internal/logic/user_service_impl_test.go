package logic

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"hands_on_go/internal/uaerr"
	"testing"
)

func TestUserServiceImpl_GetByID_RepoErr_ReturnsErr(t *testing.T) {
	inputID := 123

	userRepo := &userRepositoryStub{}
	userRepo.GetUserFunc = func(id int) (*User, error) {
		assert.Equal(t, inputID, id)

		return nil, errors.New("oh no")
	}

	service := NewUserService(userRepo)

	_, err := service.GetByID(inputID)

	assert.NotNil(t, err)
}

func TestUserServiceImpl_GetByID_Success_ReturnsUser(t *testing.T) {
	inputID := 123
	expectedUser := &User{}

	userRepo := &userRepositoryStub{}
	userRepo.GetUserFunc = func(id int) (*User, error) {
		assert.Equal(t, inputID, id)

		return expectedUser, nil
	}

	service := NewUserService(userRepo)

	result, err := service.GetByID(inputID)

	if assert.Nil(t, err) {
		assert.True(t, result == expectedUser)
	}
}

func TestUserServiceImpl_DeleteByID_UserExists_UserIsDeleted(t *testing.T) {
	inputID := 123
	userIsDeleted := false

	userRepo := &userRepositoryStub{}
	userRepo.CheckUserExistsFunc = func(id int) (bool, error) {
		return true, nil
	}
	userRepo.DeleteUserFunc = func(id int) error {
		assert.Equal(t, inputID, id)
		userIsDeleted = true

		return nil
	}

	userService := NewUserService(userRepo)
	err := userService.DeleteByID(inputID)

	assert.Nil(t, err)
	assert.True(t, userIsDeleted)
}

// TODO When user not exists
// 		-> return TypeNotFound error with user message

func TestUserServiceImpl_DeleteByID_UserNotExists_NotFoundError(t *testing.T) {
	userRepo := &userRepositoryStub{}
	userRepo.CheckUserExistsFunc = func(id int) (bool, error) {
		return false, nil
	}
	userRepo.DeleteUserFunc = func(id int) error {
		assert.Fail(t, "delete should not be called")
		return nil
	}

	userService := NewUserService(userRepo)
	err := userService.DeleteByID(123)

	assert.NotNil(t, err)
	assert.Equal(t, uaerr.TypeNotFound, uaerr.GetType(err))
	msg, ok := uaerr.GetUserMsg(err)
	assert.True(t, ok)
	assert.Equal(t, "user not found", msg)
}

type userRepositoryStub struct {
	GetUserFunc         func(id int) (*User, error)
	CreateUserFunc      func(user *User) (int, error)
	CheckUserExistsFunc func(id int) (bool, error)
	DeleteUserFunc      func(id int) error
}

func (s *userRepositoryStub) CreateUser(user *User) (int, error) {
	return s.CreateUserFunc(user)
}

func (s *userRepositoryStub) GetUser(id int) (*User, error) {
	return s.GetUserFunc(id)
}

func (s *userRepositoryStub) CheckUserExists(id int) (bool, error) {
	return s.CheckUserExistsFunc(id)
}

func (s *userRepositoryStub) DeleteUser(id int) error {
	return s.DeleteUserFunc(id)
}

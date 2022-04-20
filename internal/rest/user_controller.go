package rest

import (
	"encoding/json"
	"errors"
	"github.com/phonaputer/hands_on_go/internal/blerr"
	"github.com/phonaputer/hands_on_go/internal/logic"
	"github.com/phonaputer/hands_on_go/internal/model"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UserController struct {
	validator   userValidator
	userService logic.UserService
}

func NewUserController(userValidator userValidator,
	userService logic.UserService) *UserController {

	return &UserController{
		validator:   userValidator,
		userService: userService,
	}
}

func (u *UserController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id, err := u.validator.ValidateDeleteUserByID(r)
	if errors.Is(err, blerr.ErrInvalidInput) {
		w.WriteHeader(400)
		w.Write([]byte("request not valid"))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("error validating delete user by ID")
		w.WriteHeader(500)
		w.Write([]byte("an unexpected error has occurred"))
		return
	}

	err = u.userService.DeleteByID(id)
	if errors.Is(err, blerr.ErrUserNotFound) {
		w.WriteHeader(404)
		w.Write([]byte("user not found"))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("error deleting user by ID")
		w.WriteHeader(500)
		w.Write([]byte("an unexpected error has occurred"))
		return
	}

	w.WriteHeader(204)
}

func (u *UserController) GetByID(w http.ResponseWriter, r *http.Request) {

	// 1.
	id, err := u.validator.ValidateGetUserByID(r)
	if errors.Is(err, blerr.ErrInvalidInput) {
		w.WriteHeader(400)
		w.Write([]byte("request not valid"))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("error validating get user by ID")
		w.WriteHeader(500)
		w.Write([]byte("an unexpected error has occurred"))
		return
	}

	// 2.
	userModel, err := u.userService.GetByID(id)
	if errors.Is(err, blerr.ErrUserNotFound) {
		w.WriteHeader(404)
		w.Write([]byte("user not found"))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("error getting user by ID")
		w.WriteHeader(500)
		w.Write([]byte("an unexpected error has occurred"))
		return
	}

	// 3.
	responseBody := u.toGetByIDResponse(userModel)

	// 4.
	bodyBytes, err := json.Marshal(responseBody)
	if err != nil {
		logrus.WithError(err).Error("error serializing get user by ID response JSON")
		w.WriteHeader(500)
		w.Write([]byte("an unexpected error has occurred"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bodyBytes)
}

func (u *UserController) toGetByIDResponse(user *model.User) *getUserByIDResponse {
	return &getUserByIDResponse{
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Age:             json.Number(strconv.Itoa(user.Age)),
		PhoneNumber:     user.PhoneNumber,
		IsPhoneVerified: user.IsPhoneVerified,
	}
}

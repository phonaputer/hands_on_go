package presentation

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"hands_on_go/internal/logic"
	"net/http"
)

type UserController struct {
	validator   userValidator
	userService logic.UserService
}

func NewUserController(
	validator userValidator,
	userService logic.UserService,
) *UserController {
	return &UserController{
		validator:   validator,
		userService: userService,
	}
}

func (u *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {

	// 1. Get ID from request (query string)
	id, err := u.validator.ValidateGetUserByID(r)

	// 1.1. if this is invalid -> return 400 response
	if errors.Is(err, errInvalidInput) {
		logrus.WithError(err).Error("invalid input to get user by ID")
		w.WriteHeader(400)
		return
	}

	// unexpected error occurred -> return 500
	if err != nil {
		logrus.WithError(err).Error("unexpected validation error in get user by ID")
		w.WriteHeader(500)
		return
	}

	// 2. Pass ID to business logic layer & get back a user
	user, err := u.userService.GetByID(id)

	// 2.1. If user is not found -> return 404 response
	if errors.Is(err, logic.ErrNotFound) {
		logrus.WithError(err).Error("user not found in get user by ID")
		w.WriteHeader(404)
		return
	}

	// unexpected error occurred -> return 500
	if err != nil {
		logrus.WithError(err).Error("unexpected service error in get user by ID")
		w.WriteHeader(500)
		return
	}

	// 3. Map user model to user JSON response model
	responseBody := u.toGetByIDResponse(user)

	bodyBytes, err := json.Marshal(responseBody)
	if err != nil {
		logrus.WithError(err).Error("error serializing JSON in get user by ID")
		w.WriteHeader(500)
		return
	}

	// 4. Serialize JSON response model and write 200 response with JSON
	//	body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bodyBytes)
}

func (u *UserController) toGetByIDResponse(user *logic.User) *getUserByIDResponse {
	return &getUserByIDResponse{
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Age:             user.Age,
		PhoneNumber:     user.PhoneNumber,
		IsPhoneVerified: user.PhoneVerified,
	}
}

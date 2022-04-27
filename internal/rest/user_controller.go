package rest

import (
	"encoding/json"
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

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) error {
	validatedRequest, err := u.validator.ValidateCreateUser(r)
	if err != nil {
		return err
	}

	user := u.createRequestToUserModel(validatedRequest)

	id, err := u.userService.Create(user)
	if err != nil {
		return err
	}

	bodyBytes, err := json.Marshal(&createUserResponse{ID: id})
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(bodyBytes)

	return nil
}

func (u *UserController) createRequestToUserModel(req *createUserRequest) *model.User {
	age, err := req.Age.Int64()
	if err != nil {
		logrus.WithError(err).Error("unexpected invalid age in user create request to model mapping!")
	}

	return &model.User{
		FirstName:       *req.FirstName,
		LastName:        *req.LastName,
		Age:             int(age),
		PhoneNumber:     *req.PhoneNumber,
		IsPhoneVerified: *req.IsPhoneNumberVerified,
	}
}

func (u *UserController) DeleteByID(w http.ResponseWriter, r *http.Request) error {
	id, err := u.validator.ValidateDeleteUserByID(r)
	if err != nil {
		return err
	}

	err = u.userService.DeleteByID(id)
	if err != nil {
		return err
	}

	w.WriteHeader(204)

	return nil
}

func (u *UserController) GetByID(w http.ResponseWriter, r *http.Request) error {

	// 1.
	id, err := u.validator.ValidateGetUserByID(r)
	if err != nil {
		return err
	}

	// 2.
	userModel, err := u.userService.GetByID(id)
	if err != nil {
		return err
	}

	// 3.
	responseBody := u.toGetByIDResponse(userModel)

	// 4.
	bodyBytes, err := json.Marshal(responseBody)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bodyBytes)

	return nil
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

package presentation

import (
	"encoding/json"
	"fmt"
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

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) error {

	// 1. parse & validate request data
	reqData, err := u.validator.ValidateCreateUser(r)
	if err != nil {
		return fmt.Errorf("validate request: %w", err)
	}

	// 2. Map create request data to User logic struct
	user := &logic.User{
		FirstName:     *reqData.FirstName,
		LastName:      *reqData.LastName,
		Age:           *reqData.Age,
		PhoneNumber:   *reqData.PhoneNumber,
		PhoneVerified: *reqData.IsPhoneVerified,
	}

	// 3. Create user in business logic layer
	id, err := u.userService.CreateUser(user)
	if err != nil {
		return fmt.Errorf("user service create user: %w", err)
	}

	// 4. Map user ID to create user response
	responseBody := createUserResponse{ID: id}

	bodyBytes, err := json.Marshal(responseBody)
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}

	// 5. Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(bodyBytes)

	return nil
}

func (u *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) error {

	// 1. Get ID from request (query string)
	id, err := u.validator.ValidateGetUserByID(r)
	if err != nil {
		return fmt.Errorf("validate request: %w", err)
	}

	// 2. Pass ID to business logic layer & get back a user
	user, err := u.userService.GetByID(id)
	if err != nil {
		return fmt.Errorf("user service get by ID: %w", err)
	}

	// 3. Map user model to user JSON response model
	responseBody := u.toGetByIDResponse(user)

	bodyBytes, err := json.Marshal(responseBody)
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}

	// 4. Serialize JSON response model and write 200 response with JSON
	//	body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bodyBytes)

	return nil
}

func (u *UserController) DeleteUserByID(w http.ResponseWriter, r *http.Request) error {

	// 1. Get ID from request (query string)
	id, err := u.validator.ValidateDeleteUserByID(r)
	if err != nil {
		return fmt.Errorf("validate request: %w", err)
	}

	// 2. Pass ID to business logic layer & get back a user
	err = u.userService.DeleteByID(id)
	if err != nil {
		return fmt.Errorf("user service delete by ID: %w", err)
	}

	// 3. write response
	w.WriteHeader(204)

	return nil
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

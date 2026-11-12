package restapi

import (
	"fmt"
	"net/http"
	"strconv"
	"userapp/internal/userapp/logic"
)

type UsersController struct {
	crud      logic.UserCRUD
	mapper    usersMapper
	validator usersValidator
}

func NewUsersController(
	crud logic.UserCRUD,
	mapper usersMapper,
	validator usersValidator,
) *UsersController {
	return &UsersController{
		crud:      crud,
		mapper:    mapper,
		validator: validator,
	}
}

func (c *UsersController) GetUser(w http.ResponseWriter, r *http.Request) error {
	// 1. Get & validate input
	userIDStr := r.PathValue("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		err = fmt.Errorf("%w: convert id to integer: %w", logic.ErrInvalidInput, err)
		return logic.SetEndUserMsg(err, "User ID must be an integer!")
	}

	// 2. Get actual user data from logic layer
	user, err := c.crud.GetByID(r.Context(), userID)
	if err != nil {
		return fmt.Errorf("get by id: %w", err)
	}

	// 3. Map result to HTTP response and send the response back
	response := c.mapper.User2UserResponse(user)

	return jsonResponse(w, http.StatusOK, response)
}

func (c *UsersController) Create(w http.ResponseWriter, r *http.Request) error {
	// Hexagonal architecture

	// 1. Receive request, parse it, verify it is valid
	user, err := c.validator.Create(r)
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	// 2. Call business logic to create the user
	userID, err := c.crud.Create(r.Context(), user)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	// 3. Serialize the resulting user ID to response
	return jsonResponse(w, http.StatusCreated, &userIDResponse{ID: userID})
}

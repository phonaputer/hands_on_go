package restapi

import (
	"userapp/internal/userapp/logic"
)

type usersMapper interface {
	User2UserResponse(user *logic.User) *userResponse
}

type jsonUsersMapper struct{}

func newJSONUsersMapper() jsonUsersMapper {
	return jsonUsersMapper{}
}

func (m *jsonUsersMapper) User2UserResponse(user *logic.User) *userResponse {
	result := userResponse(*user)

	return &result
}

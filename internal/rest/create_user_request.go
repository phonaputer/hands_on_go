package rest

import "encoding/json"

type createUserRequest struct {
	FirstName             *string      `json:"firstName"`
	LastName              *string      `json:"lastName"`
	Age                   *json.Number `json:"age"`
	PhoneNumber           *string      `json:"phoneNumber"`
	IsPhoneNumberVerified *bool        `json:"isPhoneNumberVerified"`
}

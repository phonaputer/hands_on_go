package rest

import "encoding/json"

type getUserByIDResponse struct {
	FirstName       string      `json:"firstName"`
	LastName        string      `json:"lastName"`
	Age             json.Number `json:"age"`
	PhoneNumber     string      `json:"phoneNumber"`
	IsPhoneVerified bool        `json:"isPhoneVerified"`
}

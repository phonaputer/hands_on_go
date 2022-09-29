package presentation

type createUserRequest struct {
	FirstName       *string `json:"firstName"`
	LastName        *string `json:"lastName"`
	Age             *int    `json:"age"`
	PhoneNumber     *string `json:"phoneNumber"`
	IsPhoneVerified *bool   `json:"isPhoneVerified"`
}

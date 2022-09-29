package logic

type UserServiceStubImpl struct{}

func (u *UserServiceStubImpl) GetByID(id int) (*User, error) {
	return &User{
		ID:            123,
		FirstName:     "first",
		LastName:      "last",
		Age:           456,
		PhoneNumber:   "+81 555-5555",
		PhoneVerified: true,
	}, nil

	// Use this return to test the Not Found error case:
	//return nil, ErrNotFound

	// Use this return to test the unexpected error case:
	// return nil, errors.New("something has failed")
}

func (u *UserServiceStubImpl) CreateUser(user *User) (int, error) {
	return 123, nil
}

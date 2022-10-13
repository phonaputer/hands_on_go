package presentation

import (
	"github.com/stretchr/testify/assert"
	"hands_on_go/internal/uaerr"
	"io"
	"net/http"
	"strings"
	"testing"
)

var TestUserValidatorImpl_ValidateCreateUser_InvalidData_ReturnsError_table = []struct {
	Name            string
	InputJSON       string
	ExpectedUserMsg string
}{
	{
		Name: "phone number too short",
		InputJSON: `{
			"firstName": "first",
			"lastName": "last",
			"age": 1,
			"phoneNumber": "",
			"isPhoneVerified": false
		}`,
		ExpectedUserMsg: "phoneNumber has invalid length",
	},
	{
		Name: "first name too short",
		InputJSON: `{
			"firstName": "",
			"lastName": "last",
			"age": 1,
			"phoneNumber": "a",
			"isPhoneVerified": false
		}`,
		ExpectedUserMsg: "firstName has invalid length",
	},
	{
		Name: "last name is missing",
		InputJSON: `{
			"firstName": "a",
			"age": 1,
			"phoneNumber": "a",
			"isPhoneVerified": false
		}`,
		ExpectedUserMsg: "lastName is required",
	},
	{
		Name: "age is too large",
		InputJSON: `{
			"firstName": "a",
			"lastName": "last",
			"age": 201,
			"phoneNumber": "a",
			"isPhoneVerified": false
		}`,
		ExpectedUserMsg: "age has invalid size",
	},
	{
		Name: "age is too small",
		InputJSON: `{
			"firstName": "a",
			"lastName": "last",
			"age": 0,
			"phoneNumber": "a",
			"isPhoneVerified": false
		}`,
		ExpectedUserMsg: "age has invalid size",
	},
	{
		Name:            "invalid json",
		InputJSON:       `{`,
		ExpectedUserMsg: "invalid json body",
	},
}

func TestUserValidatorImpl_ValidateCreateUser_InvalidData_ReturnsError(t *testing.T) {
	for _, testCase := range TestUserValidatorImpl_ValidateCreateUser_InvalidData_ReturnsError_table {
		t.Run(testCase.Name, func(t *testing.T) {
			body := io.NopCloser(strings.NewReader(testCase.InputJSON))

			_, err := (&UserValidatorImpl{}).ValidateCreateUser(&http.Request{Body: body})

			assert.Equal(t, uaerr.TypeInvalidInput, uaerr.GetType(err))
			msg, ok := uaerr.GetUserMsg(err)
			assert.True(t, ok)
			assert.Equal(t, testCase.ExpectedUserMsg, msg)
		})
	}
}

// TODO add success cases for age == 1 and age == 200

var TestUserValidatorImpl_ValidateCreateUser_ValidRequest_NoError_table = []struct {
	Name      string
	InputJSON string
	// not shown here: check that the fields of the struct result are equal to the values from the JSON
}{
	{
		Name: "age is 1",
		InputJSON: `{
			"firstName": "a",
			"lastName": "a",
			"age": 1,
			"phoneNumber": "a",
			"isPhoneVerified": false
		}`,
	},
	{
		Name: "age is 200",
		InputJSON: `{
			"firstName": "a",
			"lastName": "a",
			"age": 200,
			"phoneNumber": "a",
			"isPhoneVerified": false
		}`,
	},
}

func TestUserValidatorImpl_ValidateCreateUser_ValidRequest_NoError(t *testing.T) {
	for _, testCase := range TestUserValidatorImpl_ValidateCreateUser_ValidRequest_NoError_table {
		t.Run(testCase.Name, func(t *testing.T) {
			body := io.NopCloser(strings.NewReader(testCase.InputJSON))

			_, err := (&UserValidatorImpl{}).ValidateCreateUser(&http.Request{Body: body})

			assert.Nil(t, err)
			// not shown here: check that the fields of the struct result are equal to the values from the JSON
		})
	}
}

func TestUserValidatorImpl_ValidateCreateUser_PhoneNumberTooShort_ReturnsError(t *testing.T) {
	inputJSON := `{
			"firstName": "first",
			"lastName": "last",
			"age": 1,
			"phoneNumber": "",
			"isPhoneVerified": false
		}`

	body := io.NopCloser(strings.NewReader(inputJSON))

	_, err := (&UserValidatorImpl{}).ValidateCreateUser(&http.Request{Body: body})

	// Expecting an error with a specific user message
	// and type uaerr.TypeInvalidInput
	assert.Equal(t, uaerr.TypeInvalidInput, uaerr.GetType(err))
	msg, ok := uaerr.GetUserMsg(err)
	assert.True(t, ok)
	assert.Equal(t, "phoneNumber has invalid length", msg)
}

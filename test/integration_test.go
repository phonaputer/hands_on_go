package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type getNameJSONResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func TestIntegration_GetUserID1_ShouldReturnCorrectJSONData(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/get-user?id=1")
	if !assert.Nil(t, err, "failed to send HTTP request to localhost:8080/get-user: %s", err) {
		return
	}
	defer resp.Body.Close()

	// Verify request
	if !assert.Equal(t, 200, resp.StatusCode, "unexpected status code") {
		return
	}

	var responseBody getNameJSONResponse
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if !assert.Nil(t, err, "failed to JSON parse response body: %s", err) {
		return
	}

	assert.Equal(t, 1, responseBody.ID)
	assert.Equal(t, "Risa", responseBody.FirstName)
	assert.Equal(t, "Rakuten", responseBody.LastName)
}

func TestIntegration_GetUserID2_ShouldReturnCorrectJSONData(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/get-user?id=2")
	if !assert.Nil(t, err, "failed to send HTTP request to localhost:8080/get-user: %s", err) {
		return
	}
	defer resp.Body.Close()

	// Verify request
	if !assert.Equal(t, 200, resp.StatusCode, "unexpected status code") {
		return
	}

	var responseBody getNameJSONResponse
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if !assert.Nil(t, err, "failed to JSON parse response body: %s", err) {
		return
	}

	assert.Equal(t, 2, responseBody.ID)
	assert.Equal(t, "Taro", responseBody.FirstName)
	assert.Equal(t, "Rakuten", responseBody.LastName)
}

func TestIntegration_GetUserID3_ShouldReturnCorrectJSONData(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/get-user?id=3")
	if !assert.Nil(t, err, "failed to send HTTP request to localhost:8080/get-user: %s", err) {
		return
	}
	defer resp.Body.Close()

	// Verify request was successful
	if !assert.Equal(t, 200, resp.StatusCode, "unexpected status code") {
		return
	}

	// Verify JSON body
	var responseBody getNameJSONResponse
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if !assert.Nil(t, err, "failed to JSON parse response body: %s", err) {
		return
	}

	assert.Equal(t, 3, responseBody.ID)
	assert.Equal(t, "Quiz", responseBody.FirstName)
	assert.Equal(t, "Testersson", responseBody.LastName)
}

func TestIntegration_GetNonexistentUser_ShouldReturn500(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/get-user?id=123")
	if !assert.Nil(t, err) {
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, 500, resp.StatusCode, "unexpected status code")
}

func TestIntegration_GetWithoutID_ShouldReturn500(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/get-user")
	if !assert.Nil(t, err) {
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, 500, resp.StatusCode)
}

func TestIntegration_GetWithoNonIntegerID_ShouldReturn500(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/get-user?id=abc")
	if !assert.Nil(t, err) {
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, 500, resp.StatusCode, "unexpected status code")
}

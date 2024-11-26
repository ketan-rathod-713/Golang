package test

import (
	"auth/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {
	Api, err := InitialiseTestEnvironment()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// create request and then accept
	var reqBody models.SignupRequest = models.SignupRequest{
		Email:    "trial@gmail.com",
		Password: "Trial@123",
		Name:     "Trial 1",
		Phone:    "12345",
		Address:  "Times square thaltej",
		City:     "Ahmedabad",
		State:    "Gujarat",
		Country:  "India",
		Zip:      "398282",
		Standard: "12",
		Role:     "user",
	}

	// Convert the request body to JSON
	reqBodyMarshalled, err := json.Marshal(reqBody)
	if err != nil {
		t.Errorf("Error marshalling request body: %v", err)
		return
	}

	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(reqBodyMarshalled))
	req.Header.Set("Content-Type", "application/json")

	responseRecorder := httptest.NewRecorder()

	Api.AuthApi.HandleSignUp(responseRecorder, req)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", responseRecorder.Code)
	}
	// Api.AuthApi.HandleSignUp // This is my handler
}

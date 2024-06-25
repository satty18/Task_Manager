package controllers

import (
	"Task_Manager/models"
	"Task_Manager/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	utils.InitDB()

	user := models.User{
		Username: "testuser",
		Password: "testpassword",
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonUser))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	if _, ok := response["token"]; !ok {
		t.Errorf("Expected token in response")
	}
}

func TestLogin(t *testing.T) {
	utils.InitDB()

	user := models.User{
		Username: "testuser",
		Password: "testpassword",
	}

	// Register the user first
	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonUser))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)
	handler.ServeHTTP(rr, req)

	// Attempt to login
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUser))
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Login)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	if _, ok := response["token"]; !ok {
		t.Errorf("Expected token in response")
	}
}

package controllers

import (
	"Task_Manager/middleware"
	"Task_Manager/models"
	"Task_Manager/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTask(t *testing.T) {
	utils.InitDB()

	// Create a user and get a token
	user := models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonUser))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)
	handler.ServeHTTP(rr, req)
	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	token := response["token"]

	task := models.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "To Do",
	}
	jsonTask, _ := json.Marshal(task)
	req, _ = http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonTask))
	req.Header.Set("Authorization", "Bearer "+token)
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(CreateTask)
	middleware.AuthMiddleware(handler).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var createdTask models.Task
	json.NewDecoder(rr.Body).Decode(&createdTask)
	if createdTask.Title != task.Title {
		t.Errorf("Expected task title to be %v, got %v", task.Title, createdTask.Title)
	}
}

func TestGetTasks(t *testing.T) {
	utils.InitDB()

	// Create a user and get a token
	user := models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonUser))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)
	handler.ServeHTTP(rr, req)
	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	token := response["token"]

	req, _ = http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetTasks)
	middleware.AuthMiddleware(handler).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var tasks []models.Task
	json.NewDecoder(rr.Body).Decode(&tasks)
	if len(tasks) != 0 {
		t.Errorf("Expected no tasks, got %v", len(tasks))
	}
}

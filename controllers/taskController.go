package controllers

import (
	"Task_Manager/models"
	"Task_Manager/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	var tasks []models.Task
	utils.DB.Where("user_id = ?", userID).Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	if task.Title == "" || task.Status == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	task.UserID = userID
	utils.DB.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	utils.DB.First(&task, params["id"])
	json.NewDecoder(r.Body).Decode(&task)
	utils.DB.Save(&task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	utils.DB.Delete(&task, params["id"])
	json.NewEncoder(w).Encode("Task deleted")
}

func SearchTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	query := r.URL.Query().Get("q")
	var tasks []models.Task
	utils.DB.Where("user_id = ? AND (title LIKE ? OR description LIKE ?)", userID, "%"+query+"%", "%"+query+"%").Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func SortTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	sortBy := r.URL.Query().Get("sort_by")
	var tasks []models.Task
	utils.DB.Where("user_id = ?", userID).Order(sortBy).Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

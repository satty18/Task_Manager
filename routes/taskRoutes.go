package routes

import (
	"Task_Manager/controllers"

	"github.com/gorilla/mux"
)

func InitTaskRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/search", controllers.SearchTasks).Methods("GET")
	r.HandleFunc("/tasks/sort", controllers.SortTasks).Methods("GET")
}

package routes

import (
	"Task_Manager/controllers"

	"github.com/gorilla/mux"
)

func InitUserRoutes(r *mux.Router) {
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/profile", controllers.UpdateProfile).Methods("PUT")
}

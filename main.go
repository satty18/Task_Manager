package main

import (
	"Task_Manager/middleware"
	"Task_Manager/routes"
	"Task_Manager/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	utils.InitDB()

	// Create a new router
	r := mux.NewRouter()

	// Initialize user routes
	routes.InitUserRoutes(r)

	// Initialize task routes with authentication middleware
	taskRouter := r.PathPrefix("/tasks").Subrouter()
	taskRouter.Use(middleware.AuthMiddleware)
	routes.InitTaskRoutes(taskRouter)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

## Task Manager Application Overview 
The Task Manager Application is a
full-stack application that allows users to create, update, delete, and
manage tasks. Users can also filter tasks by status, search tasks by
query, sort tasks, and manage their profiles with avatars. The backend
is built using Golang and the frontend is built using React. 

## Table of Contents • Backend Setup

• Frontend Setup
• Features
• API Endpoints
• Testing
• Environment Variables
• License

## Backend Setup Prerequisites

• Golang (version 1.16 or higher)

## Installation 
• Clone the repository:git clone \<repository-url\>
cd task-manager-backend

• Install dependencies: go mod tidy

• Run the server:go run main.go

The server will start on http://localhost:8080. API Endpoints User

## Authentication 
• Register a new user 

POST `/register` 

Request Body:
`code{ \"username\": \"example\", \"password\": \"password\" }
`

• Login and get a JWT token

`POST /login`

`Request Body:{
\"username\": \"example\", \"password\": \"password\" }`

* Update user profile (requires JWT token)

`PUT /profile Request Body:{
\"username\": \"newUsername\", \"avatar\": \"avatarURL\" }
`
## Task Management (Protected)
• Get all tasks
`GET /tasks`

Create a new task
`POST /tasks`

`Request Body:{ \"title\":
\"New Task\", \"description\": \"Task description\", \"status\": \"To
Do\", \"due_date\": \"2022-12-31\" }
`

• Update a task `PUT /tasks/{id}`

` Request Body:{ \"title\":
\"Updated Task\", \"description\": \"Updated description\", \"status\":
\"In Progress\", \"due_date\": \"2023-01-31\" }
`

• Delete a task

 `DELETE /tasks/{id}`

• Search tasks by query 
  `GET /tasks/search?q=query`

• Sort tasks by specified field

  `GET /tasks/sort?sort_by=field`

Example Request Headers To access protected routes, include the JWT
token in the Authorization header: 
Authorization: Bearer
\<your-jwt-token\>

## Database 
The application uses SQLite for data storage. The database file
is named tasks.db and is automatically created in the project root
directory. Security • Passwords are hashed using bcrypt before being
stored in the database.

• JWT tokens are used for authentication and authorization.

## Frontend Setup Prerequisites
• Node.js (version 14 or higher)

• npm (version 6 or higher) or yarn (version 1.22 or higher)

Installation • Clone the repository:git clone \<repository-url\>
cd task-manager-frontend

• Install dependencies: npm install

• Run the application:npm start

The application will start on http://localhost:3000.

## Features 
### Task Management 
• Create, update, and delete tasks

• Filter tasks by status

• Search tasks by query

• Sort tasks by specified fields

• Set due dates for tasks

User Management 
• Register and login
• Update user profile with avatars

Responsive Design 
• The application is responsive and works well on both
desktop and mobile devices.

API Integration The frontend interacts with the backend API hosted at
http://localhost:8080. Ensure the backend server is running before
starting the frontend application. Testing Backend Testing To run tests
for the backend, use the go test command: go test ./\...

Frontend Testing To run tests for the frontend, use the following
commands: npm run lint npm run format Environment Variables
Create a .env file in the root of your project to store environment
variables: REACT_APP_API_URL=http://localhost:8080 License This
project is licensed under the MIT License.

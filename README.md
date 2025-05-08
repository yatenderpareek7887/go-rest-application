# go-rest-application
A simple RESTful API built with Go.This project demonstrates a clean architecture with separated concerns for controllers, services, models, and data transfer objects (DTOs).

Table of Contents::
User CRUD operations (Create, Read, Update, Delete)
JSON-based API responses
Swagger documentation for API endpoints
Modular and scalable codebase
Error handling with generic error types
Configuration for storage as a database

Tech::
Go: Backend programming language
Gin: HTTP web framework
Swagger: API documentation
JSON: Data storage (users.json)

Project Structure::
.
├── DB
│   └── users.json              # JSON file for user data storage
├── README.md                   # Project documentation
├── go.mod                      # Go module dependencies
├── go.sum                      # Dependency checksums
└── src
    ├── config
    │   ├── storage             # Storage configuration
    │   │   └── storage.go
    │   └── swagger             # Swagger setup
    │       └── SetupSwagger.go
    ├── controllers
    │   └── user-controller     # User-related API endpoints
    │       └── user-controller.go
    ├── docs                    # Swagger documentation files
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── dto
    │   └── user                # Data transfer objects for users
    │       └── user_dto.go
    ├── genric_error            # Generic error handling
    │   └── genric_error.go
    ├── main                    # Main application entry
    ├── main.go                 # Entry point for the application
    ├── models
    │   └── user                # User data model
    │       └── user.go
    ├── routes                  # API route definitions
    │   └── routes.go
    ├── services                # Business logic for user operations
    │   └── user_service.go
    └── utils                   # Utility functions


Setup:: 
Clone the repository:
git clone https://github.com/yatenderpareek7887/go-rest-application.git
cd go-rest-application

Ensure Go is installed::
Install dependencies:
go mod tidy

Running the Application::
Start the server
go run src/main.go

Access::
The API will be available::
http://localhost:8080/*

Access Swagger documentation::
For detailed endpoint information, refer to the Swagger UI at
http://localhost:8080/swagger/api/docs/index.html

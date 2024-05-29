# RBAC API

This project is a simple Role-Based Access Control (RBAC) API built with Go. The API provides basic user management, authentication using JWT, and role-based authorization.

## Project Structure

```bash
rbac-api/
├── main.go
├── controllers/
│ └── userController.go
├── models/
│ └── user.go
├── routes/
│ └── routes.go
├── utils/
│ └── db.go
├── middlewares/
│ └── authMiddleware.go
└── auth/
└── auth.go
```

## Features

- User registration
- User login with JWT authentication
- Role-based authorization for protected routes

## Requirements

- Go 1.22+
- Gorilla Mux for HTTP routing
- GORM for ORM
- SQLite as the database (can be replaced with any other database supported by GORM)
- bcrypt for password hashing
- jwt-go for JWT handling

## Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/rbac-api.git
   cd rbac-api
   ```
2. Install dependencies:
   ```sh
   go mod download
   or
   go mod tidy
   ```
3. Run the application:
   ```sh
   go run main.go
   ```
The application will start listening on port 3000.

## API Endpoints

### User Registration
- URL: `/users`
- Method: `POST`
- Body:
  ```json
  {
    "username": "admin",
    "password": "password",
    "role_id": 1
  }
  ```
- Successful response:
  - Code: `200 OK`
  - Content:
    ```json
    {
      "ID": 1,
      "Username": "admin",
      "RoleID": 1,
      "CreatedAt": "2024-05-28T00:00:00Z",
      "UpdatedAt": "2024-05-28T00:00:00Z"
    }
    ```
### User Login
- URL: `/login`
- Method: `POST`
- Body:
  ```json
  {
    "username": "admin",
    "password": "password"
  }
  ```
- Successful response:
  - Code: `200 OK`
  - Content:
    ```json
    {
    "token": "your.jwt.token.here"
    }
    ```
### Get Users (Protected)
- URL: `/users`
- Method: `GET`
- Headers:
  - Authorization: `Bearer your.jwt.token.here`
- Successful response:
  - Code: `200 OK`
  - Content:
    ```json
    [
      {
        "ID": 1,
        "Username": "admin",
        "Role": {
            "ID": 1,
            "Name": "admin"
        },
        "CreatedAt": "2024-05-28T00:00:00Z",
        "UpdatedAt": "2024-05-28T00:00:00Z"
      }
    ]
    ```
## Middleware

### Authentication Middleware

- The `IsAuthorized` middleware checks if the user the required role to access the route.
- Example usage:
```go
router.HandleFunc("/users", middlewares.IsAuthorized("admin", controllers.GetUsers)).Methods("GET")
```

## Author

- [Aleksander Palamar](https://aleksanderpalamar.dev)
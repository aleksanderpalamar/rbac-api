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
3. Environment variables:
   ```sh
   DB_HOST=db
   DB_PORT=5432
   DB_USER=<your-username>
   DB_PASSWORD=<your-password>
   DB_NAME=<your-database-name>
   ```
4. Run the application:
   ```sh
   go run main.go
   ```
The application will start listening on port 3000.

## Docker and Docker Compose

This project includes a Dockerfile and docker-compose.yml file to help you manage common task more easily. Here's a list of the available commands and a brief description of what they do:

- `docker build -t rbac-api .`: Builds the Docker image.
- `docker run -d -p 3001:3001 -e DB_HOST=db -e DB_PORT=5432 -e DB_USER=admin -e DB_PASSWORD=password -e DB_NAME=rbac rbac-api`: Runs the Docker container.
- `docker-compose up -d`: Builds and runs the Docker containers.
- `docker-compose down`: Stops and removes the Docker containers.

## API Endpoints

### User Registration
- URL: `/users`
- Method: `POST`
- Body:
  ```json
  {
    "username": "admin",
    "password": "password",
    "role": "admin",
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
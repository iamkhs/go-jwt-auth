# User Authentication and JWT-based API with Gin and GORM

This is a simple REST API built using **Gin** (a web framework for Go), **GORM** (Go ORM), and **JWT** (JSON Web Token) for authentication. The API allows users to register, log in, and access protected routes. The routes are secured using JWT, which is issued upon successful login and must be included in the Authorization header for access to protected routes.

## Features
- **User Registration**: Allows users to register with a username, email, and password. Passwords are hashed using **bcrypt** before being stored.
- **User Login**: Users can log in with their email and password. If credentials are valid, a JWT token is returned.
- **JWT Authentication**: Protects specific routes with JWT validation. Only authenticated users with a valid token can access protected routes.
- **Profile Endpoint**: A protected route to get the user's profile information (username, email) after authentication.

## Technologies Used
- **Go**: Programming language to build the API.
- **Gin**: Web framework for Go to handle routing and middleware.
- **GORM**: Object-relational mapping (ORM) library for interacting with MySQL.
- **JWT**: For authentication using JSON Web Tokens.
- **bcrypt**: For securely hashing and comparing passwords.
- **MySQL**: Relational database to store user credentials.
- **dotenv**: To load environment variables from a `.env` file.

## Prerequisites
- Go (1.18 or higher)
- MySQL database
- Postman or any API testing tool (optional)

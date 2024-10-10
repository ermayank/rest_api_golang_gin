# Event Management System - Golang & Gin

Event Management System is a REST API built using Golang and the Gin framework to streamline the process of managing events. This system allows users to register, log in, and manage events with full CRUD functionality. The API provides secure user authentication using JWT, ensuring that only authorized users can create, update, or delete events. Public events are accessible to all users, making it easy to view upcoming events without logging in.

Designed for scalability and security, this API is suitable for any event-driven application, whether it's for organizing conferences, workshops, or community gatherings.

## Features

- **User Registration & Authentication**: Users can register and log in using secure JWT tokens.
- **Event Creation & Management**: Authenticated users can create, update, and delete events.
- **Event Listing**: Publicly available events can be viewed by any user.
- **JWT Token Authentication**: Certain routes are protected and require valid JWT tokens for access.
- **CRUD Operations for Events**: Full control over event lifecycle—Create, Read, Update, Delete.
- **Clean & Scalable API Design**: The codebase follows a modular structure, making it easy to scale and maintain.

## Installation & Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/ermayank/rest_api_golang_gin.git
   ```
2. **Install Dependencies**: Make sure you have Go installed. Then install the required Go packages:
    ```bash
   go mod tidy
   ```
3. **Install Dependencies**:
    ```bash
   go run main.go
   ```
  
 ### Available APIs

- GET /events
- GET /events/<id>
- POST /events
- PUT /events/<id>
- DELETE /events/<id>
- POST /signup
- POST /login
- POST /events/<id>/register
- DELETE /events/<id>/register
   

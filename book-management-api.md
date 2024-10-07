
# Book Management API with JWT Authentication

This is a simple REST API for managing books, implemented in Go. It features JWT (JSON Web Token) authentication for securing endpoints and supports operations such as creating, updating, deleting, and fetching books.

## Features

- **Authentication**: JWT-based authentication.
- **CRUD Operations**: Create, Read, Update, Delete books.
- **Middleware**: Secures certain routes with JWT middleware.
- **Gorilla Mux**: HTTP request router and dispatcher.

## Prerequisites

To run this application, you need:

- Go 1.16 or later
- [Gorilla Mux](https://github.com/gorilla/mux)
- [JWT-go](https://github.com/golang-jwt/jwt)

## Getting Started

1. **Clone the repository:**

    ```bash
    git clone https://github.com/AdomBoateng/book-management-api.git
    cd book-management-api
    ```

2. **Install the dependencies:**

    ```bash
    go mod tidy
    ```

3. **Run the application:**

    ```bash
    go run main.go
    ```

4. **API will be available at**: `http://localhost:8080`

## API Endpoints

### Authentication

- **Login:**
  - URL: `/login`
  - Method: `POST`
  - Payload:
    ```json
    {
      "username": "admin",
      "password": "password"
    }
    ```
  - Response: Returns a JWT token if login is successful.

### Books

- **Get all books:**
  - URL: `/books`
  - Method: `GET`
  - Public endpoint, no token required.

- **Get a single book by ID:**
  - URL: `/books/{id}`
  - Method: `GET`
  - Public endpoint, no token required.
  
- **Create a new book:**
  - URL: `/books`
  - Method: `POST`
  - Secured with JWT. Must provide token in the `Authorization` header.
  - Payload:
    ```json
    {
      "title": "Book Title",
      "author": "Author Name",
      "year": "2023"
    }
    ```

- **Update an existing book by ID:**
  - URL: `/books/{id}`
  - Method: `PUT`
  - Secured with JWT. Must provide token in the `Authorization` header.
  - Payload:
    ```json
    {
      "title": "Updated Book Title",
      "author": "Updated Author Name",
      "year": "2023"
    }
    ```

- **Delete a book by ID:**
  - URL: `/books/{id}`
  - Method: `DELETE`
  - Secured with JWT. Must provide token in the `Authorization` header.

## JWT Authentication

The API requires a valid JWT token for certain endpoints (Create, Update, and Delete books). To get a token, send a POST request to `/login` with the credentials:
- Username: `admin`
- Password: `password`

The token will be returned in the response, and should be used in the `Authorization` header for secured endpoints:

`Authorization`: Bearer<your_token>

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---


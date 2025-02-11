# BookStore API

This is a simple Book Store API built with **Golang** and **GORM** for database management. It allows users to manage books with CRUD operations.

## Features
- Add, update, delete, and retrieve books
- Uses **GORM** for database interaction
- Secure database credentials using `.env` file

## Getting Started

### Prerequisites
Ensure you have the following installed:
- [Golang](https://go.dev/doc/install)
- [MySQL](https://www.mysql.com/downloads/)
- [Git](https://git-scm.com/downloads)

### Clone the Repository
```sh
git clone https://github.com/agrawal-2005/bookStore.git
cd bookStore
```

### Setup Environment Variables
Create a `.env` file in the root directory and add your database credentials:
```ini
DB_USER=your_username
DB_PASSWORD=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=bookstore
```

### Install Dependencies
```sh
go mod tidy
```

### Run the Application
```sh
go run cmd/main.go
```

## Database Migration
To create and migrate tables, use:
```sh
go run cmd/migrate.go
```

## API Endpoints
| Method | Endpoint      | Description              |
|--------|--------------|--------------------------|
| GET    | `/books`     | Get all books            |
| GET    | `/books/:id` | Get book by ID           |
| POST   | `/books`     | Add a new book           |
| PUT    | `/books/:id` | Update book by ID        |
| DELETE | `/books/:id` | Delete book by ID        |

## Contributing
Feel free to submit pull requests or open issues to improve this project!

## License
This project is licensed under the MIT License.

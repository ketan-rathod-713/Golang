## Task 6 : MUX - GORM CRUD

This project implements a RESTful API for managing books. It is built using GORM (a Go ORM) for database operations and Mux for routing HTTP requests. The project follows a clean folder structure to maintain code organization.

## EndPoints

### 1. /
| Method    | Route     | Description |
| ----------| --------- | ----------- |
| GET       | /         | Api Status  |

### 2. /book
| Method    | Route     | Description       |
| ----------| --------- | ----------------- |
| GET       | /book/    | Get All Books     |
| GET       | /book/{id}| Get One Book By Id|
| POST      | /book/    | Create Book       |
| PUT       | /book/    | Update One Book   |
| DELETE    | /book/{id}| Delete One Book   |


## Example Of Required .env Variables 

```
PORT="8080"
DB_PORT="5432"
DATABASE="bacancy"
HOST="localhost"
DB_USER="bacancy"
DB_USER_PASSWORD="admin"
DB_SCHEMA_NAME="task6muxgorm"
```

## Running The Project

1. Make sure you have Go and Postgress installed on your system.
2. Clone the repository.
3. Create a .env file with necessary environment variables.
4. Run `go run main.go` to start the server.

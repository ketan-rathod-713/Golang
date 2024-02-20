## Task 6 : MUX - GORM CRUD

This project implements a RESTful API for managing books. It is built using GORM (a Go ORM) for database operations and Mux for routing HTTP requests. The project follows a clean folder structure to maintain code organization.

## EndPoints

### 1. /
| Method    | Route     | Description |
| ----------| --------- | ----------- |
| GET       | /         | Api Status  |

### 2. /book
| Method    | Route     | Description       | ACCESS |
| ----------| --------- | ----------------- | ------ |
| GET       | /book/    | Get All Books     | anyone |
| GET       | /book/{id}| Get One Book By Id| anyone |
| POST      | /book/    | Create Book       | admin  |
| PUT       | /book/    | Update One Book   | admin  |
| DELETE    | /book/{id}| Delete One Book   | admin  |


### 2. /user
| Method    | Route     | Description       | ACCESS |
| ----------| --------- | ----------------- | ------ |
| GET       | /user/    | Get All Users     | user   |
| GET       | /user/{id}| Get One User By Id| user   |
| POST      | /user/    | Create User       | admin  |
| PUT       | /user/    | Update One User   | admin  |
| DELETE    | /user/{id}| Delete One User   | admin  |
| GET       | /user/{userId}/book/{bookId}| Issue One Book   | admin |

### 3. /auth

| Method    | Route     | Description       | ACCESS |
| ----------| --------- | ----------------- | ------ |
| POST       | /auth/login    | Login by userId & Password (client will get token as cookie for further authorization purpose)    | anyone |

## Example Of Required .env Variables 

```
PORT="8080"
DB_PORT="5432"
DATABASE="bacancy"
HOST="localhost"
DB_USER="bacancy"
DB_USER_PASSWORD="admin"
DB_SCHEMA_NAME="task6muxgorm"
JWT_SECRET="secret"
```

## Running The Project

1. Make sure you have Go and Postgress installed on your system.
2. Clone the repository.
3. Create a .env file with necessary environment variables.
4. Run `go run main.go` to start the server.

## Example Api Requests

If you are using thunder client then you can import some api requests using thunderClientApi.json file provided.
# GoLang Fiber MongoDB API for User CRUD Operations

This is a simple API built using GoLang, Fiber, and MongoDB to perform CRUD (Create, Read, Update, Delete) operations on a user table.

## Prerequisites
Before running the API, make sure you have the following installed:

- GoLang
- MongoDB
- Fiber (Fiber is a web framework for Go)

## API EndPoints

### 1. /
| Method    | Route     | Description |
| ----------| --------- | ----------- |
| GET       | /         | Api Status  |

### 2. /user
| Method    | Route| Description       |
| ----------| -----| ----------------- |
| GET       | /    | Get All Users     |
| GET       | /{id}| Get One User By Id|
| POST      | /    | Create User       | 
| PUT       | /{id}| Update One User   | 
| DELETE    | /{id}| Delete One User   | 


## Example Of Required .env Variables 

```
PORT="8080"
MONGO_URI="mongodb://localhost:27017"
DATABASE="bacancy"
HOST="localhost"
```

## Running The Project

1. Make sure you have Go and Mongodb installed on your system.
2. Clone the repository.
3. Create a .env file with necessary environment variables.
4. Run `go run main.go` to start the server.

## Example Api Requests




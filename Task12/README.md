# GRPC Project

In this project i am going to develop 3 services, which then will communicate using grpc.

## Project Structure

- /proto Folder will be having all the proto files and their golang code will be stored inside /proto folder.

## Prerequisites
Before running the API, make sure you have the following installed:

- GoLang
- MongoDB
- Mux Router

## API EndPoints

### 1. AUTH SERVER 

#### REST API :  /api/v1/

| Method    | Route| Description       |
| ----------| -----| ----------------- |
| POST       | /signin    | Signin User and return jwt token for authorization    |
| POST       | /signup| Signup user|

#### GRPC Services

```
service Auth {
    rpc AuthoriseUser(AuthoriseRequest) returns (AuthoriseResponse);
    rpc GetUserDetails(UserDetailsRequest) returns (UserDetailsResponse);
};
```

### 2. BOOK SERVER

#### REST API :  /api/v1/

| Method    | Route| Description       |
| ----------| -----| ----------------- |
| POST       | /book?token=""    | Create Book If User is admin and is authorized |

#### GRPC Services

Book Service uses auth service for authorization purposes.

## Example Of Required .env Variables 

View .env-example Files for more details.

## Running The Project

1. Clone the repository.
2. Create a .env file with necessary environment variables. (view .env-example for getting started)
3. Start Auth Service
4. cd auth-server
5. go run main.go
6. Start Book Service
7. cd book-server
8. go run main.go

## TODO

- Should i use different database for different microservices ? How should i manage it ?!
- What if particular service is stopped ! How data will be reflected in other service when it starts ?
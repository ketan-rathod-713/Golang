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

## Models

### User

| Name      | Type |
| ----------| -----|
| Id        | primitive.ObjectId  |
| Name      | string |
| Location  | string   | 
| Title     | string| 
| Age       | int| 
| FavoriteGames | []string | 
| Hobby     | Hobby| 


### Hobby

| Name      | Type |
| ----------| -----|
| Name      | string|
| Years     | int |

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

You can import example api requests in postman using `task8fiber_postman_collection.json` file in current directory.

## Notes

### Mongodb

#### Array Manipulation Techniques

- To project or return all the elements from array use $ ( positional operator )
- To update all elements in an array, see the all positional operator $[] instead.
- To update all elements that match an array filter condition or conditions, see the filtered positional operator instead $[<identifier>]

```
db.students.updateOne(
   { _id: 1, grades: 80 },
   { $set: { "grades.$" : 82 } }
)
```
Above one will update first grade element whose value is 80.

We can also update embedded documents using "array.$.field"

##### Update Embedded Documents Using Multiple Field Matches

```
 $elemMatch: { grade: { $lte: 90 }, mean: { $gt: 80 } }
 ```
 Hence for multiple field matches of array we need to use $elemMatch Operator.

 ##### The all positional operator $[]

 The update operator will modify all array elements.
 
 `{ <update operator>: { "<array>.$[]" : value } }`
 Example `  { $inc: { "grades.$[].std" : -2 } },`

- It can be used with nested arrays. 

- update all the array elements which are not equal to 100
```
db.results.updateMany(
   { "grades" : { $ne: 100 } },
   { $inc: { "grades.$[]": 10 } },
)
```

#####

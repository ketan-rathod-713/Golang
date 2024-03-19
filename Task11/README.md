# Task 11 Mongodb Trains 


## Endpoints

| Method    | Route     | Description |
| ----------| --------- | ----------- |
| GET       | /         | Api Status  |
| GET       | /train         | Get all trains  |

### Query Parameters For GET /train

- page : for getting page number
- limit : limit number of records per page
- sort : sort field name. for eg. trainName, number etc.
- order : for ASC 1 and for DESC -1
- search : search text

## Example .env File

```
PORT=8080
DB_URL="mongodb://localhost:27017"
```

## Running Project

### Start Server

```
cd Task11
go run main.go --readcsv --path All_Indian_Trains.csv // Do it only once
go run main.go  // for starting server
```

### Start React Server

```
cd Task11
cd frontend
npm install
npm start
```
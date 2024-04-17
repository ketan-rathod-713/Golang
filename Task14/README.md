# Task 14 Wikepedia Search

Search Functionality for wikipedia. Search any keyword and get scrapped data from wikipedia which will include Images and Paragraphs information.

## Running Code

```
// start backend
cd backend
go run main.go

// start client
cd client
start live server for html page.
```

## Testing Code

```
cd Task14/backend/handlers
go test -coverprofile cover.out
go tool cover -html cover.out
```

## Valid Api Routes

- localhost:8080/wikipedia_scrap
    - require: {"url": "Enter url to be scrapped"}
    - Content-Type : application/json

- localhost:8080/wikipedia_search
    - require: {"searchText":"Enter the keyword you want to search"}
    - Content-Type : application/json
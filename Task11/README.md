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
```
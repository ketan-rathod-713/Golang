# Graphql

It is the simple example of meetup show casing the usecase of graphql. 


## Some Important Commands

```
source .env
migrate create -ext sql -dir postgres/migrations create_meetups;

// Finally to migrate
migrate --path "postgres/migrations" --database "$POSTGRESQL_URL" up

// For using dataloaden run below command inside graph directory.
go run github.com/vektah/dataloaden UserLoader string '[]*meetmeup/models.User';
```

## Dataloader

It effectively reduces the number of calls to database.

## Go Pg Package

It is a GORM used in this project.

## Running The Program

1. create database and migrate using below commands

```
source .env
migrate --path "postgres/migrations" --database "$POSTGRESQL_URL" up
```

2. start server

```
go run main.go
```

3. Open URL shown in console and see playground and try making different queries.
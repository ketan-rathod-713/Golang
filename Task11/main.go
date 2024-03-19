package main

import (
	"context"
	"encoding/csv"
	"errors"
	"flag"
	"io"
	"log"
	"mongodbpagination/app"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Use this boolean flag to decide wheather to read csv and upload to database or start server and listen.
	var flagCsvRead bool
	var flagPath string // from which path to read

	flag.BoolVar(&flagCsvRead, "readcsv", false, "Use this flag when you want to load trains data from csv. It should be done once only. mention file name as the value")
	flag.StringVar(&flagPath, "path", "All_Indian_Trains.csv", "From which path should i read")

	flag.Parse()

	if flagCsvRead {
		// Do connection with mongodb
		// read line by line csv
		// insert to mongodb
		env := app.LoadEnv()

		client, err := app.InitDB(env)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("database connected successfully.")

		db := app.GetDB(client)
		trains := app.GetCollectionTrain(db)

		// Now start reading csv
		file, err := os.Open(flagPath)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(trains, file.Name())

		// create csv reader to read csv file line by line
		reader := csv.NewReader(file)

		// todo Begin transaction from here

		session, err := client.StartSession()
		if err != nil {
			log.Fatal(err)
		}

		session.StartTransaction()

		for {
			record, err := reader.Read()
			// if end of file then break
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				log.Fatal(err)

				// todo roll back transaction
				session.AbortTransaction(context.Background())
				session.EndSession(context.Background())
			}

			//TODO ingore first record
			if record[0] == "" {
				continue
			}

			log.Println(record)

			trainNumber, err := strconv.Atoi(record[1])
			if err != nil {
				log.Fatal(err)

				// todo roll back transaction
				session.AbortTransaction(context.Background())
				session.EndSession(context.Background())
			}
			// insert data to mongodb trains collection
			trains.InsertOne(context.Background(), bson.M{
				"number":      trainNumber,
				"name":        record[2],
				"source":      record[3],
				"destination": record[4],
			})
		}

		session.CommitTransaction(context.Background())
		session.EndSession(context.Background())
		// todo End transaction here

		log.Println("Data Uploaded Successfully")

	} else {
		// start main server
		app.StartServer()
	}
}

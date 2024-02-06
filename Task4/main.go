package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// user.json unmarshal struct
type User struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type Address struct {
	Area    string `json:"area"`
	Country string `json:"country"`
}

// tech.json unmarshal struct
type Tech struct {
	Id       int          `json:"id"`
	TechDets []TechDetail `json:"techDets`
}

type TechDetail struct {
	Tech string  `json:"tech"`
	Exp  float64 `json:"exp"`
}

func main() {
	// Get data of all 3 files

	// Then Merge them based on id
	// Normal case -> all files have equal and all ids
	// Some files have some id and some have some other ids
	// Let user.json -> 1 to 3 id and tech.json this id and so on

	// Get all the id's
	// Data structure for storing id and relevent other information for particular file
	// Then can easily iterate over all id's and gather data

	// Lets start

	userFile, err := os.ReadFile("user.json")

	if err != nil {
		panic(err)
	}

	var user []User
	err = json.Unmarshal([]byte(userFile), &user)
	if err != nil {
		panic(err)
	}
	// got data in user
	fmt.Println(user)

	// Same for other 2 files
	techFile, err := os.ReadFile("tech.json")

	if err != nil {
		panic(err)
	}

	var tech []Tech
	err = json.Unmarshal([]byte(techFile), &tech)
	if err != nil {
		panic(err)
	}
	// got data in user
	fmt.Println(tech)

}

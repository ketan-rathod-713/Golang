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
	Id         int          `json:"id"`
	TechDetail []TechDetail `json:"techDets"`
}

type TechDetail struct {
	Tech string  `json:"tech"`
	Exp  float64 `json:"exp"`
}

// for final output Marshaling
type TechDetailFinal struct {
	Tech string  `json:"techdata"`
	Exp  float64 `json:"exp"`
}

// contact.json unmarshal struct
type Contact struct {
	Id             int            `json:"id"`
	ContactDetails ContactDetails `json:"contactDets"`
}

type ContactDetails struct {
	Email string `json:"email"`
	Phone Phone  `json:"phone"`
}

// Final output struct
type UserInfo struct {
	UserId      int               `json:"Userid"`
	Name        string            `json:"Name"`
	Address     Address           `json:"Address"`
	TechDetails []TechDetailFinal `json:"TechDetails"`
	Email       string            `json:"Email"`
	Phone       Phone             `json:"Phone"`
}

// Special marshal logic for phone

type Phone string

// func (p Phone) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(fmt.Sprintf("+91-%v", p))
// }

// do marshal on userInfo // Data structure to json creation
var COUNTRY_CODE map[string]string = map[string]string{
	"IND": "+91",
	"UK":  "+41",
}

func (u *UserInfo) MarshalJSON() ([]byte, error) { // With Pointers it will work
	countryCode := COUNTRY_CODE[u.Address.Country]

	// Pre-processing before marshaling:
	u.Phone = Phone(fmt.Sprintf("%v-%v", countryCode, u.Phone))

	type UserInfo2 UserInfo              // creating different type
	return json.Marshal((*UserInfo2)(u)) // Passing as a different type so that it calls default method and not this one only

	// return json.Marshal(newUser) // it will give stack overflow as this looks for custom marshalJson method each time for given type and thus producing error hence I need to create a new type and then marshal it.
}

func main() {

	userFile, err := os.ReadFile("user.json")

	if err != nil {
		panic(err)
	}

	var users []User
	err = json.Unmarshal([]byte(userFile), &users)
	if err != nil {
		panic(err)
	}
	// got data in user
	// fmt.Println(users)

	// Same for Tech
	techFile, err := os.ReadFile("tech.json")

	if err != nil {
		panic(err)
	}

	var techs []Tech
	err = json.Unmarshal([]byte(techFile), &techs)
	if err != nil {
		panic(err)
	}
	// got data in user
	fmt.Println(techs)

	// Same for contact.json
	contactFile, err := os.ReadFile("contact.json")

	if err != nil {
		panic(err)
	}

	var contacts []Contact
	err = json.Unmarshal([]byte(contactFile), &contacts)
	if err != nil {
		panic(err)
	}
	// got data in user
	// fmt.Println(contacts)

	// Now merge all
	// Simple Join of users, contacts, and techs

	var userInfos []UserInfo
	for _, user := range users {
		for _, tech := range techs {
			for _, contact := range contacts {
				if user.Id == tech.Id && user.Id == contact.Id { // if tino referencing to same one then merge

					// modify tech details
					// Transform TechDetail into TechDetailFinal
					fmt.Println("Tech details", tech.TechDetail)
					var techFinalDetails []TechDetailFinal
					for _, t := range tech.TechDetail {
						fmt.Println(t)
						techFinalDetails = append(techFinalDetails, TechDetailFinal{t.Tech, t.Exp})
					}

					newUserInfo := UserInfo{UserId: user.Id, Name: user.Name, Address: user.Address, TechDetails: techFinalDetails, Email: contact.ContactDetails.Email, Phone: contact.ContactDetails.Phone}
					userInfos = append(userInfos, newUserInfo)
				}
			}
		}
	}

	fmt.Println(userInfos)

	userInfoBytes, err := json.Marshal(userInfos)
	if err != nil {
		panic(err)
	}

	WriteJSONToFile(userInfoBytes)

	// Alternative 2
	// can we do it using map // store information got in map and update it

}

func WriteJSONToFile(jsonBytes []byte) error {
	err := os.WriteFile("result.json", jsonBytes, os.FileMode(0644))
	fmt.Println("\nFinal Output is in result.json\n")

	return err
}

package main

import (
	"database/sql"
	"log"
	"main/database"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type User struct { // if * then it can result in nil pointer reference when using with scan
	Id          string
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	DateOfBirth string
}

type Student struct {
	Name string
	Age  int
}

// Global db variable for database operations
var db *sql.DB

func main() {

	db = database.ConnectDb()

	// File Server To host static files.
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs) //TODO: StripPrefix creates a new handler. And how it will see it relative to our folder and price.

	// Form Handler
	http.HandleFunc("/form", formHandler)

	// Users Data Handler
	http.HandleFunc("/users", usersDataHandler)

	log.Fatal(http.ListenAndServe(":8080", nil)) // Here nil  because we are setting up http2 here hence no need to define it.

	log.Println("Server Started On Port :8080")
}

// Show All The Users Data To Client
func usersDataHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT *  FROM httpnet.user;"

	rows, err := db.Query(query)

	if err != nil {
		http.Error(w, "An Error Occured", http.StatusInternalServerError)
	}

	defer rows.Close()

	users := []User{} // array of users
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.DateOfBirth)

		if err != nil {
			http.Error(w, "An Error Occured", http.StatusInternalServerError)
		}

		users = append(users, user)
	}

	t, err := template.ParseFiles("users.html")

	if err != nil {
		http.Error(w, "An Error Occured In Parsing Template", http.StatusInternalServerError)
	}

	t.Execute(w, users)
}

// Handles Post Request To get Form data and insert it to database
func formHandler(w http.ResponseWriter, r *http.Request) {

	// Form Data Upload Post Request
	if r.Method == http.MethodPost {

		// Parse Form
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error Occured During Parsing", http.StatusBadRequest)
			return
		}

		log.Println("POST SUCCESS")

		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		dateofbirth := r.FormValue("dateofbirth")
		email := r.FormValue("email")
		phone := r.FormValue("phone")

		log.Println("GOT", firstname, lastname, dateofbirth, email, phone)

		// Now connected hence submit form // yahi pe stop ho jao
		err := submitForm(firstname, lastname, dateofbirth, email, phone)

		if err != nil {
			http.Error(w, "Invalid Entry", http.StatusBadRequest)
		} else {
			http.ServeFile(w, r, "./static/success.html")
		}
	} else {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
	}
}

func submitForm(firstname string, lastname string, dateofbirth string, email string, phone string) error {
	query := `INSERT INTO httpnet.user(firstname, lastname, dateofbirth, email, phone) VALUES($1, $2, $3, $4, $5);`

	_, err := db.Exec(query, firstname, lastname, dateofbirth, email, phone)
	if err != nil {
		return err
	}

	return nil
}

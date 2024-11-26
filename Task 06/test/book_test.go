package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"task6MuxGorm/models"
	"testing"
)

type testcase struct {
	Data interface{} 
	Name string
	RequiredStatusCode int
}

const URL = "http://localhost:8080"

func TestBookCreate(t *testing.T) {
	testCases := []testcase{
		testcase{
			Data: models.Book{
				Title: "testing book",
				Author: "someone",
			},
			Name: "create-book-without-isbn",
			RequiredStatusCode: 200,
		},
		testcase{
			Data: models.Book{
				Title: "testing book",
				Author: "someone",
				ISBN: "abcd",
			},
			Name: "create-book-without-isbn",
			RequiredStatusCode: 200,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.Name, func(t *testing.T) {
			body, _ := json.Marshal(testcase.Data)
            resp, err := http.Post(URL + "/book/", "application/json", bytes.NewReader(body))
            if err!= nil {
                t.Fatalf("failed to create book: %v", err)
            }
            defer resp.Body.Close()

            if resp.StatusCode!= testcase.RequiredStatusCode {
                t.Fatalf("expected status code %d, got %d", testcase.RequiredStatusCode, resp.StatusCode)
            } else {
				fmt.Println("test case passed for creating book")
			}
        })
	}
}

func TestBookGetAll(t *testing.T) {

	res, err := http.Get(fmt.Sprintf("%v/book/", URL))
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, res.StatusCode)
	} else {
		fmt.Println("Book Get All Test Case Passed mahesh")
	}
}

func TestBookUpdate(t *testing.T) {

	var testcase models.Book = models.Book{
		ID:     1,
		Title:  "good",
		Author: "mice",
	}

	body, err := json.Marshal(testcase)

	req, err := http.NewRequest("PUT", URL+"/book/", bytes.NewReader(body))
	if err != nil {
		t.Error(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error("error occured")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Error("status codes do not match")
	} else {
		fmt.Println("Book Update book passed")
	}
}

func TestBookDelete(t *testing.T) {

	client := &http.Client{}
	const Id = 19
	req, err := http.NewRequest("DELETE", URL+"/book/"+fmt.Sprintf("%v", Id), nil)
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, res.StatusCode)
	} else {
		fmt.Println("Book Delete Test Case Failed")
	}
}

package test

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

// Check consistency of database by issueing books at the same time.
func issueBook(userId uint64) {

	for i := 1; i <= 11; i++ {
		query := fmt.Sprintf("http://localhost:8081/user/%v/book/%v", userId, i)
		_, err := http.Get(query)
		if err != nil {
			log.Println("ERROR OCCURED", userId, i)
		}

		fmt.Println("DONE FOR ", userId, i)
	}
}

func TestIssueBook(t *testing.T) {
	// make 100 requests to database server to issue book from 100 diffferent users and we will be having 10 books with different quantities less then 3.
	// all will try to take all 10 books
	// lets see if this is consistent

	// for consistency book initial == book issued

	for i := 1; i <= 11; i++ {
		go issueBook(uint64(i))
	}

	time.Sleep(time.Second * 4)
}

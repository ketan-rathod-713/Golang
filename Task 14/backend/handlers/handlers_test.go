package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wikipediasearch/models"
)

// Test wikipedia search handler
func TestHandleWikipediaSearch(t *testing.T) {
	// Test casese with json payloads
	testCases := []struct {
		name          string
		payload       string
		expectedCode  int
		expectedBody  string
		expectedError bool
	}{
		{
			name:          "Valid JSON",
			payload:       `{"searchText": "usa"}`,
			expectedCode:  http.StatusOK,
			expectedError: false,
		},
		{
			name:          "Valid JSON",
			payload:       `{"searchText": "usain"}`,
			expectedCode:  http.StatusOK,
			expectedError: false,
		},
		{
			name:          "Empty Search Text",
			payload:       `{"searchText": ""}`,
			expectedCode:  http.StatusBadRequest,
			expectedError: true,
		},
		{
			name:          "Invalid JSON",
			payload:       `{"searchTextu": ""}`,
			expectedCode:  http.StatusBadRequest,
			expectedError: true,
		},
		{
			name:          "Empty Payload",
			payload:       ``,
			expectedCode:  http.StatusBadRequest,
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var searchInput models.SearchInput

			// for testing purpose // if get status ok then test
			json.Unmarshal([]byte(tc.payload), &searchInput)

			reader := strings.NewReader(tc.payload)
			request := httptest.NewRequest("POST", "/wikipedia_search", reader)

			request.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder to record the handler's response
			responseRecorder := httptest.NewRecorder()

			// Call the handler function with the recorder and request
			HandleWikipediaSearch(responseRecorder, request)

			// Check the response status code
			if responseRecorder.Code != tc.expectedCode {
				t.Errorf("Expected status code %d; got %d", tc.expectedCode, responseRecorder.Code)
			} else {

				// if status code is expected then
				// Read response body
				body, err := ioutil.ReadAll(responseRecorder.Body)
				if err != nil {
					t.Errorf("Error reading response body: %v", err)
				}

				var apiResponse models.ApiResponse
				err = json.Unmarshal(body, &apiResponse)

				if err != nil {
					t.Errorf("Error decoding response body: %v body %v", err, string(body))
				}

				// i will get map[string]interface{}
				data, valid := apiResponse.Data.(map[string]interface{})

				if !valid && tc.expectedError == false {
					t.Errorf("Not valid api response")
				} else {
					var gotSearchText string
					if _, ok := data["searchText"].(string); ok {
						gotSearchText = data["searchText"].(string)
					}

					var gotSearchTitles []string
					if _, ok := data["searchTitles"].([]string); ok {
						gotSearchTitles = data["searchTitles"].([]string)
					}

					var gotSearchUrls []string
					if _, ok := data["searchURLs"].([]string); ok {
						gotSearchUrls = data["searchURLs"].([]string)
					}

					// if length is not equal to 10 then error
					if len(gotSearchTitles) != len(gotSearchUrls) {
						t.Errorf("Search titles is not equal to search urls")
					}

					if gotSearchText != searchInput.SearchText {
						t.Errorf("Search text passed is different from the got.")
					}
				}

			}

		})
	}

}


// Test wikipediascrap handlers
func TestHandleWikipediaScrap(t *testing.T) {
	testCases := []struct {
		name          string
		payload       string
		expectedCode  int
		expectedBody  string
		expectedError bool
	}{
		{
			name:          "Valid JSON Invalid URL",
			payload:       `{"url": "usa"}`,
			expectedCode:  http.StatusBadRequest,
			expectedError: true,
		},
		{
			name:          "Valid JSON Valid URL",
			payload:       `{"url":  "https://en.wikipedia.org/wiki/Usa"}`,
			expectedCode:  http.StatusOK,
			expectedError: false,
		},
		{
			name:          "Valid JSON Valid URL",
			payload:       `{"url":  "https://en.wikipedia.org/wiki/India"}`,
			expectedCode:  http.StatusOK,
			expectedError: false,
		},
		{
			name:          "Invalid JSON Valid URL",
			payload:       `{"urll":  "https://en.wikipedia.org/wiki/Usa"}`,
			expectedCode:  http.StatusBadRequest,
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var scrapInput models.ScrapInput

			// for testing purpose only // if get status ok then check
			json.Unmarshal([]byte(tc.payload), &scrapInput)

			// get reader from string
			reader := strings.NewReader(tc.payload)
			request := httptest.NewRequest("POST", "/wikipedia_search", reader)

			request.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder to record the handler's response
			responseRecorder := httptest.NewRecorder()

			// Call the handler function with the recorder and request
			HandleWikipediaScrap(responseRecorder, request)

			// Check the response status code
			if responseRecorder.Code != tc.expectedCode {
				t.Errorf("Expected status code %d; got %d", tc.expectedCode, responseRecorder.Code)
			}

			// Read response body
			body, err := ioutil.ReadAll(responseRecorder.Body)
			if err != nil {
				t.Errorf("Error reading response body: %v", err)
			}

			if tc.expectedError == false {
				var apiResponse models.ApiResponse
				err = json.Unmarshal(body, &apiResponse)

				if err != nil {
					t.Errorf("Error decoding response body: %v body %v", err, string(body))
				}

				// i will get map[string]interface{}
				data, valid := apiResponse.Data.(map[string]interface{})

				// fmt.Println(reflect.TypeOf(data["url"]))
				if !valid {
					t.Errorf("Not valid api response")
				} else {
					var gotImageLinks []interface{}
					var gotContents []interface{}
					var gotUrl string

					if _, ok := data["contentParagraphs"].([]interface{}); ok {
						gotContents = data["contentParagraphs"].([]interface{})
					}

					if _, ok := data["imageLinks"].([]interface{}); ok {
						gotImageLinks = data["imageLinks"].([]interface{})
					}

					if _, ok := data["url"].(string); ok {
						gotUrl = data["url"].(string)
					}

					if len(gotImageLinks) == 0 {
						fmt.Println(gotImageLinks)
						t.Errorf("Empty image links")
					}

					if len(gotContents) <= 5 {
						fmt.Println(gotContents)
						t.Errorf("Length of cotent is very less")
					}

					if gotUrl != scrapInput.URL {
						t.Errorf("Input url and got url doesn't match")
					}
				}
			}
		})
	}
}

func TestCatchAllRouters(t *testing.T) {

	request := httptest.NewRequest("POST", "/wikipedia_search", nil)
	request.Header.Set("Content-Type", "application/json")

	responseRecorder := httptest.NewRecorder()

	CatchAllRouters(responseRecorder, request)

	// Check the response status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d; got %d", http.StatusOK, responseRecorder.Code)
	}
}

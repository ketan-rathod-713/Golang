package models

import (
	"encoding/json"
)

type Config struct {
	PORT     string
	WIKI_URL string
}

// search input for the handler
type SearchInput struct {
	SearchText string `json:"searchText"`
}

type ScrapInput struct {
	URL string `json:"url"`
}

// raw result getting from wikipedia api
type SearchResult []interface{}

// our structured data that we are sending to client
type SearchResultList struct {
	SearchText   string        `json:"searchText"`
	SearchTitles []interface{} `json:"searchTitles"`
	SearchURLs   []interface{} `json:"searchURLs"`
}

// string to data type conversion
func (s *SearchResultList) UnmarshalJSON(b []byte) error {
	// first of all unmarshal it to list
	var searchResult SearchResult
	err := json.Unmarshal(b, &searchResult)

	if err != nil {
		return err
	}

	// Now define custom logic
	s.SearchText = searchResult[0].(string)
	// fmt.Println(reflect.TypeOf(searchResult[1]))
	searchTitles, ok := searchResult[1].([]interface{})

	if ok {
		s.SearchTitles = searchTitles
	}

	searchUrls, ok := searchResult[3].([]interface{})

	if ok {
		s.SearchURLs = searchUrls
	}

	return nil
}

// only use for testing searchResultListTest
type SearchResultListTest struct {
	SearchText   string        `json:"searchText"`
	SearchTitles []interface{} `json:"searchTitles"`
	SearchURLs   []interface{} `json:"searchURLs"`
	Message      string        `json:"message"`
}

// scrapped data

type ScrappedData struct {
	URL               string   `json:"url"`
	PageLinks         []string `json:"pageLinks"`
	ImageLinks        []string `json:"imageLinks"`
	ContentParagraphs []string `json:"contentParagraphs"`
}

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

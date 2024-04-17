package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"wikipediasearch/models"

	"github.com/PuerkitoBio/goquery"
)

// url for searching on wikipedia
const WIKI_URL = "https://en.wikipedia.org/w/api.php?action=opensearch&search"

// It will search a text from wikipedia url and give search results to user.
// Require json input like {"searchInput": "...."}
// Method : POST
func HandleWikipediaSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.SearchInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		var out = models.ApiResponse{
			Message: "error decoding search text",
			Data:    nil,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&out)
		return
	}

	if input.SearchText == "" {
		var out = models.ApiResponse{
			Message: "Error Getting Search Text",
			Data:    nil,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&out)
		return
	}

	var url = fmt.Sprintf("%v=%v", WIKI_URL, input.SearchText)
	response, err := http.Get(url)

	if err != nil {
		log.Fatal("an error occured", err)
		return
	}

	defer response.Body.Close()

	contentBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("an error occured", err)
		return
	}

	var searchResult models.SearchResultList
	err = json.Unmarshal(contentBytes, &searchResult)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(&models.ApiResponse{
		Message: "Data fetched successfully",
		Data:    searchResult,
	})
}

// It will scrap data from the wikipedia url provided by client.
// Require json input like {"url": "...."}
// Method : POST
func HandleWikipediaScrap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.ScrapInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		var out = models.ApiResponse{
			Message: "error decoding url from payload",
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&out)
		return
	}

	if input.URL == "" {
		var out = models.ApiResponse{
			Message: "empty url not accepted",
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&out)
		return
	}

	var url = input.URL
	response, err := http.Get(url)
	if err != nil {
		var out = models.ApiResponse{
			Message: "not able to get data from wikipedia url",
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&out)
		return
	}
	defer response.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		var out = models.ApiResponse{
			Message: "error scrapping data from wikipedia page",
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&out)
		return
	}

	// scrap data
	var scrappedData models.ScrappedData = models.ScrappedData{}

	// get all image links
	var links = make([]string, 0)
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("src")
		if exists {
			if strings.Contains(link, "icon") || strings.Contains(link, "svg") || strings.Contains(link, "static/images") {
				// do nothing
			} else {
				// by default it gives relative url so refactor it.
				links = append(links, fmt.Sprintf("https:%v", link))

			}
		}
	})

	var texts []string
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		if len(text) >= 100 {
			texts = append(texts, text)
		}
	})

	doc.Find(".references").Each(func(i int, s *goquery.Selection) {
		anchorTag := s.Find("a")

		href, exist := anchorTag.Attr("href")
		if exist {
			fmt.Println("href", href)
		}
	})

	scrappedData.ContentParagraphs = texts
	scrappedData.ImageLinks = links
	scrappedData.URL = input.URL

	json.NewEncoder(w).Encode(&models.ApiResponse{
		Message: "successfully scrapped data",
		Data:    scrappedData,
	})
}

// handle all other routes
func CatchAllRouters(w http.ResponseWriter, r *http.Request) {
	var out = models.ApiResponse{
		Message: "Route is not defined",
		Data:    nil,
	}
	json.NewEncoder(w).Encode(&out)
}

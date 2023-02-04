package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Article represents a blog article
type Article struct {
	Title        string `json:"title"`
	Nickname     string `json:"nickname"`
	CreationDate string `json:"creation_date"`
}

// Articles represents a collection of articles
type Articles []Article

func main() {
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		articles := Articles{
			Article{Title: "Article 1", Nickname: "Author 1", CreationDate: "2022-01-01"},
			Article{Title: "Article 2", Nickname: "Author 2", CreationDate: "2022-01-02"},
			// add up to 20 articles here
		}

		// convert articles to JSON
		json, err := json.Marshal(articles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set the content type and write the response
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

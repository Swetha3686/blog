package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

var articles []Article

func main() {
	http.HandleFunc("/article/", handleArticle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Path[len("/article/"):])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, article := range articles {
		if article.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

package main

import (
	"encoding/json"
	"net/http"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

var articles []Article

func main() {
	http.HandleFunc("/article", handleArticle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var article Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article.ID = len(articles) + 1
	articles = append(articles, article)
	w.WriteHeader(http.StatusCreated)
}

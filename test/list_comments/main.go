package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Comment struct {
	ID        int    `json:"id"`
	ArticleID int    `json:"article_id"`
	Author    string `json:"author"`
	Text      string `json:"text"`
}

var comments []Comment

func main() {
	http.HandleFunc("/comment", handleComment)
	http.HandleFunc("/comments", handleComments)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handleComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comments = append(comments, comment)
	w.WriteHeader(http.StatusCreated)
}

func handleComments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	articleID, err := strconv.Atoi(r.URL.Query().Get("article_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var articleComments []Comment
	for _, comment := range comments {
		if comment.ArticleID == articleID {
			articleComments = append(articleComments, comment)
		}
	}

	json.NewEncoder(w).Encode(articleComments)
}

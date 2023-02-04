package main

import (
	"encoding/json"
	"net/http"
)

type Comment struct {
	ArticleID int    `json:"article_id"`
	Author    string `json:"author"`
	Text      string `json:"text"`
}

var comments []Comment

func main() {
	http.HandleFunc("/comment", HandleComment)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func HandleComment(w http.ResponseWriter, r *http.Request) {
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

package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Comment struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Author   string `json:"author"`
	Text     string `json:"text"`
}

var comments []Comment
var idCounter int

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

	idCounter++
	comment.ID = idCounter
	comments = append(comments, comment)
	w.WriteHeader(http.StatusCreated)
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, comment := range comments {
		if comment.ID == id {
			json.NewEncoder(w).Encode(comment)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

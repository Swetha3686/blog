package main

import (
	"testing"
)

func TestListArticles(t *testing.T) {
	var comment Article
	expectedComment := Article{
		Nickname: "samy",
		Title:    "Hello",
	}
	var err error
	if err != nil {
		t.Errorf("Unexpected error while commenting on comment: %v", err)
	}

	if comment.Nickname != expectedComment.Nickname {
		t.Errorf("Expected comment nickname to be %s, but got %s", expectedComment.Nickname, comment.Nickname)
	}

	if comment.Title != expectedComment.Title {
		t.Errorf("Expected comment content to be %s, but got %s", expectedComment.Title, comment.Title)
	}
}

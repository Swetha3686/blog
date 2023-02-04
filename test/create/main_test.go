package main

import (
	"testing"
)

func TestCreateArticle(t *testing.T) {
	var comment Article
	expectedComment := Article{
		Title:   "Test article",
		Content: "This is a test article.",
	}
	var err error

	if err != nil {
		t.Errorf("Unexpected error while creating article: %v", err)
	}

	if comment.Title != expectedComment.Title {
		t.Errorf("Expected title to be %s, but got %s", comment.Title, expectedComment.Title)
	}

	if comment.Content != expectedComment.Content {
		t.Errorf("Expected content to be %s, but got %s", comment.Content, expectedComment.Content)
	}

}

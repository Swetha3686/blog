package main

import (
	"testing"
)

func TestListArticleComments(t *testing.T) {
	var comment Comment
	expectedComment := Comment{
		Author: "samy",
		Text:   "Hello",
	}
	var err error
	if err != nil {
		t.Errorf("Unexpected error while commenting on comment: %v", err)
	}

	if comment.Author != expectedComment.Author {
		t.Errorf("Expected comment nickname to be %s, but got %s", expectedComment.Author, comment.Author)
	}

	if comment.Text != expectedComment.Text {
		t.Errorf("Expected comment content to be %s, but got %s", expectedComment.Text, comment.Text)
	}
}

package main

import (
	"testing"
)

func TestCommentOnArticle(t *testing.T) {
	var comment Comment
	expectedComment := Comment{
		ArticleID: 46,
		Text:      "Hello",
	}
	var err error
	if err != nil {
		t.Errorf("Unexpected error while commenting on comment: %v", err)
	}

	if comment.ArticleID != expectedComment.ArticleID {
		t.Errorf("Expected comment nickname to be %d, but got %d", expectedComment.ArticleID, comment.ArticleID)
	}

	if comment.Text != expectedComment.Text {
		t.Errorf("Expected comment content to be %s, but got %s", expectedComment.Text, comment.Text)
	}
}

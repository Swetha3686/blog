package main

import "testing"

func TestGetArticleContent(t *testing.T) {
	content := "hello"
	expectedContent := "This is a test article."
	var err error

	if err != nil {
		t.Errorf("Unexpected error while getting article content: %v", err)
	}

	if content != expectedContent {
		t.Errorf("Expected content to be %s, but got %s", expectedContent, content)
	}
}

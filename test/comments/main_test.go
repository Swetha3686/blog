package main

import "testing"

func TestCommentOnComment(t *testing.T) {
	var comment Comment
	expectedComment := Comment{
		ID:       1,
		ParentID: 46,
		Text:     "Hello",
	}
	var err error
	if err != nil {
		t.Errorf("Unexpected error while commenting on comment: %v", err)
	}

	if comment.ID != expectedComment.ID {
		t.Errorf("Expected comment ID to be %d, but got %d", expectedComment.ID, comment.ID)
	}

	if comment.ParentID != expectedComment.ParentID {
		t.Errorf("Expected comment nickname to be %d, but got %d", expectedComment.ParentID, comment.ParentID)
	}

	if comment.Text != expectedComment.Text {
		t.Errorf("Expected comment content to be %s, but got %s", expectedComment.Text, comment.Text)
	}
}

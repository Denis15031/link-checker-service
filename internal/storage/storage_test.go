package storage

import (
	"testing"
)

func TestAddRequestAndGetRequestByID(t *testing.T) {
	// Очищаем состояние перед тестом
	requests = nil
	nextID = 0

	links := []string{"google.com", "example.com"}
	id := AddRequest(links)

	req, exists := GetRequestByID(id)
	if !exists {
		t.Fatalf("Expected request with ID %d to exist", id)
	}

	if req.ID != id {
		t.Errorf("Expected ID %d, got %d", id, req.ID)
	}

	if len(req.Links) != len(links) {
		t.Errorf("Expected %d links, got %d", len(links), len(req.Links))
	}

	for i, link := range req.Links {
		if link != links[i] {
			t.Errorf("Expected link %s at index %d, got %s", links[i], i, link)
		}
	}
}

func TestGetRequestByID_NonExistent(t *testing.T) {
	_, exists := GetRequestByID(999)
	if exists {
		t.Errorf("Expected non-existent ID to return false")
	}
}

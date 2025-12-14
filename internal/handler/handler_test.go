package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckHandler(t *testing.T) {
	body := `{"links":["google.com"]}`
	req := httptest.NewRequest(http.MethodPost, "/check", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	CheckHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %d", w.Code)
	}

	var resp CheckResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if resp.LinksNum != 1 {
		t.Errorf("Expected 1 link, got %d", resp.LinksNum)
	}
}

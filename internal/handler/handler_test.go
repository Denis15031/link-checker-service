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

func TestReportHandler(t *testing.T) {
	// NOTE: Требуется подготовить данные в storage, как в pdfgen_test
	// или замокать pdfgen

	body := `{"links_list":[0]}`
	req := httptest.NewRequest(http.MethodPost, "/report", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	ReportHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/pdf" {
		t.Errorf("Expected Content-Type application/pdf; got %s", contentType)
	}
}

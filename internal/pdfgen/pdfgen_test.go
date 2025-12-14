package pdfgen

import (
	"link-checker-service/internal/storage"
	"testing"
)

func TestGenerateReport(t *testing.T) {
	// Подготовим фиктивные данные в storage
	storage.AddRequest([]string{"test.com"})
	storage.AddRequest([]string{"another-test.com"})

	requestIDs := []int{0, 1}

	pdfData, err := GenerateReport(requestIDs)
	if err != nil {
		t.Fatalf("GenerateReport failed: %v", err)
	}

	if len(pdfData) == 0 {
		t.Error("Generated PDF is empty")
	}
}

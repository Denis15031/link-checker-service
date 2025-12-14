package pdfgen

import (
	"bytes"
	"fmt"
	"link-checker-service/internal/storage"

	"github.com/phpdave11/gofpdf"
)

func GenerateReport(requestIDs []int) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 16)

	for _, id := range requestIDs {
		req, exists := storage.GetRequestByID(id)
		if !exists {
			continue
		}

		title := fmt.Sprintf("Request ID: %d", id)
		pdf.CellFormat(40, 10, title, "1", 0, "", false, 0, "")
		pdf.Ln(-1)

		for _, link := range req.Links {
			pdf.CellFormat(40, 8, link, "0", 0, "", false, 0, "")
			pdf.Ln(-1)
		}
		pdf.Ln(5)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

package handler

import (
	"encoding/json"
	"link-checker-service/internal/checker"
	"link-checker-service/internal/pdfgen"
	"link-checker-service/internal/storage"
	"net/http"
)

type CheckRequest struct {
	Links []string `json:"links"`
}

type CheckResponse struct {
	Links     map[string]string `json:"links"`
	LinksNum  int               `json:"links_num"`
	RequestID int               `json:"request_id"`
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CheckRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result := make(map[string]string)
	for _, link := range req.Links {
		result[link] = checker.CheckURL(link)
	}

	id := storage.AddRequest(req.Links)

	resp := CheckResponse{
		Links:     result,
		LinksNum:  len(req.Links),
		RequestID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

type ReportRequest struct {
	LinksList []int `json:"links_list"`
}

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ReportRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	pdfData, err := pdfgen.GenerateReport(req.LinksList)
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=report.pdf")
	w.Write(pdfData)
}

package main

import (
	"link-checker-service/internal/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/check", handler.CheckHandler)
	http.HandleFunc("/report", handler.ReportHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

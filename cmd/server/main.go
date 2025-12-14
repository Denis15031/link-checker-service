package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"link-checker-service/internal/checker"
	"link-checker-service/internal/handler"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем .env, если файл существует
	_ = godotenv.Load()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	timeoutStr := os.Getenv("CHECK_TIMEOUT_SECONDS")
	if timeoutStr == "" {
		timeoutStr = "5"
	}

	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Fatal("Invalid CHECK_TIMEOUT_SECONDS in .env")
	}

	// Установим таймаут в checker
	checker.SetDefaultTimeout(time.Duration(timeout) * time.Second)

	http.HandleFunc("/check", handler.CheckHandler)
	http.HandleFunc("/report", handler.ReportHandler)

	addr := ":" + port
	log.Printf("Server running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

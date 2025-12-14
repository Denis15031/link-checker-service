package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"link-checker-service/internal/checker"
	"link-checker-service/internal/handler"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		// log.Printf("Warning: .env file not loaded: %v", err)
		// или просто игнорируем ошибку, если .env не обязателен
	}

	// Логгер
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)) // ← исправлено: добавлен второй аргумент

	// Порт
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Таймаут
	timeoutStr := os.Getenv("CHECK_TIMEOUT_SECONDS")
	if timeoutStr == "" {
		timeoutStr = "5"
	}

	timeout, err := time.ParseDuration(timeoutStr + "s")
	if err != nil {
		logger.Error("Invalid CHECK_TIMEOUT_SECONDS in .env", "error", err)
		os.Exit(1)
	}

	checker.SetDefaultTimeout(timeout)

	mux := http.NewServeMux()
	mux.HandleFunc("/check", handler.CheckHandler)
	mux.HandleFunc("/report", handler.ReportHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go func() {
		logger.Info("Starting server", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server error", "error", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logger.Info("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}

	logger.Info("Server exited")
}

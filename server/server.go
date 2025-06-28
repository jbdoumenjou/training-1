package main

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type ItemHandler struct{}

func (h ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"msg": "Item #" + itemID + " retrieved successfully!",
	})
}

// withCORS is an HTTP middleware that adds CORS headers to the response
// and handles preflight (OPTIONS) requests for cross-origin support.
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Allow all origins to access the server.
		// In production, you should replace "*" with a specific domain (e.g., "https://myapp.com").
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Specify which HTTP methods are allowed when accessing the resource.
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Specify which custom headers are allowed in the request.
		// Common ones include "Content-Type" and "Authorization".
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If the request method is OPTIONS, it's a preflight request.
		// Respond immediately without calling the next handler.
		// Browsers automatically send these before actual requests in cross-origin scenarios.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent) // 204 No Content
			return
		}

		// For all other requests, proceed to the next handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	itemHandler := ItemHandler{}
	// https://pkg.go.dev/net/http#ServeMux
	router.HandleFunc("GET /item/{id}", itemHandler.GetItem)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      withCORS(router),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		slog.Info("Starting server on :8080")

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server failed", "error", err)
			stop()
		}
	}()

	<-ctx.Done()

	slog.Info("Shutdown signal received")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	} else {
		slog.Info("Server shutdown gracefully")
	}
}

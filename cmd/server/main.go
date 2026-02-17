package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tasks-api/internal/handlers"
	"tasks-api/internal/storage"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	store := storage.NewMemory()
	h := handlers.New(store)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", h.TasksCollection)
	mux.HandleFunc("/tasks/", h.TaskItem)
	mux.HandleFunc("/health", h.Health)

	server := &http.Server{
		Addr:    ":8080",
		Handler: loggingMiddleware(mux),
	}

	// Запуск сервера в goroutine
	go func() {
		log.Println("server listening on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	// Ожидание сигнала остановки
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server forced to shutdown: %v\n", err)
	}

	log.Println("server exited properly")
}

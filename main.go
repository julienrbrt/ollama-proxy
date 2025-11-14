package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	portEnv   = "OP_PORT"
	authEnv   = "OP_AUTH_TOKEN"
	ollamaEnv = "OP_OLLAMA_URL"
)

func main() {
	port := os.Getenv(portEnv)
	if port == "" {
		port = "11433"
	}

	ollamaURL := os.Getenv(ollamaEnv)
	if ollamaURL == "" {
		ollamaURL = "http://localhost:11434"
	}

	token := os.Getenv(authEnv)
	if token == "" {
		log.Fatal("AUTH_TOKEN environment variable is required")
	}

	proxy, err := newProxy(ollamaURL)
	if err != nil {
		log.Fatalf("failed to create proxy: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(authMiddleware(token))

	r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Printf("Starting proxy server on port %s, forwarding to %s", port, ollamaURL)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// authMiddleware validates Bearer token authorization
func authMiddleware(token string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if parts[1] == "" || parts[1] != token {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// newProxy creates a reverse proxy for the target URL
func newProxy(targetURL string) (*httputil.ReverseProxy, error) {
	if targetURL == "" {
		return nil, fmt.Errorf("target URL cannot be empty")
	}

	target, err := url.Parse(targetURL)
	if err != nil {
		return nil, fmt.Errorf("invalid target URL: %w", err)
	}

	if target.Scheme == "" || target.Host == "" {
		return nil, fmt.Errorf("invalid target URL: missing scheme or host")
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy, nil
}

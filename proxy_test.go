package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/v3/assert"
)

func TestAuthMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		token          string
		authHeader     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid token",
			token:          "test-token-123",
			authHeader:     "Bearer test-token-123",
			expectedStatus: http.StatusOK,
			expectedBody:   "OK",
		},
		{
			name:           "invalid token",
			token:          "test-token-123",
			authHeader:     "Bearer wrong-token",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Unauthorized\n",
		},
		{
			name:           "missing bearer prefix",
			token:          "test-token-123",
			authHeader:     "test-token-123",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Unauthorized\n",
		},
		{
			name:           "missing authorization header",
			token:          "test-token-123",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Unauthorized\n",
		},
		{
			name:           "empty token value",
			token:          "test-token-123",
			authHeader:     "Bearer ",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Unauthorized\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("OK"))
			})

			middleware := authMiddleware(tt.token)
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			rr := httptest.NewRecorder()
			middleware(handler).ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestNewProxy(t *testing.T) {
	tests := []struct {
		name      string
		targetURL string
		wantErr   bool
	}{
		{
			name:      "valid http URL",
			targetURL: "http://localhost:11434",
			wantErr:   false,
		},
		{
			name:      "valid https URL",
			targetURL: "https://example.com",
			wantErr:   false,
		},
		{
			name:      "invalid URL",
			targetURL: "://invalid",
			wantErr:   true,
		},
		{
			name:      "empty URL",
			targetURL: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proxy, err := newProxy(tt.targetURL)
			if tt.wantErr {
				assert.Assert(t, err != nil)
				assert.Assert(t, proxy == nil)
			} else {
				assert.NilError(t, err)
				assert.Assert(t, proxy != nil)
			}
		})
	}
}

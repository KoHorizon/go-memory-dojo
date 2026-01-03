// Package http provides comprehensive documentation and working examples
// for Go's net/http package - the foundation of HTTP in Go.
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"
)
// Common HTTP status codes to memorize
var (
	StatusOK                  = http.StatusOK                  // 200
	StatusCreated             = http.StatusCreated             // 201
	StatusAccepted            = http.StatusAccepted            // 202
	StatusNoContent           = http.StatusNoContent           // 204
	StatusMovedPermanently    = http.StatusMovedPermanently    // 301
	StatusFound               = http.StatusFound               // 302
	StatusNotModified         = http.StatusNotModified         // 304
	StatusBadRequest          = http.StatusBadRequest          // 400
	StatusUnauthorized        = http.StatusUnauthorized        // 401
	StatusForbidden           = http.StatusForbidden           // 403
	StatusNotFound            = http.StatusNotFound            // 404
	StatusMethodNotAllowed    = http.StatusMethodNotAllowed    // 405
	StatusInternalServerError = http.StatusInternalServerError // 500
	StatusBadGateway          = http.StatusBadGateway          // 502
	StatusServiceUnavailable  = http.StatusServiceUnavailable  // 503
)

// Common HTTP methods
var (
	MethodGet     = http.MethodGet     // "GET"
	MethodPost    = http.MethodPost    // "POST"
	MethodPut     = http.MethodPut     // "PUT"
	MethodPatch   = http.MethodPatch   // "PATCH"
	MethodDelete  = http.MethodDelete  // "DELETE"
	MethodHead    = http.MethodHead    // "HEAD"
	MethodOptions = http.MethodOptions // "OPTIONS"
)

// Common headers to know
var CommonHeaders = map[string]string{
	"Content-Type":     "application/json",
	"Accept":           "application/json",
	"Authorization":    "Bearer <token>",
	"X-Request-ID":     "unique-id",
	"Cache-Control":    "no-cache",
	"Content-Length":   "123",
	"User-Agent":       "MyApp/1.0",
}

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF NET/HTTP
// =============================================================================
//
// Go's net/http package embodies Go's philosophy: simple primitives that
// compose well. Unlike frameworks that hide complexity, net/http exposes
// the HTTP protocol directly while providing sensible defaults.
//
// THE TWO SIDES OF HTTP
//
// CLIENT SIDE: "I want to fetch data from a server"
//   - Create a Request
//   - Send it using a Client
//   - Receive a Response
//   - Read the Response body
//
// SERVER SIDE: "I want to respond to incoming requests"
//   - Define Handlers (functions that process requests)
//   - Register them with a multiplexer (router)
//   - Start a Server that listens for connections
//
// THE KEY INSIGHT
//
// Everything revolves around two interfaces:
//   - http.Handler: anything that can handle an HTTP request
//   - io.Reader/io.Writer: request and response bodies are streams
//
// =============================================================================

// =============================================================================
// SECTION 2: THE HTTP CLIENT - BASIC REQUESTS
// =============================================================================
//
// THE DEFAULT CLIENT
//
// http.DefaultClient is a global *http.Client ready to use. It's convenient
// but has NO TIMEOUT by default - requests can hang forever!
// In production, always create your own client with timeouts.
//
// THE REQUEST-RESPONSE CYCLE
//
// 1. Create/configure a request (or use helper functions)
// 2. Send it via a Client
// 3. Check for errors (network errors, DNS failures, etc.)
// 4. Check the status code (200, 404, 500, etc.)
// 5. Read the response body
// 6. CLOSE THE BODY (critical for connection reuse!)
//
// WHY MUST YOU CLOSE THE BODY?
//
// HTTP/1.1 uses persistent connections. The connection can only be reused
// AFTER the current response body is fully read and closed. If you don't:
//   - Connections stay open, wasting resources
//   - You'll run out of file descriptors
//   - Connection pooling breaks down
//
// The pattern: defer resp.Body.Close() immediately after error check.
//
// =============================================================================

func DemonstrateBasicClient() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello!",
			"method":  r.Method,
		})
	}))
	defer server.Close()

	// Simple GET request
	resp, err := http.Get(server.URL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close() // CRITICAL: Always close!

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %d, Body: %s\n", resp.StatusCode, body)

	// POST with form data
	formData := url.Values{
		"username": {"alice"},
		"password": {"secret"},
	}
	resp2, _ := http.PostForm(server.URL+"/login", formData)
	defer resp2.Body.Close()
	body2, _ := io.ReadAll(resp2.Body)
	fmt.Printf("Form POST: %s\n", body2)

	// POST with JSON body
	jsonData := []byte(`{"name": "Bob", "age": 30}`)
	resp3, _ := http.Post(server.URL+"/users", "application/json", bytes.NewBuffer(jsonData))
	defer resp3.Body.Close()
	body3, _ := io.ReadAll(resp3.Body)
	fmt.Printf("JSON POST: %s\n", body3)
}

// =============================================================================
// SECTION 3: CREATING CUSTOM REQUESTS
// =============================================================================
//
// WHY CREATE REQUESTS MANUALLY?
//
// The helper functions (Get, Post, PostForm) are limited:
//   - Can't set custom headers
//   - Can't use methods like PUT, DELETE, PATCH
//   - Can't add context for cancellation/timeouts
//   - Can't add authentication
//
// http.NewRequestWithContext vs http.NewRequest
//
// ALWAYS prefer NewRequestWithContext. It accepts a context.Context that:
//   - Enables request cancellation
//   - Enables request timeouts
//   - Propagates deadlines through your application
//
// http.NewRequest uses context.Background() internally - no cancellation!
//
// =============================================================================

func DemonstrateCustomRequests() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Method=%s Auth=%s", r.Method, r.Header.Get("Authorization"))
	}))
	defer server.Close()

	ctx := context.Background()
	client := &http.Client{Timeout: 10 * time.Second}

	// Custom GET with headers
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/api/data", nil)
	req.Header.Set("Authorization", "Bearer token123")
	req.Header.Set("Accept", "application/json")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("GET with auth: %s\n", body)

	// PUT with JSON body
	userData := map[string]interface{}{"id": 123, "name": "Updated"}
	jsonBody, _ := json.Marshal(userData)

	req2, _ := http.NewRequestWithContext(ctx, http.MethodPut, server.URL+"/users/123", bytes.NewReader(jsonBody))
	req2.Header.Set("Content-Type", "application/json")

	resp2, _ := client.Do(req2)
	defer resp2.Body.Close()
	body2, _ := io.ReadAll(resp2.Body)
	fmt.Printf("PUT: %s\n", body2)

	// DELETE
	req3, _ := http.NewRequestWithContext(ctx, http.MethodDelete, server.URL+"/users/123", nil)
	resp3, _ := client.Do(req3)
	defer resp3.Body.Close()
	body3, _ := io.ReadAll(resp3.Body)
	fmt.Printf("DELETE: %s\n", body3)
}

// =============================================================================
// SECTION 4: THE HTTP CLIENT - CONFIGURATION
// =============================================================================
//
// WHY CREATE A CUSTOM CLIENT?
//
// http.DefaultClient has NO TIMEOUTS. A request can hang forever waiting
// for a slow or unresponsive server. In production, this is catastrophic.
//
// THE TIMEOUT HIERARCHY
//
//   Client.Timeout: Total time from request start to response body read
//   │
//   ├─ Transport.DialContext: Time to establish TCP connection
//   ├─ Transport.TLSHandshakeTimeout: Time for TLS negotiation
//   ├─ Transport.ResponseHeaderTimeout: Time waiting for headers
//   └─ Response body read: Remaining time from Client.Timeout
//
// THE TRANSPORT
//
// http.Transport controls the low-level connection handling:
//   - Connection pooling (MaxIdleConns, MaxIdleConnsPerHost)
//   - Keep-alive behavior
//   - Proxy settings
//   - TLS configuration
//
// =============================================================================

func DemonstrateClientConfiguration() {
	// Production-ready client configuration
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		fmt.Fprintf(w, "OK")
	}))
	defer server.Close()

	// Per-request timeout with context (50ms - will timeout)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	_, err := client.Do(req)
	if err != nil {
		fmt.Printf("Expected timeout: %v\n", err)
	}

	// Longer timeout - succeeds
	ctx2, cancel2 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel2()

	req2, _ := http.NewRequestWithContext(ctx2, http.MethodGet, server.URL, nil)
	resp, _ := client.Do(req2)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Success: %s\n", body)
}

// =============================================================================
// SECTION 5: HANDLING RESPONSES
// =============================================================================
//
// STATUS CODES: WHAT THEY MEAN
//
// 1xx - Informational (rare)
// 2xx - Success: 200 OK, 201 Created, 204 No Content
// 3xx - Redirection: 301 Moved, 302 Found, 304 Not Modified
// 4xx - Client Error: 400 Bad Request, 401 Unauthorized, 403 Forbidden, 404 Not Found
// 5xx - Server Error: 500 Internal, 502 Bad Gateway, 503 Unavailable
//
// IMPORTANT: HTTP ERRORS ARE NOT GO ERRORS
//
// A 404 or 500 response is still a valid HTTP response. client.Do() returns
// an error only for network-level failures. You MUST check resp.StatusCode!
//
// =============================================================================

func DemonstrateResponseHandling() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Custom", "value")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	})
	mux.HandleFunc("/notfound", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
	})
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	server := httptest.NewServer(mux)
	defer server.Close()
	client := &http.Client{Timeout: 10 * time.Second}

	// Success response
	resp, _ := client.Get(server.URL + "/ok")
	fmt.Printf("Status: %d, Custom Header: %s\n", resp.StatusCode, resp.Header.Get("X-Custom"))
	resp.Body.Close()

	// 404 - NOT a Go error!
	resp2, err := client.Get(server.URL + "/notfound")
	fmt.Printf("404 err is nil: %v, Status: %d\n", err == nil, resp2.StatusCode)
	resp2.Body.Close()

	// Proper error handling pattern
	resp3, err := client.Get(server.URL + "/error")
	if err != nil {
		fmt.Printf("Network error: %v\n", err)
		return
	}
	defer resp3.Body.Close()

	if resp3.StatusCode >= 400 {
		fmt.Printf("HTTP error: %d\n", resp3.StatusCode)
		return
	}
}

// =============================================================================
// SECTION 6: THE HTTP SERVER - HANDLERS
// =============================================================================
//
// WHAT IS A HANDLER?
//
// An http.Handler is anything that implements:
//
//   ServeHTTP(w http.ResponseWriter, r *http.Request)
//
// http.ResponseWriter: Where you write your response
//   - Header(): Get response headers (set before WriteHeader)
//   - WriteHeader(statusCode): Set the status code
//   - Write([]byte): Write body (implicitly calls WriteHeader(200))
//
// http.HandlerFunc: THE ADAPTER PATTERN
//
// Converts a function to a Handler. This lets you use plain functions:
//
//   http.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
//       fmt.Fprintf(w, "Hello!")
//   })
//
// =============================================================================

// Handler as a struct
type GreetHandler struct {
	greeting string
}

func (h *GreetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, visitor!", h.greeting)
}

func DemonstrateHandlers() {
	// Handler interface implementation
	handler1 := &GreetHandler{greeting: "Welcome"}
	server1 := httptest.NewServer(handler1)
	resp1, _ := http.Get(server1.URL)
	body1, _ := io.ReadAll(resp1.Body)
	resp1.Body.Close()
	server1.Close()
	fmt.Printf("Struct handler: %s\n", body1)

	// HandlerFunc adapter (most common)
	handler2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello!"})
	})
	server2 := httptest.NewServer(handler2)
	resp2, _ := http.Get(server2.URL)
	body2, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()
	server2.Close()
	fmt.Printf("Function handler: %s\n", body2)
}

// =============================================================================
// SECTION 7: THE HTTP SERVER - ROUTING (MUX)
// =============================================================================
//
// WHAT IS A MUX?
//
// "Mux" is short for "multiplexer" - it routes requests to handlers
// based on the URL path.
//
// PATTERN MATCHING
//
//   "/exact"   - Matches only /exact
//   "/prefix/" - Matches /prefix/ and anything below it
//
// The trailing slash matters!
//
// Go 1.22+ ENHANCED ROUTING
//
//   "GET /users"        - Only matches GET requests
//   "POST /users"       - Only matches POST requests
//   "/users/{id}"       - Captures path parameter
//   "/files/{path...}"  - Captures rest of path (wildcard)
//
// Access parameters with: r.PathValue("id")
//
// =============================================================================

func DemonstrateRouting() {
	mux := http.NewServeMux()

	// Basic routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Home")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About")
	})

	// Prefix matching
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API: %s", r.URL.Path)
	})

	// Go 1.22+ method-specific routes
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "List users")
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Create user")
	})

	// Go 1.22+ path parameters
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "User %s", id)
	})

	// Wildcard
	mux.HandleFunc("/files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		path := r.PathValue("path")
		fmt.Fprintf(w, "File: %s", path)
	})

	server := httptest.NewServer(mux)
	defer server.Close()
	client := &http.Client{}

	// Test routes
	routes := []struct {
		method, path string
	}{
		{"GET", "/"},
		{"GET", "/about"},
		{"GET", "/api/v1/data"},
		{"GET", "/users"},
		{"POST", "/users"},
		{"GET", "/users/123"},
		{"GET", "/files/images/photo.jpg"},
	}

	for _, r := range routes {
		req, _ := http.NewRequest(r.method, server.URL+r.path, nil)
		resp, _ := client.Do(req)
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s %s -> %s\n", r.method, r.path, body)
	}
}

// =============================================================================
// SECTION 8: THE HTTP SERVER - WRITING RESPONSES
// =============================================================================
//
// THE ORDER MATTERS
//
// 1. Set headers (w.Header().Set(...))
// 2. Set status code (w.WriteHeader(code))
// 3. Write body (w.Write(data) or fmt.Fprintf(w, ...))
//
// Once WriteHeader() is called, headers are sent. You can't change them.
// Once Write() is called, it implicitly calls WriteHeader(200) if not yet called.
//
// COMMON PATTERNS
//
// JSON: w.Header().Set("Content-Type", "application/json")
//       json.NewEncoder(w).Encode(data)
//
// Error: http.Error(w, "message", http.StatusBadRequest)
//
// Redirect: http.Redirect(w, r, "/new-path", http.StatusFound)
//
// =============================================================================

func DemonstrateResponses() {
	mux := http.NewServeMux()

	// JSON response
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"name": "Alice",
			"age":  30,
		})
	})

	// Created with Location header
	mux.HandleFunc("/created", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "/users/123")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"id": 123}`)
	})

	// Error response
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	})

	// JSON error
	mux.HandleFunc("/json-error", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "validation failed"})
	})

	// Redirect
	mux.HandleFunc("/old", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new", http.StatusMovedPermanently)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	// Test JSON
	resp, _ := http.Get(server.URL + "/json")
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("JSON: %s\n", strings.TrimSpace(string(body)))

	// Test Created
	resp2, _ := http.Get(server.URL + "/created")
	fmt.Printf("Created: %d, Location: %s\n", resp2.StatusCode, resp2.Header.Get("Location"))
	resp2.Body.Close()
}

// =============================================================================
// SECTION 9: READING REQUEST DATA
// =============================================================================
//
// WHERE DOES DATA COME FROM?
//
// URL PATH: /users/123/orders/456
//   Go 1.22+: r.PathValue("id")
//
// QUERY STRING: /search?q=golang&limit=10
//   r.URL.Query().Get("q")
//
// REQUEST BODY:
//   r.Body is an io.ReadCloser - read once, then consumed
//
// FORM DATA:
//   r.ParseForm() then r.FormValue("field")
//
// HEADERS:
//   r.Header.Get("Authorization")
//
// =============================================================================

func DemonstrateReadingRequests() {
	mux := http.NewServeMux()

	// Path parameters
	mux.HandleFunc("GET /users/{id}/orders/{orderId}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User=%s Order=%s", r.PathValue("id"), r.PathValue("orderId"))
	})

	// Query parameters
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		limit := r.URL.Query().Get("limit")
		tags := r.URL.Query()["tags"] // Multiple values
		fmt.Fprintf(w, "q=%s limit=%s tags=%v", q, limit, tags)
	})

	// JSON body
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var user struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Created: %s <%s>", user.Name, user.Email)
	})

	// Form data
	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintf(w, "Login: %s", r.FormValue("username"))
	})

	// Headers
	mux.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Auth: %s", r.Header.Get("Authorization"))
	})

	server := httptest.NewServer(mux)
	defer server.Close()
	client := &http.Client{}

	// Test path params
	resp, _ := client.Get(server.URL + "/users/123/orders/456")
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Path: %s\n", body)

	// Test query params
	resp2, _ := client.Get(server.URL + "/search?q=golang&limit=10&tags=web&tags=api")
	body2, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()
	fmt.Printf("Query: %s\n", body2)

	// Test JSON body
	resp3, _ := client.Post(server.URL+"/users", "application/json",
		strings.NewReader(`{"name":"Alice","email":"alice@example.com"}`))
	body3, _ := io.ReadAll(resp3.Body)
	resp3.Body.Close()
	fmt.Printf("JSON: %s\n", body3)
}

// =============================================================================
// SECTION 10: MIDDLEWARE
// =============================================================================
//
// WHAT IS MIDDLEWARE?
//
// Middleware wraps a handler to add functionality:
//   - Logging, Authentication, Rate limiting, CORS, Compression
//
// THE MIDDLEWARE PATTERN
//
//   func Middleware(next http.Handler) http.Handler {
//       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//           // Before
//           next.ServeHTTP(w, r)
//           // After
//       })
//   }
//
// CHAINING: handler := Logging(Auth(RateLimit(actualHandler)))
//
// =============================================================================

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[LOG] %s %s %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func DemonstrateMiddleware() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/panic" {
			panic("boom!")
		}
		fmt.Fprintf(w, "OK")
	})

	// Stack: Recovery -> Logging -> Auth -> Handler
	wrapped := RecoveryMiddleware(LoggingMiddleware(AuthMiddleware(handler)))

	server := httptest.NewServer(wrapped)
	defer server.Close()
	client := &http.Client{}

	// Without auth - blocked
	resp, _ := client.Get(server.URL + "/test")
	fmt.Printf("No auth: %d\n", resp.StatusCode)
	resp.Body.Close()

	// With auth - success
	req, _ := http.NewRequest("GET", server.URL+"/test", nil)
	req.Header.Set("Authorization", "Bearer token")
	resp2, _ := client.Do(req)
	body, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()
	fmt.Printf("With auth: %s\n", body)

	// Panic - handled by recovery
	req2, _ := http.NewRequest("GET", server.URL+"/panic", nil)
	req2.Header.Set("Authorization", "Bearer token")
	resp3, _ := client.Do(req2)
	fmt.Printf("Panic recovered: %d\n", resp3.StatusCode)
	resp3.Body.Close()
}

// =============================================================================
// SECTION 11: CONTEXT IN HTTP
// =============================================================================
//
// WHY CONTEXT MATTERS
//
// HTTP requests can be cancelled (client disconnects, timeout, navigation).
// Without context, your handler keeps running after the client is gone.
//
// REQUEST CONTEXT
//
// Every *http.Request has a context: r.Context()
// Cancelled when client closes connection or timeout is exceeded.
//
// PASSING VALUES VIA CONTEXT
//
//   ctx := context.WithValue(r.Context(), "userID", 123)
//   r = r.WithContext(ctx)
//
// Then: userID := r.Context().Value("userID")
//
// =============================================================================

type contextKey string

const requestIDKey contextKey = "requestID"

func RequestIDMiddleware(next http.Handler) http.Handler {
	counter := 0
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		counter++
		ctx := context.WithValue(r.Context(), requestIDKey, fmt.Sprintf("req-%d", counter))
		w.Header().Set("X-Request-ID", fmt.Sprintf("req-%d", counter))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func DemonstrateContext() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Context().Value(requestIDKey).(string)
		fmt.Fprintf(w, "Request ID: %s", reqID)
	})

	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(100 * time.Millisecond):
			fmt.Fprintf(w, "Done")
		case <-r.Context().Done():
			return // Client cancelled
		}
	})

	handler := RequestIDMiddleware(mux)
	server := httptest.NewServer(handler)
	defer server.Close()
	client := &http.Client{}

	// Request ID from context
	resp, _ := client.Get(server.URL + "/")
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Context: %s, Header: %s\n", body, resp.Header.Get("X-Request-ID"))

	// With timeout that expires
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", server.URL+"/slow", nil)
	_, err := client.Do(req)
	fmt.Printf("Timeout: %v\n", err != nil)
}

// =============================================================================
// SECTION 12: SERVER CONFIGURATION
// =============================================================================
//
// BEYOND LISTENANDSERVE
//
// Use http.Server directly for production:
//
//   server := &http.Server{
//       Addr:         ":8080",
//       Handler:      mux,
//       ReadTimeout:  5 * time.Second,
//       WriteTimeout: 10 * time.Second,
//       IdleTimeout:  120 * time.Second,
//   }
//
// GRACEFUL SHUTDOWN
//
// server.Shutdown(ctx) gracefully shuts down:
// 1. Stops accepting new connections
// 2. Waits for active requests to complete
// 3. Respects the context's deadline
//
// =============================================================================

func DemonstrateServerConfig() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	// Production server configuration
	_ = &http.Server{
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	fmt.Println("Server config: ReadTimeout=5s WriteTimeout=10s IdleTimeout=120s")

	// Graceful shutdown pattern (conceptual)
	// go server.ListenAndServe()
	// <-quit // wait for signal
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// server.Shutdown(ctx)
}

// =============================================================================
// SECTION 13: TESTING HTTP
// =============================================================================
//
// HTTPTEST PACKAGE
//
// httptest.NewServer(handler): Creates a real HTTP server on random port
// httptest.NewRecorder(): Captures what a handler writes (no network)
// httptest.NewRequest(): Creates a request for testing
//
// NewRecorder: Unit testing handlers (fast, no network)
// NewServer: Integration testing (real HTTP, tests middleware/routing)
//
// =============================================================================

func DemonstrateTesting() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var data struct {
			Name string `json:"name"`
		}
		json.NewDecoder(r.Body).Decode(&data)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello, " + data.Name})
	})

	// Using ResponseRecorder (unit test)
	req := httptest.NewRequest(http.MethodPost, "/greet", strings.NewReader(`{"name":"Alice"}`))
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	fmt.Printf("Recorder: status=%d body=%s\n", rr.Code, strings.TrimSpace(rr.Body.String()))

	// Using test server (integration test)
	server := httptest.NewServer(handler)
	defer server.Close()
	resp, _ := http.Post(server.URL, "application/json", strings.NewReader(`{"name":"Bob"}`))
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Server: status=%d body=%s\n", resp.StatusCode, strings.TrimSpace(string(body)))
}

// =============================================================================
// SECTION 14: COMMON PATTERNS
// =============================================================================
//
// Real-world patterns that appear constantly in production code.
//
// =============================================================================

func DemonstrateCommonPatterns() {
	// Pattern 1: JSON response helper
	respondJSON := func(w http.ResponseWriter, status int, data interface{}) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(data)
	}

	// Pattern 2: Error response helper
	respondError := func(w http.ResponseWriter, status int, message string) {
		respondJSON(w, status, map[string]string{"error": message})
	}

	mux := http.NewServeMux()

	// Pattern 3: REST resource
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusOK, []string{"alice", "bob"})
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var user struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			respondError(w, http.StatusBadRequest, "Invalid JSON")
			return
		}
		if user.Name == "" {
			respondError(w, http.StatusBadRequest, "Name required")
			return
		}
		respondJSON(w, http.StatusCreated, map[string]string{"id": "123", "name": user.Name})
	})

	// Pattern 4: Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	resp, _ := http.Get(server.URL + "/health")
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Health: %s\n", strings.TrimSpace(string(body)))

	resp2, _ := http.Get(server.URL + "/users")
	body2, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()
	fmt.Printf("Users: %s\n", strings.TrimSpace(string(body2)))
}

// RunAllDemonstrations executes all examples in order.
func RunAllDemonstrations() {
	DemonstrateBasicClient()
	DemonstrateCustomRequests()
	DemonstrateClientConfiguration()
	DemonstrateResponseHandling()
	DemonstrateHandlers()
	DemonstrateRouting()
	DemonstrateResponses()
	DemonstrateReadingRequests()
	DemonstrateMiddleware()
	DemonstrateContext()
	DemonstrateServerConfig()
	DemonstrateTesting()
	DemonstrateCommonPatterns()
}

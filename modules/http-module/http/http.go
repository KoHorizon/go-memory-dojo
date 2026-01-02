// Package http provides comprehensive documentation and working examples
// for Go's net/http package - the foundation of HTTP in Go.
//
// This file covers both client and server sides, with deep explanations
// of WHY things work the way they do.
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

// ============================================================================
// 🎯 TRAINING VARIABLES - HTTP Module
// ============================================================================

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
// WHY IS NET/HTTP DESIGNED THIS WAY?
//
// Go's net/http package embodies Go's philosophy: simple primitives that
// compose well. Unlike frameworks in other languages that hide complexity,
// net/http exposes the HTTP protocol directly while providing sensible
// defaults.
//
// THE TWO SIDES OF HTTP
//
// HTTP is a request-response protocol. There are two perspectives:
//
// CLIENT SIDE: "I want to fetch data from a server"
//   - You create a Request
//   - You send it using a Client
//   - You receive a Response
//   - You read the Response body
//
// SERVER SIDE: "I want to respond to incoming requests"
//   - You define Handlers (functions that process requests)
//   - You register them with a multiplexer (router)
//   - You start a Server that listens for connections
//
// THE KEY INSIGHT
//
// Everything in net/http revolves around two interfaces:
//   - http.Handler: anything that can handle an HTTP request
//   - io.Reader/io.Writer: request and response bodies are streams
//
// Understanding these interfaces unlocks the entire package.
//
// =============================================================================

// =============================================================================
// SECTION 2: THE HTTP CLIENT - BASIC REQUESTS
// =============================================================================
//
// THE DEFAULT CLIENT
//
// http.DefaultClient is a global *http.Client ready to use. It's convenient
// but has NO TIMEOUT by default - requests can hang forever! In production,
// always create your own client with timeouts.
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
// HTTP/1.1 uses persistent connections by default. The connection can only
// be reused for the next request AFTER the current response body is fully
// read and closed. If you don't close it:
//   - The connection stays open, wasting resources
//   - You'll eventually run out of file descriptors
//   - Connection pooling breaks down
//
// The pattern: defer resp.Body.Close() immediately after error check.
//
// =============================================================================

func DemonstrateBasicClient() {
	fmt.Println("=== BASIC HTTP CLIENT ===")
	fmt.Println()

	// We'll use httptest to create a local test server
	// This avoids network dependencies in examples
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "Hello from server", "method": "%s"}`, r.Method)
	}))
	defer server.Close()

	// SIMPLE GET - the http.Get helper function
	// Under the hood: creates a Request, uses DefaultClient, returns Response
	fmt.Println("1. Simple GET request:")
	resp, err := http.Get(server.URL)
	if err != nil {
		// This catches: DNS failures, connection refused, network errors
		// It does NOT catch HTTP errors like 404, 500
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp.Body.Close() // CRITICAL: Always close the body!

	// Read the body - it's an io.Reader, we need to consume it
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("   Error reading body: %v\n", err)
		return
	}

	fmt.Printf("   Status: %s\n", resp.Status)           // "200 OK"
	fmt.Printf("   Status Code: %d\n", resp.StatusCode)  // 200
	fmt.Printf("   Body: %s\n", string(body))
	fmt.Println()

	// SIMPLE POST with form data
	fmt.Println("2. POST with form data:")
	formData := url.Values{}
	formData.Set("username", "alice")
	formData.Set("password", "secret")

	resp2, err := http.PostForm(server.URL+"/login", formData)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp2.Body.Close()

	body2, _ := io.ReadAll(resp2.Body)
	fmt.Printf("   Response: %s\n", string(body2))
	fmt.Println()

	// SIMPLE POST with JSON body
	fmt.Println("3. POST with JSON body:")
	jsonData := []byte(`{"name": "Bob", "age": 30}`)
	resp3, err := http.Post(server.URL+"/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp3.Body.Close()

	body3, _ := io.ReadAll(resp3.Body)
	fmt.Printf("   Response: %s\n", string(body3))
	fmt.Println()
}

// =============================================================================
// SECTION 3: CREATING CUSTOM REQUESTS
// =============================================================================
//
// WHY CREATE REQUESTS MANUALLY?
//
// The helper functions (Get, Post, PostForm) are convenient but limited:
//   - Can't set custom headers
//   - Can't use methods like PUT, DELETE, PATCH
//   - Can't add context for cancellation/timeouts
//   - Can't add authentication
//
// http.NewRequest creates a *Request you can fully customize.
//
// THE REQUEST STRUCT
//
// Key fields you'll use:
//   Method: GET, POST, PUT, DELETE, PATCH, etc.
//   URL:    The target URL (parsed *url.URL)
//   Header: http.Header (map[string][]string)
//   Body:   io.ReadCloser (the request body, if any)
//
// Note: Headers are string slices because HTTP allows multiple values
// for the same header (e.g., multiple Set-Cookie headers).
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
	fmt.Println("=== CUSTOM HTTP REQUESTS ===")
	fmt.Println()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Method: %s, Auth: %s, Custom: %s",
			r.Method,
			r.Header.Get("Authorization"),
			r.Header.Get("X-Custom-Header"))
	}))
	defer server.Close()

	// Creating a custom request with context
	fmt.Println("1. Custom GET with headers:")

	ctx := context.Background() // In real code, use a proper context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/api/data", nil)
	if err != nil {
		fmt.Printf("   Error creating request: %v\n", err)
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer token123")
	req.Header.Set("X-Custom-Header", "custom-value")
	req.Header.Set("Accept", "application/json")

	// Use a client to send the request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("   Response: %s\n", string(body))
	fmt.Println()

	// PUT request with JSON body
	fmt.Println("2. PUT request with JSON body:")

	userData := map[string]interface{}{
		"id":    123,
		"name":  "Updated Name",
		"email": "new@example.com",
	}
	jsonBody, _ := json.Marshal(userData)

	req2, err := http.NewRequestWithContext(ctx, http.MethodPut, server.URL+"/users/123", bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	req2.Header.Set("Content-Type", "application/json")

	resp2, err := client.Do(req2)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp2.Body.Close()

	body2, _ := io.ReadAll(resp2.Body)
	fmt.Printf("   Response: %s\n", string(body2))
	fmt.Println()

	// DELETE request
	fmt.Println("3. DELETE request:")

	req3, _ := http.NewRequestWithContext(ctx, http.MethodDelete, server.URL+"/users/123", nil)
	resp3, err := client.Do(req3)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp3.Body.Close()

	body3, _ := io.ReadAll(resp3.Body)
	fmt.Printf("   Response: %s\n", string(body3))
	fmt.Println()
}

// =============================================================================
// SECTION 4: THE HTTP CLIENT - CONFIGURATION
// =============================================================================
//
// WHY CREATE A CUSTOM CLIENT?
//
// http.DefaultClient has NO TIMEOUTS. A request can hang forever waiting
// for a slow or unresponsive server. In production, this is catastrophic:
//   - Goroutines pile up waiting
//   - Memory grows unbounded
//   - Your service becomes unresponsive
//
// THE TIMEOUT HIERARCHY
//
// There are multiple timeout points:
//
//   Client.Timeout: Total time from request start to response body read
//   │
//   ├─ Transport.DialContext: Time to establish TCP connection
//   ├─ Transport.TLSHandshakeTimeout: Time for TLS negotiation
//   ├─ Transport.ResponseHeaderTimeout: Time waiting for headers
//   └─ Response body read: Remaining time from Client.Timeout
//
// Client.Timeout is the overall limit. If ANY part takes too long, the
// entire request fails.
//
// THE TRANSPORT
//
// http.Transport controls the low-level connection handling:
//   - Connection pooling (MaxIdleConns, MaxIdleConnsPerHost)
//   - Keep-alive behavior
//   - Proxy settings
//   - TLS configuration
//
// The default transport maintains a pool of connections for reuse.
// This is why closing response bodies matters - connections return to the pool.
//
// =============================================================================

func DemonstrateClientConfiguration() {
	fmt.Println("=== CLIENT CONFIGURATION ===")
	fmt.Println()

	// Create a properly configured client
	client := &http.Client{
		Timeout: 30 * time.Second, // Total request timeout

		Transport: &http.Transport{
			MaxIdleConns:        100,              // Total idle connections
			MaxIdleConnsPerHost: 10,               // Idle connections per host
			IdleConnTimeout:     90 * time.Second, // How long idle conns live

			// Connection timeouts
			// DialContext would be set here for dial timeout
			// TLSHandshakeTimeout: 10 * time.Second,
			// ResponseHeaderTimeout: 10 * time.Second,
		},
	}

	fmt.Println("Configured client with:")
	fmt.Println("  - 30s total timeout")
	fmt.Println("  - 100 max idle connections")
	fmt.Println("  - 10 max idle per host")
	fmt.Println("  - 90s idle connection timeout")
	fmt.Println()

	// Using context for per-request timeout
	fmt.Println("Per-request timeout with context:")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response
		time.Sleep(100 * time.Millisecond)
		fmt.Fprintf(w, "Response after delay")
	}))
	defer server.Close()

	// This context will timeout in 50ms - before the server responds
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	_, err := client.Do(req)
	if err != nil {
		fmt.Printf("  Request failed (expected): %v\n", err)
	}
	fmt.Println()

	// Longer timeout - succeeds
	ctx2, cancel2 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel2()

	req2, _ := http.NewRequestWithContext(ctx2, http.MethodGet, server.URL, nil)
	resp, err := client.Do(req2)
	if err != nil {
		fmt.Printf("  Unexpected error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("  Success: %s\n", string(body))
	fmt.Println()
}

// =============================================================================
// SECTION 5: HANDLING RESPONSES
// =============================================================================
//
// STATUS CODES: WHAT THEY MEAN
//
// HTTP status codes are grouped by their first digit:
//
// 1xx - Informational (rare in practice)
//   100 Continue: Server received headers, send the body
//
// 2xx - Success
//   200 OK: Standard success
//   201 Created: Resource created (POST success)
//   204 No Content: Success, but no body (DELETE success)
//
// 3xx - Redirection
//   301 Moved Permanently: Resource moved, update your bookmarks
//   302 Found: Temporary redirect
//   304 Not Modified: Use your cached copy
//
// 4xx - Client Error (YOUR fault)
//   400 Bad Request: Malformed request
//   401 Unauthorized: Need to authenticate
//   403 Forbidden: Authenticated but not allowed
//   404 Not Found: Resource doesn't exist
//   429 Too Many Requests: Rate limited
//
// 5xx - Server Error (THEIR fault)
//   500 Internal Server Error: Server crashed
//   502 Bad Gateway: Upstream server failed
//   503 Service Unavailable: Server overloaded
//   504 Gateway Timeout: Upstream timeout
//
// IMPORTANT: HTTP ERRORS ARE NOT GO ERRORS
//
// A 404 or 500 response is still a valid HTTP response. The client.Do()
// method returns an error only for network-level failures (DNS, connection
// refused, timeout). You MUST check resp.StatusCode separately!
//
// READING RESPONSE HEADERS
//
// resp.Header is an http.Header (map[string][]string).
// Use .Get() for single value, direct access for multiple values.
// Header names are canonicalized: "content-type" becomes "Content-Type".
//
// =============================================================================

func DemonstrateResponseHandling() {
	fmt.Println("=== HANDLING RESPONSES ===")
	fmt.Println()

	// Create a server that returns different status codes
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/ok":
			w.Header().Set("X-Custom", "value1")
			w.Header().Add("Set-Cookie", "session=abc")
			w.Header().Add("Set-Cookie", "user=bob")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"status": "success"}`)
		case "/created":
			w.Header().Set("Location", "/users/123")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, `{"id": 123}`)
		case "/notfound":
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error": "not found"}`)
		case "/error":
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "internal error"}`)
		}
	}))
	defer server.Close()

	client := &http.Client{Timeout: 10 * time.Second}

	// Demonstrating status code checks
	fmt.Println("1. Successful response (200):")
	resp, _ := client.Get(server.URL + "/ok")
	defer resp.Body.Close()

	fmt.Printf("   Status: %s\n", resp.Status)
	fmt.Printf("   Status Code: %d\n", resp.StatusCode)
	fmt.Printf("   Is success? %v\n", resp.StatusCode >= 200 && resp.StatusCode < 300)

	// Reading headers
	fmt.Printf("   X-Custom header: %s\n", resp.Header.Get("X-Custom"))
	fmt.Printf("   All Set-Cookie headers: %v\n", resp.Header["Set-Cookie"])
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("   Body: %s\n", string(body))
	fmt.Println()

	// 201 Created - check Location header
	fmt.Println("2. Created response (201):")
	resp2, _ := client.Get(server.URL + "/created")
	defer resp2.Body.Close()
	fmt.Printf("   Status: %d\n", resp2.StatusCode)
	fmt.Printf("   Location: %s\n", resp2.Header.Get("Location"))
	fmt.Println()

	// 404 - NOT a Go error!
	fmt.Println("3. Not Found (404) - Note: NOT a Go error:")
	resp3, err := client.Get(server.URL + "/notfound")
	if err != nil {
		fmt.Printf("   Network error: %v\n", err) // Won't happen for 404
		return
	}
	defer resp3.Body.Close()
	fmt.Printf("   err is nil: %v\n", err == nil) // true!
	fmt.Printf("   Status: %d\n", resp3.StatusCode)
	body3, _ := io.ReadAll(resp3.Body)
	fmt.Printf("   Body: %s\n", string(body3))
	fmt.Println()

	// Proper error handling pattern
	fmt.Println("4. Proper error handling pattern:")
	resp4, err := client.Get(server.URL + "/error")
	if err != nil {
		// Network error
		fmt.Printf("   Network error: %v\n", err)
		return
	}
	defer resp4.Body.Close()

	if resp4.StatusCode >= 400 {
		// HTTP error - read body for error details
		body4, _ := io.ReadAll(resp4.Body)
		fmt.Printf("   HTTP error %d: %s\n", resp4.StatusCode, string(body4))
		return
	}
	fmt.Println("   Success!")
	fmt.Println()
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
// That's it. One method. This simple interface is the foundation of all
// HTTP servers in Go.
//
// http.ResponseWriter: Where you write your response
//   - Header(): Get the response headers (must set before WriteHeader)
//   - WriteHeader(statusCode): Set the status code
//   - Write([]byte): Write body bytes (implicitly calls WriteHeader(200))
//
// *http.Request: The incoming request
//   - Method: GET, POST, etc.
//   - URL: The request URL
//   - Header: Request headers
//   - Body: Request body (io.ReadCloser)
//
// http.HandlerFunc: THE ADAPTER PATTERN
//
// Writing a struct with ServeHTTP for every handler is tedious.
// http.HandlerFunc is a type that converts a function to a Handler:
//
//   type HandlerFunc func(ResponseWriter, *Request)
//
// It implements Handler by calling itself. This lets you use plain
// functions as handlers:
//
//   http.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
//       fmt.Fprintf(w, "Hello!")
//   })
//
// =============================================================================

// ExampleHandler demonstrates implementing the Handler interface
type ExampleHandler struct {
	greeting string
}

func (h *ExampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, visitor!", h.greeting)
}

func DemonstrateHandlers() {
	fmt.Println("=== HTTP HANDLERS ===")
	fmt.Println()

	// Method 1: Implement Handler interface
	fmt.Println("1. Handler interface implementation:")
	handler1 := &ExampleHandler{greeting: "Welcome"}

	server1 := httptest.NewServer(handler1)
	defer server1.Close()

	resp1, _ := http.Get(server1.URL)
	body1, _ := io.ReadAll(resp1.Body)
	resp1.Body.Close()
	fmt.Printf("   Response: %s\n", string(body1))
	fmt.Println()

	// Method 2: HandlerFunc adapter
	fmt.Println("2. HandlerFunc adapter:")
	handler2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from function!")
	})

	server2 := httptest.NewServer(handler2)
	defer server2.Close()

	resp2, _ := http.Get(server2.URL)
	body2, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()
	fmt.Printf("   Response: %s\n", string(body2))
	fmt.Println()

	// Method 3: Inline function (most common)
	fmt.Println("3. Inline function (common pattern):")
	server3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "Inline handler!"}`)
	}))
	defer server3.Close()

	resp3, _ := http.Get(server3.URL)
	body3, _ := io.ReadAll(resp3.Body)
	resp3.Body.Close()
	fmt.Printf("   Response: %s\n", string(body3))
	fmt.Println()
}

// =============================================================================
// SECTION 7: THE HTTP SERVER - ROUTING (MUX)
// =============================================================================
//
// WHAT IS A MUX?
//
// "Mux" is short for "multiplexer" - it routes requests to handlers based
// on the URL path. http.ServeMux is Go's built-in router.
//
// HOW PATTERN MATCHING WORKS
//
// Patterns can be:
//   "/exact"   - Matches only /exact
//   "/prefix/" - Matches /prefix/ and anything below (/prefix/foo, etc.)
//
// The trailing slash matters:
//   "/api"  matches only "/api"
//   "/api/" matches "/api/", "/api/users", "/api/users/123", etc.
//
// Longer patterns take precedence:
//   "/api/users/" beats "/api/"
//   "/api/users/admin" beats "/api/users/"
//
// Go 1.22+ ENHANCED ROUTING
//
// Go 1.22 added method and path parameter support:
//   "GET /users"        - Only matches GET requests
//   "POST /users"       - Only matches POST requests
//   "/users/{id}"       - Captures path parameter
//   "/files/{path...}"  - Captures rest of path (wildcard)
//
// Access parameters with: r.PathValue("id")
//
// DEFAULT MUX
//
// http.HandleFunc() and http.ListenAndServe() use http.DefaultServeMux.
// For production, create your own mux for:
//   - Avoiding global state
//   - Testing handlers in isolation
//   - Running multiple servers
//
// =============================================================================

func DemonstrateRouting() {
	fmt.Println("=== HTTP ROUTING (MUX) ===")
	fmt.Println()

	// Create a new mux (don't use DefaultServeMux in production)
	mux := http.NewServeMux()

	// Basic routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// "/" matches everything not matched by other patterns
		// Check for exact match if needed
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Home page")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About page")
	})

	// Prefix matching (note the trailing slash)
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API endpoint: %s", r.URL.Path)
	})

	// Go 1.22+ method-specific routes
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "List all users")
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Create user")
	})

	// Go 1.22+ path parameters
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Get user: %s", id)
	})

	mux.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Update user: %s", id)
	})

	mux.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Delete user: %s", id)
	})

	// Wildcard path parameter
	mux.HandleFunc("/files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		path := r.PathValue("path")
		fmt.Fprintf(w, "File path: %s", path)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	// Test the routes
	testRoutes := []struct {
		method string
		path   string
	}{
		{"GET", "/"},
		{"GET", "/about"},
		{"GET", "/api/v1/data"},
		{"GET", "/users"},
		{"POST", "/users"},
		{"GET", "/users/123"},
		{"PUT", "/users/456"},
		{"DELETE", "/users/789"},
		{"GET", "/files/images/photo.jpg"},
	}

	client := &http.Client{}
	for _, tc := range testRoutes {
		req, _ := http.NewRequest(tc.method, server.URL+tc.path, nil)
		resp, _ := client.Do(req)
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("   %s %s -> %s\n", tc.method, tc.path, string(body))
	}
	fmt.Println()
}

// =============================================================================
// SECTION 8: THE HTTP SERVER - WRITING RESPONSES
// =============================================================================
//
// THE ORDER MATTERS
//
// When writing a response, order is critical:
//   1. Set headers (w.Header().Set(...))
//   2. Set status code (w.WriteHeader(code))
//   3. Write body (w.Write(data) or fmt.Fprintf(w, ...))
//
// Once you call WriteHeader(), headers are sent. You can't change them.
// Once you call Write(), it implicitly calls WriteHeader(200) if not yet called.
//
// COMMON PATTERNS
//
// JSON Response:
//   w.Header().Set("Content-Type", "application/json")
//   json.NewEncoder(w).Encode(data)
//
// Error Response:
//   http.Error(w, "message", http.StatusBadRequest)
//
// Redirect:
//   http.Redirect(w, r, "/new-path", http.StatusFound)
//
// File Serving:
//   http.ServeFile(w, r, "path/to/file")
//
// NOT FOUND HANDLING
//
// http.NotFound(w, r) writes a 404 response.
// http.NotFoundHandler() returns a handler that does this.
//
// =============================================================================

func DemonstrateResponses() {
	fmt.Println("=== WRITING RESPONSES ===")
	fmt.Println()

	mux := http.NewServeMux()

	// JSON response
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"name":  "Alice",
			"age":   30,
			"email": "alice@example.com",
		}
		w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder writes directly to w
		json.NewEncoder(w).Encode(data)
	})

	// Custom status code
	mux.HandleFunc("/created", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "/users/123")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
		fmt.Fprintf(w, `{"id": 123}`)
	})

	// Error response helper
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		// http.Error sets Content-Type to text/plain
		// and writes the status code and message
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	})

	// Custom error with JSON
	mux.HandleFunc("/json-error", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "validation_failed",
			"message": "Email is required",
		})
	})

	// Redirect
	mux.HandleFunc("/old-path", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new-path", http.StatusMovedPermanently)
	})

	mux.HandleFunc("/new-path", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You made it to the new path!")
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	// Don't follow redirects automatically for demonstration
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// Test responses
	fmt.Println("1. JSON response:")
	resp1, _ := client.Get(server.URL + "/json")
	body1, _ := io.ReadAll(resp1.Body)
	resp1.Body.Close()
	fmt.Printf("   Content-Type: %s\n", resp1.Header.Get("Content-Type"))
	fmt.Printf("   Body: %s\n", strings.TrimSpace(string(body1)))
	fmt.Println()

	fmt.Println("2. Created response (201):")
	resp2, _ := client.Get(server.URL + "/created")
	body2, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()
	fmt.Printf("   Status: %d\n", resp2.StatusCode)
	fmt.Printf("   Location: %s\n", resp2.Header.Get("Location"))
	fmt.Printf("   Body: %s\n", string(body2))
	fmt.Println()

	fmt.Println("3. Error response (http.Error):")
	resp3, _ := client.Get(server.URL + "/error")
	body3, _ := io.ReadAll(resp3.Body)
	resp3.Body.Close()
	fmt.Printf("   Status: %d\n", resp3.StatusCode)
	fmt.Printf("   Body: %s\n", strings.TrimSpace(string(body3)))
	fmt.Println()

	fmt.Println("4. JSON error response:")
	resp4, _ := client.Get(server.URL + "/json-error")
	body4, _ := io.ReadAll(resp4.Body)
	resp4.Body.Close()
	fmt.Printf("   Status: %d\n", resp4.StatusCode)
	fmt.Printf("   Body: %s\n", strings.TrimSpace(string(body4)))
	fmt.Println()

	fmt.Println("5. Redirect (301):")
	resp5, _ := client.Get(server.URL + "/old-path")
	resp5.Body.Close()
	fmt.Printf("   Status: %d\n", resp5.StatusCode)
	fmt.Printf("   Location: %s\n", resp5.Header.Get("Location"))
	fmt.Println()
}

// =============================================================================
// SECTION 9: READING REQUEST DATA
// =============================================================================
//
// WHERE DOES DATA COME FROM?
//
// HTTP requests can carry data in several places:
//
// URL PATH: /users/123/orders/456
//   - Go 1.22+: r.PathValue("id")
//   - Pre-1.22: Parse from r.URL.Path manually
//
// QUERY STRING: /search?q=golang&limit=10
//   - r.URL.Query() returns url.Values
//   - r.URL.Query().Get("q") for single value
//   - r.URL.Query()["tags"] for multiple values
//
// REQUEST BODY:
//   - r.Body is an io.ReadCloser
//   - Read once, then it's consumed
//   - Close it when done (though servers handle this)
//
// FORM DATA: application/x-www-form-urlencoded or multipart/form-data
//   - Call r.ParseForm() first
//   - Then r.FormValue("field") or r.Form["field"]
//   - For files: r.ParseMultipartForm(maxMemory), then r.FormFile("field")
//
// HEADERS:
//   - r.Header.Get("Authorization")
//   - r.Header["Accept"] for multiple values
//
// COMMON GOTCHA: BODY CONSUMPTION
//
// The request body is a stream. Once read, it's gone. If you need to
// read it multiple times, you must buffer it yourself:
//
//   body, _ := io.ReadAll(r.Body)
//   r.Body = io.NopCloser(bytes.NewBuffer(body)) // Put it back
//
// =============================================================================

func DemonstrateReadingRequests() {
	fmt.Println("=== READING REQUEST DATA ===")
	fmt.Println()

	mux := http.NewServeMux()

	// Path parameters (Go 1.22+)
	mux.HandleFunc("GET /users/{id}/orders/{orderId}", func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("id")
		orderId := r.PathValue("orderId")
		fmt.Fprintf(w, "User: %s, Order: %s", userId, orderId)
	})

	// Query parameters
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		q := query.Get("q")
		limit := query.Get("limit")
		tags := query["tags"] // Multiple values

		fmt.Fprintf(w, "Search: q=%s, limit=%s, tags=%v", q, limit, tags)
	})

	// JSON body
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var user struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		// Decode JSON from request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Created user: %s <%s>", user.Name, user.Email)
	})

	// Form data
	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		// ParseForm populates r.Form and r.PostForm
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form", http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		remember := r.FormValue("remember") == "true"

		fmt.Fprintf(w, "Login: %s (remember: %v)", username, remember)
		_ = password // Don't log passwords!
	})

	// Headers
	mux.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		userAgent := r.Header.Get("User-Agent")
		accept := r.Header.Get("Accept")

		fmt.Fprintf(w, "Auth: %s, UA: %s, Accept: %s", auth, userAgent, accept)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := &http.Client{}

	// Test path parameters
	fmt.Println("1. Path parameters:")
	resp1, _ := client.Get(server.URL + "/users/123/orders/456")
	body1, _ := io.ReadAll(resp1.Body)
	resp1.Body.Close()
	fmt.Printf("   %s\n", string(body1))
	fmt.Println()

	// Test query parameters
	fmt.Println("2. Query parameters:")
	resp2, _ := client.Get(server.URL + "/search?q=golang&limit=10&tags=web&tags=api")
	body2, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()
	fmt.Printf("   %s\n", string(body2))
	fmt.Println()

	// Test JSON body
	fmt.Println("3. JSON body:")
	jsonBody := `{"name": "Alice", "email": "alice@example.com"}`
	resp3, _ := client.Post(server.URL+"/users", "application/json", strings.NewReader(jsonBody))
	body3, _ := io.ReadAll(resp3.Body)
	resp3.Body.Close()
	fmt.Printf("   %s\n", string(body3))
	fmt.Println()

	// Test form data
	fmt.Println("4. Form data:")
	formData := url.Values{}
	formData.Set("username", "bob")
	formData.Set("password", "secret")
	formData.Set("remember", "true")
	resp4, _ := client.PostForm(server.URL+"/login", formData)
	body4, _ := io.ReadAll(resp4.Body)
	resp4.Body.Close()
	fmt.Printf("   %s\n", string(body4))
	fmt.Println()

	// Test headers
	fmt.Println("5. Headers:")
	req5, _ := http.NewRequest("GET", server.URL+"/auth", nil)
	req5.Header.Set("Authorization", "Bearer token123")
	req5.Header.Set("Accept", "application/json")
	resp5, _ := client.Do(req5)
	body5, _ := io.ReadAll(resp5.Body)
	resp5.Body.Close()
	fmt.Printf("   %s\n", string(body5))
	fmt.Println()
}

// =============================================================================
// SECTION 10: MIDDLEWARE
// =============================================================================
//
// WHAT IS MIDDLEWARE?
//
// Middleware is code that wraps a handler to add functionality:
//   - Logging
//   - Authentication
//   - Rate limiting
//   - CORS headers
//   - Compression
//   - Request ID tracking
//   - Panic recovery
//
// THE MIDDLEWARE PATTERN
//
// A middleware is a function that takes a Handler and returns a Handler:
//
//   func Middleware(next http.Handler) http.Handler {
//       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//           // Before the request
//           next.ServeHTTP(w, r)
//           // After the request
//       })
//   }
//
// MIDDLEWARE CHAINING
//
// Middleware composes naturally:
//   handler := Logging(Auth(RateLimit(actualHandler)))
//
// Request flows: Logging -> Auth -> RateLimit -> actualHandler
// Response flows back: actualHandler -> RateLimit -> Auth -> Logging
//
// THE RESPONSEWRITER CAPTURE PATTERN
//
// To log response status codes or measure response size, you need to
// capture what the handler writes. This requires wrapping ResponseWriter:
//
//   type responseCapture struct {
//       http.ResponseWriter
//       statusCode int
//   }
//
// =============================================================================

// LoggingMiddleware logs request method, path, and response time
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log after request completes
		duration := time.Since(start)
		fmt.Printf("   [LOG] %s %s - %v\n", r.Method, r.URL.Path, duration)
	})
}

// AuthMiddleware checks for an Authorization header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// In real code, validate the token here
		// Then add user info to context

		next.ServeHTTP(w, r)
	})
}

// RecoveryMiddleware catches panics and returns 500
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("   [PANIC] %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func DemonstrateMiddleware() {
	fmt.Println("=== MIDDLEWARE ===")
	fmt.Println()

	// The actual handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/panic" {
			panic("intentional panic!")
		}
		fmt.Fprintf(w, "Hello!")
	})

	// Stack middleware: Recovery -> Logging -> Auth -> Handler
	// (Recovery must be outermost to catch all panics)
	wrapped := RecoveryMiddleware(LoggingMiddleware(AuthMiddleware(handler)))

	server := httptest.NewServer(wrapped)
	defer server.Close()

	client := &http.Client{}

	// Request without auth - blocked by AuthMiddleware
	fmt.Println("1. Request without auth:")
	resp1, _ := client.Get(server.URL + "/test")
	body1, _ := io.ReadAll(resp1.Body)
	resp1.Body.Close()
	fmt.Printf("   Status: %d, Body: %s\n", resp1.StatusCode, strings.TrimSpace(string(body1)))
	fmt.Println()

	// Request with auth - succeeds
	fmt.Println("2. Request with auth:")
	req2, _ := http.NewRequest("GET", server.URL+"/test", nil)
	req2.Header.Set("Authorization", "Bearer token")
	resp2, _ := client.Do(req2)
	body2, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()
	fmt.Printf("   Status: %d, Body: %s\n", resp2.StatusCode, string(body2))
	fmt.Println()

	// Request that causes panic - handled by RecoveryMiddleware
	fmt.Println("3. Request that panics:")
	req3, _ := http.NewRequest("GET", server.URL+"/panic", nil)
	req3.Header.Set("Authorization", "Bearer token")
	resp3, _ := client.Do(req3)
	body3, _ := io.ReadAll(resp3.Body)
	resp3.Body.Close()
	fmt.Printf("   Status: %d, Body: %s\n", resp3.StatusCode, strings.TrimSpace(string(body3)))
	fmt.Println()
}

// =============================================================================
// SECTION 11: CONTEXT IN HTTP
// =============================================================================
//
// WHY CONTEXT MATTERS FOR HTTP
//
// HTTP requests can be cancelled:
//   - Client disconnects
//   - Timeout expires
//   - User navigates away
//   - Upstream service times out
//
// Without context, your handler keeps running after the client is gone,
// wasting resources on work no one will ever see.
//
// REQUEST CONTEXT
//
// Every *http.Request has a context: r.Context()
//
// This context is cancelled when:
//   - The client closes the connection
//   - The server's read timeout is exceeded
//   - You explicitly cancel it
//
// PASSING VALUES VIA CONTEXT
//
// Middleware often needs to pass data to handlers (user ID, request ID).
// Use context.WithValue():
//
//   ctx := context.WithValue(r.Context(), "userID", 123)
//   r = r.WithContext(ctx)
//
// Then in handlers:
//   userID := r.Context().Value("userID")
//
// BEST PRACTICE: Use typed keys to avoid collisions:
//   type contextKey string
//   const userIDKey contextKey = "userID"
//
// =============================================================================

type contextKey string

const requestIDKey contextKey = "requestID"

// RequestIDMiddleware adds a unique ID to each request
func RequestIDMiddleware(next http.Handler) http.Handler {
	counter := 0
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		counter++
		requestID := fmt.Sprintf("req-%d", counter)

		// Add to context
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		r = r.WithContext(ctx)

		// Add to response headers
		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r)
	})
}

func DemonstrateContext() {
	fmt.Println("=== CONTEXT IN HTTP ===")
	fmt.Println()

	mux := http.NewServeMux()

	// Handler that uses context values
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Context().Value(requestIDKey).(string)
		fmt.Fprintf(w, "Request ID: %s", requestID)
	})

	// Handler that respects cancellation
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Simulate slow work that checks for cancellation
		select {
		case <-time.After(100 * time.Millisecond):
			fmt.Fprintf(w, "Completed!")
		case <-ctx.Done():
			// Client disconnected or timeout
			fmt.Printf("   [CANCELLED] %v\n", ctx.Err())
			return
		}
	})

	// Wrap with request ID middleware
	handler := RequestIDMiddleware(mux)

	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{}

	// Test request ID
	fmt.Println("1. Request ID from context:")
	resp1, _ := client.Get(server.URL + "/")
	body1, _ := io.ReadAll(resp1.Body)
	resp1.Body.Close()
	fmt.Printf("   Body: %s\n", string(body1))
	fmt.Printf("   X-Request-ID header: %s\n", resp1.Header.Get("X-Request-ID"))
	fmt.Println()

	// Test with client timeout
	fmt.Println("2. Request with context timeout:")

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", server.URL+"/slow", nil)
	_, err := client.Do(req)
	if err != nil {
		fmt.Printf("   Client error (expected): %v\n", err)
	}
	fmt.Println()

	// Test without timeout - completes
	fmt.Println("3. Request without timeout:")
	resp3, _ := client.Get(server.URL + "/slow")
	body3, _ := io.ReadAll(resp3.Body)
	resp3.Body.Close()
	fmt.Printf("   Body: %s\n", string(body3))
	fmt.Println()
}

// =============================================================================
// SECTION 12: SERVER CONFIGURATION
// =============================================================================
//
// BEYOND LISTENANDSERVE
//
// http.ListenAndServe is a convenience function. For production, use
// http.Server directly for full control:
//
//   server := &http.Server{
//       Addr:         ":8080",
//       Handler:      mux,
//       ReadTimeout:  5 * time.Second,
//       WriteTimeout: 10 * time.Second,
//       IdleTimeout:  120 * time.Second,
//   }
//   server.ListenAndServe()
//
// TIMEOUT MEANINGS
//
// ReadTimeout: Max time to read entire request (headers + body)
//   Too short: Large uploads fail
//   Too long: Slow clients tie up connections
//
// WriteTimeout: Max time to write entire response
//   Too short: Large downloads fail
//   Too long: Slow clients tie up connections
//
// IdleTimeout: Max time between requests on keep-alive connections
//   Frees up connections from idle clients
//
// ReadHeaderTimeout: Max time to read just the headers
//   Protects against slowloris attacks
//
// GRACEFUL SHUTDOWN
//
// server.Shutdown(ctx) gracefully shuts down:
//   1. Stops accepting new connections
//   2. Waits for active requests to complete
//   3. Respects the context's deadline
//
// This is essential for production. Without it, you kill active requests
// when deploying, which causes errors for your users.
//
// =============================================================================

func DemonstrateServerConfig() {
	fmt.Println("=== SERVER CONFIGURATION ===")
	fmt.Println()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		fmt.Fprintf(w, "Done!")
	})

	// Create a properly configured server
	server := &http.Server{
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	// Use httptest to get a test server with our config
	testServer := httptest.NewUnstartedServer(mux)
	testServer.Config = server
	testServer.Start()
	defer testServer.Close()

	fmt.Println("Server configured with:")
	fmt.Println("  ReadTimeout:       5s")
	fmt.Println("  ReadHeaderTimeout: 2s")
	fmt.Println("  WriteTimeout:      10s")
	fmt.Println("  IdleTimeout:       120s")
	fmt.Println()

	// Demonstrate graceful shutdown
	fmt.Println("Graceful shutdown pattern:")
	fmt.Println(`
   // In production code:
   go func() {
       if err := server.ListenAndServe(); err != http.ErrServerClosed {
           log.Fatal(err)
       }
   }()

   // Wait for interrupt signal
   quit := make(chan os.Signal, 1)
   signal.Notify(quit, os.Interrupt)
   <-quit

   // Graceful shutdown with 30s timeout
   ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
   defer cancel()
   server.Shutdown(ctx)
   `)
	fmt.Println()
}

// =============================================================================
// SECTION 13: TESTING HTTP
// =============================================================================
//
// HTTPTEST PACKAGE
//
// net/http/httptest provides utilities for testing HTTP code:
//
// httptest.NewServer(handler):
//   Creates a real HTTP server on a random port.
//   Returns *httptest.Server with .URL for the base URL.
//   Don't forget defer server.Close()
//
// httptest.NewRecorder():
//   Captures what a handler writes without a real server.
//   Implements http.ResponseWriter.
//   Access results via .Code, .Body, .Header()
//
// httptest.NewRequest(method, target, body):
//   Creates an *http.Request for testing.
//   No network involved.
//
// WHEN TO USE EACH
//
// NewRecorder: Unit testing handlers in isolation
//   - Fast, no network
//   - Tests handler logic directly
//
// NewServer: Integration testing
//   - Tests actual HTTP behavior
//   - Tests middleware, routing, serialization
//   - Required for testing client code
//
// =============================================================================

func DemonstrateTesting() {
	fmt.Println("=== TESTING HTTP ===")
	fmt.Println()

	// The handler we want to test
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var data struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, " + data.Name,
		})
	})

	// Test 1: Using ResponseRecorder (unit test style)
	fmt.Println("1. Using httptest.NewRecorder (unit test):")

	// Create a request
	reqBody := strings.NewReader(`{"name": "Alice"}`)
	req := httptest.NewRequest(http.MethodPost, "/greet", reqBody)
	req.Header.Set("Content-Type", "application/json")

	// Create a recorder
	rr := httptest.NewRecorder()

	// Call the handler directly
	handler.ServeHTTP(rr, req)

	// Check the results
	fmt.Printf("   Status: %d (want 201)\n", rr.Code)
	fmt.Printf("   Content-Type: %s\n", rr.Header().Get("Content-Type"))
	fmt.Printf("   Body: %s\n", strings.TrimSpace(rr.Body.String()))
	fmt.Println()

	// Test 2: Wrong method
	fmt.Println("2. Testing error case (wrong method):")
	req2 := httptest.NewRequest(http.MethodGet, "/greet", nil)
	rr2 := httptest.NewRecorder()
	handler.ServeHTTP(rr2, req2)
	fmt.Printf("   Status: %d (want 405)\n", rr2.Code)
	fmt.Println()

	// Test 3: Using real server (integration test style)
	fmt.Println("3. Using httptest.NewServer (integration test):")
	server := httptest.NewServer(handler)
	defer server.Close()

	// Use real HTTP client
	resp, err := http.Post(
		server.URL+"/greet",
		"application/json",
		strings.NewReader(`{"name": "Bob"}`),
	)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("   Status: %d\n", resp.StatusCode)
	fmt.Printf("   Body: %s\n", strings.TrimSpace(string(body)))
	fmt.Println()
}

// =============================================================================
// SECTION 14: COMMON PATTERNS
// =============================================================================
//
// These patterns appear constantly in production Go HTTP code.
//
// =============================================================================

func DemonstrateCommonPatterns() {
	fmt.Println("=== COMMON PATTERNS ===")
	fmt.Println()

	// Pattern 1: REST-style resource handler
	fmt.Println("1. REST resource pattern:")
	fmt.Println(`
   mux.HandleFunc("GET /users", listUsers)
   mux.HandleFunc("POST /users", createUser)
   mux.HandleFunc("GET /users/{id}", getUser)
   mux.HandleFunc("PUT /users/{id}", updateUser)
   mux.HandleFunc("DELETE /users/{id}", deleteUser)
   `)

	// Pattern 2: Middleware chain
	fmt.Println("2. Middleware chain pattern:")
	fmt.Println(`
   handler := Recovery(
       Logging(
           RateLimit(
               Auth(
                   actualHandler,
               ),
           ),
       ),
   )
   `)

	// Pattern 3: Structured JSON responses
	fmt.Println("3. JSON response helper:")
	fmt.Println(`
   func respondJSON(w http.ResponseWriter, status int, data interface{}) {
       w.Header().Set("Content-Type", "application/json")
       w.WriteHeader(status)
       json.NewEncoder(w).Encode(data)
   }

   func respondError(w http.ResponseWriter, status int, message string) {
       respondJSON(w, status, map[string]string{"error": message})
   }
   `)

	// Pattern 4: Request validation
	fmt.Println("4. Request validation pattern:")
	fmt.Println(`
   func createUser(w http.ResponseWriter, r *http.Request) {
       var req CreateUserRequest
       if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
           respondError(w, http.StatusBadRequest, "Invalid JSON")
           return
       }
       if err := req.Validate(); err != nil {
           respondError(w, http.StatusBadRequest, err.Error())
           return
       }
       // ... process valid request
   }
   `)

	// Pattern 5: Client with retry
	fmt.Println("5. HTTP client with retry:")
	fmt.Println(`
   func doWithRetry(req *http.Request, maxRetries int) (*http.Response, error) {
       var resp *http.Response
       var err error

       for i := 0; i < maxRetries; i++ {
           resp, err = client.Do(req)
           if err == nil && resp.StatusCode < 500 {
               return resp, nil
           }
           if resp != nil {
               resp.Body.Close()
           }
           time.Sleep(time.Duration(i+1) * time.Second)
       }
       return resp, err
   }
   `)

	// Pattern 6: Health check endpoint
	fmt.Println("6. Health check pattern:")
	fmt.Println(`
   mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(map[string]string{
           "status": "healthy",
           "time":   time.Now().Format(time.RFC3339),
       })
   })
   `)
	fmt.Println()
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
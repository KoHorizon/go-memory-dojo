// Package http_practice is your daily practice space for the net/http package.
//
// INSTRUCTIONS:
// 1. Each morning, open this file fresh
// 2. Fill in all the TODOs from memory (no peeking at http.go!)
// 3. Run with: go run cmd/main.go --module http --practice
// 4. Check your answers against the theory file
// 5. Note what you missed - focus on those tomorrow
package http_practice

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// =============================================================================
// EXERCISE 1: BASIC HTTP CLIENT
// =============================================================================
// Fill in the code to make HTTP requests.

func PracticeBasicClient() {
	fmt.Println("=== PRACTICE: BASIC HTTP CLIENT ===")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Method: %s, Path: %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	// TODO: Make a GET request to server.URL
	// Hint: Use http.Get()
	// resp, err := ???

	// TODO: Check for errors
	// if err != nil { ... }

	// TODO: CRITICAL - Close the response body
	// Why is this critical? (Think about connection pooling)
	// defer ???

	// TODO: Read the response body
	// Hint: Use io.ReadAll()
	// body, err := ???

	// TODO: Print the status code and body
	// fmt.Printf("Status: %d\n", ???)
	// fmt.Printf("Body: %s\n", ???)

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 2: CUSTOM REQUESTS
// =============================================================================
// Create requests with custom methods, headers, and bodies.

func PracticeCustomRequests() {
	fmt.Println("=== PRACTICE: CUSTOM REQUESTS ===")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Method: %s, Auth: %s", r.Method, r.Header.Get("Authorization"))
	}))
	defer server.Close()

	// TODO: Create a new request with context
	// What function should you use instead of http.NewRequest?
	// Hint: It includes "Context" in the name
	//
	// ctx := context.Background()
	// req, err := ???(ctx, http.MethodPut, server.URL+"/users/123", nil)

	// TODO: Set headers on the request
	// req.Header.Set("Authorization", "Bearer token123")
	// req.Header.Set("Content-Type", "application/json")

	// TODO: Create a client with timeout
	// client := &http.Client{Timeout: ???}

	// TODO: Send the request
	// resp, err := client.Do(???)

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 3: CLIENT CONFIGURATION
// =============================================================================
// Configure a production-ready HTTP client.

func PracticeClientConfig() {
	fmt.Println("=== PRACTICE: CLIENT CONFIGURATION ===")

	// TODO: Create a properly configured client
	// Fill in the struct fields:
	//
	// client := &http.Client{
	//     Timeout: ???,  // Total request timeout
	//
	//     Transport: &http.Transport{
	//         MaxIdleConns:        ???,  // Total idle connections to keep
	//         MaxIdleConnsPerHost: ???,  // Idle connections per host
	//         IdleConnTimeout:     ???,  // How long idle connections live
	//     },
	// }

	// QUESTION: Why is it dangerous to use http.DefaultClient in production?
	// Answer: ???

	// QUESTION: What happens if you don't close response bodies?
	// Answer: ???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 4: STATUS CODES
// =============================================================================
// Understand what status codes mean and how to check them.

func PracticeStatusCodes() {
	fmt.Println("=== PRACTICE: STATUS CODES ===")

	// TODO: Match the status code to its meaning:
	// Write the correct category (1xx/2xx/3xx/4xx/5xx) and meaning

	// 200 = ??? (Category: ???, Meaning: ???)
	// 201 = ??? (Category: ???, Meaning: ???)
	// 204 = ??? (Category: ???, Meaning: ???)
	// 301 = ??? (Category: ???, Meaning: ???)
	// 400 = ??? (Category: ???, Meaning: ???)
	// 401 = ??? (Category: ???, Meaning: ???)
	// 403 = ??? (Category: ???, Meaning: ???)
	// 404 = ??? (Category: ???, Meaning: ???)
	// 429 = ??? (Category: ???, Meaning: ???)
	// 500 = ??? (Category: ???, Meaning: ???)
	// 502 = ??? (Category: ???, Meaning: ???)
	// 503 = ??? (Category: ???, Meaning: ???)

	// CRITICAL QUESTION:
	// Does http.Get() return an error for a 404 response?
	// Answer: ???
	// Why: ???

	fmt.Println("   (Fill in the answers above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 5: HANDLERS
// =============================================================================
// Implement HTTP handlers.

func PracticeHandlers() {
	fmt.Println("=== PRACTICE: HANDLERS ===")

	// TODO: What interface does an HTTP handler implement?
	// type Handler interface {
	//     ???
	// }

	// TODO: Implement a handler as a struct
	// type GreetHandler struct {
	//     greeting string
	// }
	//
	// func (h *GreetHandler) ???(w http.ResponseWriter, r *http.Request) {
	//     fmt.Fprintf(w, "%s!", h.greeting)
	// }

	// TODO: Write an inline handler function
	// handler := http.HandlerFunc(func(w ???, r ???) {
	//     fmt.Fprintf(w, "Hello!")
	// })

	// QUESTION: What is http.HandlerFunc?
	// Answer: ???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 6: ROUTING (MUX)
// =============================================================================
// Set up routes with the standard mux.

func PracticeRouting() {
	fmt.Println("=== PRACTICE: ROUTING ===")

	// TODO: Create a new mux (don't use DefaultServeMux!)
	// mux := ???

	// TODO: Register routes
	// Note the difference between "/api" and "/api/" (trailing slash)

	// Basic routes:
	// mux.HandleFunc("/", homeHandler)
	// mux.HandleFunc("/about", aboutHandler)

	// TODO: Go 1.22+ method-specific routes
	// How do you make a route that only accepts GET?
	// mux.HandleFunc("??? /users", listUsers)

	// TODO: Go 1.22+ path parameters
	// How do you capture {id} from the path?
	// mux.HandleFunc("GET /users/???", getUser)
	// Inside handler: id := r.PathValue("???")

	// TODO: Wildcard path (captures rest of path)
	// mux.HandleFunc("/files/???", serveFile)

	// QUESTION: Does "/api" match "/api/users"?
	// Answer: ???
	// QUESTION: Does "/api/" match "/api/users"?
	// Answer: ???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 7: WRITING RESPONSES
// =============================================================================
// Send different types of responses.

func PracticeResponses() {
	fmt.Println("=== PRACTICE: WRITING RESPONSES ===")

	// CRITICAL: What is the correct ORDER for writing responses?
	// 1. ???
	// 2. ???
	// 3. ???

	// TODO: Write a JSON response
	// func jsonHandler(w http.ResponseWriter, r *http.Request) {
	//     data := map[string]string{"message": "hello"}
	//     w.Header().Set("???", "application/json")
	//     ???.NewEncoder(w).Encode(data)
	// }

	// TODO: Write an error response
	// func errorHandler(w http.ResponseWriter, r *http.Request) {
	//     http.???(w, "Something went wrong", http.StatusInternalServerError)
	// }

	// TODO: Redirect to another URL
	// func redirectHandler(w http.ResponseWriter, r *http.Request) {
	//     http.???(w, r, "/new-path", http.StatusMovedPermanently)
	// }

	// TODO: Return a 201 Created with Location header
	// func createHandler(w http.ResponseWriter, r *http.Request) {
	//     w.Header().Set("???", "/users/123")
	//     w.WriteHeader(http.???)
	//     fmt.Fprintf(w, `{"id": 123}`)
	// }

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 8: READING REQUEST DATA
// =============================================================================
// Extract data from incoming requests.

func PracticeReadingRequests() {
	fmt.Println("=== PRACTICE: READING REQUEST DATA ===")

	// TODO: Get path parameters (Go 1.22+)
	// In handler for "/users/{id}":
	// id := r.???(???)

	// TODO: Get query parameters from /search?q=golang&limit=10
	// query := r.URL.???()
	// q := query.Get("???")
	// limit := query.Get("???")

	// TODO: Read JSON body
	// var data struct {
	//     Name string `json:"name"`
	// }
	// err := ???.NewDecoder(r.???).Decode(&data)

	// TODO: Read form data
	// First: r.???()  // Parse the form
	// Then: username := r.FormValue("???")

	// TODO: Get a header value
	// auth := r.???.Get("Authorization")

	// QUESTION: Can you read the request body twice?
	// Answer: ???
	// Why: ???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 9: MIDDLEWARE
// =============================================================================
// Write middleware functions.

func PracticeMiddleware() {
	fmt.Println("=== PRACTICE: MIDDLEWARE ===")

	// TODO: Write the middleware signature
	// func Middleware(next ???) ??? {
	//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//         // Before the request
	//         ???.ServeHTTP(w, r)
	//         // After the request
	//     })
	// }

	// TODO: Write a logging middleware
	// func LoggingMiddleware(next http.Handler) http.Handler {
	//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//         start := time.Now()
	//         ???
	//         duration := time.Since(start)
	//         log.Printf("%s %s - %v", r.Method, r.URL.Path, duration)
	//     })
	// }

	// TODO: Write an auth middleware that blocks unauthenticated requests
	// func AuthMiddleware(next http.Handler) http.Handler {
	//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//         auth := r.Header.Get("???")
	//         if auth == "" {
	//             http.Error(w, "Unauthorized", http.???)
	//             return  // Don't call next!
	//         }
	//         ???
	//     })
	// }

	// TODO: Chain middleware
	// Order: Recovery -> Logging -> Auth -> Handler
	// handler := ???(???(???(actualHandler)))

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 10: CONTEXT
// =============================================================================
// Use context for cancellation and values.

func PracticeContext() {
	fmt.Println("=== PRACTICE: CONTEXT ===")

	// TODO: Get the request context
	// ctx := r.???()

	// TODO: Check if the context is cancelled
	// select {
	// case <-time.After(time.Second):
	//     fmt.Fprintf(w, "Done!")
	// case <-ctx.???():
	//     // Client disconnected
	//     return
	// }

	// TODO: Add a value to the context (in middleware)
	// ctx := context.WithValue(r.Context(), "userID", 123)
	// r = r.???(ctx)

	// TODO: Read a value from the context (in handler)
	// userID := r.Context().???("userID")

	// QUESTION: Why should you use typed keys for context values?
	// Answer: ???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 11: SERVER CONFIGURATION
// =============================================================================
// Configure a production server.

func PracticeServerConfig() {
	fmt.Println("=== PRACTICE: SERVER CONFIGURATION ===")

	// TODO: Create a configured server
	// server := &http.Server{
	//     Addr:              ":8080",
	//     Handler:           mux,
	//     ???Timeout:         5 * time.Second,   // Time to read entire request
	//     ???Timeout:         10 * time.Second,  // Time to write response
	//     ???Timeout:         120 * time.Second, // Keep-alive timeout
	// }

	// TODO: What method starts the server?
	// server.???()

	// TODO: What method gracefully shuts down?
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	// server.???(ctx)

	// QUESTION: Why is graceful shutdown important?
	// Answer: ???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 12: TESTING HTTP
// =============================================================================
// Test handlers and servers.

func PracticeTesting() {
	fmt.Println("=== PRACTICE: TESTING HTTP ===")

	// TODO: Create a test request (no network)
	// req := httptest.???(http.MethodGet, "/path", nil)

	// TODO: Create a response recorder
	// rr := httptest.???()

	// TODO: Call the handler directly
	// handler.???(rr, req)

	// TODO: Check the response
	// status := rr.???
	// body := rr.???().String()
	// headers := rr.???()

	// TODO: Create a test server (with network)
	// server := httptest.???(handler)
	// defer server.???()
	// resp, err := http.Get(server.??? + "/path")

	// QUESTION: When do you use NewRecorder vs NewServer?
	// NewRecorder: ???
	// NewServer: ???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================
// Answer these from memory.

func SelfTest() {
	fmt.Println("=== SELF-TEST ===")
	fmt.Println("Answer from memory:")
	fmt.Println()

	// 1. What is the danger of using http.DefaultClient?
	_ = "???"

	// 2. What must you always do with response bodies?
	_ = "???"

	// 3. What method does http.Handler require?
	_ = "???"

	// 4. Does a 404 response cause an error from http.Get()?
	_ = "???"

	// 5. What's the order for writing a response?
	_ = "1. ??? 2. ??? 3. ???"

	// 6. How do you get path parameters in Go 1.22+?
	_ = "???"

	// 7. What's the middleware function signature?
	_ = "func(http.???) http.???"

	// 8. How do you get the request context?
	_ = "r.???()"

	// 9. What function creates a test recorder?
	_ = "httptest.???()"

	// 10. What are the three timeout types for http.Server?
	_ = "???, ???, ???"

	fmt.Println("Check your answers against http/http.go!")
	fmt.Println()
}

// =============================================================================
// MINI PROJECT: BUILD A SIMPLE API
// =============================================================================
// Combine everything to build a mini REST API.

func MiniProject() {
	fmt.Println("=== MINI PROJECT: BUILD A REST API ===")
	fmt.Println()
	fmt.Println("Build an API with these endpoints:")
	fmt.Println("  GET  /health        - Return {\"status\": \"ok\"}")
	fmt.Println("  GET  /users         - List users (return JSON array)")
	fmt.Println("  POST /users         - Create user (read JSON body)")
	fmt.Println("  GET  /users/{id}    - Get single user")
	fmt.Println()
	fmt.Println("Requirements:")
	fmt.Println("  - Create your own mux (not DefaultServeMux)")
	fmt.Println("  - Add logging middleware")
	fmt.Println("  - Return proper status codes (200, 201, 404)")
	fmt.Println("  - Set Content-Type: application/json")
	fmt.Println("  - Write tests using httptest")
	fmt.Println()

	// Scaffold for your solution:
	// mux := http.NewServeMux()
	//
	// mux.HandleFunc("GET /health", ...)
	// mux.HandleFunc("GET /users", ...)
	// mux.HandleFunc("POST /users", ...)
	// mux.HandleFunc("GET /users/{id}", ...)
	//
	// handler := LoggingMiddleware(mux)
	//
	// server := httptest.NewServer(handler)
	// defer server.Close()
	//
	// // Test your API here

	fmt.Println("   (Implement the API above)")
	fmt.Println()
}

// RunAllPractice executes all practice exercises.
func RunAllPractice() {
	PracticeBasicClient()
	PracticeCustomRequests()
	PracticeClientConfig()
	PracticeStatusCodes()
	PracticeHandlers()
	PracticeRouting()
	PracticeResponses()
	PracticeReadingRequests()
	PracticeMiddleware()
	PracticeContext()
	PracticeServerConfig()
	PracticeTesting()
	SelfTest()
	MiniProject()

	fmt.Println("=====================================")
	fmt.Println("Practice complete!")
	fmt.Println("Now check your answers against:")
	fmt.Println("  modules/http-module/http/http.go")
	fmt.Println("=====================================")
}

// Helper function to suppress unused import warnings
func init() {
	_ = io.ReadAll
	_ = httptest.NewServer
}
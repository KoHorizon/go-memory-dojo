// Package context provides comprehensive documentation and working examples
// for Go's context package - the foundation of cancellation and request-scoped data.
package context

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// =============================================================================
// SECTION 1: WHY CONTEXT EXISTS
// =============================================================================
//
// THE PROBLEM: RUNAWAY GOROUTINES
//
// Imagine a web server handling a request:
//   1. User makes request
//   2. Server spawns goroutines: database query, API call, file read
//   3. User closes browser (cancels request)
//   4. ...but those goroutines keep running, wasting resources
//
// Without context, there's no standard way to tell goroutines "stop working,
// nobody cares about your result anymore."
//
// THE SOLUTION: CONTEXT
//
// Context provides:
//   1. CANCELLATION - Signal goroutines to stop
//   2. DEADLINES - Automatic cancellation after a time
//   3. VALUES - Request-scoped data (user ID, trace ID, etc.)
//
// THE CONTEXT INTERFACE
//
//   type Context interface {
//       Done() <-chan struct{}           // Closed when cancelled
//       Err() error                      // Why it was cancelled
//       Deadline() (time.Time, bool)     // When it will be cancelled
//       Value(key any) any               // Request-scoped values
//   }
//
// THE RULES
//
// 1. Pass context as first parameter: func DoThing(ctx context.Context, ...)
// 2. Don't store context in structs
// 3. Don't pass nil context - use context.TODO() if unsure
// 4. Context values are for request-scoped data, not function parameters
//
// =============================================================================

// =============================================================================
// SECTION 2: CREATING CONTEXTS
// =============================================================================
//
// ROOT CONTEXTS
//
//   context.Background() - The root, never cancelled
//                          Use at: main, init, tests, top of request handler
//
//   context.TODO()       - Placeholder when unsure which context to use
//                          Use during refactoring, will be replaced later
//
// DERIVED CONTEXTS
//
//   context.WithCancel(parent)           - Cancel manually
//   context.WithTimeout(parent, dur)     - Cancel after duration
//   context.WithDeadline(parent, time)   - Cancel at specific time
//   context.WithValue(parent, key, val)  - Add a value
//
// CANCELLATION IS INHERITED
//
// If parent is cancelled, all children are cancelled too.
// But cancelling a child doesn't affect the parent or siblings.
//
// =============================================================================

func DemonstrateCreatingContexts() {
	// Root contexts
	bg := context.Background()
	todo := context.TODO()
	fmt.Printf("Background: %v\n", bg)
	fmt.Printf("TODO: %v\n", todo)

	// WithCancel - manual cancellation
	ctx, cancel := context.WithCancel(bg)
	fmt.Printf("Before cancel - Done closed: %v\n", isClosed(ctx.Done()))
	cancel()
	fmt.Printf("After cancel - Done closed: %v\n", isClosed(ctx.Done()))
	fmt.Printf("Err: %v\n", ctx.Err())

	// WithTimeout - auto-cancel after duration
	ctx2, cancel2 := context.WithTimeout(bg, 100*time.Millisecond)
	defer cancel2() // Always call cancel to release resources!
	fmt.Printf("Timeout context created, waiting...\n")
	<-ctx2.Done()
	fmt.Printf("Timeout err: %v\n", ctx2.Err())

	// WithDeadline - auto-cancel at specific time
	deadline := time.Now().Add(50 * time.Millisecond)
	ctx3, cancel3 := context.WithDeadline(bg, deadline)
	defer cancel3()
	dl, ok := ctx3.Deadline()
	fmt.Printf("Has deadline: %v, Deadline: %v\n", ok, dl.Format(time.RFC3339Nano))
}

func isClosed(ch <-chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

// =============================================================================
// SECTION 3: CANCELLATION PATTERN
// =============================================================================
//
// THE PATTERN
//
//   ctx, cancel := context.WithCancel(parent)
//   defer cancel()  // ALWAYS defer cancel!
//
//   go func() {
//       select {
//       case <-ctx.Done():
//           return  // Cancelled, clean up and exit
//       case result := <-work:
//           // Use result
//       }
//   }()
//
// WHY ALWAYS DEFER CANCEL?
//
// Even if context times out or parent is cancelled, calling cancel()
// releases resources associated with the context. It's always safe
// to call cancel multiple times.
//
// CHECKING FOR CANCELLATION
//
// Option 1: Select on ctx.Done()
//   select {
//   case <-ctx.Done():
//       return ctx.Err()
//   default:
//       // Continue working
//   }
//
// Option 2: Check ctx.Err() directly
//   if ctx.Err() != nil {
//       return ctx.Err()
//   }
//
// =============================================================================

func DemonstrateCancellation() {
	ctx, cancel := context.WithCancel(context.Background())

	// Start a worker
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker stopped after %d iterations: %v\n", i, ctx.Err())
				return
			default:
				// Simulate work
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	// Let it run for a bit, then cancel
	time.Sleep(50 * time.Millisecond)
	cancel()
	wg.Wait()

	// Cancelling multiple goroutines at once
	ctx2, cancel2 := context.WithCancel(context.Background())
	var wg2 sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			<-ctx2.Done()
			fmt.Printf("Goroutine %d cancelled\n", id)
		}(i)
	}

	time.Sleep(10 * time.Millisecond)
	cancel2() // One call cancels all
	wg2.Wait()
}

// =============================================================================
// SECTION 4: TIMEOUTS AND DEADLINES
// =============================================================================
//
// WithTimeout vs WithDeadline
//
//   WithTimeout(ctx, 5*time.Second)  - Cancel 5 seconds from now
//   WithDeadline(ctx, specificTime)  - Cancel at that exact time
//
// WithTimeout is just WithDeadline(ctx, time.Now().Add(duration))
//
// DEADLINE PROPAGATION
//
// If parent has a deadline, child cannot extend it:
//   parent: deadline in 10s
//   child with WithTimeout(parent, 20s): still 10s (inherits parent's)
//   child with WithTimeout(parent, 5s): 5s (shorter is allowed)
//
// USE CASES
//
// - HTTP request: "complete within 30 seconds"
// - Database query: "timeout after 5 seconds"
// - External API: "give up after 10 seconds"
//
// =============================================================================

func DemonstrateTimeouts() {
	// Basic timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	<-ctx.Done()
	fmt.Printf("Timeout after: %v\n", time.Since(start))
	fmt.Printf("Error: %v\n", ctx.Err())

	// Deadline - specific time
	deadline := time.Now().Add(50 * time.Millisecond)
	ctx2, cancel2 := context.WithDeadline(context.Background(), deadline)
	defer cancel2()

	dl, _ := ctx2.Deadline()
	fmt.Printf("Deadline set to: %v\n", dl.Format("15:04:05.000"))

	// Simulating work with timeout
	ctx3, cancel3 := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel3()

	result := doWorkWithTimeout(ctx3, 30*time.Millisecond) // Completes in time
	fmt.Printf("Fast work result: %v\n", result)

	ctx4, cancel4 := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel4()

	result2 := doWorkWithTimeout(ctx4, 100*time.Millisecond) // Too slow
	fmt.Printf("Slow work result: %v\n", result2)
}

func doWorkWithTimeout(ctx context.Context, workDuration time.Duration) string {
	select {
	case <-time.After(workDuration):
		return "completed"
	case <-ctx.Done():
		return fmt.Sprintf("cancelled: %v", ctx.Err())
	}
}

// =============================================================================
// SECTION 5: CONTEXT VALUES
// =============================================================================
//
// WHAT ARE CONTEXT VALUES?
//
// Key-value pairs attached to a context, inherited by children.
// Used for request-scoped data that crosses API boundaries.
//
// CREATING
//
//   ctx := context.WithValue(parent, key, value)
//
// RETRIEVING
//
//   value := ctx.Value(key)  // Returns nil if not found
//
// KEY BEST PRACTICES
//
// Use unexported type for keys to prevent collisions:
//
//   type contextKey string
//   const userIDKey contextKey = "userID"
//
// Or use a struct type:
//
//   type userIDKeyType struct{}
//   var userIDKey = userIDKeyType{}
//
// WHAT TO STORE (AND NOT STORE)
//
// Good:
//   - Request ID / Trace ID
//   - User ID / Auth info
//   - Locale / Language
//
// Bad:
//   - Function parameters (pass explicitly)
//   - Optional config (use functional options)
//   - Database connections (use dependency injection)
//
// =============================================================================

type contextKey string

const (
	userIDKey    contextKey = "userID"
	requestIDKey contextKey = "requestID"
)

func DemonstrateContextValues() {
	// Setting values
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIDKey, 12345)
	ctx = context.WithValue(ctx, requestIDKey, "req-abc-123")

	// Retrieving values
	userID := ctx.Value(userIDKey)
	requestID := ctx.Value(requestIDKey)
	missing := ctx.Value(contextKey("missing"))

	fmt.Printf("UserID: %v\n", userID)
	fmt.Printf("RequestID: %v\n", requestID)
	fmt.Printf("Missing key: %v\n", missing)

	// Type assertion pattern
	if uid, ok := ctx.Value(userIDKey).(int); ok {
		fmt.Printf("User ID (typed): %d\n", uid)
	}

	// Values are inherited
	child := context.WithValue(ctx, contextKey("extra"), "child-value")
	fmt.Printf("Child sees parent value: %v\n", child.Value(userIDKey))
	fmt.Printf("Child sees own value: %v\n", child.Value(contextKey("extra")))
	fmt.Printf("Parent doesn't see child value: %v\n", ctx.Value(contextKey("extra")))

	// Using in a function chain
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	userID := ctx.Value(userIDKey)
	requestID := ctx.Value(requestIDKey)
	fmt.Printf("Processing request %v for user %v\n", requestID, userID)
}

// =============================================================================
// SECTION 6: CONTEXT ERRORS
// =============================================================================
//
// TWO ERROR VALUES
//
//   context.Canceled          - Returned when cancel() is called
//   context.DeadlineExceeded  - Returned when timeout/deadline is reached
//
// These are sentinel errors - check with errors.Is():
//
//   if errors.Is(ctx.Err(), context.Canceled) { ... }
//   if errors.Is(ctx.Err(), context.DeadlineExceeded) { ... }
//
// WHEN IS ERR SET?
//
// ctx.Err() returns nil until the context is cancelled.
// After cancellation, it returns one of the two errors above.
//
// =============================================================================

func DemonstrateContextErrors() {
	// Canceled error
	ctx1, cancel1 := context.WithCancel(context.Background())
	fmt.Printf("Before cancel: err=%v\n", ctx1.Err())
	cancel1()
	fmt.Printf("After cancel: err=%v\n", ctx1.Err())
	fmt.Printf("Is Canceled: %v\n", errors.Is(ctx1.Err(), context.Canceled))

	// DeadlineExceeded error
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel2()
	<-ctx2.Done()
	fmt.Printf("After timeout: err=%v\n", ctx2.Err())
	fmt.Printf("Is DeadlineExceeded: %v\n", errors.Is(ctx2.Err(), context.DeadlineExceeded))

	// Distinguishing error types
	handleError := func(err error) {
		switch {
		case errors.Is(err, context.Canceled):
			fmt.Println("-> Request was cancelled by client")
		case errors.Is(err, context.DeadlineExceeded):
			fmt.Println("-> Request timed out")
		default:
			fmt.Printf("-> Other error: %v\n", err)
		}
	}

	handleError(ctx1.Err())
	handleError(ctx2.Err())
}

// =============================================================================
// SECTION 7: PROPAGATING CONTEXT
// =============================================================================
//
// THE RULE: FIRST PARAMETER
//
// Context should be the first parameter of any function that needs it:
//
//   func DoSomething(ctx context.Context, arg1 string, arg2 int) error
//
// NOT in a struct:
//
//   type Service struct {
//       ctx context.Context  // DON'T DO THIS
//   }
//
// WHY FIRST PARAMETER?
//
// - Convention makes it easy to spot
// - Clear that function supports cancellation
// - Easy to propagate through call chain
//
// =============================================================================

func DemonstratePropagation() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Context flows through the call chain
	err := handleRequest(ctx, "some data")
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	}
}

func handleRequest(ctx context.Context, data string) error {
	// Check if already cancelled
	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Call downstream services with same context
	if err := queryDatabase(ctx, data); err != nil {
		return fmt.Errorf("database: %w", err)
	}

	if err := callExternalAPI(ctx, data); err != nil {
		return fmt.Errorf("api: %w", err)
	}

	return nil
}

func queryDatabase(ctx context.Context, query string) error {
	select {
	case <-time.After(30 * time.Millisecond):
		fmt.Println("Database query completed")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func callExternalAPI(ctx context.Context, data string) error {
	select {
	case <-time.After(30 * time.Millisecond):
		fmt.Println("API call completed")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// =============================================================================
// SECTION 8: REAL-WORLD PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: HTTP handler with timeout
	_ = `
	func handler(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()  // Already has cancellation from client disconnect
		
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		
		result, err := doWork(ctx)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				http.Error(w, "timeout", http.StatusGatewayTimeout)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)
	}
	`

	// Pattern 2: Graceful shutdown
	_ = `
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
	`

	// Pattern 3: Database query with timeout
	_ = `
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	rows, err := db.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", id)
	`

	// Pattern 4: Parallel operations with shared cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := make(chan string, 2)
	errs := make(chan error, 2)

	go func() {
		select {
		case <-time.After(20 * time.Millisecond):
			results <- "service1 done"
		case <-ctx.Done():
			errs <- ctx.Err()
		}
	}()

	go func() {
		select {
		case <-time.After(30 * time.Millisecond):
			results <- "service2 done"
		case <-ctx.Done():
			errs <- ctx.Err()
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case r := <-results:
			fmt.Printf("Pattern 4: %s\n", r)
		case e := <-errs:
			fmt.Printf("Pattern 4 error: %v\n", e)
		}
	}

	// Pattern 5: Request ID middleware
	_ = `
	func RequestIDMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.New().String()
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	`
}

// =============================================================================
// SECTION 9: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Forgetting to call cancel()
	_ = `
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()  <- MISSING! Resource leak!
	`
	fmt.Println("Mistake 1: Always defer cancel()")

	// Mistake 2: Storing context in struct
	_ = `
	type Server struct {
		ctx context.Context  // DON'T DO THIS
	}
	`
	fmt.Println("Mistake 2: Don't store context in structs")

	// Mistake 3: Passing nil context
	_ = `
	doWork(nil, data)  // DON'T! Use context.TODO() if unsure
	`
	fmt.Println("Mistake 3: Never pass nil, use context.TODO()")

	// Mistake 4: Using context for function parameters
	_ = `
	ctx = context.WithValue(ctx, "config", config)  // DON'T!
	// Pass config as explicit parameter instead
	`
	fmt.Println("Mistake 4: Don't use values for function parameters")

	// Mistake 5: Not checking ctx.Err() in loops
	_ = `
	for _, item := range items {
		process(item)  // Should check ctx.Err() periodically!
	}
	`
	fmt.Println("Mistake 5: Check ctx.Err() in long loops")

	// Mistake 6: Using string keys
	_ = `
	ctx = context.WithValue(ctx, "userID", 123)  // Risk of collision!
	// Use typed key instead
	`
	fmt.Println("Mistake 6: Use typed keys, not strings")
}

// =============================================================================
// SECTION 10: context.AfterFunc (Go 1.21+)
// =============================================================================
//
// AfterFunc schedules a function to run after context is done.
// Returns a stop function to prevent the callback.
//
//   stop := context.AfterFunc(ctx, func() {
//       // Cleanup code
//   })
//   defer stop()  // Cancel the callback if not needed
//
// =============================================================================

func DemonstrateAfterFunc() {
	ctx, cancel := context.WithCancel(context.Background())

	// Schedule cleanup when context is done
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("AfterFunc: Context was cancelled, cleaning up")
	})
	_ = stop // Can call stop() to prevent the callback

	cancel()
	time.Sleep(10 * time.Millisecond) // Give AfterFunc time to run
}

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateCreatingContexts()
	DemonstrateCancellation()
	DemonstrateTimeouts()
	DemonstrateContextValues()
	DemonstrateContextErrors()
	DemonstratePropagation()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
	DemonstrateAfterFunc()
}
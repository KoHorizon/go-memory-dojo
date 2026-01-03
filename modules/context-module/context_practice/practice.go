// Package context_practice is your daily practice space for the context package.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against context/context.go
// 3. Note what you missed - focus on those tomorrow
package context_practice

// =============================================================================
// EXERCISE 1: WHY CONTEXT EXISTS
// =============================================================================

func PracticeWhyContext() {
	// TODO: What problem does context solve?
	// Answer: ???

	// TODO: What three things does context provide?
	// 1. ???
	// 2. ???
	// 3. ???

	// TODO: Write the Context interface
	// type Context interface {
	//     ???() <-chan struct{}      // Closed when cancelled
	//     ???() error                // Why it was cancelled
	//     ???() (time.Time, bool)    // When it will be cancelled
	//     ???(key any) any           // Request-scoped values
	// }
}

// =============================================================================
// EXERCISE 2: CREATING CONTEXTS
// =============================================================================

func PracticeCreatingContexts() {
	// TODO: Create a root context (never cancelled)
	// ctx := context.???()

	// TODO: Create a placeholder context (for refactoring)
	// ctx := context.???()

	// TODO: Create a cancellable context
	// ctx, cancel := context.???(parent)
	// defer ???()

	// TODO: Create a context with timeout
	// ctx, cancel := context.???(parent, 5*time.Second)
	// defer ???()

	// TODO: Create a context with deadline
	// deadline := time.Now().Add(5 * time.Second)
	// ctx, cancel := context.???(parent, deadline)
	// defer ???()

	// TODO: Create a context with a value
	// ctx := context.???(parent, key, value)

	// QUESTION: If parent is cancelled, what happens to children?
	// Answer: ???

	// QUESTION: If child is cancelled, what happens to parent?
	// Answer: ???
}

// =============================================================================
// EXERCISE 3: CANCELLATION PATTERN
// =============================================================================

func PracticeCancellation() {
	// TODO: Write the standard cancellation pattern
	// ctx, cancel := context.WithCancel(parent)
	// defer ???()
	//
	// go func() {
	//     select {
	//     case <-ctx.???():
	//         return  // Cancelled
	//     case result := <-work:
	//         // Use result
	//     }
	// }()

	// TODO: Why must you always defer cancel()?
	// Answer: ???

	// TODO: Two ways to check if context is cancelled:
	// Option 1: select on ???
	// Option 2: check ???
}

// =============================================================================
// EXERCISE 4: TIMEOUTS AND DEADLINES
// =============================================================================

func PracticeTimeouts() {
	// TODO: What's the difference between WithTimeout and WithDeadline?
	// WithTimeout: ???
	// WithDeadline: ???

	// TODO: If parent has 10s timeout, child requests 20s, what happens?
	// Answer: ???

	// TODO: If parent has 10s timeout, child requests 5s, what happens?
	// Answer: ???

	// TODO: Check if context has a deadline
	// deadline, ok := ctx.???()
	// if ok {
	//     // Has deadline
	// }
}

// =============================================================================
// EXERCISE 5: CONTEXT VALUES
// =============================================================================

func PracticeContextValues() {
	// TODO: Set a value on context
	// ctx := context.???(parent, key, value)

	// TODO: Get a value from context
	// value := ctx.???(key)

	// TODO: Why should you use typed keys?
	// Answer: ???

	// TODO: Write the typed key pattern
	// type ??? string
	// const userIDKey ??? = "userID"

	// TODO: Type assertion pattern for values
	// if uid, ok := ctx.Value(key).(???); ok {
	//     // Use uid
	// }

	// TODO: What SHOULD you store in context values?
	// Good: ???

	// TODO: What should you NOT store?
	// Bad: ???
}

// =============================================================================
// EXERCISE 6: CONTEXT ERRORS
// =============================================================================

func PracticeContextErrors() {
	// TODO: What are the two context error values?
	// 1. context.??? - when cancel() is called
	// 2. context.??? - when timeout/deadline reached

	// TODO: How do you check which error occurred?
	// if errors.Is(ctx.Err(), context.???) { ... }
	// if errors.Is(ctx.Err(), context.???) { ... }

	// TODO: When does ctx.Err() return nil?
	// Answer: ???
}

// =============================================================================
// EXERCISE 7: PROPAGATING CONTEXT
// =============================================================================

func PracticePropagation() {
	// TODO: Where should context be in the function signature?
	// Answer: ???

	// TODO: Should you store context in a struct?
	// Answer: ???

	// TODO: Write the correct function signature
	// func DoWork(??? context.Context, data string) error

	// TODO: What do you pass if you don't have a context yet?
	// Answer: context.???() or context.???()
}

// =============================================================================
// EXERCISE 8: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// doWork(ctx)
	// Answer: ???

	// Mistake 2: What's wrong?
	// type Server struct {
	//     ctx context.Context
	// }
	// Answer: ???

	// Mistake 3: What's wrong?
	// doWork(nil, data)
	// Answer: ???

	// Mistake 4: What's wrong?
	// ctx = context.WithValue(ctx, "userID", 123)
	// Answer: ???

	// Mistake 5: What's wrong?
	// for _, item := range items {
	//     process(item)
	// }
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// 1. Root context that's never cancelled?
	_ = "context.???()"

	// 2. Placeholder context?
	_ = "context.???()"

	// 3. Context with manual cancellation?
	_ = "context.???(parent)"

	// 4. Context with timeout?
	_ = "context.???(parent, duration)"

	// 5. Context with deadline?
	_ = "context.???(parent, time)"

	// 6. Context with value?
	_ = "context.???(parent, key, value)"

	// 7. Channel that closes on cancellation?
	_ = "ctx.???()"

	// 8. Error after cancellation?
	_ = "ctx.???()"

	// 9. Get value from context?
	_ = "ctx.???(key)"

	// 10. Two context errors?
	_ = "context.???, context.???"
}

// =============================================================================
// MINI PROJECT: REQUEST HANDLER WITH TIMEOUT
// =============================================================================

func MiniProject() {
	// Build a request handler that:
	//
	// 1. Creates a context with 5 second timeout
	// 2. Adds request ID to context
	// 3. Calls a "database" function
	// 4. Calls an "API" function
	// 5. Handles cancellation/timeout gracefully
	//
	// Scaffold:
	//
	// type contextKey string
	// const requestIDKey contextKey = "requestID"
	//
	// func handleRequest(requestID string, data string) error {
	//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//     defer cancel()
	//
	//     ctx = context.WithValue(ctx, requestIDKey, requestID)
	//
	//     if err := queryDB(ctx, data); err != nil {
	//         return fmt.Errorf("db: %w", err)
	//     }
	//
	//     if err := callAPI(ctx, data); err != nil {
	//         return fmt.Errorf("api: %w", err)
	//     }
	//
	//     return nil
	// }
	//
	// func queryDB(ctx context.Context, query string) error {
	//     select {
	//     case <-time.After(1 * time.Second):
	//         reqID := ctx.Value(requestIDKey)
	//         fmt.Printf("[%v] DB query done\n", reqID)
	//         return nil
	//     case <-ctx.Done():
	//         return ctx.Err()
	//     }
	// }
	//
	// func callAPI(ctx context.Context, data string) error {
	//     select {
	//     case <-time.After(1 * time.Second):
	//         reqID := ctx.Value(requestIDKey)
	//         fmt.Printf("[%v] API call done\n", reqID)
	//         return nil
	//     case <-ctx.Done():
	//         return ctx.Err()
	//     }
	// }
}

// RunAllPractice is not meant to be called - this file is for reading.
func RunAllPractice() {
	PracticeWhyContext()
	PracticeCreatingContexts()
	PracticeCancellation()
	PracticeTimeouts()
	PracticeContextValues()
	PracticeContextErrors()
	PracticePropagation()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
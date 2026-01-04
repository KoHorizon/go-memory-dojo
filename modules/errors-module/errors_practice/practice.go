// Package errors_practice is your daily practice space for Go error handling.
//
// INSTRUCTIONS:
// 1. Each morning, open this file fresh
// 2. Fill in all the TODOs from memory (no peeking at errors.go!)
// 3. Run with: go run cmd/main.go --module errors --practice
// 4. Check your answers against the theory file
// 5. Note what you missed - focus on those tomorrow
package errors_practice

import (
	"errors"
	"fmt"
)

// =============================================================================
// EXERCISE 1: THE ERROR INTERFACE
// =============================================================================
// Understand the fundamental error interface.

func PracticeErrorInterface() {
	fmt.Println("=== PRACTICE: ERROR INTERFACE ===")

	// TODO: Write the error interface definition
	// type error interface {
	//     ???
	// }

	// TODO: What is the zero value of error?
	// Answer: ???

	// TODO: How do you check if an error occurred?
	// if err ??? nil { ... }

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 2: CREATING ERRORS
// =============================================================================
// Practice creating simple errors.

func PracticeCreatingErrors() {
	fmt.Println("=== PRACTICE: CREATING ERRORS ===")

	// TODO: Create a simple error with a message
	// err1 := ???.New("something went wrong")
	_ = errors.New("placeholder")

	// TODO: Create a formatted error
	// name := "config.yaml"
	// err2 := ???.Errorf("failed to load %s", name)
	_ = fmt.Errorf("placeholder")

	// TODO: What function from the errors package creates a simple error?
	// Answer: errors.???

	// TODO: What function from fmt creates a formatted error?
	// Answer: fmt.???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 3: SENTINEL ERRORS
// =============================================================================
// Define and use sentinel errors.

func PracticeSentinelErrors() {
	fmt.Println("=== PRACTICE: SENTINEL ERRORS ===")

	// TODO: Define a sentinel error
	// var ErrNotFound = ???.???(???)

	// TODO: Name 3 sentinel errors from the standard library
	// 1. io.???
	// 2. sql.???
	// 3. os.???

	// TODO: Why should you NOT compare sentinel errors with ==?
	// Answer: ???

	// TODO: What function should you use instead of ==?
	// Answer: errors.???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 4: ERROR WRAPPING WITH %w
// =============================================================================
// Practice wrapping errors with context.

func PracticeWrapping() {
	fmt.Println("=== PRACTICE: ERROR WRAPPING ===")

	// TODO: Wrap an error with context
	// original := errors.New("connection refused")
	// wrapped := fmt.Errorf("database error: %?", original)
	// What character goes after %?

	// TODO: What's the difference between %w and %v?
	// %w: ???
	// %v: ???

	// TODO: When should you use %w?
	// Answer: ???

	// TODO: When should you NOT wrap an error?
	// Answer: ???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 5: errors.Is
// =============================================================================
// Practice checking error identity.

func PracticeErrorsIs() {
	fmt.Println("=== PRACTICE: errors.Is ===")

	// Setup
	var ErrTimeout = errors.New("timeout")
	wrapped := fmt.Errorf("operation failed: %w", ErrTimeout)

	// TODO: Check if wrapped contains ErrTimeout
	// if errors.???(???, ???) {
	//     fmt.Println("It's a timeout!")
	// }
	_ = wrapped

	// TODO: What does errors.Is do with wrapped errors?
	// Answer: ???

	// TODO: Can a custom error type control Is behavior?
	// Answer: ??? by implementing ???

	// TODO: Write the Is method signature for custom errors
	// func (e *MyError) ???(target ???) ??? {

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 6: errors.As
// =============================================================================
// Practice extracting error types.

func PracticeErrorsAs() {
	fmt.Println("=== PRACTICE: errors.As ===")

	// TODO: What's the difference between errors.Is and errors.As?
	// errors.Is: ???
	// errors.As: ???

	// TODO: Extract a *PathError from an error
	// var pathErr ???
	// if errors.???(err, ???) {
	//     fmt.Println(pathErr.Path)
	// }

	// TODO: CRITICAL: Why must the target be a pointer to a pointer?
	// var pathErr *os.PathError  // This is already a pointer
	// errors.As(err, &pathErr)   // We pass ??? of it
	// Answer: Because errors.As needs to ??? the target

	// TODO: Write the As method signature for custom errors
	// func (e *MyError) ???(target ???) ??? {

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 7: errors.Unwrap
// =============================================================================
// Practice unwrapping errors.

func PracticeUnwrap() {
	fmt.Println("=== PRACTICE: errors.Unwrap ===")

	// TODO: Unwrap an error
	// inner := errors.New("root cause")
	// outer := fmt.Errorf("context: %w", inner)
	// unwrapped := errors.???(outer)

	// TODO: What does Unwrap return if there's no wrapped error?
	// Answer: ???

	// TODO: When would you call Unwrap directly?
	// Answer: ???

	// TODO: Write the Unwrap method signature for custom errors
	// func (e *MyError) ???() ??? {

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 8: errors.Join (Go 1.20+)
// =============================================================================
// Practice combining multiple errors.

func PracticeErrorsJoin() {
	fmt.Println("=== PRACTICE: errors.Join ===")

	// TODO: Combine multiple errors
	// err1 := errors.New("error 1")
	// err2 := errors.New("error 2")
	// combined := errors.???(???, ???)

	// TODO: Does errors.Is work with joined errors?
	// Answer: ???

	// TODO: What happens if you join nil errors?
	// errors.Join(nil, nil) returns ???

	// TODO: Name 3 use cases for errors.Join
	// 1. ???
	// 2. ???
	// 3. ???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 9: CUSTOM ERROR TYPES
// =============================================================================
// Practice creating custom error types.

func PracticeCustomErrors() {
	fmt.Println("=== PRACTICE: CUSTOM ERROR TYPES ===")

	// TODO: Create a custom error type with Code and Message fields
	// type APIError struct {
	//     ???
	//     ???
	// }
	//
	// func (e *APIError) ???() ??? {
	//     return fmt.Sprintf("[%d] %s", e.Code, e.Message)
	// }

	// TODO: Add Unwrap method to wrap an underlying error
	// type WrapperError struct {
	//     Message string
	//     Err     error
	// }
	//
	// func (e *WrapperError) ???() ??? {
	//     ???
	// }

	// TODO: When should you create a custom error type?
	// 1. ???
	// 2. ???
	// 3. ???

	// TODO: Should you use pointer or value receivers?
	// Answer: ??? because ???

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 10: ERROR HANDLING PATTERNS
// =============================================================================
// Practice common error handling patterns.

func PracticePatterns() {
	fmt.Println("=== PRACTICE: ERROR HANDLING PATTERNS ===")

	// TODO: The three options when you receive an error
	// 1. ???: when you can recover
	// 2. ???: when adding context helps
	// 3. ???: when no context to add

	// TODO: Write the defer cleanup pattern with error handling
	// func processFile(path string) (err error) {
	//     f, err := os.Open(path)
	//     if err != nil {
	//         return err
	//     }
	//     defer func() {
	//         if cerr := f.Close(); cerr != nil {
	//             err = errors.???(err, cerr)
	//         }
	//     }()
	//     return nil
	// }

	// TODO: Write the error type switch pattern
	// var httpErr *HTTPError
	// var pathErr *os.PathError
	//
	// switch {
	// case errors.???(err, &httpErr):
	//     // handle HTTP error
	// case errors.???(err, &pathErr):
	//     // handle path error
	// }

	fmt.Println("   (Fill in the TODOs above)")
	fmt.Println()
}

// =============================================================================
// EXERCISE 11: COMMON MISTAKES
// =============================================================================
// Know what NOT to do.

func PracticeMistakes() {
	fmt.Println("=== PRACTICE: COMMON MISTAKES ===")

	// TODO: What's wrong with this?
	// if err == ErrNotFound { ... }
	// Answer: ???
	// Fix: ???

	// TODO: What's wrong with this?
	// return fmt.Errorf("failed: %v", err)
	// Answer: ???
	// Fix: ???

	// TODO: What's wrong with this?
	// if err.Error() == "not found" { ... }
	// Answer: ???
	// Fix: ???

	// TODO: What's wrong with this?
	// var pathErr os.PathError
	// errors.As(err, &pathErr)
	// Answer: ???
	// Fix: ???

	// TODO: What's wrong with this?
	// return fmt.Errorf("context: %w", err)  // when err might be nil
	// Answer: ???
	// Fix: ???

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

	// 1. What method does the error interface require?
	_ = "???"

	// 2. What function creates a simple error?
	_ = "errors.???"

	// 3. What verb wraps an error in fmt.Errorf?
	_ = "%?"

	// 4. What function checks if an error IS a specific error?
	_ = "errors.???"

	// 5. What function extracts an error TYPE?
	_ = "errors.???"

	// 6. What function combines multiple errors?
	_ = "errors.???"

	// 7. What method do custom errors implement to support unwrapping?
	_ = "???"

	// 8. What's the difference between %w and %v in fmt.Errorf?
	_ = "%w: ???, %v: ???"

	// 9. Why use a pointer receiver for custom error types?
	_ = "???"

	// 10. What does errors.Is return for a wrapped error?
	_ = "???"

	fmt.Println("Check your answers against errors/errors.go!")
	fmt.Println()
}

// =============================================================================
// MINI PROJECT: BUILD AN ERROR SYSTEM
// =============================================================================
// Combine everything to build a complete error handling system.

func MiniProject() {
	fmt.Println("=== MINI PROJECT: ERROR SYSTEM ===")
	fmt.Println()
	fmt.Println("Build an error system for a user service:")
	fmt.Println()
	fmt.Println("1. Define sentinel errors:")
	fmt.Println("   - ErrUserNotFound")
	fmt.Println("   - ErrInvalidEmail")
	fmt.Println("   - ErrDuplicateUser")
	fmt.Println()
	fmt.Println("2. Create a ValidationError type:")
	fmt.Println("   - Field string")
	fmt.Println("   - Message string")
	fmt.Println("   - Implement Error()")
	fmt.Println()
	fmt.Println("3. Create a ServiceError type:")
	fmt.Println("   - Code string (e.g., \"USER_001\")")
	fmt.Println("   - Message string")
	fmt.Println("   - Err error (underlying)")
	fmt.Println("   - Implement Error() and Unwrap()")
	fmt.Println()
	fmt.Println("4. Write a validateUser function:")
	fmt.Println("   - Check name (not empty)")
	fmt.Println("   - Check email (contains @)")
	fmt.Println("   - Return joined ValidationErrors")
	fmt.Println()
	fmt.Println("5. Write a createUser function:")
	fmt.Println("   - Call validateUser")
	fmt.Println("   - Wrap validation errors in ServiceError")
	fmt.Println("   - Check for duplicates (return wrapped ErrDuplicateUser)")
	fmt.Println()
	fmt.Println("6. Write error handling code that:")
	fmt.Println("   - Uses errors.Is to check for ErrDuplicateUser")
	fmt.Println("   - Uses errors.As to extract ServiceError")
	fmt.Println("   - Uses errors.As to extract ValidationError")
	fmt.Println()

	// Scaffold:
	// var (
	//     ErrUserNotFound = errors.New("user not found")
	//     ErrInvalidEmail = errors.New("invalid email")
	//     ErrDuplicateUser = errors.New("duplicate user")
	// )
	//
	// type ValidationError struct { ... }
	// type ServiceError struct { ... }
	//
	// func validateUser(name, email string) error { ... }
	// func createUser(name, email string) error { ... }

	fmt.Println("   (Implement the system above)")
	fmt.Println()
}

// AllPractice executes all practice exercises.
func AllPractice() {
	PracticeErrorInterface()
	PracticeCreatingErrors()
	PracticeSentinelErrors()
	PracticeWrapping()
	PracticeErrorsIs()
	PracticeErrorsAs()
	PracticeUnwrap()
	PracticeErrorsJoin()
	PracticeCustomErrors()
	PracticePatterns()
	PracticeMistakes()
	SelfTest()
	MiniProject()

	fmt.Println("=====================================")
	fmt.Println("Practice complete!")
	fmt.Println("Now check your answers against:")
	fmt.Println("  modules/errors-module/errors/errors.go")
	fmt.Println("=====================================")
}
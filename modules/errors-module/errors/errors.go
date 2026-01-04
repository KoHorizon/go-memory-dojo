// Package errors provides comprehensive documentation and working examples
// for Go's error handling - the foundation of robust Go programs.
package errors

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF ERRORS IN GO
// =============================================================================
//
// WHY ERRORS ARE VALUES, NOT EXCEPTIONS
//
// In many languages, errors are "thrown" and "caught" with try/catch.
// Go takes a different approach: errors are just values returned from
// functions. This is deliberate:
//
// 1. EXPLICIT HANDLING: You can't accidentally ignore an error.
// 2. CONTROL FLOW IS CLEAR: No hidden jumps. Code flows top-to-bottom.
// 3. ERRORS ARE DATA: Store, compare, wrap, inspect them.
//
// THE ERROR INTERFACE
//
//   type error interface {
//       Error() string
//   }
//
// Any type with an Error() string method IS an error.
// The zero value of error is nil (meaning "no error" / "success").
//
// =============================================================================

func DemonstrateErrorBasics() {
	// Creating simple errors
	err1 := errors.New("something went wrong")
	err2 := fmt.Errorf("failed to load %s", "config.yaml")

	fmt.Printf("errors.New: %v\n", err1)
	fmt.Printf("fmt.Errorf: %v\n", err2)

	// Basic error handling
	result, err := mightFail(false)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Success: %d\n", result)

	_, err = mightFail(true)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func mightFail(shouldFail bool) (int, error) {
	if shouldFail {
		return 0, errors.New("intentional failure")
	}
	return 42, nil
}

// =============================================================================
// SECTION 2: SENTINEL ERRORS
// =============================================================================
//
// WHAT ARE SENTINEL ERRORS?
//
// Package-level variables that represent specific error conditions.
// They allow callers to check for SPECIFIC error types.
//
// Standard library examples:
//   io.EOF           - End of input
//   sql.ErrNoRows    - Query returned no results
//   os.ErrNotExist   - File doesn't exist
//   os.ErrPermission - Permission denied
//
// NAMING CONVENTION: Err<Condition>
//   ErrNotFound, ErrInvalidInput, ErrTimeout
//
// IMPORTANT: Compare with errors.Is, NOT ==
//
// Wrapped errors are NOT equal with ==, but errors.Is understands wrapping:
//   wrappedErr := fmt.Errorf("context: %w", io.EOF)
//   wrappedErr == io.EOF          // false!
//   errors.Is(wrappedErr, io.EOF) // true!
//
// =============================================================================

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrUnauthorized = errors.New("unauthorized")
)

func DemonstrateSentinelErrors() {
	// Standard library sentinels
	fmt.Printf("io.EOF: %v\n", io.EOF)
	fmt.Printf("sql.ErrNoRows: %v\n", sql.ErrNoRows)
	fmt.Printf("os.ErrNotExist: %v\n", os.ErrNotExist)

	// Using sentinels
	err := findUser(999)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not found, showing 404 page")
	}

	// Why == fails with wrapped errors
	wrapped := fmt.Errorf("db query: %w", ErrNotFound)
	fmt.Printf("wrapped == ErrNotFound: %v\n", wrapped == ErrNotFound)
	fmt.Printf("errors.Is(wrapped, ErrNotFound): %v\n", errors.Is(wrapped, ErrNotFound))

	// Real pattern: reading until EOF
	data := []byte("hello")
	reader := &simpleReader{data: data}
	for {
		buf := make([]byte, 10)
		_, err := reader.Read(buf)
		if errors.Is(err, io.EOF) {
			fmt.Println("Finished reading")
			break
		}
		if err != nil {
			fmt.Printf("Read error: %v\n", err)
			break
		}
	}
}

type simpleReader struct {
	data []byte
	pos  int
}

func (r *simpleReader) Read(p []byte) (int, error) {
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}
	n := copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}

func findUser(id int) error {
	if id == 999 {
		return ErrNotFound
	}
	return nil
}

// =============================================================================
// SECTION 3: ERROR WRAPPING WITH %w
// =============================================================================
//
// WHAT IS ERROR WRAPPING?
//
// Adding context to an error while preserving the original.
// Creates an "error chain" that can be inspected later.
//
// THE %w VERB
//
//   original := errors.New("connection refused")
//   wrapped := fmt.Errorf("database: %w", original)
//
// %w vs %v:
//   %w: Creates chain - errors.Is/As can find original
//   %v: No chain - original is lost for checking purposes
//
// WHEN TO WRAP:
//   - Crossing function/package boundaries
//   - Adding context helps debugging
//
// WHEN NOT TO WRAP:
//   - Creating a completely new error condition
//   - Original error is an implementation detail
//
// =============================================================================

func DemonstrateErrorWrapping() {
	// Basic wrapping
	original := errors.New("connection refused")
	wrapped := fmt.Errorf("database error: %w", original)
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Wrapped: %v\n", wrapped)

	// Multi-level wrapping (error chain)
	level1 := errors.New("disk full")
	level2 := fmt.Errorf("write file: %w", level1)
	level3 := fmt.Errorf("save user: %w", level2)
	level4 := fmt.Errorf("registration: %w", level3)
	fmt.Printf("Full chain: %v\n", level4)

	// %w vs %v - critical difference!
	errW := fmt.Errorf("with %%w: %w", ErrNotFound)
	errV := fmt.Errorf("with %%v: %v", ErrNotFound)
	fmt.Printf("errors.Is with %%w: %v\n", errors.Is(errW, ErrNotFound)) // true
	fmt.Printf("errors.Is with %%v: %v\n", errors.Is(errV, ErrNotFound)) // false!

	// Real pattern: adding context at each layer
	err := saveUserProfile(123, "invalid-email")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		if errors.Is(err, ErrInvalidInput) {
			fmt.Println("Root cause: invalid input")
		}
	}
}

func saveUserProfile(userID int, email string) error {
	if err := validateEmail(email); err != nil {
		return fmt.Errorf("saveUserProfile(user=%d): %w", userID, err)
	}
	return nil
}

func validateEmail(email string) error {
	if email == "invalid-email" {
		return fmt.Errorf("bad email '%s': %w", email, ErrInvalidInput)
	}
	return nil
}

// =============================================================================
// SECTION 4: errors.Is - CHECKING ERROR IDENTITY
// =============================================================================
//
// WHAT DOES errors.Is DO?
//
// Reports whether any error in the chain matches a target.
// It "unwraps" the chain, checking each level.
//
// HOW IT WORKS:
// 1. Is err == target?
// 2. Does err implement Is(target) returning true?
// 3. If wrapped, unwrap and repeat
//
// CUSTOM Is METHOD:
//
//   func (e *MyError) Is(target error) bool {
//       t, ok := target.(*MyError)
//       return ok && e.Code == t.Code
//   }
//
// =============================================================================

func DemonstrateErrorsIs() {
	// Basic usage
	err := fmt.Errorf("operation failed: %w", ErrUnauthorized)
	if errors.Is(err, ErrUnauthorized) {
		fmt.Println("Unauthorized - redirect to login")
	}

	// Deep chain
	deepErr := fmt.Errorf("layer3: %w",
		fmt.Errorf("layer2: %w",
			fmt.Errorf("layer1: %w", ErrNotFound)))
	fmt.Printf("errors.Is finds deep error: %v\n", errors.Is(deepErr, ErrNotFound))

	// File system errors
	_, err = os.Open("/nonexistent/path")
	fmt.Printf("os.ErrNotExist: %v\n", errors.Is(err, os.ErrNotExist))
	fmt.Printf("fs.ErrNotExist: %v\n", errors.Is(err, fs.ErrNotExist)) // same!

	// Custom Is method
	err1 := &HTTPError{Code: 404, Message: "page not found"}
	err2 := &HTTPError{Code: 404, Message: "user not found"}
	fmt.Printf("Same code matches: %v\n", errors.Is(err1, err2)) // true
}

type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}

func (e *HTTPError) Is(target error) bool {
	t, ok := target.(*HTTPError)
	return ok && e.Code == t.Code
}

// =============================================================================
// SECTION 5: errors.As - EXTRACTING ERROR TYPES
// =============================================================================
//
// WHAT DOES errors.As DO?
//
// Finds the first error in the chain matching a target TYPE
// and sets the target to that value.
//
// errors.Is: "Is this (or caused by) THIS specific error?"
// errors.As: "Is this (or caused by) THIS TYPE of error?"
//
// CRITICAL: TARGET MUST BE A POINTER
//
//   var pathErr *os.PathError
//   if errors.As(err, &pathErr) {
//       fmt.Println(pathErr.Path) // Access the field!
//   }
//
// =============================================================================

func DemonstrateErrorsAs() {
	// Extract *os.PathError
	_, err := os.Open("/nonexistent/path/file.txt")

	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		fmt.Printf("PathError - Op: %s, Path: %s\n", pathErr.Op, pathErr.Path)
	}

	// errors.As finds wrapped errors too
	wrapped := fmt.Errorf("config load failed: %w", err)
	var pathErr2 *os.PathError
	if errors.As(wrapped, &pathErr2) {
		fmt.Printf("Found through wrapping: %s\n", pathErr2.Path)
	}

	// Custom error type
	apiErr := callAPI()
	var httpErr *HTTPError
	if errors.As(apiErr, &httpErr) {
		switch httpErr.Code {
		case 400:
			fmt.Println("Bad request - check input")
		case 404:
			fmt.Println("Not found - resource missing")
		case 500:
			fmt.Println("Server error - retry later")
		}
	}

	// Type switching pattern
	err2 := processFile()
	var synErr *SyntaxError
	var ioErr *IOError
	switch {
	case errors.As(err2, &synErr):
		fmt.Printf("Syntax error at line %d\n", synErr.Line)
	case errors.As(err2, &ioErr):
		fmt.Printf("IO error on %s\n", ioErr.Path)
	}
}

func callAPI() error {
	return fmt.Errorf("API request: %w", &HTTPError{Code: 404, Message: "user not found"})
}

type SyntaxError struct {
	Line    int
	Message string
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("line %d: %s", e.Line, e.Message)
}

type IOError struct {
	Path    string
	Message string
}

func (e *IOError) Error() string {
	return fmt.Sprintf("%s: %s", e.Path, e.Message)
}

func processFile() error {
	return &SyntaxError{Line: 42, Message: "unexpected token"}
}

// =============================================================================
// SECTION 6: errors.Unwrap
// =============================================================================
//
// WHAT IS UNWRAPPING?
//
// Extracting the underlying error from a wrapped error.
// Like peeling an onion - each layer reveals the next.
//
// errors.Unwrap(err) returns:
//   - The wrapped error if err implements Unwrap() error
//   - nil if there's no wrapped error
//
// Usually you DON'T call Unwrap directly - errors.Is/As do it for you.
// Direct unwrapping is for debugging or building utilities.
//
// =============================================================================

func DemonstrateUnwrap() {
	inner := errors.New("connection refused")
	outer := fmt.Errorf("database: %w", inner)

	unwrapped := errors.Unwrap(outer)
	fmt.Printf("Outer: %v\n", outer)
	fmt.Printf("Unwrapped: %v\n", unwrapped)
	fmt.Printf("Same as inner: %v\n", unwrapped == inner)

	// Walking the error chain
	chain := fmt.Errorf("l3: %w", fmt.Errorf("l2: %w", fmt.Errorf("l1: %w", errors.New("root"))))
	for err := chain; err != nil; err = errors.Unwrap(err) {
		fmt.Printf("-> %v\n", err)
	}

	// Custom Unwrap
	dbErr := &DatabaseError{
		Operation: "INSERT",
		Table:     "users",
		Err:       errors.New("duplicate key"),
	}
	fmt.Printf("DatabaseError: %v\n", dbErr)
	fmt.Printf("Underlying: %v\n", errors.Unwrap(dbErr))
}

type DatabaseError struct {
	Operation string
	Table     string
	Err       error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("%s on %s: %v", e.Operation, e.Table, e.Err)
}

func (e *DatabaseError) Unwrap() error {
	return e.Err
}

// =============================================================================
// SECTION 7: errors.Join (Go 1.20+)
// =============================================================================
//
// WHAT IS errors.Join?
//
// Combines multiple errors into one. The result:
//   - err.Error() shows all messages
//   - errors.Is works for any wrapped error
//   - errors.As finds any matching type
//
// USE CASES:
//   - Cleanup: closing multiple resources
//   - Validation: reporting all failures
//   - Parallel ops: collecting goroutine errors
//   - Batch processing: some items failed
//
// =============================================================================

func DemonstrateErrorsJoin() {
	// Basic join
	err1 := errors.New("failed to close file")
	err2 := errors.New("failed to flush buffer")
	combined := errors.Join(err1, err2)
	fmt.Printf("Combined:\n%v\n", combined)

	// errors.Is works with joined
	fmt.Printf("Is err1: %v\n", errors.Is(combined, err1))
	fmt.Printf("Is err2: %v\n", errors.Is(combined, err2))

	// Real pattern: cleanup errors
	err := closeResources()
	if err != nil {
		fmt.Printf("Cleanup errors:\n%v\n", err)
	}

	// Real pattern: validation
	valErr := validateUser(User{Name: "", Email: "bad", Age: -5})
	if valErr != nil {
		fmt.Printf("Validation:\n%v\n", valErr)
	}

	// Nil handling
	result := errors.Join(nil, errors.New("only error"), nil)
	fmt.Printf("With nils: %v\n", result)

	allNil := errors.Join(nil, nil)
	fmt.Printf("All nil: %v\n", allNil)
}

func closeResources() error {
	var errs []error
	if err := closeDB(); err != nil {
		errs = append(errs, fmt.Errorf("db: %w", err))
	}
	if err := closeCache(); err != nil {
		errs = append(errs, fmt.Errorf("cache: %w", err))
	}
	return errors.Join(errs...)
}

func closeDB() error    { return errors.New("timeout") }
func closeCache() error { return nil }

type User struct {
	Name, Email string
	Age         int
}

func validateUser(u User) error {
	var errs []error
	if u.Name == "" {
		errs = append(errs, errors.New("name required"))
	}
	if u.Email == "bad" {
		errs = append(errs, errors.New("invalid email"))
	}
	if u.Age < 0 {
		errs = append(errs, fmt.Errorf("invalid age: %d", u.Age))
	}
	return errors.Join(errs...)
}

// =============================================================================
// SECTION 8: CUSTOM ERROR TYPES
// =============================================================================
//
// WHEN TO CREATE CUSTOM ERROR TYPES
//
// 1. Need additional data (codes, IDs, paths)
// 2. Need to distinguish error categories
// 3. Need custom Is() or As() behavior
// 4. Want structured error info for logging
//
// IMPLEMENT:
//   Error() string          - Required
//   Unwrap() error          - For wrapping
//   Is(target error) bool   - Custom equality
//
// Use pointer receivers for consistency.
//
// =============================================================================

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: %s - %s", e.Field, e.Message)
}

type RetryableError struct {
	Err        error
	RetryAfter int
}

func (e *RetryableError) Error() string {
	return fmt.Sprintf("%v (retry after %ds)", e.Err, e.RetryAfter)
}

func (e *RetryableError) Unwrap() error {
	return e.Err
}

func (e *RetryableError) Is(target error) bool {
	_, ok := target.(*RetryableError)
	return ok
}

type ErrorCode struct {
	Code    string
	Message string
	Details map[string]interface{}
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func DemonstrateCustomErrors() {
	// ValidationError
	valErr := &ValidationError{Field: "email", Message: "invalid format"}
	fmt.Printf("%v\n", valErr)

	// RetryableError with custom Is
	retryErr := &RetryableError{
		Err:        errors.New("service unavailable"),
		RetryAfter: 30,
	}
	wrapped := fmt.Errorf("API: %w", retryErr)

	var r *RetryableError
	if errors.As(wrapped, &r) {
		fmt.Printf("Retryable! Wait %d seconds\n", r.RetryAfter)
	}

	// ErrorCode for APIs
	apiErr := &ErrorCode{
		Code:    "USER_NOT_FOUND",
		Message: "user does not exist",
		Details: map[string]interface{}{"user_id": 123},
	}
	fmt.Printf("%v (details: %v)\n", apiErr, apiErr.Details)
}

// =============================================================================
// SECTION 9: ERROR HANDLING PATTERNS
// =============================================================================
//
// THE THREE OPTIONS
//
// 1. HANDLE: When you can recover
// 2. WRAP: When adding context helps
// 3. RETURN: When no context to add
//
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Handle, wrap, or return
	_ = `
	// HANDLE
	data, err := loadConfig()
	if err != nil {
		log.Printf("using defaults: %v", err)
		data = defaultConfig
	}

	// WRAP
	if err != nil {
		return fmt.Errorf("loadConfig: %w", err)
	}

	// RETURN
	if err != nil {
		return err
	}
	`

	// Pattern 2: Defer cleanup with error joining
	_ = `
	func process(path string) (err error) {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func() {
			if cerr := f.Close(); cerr != nil {
				err = errors.Join(err, cerr)
			}
		}()
		// ...
	}
	`

	// Pattern 3: Type switch
	err := simulateError()
	var httpErr *HTTPError
	var pathErr *os.PathError
	switch {
	case errors.As(err, &httpErr):
		fmt.Printf("HTTP %d: handle accordingly\n", httpErr.Code)
	case errors.As(err, &pathErr):
		fmt.Printf("Path %s: handle accordingly\n", pathErr.Path)
	case errors.Is(err, ErrNotFound):
		fmt.Println("Not found: show 404")
	default:
		fmt.Printf("Unknown: %v\n", err)
	}

	// Pattern 4: Error accumulation
	errs := processItems([]string{"good", "bad1", "good", "bad2"})
	if errs != nil {
		fmt.Printf("Batch errors:\n%v\n", errs)
	}
}

func simulateError() error {
	return &HTTPError{Code: 503, Message: "unavailable"}
}

func processItems(items []string) error {
	var errs []error
	for i, item := range items {
		if err := processItem(item); err != nil {
			errs = append(errs, fmt.Errorf("item %d: %w", i, err))
		}
	}
	return errors.Join(errs...)
}

func processItem(item string) error {
	if item == "bad1" || item == "bad2" {
		return errors.New("invalid")
	}
	return nil
}

// =============================================================================
// SECTION 10: COMMON MISTAKES
// =============================================================================
//
// 1. Using == instead of errors.Is
// 2. Using %v instead of %w
// 3. Comparing error strings
// 4. Wrong pointer type with errors.As
// 5. Wrapping nil errors
//
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: == vs errors.Is
	wrapped := fmt.Errorf("wrapped: %w", ErrNotFound)
	fmt.Printf("== (wrong): %v\n", wrapped == ErrNotFound)
	fmt.Printf("errors.Is (right): %v\n", errors.Is(wrapped, ErrNotFound))

	// Mistake 2: %v vs %w
	errV := fmt.Errorf("context: %v", ErrNotFound)
	errW := fmt.Errorf("context: %w", ErrNotFound)
	fmt.Printf("%%v breaks chain: %v\n", errors.Is(errV, ErrNotFound))
	fmt.Printf("%%w preserves chain: %v\n", errors.Is(errW, ErrNotFound))

	// Mistake 3: Comparing strings (don't!)
	// if err.Error() == "not found" { } // FRAGILE!

	// Mistake 4: Wrong pointer type
	// var pathErr os.PathError          // WRONG: not a pointer
	// var pathErr *os.PathError         // RIGHT: pointer type
	// errors.As(err, &pathErr)          // Pass address of pointer

	// Mistake 5: Wrapping nil
	wrappedNil := fmt.Errorf("context: %w", error(nil))
	fmt.Printf("Wrapped nil is NOT nil: %v\n", wrappedNil != nil)
}

// =============================================================================
// SECTION 11: PARSING AND CONVERTING ERRORS
// =============================================================================

func DemonstrateConversions() {
	// Error to string
	err := &HTTPError{Code: 500, Message: "crash"}
	str := err.Error()
	fmt.Printf("Error string: %s\n", str)

	// Parsing errors from stdlib
	_, err2 := strconv.Atoi("not-a-number")
	var numErr *strconv.NumError
	if errors.As(err2, &numErr) {
		fmt.Printf("NumError: Func=%s Num=%s Err=%v\n",
			numErr.Func, numErr.Num, numErr.Err)
	}
}

// AllDemonstrations executes all examples.
func AllDemonstrations() {
	DemonstrateErrorBasics()
	DemonstrateSentinelErrors()
	DemonstrateErrorWrapping()
	DemonstrateErrorsIs()
	DemonstrateErrorsAs()
	DemonstrateUnwrap()
	DemonstrateErrorsJoin()
	DemonstrateCustomErrors()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
	DemonstrateConversions()
}
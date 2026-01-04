// Package testingmod provides comprehensive documentation and working examples
// for Go's testing package - unit tests, benchmarks, examples, and fuzzing.
package testingmod

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF GO TESTING
// =============================================================================
//
// Go's testing philosophy is refreshingly simple:
//
// 1. No external framework needed - testing is built in
// 2. Tests are just Go code in _test.go files
// 3. Run with: go test
//
// TEST FILE NAMING
//
//   foo.go      → foo_test.go
//   math.go     → math_test.go
//
// TEST FUNCTION NAMING
//
//   func TestXxx(t *testing.T)     - Unit tests (Xxx is exported name)
//   func BenchmarkXxx(b *testing.B) - Benchmarks
//   func ExampleXxx()               - Testable examples
//   func FuzzXxx(f *testing.F)      - Fuzz tests (Go 1.18+)
//
// RUNNING TESTS
//
//   go test              - Run all tests in current package
//   go test ./...        - Run all tests in all packages
//   go test -v           - Verbose output
//   go test -run Regex   - Run matching tests
//   go test -count=1     - Disable test caching
//
// =============================================================================

// =============================================================================
// SECTION 2: BASIC TESTS
// =============================================================================
//
// TEST FUNCTION SIGNATURE
//
//   func TestXxx(t *testing.T)
//
// - Must start with Test
// - Xxx must start with capital letter
// - Takes *testing.T parameter
//
// REPORTING FAILURES
//
//   t.Error(args...)     - Log and continue
//   t.Errorf(format, args...) - Formatted, continue
//   t.Fatal(args...)     - Log and stop this test
//   t.Fatalf(format, args...) - Formatted, stop this test
//   t.Fail()             - Mark failed, continue
//   t.FailNow()          - Mark failed, stop
//
// LOGGING
//
//   t.Log(args...)       - Log (shown with -v or on failure)
//   t.Logf(format, args...) - Formatted log
//
// =============================================================================

// Add is a simple function to test
func Add(a, b int) int {
	return a + b
}

// Divide returns a/b or error if b is zero
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}

func TestAddNegative(t *testing.T) {
	result := Add(-1, -2)
	if result != -3 {
		t.Errorf("Add(-1, -2) = %d; want -3", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("expected error for division by zero")
	}
}

// =============================================================================
// SECTION 3: TABLE-DRIVEN TESTS
// =============================================================================
//
// THE PATTERN
//
// Table-driven tests are the idiomatic way to test multiple cases:
//
//   func TestXxx(t *testing.T) {
//       tests := []struct {
//           name     string
//           input    InputType
//           expected OutputType
//       }{
//           {"case1", input1, expected1},
//           {"case2", input2, expected2},
//       }
//
//       for _, tt := range tests {
//           t.Run(tt.name, func(t *testing.T) {
//               result := Function(tt.input)
//               if result != tt.expected {
//                   t.Errorf("got %v; want %v", result, tt.expected)
//               }
//           })
//       }
//   }
//
// WHY TABLE-DRIVEN?
//
// - Easy to add new test cases
// - Clear structure
// - t.Run creates subtests (can run individually)
// - Parallel execution possible
//
// =============================================================================

func TestAddTable(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 2, 3, 5},
		{"negative", -1, -2, -3},
		{"mixed", -1, 5, 4},
		{"zeros", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivideTable(t *testing.T) {
	tests := []struct {
		name      string
		a, b      int
		expected  int
		wantErr   bool
	}{
		{"simple", 10, 2, 5, false},
		{"integer division", 7, 2, 3, false},
		{"by zero", 10, 0, 0, true},
		{"negative", -10, 2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v; wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// =============================================================================
// SECTION 4: SUBTESTS AND t.Run
// =============================================================================
//
// t.Run(name, func(t *testing.T))
//
// Creates a subtest that:
// - Can be run individually: go test -run TestParent/SubtestName
// - Has its own t for failures
// - Supports parallel execution
//
// RUNNING SPECIFIC SUBTESTS
//
//   go test -run TestAdd/positive
//   go test -run TestAdd/negative
//   go test -run "TestAdd/.*"
//
// =============================================================================

func TestSubtests(t *testing.T) {
	t.Run("group1", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			// Test case 1
			if Add(1, 1) != 2 {
				t.Error("1+1 should be 2")
			}
		})
		t.Run("case2", func(t *testing.T) {
			// Test case 2
			if Add(2, 2) != 4 {
				t.Error("2+2 should be 4")
			}
		})
	})

	t.Run("group2", func(t *testing.T) {
		t.Log("Running group2 tests")
	})
}

// =============================================================================
// SECTION 5: PARALLEL TESTS
// =============================================================================
//
// t.Parallel() marks a test to run in parallel with other parallel tests.
//
// RULES
//
// 1. Call t.Parallel() at the start of the test
// 2. In table tests, capture loop variable (or use Go 1.22+ loop semantics)
// 3. Don't share mutable state between parallel tests
//
// GOTCHA: LOOP VARIABLE CAPTURE
//
//   // WRONG (before Go 1.22)
//   for _, tt := range tests {
//       t.Run(tt.name, func(t *testing.T) {
//           t.Parallel()
//           use(tt)  // tt may have changed!
//       })
//   }
//
//   // RIGHT
//   for _, tt := range tests {
//       tt := tt  // Capture
//       t.Run(tt.name, func(t *testing.T) {
//           t.Parallel()
//           use(tt)
//       })
//   }
//
// =============================================================================

func TestParallel(t *testing.T) {
	tests := []struct {
		name string
		val  int
	}{
		{"test1", 1},
		{"test2", 2},
		{"test3", 3},
	}

	for _, tt := range tests {
		tt := tt // Capture for Go < 1.22
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Simulate work
			time.Sleep(10 * time.Millisecond)
			t.Logf("Running %s with value %d", tt.name, tt.val)
		})
	}
}

// =============================================================================
// SECTION 6: TEST HELPERS
// =============================================================================
//
// t.Helper() marks a function as a test helper.
// When the helper reports a failure, the line number points to the
// calling test, not the helper itself.
//
// COMMON HELPERS
//
// - Assertion functions
// - Setup/teardown
// - Test data creation
//
// =============================================================================

func assertEqual(t *testing.T, got, want int) {
	t.Helper() // Makes error point to caller
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestWithHelpers(t *testing.T) {
	result := Add(2, 3)
	assertEqual(t, result, 5) // If fails, points here, not inside assertEqual

	_, err := Divide(10, 2)
	assertNoError(t, err)

	_, err = Divide(10, 0)
	assertError(t, err)
}

// =============================================================================
// SECTION 7: SETUP AND TEARDOWN
// =============================================================================
//
// t.Cleanup(func()) registers a function to run after the test completes.
// Multiple cleanups run in LIFO order (like defer).
//
// TestMain(m *testing.M) runs before any tests, for package-level setup.
//
// =============================================================================

func TestWithCleanup(t *testing.T) {
	// Setup
	resource := "test resource"
	t.Logf("Setup: created %s", resource)

	// Register cleanup
	t.Cleanup(func() {
		t.Logf("Cleanup: releasing %s", resource)
	})

	// Test
	if resource == "" {
		t.Error("resource should not be empty")
	}
	// Cleanup runs automatically after test
}

func setupTestDB(t *testing.T) *strings.Builder {
	t.Helper()
	db := &strings.Builder{}
	db.WriteString("test_db_connection")

	t.Cleanup(func() {
		db.Reset()
		t.Log("DB connection closed")
	})

	return db
}

func TestWithSetupHelper(t *testing.T) {
	db := setupTestDB(t)
	if db.Len() == 0 {
		t.Error("DB should be initialized")
	}
}

// TestMain is package-level setup/teardown
// Uncomment to use:
//
// func TestMain(m *testing.M) {
//     // Setup
//     fmt.Println("Setting up tests...")
//
//     // Run tests
//     code := m.Run()
//
//     // Teardown
//     fmt.Println("Cleaning up...")
//
//     os.Exit(code)
// }

// =============================================================================
// SECTION 8: SKIPPING TESTS
// =============================================================================
//
// t.Skip(reason)      - Skip this test
// t.Skipf(format, args...) - Skip with formatted message
// t.SkipNow()         - Skip immediately
//
// testing.Short()     - True if -short flag passed
//
// COMMON SKIP CONDITIONS
//
// - Integration tests in short mode
// - Platform-specific tests
// - Tests requiring external resources
//
// =============================================================================

func TestSkipExample(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	// Long-running test here
	time.Sleep(10 * time.Millisecond)
}

func TestSkipPlatform(t *testing.T) {
	// Example: skip if not on target platform
	// if runtime.GOOS != "linux" {
	//     t.Skip("linux only test")
	// }
}

// =============================================================================
// SECTION 9: BENCHMARKS
// =============================================================================
//
// BENCHMARK FUNCTION SIGNATURE
//
//   func BenchmarkXxx(b *testing.B)
//
// RUNNING BENCHMARKS
//
//   go test -bench=.           - Run all benchmarks
//   go test -bench=BenchmarkX  - Run matching
//   go test -bench=. -benchmem - Include memory stats
//   go test -bench=. -count=5  - Run 5 times
//
// THE PATTERN
//
//   func BenchmarkXxx(b *testing.B) {
//       for i := 0; i < b.N; i++ {
//           // Code to benchmark
//       }
//   }
//
// b.N is adjusted by the test framework to get accurate timing.
//
// BENCHMARK METHODS
//
//   b.ResetTimer()    - Reset after expensive setup
//   b.StopTimer()     - Pause timing
//   b.StartTimer()    - Resume timing
//   b.ReportAllocs()  - Report memory allocations
//   b.SetBytes(n)     - For throughput benchmarks
//
// =============================================================================

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

func BenchmarkAddWithSetup(b *testing.B) {
	// Expensive setup
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	b.ResetTimer() // Don't count setup time

	for i := 0; i < b.N; i++ {
		Add(data[0], data[1])
	}
}

// Concat concatenates strings
func Concat(strs ...string) string {
	var result string
	for _, s := range strs {
		result += s
	}
	return result
}

// ConcatBuilder uses strings.Builder
func ConcatBuilder(strs ...string) string {
	var b strings.Builder
	for _, s := range strs {
		b.WriteString(s)
	}
	return b.String()
}

func BenchmarkConcat(b *testing.B) {
	strs := []string{"hello", "world", "this", "is", "a", "test"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Concat(strs...)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	strs := []string{"hello", "world", "this", "is", "a", "test"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ConcatBuilder(strs...)
	}
}

// Sub-benchmarks
func BenchmarkAddSizes(b *testing.B) {
	sizes := []struct {
		name string
		a, z int
	}{
		{"small", 1, 2},
		{"medium", 1000, 2000},
		{"large", 1000000, 2000000},
	}

	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Add(size.a, size.z)
			}
		})
	}
}

// =============================================================================
// SECTION 10: EXAMPLES
// =============================================================================
//
// EXAMPLE FUNCTION NAMING
//
//   func Example()           - Package example
//   func ExampleFunction()   - Function example
//   func ExampleType()       - Type example
//   func ExampleType_Method() - Method example
//   func ExampleFunction_suffix() - Multiple examples
//
// OUTPUT COMMENT
//
// Examples with // Output: comments are run as tests.
// The output must match exactly.
//
//   func ExampleAdd() {
//       fmt.Println(Add(2, 3))
//       // Output: 5
//   }
//
// For unordered output:
//   // Unordered output:
//
// =============================================================================

func ExampleAdd() {
	fmt.Println(Add(2, 3))
	// Output: 5
}

func ExampleAdd_negative() {
	fmt.Println(Add(-1, -2))
	// Output: -3
}

func ExampleDivide() {
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(result)
	// Output: 5
}

// =============================================================================
// SECTION 11: TEST DOUBLES (MOCKS, STUBS, FAKES)
// =============================================================================
//
// INTERFACE-BASED TESTING
//
// Define interfaces for dependencies, then provide test implementations.
//
// TYPES OF TEST DOUBLES
//
//   Stub  - Returns canned answers
//   Mock  - Verifies interactions
//   Fake  - Working implementation (e.g., in-memory DB)
//   Spy   - Records calls for later verification
//
// =============================================================================

// UserStore is an interface for user persistence
type UserStore interface {
	Get(id int) (string, error)
	Save(id int, name string) error
}

// UserService depends on UserStore
type UserService struct {
	store UserStore
}

func (s *UserService) GetUserName(id int) (string, error) {
	return s.store.Get(id)
}

// StubUserStore is a test stub
type StubUserStore struct {
	users map[int]string
}

func (s *StubUserStore) Get(id int) (string, error) {
	name, ok := s.users[id]
	if !ok {
		return "", errors.New("not found")
	}
	return name, nil
}

func (s *StubUserStore) Save(id int, name string) error {
	s.users[id] = name
	return nil
}

func TestUserService(t *testing.T) {
	store := &StubUserStore{
		users: map[int]string{
			1: "Alice",
			2: "Bob",
		},
	}
	service := &UserService{store: store}

	t.Run("existing user", func(t *testing.T) {
		name, err := service.GetUserName(1)
		assertNoError(t, err)
		if name != "Alice" {
			t.Errorf("got %s; want Alice", name)
		}
	})

	t.Run("missing user", func(t *testing.T) {
		_, err := service.GetUserName(999)
		assertError(t, err)
	})
}

// Reader testing with io.Reader
func ProcessInput(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(string(data)), nil
}

func TestProcessInput(t *testing.T) {
	// strings.Reader implements io.Reader
	input := strings.NewReader("hello")
	result, err := ProcessInput(input)
	assertNoError(t, err)
	if result != "HELLO" {
		t.Errorf("got %s; want HELLO", result)
	}
}

// =============================================================================
// SECTION 12: TESTING HTTP
// =============================================================================
//
// net/http/httptest provides:
//
//   httptest.NewRecorder() - Captures response for handler testing
//   httptest.NewServer(handler) - Creates real HTTP server for integration tests
//
// =============================================================================

// Example shown as comments since we can't import net/http in this demo
// without creating complex dependencies

/*
import (
    "net/http"
    "net/http/httptest"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, World!"))
}

func TestHelloHandler(t *testing.T) {
    // Create request
    req := httptest.NewRequest("GET", "/hello", nil)

    // Create recorder
    rr := httptest.NewRecorder()

    // Call handler
    HelloHandler(rr, req)

    // Check status
    if rr.Code != http.StatusOK {
        t.Errorf("status = %d; want %d", rr.Code, http.StatusOK)
    }

    // Check body
    if rr.Body.String() != "Hello, World!" {
        t.Errorf("body = %s; want Hello, World!", rr.Body.String())
    }
}

func TestWithServer(t *testing.T) {
    // Create test server
    ts := httptest.NewServer(http.HandlerFunc(HelloHandler))
    defer ts.Close()

    // Make real HTTP request
    resp, err := http.Get(ts.URL + "/hello")
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    // Check response
    if resp.StatusCode != http.StatusOK {
        t.Errorf("status = %d; want %d", resp.StatusCode, http.StatusOK)
    }
}
*/

// =============================================================================
// SECTION 13: FUZZING (Go 1.18+)
// =============================================================================
//
// FUZZ FUNCTION SIGNATURE
//
//   func FuzzXxx(f *testing.F)
//
// RUNNING FUZZING
//
//   go test -fuzz=FuzzXxx         - Run until stopped (Ctrl+C)
//   go test -fuzz=FuzzXxx -fuzztime=30s  - Run for 30 seconds
//
// THE PATTERN
//
//   func FuzzXxx(f *testing.F) {
//       // Seed corpus with initial values
//       f.Add("initial", "values")
//
//       // Fuzz target
//       f.Fuzz(func(t *testing.T, a string, b string) {
//           // Test with random inputs
//           // Panic or t.Error on failure
//       })
//   }
//
// =============================================================================

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FuzzReverse(f *testing.F) {
	// Seed corpus
	f.Add("hello")
	f.Add("世界")
	f.Add("")

	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)

		// Reversing twice should give original
		if orig != doubleRev {
			t.Errorf("double reverse mismatch: %q -> %q -> %q", orig, rev, doubleRev)
		}

		// Reverse should preserve length
		if len([]rune(orig)) != len([]rune(rev)) {
			t.Errorf("length mismatch: %d vs %d", len([]rune(orig)), len([]rune(rev)))
		}
	})
}

// =============================================================================
// SECTION 14: COMMON TESTING PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Golden files
	_ = `
	func TestOutput(t *testing.T) {
		result := GenerateOutput()

		golden := filepath.Join("testdata", "output.golden")
		if *update {
			os.WriteFile(golden, []byte(result), 0644)
		}

		expected, _ := os.ReadFile(golden)
		if result != string(expected) {
			t.Errorf("mismatch with golden file")
		}
	}
	`

	// Pattern 2: Test fixtures in testdata/
	_ = `
	// testdata/ directory is ignored by go build
	data, _ := os.ReadFile("testdata/input.json")
	`

	// Pattern 3: Environment-based skip
	_ = `
	func TestIntegration(t *testing.T) {
		if os.Getenv("INTEGRATION") == "" {
			t.Skip("set INTEGRATION to run")
		}
		// Integration test
	}
	`

	// Pattern 4: Table test with error check
	_ = `
	tests := []struct{
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"valid", "input", "output", false},
		{"invalid", "bad", "", true},
	}
	`
}

// =============================================================================
// SECTION 15: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes(t *testing.T) {
	// Mistake 1: Not using t.Helper() in helpers
	_ = `
	// WRONG - error points to helper, not caller
	func check(t *testing.T, got, want int) {
		if got != want {
			t.Errorf("got %d; want %d", got, want)
		}
	}

	// RIGHT
	func check(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d; want %d", got, want)
		}
	}
	`

	// Mistake 2: Loop variable capture in parallel tests
	_ = `
	// WRONG (before Go 1.22)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			use(tt)  // All goroutines see last tt!
		})
	}

	// RIGHT
	for _, tt := range tests {
		tt := tt  // Capture
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			use(tt)
		})
	}
	`

	// Mistake 3: Not cleaning up resources
	_ = `
	// WRONG
	func TestX(t *testing.T) {
		f, _ := os.Create("temp")
		// Test runs...
		// File never closed!
	}

	// RIGHT
	func TestX(t *testing.T) {
		f, _ := os.Create("temp")
		t.Cleanup(func() {
			f.Close()
			os.Remove("temp")
		})
	}
	`

	// Mistake 4: Using Fatal in goroutines
	_ = `
	// WRONG - Fatal/FailNow in goroutine doesn't stop test
	go func() {
		t.Fatal("error")  // Only stops goroutine, not test!
	}()

	// RIGHT - Signal main goroutine
	errCh := make(chan error)
	go func() {
		errCh <- someOperation()
	}()
	if err := <-errCh; err != nil {
		t.Fatal(err)
	}
	`

	// Mistake 5: Testing unexported functions directly
	_ = `
	// Consider testing through exported API instead
	// Or use _test.go in same package for whitebox testing
	`
}

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	// This package's functions are meant to be run via go test
	DemonstratePatterns()
}
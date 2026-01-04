// Package testing_practice is your daily practice space for Go testing.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against testing/testing.go
// 3. Note what you missed - focus on those tomorrow
package testing_practice

// =============================================================================
// EXERCISE 1: BASICS
// =============================================================================

func PracticeBasics() {
	// TODO: Test file naming convention
	// foo.go → ???

	// TODO: Test function signature
	// func ???(t *???.T)

	// TODO: What must test function name start with?
	// Answer: ???

	// TODO: Running tests
	// go ???              - run all tests
	// go ??? -v           - verbose
	// go ??? -run Regex   - run matching
	// go ??? ./...        - all packages
}

// =============================================================================
// EXERCISE 2: REPORTING
// =============================================================================

func PracticeReporting() {
	// TODO: Failure methods
	// t.???(args...)       - log and continue
	// t.???(fmt, args...)  - formatted, continue
	// t.???(args...)       - log and STOP test
	// t.???(fmt, args...)  - formatted, STOP test

	// TODO: Logging (shown with -v or on failure)
	// t.???(args...)
	// t.???(fmt, args...)

	// TODO: When use Error vs Fatal?
	// Error: ???
	// Fatal: ???
}

// =============================================================================
// EXERCISE 3: TABLE-DRIVEN TESTS
// =============================================================================

func PracticeTableDriven() {
	// TODO: Write the table test pattern
	// func TestXxx(t *testing.T) {
	//     tests := []struct {
	//         ???  string
	//         input  ???
	//         expected ???
	//     }{
	//         {"case1", input1, expected1},
	//     }
	//
	//     for _, tt := range tests {
	//         t.???(tt.name, func(t *testing.T) {
	//             result := Function(tt.input)
	//             if result != tt.expected {
	//                 t.???(...)
	//             }
	//         })
	//     }
	// }

	// TODO: Why use table-driven tests?
	// 1. ???
	// 2. ???
	// 3. ???
}

// =============================================================================
// EXERCISE 4: SUBTESTS
// =============================================================================

func PracticeSubtests() {
	// TODO: Create a subtest
	// t.???(name, func(t *testing.T) { ... })

	// TODO: Run specific subtest from command line
	// go test -run ???/???

	// TODO: What does t.Run return?
	// Answer: ???
}

// =============================================================================
// EXERCISE 5: PARALLEL TESTS
// =============================================================================

func PracticeParallel() {
	// TODO: Mark test as parallel
	// t.???()

	// TODO: Loop variable capture (before Go 1.22)
	// for _, tt := range tests {
	//     ??? := ???  // Capture!
	//     t.Run(tt.name, func(t *testing.T) {
	//         t.Parallel()
	//         use(tt)
	//     })
	// }

	// TODO: Why capture loop variable?
	// Answer: ???
}

// =============================================================================
// EXERCISE 6: TEST HELPERS
// =============================================================================

func PracticeHelpers() {
	// TODO: Mark function as helper
	// t.???()

	// TODO: Why use t.Helper()?
	// Answer: ???

	// TODO: Write a helper function
	// func assertEqual(t *testing.T, got, want int) {
	//     t.???()
	//     if got != want {
	//         t.???(...)
	//     }
	// }
}

// =============================================================================
// EXERCISE 7: SETUP AND CLEANUP
// =============================================================================

func PracticeSetupCleanup() {
	// TODO: Register cleanup function
	// t.???(func() {
	//     // Cleanup code
	// })

	// TODO: When does cleanup run?
	// Answer: ???

	// TODO: Multiple cleanups run in what order?
	// Answer: ???

	// TODO: Package-level setup function name?
	// func ???(m *testing.M)
}

// =============================================================================
// EXERCISE 8: SKIPPING
// =============================================================================

func PracticeSkipping() {
	// TODO: Skip a test
	// t.???(reason)
	// t.???(format, args...)

	// TODO: Check if -short flag passed
	// if testing.???() {
	//     t.Skip("skipping in short mode")
	// }

	// TODO: Run tests in short mode
	// go test ???
}

// =============================================================================
// EXERCISE 9: BENCHMARKS
// =============================================================================

func PracticeBenchmarks() {
	// TODO: Benchmark function signature
	// func ???(b *testing.B)

	// TODO: The benchmark loop
	// for i := 0; i < b.??? ; i++ {
	//     // Code to benchmark
	// }

	// TODO: Running benchmarks
	// go test ???=.           - run all
	// go test ???=. -benchmem  - with memory

	// TODO: Reset timer after setup
	// b.???()

	// TODO: Report memory allocations
	// b.???()
}

// =============================================================================
// EXERCISE 10: EXAMPLES
// =============================================================================

func PracticeExamples() {
	// TODO: Example function naming
	// func ???()             - package example
	// func ???Function()     - function example
	// func ???Type_Method()  - method example

	// TODO: Make example testable (checked by go test)
	// func ExampleAdd() {
	//     fmt.Println(Add(2, 3))
	//     // ???: 5
	// }

	// TODO: For unordered output
	// // ??? output:
}

// =============================================================================
// EXERCISE 11: TEST DOUBLES
// =============================================================================

func PracticeTestDoubles() {
	// TODO: Name 4 types of test doubles
	// 1. ??? - returns canned answers
	// 2. ??? - verifies interactions
	// 3. ??? - working implementation
	// 4. ??? - records calls

	// TODO: What Go feature enables test doubles?
	// Answer: ???

	// TODO: Common test double for io.Reader?
	// strings.???(str)
}

// =============================================================================
// EXERCISE 12: HTTP TESTING
// =============================================================================

func PracticeHTTPTesting() {
	// TODO: Package for HTTP testing helpers
	// net/http/???

	// TODO: Create response recorder
	// rr := httptest.???()

	// TODO: Create test request
	// req := httptest.???(method, url, body)

	// TODO: Create test server
	// ts := httptest.???(handler)
	// defer ts.???()
}

// =============================================================================
// EXERCISE 13: FUZZING
// =============================================================================

func PracticeFuzzing() {
	// TODO: Fuzz function signature
	// func ???(f *testing.F)

	// TODO: Add seed values
	// f.???(values...)

	// TODO: Fuzz target
	// f.???(func(t *testing.T, args...) {
	//     // Test with random inputs
	// })

	// TODO: Run fuzzing
	// go test ???=FuzzXxx
	// go test ???=FuzzXxx -fuzztime=30s
}

// =============================================================================
// EXERCISE 14: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// func check(t *testing.T, got, want int) {
	//     if got != want {
	//         t.Errorf("...")
	//     }
	// }
	// Answer: ???

	// Mistake 2: What's wrong?
	// for _, tt := range tests {
	//     t.Run(tt.name, func(t *testing.T) {
	//         t.Parallel()
	//         use(tt)
	//     })
	// }
	// Answer: ???

	// Mistake 3: What's wrong?
	// go func() {
	//     t.Fatal("error")
	// }()
	// Answer: ???

	// Mistake 4: What's wrong?
	// func TestX(t *testing.T) {
	//     f, _ := os.Create("temp")
	//     // test code
	// }
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// 1. Test function signature?
	_ = "func ???(t *testing.T)"

	// 2. Benchmark function signature?
	_ = "func ???(b *testing.B)"

	// 3. Table test subtest?
	_ = "t.???(name, func(t *testing.T){...})"

	// 4. Mark parallel?
	_ = "t.???()"

	// 5. Mark helper?
	_ = "t.???()"

	// 6. Register cleanup?
	_ = "t.???(func(){...})"

	// 7. Skip test?
	_ = "t.???(reason)"

	// 8. Check short mode?
	_ = "testing.???()"

	// 9. Reset benchmark timer?
	_ = "b.???()"

	// 10. Run benchmarks?
	_ = "go test -???=."

	// 11. Example output comment?
	_ = "// ???: value"

	// 12. HTTP test recorder?
	_ = "httptest.???()"

	// 13. Fuzz seed corpus?
	_ = "f.???(values)"

	// 14. Loop capture fix?
	_ = "tt := ???"
}

// =============================================================================
// MINI PROJECT: CALCULATOR TEST SUITE
// =============================================================================

func MiniProject() {
	// Build a complete test suite for a Calculator type that:
	//
	// 1. Has table-driven tests for basic operations
	// 2. Uses subtests for organization
	// 3. Includes benchmarks
	// 4. Has testable examples
	// 5. Uses helpers for assertions
	// 6. Includes a fuzz test
	//
	// Scaffold:
	//
	// // calculator.go
	// type Calculator struct{}
	//
	// func (c *Calculator) Add(a, b int) int { return a + b }
	// func (c *Calculator) Divide(a, b int) (int, error) {
	//     if b == 0 { return 0, errors.New("divide by zero") }
	//     return a / b, nil
	// }
	//
	// // calculator_test.go
	// func assertEqual(t *testing.T, got, want int) {
	//     t.Helper()
	//     if got != want {
	//         t.Errorf("got %d; want %d", got, want)
	//     }
	// }
	//
	// func TestCalculator_Add(t *testing.T) {
	//     tests := []struct {
	//         name     string
	//         a, b     int
	//         expected int
	//     }{
	//         {"positive", 2, 3, 5},
	//         {"negative", -1, -2, -3},
	//         {"zero", 0, 0, 0},
	//     }
	//
	//     c := &Calculator{}
	//     for _, tt := range tests {
	//         t.Run(tt.name, func(t *testing.T) {
	//             assertEqual(t, c.Add(tt.a, tt.b), tt.expected)
	//         })
	//     }
	// }
	//
	// func TestCalculator_Divide(t *testing.T) {
	//     c := &Calculator{}
	//
	//     t.Run("valid", func(t *testing.T) {
	//         result, err := c.Divide(10, 2)
	//         if err != nil {
	//             t.Fatalf("unexpected error: %v", err)
	//         }
	//         assertEqual(t, result, 5)
	//     })
	//
	//     t.Run("divide by zero", func(t *testing.T) {
	//         _, err := c.Divide(10, 0)
	//         if err == nil {
	//             t.Error("expected error")
	//         }
	//     })
	// }
	//
	// func BenchmarkCalculator_Add(b *testing.B) {
	//     c := &Calculator{}
	//     for i := 0; i < b.N; i++ {
	//         c.Add(2, 3)
	//     }
	// }
	//
	// func ExampleCalculator_Add() {
	//     c := &Calculator{}
	//     fmt.Println(c.Add(2, 3))
	//     // Output: 5
	// }
	//
	// func FuzzCalculator_Add(f *testing.F) {
	//     f.Add(1, 2)
	//     f.Add(-1, 1)
	//     f.Add(0, 0)
	//
	//     c := &Calculator{}
	//     f.Fuzz(func(t *testing.T, a, b int) {
	//         result := c.Add(a, b)
	//         // Verify commutativity
	//         if result != c.Add(b, a) {
	//             t.Errorf("Add not commutative: %d+%d vs %d+%d", a, b, b, a)
	//         }
	//     })
	// }
}

// RunAllPractice is not meant to be called - this file is for reading.
func AllPractice() {
	PracticeBasics()
	PracticeReporting()
	PracticeTableDriven()
	PracticeSubtests()
	PracticeParallel()
	PracticeHelpers()
	PracticeSetupCleanup()
	PracticeSkipping()
	PracticeBenchmarks()
	PracticeExamples()
	PracticeTestDoubles()
	PracticeHTTPTesting()
	PracticeFuzzing()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
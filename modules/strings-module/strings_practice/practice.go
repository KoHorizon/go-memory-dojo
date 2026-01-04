// Package strings_practice is your daily practice space for the fmt package.
//
// INSTRUCTIONS:
// 1. Each morning, open this file fresh
// 2. Fill in all the TODOs from memory (no peeking at strings.go!)
// 3. Run with: go run cmd/main.go --practice
// 4. Check your answers against the theory file
// 5. Note what you missed - focus on those tomorrow
//
// The goal is to make this automatic - muscle memory, not conscious recall.
package strings_practice

import (
	"fmt"
)

// =============================================================================
// EXERCISE 1: GENERAL VERBS
// =============================================================================
// Fill in the format verbs for general-purpose printing.

func PracticeGeneralVerbs() {
	fmt.Println("=== PRACTICE: GENERAL VERBS ===")

	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}

	// TODO: Print p with default formatting
	// Expected: {Alice 30}
	fmt.Printf("Default value: %v\n", p) // TODO: Replace %v with your answer

	// TODO: Print p with field names included
	// Expected: {Name:Alice Age:30}
	fmt.Printf("With field names: %v\n", p) // TODO: What verb shows field names?

	// TODO: Print p in Go syntax (paste-able code)
	// Expected: strings_practice.Person{Name:"Alice", Age:30}
	fmt.Printf("Go syntax: %v\n", p) // TODO: What verb shows Go syntax?

	// TODO: Print just the type of p
	// Expected: strings_practice.Person
	fmt.Printf("Type only: %v\n", p) // TODO: What verb shows the type?

	fmt.Println()
}

// =============================================================================
// EXERCISE 2: INTEGER VERBS
// =============================================================================
// Fill in the format verbs for integer printing.

func PracticeIntegerVerbs() {
	fmt.Println("=== PRACTICE: INTEGER VERBS ===")

	n := 255

	// TODO: Fill in the correct verbs
	fmt.Printf("Decimal:     %v\n", n) // TODO: base 10
	fmt.Printf("Binary:      %v\n", n) // TODO: base 2
	fmt.Printf("Octal:       %v\n", n) // TODO: base 8
	fmt.Printf("Hex (lower): %v\n", n) // TODO: base 16, lowercase
	fmt.Printf("Hex (upper): %v\n", n) // TODO: base 16, uppercase

	// Character printing
	char := 65 // This is 'A'

	fmt.Printf("As character: %v\n", char) // TODO: print as Unicode character
	fmt.Printf("Unicode:      %v\n", char) // TODO: print as U+XXXX
	fmt.Printf("Quoted char:  %v\n", char) // TODO: print as 'A'

	fmt.Println()
}

// =============================================================================
// EXERCISE 3: FLOATING-POINT VERBS
// =============================================================================
// Fill in the format verbs for float printing.

func PracticeFloatVerbs() {
	fmt.Println("=== PRACTICE: FLOATING-POINT VERBS ===")

	pi := 3.14159265358979

	// TODO: Fill in the correct verbs
	fmt.Printf("Decimal notation:    %v\n", pi) // TODO: e.g., 3.141593
	fmt.Printf("Scientific (lower):  %v\n", pi) // TODO: e.g., 3.141593e+00
	fmt.Printf("Scientific (upper):  %v\n", pi) // TODO: e.g., 3.141593E+00
	fmt.Printf("Smart format:        %v\n", pi) // TODO: auto-chooses best

	// Precision practice
	// TODO: Print pi with exactly 2 decimal places
	fmt.Printf("Two decimals: %v\n", pi) // TODO: should print 3.14

	// TODO: Print pi with 5 decimal places
	fmt.Printf("Five decimals: %v\n", pi) // TODO: should print 3.14159

	fmt.Println()
}

// =============================================================================
// EXERCISE 4: STRING AND BYTE VERBS
// =============================================================================
// Fill in the format verbs for strings and bytes.

func PracticeStringVerbs() {
	fmt.Println("=== PRACTICE: STRING VERBS ===")

	s := "Hello\tWorld"

	// TODO: Fill in the correct verbs
	fmt.Printf("Plain string: %v\n", s) // TODO: just the string
	fmt.Printf("Quoted:       %v\n", s) // TODO: with quotes and escapes visible
	fmt.Printf("Hex dump:     %v\n", s) // TODO: as hex bytes

	// []byte difference
	bytes := []byte("Hi")

	fmt.Printf("Bytes as text:   %v\n", bytes) // TODO: prints "Hi"
	fmt.Printf("Bytes as values: %v\n", bytes) // TODO: prints [72 105]

	fmt.Println()
}

// =============================================================================
// EXERCISE 5: WIDTH AND PRECISION
// =============================================================================
// Practice width and precision specifiers.

func PracticeWidthPrecision() {
	fmt.Println("=== PRACTICE: WIDTH AND PRECISION ===")

	n := 42
	pi := 3.14159

	// TODO: Print n with minimum width of 6, right-aligned
	// Expected: |    42|
	fmt.Printf("|%v|\n", n) // TODO: add width specifier

	// TODO: Print n with minimum width of 6, left-aligned
	// Expected: |42    |
	fmt.Printf("|%v|\n", n) // TODO: add width and alignment flag

	// TODO: Print n with minimum width of 6, zero-padded
	// Expected: |000042|
	fmt.Printf("|%v|\n", n) // TODO: add width and padding flag

	// TODO: Print pi with 2 decimal places
	// Expected: 3.14
	fmt.Printf("Pi: %v\n", pi) // TODO: add precision

	// TODO: Print pi with width 8, 2 decimal places
	// Expected: |    3.14|
	fmt.Printf("|%v|\n", pi) // TODO: add width and precision

	fmt.Println()
}

// =============================================================================
// EXERCISE 6: FLAGS
// =============================================================================
// Practice the various flags.

func PracticeFlags() {
	fmt.Println("=== PRACTICE: FLAGS ===")

	pos := 42
	neg := -42
	hex := 255

	// TODO: Print pos with a + sign always shown
	// Expected: +42
	fmt.Printf("With sign: %v\n", pos) // TODO: what flag always shows sign?

	// TODO: Print both pos and neg aligned (positive with space for sign)
	// Expected: " 42" and "-42" (note the space before 42)
	fmt.Printf("Aligned: %v and %v\n", pos, neg) // TODO: what flag adds space?

	// TODO: Print hex in hex format with 0x prefix
	// Expected: 0xff
	fmt.Printf("With prefix: %v\n", hex) // TODO: what flag adds alternate format?

	// TODO: Print time components with zero padding
	hour, min, sec := 9, 5, 1
	// Expected: 09:05:01
	fmt.Printf("Time: %v:%v:%v\n", hour, min, sec) // TODO: add formatting

	fmt.Println()
}

// =============================================================================
// EXERCISE 7: ESCAPING PERCENT
// =============================================================================
// Practice printing literal percent signs.

func PracticeEscaping() {
	fmt.Println("=== PRACTICE: ESCAPING ===")

	progress := 75.5

	// TODO: Print "Progress: 75.5%"
	// Hint: How do you print a literal % in a format string?
	fmt.Printf("Progress: %.1f\n", progress) // TODO: add the percent sign!

	// TODO: Print "100% complete!"
	fmt.Printf("complete!\n") // TODO: add "100%" at the beginning

	fmt.Println()
}

// =============================================================================
// EXERCISE 8: SPRINT FAMILY
// =============================================================================
// Practice returning strings instead of printing.

func PracticeSprintFamily() {
	fmt.Println("=== PRACTICE: SPRINT FAMILY ===")

	name := "server-01"
	port := 8080

	// TODO: Use Sprintf to create a URL string
	// Expected: "http://server-01:8080/api"
	// Hint: Replace the empty string with fmt.Sprintf(...)
	url := fmt.Sprintf("TODO: format string here", name, port)

	fmt.Printf("URL: %s\n", url)

	// TODO: Use Sprint to concatenate values
	// Expected: "Hello World 42"
	// Hint: fmt.Sprint concatenates its arguments with spaces
	result := fmt.Sprint("TODO") // TODO: use fmt.Sprint with multiple args

	fmt.Printf("Sprint result: %s\n", result)

	fmt.Println()
}

// =============================================================================
// EXERCISE 9: ARGUMENT INDEXING
// =============================================================================
// Practice referencing arguments by position.

func PracticeArgumentIndexing() {
	fmt.Println("=== PRACTICE: ARGUMENT INDEXING ===")

	n := 255

	// TODO: Print n in decimal, hex, binary, and octal - using just ONE argument
	// Expected: "Decimal: 255, Hex: ff, Binary: 11111111, Octal: 377"
	// Hint: Use [1] to reference the first argument multiple times
	fmt.Printf("Decimal: %d, Hex: %x, Binary: %b, Octal: %o\n", n, n, n, n) // TODO: rewrite with indexing

	// TODO: Print these out of order: "C, A, B" using arguments ("A", "B", "C")
	fmt.Printf("%s, %s, %s\n", "A", "B", "C") // TODO: rewrite to print C, A, B

	fmt.Println()
}

// =============================================================================
// EXERCISE 10: COMMON PATTERNS
// =============================================================================
// Practice real-world formatting patterns.

func PracticeCommonPatterns() {
	fmt.Println("=== PRACTICE: COMMON PATTERNS ===")

	// Pattern 1: Aligned table
	// TODO: Create a table with aligned columns
	// Expected output:
	// | Name       |  Size | Type |
	// | readme.md  |  1234 | file |
	// | src        |  4096 | dir  |

	// TODO: Fill in the Printf calls to create the table
	// fmt.Printf(...)

	// Pattern 2: Debug struct output
	type Config struct {
		Host string
		Port int
	}
	cfg := Config{"localhost", 8080}

	// TODO: Print cfg in a format you could paste into Go code
	fmt.Printf("Debug: %v\n", cfg) // TODO: what verb for Go syntax?

	// Pattern 3: Hex dump with spaces
	data := []byte("Hello")

	// TODO: Print as hex with spaces between bytes
	// Expected: "48 65 6c 6c 6f"
	fmt.Printf("Hex dump: %v\n", data) // TODO: what format shows hex with spaces?

	fmt.Println()
}

// =============================================================================
// BONUS CHALLENGE: FULL FORMAT SPECIFICATION
// =============================================================================
// The complete format: %[flags][width][.precision]verb
// Can you write one format string that uses ALL of these?

func BonusChallenge() {
	fmt.Println("=== BONUS CHALLENGE ===")

	value := 42.5

	// TODO: Write a format that:
	// - Always shows the sign (+/-)
	// - Has minimum width of 10
	// - Shows 2 decimal places
	// - Uses 'f' verb
	// Expected: |    +42.50|

	fmt.Printf("|%v|\n", value) // TODO: the full format specification

	fmt.Println()
}

// =============================================================================
// SELF-TEST: WRITE FROM MEMORY
// =============================================================================
// Without looking at the theory file, fill in these from pure memory.

func SelfTest() {
	fmt.Println("=== SELF-TEST ===")
	fmt.Println("Fill these in from memory. Be honest with yourself!")
	fmt.Println()

	// What verb prints the type of a value?
	_ = "%?" // TODO

	// What verb prints a struct with field names?
	_ = "%?" // TODO

	// What verb prints Go syntax representation?
	_ = "%?" // TODO

	// What verb prints integers in binary?
	_ = "%?" // TODO

	// What verb prints floats in scientific notation?
	_ = "%?" // TODO

	// What flag always shows the sign for numbers?
	_ = "%?" // TODO

	// What flag adds 0x prefix to hex?
	_ = "%?" // TODO

	// How do you print a literal percent sign?
	_ = "%?" // TODO

	// How do you reference the first argument explicitly?
	_ = "%?" // TODO

	fmt.Println("Check your answers against strings/strings.go!")
	fmt.Println()
}

// AllPractice executes all practice exercises.
func AllPractice() {
	PracticeGeneralVerbs()
	PracticeIntegerVerbs()
	PracticeFloatVerbs()
	PracticeStringVerbs()
	PracticeWidthPrecision()
	PracticeFlags()
	PracticeEscaping()
	PracticeSprintFamily()
	PracticeArgumentIndexing()
	PracticeCommonPatterns()
	BonusChallenge()
	SelfTest()

	fmt.Println("=====================================")
	fmt.Println("Practice complete!")
	fmt.Println("Now check your answers against:")
	fmt.Println("  modules/strings-module/strings/strings.go")
	fmt.Println("=====================================")
}
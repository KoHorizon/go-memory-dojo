// Package strings provides comprehensive documentation and working examples
// for Go's fmt package - the foundation of formatted I/O in Go.
package strings

import (
	"fmt"
	"os"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF FMT
// =============================================================================
//
// WHY DOES FMT EXIST?
//
// The fmt package implements formatted I/O, similar to C's printf and scanf.
// But why do we need "format verbs" at all? Why not just print things directly?
//
// The answer lies in how computers represent data internally vs. how humans
// read it. A number like 42 is stored as binary (101010), but we want to see
// "42" or maybe "0x2A" or "052" (octal). A pointer is just a memory address,
// but sometimes we want to see what it points TO.
//
// Format verbs are the bridge between internal representation and human-readable
// output. They tell Go: "Take this internal data and present it THIS way."
//
// THE ANATOMY OF A FORMAT STRING
//
// A format string like "Hello, %s! You are %d years old." contains:
// - Literal text: "Hello, ", "! You are ", " years old."
// - Format verbs: %s, %d
//
// When fmt.Printf processes this, it:
// 1. Copies literal text directly to output
// 2. When it hits %, it reads the next character(s) to determine the verb
// 3. Consumes the next argument and formats it according to the verb
// 4. Continues until the format string is exhausted
//
// =============================================================================

// =============================================================================
// SECTION 2: THE GENERAL VERBS
// =============================================================================
//
// Sometimes you don't know (or don't care) what type you're printing.
// General verbs work with ANY type, making them invaluable for debugging,
// logging, and generic code.
//
// %v  - Value in default format
//       bool:    true or false
//       int:     base 10 number
//       float:   decimal notation
//       string:  the string itself (no quotes)
//       slice:   [elem1 elem2 elem3]
//       struct:  {field1 field2 field3}
//       pointer: memory address in hex
//
// %+v - Like %v, but adds field names to structs
//       Useful because {John 30} is ambiguous, {Name:John Age:30} is clear
//
// %#v - Go syntax representation
//       Output you could paste back into Go code
//       Invaluable for debugging - copy output as test fixture
//
// %T  - Type of the value, not the value itself
//       Essential for debugging interface{} or understanding return types
//
// =============================================================================

func DemonstrateGeneralVerbs() {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Bob", Age: 25}

	// %v - default format
	fmt.Printf("User: %v\n", p) // {Bob 25}

	// %+v - with field names (much clearer!)
	fmt.Printf("User: %+v\n", p) // {Name:Bob Age:25}

	// %#v - Go syntax (paste-able code)
	fmt.Printf("Debug: %#v\n", p) // strings.Person{Name:"Bob", Age:25}

	// %T - reveal the type
	var x interface{} = []int{1, 2, 3}
	fmt.Printf("Type is: %T\n", x) // []int
}

// =============================================================================
// SECTION 3: INTEGER VERBS
// =============================================================================
//
// Humans invented multiple number systems, and computers use them all:
//
// %d - DECIMAL (Base 10)
//      The system we use daily. Each digit is 0-9.
//
// %b - BINARY (Base 2)
//      How computers ACTUALLY store numbers. Each digit is 0 or 1.
//      Used for: bit manipulation, understanding memory, flags
//
// %o - OCTAL (Base 8)
//      Each digit is 0-7. Historically used because 3 bits = 1 octal digit.
//      Used for: Unix file permissions (755, 644, etc.)
//
// %x, %X - HEXADECIMAL (Base 16)
//      Digits are 0-9, A-F. Each hex digit = exactly 4 bits.
//      Used for: memory addresses, colors (#FF0000), byte values
//      %x = lowercase (ff), %X = uppercase (FF)
//
// %c - CHARACTER
//      Prints the Unicode character for this number.
//      Works because characters ARE numbers (Unicode code points)
//
// %U - UNICODE
//      Prints the Unicode code point format: U+XXXX
//
// %q - QUOTED CHARACTER
//      Prints a quoted Go character literal with escape sequences
//
// =============================================================================

func DemonstrateIntegerVerbs() {
	n := 255 // A full byte: 11111111 in binary

	fmt.Printf("Decimal: %d\n", n)         // 255
	fmt.Printf("Binary:  %b\n", n)         // 11111111
	fmt.Printf("Octal:   %o\n", n)         // 377
	fmt.Printf("Hex:     %x / %X\n", n, n) // ff / FF

	// Unix file permissions use octal
	perms := 0755
	fmt.Printf("chmod %o\n", perms) // chmod 755

	// Colors in hex
	red := 0xFF0000
	fmt.Printf("Color: #%06X\n", red) // #FF0000

	// Characters are just numbers
	char := 65
	fmt.Printf("%c = U+%04X\n", char, char) // A = U+0041
}

// =============================================================================
// SECTION 4: FLOATING-POINT VERBS
// =============================================================================
//
// Floating-point numbers can represent a HUGE range of values:
// from 0.0000000000000001 to 99999999999999999999.0
//
// The challenge: how do you display them readably?
//
// %f - DECIMAL NOTATION
//      Always shows standard decimal form: 1234.5678
//      Problem: Very large/small numbers become unreadable
//
// %e, %E - SCIENTIFIC NOTATION
//      Shows as: coefficient × 10^exponent
//      Written as: 1.234568e+03 (meaning 1.234568 × 10³)
//      Good for: very large or very small numbers
//
// %g, %G - "SMART" FORMAT
//      Automatically chooses %e or %f based on the exponent
//      Usually what you want for "just show me the number nicely"
//
// PRECISION
//
// The precision specifier (%.Nf) means different things:
//   %f, %e, %E: N digits after the decimal point
//   %g, %G: N total significant digits
//
// Remember: floats have limited precision (~15-17 digits for float64).
// Asking for more precision than exists shows meaningless digits.
//
// =============================================================================

func DemonstrateFloatVerbs() {
	pi := 3.14159265358979
	avogadro := 6.02214076e23
	planck := 6.62607015e-34

	// %f for everyday numbers
	fmt.Printf("Pi: %.4f\n", pi) // 3.1416

	// %e for scientific values
	fmt.Printf("Avogadro: %e\n", avogadro) // 6.022141e+23
	fmt.Printf("Planck:   %e\n", planck)   // 6.626070e-34

	// %g auto-selects the best format
	fmt.Printf("Auto: %g\n", pi)       // 3.14159265358979
	fmt.Printf("Auto: %g\n", avogadro) // 6.02214076e+23

	// Precision control
	price := 19.99
	fmt.Printf("Price: $%.2f\n", price) // $19.99
}

// =============================================================================
// SECTION 5: STRING AND BYTE VERBS
// =============================================================================
//
// STRINGS VS BYTES IN GO
//
// - A string is an immutable sequence of bytes
// - A []byte is a mutable sequence of bytes
// - Both can represent text, but they're used differently
//
// %s - STRING
//      For strings and []byte, prints the raw bytes as text
//      WARNING: If bytes aren't valid UTF-8, you get garbled output
//
// %q - QUOTED STRING
//      Wraps in double quotes and escapes special characters
//      Essential for: seeing invisible characters, debugging encoding
//
// %x, %X - HEX DUMP
//      Shows each byte as two hex digits
//      "Hi" = 4869 (H=0x48, i=0x69)
//      With space flag: "% x" = "48 69"
//
// THE DIFFERENCE BETWEEN %s AND %v FOR []byte
//
// For plain strings: identical
// For []byte: %s prints as text, %v prints as [72 101 108 108 111]
//
// =============================================================================

func DemonstrateStringVerbs() {
	greeting := "Hello, 世界!"
	hidden := "Line1\tLine2\nLine3"

	fmt.Printf("%s\n", greeting)  // Hello, 世界!
	fmt.Printf("%q\n", greeting)  // "Hello, 世界!"
	fmt.Printf("% x\n", greeting) // 48 65 6c 6c 6f ...

	// %q reveals hidden characters
	fmt.Printf("Visible: %s\n", hidden) // Line1	Line2 (tab invisible!)
	fmt.Printf("Quoted:  %q\n", hidden) // "Line1\tLine2\nLine3"

	// []byte difference
	bytes := []byte("Hello")
	fmt.Printf("%%s: %s\n", bytes) // Hello
	fmt.Printf("%%v: %v\n", bytes) // [72 101 108 108 111]
}

// =============================================================================
// SECTION 6: POINTER VERBS
// =============================================================================
//
// A pointer is a variable that stores a memory address.
//
// %p - POINTER
//      Prints the memory address in hexadecimal with 0x prefix
//      Only works with pointer types, slices, maps, channels, functions
//
// WHY WOULD YOU PRINT A POINTER?
//
// 1. Debugging: "Are these two variables pointing to the same data?"
// 2. Understanding: "Is this a copy or the original?"
// 3. Troubleshooting: "Why is my data being modified unexpectedly?"
//
// =============================================================================

func DemonstratePointerVerbs() {
	x := 42
	ptr := &x
	fmt.Printf("Value: %d, Address: %p\n", x, ptr)

	// Are these pointing to the same underlying array?
	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice3 := make([]int, 3)
	copy(slice3, slice1)

	fmt.Printf("slice1: %p\n", slice1) // 0xc000...
	fmt.Printf("slice2: %p\n", slice2) // same address!
	fmt.Printf("slice3: %p\n", slice3) // different address
}

// =============================================================================
// SECTION 7: BOOLEAN VERBS
// =============================================================================
//
// Go is strict about booleans. Unlike C or JavaScript, Go booleans are
// ONLY true or false. Period.
//
// %t - BOOLEAN
//      Simply prints "true" or "false"
//
// =============================================================================

func DemonstrateBooleanVerbs() {
	isAdmin := true
	hasAccess := false

	fmt.Printf("Admin: %t, Access: %t\n", isAdmin, hasAccess)

	// Common pattern: conditional logging
	debug := true
	if debug {
		fmt.Printf("[DEBUG] Feature enabled: %t\n", hasAccess)
	}
}

// =============================================================================
// SECTION 8: WIDTH AND PRECISION
// =============================================================================
//
// THE ANATOMY OF A FORMAT SPECIFICATION
//
// Full format: %[flags][width][.precision]verb
//
// Example: %+08.2f
//   % = start of format verb
//   + = flag (always show sign)
//   0 = flag (pad with zeros)
//   8 = width (minimum 8 characters)
//   .2 = precision (2 decimal places)
//   f = verb (decimal float)
//
// WIDTH
//
// Width is the MINIMUM number of characters to output.
// If shorter, it gets padded. If longer, it overflows (width is not max).
// Default: spaces on left (right-aligned)
// Use '-' flag to left-align (pad on right)
// Use '0' flag to pad with zeros (numbers only)
//
// PRECISION
//
// Meaning varies by type:
//   - Floats (%f, %e): digits after decimal point
//   - Floats (%g): total significant digits
//   - Strings (%s): maximum characters to print (truncates!)
//   - Integers: minimum digits (pads with zeros)
//
// DYNAMIC WIDTH AND PRECISION
//
// Use * to read width/precision from arguments:
//   fmt.Printf("%*d", 5, 42)       = "   42" (width 5)
//   fmt.Printf("%.*f", 2, 3.14159) = "3.14" (precision 2)
//
// =============================================================================

func DemonstrateWidthAndPrecision() {
	// Width for alignment
	fmt.Printf("|%10s|%10s|\n", "Name", "Score")
	fmt.Printf("|%10s|%10d|\n", "Alice", 95)
	fmt.Printf("|%10s|%10d|\n", "Bob", 87)

	// Left-align with -
	fmt.Printf("|%-10s|%-10d|\n", "Alice", 95)

	// Zero-padding for IDs and timestamps
	id := 42
	fmt.Printf("ID: %05d\n", id) // ID: 00042

	hour, min, sec := 9, 5, 1
	fmt.Printf("Time: %02d:%02d:%02d\n", hour, min, sec) // 09:05:01

	// Precision for floats
	pi := 3.14159265
	fmt.Printf("$%.2f\n", pi) // $3.14

	// Precision for strings (truncation!)
	long := "Hello, World!"
	fmt.Printf("%.5s...\n", long) // Hello...

	// Dynamic width
	for _, width := range []int{5, 10, 15} {
		fmt.Printf("%*s|\n", width, "Hi")
	}
}

// =============================================================================
// SECTION 9: FLAGS
// =============================================================================
//
// FLAGS MODIFY VERB BEHAVIOR
//
// Flags come after % and before width. Multiple flags can be combined.
//
// + (PLUS)
//   Always print a sign for numeric values.
//   Use case: bank statements, temperature changes
//
// - (MINUS)
//   Left-justify within the field width.
//   Use case: tables, aligned output
//
// (SPACE)
//   Leave a space for positive numbers (where + would go).
//   Helps align positive/negative numbers in columns.
//
// # (HASH/SHARP)
//   "Alternate format" - meaning varies by verb:
//   %#o = adds leading 0 (0377 instead of 377)
//   %#x = adds leading 0x (0xff instead of ff)
//   %#v = Go syntax representation
//
// 0 (ZERO)
//   Pad with zeros instead of spaces.
//   Use case: timestamps, codes, IDs
//
// =============================================================================

func DemonstrateFlags() {
	// + flag: always show sign
	change := 42
	fmt.Printf("Balance change: %+d\n", change)  // +42
	fmt.Printf("Balance change: %+d\n", -change) // -42

	// space flag: align positive/negative
	values := []int{42, -17, 8, -3}
	for _, v := range values {
		fmt.Printf("% d\n", v) // space before positive, - before negative
	}

	// # flag: alternate format
	n := 255
	fmt.Printf("Octal:  %#o\n", n) // 0377
	fmt.Printf("Hex:    %#x\n", n) // 0xff
	fmt.Printf("Hex:    %#X\n", n) // 0XFF

	// Combining flags: zero-padded, always signed
	temp := 7
	fmt.Printf("Temp: %+03d°C\n", temp) // +07°C
}

// =============================================================================
// SECTION 10: ESCAPING PERCENT SIGNS
// =============================================================================
//
// THE PROBLEM
//
// The % character is special - it starts a format verb.
// So how do you print a literal percent sign?
//
// THE SOLUTION: %%
//
// Use %% to print a single %.
// The first % "consumes" the escape meaning, the second is literal.
//
// =============================================================================

func DemonstrateEscaping() {
	progress := 75.5
	fmt.Printf("Progress: %.1f%%\n", progress) // 75.5%
	fmt.Printf("100%% complete!\n")

	// Common pattern: percentage display
	passed, total := 85, 100
	pct := float64(passed) / float64(total) * 100
	fmt.Printf("Score: %d/%d (%.0f%%)\n", passed, total, pct)
}

// =============================================================================
// SECTION 11: PRINT FAMILY FUNCTIONS
// =============================================================================
//
// THE PRINT FAMILY TREE
//
// BASE FUNCTIONS (stdout):
//   Print   - prints arguments, no newline, spaces between non-strings
//   Println - prints arguments, adds newline, always spaces between
//   Printf  - prints formatted string (format verbs)
//
// F-PREFIX (to io.Writer):
//   Fprint, Fprintln, Fprintf
//   First argument is an io.Writer (file, buffer, network connection)
//
// S-PREFIX (returns string):
//   Sprint, Sprintln, Sprintf
//   Returns the string instead of printing it
//   Use case: building strings, logging, testing
//
// =============================================================================

func DemonstratePrintFamily() {
	name := "Alice"
	age := 30

	// Printf - formatted output
	fmt.Printf("%s is %d years old\n", name, age)

	// Sprintf - returns string (most common for building strings)
	url := fmt.Sprintf("https://api.example.com/users/%s?age=%d", name, age)
	fmt.Println(url)

	// Fprintf - write to any io.Writer
	fmt.Fprintf(os.Stderr, "[ERROR] User %s not found\n", name)

	// Sprint variants for string building
	msg := fmt.Sprintf("Hello, %s!", name)
	_ = msg
}

// =============================================================================
// SECTION 12: COMMON PATTERNS AND IDIOMS
// =============================================================================
//
// These patterns come up constantly in real Go code.
//
// =============================================================================

func DemonstrateCommonPatterns() {
	// Pattern 1: Building URLs and paths
	host := "api.example.com"
	version := "v2"
	resource := "users"
	id := 123
	url := fmt.Sprintf("https://%s/%s/%s/%d", host, version, resource, id)
	_ = url

	// Pattern 2: Formatted logging
	level := "INFO"
	component := "auth"
	message := "user logged in"
	fmt.Printf("[%s] %s: %s\n", level, component, message)

	// Pattern 3: Debug output with %#v
	type Request struct {
		Method string
		Path   string
		Body   []byte
	}
	req := Request{"POST", "/api/users", []byte(`{"name":"bob"}`)}
	fmt.Printf("DEBUG: %#v\n", req)

	// Pattern 4: Aligned table output
	data := []struct {
		Name string
		Size int
		Type string
	}{
		{"readme.md", 1234, "file"},
		{"src", 4096, "dir"},
		{"main.go", 567, "file"},
	}
	fmt.Printf("%-12s %8s %6s\n", "Name", "Size", "Type")
	for _, d := range data {
		fmt.Printf("%-12s %8d %6s\n", d.Name, d.Size, d.Type)
	}

	// Pattern 5: Error messages with context
	filename := "config.yaml"
	line := 42
	errMsg := "unexpected token"
	fmt.Printf("%s:%d: %s\n", filename, line, errMsg)

	// Pattern 6: Hex dump for debugging bytes
	data2 := []byte("Hello!")
	fmt.Printf("Hex: % x\n", data2)
}

// =============================================================================
// SECTION 13: ARGUMENT INDEXING
// =============================================================================
//
// EXPLICIT ARGUMENT INDEXING
//
// By default, Printf consumes arguments left-to-right. But sometimes you
// need to use the same argument multiple times, or use them out of order.
//
// SYNTAX: %[n]verb
//   [n] specifies which argument to use (1-indexed)
//
// WHY WOULD YOU NEED THIS?
//
// 1. Reusing arguments: same value in multiple formats
// 2. Localization: different languages have different word orders
// 3. Complex formatting: show same value in decimal and hex
//
// =============================================================================

func DemonstrateArgumentIndexing() {
	// Show same value in multiple formats
	n := 255
	fmt.Printf("Decimal: %[1]d, Hex: %[1]x, Binary: %[1]b\n", n)

	// Repeat a value
	name := "Alert"
	code := 500
	fmt.Printf("%[1]s! Code %[2]d. Repeat: %[1]s!\n", name, code)

	// Reorder arguments
	fmt.Printf("Third=%[3]s First=%[1]s Second=%[2]s\n", "A", "B", "C")
}

// =============================================================================
// SECTION 14: ERROR HANDLING IN FMT
// =============================================================================
//
// WHAT HAPPENS WHEN THINGS GO WRONG?
//
// fmt is very forgiving - it tries to print SOMETHING rather than crash.
//
// WRONG TYPE:
//   %d with a string → %!d(string=hello)
//   The %! prefix signals a formatting error
//
// MISSING ARGUMENT:
//   %d with no argument → %!d(MISSING)
//
// EXTRA ARGUMENTS:
//   More args than verbs → extra args printed with %!(EXTRA ...)
//
// These error strings appear inline in output. They don't cause panics.
// This is by design: fmt prioritizes "always produce output" over
// "fail loudly on mistakes."
//
// =============================================================================

func DemonstrateErrorHandling() {
	// These demonstrate what errors look like - don't do this intentionally!
	fmt.Printf("Wrong type: %d\n", "hello") // %!d(string=hello)
	fmt.Printf("Missing: %d %d\n", 42)      // 42 %!d(MISSING)
	fmt.Printf("Extra: %d\n", 1, 2, 3)      // 1 %!(EXTRA int=2, int=3)
}

// =============================================================================
// SECTION 15: PERFORMANCE CONSIDERATIONS
// =============================================================================
//
// WHEN DOES PERFORMANCE MATTER?
//
// For most uses, fmt is plenty fast. But in hot loops or high-throughput
// scenarios, these considerations matter:
//
// SPRINTF ALLOCATES
//   Every Sprintf call allocates a new string. In a loop, this creates
//   garbage collection pressure.
//
// INTERFACE{} OVERHEAD
//   Printf accepts interface{} arguments, which means values get boxed.
//   For maximum performance in critical paths, use strconv directly:
//     strconv.Itoa(42) is faster than fmt.Sprintf("%d", 42)
//
// But measure first - premature optimization is the root of all evil.
//
// =============================================================================

func DemonstratePerformance() {
	// For hot paths, consider strconv
	// import "strconv"
	// s := strconv.Itoa(42)         // faster than fmt.Sprintf("%d", 42)
	// s := strconv.FormatFloat(...) // faster than fmt.Sprintf("%f", ...)

	// For building many strings, use strings.Builder
	// var b strings.Builder
	// for i := 0; i < 1000; i++ {
	//     fmt.Fprintf(&b, "item %d\n", i)
	// }
	// result := b.String()
}

// RunAllDemonstrations executes all examples in order.
func RunAllDemonstrations() {
	DemonstrateGeneralVerbs()
	DemonstrateIntegerVerbs()
	DemonstrateFloatVerbs()
	DemonstrateStringVerbs()
	DemonstratePointerVerbs()
	DemonstrateBooleanVerbs()
	DemonstrateWidthAndPrecision()
	DemonstrateFlags()
	DemonstrateEscaping()
	DemonstratePrintFamily()
	DemonstrateCommonPatterns()
	DemonstrateArgumentIndexing()
	DemonstrateErrorHandling()
	DemonstratePerformance()
}
// Package strings provides comprehensive documentation and working examples
// for Go's fmt package - the foundation of formatted I/O in Go.
//
// This file is designed to be read, understood, and practiced until the
// concepts become second nature.
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
// WHY DO GENERAL VERBS EXIST?
//
// Sometimes you don't know (or don't care) what type you're printing.
// General verbs work with ANY type, making them invaluable for debugging,
// logging, and generic code.
//
// %v - THE "VALUE" VERB
//
// %v stands for "value" - it prints the value in its default format.
// But what IS the "default format"? It depends on the type:
//   - bool:    true or false
//   - int:     base 10 number
//   - float:   decimal notation (e.g., 3.14)
//   - string:  the string itself (no quotes)
//   - slice:   [elem1 elem2 elem3]
//   - struct:  {field1 field2 field3}
//   - pointer: memory address in hex
//
// WHY %+v AND %#v?
//
// %+v adds field names to structs. This exists because {John 30} is
// ambiguous - is 30 an age? An ID? A score? {Name:John Age:30} is clear.
//
// %#v gives "Go syntax representation" - output you could paste back into
// Go code. This is invaluable for debugging because you can literally
// copy the output and use it as a test fixture.
//
// %T - THE "TYPE" VERB
//
// Prints the TYPE of a value, not the value itself. Essential for debugging
// interface{} values or understanding what a function actually returned.
//
// =============================================================================

func DemonstrateGeneralVerbs() {
	fmt.Println("=== GENERAL VERBS ===")
	fmt.Println()

	// Basic types with %v
	name := "Alice"
	age := 30
	height := 5.8
	active := true

	fmt.Printf("%%v with string:  %v\n", name)   // Alice
	fmt.Printf("%%v with int:     %v\n", age)    // 30
	fmt.Printf("%%v with float:   %v\n", height) // 5.8
	fmt.Printf("%%v with bool:    %v\n", active) // true
	fmt.Println()

	// Structs show why %+v matters
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Bob", Age: 25}

	fmt.Printf("%%v  (struct):  %v\n", p)  // {Bob 25} - ambiguous!
	fmt.Printf("%%+v (struct):  %+v\n", p) // {Name:Bob Age:25} - clear!
	fmt.Printf("%%#v (struct):  %#v\n", p) // strings.Person{Name:"Bob", Age:25} - Go syntax!
	fmt.Println()

	// %T reveals the type
	var x interface{} = 42
	fmt.Printf("%%T reveals type: %T\n", x)         // int
	fmt.Printf("%%T on slice:     %T\n", []int{})   // []int
	fmt.Printf("%%T on map:       %T\n", map[string]int{}) // map[string]int
	fmt.Println()
}

// =============================================================================
// SECTION 3: INTEGER VERBS
// =============================================================================
//
// WHY SO MANY INTEGER FORMATS?
//
// Humans invented multiple number systems, and computers use them all:
//
// DECIMAL (%d) - Base 10
//   The system we use daily. Each digit is 0-9.
//   Example: 255 means 2×100 + 5×10 + 5×1
//
// BINARY (%b) - Base 2
//   How computers ACTUALLY store numbers. Each digit is 0 or 1.
//   Example: 255 = 11111111 (eight 1s = 8 bits = 1 byte all "on")
//   Used for: bit manipulation, understanding memory, flags
//
// OCTAL (%o) - Base 8
//   Each digit is 0-7. Historically used because 3 bits = 1 octal digit.
//   Example: 255 = 377 (3×64 + 7×8 + 7×1)
//   Used for: Unix file permissions (755, 644, etc.)
//
// HEXADECIMAL (%x, %X) - Base 16
//   Digits are 0-9, A-F. Each hex digit = exactly 4 bits.
//   Example: 255 = FF (15×16 + 15×1)
//   Used for: memory addresses, colors (#FF0000), byte values
//   %x = lowercase (ff), %X = uppercase (FF)
//
// CHARACTER (%c)
//   Prints the Unicode character for this number.
//   Example: 65 = 'A', 128512 = '😀'
//   This works because characters ARE numbers (Unicode code points)
//
// UNICODE (%U)
//   Prints the Unicode code point format: U+XXXX
//   Example: 65 = U+0041
//   Used for: documentation, debugging character issues
//
// QUOTED CHARACTER (%q)
//   Prints a quoted Go character literal.
//   Example: 65 = 'A', 9 = '\t' (tab character)
//   Shows escape sequences for non-printable characters
//
// =============================================================================

func DemonstrateIntegerVerbs() {
	fmt.Println("=== INTEGER VERBS ===")
	fmt.Println()

	n := 255 // A nice number: it's 11111111 in binary (one full byte)

	fmt.Printf("Decimal %%d:     %d\n", n)  // 255
	fmt.Printf("Binary %%b:      %b\n", n)  // 11111111
	fmt.Printf("Octal %%o:       %o\n", n)  // 377
	fmt.Printf("Hex lower %%x:   %x\n", n)  // ff
	fmt.Printf("Hex upper %%X:   %X\n", n)  // FF
	fmt.Println()

	// Character representations
	char := 65 // ASCII/Unicode for 'A'
	fmt.Printf("Character %%c:    %c\n", char) // A
	fmt.Printf("Unicode %%U:      %U\n", char) // U+0041
	fmt.Printf("Quoted %%q:       %q\n", char) // 'A'
	fmt.Println()

	// Why octal matters: Unix permissions
	perms := 0755 // rwxr-xr-x
	fmt.Printf("File permissions: %o (octal) = %d (decimal)\n", perms, perms)
	fmt.Println()

	// Why hex matters: colors, memory
	red := 0xFF0000
	fmt.Printf("Red color: #%06X\n", red) // #FF0000 (padded to 6 digits)
	fmt.Println()
}

// =============================================================================
// SECTION 4: FLOATING-POINT VERBS
// =============================================================================
//
// WHY MULTIPLE FLOAT FORMATS?
//
// Floating-point numbers can represent a HUGE range of values:
// from 0.0000000000000001 to 99999999999999999999.0
//
// The challenge: how do you display them readably?
//
// %f - DECIMAL NOTATION
//   Always shows the number in standard decimal form.
//   Example: 1234.5678
//   Problem: Very large numbers become unreadable (1234567890000.000000)
//            Very small numbers show mostly zeros (0.000001)
//
// %e, %E - SCIENTIFIC NOTATION
//   Shows as: coefficient × 10^exponent
//   Written as: 1.234568e+03 (meaning 1.234568 × 10³ = 1234.568)
//   %e = lowercase 'e', %E = uppercase 'E'
//   Good for: very large or very small numbers
//
// %g, %G - "SMART" FORMAT (General)
//   Automatically chooses %e or %f based on the exponent.
//   Rule: Uses %e for exponents < -4 or >= precision
//   This is usually what you want for "just show me the number nicely"
//
// %x, %X - HEXADECIMAL FLOAT
//   Shows the exact binary representation in hex.
//   Example: 0x1.921fb54442d18p+001 for π
//   Used for: precise reproduction of floating-point values,
//            debugging float representation issues
//
// PRECISION AND FLOATING-POINT
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
	fmt.Println("=== FLOATING-POINT VERBS ===")
	fmt.Println()

	pi := 3.14159265358979323846
	avogadro := 6.02214076e23
	planck := 6.62607015e-34

	fmt.Println("Standard decimal %f:")
	fmt.Printf("  Pi:        %f\n", pi)       // 3.141593
	fmt.Printf("  Avogadro:  %f\n", avogadro) // huge, unreadable
	fmt.Printf("  Planck:    %f\n", planck)   // 0.000000 (too small!)
	fmt.Println()

	fmt.Println("Scientific notation %e:")
	fmt.Printf("  Pi:        %e\n", pi)       // 3.141593e+00
	fmt.Printf("  Avogadro:  %e\n", avogadro) // 6.022141e+23
	fmt.Printf("  Planck:    %e\n", planck)   // 6.626070e-34
	fmt.Println()

	fmt.Println("Smart format %g (chooses best):")
	fmt.Printf("  Pi:        %g\n", pi)       // 3.14159265358979
	fmt.Printf("  Avogadro:  %g\n", avogadro) // 6.02214076e+23
	fmt.Printf("  Planck:    %g\n", planck)   // 6.62607015e-34
	fmt.Println()

	fmt.Println("Precision control:")
	fmt.Printf("  Pi with %%.2f:  %.2f\n", pi) // 3.14
	fmt.Printf("  Pi with %%.5f:  %.5f\n", pi) // 3.14159
	fmt.Printf("  Pi with %%.2e:  %.2e\n", pi) // 3.14e+00
	fmt.Printf("  Pi with %%.5g:  %.5g\n", pi) // 3.1416 (5 sig figs)
	fmt.Println()
}

// =============================================================================
// SECTION 5: STRING AND BYTE VERBS
// =============================================================================
//
// STRINGS VS BYTES IN GO
//
// This is crucial to understand:
// - A string is an immutable sequence of bytes
// - A []byte is a mutable sequence of bytes
// - Both can represent text, but they're used differently
//
// Under the hood, strings are just (pointer, length) pairs pointing to
// byte arrays. When you iterate over a string with range, Go decodes
// UTF-8 and gives you runes (Unicode code points).
//
// %s - STRING
//   For strings and []byte, prints the raw bytes as text.
//   WARNING: If the bytes aren't valid UTF-8, you get garbled output.
//
// %q - QUOTED STRING
//   Wraps the string in double quotes and escapes special characters.
//   Example: "Hello\tWorld" instead of Hello	World
//   Essential for: seeing invisible characters, generating Go code,
//                 debugging encoding issues
//
// %x, %X - HEX DUMP
//   Shows each byte as two hex digits.
//   "Hi" = 4869 (H=0x48, i=0x69)
//   Used for: debugging binary data, seeing raw bytes
//   With space flag: "% x" = "48 69" (space between bytes)
//
// THE DIFFERENCE BETWEEN %s AND %v FOR STRINGS
//
// For plain strings: identical output
// For []byte: %s prints as text, %v prints as [72 101 108 108 111]
// This matters when debugging - sometimes you WANT to see the bytes!
//
// =============================================================================

func DemonstrateStringVerbs() {
	fmt.Println("=== STRING AND BYTE VERBS ===")
	fmt.Println()

	greeting := "Hello, 世界!" // Mix of ASCII and UTF-8
	hidden := "Line1\tLine2\nLine3"

	fmt.Printf("%%s (string):     %s\n", greeting) // Hello, 世界!
	fmt.Printf("%%q (quoted):     %q\n", greeting) // "Hello, 世界!"
	fmt.Printf("%%x (hex):        %x\n", greeting) // 48656c6c6f2c20e4b896e7958c21
	fmt.Printf("%% x (hex+space): % x\n", greeting) // 48 65 6c 6c 6f 2c 20 e4 b8 96...
	fmt.Println()

	fmt.Println("Why %q matters (reveals hidden characters):")
	fmt.Printf("  %%s: %s\n", hidden)   // Line1	Line2 (tab invisible!)
	fmt.Printf("  %%q: %q\n", hidden)   // "Line1\tLine2\nLine3"
	fmt.Println()

	// The []byte difference
	bytes := []byte{72, 101, 108, 108, 111} // "Hello"
	fmt.Println("[]byte printing difference:")
	fmt.Printf("  %%s: %s\n", bytes)  // Hello (as text)
	fmt.Printf("  %%v: %v\n", bytes)  // [72 101 108 108 111] (as bytes)
	fmt.Printf("  %%q: %q\n", bytes)  // "Hello" (quoted text)
	fmt.Println()
}

// =============================================================================
// SECTION 6: POINTER VERBS
// =============================================================================
//
// WHAT IS A POINTER?
//
// A pointer is a variable that stores a memory address. When you see
// 0xc0000140b8, that's saying "look at memory location 0xc0000140b8".
//
// %p - POINTER
//   Prints the memory address in hexadecimal with 0x prefix.
//   Only works with pointer types, slices, maps, channels, and functions.
//
// WHY WOULD YOU PRINT A POINTER?
//
// 1. Debugging: "Are these two variables pointing to the same data?"
// 2. Understanding: "Is this a copy or the original?"
// 3. Troubleshooting: "Why is my data being modified unexpectedly?"
//
// COMMON GOTCHA
//
// In Go, slices, maps, and channels are "reference types" - the variable
// itself contains a pointer to the underlying data. So when you pass a
// slice to a function, you're passing a copy of the slice header, but
// both headers point to the same underlying array!
//
// =============================================================================

func DemonstratePointerVerbs() {
	fmt.Println("=== POINTER VERBS ===")
	fmt.Println()

	x := 42
	ptr := &x

	fmt.Printf("Value of x:           %d\n", x)
	fmt.Printf("Address of x (%%p):    %p\n", ptr)
	fmt.Println()

	// Demonstrating when pointers matter
	slice1 := []int{1, 2, 3}
	slice2 := slice1 // Same underlying array!

	fmt.Println("Slice reference demonstration:")
	fmt.Printf("  slice1 header points to: %p\n", slice1)
	fmt.Printf("  slice2 header points to: %p\n", slice2)
	fmt.Println("  They're the same! Modifying one affects the other.")
	fmt.Println()

	// Contrast with a real copy
	slice3 := make([]int, len(slice1))
	copy(slice3, slice1)
	fmt.Printf("  slice3 (real copy) points to: %p\n", slice3)
	fmt.Println("  Different address - truly independent.")
	fmt.Println()
}

// =============================================================================
// SECTION 7: BOOLEAN VERBS
// =============================================================================
//
// GO'S BOOLEAN PHILOSOPHY
//
// Go is strict about booleans. Unlike C (where 0 is false, non-zero is true)
// or JavaScript (with its "truthy/falsy" values), Go booleans are ONLY
// true or false. Period.
//
// %t - BOOLEAN
//   Simply prints "true" or "false".
//   That's it. No 1/0, no yes/no, no on/off.
//
// WHY SO SIMPLE?
//
// Clarity over cleverness. When you see "true" or "false", there's no
// ambiguity. This matches Go's philosophy of explicit over implicit.
//
// =============================================================================

func DemonstrateBooleanVerbs() {
	fmt.Println("=== BOOLEAN VERBS ===")
	fmt.Println()

	yes := true
	no := false

	fmt.Printf("%%t with true:   %t\n", yes)  // true
	fmt.Printf("%%t with false:  %t\n", no)   // false
	fmt.Println()

	// Common pattern: conditional output
	isAdmin := true
	fmt.Printf("User is admin: %t\n", isAdmin)
	fmt.Println()
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
// Width is the MINIMUM number of characters to output. If the value is
// shorter, it gets padded. If longer, it overflows (width is not max).
//
// By default, padding is with spaces on the left (right-aligned).
// Use '-' flag to left-align (pad on right).
// Use '0' flag to pad with zeros (numbers only).
//
// PRECISION
//
// Precision means different things for different types:
//   - Floats (%f, %e): digits after decimal point
//   - Floats (%g): total significant digits
//   - Strings (%s): maximum characters to print (truncates!)
//   - Integers: minimum digits (pads with zeros)
//
// DYNAMIC WIDTH AND PRECISION
//
// Use * to read width/precision from arguments:
//   fmt.Printf("%*d", 5, 42)     = "   42" (width 5)
//   fmt.Printf("%.*f", 2, 3.14159) = "3.14" (precision 2)
//   fmt.Printf("%*.*f", 8, 2, 3.14) = "    3.14" (width 8, precision 2)
//
// =============================================================================

func DemonstrateWidthAndPrecision() {
	fmt.Println("=== WIDTH AND PRECISION ===")
	fmt.Println()

	// Width examples
	fmt.Println("Width (minimum characters):")
	fmt.Printf("  |%5d|  (width 5, right-aligned)\n", 42)
	fmt.Printf("  |%-5d| (width 5, left-aligned with -)\n", 42)
	fmt.Printf("  |%05d| (width 5, zero-padded)\n", 42)
	fmt.Println()

	// Precision for floats
	pi := 3.14159265358979
	fmt.Println("Precision for floats:")
	fmt.Printf("  %%.2f:  %.2f\n", pi)   // 3.14
	fmt.Printf("  %%.4f:  %.4f\n", pi)   // 3.1416
	fmt.Printf("  %%.0f:  %.0f\n", pi)   // 3 (no decimal)
	fmt.Printf("  %%8.2f: %8.2f (width 8, precision 2)\n", pi)
	fmt.Println()

	// Precision for strings (TRUNCATION!)
	longStr := "Hello, World!"
	fmt.Println("Precision for strings (truncates!):")
	fmt.Printf("  %%.5s: %.5s\n", longStr)   // Hello
	fmt.Printf("  %%10.5s: %10.5s\n", longStr) // "     Hello" (width 10, max 5 chars)
	fmt.Println()

	// Dynamic width and precision
	fmt.Println("Dynamic width/precision with *:")
	fmt.Printf("  %%*d with 6:    |%*d|\n", 6, 42)
	fmt.Printf("  %%.*f with 3:   %.3f\n", 3, pi)
	fmt.Printf("  %%*.*f with 10,2: |%*.*f|\n", 10, 2, pi)
	fmt.Println()
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
//   Positive numbers get +, negative get -.
//   Example: %+d with 42 = "+42"
//   Use case: bank statements, temperature changes
//
// - (MINUS)
//   Left-justify within the field width.
//   Default is right-justify.
//   Example: %-5d with 42 = "42   "
//   Use case: tables, aligned output
//
// (SPACE)
//   Leave a space for positive numbers (where + would go).
//   Helps align positive/negative numbers in columns.
//   Example: % d with 42 = " 42", with -42 = "-42"
//
// # (HASH/SHARP)
//   "Alternate format" - meaning varies by verb:
//   %#o = adds leading 0 (0377 instead of 377)
//   %#x = adds leading 0x (0xff instead of ff)
//   %#X = adds leading 0X (0XFF instead of FF)
//   %#v = Go syntax representation
//   %#q = backquoted string if possible
//
// 0 (ZERO)
//   Pad with zeros instead of spaces.
//   Only applies to numbers.
//   The - flag overrides 0 flag.
//   Example: %05d with 42 = "00042"
//   Use case: timestamps (09:05:01), codes, IDs
//
// =============================================================================

func DemonstrateFlags() {
	fmt.Println("=== FLAGS ===")
	fmt.Println()

	pos := 42
	neg := -42

	fmt.Println("+ flag (always show sign):")
	fmt.Printf("  %%+d:  %+d, %+d\n", pos, neg) // +42, -42
	fmt.Println()

	fmt.Println("- flag (left-justify):")
	fmt.Printf("  |%%5d|:  |%5d|\n", pos)   // |   42|
	fmt.Printf("  |%%-5d|: |%-5d|\n", pos)  // |42   |
	fmt.Println()

	fmt.Println("space flag (space for positive):")
	fmt.Printf("  %% d:  % d\n", pos)  //  42 (note leading space)
	fmt.Printf("  %% d:  % d\n", neg)  // -42
	fmt.Println()

	fmt.Println("# flag (alternate format):")
	fmt.Printf("  %%o:   %o\n", 255)    // 377
	fmt.Printf("  %%#o:  %#o\n", 255)   // 0377
	fmt.Printf("  %%x:   %x\n", 255)    // ff
	fmt.Printf("  %%#x:  %#x\n", 255)   // 0xff
	fmt.Printf("  %%#X:  %#X\n", 255)   // 0XFF
	fmt.Println()

	fmt.Println("0 flag (zero padding):")
	fmt.Printf("  %%5d:  %5d\n", pos)   //    42
	fmt.Printf("  %%05d: %05d\n", pos)  // 00042
	fmt.Println()

	// Practical example: timestamp
	hour, min, sec := 9, 5, 1
	fmt.Printf("Timestamp: %02d:%02d:%02d\n", hour, min, sec) // 09:05:01
	fmt.Println()
}

// =============================================================================
// SECTION 10: ESCAPING PERCENT SIGNS
// =============================================================================
//
// THE PROBLEM
//
// The % character is special in format strings - it starts a format verb.
// So how do you print a literal percent sign?
//
// THE SOLUTION: %%
//
// Use %% to print a single %.
//
// WHY DOUBLE PERCENT?
//
// This is a common pattern in string escaping:
// - In regex, \\ prints a single \
// - In SQL, '' prints a single '
// - In format strings, %% prints a single %
//
// The first character "consumes" the escape meaning, the second is literal.
//
// Think of it as: "% says 'interpret the next character specially'.
// When that next character is %, the special interpretation is 'print %'."
//
// =============================================================================

func DemonstrateEscaping() {
	fmt.Println("=== ESCAPING PERCENT SIGNS ===")
	fmt.Println()

	progress := 75.5

	// WRONG: This would try to parse 'p' as a verb
	// fmt.Printf("Progress: %f%\n", progress) // ERROR!

	// RIGHT: Use %% to print %
	fmt.Printf("Progress: %.1f%%\n", progress) // Progress: 75.5%
	fmt.Printf("100%% complete!\n")            // 100% complete!
	fmt.Println()

	// Common use case: percentage display
	passed := 85
	total := 100
	pct := float64(passed) / float64(total) * 100
	fmt.Printf("Score: %d/%d (%.0f%%)\n", passed, total, pct)
	fmt.Println()
}

// =============================================================================
// SECTION 11: PRINT FAMILY FUNCTIONS
// =============================================================================
//
// THE PRINT FAMILY TREE
//
// fmt has several printing functions. They follow a naming pattern:
//
// BASE FUNCTIONS (stdout):
//   Print   - prints arguments, no newline, spaces between non-strings
//   Println - prints arguments, adds newline, always spaces between
//   Printf  - prints formatted string (format verbs)
//
// F-PREFIX (to io.Writer):
//   Fprint, Fprintln, Fprintf
//   First argument is an io.Writer (file, buffer, network connection, etc.)
//   Use case: writing to files, HTTP responses, buffers
//
// S-PREFIX (returns string):
//   Sprint, Sprintln, Sprintf
//   Returns the string instead of printing it.
//   Use case: building strings, logging, testing
//
// WHY THREE VARIANTS?
//
// Print/Println: Simple cases, debugging, quick output
// Printf: Formatted output with control over presentation
//
// The "Ln" suffix means "line" - adds newline at end.
// The "f" suffix means "format" - takes a format string.
//
// PRINT VS PRINTLN SPACING
//
// Subtle but important:
//   Print(a, b, c)   - spaces only between non-strings
//   Println(a, b, c) - always spaces, plus newline at end
//
// =============================================================================

func DemonstratePrintFamily() {
	fmt.Println("=== PRINT FAMILY FUNCTIONS ===")
	fmt.Println()

	name := "Alice"
	age := 30

	// Print vs Println vs Printf
	fmt.Print("Print: ")
	fmt.Print(name, age) // No newline, space between
	fmt.Println()        // Manual newline

	fmt.Println("Println:", name, age) // Auto newline, always spaces

	fmt.Printf("Printf: %s is %d years old\n", name, age) // Formatted
	fmt.Println()

	// Sprint family - returns strings
	s := fmt.Sprintf("%s is %d", name, age)
	fmt.Println("Sprintf returned:", s)
	fmt.Println()

	// Fprint family - writes to io.Writer
	fmt.Fprint(os.Stdout, "Fprint to stdout: ", name, "\n")
	fmt.Fprintf(os.Stdout, "Fprintf to stdout: %s\n", name)
	fmt.Println()
}

// =============================================================================
// SECTION 12: COMMON PATTERNS AND IDIOMS
// =============================================================================
//
// These patterns come up constantly in real Go code.
//
// =============================================================================

func DemonstrateCommonPatterns() {
	fmt.Println("=== COMMON PATTERNS ===")
	fmt.Println()

	// Pattern 1: Building formatted strings
	name := "server-01"
	port := 8080
	url := fmt.Sprintf("http://%s:%d/api", name, port)
	fmt.Println("URL building:", url)
	fmt.Println()

	// Pattern 2: Debug printing with %#v
	type Config struct {
		Host    string
		Port    int
		Enabled bool
	}
	cfg := Config{"localhost", 3000, true}
	fmt.Printf("Debug: %#v\n", cfg)
	fmt.Println()

	// Pattern 3: Aligned table output
	fmt.Println("Aligned table:")
	fmt.Printf("  %-10s %8s %6s\n", "Name", "Size", "Type")
	fmt.Printf("  %-10s %8d %6s\n", "readme.md", 1234, "file")
	fmt.Printf("  %-10s %8d %6s\n", "src", 4096, "dir")
	fmt.Printf("  %-10s %8d %6s\n", "main.go", 567, "file")
	fmt.Println()

	// Pattern 4: Hex dump for debugging bytes
	data := []byte("Hello!")
	fmt.Printf("Hex dump: % x\n", data)
	fmt.Println()

	// Pattern 5: Conditional formatting
	value := 42.567
	if value >= 0 {
		fmt.Printf("Positive: +%.2f\n", value)
	} else {
		fmt.Printf("Negative: %.2f\n", value)
	}
	fmt.Println()

	// Pattern 6: Error messages with context
	filename := "config.yaml"
	lineNum := 42
	errMsg := "unexpected token"
	fmt.Printf("Error in %s:%d: %s\n", filename, lineNum, errMsg)
	fmt.Println()
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
// Example: fmt.Printf("%[2]d %[1]d", 10, 20) prints "20 10"
//
// WHY WOULD YOU NEED THIS?
//
// 1. Reusing arguments: "The %[1]s is %[2]d. I repeat, the %[1]s is %[2]d."
// 2. Localization: Different languages have different word orders
// 3. Complex formatting: Show same value in multiple formats
//
// COMBINING WITH WIDTH/PRECISION
//
// You can use [n] for width and precision too:
//   %[3]*[2].*[1]f = width from arg 3, precision from arg 2, value from arg 1
//
// =============================================================================

func DemonstrateArgumentIndexing() {
	fmt.Println("=== ARGUMENT INDEXING ===")
	fmt.Println()

	// Reuse the same argument
	name := "Alert"
	code := 500
	fmt.Printf("%[1]s! Code %[2]d. I repeat: %[1]s, code %[2]d!\n", name, code)
	fmt.Println()

	// Different formats for the same value
	n := 255
	fmt.Printf("Decimal: %[1]d, Hex: %[1]x, Binary: %[1]b, Octal: %[1]o\n", n)
	fmt.Println()

	// Out-of-order arguments
	fmt.Printf("Third: %[3]s, First: %[1]s, Second: %[2]s\n", "A", "B", "C")
	fmt.Println()
}

// =============================================================================
// SECTION 14: ERROR HANDLING IN FMT
// =============================================================================
//
// WHAT HAPPENS WHEN THINGS GO WRONG?
//
// fmt is very forgiving - it tries to print SOMETHING rather than crash.
// Understanding these error outputs helps with debugging.
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
// BAD INDEX:
//   %[0]d or %[99]d → %!(BADINDEX)
//
// BAD WIDTH/PRECISION:
//   %*d with non-int width → %!(BADWIDTH)
//   %.*f with non-int precision → %!(BADPREC)
//
// These error strings are printed INLINE with your output. They don't
// cause panics or return errors - they show up in the output string.
// This is by design: fmt prioritizes "always produce output" over
// "fail loudly on mistakes."
//
// =============================================================================

func DemonstrateErrorHandling() {
	fmt.Println("=== ERROR HANDLING IN FMT ===")
	fmt.Println()

	fmt.Println("Wrong type:")
	fmt.Printf("  %%d with string: %d\n", "hello") // %!d(string=hello)
	fmt.Println()

	fmt.Println("Missing argument:")
	fmt.Printf("  %%d %%d with one arg: %d %d\n", 42) // 42 %!d(MISSING)
	fmt.Println()

	fmt.Println("Extra arguments:")
	fmt.Printf("  One verb, three args: %d\n", 1, 2, 3) // 1 %!(EXTRA int=2 int=3)
	fmt.Println()

	fmt.Println("Note: These print inline rather than causing errors!")
	fmt.Println("This is intentional - fmt prefers output over crashes.")
	fmt.Println()
}

// =============================================================================
// SECTION 15: PERFORMANCE CONSIDERATIONS
// =============================================================================
//
// WHEN DOES PERFORMANCE MATTER?
//
// For most uses, fmt is plenty fast. But in hot loops or high-throughput
// scenarios (logging millions of entries, formatting in tight loops),
// these considerations matter:
//
// SPRINTF ALLOCATES
//   Every Sprintf call allocates a new string. In a loop, this creates
//   garbage collection pressure. Consider using a bytes.Buffer with
//   Fprintf for many writes.
//
// INTERFACE{} OVERHEAD
//   Printf accepts interface{} arguments, which means:
//   - Values get boxed (memory allocation for small types)
//   - Type assertions happen at runtime
//   For maximum performance in critical paths, use strconv directly.
//
// BUFFER REUSE
//   sync.Pool with bytes.Buffer can reduce allocations.
//   But measure first - premature optimization is the root of all evil.
//
// WHEN TO USE STRCONV
//   strconv.Itoa(42) is faster than fmt.Sprintf("%d", 42)
//   strconv.FormatFloat is faster than Sprintf for floats
//   But the difference only matters in very hot paths.
//
// =============================================================================

func DemonstratePerformance() {
	fmt.Println("=== PERFORMANCE CONSIDERATIONS ===")
	fmt.Println()

	fmt.Println("For typical use: just use fmt, it's fine!")
	fmt.Println()
	fmt.Println("For hot paths, consider:")
	fmt.Println("  - strconv for simple conversions")
	fmt.Println("  - bytes.Buffer + Fprintf for building strings")
	fmt.Println("  - sync.Pool for buffer reuse")
	fmt.Println()
	fmt.Println("But always measure before optimizing!")
	fmt.Println()
}

// RunAllDemonstrations executes all examples in order.
// This is called from the main command to show all theory examples.
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
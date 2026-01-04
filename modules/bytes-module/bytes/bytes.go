// Package bytesmod provides comprehensive documentation and working examples
// for Go's bytes package - byte slice manipulation and buffering.
package bytesmod

import (
	"bytes"
	"fmt"
	"io"
	"unicode"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF THE BYTES PACKAGE
// =============================================================================
//
// WHY BYTES VS STRINGS?
//
// Go has two types for text-like data:
//
//   string  - Immutable sequence of bytes (read-only)
//   []byte  - Mutable sequence of bytes (read-write)
//
// THE FUNDAMENTAL DIFFERENCE
//
//   s := "hello"
//   s[0] = 'H'  // COMPILE ERROR - strings are immutable
//
//   b := []byte("hello")
//   b[0] = 'H'  // OK - []byte is mutable
//
// WHEN TO USE EACH
//
//   Use string when:
//   - Data won't change
//   - Passing as function parameters (no copy needed)
//   - Map keys (strings are comparable, []byte is not)
//   - API contracts (most APIs use string)
//
//   Use []byte when:
//   - Modifying data in place
//   - Building/accumulating bytes
//   - Working with I/O (io.Reader, io.Writer)
//   - Protocol parsing (binary data)
//   - Performance-critical modifications
//
// THE BYTES PACKAGE
//
// The bytes package provides functions for []byte that mirror
// the strings package, plus:
//   - bytes.Buffer: Efficient byte accumulator (like strings.Builder)
//   - bytes.Reader: io.Reader from []byte
//   - In-place modifications possible (unlike strings)
//
// CONVERTING BETWEEN STRING AND []BYTE
//
//   s := string(b)    // []byte → string (allocates and copies)
//   b := []byte(s)    // string → []byte (allocates and copies)
//
// Both conversions COPY the data. Avoid in hot paths.
//
// =============================================================================

// =============================================================================
// SECTION 2: COMPARISON FUNCTIONS
// =============================================================================
//
// Equal(a, b []byte) bool
//   Are the slices equal?
//   Unlike strings, []byte can't use ==
//
// EqualFold(s, t []byte) bool
//   Case-insensitive comparison
//
// Compare(a, b []byte) int
//   Returns: < 0 if a < b, 0 if a == b, > 0 if a > b
//   Lexicographic comparison
//
// WHY CAN'T YOU USE == FOR []BYTE?
//
//   []byte is a slice (reference type), == compares references, not contents
//   b1 := []byte("hello")
//   b2 := []byte("hello")
//   b1 == b2  // false! Different underlying arrays
//
//   Must use: bytes.Equal(b1, b2)  // true
//
// =============================================================================

func DemonstrateComparison() {
	b1 := []byte("hello")
	b2 := []byte("hello")
	b3 := []byte("Hello")

	// Equal - the only way to compare []byte for equality
	fmt.Printf("b1 == b2: %v\n", bytes.Equal(b1, b2)) // true
	fmt.Printf("b1 == b3: %v\n", bytes.Equal(b1, b3)) // false

	// EqualFold - case-insensitive
	fmt.Printf("EqualFold(b1, b3): %v\n", bytes.EqualFold(b1, b3)) // true

	// Compare
	fmt.Printf("Compare(b1, b2): %d\n", bytes.Compare(b1, b2)) // 0
	fmt.Printf("Compare(b1, b3): %d\n", bytes.Compare(b1, b3)) // 1 (lowercase > uppercase)

	// Why == doesn't work
	fmt.Printf("b1 == b2 (reference): %v\n", b1 == nil) // false
	// b1 == b2 would be a compile error
}

// =============================================================================
// SECTION 3: SEARCHING - CONTAINS AND INDEX
// =============================================================================
//
// CONTAINS FAMILY
//
// Contains(b, subslice []byte) bool
// ContainsAny(b []byte, chars string) bool
// ContainsRune(b []byte, r rune) bool
// ContainsFunc(b []byte, f func(rune) bool) bool
//
// INDEX FAMILY
//
// Index(s, sep []byte) int
// LastIndex(s, sep []byte) int
// IndexByte(b []byte, c byte) int
// LastIndexByte(b []byte, c byte) int
// IndexRune(s []byte, r rune) int
// IndexAny(s []byte, chars string) int
// LastIndexAny(s []byte, chars string) int
// IndexFunc(s []byte, f func(rune) bool) int
// LastIndexFunc(s []byte, f func(rune) bool) int
//
// COUNT
//
// Count(s, sep []byte) int
//
// Same semantics as strings package.
//
// =============================================================================

func DemonstrateSearching() {
	data := []byte("The quick brown fox jumps over the lazy dog")

	// Contains
	fmt.Printf("Contains 'fox': %v\n", bytes.Contains(data, []byte("fox")))
	fmt.Printf("Contains 'cat': %v\n", bytes.Contains(data, []byte("cat")))

	// ContainsAny
	fmt.Printf("Contains vowel: %v\n", bytes.ContainsAny(data, "aeiou"))
	fmt.Printf("Contains digit: %v\n", bytes.ContainsAny(data, "0123456789"))

	// ContainsRune
	fmt.Printf("Contains 'x': %v\n", bytes.ContainsRune(data, 'x'))

	// ContainsFunc
	hasUpper := bytes.ContainsFunc(data, unicode.IsUpper)
	fmt.Printf("Has uppercase: %v\n", hasUpper)

	// Index
	fmt.Printf("Index of 'fox': %d\n", bytes.Index(data, []byte("fox")))
	fmt.Printf("LastIndex of 'o': %d\n", bytes.LastIndex(data, []byte("o")))

	// IndexByte - fastest for single byte
	fmt.Printf("IndexByte of 'q': %d\n", bytes.IndexByte(data, 'q'))

	// IndexFunc
	firstDigit := bytes.IndexFunc(data, unicode.IsDigit)
	fmt.Printf("First digit at: %d\n", firstDigit) // -1 (none)

	// Count
	fmt.Printf("Count of 'o': %d\n", bytes.Count(data, []byte("o")))
}

// =============================================================================
// SECTION 4: PREFIX, SUFFIX, AND CUTTING
// =============================================================================
//
// HasPrefix(s, prefix []byte) bool
// HasSuffix(s, suffix []byte) bool
// TrimPrefix(s, prefix []byte) []byte
// TrimSuffix(s, suffix []byte) []byte
// CutPrefix(s, prefix []byte) (after []byte, found bool)
// CutSuffix(s, suffix []byte) (before []byte, found bool)
// Cut(s, sep []byte) (before, after []byte, found bool)
//
// Same semantics as strings package.
//
// =============================================================================

func DemonstratePrefixSuffix() {
	filename := []byte("document.txt")

	// HasPrefix/HasSuffix
	fmt.Printf("Starts with 'doc': %v\n", bytes.HasPrefix(filename, []byte("doc")))
	fmt.Printf("Ends with '.txt': %v\n", bytes.HasSuffix(filename, []byte(".txt")))

	// TrimPrefix/TrimSuffix
	noExt := bytes.TrimSuffix(filename, []byte(".txt"))
	fmt.Printf("Without .txt: %s\n", noExt)

	// Cut - split around separator
	email := []byte("user@example.com")
	user, domain, found := bytes.Cut(email, []byte("@"))
	if found {
		fmt.Printf("User: %s, Domain: %s\n", user, domain)
	}

	// CutPrefix/CutSuffix
	if name, ok := bytes.CutSuffix(filename, []byte(".txt")); ok {
		fmt.Printf("Name: %s\n", name)
	}
}

// =============================================================================
// SECTION 5: SPLITTING AND JOINING
// =============================================================================
//
// Split(s, sep []byte) [][]byte
//   Split on every occurrence of sep
//   Returns slice of slices
//
// SplitN(s, sep []byte, n int) [][]byte
//   Split with limit
//
// SplitAfter(s, sep []byte) [][]byte
//   Split but keep separator
//
// SplitAfterN(s, sep []byte, n int) [][]byte
//
// Fields(s []byte) [][]byte
//   Split on whitespace (like strings.Fields)
//
// FieldsFunc(s []byte, f func(rune) bool) [][]byte
//   Split using custom predicate
//
// Join(s [][]byte, sep []byte) []byte
//   Join slices with separator
//
// CRITICAL DIFFERENCE FROM STRINGS
//
// Returns [][]byte (slice of slices), not []string!
// The result shares memory with the original (no copying).
// Modifying original affects the split pieces.
//
// =============================================================================

func DemonstrateSplitJoin() {
	// Basic split
	csv := []byte("apple,banana,cherry")
	fruits := bytes.Split(csv, []byte(","))
	fmt.Printf("Fruits: %s\n", fruits) // [[97 112 112 108 101] ...]

	// More readable: convert to strings for printing
	for i, fruit := range fruits {
		fmt.Printf("  %d: %s\n", i, fruit)
	}

	// SplitN - with limit
	path := []byte("a/b/c/d/e")
	parts := bytes.SplitN(path, []byte("/"), 3)
	fmt.Printf("Split with limit: %s\n", parts)

	// Fields - split on whitespace
	text := []byte("  one   two\tthree\nfour  ")
	words := bytes.Fields(text)
	fmt.Printf("Words: %d\n", len(words))
	for _, word := range words {
		fmt.Printf("  %s\n", word)
	}

	// Join - inverse of split
	joined := bytes.Join(fruits, []byte(" | "))
	fmt.Printf("Joined: %s\n", joined)

	// IMPORTANT: Split shares memory!
	original := []byte("a,b,c")
	parts2 := bytes.Split(original, []byte(","))
	parts2[0][0] = 'X' // Modifies original!
	fmt.Printf("Original after modify: %s\n", original) // "X,b,c"

	// To avoid: use Clone or copy
	original2 := []byte("a,b,c")
	parts3 := bytes.Split(original2, []byte(","))
	safePart := bytes.Clone(parts3[0])
	safePart[0] = 'X'
	fmt.Printf("Original unchanged: %s\n", original2) // "a,b,c"
}

// =============================================================================
// SECTION 6: TRIMMING
// =============================================================================
//
// Trim(s []byte, cutset string) []byte
// TrimLeft(s []byte, cutset string) []byte
// TrimRight(s []byte, cutset string) []byte
// TrimSpace(s []byte) []byte
// TrimFunc(s []byte, f func(rune) bool) []byte
// TrimLeftFunc(s []byte, f func(rune) bool) []byte
// TrimRightFunc(s []byte, f func(rune) bool) []byte
//
// Same semantics as strings package.
// Removes runes from cutset, not substring!
//
// =============================================================================

func DemonstrateTrim() {
	// TrimSpace - most common
	padded := []byte("   hello world   \n\t")
	trimmed := bytes.TrimSpace(padded)
	fmt.Printf("Trimmed: '%s'\n", trimmed)

	// Trim with cutset
	bracketed := []byte("[[hello]]")
	fmt.Printf("Trim brackets: %s\n", bytes.Trim(bracketed, "[]"))

	// TrimFunc
	digits := []byte("123hello456")
	noDigits := bytes.TrimFunc(digits, unicode.IsDigit)
	fmt.Printf("Trim digits: %s\n", noDigits)

	// Remember: Trim removes characters, not substring
	fmt.Printf("Trim 'lo': %s\n", bytes.Trim([]byte("hello"), "lo"))        // "he"
	fmt.Printf("TrimSuffix 'lo': %s\n", bytes.TrimSuffix([]byte("hello"), []byte("lo"))) // "hel"
}

// =============================================================================
// SECTION 7: CASE CONVERSION
// =============================================================================
//
// ToUpper(s []byte) []byte
// ToLower(s []byte) []byte
// ToTitle(s []byte) []byte
// ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte
// ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte
// ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte
// ToValidUTF8(s, replacement []byte) []byte
//   Replace invalid UTF-8 with replacement
//
// IMPORTANT: These return NEW slices (allocate)
// Even though []byte is mutable, these don't modify in place.
//
// =============================================================================

func DemonstrateCase() {
	text := []byte("Hello, World!")

	// Basic conversions
	upper := bytes.ToUpper(text)
	lower := bytes.ToLower(text)
	fmt.Printf("Upper: %s\n", upper)
	fmt.Printf("Lower: %s\n", lower)
	fmt.Printf("Original: %s\n", text) // Unchanged!

	// ToValidUTF8 - replace invalid sequences
	invalid := []byte{0xff, 0xfe, 0xfd}
	valid := bytes.ToValidUTF8(invalid, []byte("?"))
	fmt.Printf("Replaced invalid: %s\n", valid) // "???"

	// Pattern: In-place uppercase (if you really need it)
	// Note: Only works for ASCII
	inPlace := []byte("hello")
	for i := 0; i < len(inPlace); i++ {
		if inPlace[i] >= 'a' && inPlace[i] <= 'z' {
			inPlace[i] -= 32 // 'a' - 'A' = 32
		}
	}
	fmt.Printf("In-place upper: %s\n", inPlace)
}

// =============================================================================
// SECTION 8: REPLACE AND MAP
// =============================================================================
//
// Replace(s, old, new []byte, n int) []byte
//   Replace first n occurrences (-1 for all)
//
// ReplaceAll(s, old, new []byte) []byte
//   Replace all occurrences
//
// Map(mapping func(rune) rune, s []byte) []byte
//   Apply function to each rune
//   Return -1 to drop the rune
//
// Repeat(b []byte, count int) []byte
//   Repeat count times
//
// =============================================================================

func DemonstrateReplace() {
	text := []byte("foo bar foo baz foo")

	// Replace first N
	fmt.Printf("Replace 1: %s\n", bytes.Replace(text, []byte("foo"), []byte("XXX"), 1))
	fmt.Printf("Replace all: %s\n", bytes.ReplaceAll(text, []byte("foo"), []byte("XXX")))

	// Map - transform each rune
	rot13 := func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return 'a' + (r-'a'+13)%26
		}
		if r >= 'A' && r <= 'Z' {
			return 'A' + (r-'A'+13)%26
		}
		return r
	}
	encoded := bytes.Map(rot13, []byte("Hello"))
	fmt.Printf("ROT13: %s\n", encoded)

	// Map to remove characters
	onlyDigits := bytes.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1 // Drop
	}, []byte("abc123def456"))
	fmt.Printf("Only digits: %s\n", onlyDigits)

	// Repeat
	fmt.Printf("Repeat: %s\n", bytes.Repeat([]byte("Go"), 5))
}

// =============================================================================
// SECTION 9: RUNES OPERATIONS
// =============================================================================
//
// Runes(s []byte) []rune
//   Convert []byte to []rune (UTF-8 decode)
//   Each rune is one Unicode code point
//
// EXAMPLE:
//
//   b := []byte("Hello, 世界")
//   fmt.Println(len(b))           // 13 bytes
//   r := bytes.Runes(b)
//   fmt.Println(len(r))           // 9 runes
//
// =============================================================================

func DemonstrateRunes() {
	text := []byte("Hello, 世界!")

	// Length in bytes vs runes
	fmt.Printf("Bytes: %d\n", len(text))
	runes := bytes.Runes(text)
	fmt.Printf("Runes: %d\n", len(runes))

	// Iterate runes
	for i, r := range runes {
		fmt.Printf("  %d: %c (U+%04X)\n", i, r, r)
	}

	// Reverse by runes (not bytes!)
	reversed := make([]rune, len(runes))
	for i, r := range runes {
		reversed[len(runes)-1-i] = r
	}
	fmt.Printf("Reversed: %s\n", string(reversed))
}

// =============================================================================
// SECTION 10: bytes.Buffer - EFFICIENT BYTE ACCUMULATION
// =============================================================================
//
// bytes.Buffer is like strings.Builder but for []byte.
// It's the primary way to efficiently build byte slices.
//
// CREATING
//
//   var buf bytes.Buffer              // Zero value is ready
//   buf := new(bytes.Buffer)          // Explicit
//   buf := bytes.NewBuffer(initial)   // With initial content
//   buf := bytes.NewBufferString(s)   // From string
//
// WRITING
//
//   Write(p []byte) (n int, err error)    - Implements io.Writer
//   WriteByte(c byte) error
//   WriteRune(r rune) (n int, err error)
//   WriteString(s string) (n int, err error)
//
// READING
//
//   Read(p []byte) (n int, err error)     - Implements io.Reader
//   ReadByte() (byte, error)
//   ReadRune() (r rune, size int, err error)
//   ReadBytes(delim byte) ([]byte, error)
//   ReadString(delim byte) (string, error)
//   Next(n int) []byte                     - Read next n bytes
//
// ACCESSING
//
//   Bytes() []byte       - Returns internal buffer (don't modify!)
//   String() string      - Returns as string
//   Len() int            - Current length
//   Cap() int            - Current capacity
//   Grow(n int)          - Ensure space for n more bytes
//   Reset()              - Clear and reuse
//   Truncate(n int)      - Keep only first n bytes
//
// BUFFER IS BOTH READER AND WRITER
//
// Buffer implements io.Reader, io.Writer, io.ByteReader, io.ByteWriter, etc.
// Perfect for testing or when you need both reading and writing.
//
// =============================================================================

func DemonstrateBuffer() {
	// Creating
	var buf bytes.Buffer
	buf.WriteString("Hello, ")
	buf.WriteString("World!")
	fmt.Printf("Built: %s\n", buf.String())

	// Writing different types
	var buf2 bytes.Buffer
	buf2.Write([]byte("Bytes: "))
	buf2.WriteByte('X')
	buf2.WriteRune('世')
	buf2.WriteString(" done")
	fmt.Printf("Mixed: %s\n", buf2.String())

	// Reading
	var buf3 bytes.Buffer
	buf3.WriteString("line1\nline2\nline3")
	
	for {
		line, err := buf3.ReadString('\n')
		if err != nil {
			if err == io.EOF && line != "" {
				fmt.Printf("Last line: %s\n", line)
			}
			break
		}
		fmt.Printf("Read line: %s", line)
	}

	// Bytes() vs String()
	var buf4 bytes.Buffer
	buf4.WriteString("data")
	
	b := buf4.Bytes()   // Returns internal slice (fast, but don't modify!)
	s := buf4.String()  // Allocates new string
	fmt.Printf("Bytes: %v, String: %s\n", b, s)

	// Grow for performance
	var buf5 bytes.Buffer
	buf5.Grow(1000) // Pre-allocate for 1000 bytes
	for i := 0; i < 100; i++ {
		fmt.Fprintf(&buf5, "Line %d\n", i)
	}
	fmt.Printf("Grew buffer, length: %d\n", buf5.Len())

	// Reset for reuse
	buf5.Reset()
	buf5.WriteString("Reused")
	fmt.Printf("After reset: %s\n", buf5.String())

	// Truncate
	var buf6 bytes.Buffer
	buf6.WriteString("Hello, World!")
	buf6.Truncate(5) // Keep only "Hello"
	fmt.Printf("Truncated: %s\n", buf6.String())
}

// =============================================================================
// SECTION 11: bytes.Reader
// =============================================================================
//
// bytes.Reader provides io.Reader interface for []byte.
// Like strings.Reader but for byte slices.
//
// CREATING
//
//   r := bytes.NewReader(b []byte)
//
// METHODS (implements many interfaces)
//
//   Read(p []byte) (n int, err error)
//   ReadAt(p []byte, off int64) (n int, err error)
//   ReadByte() (byte, error)
//   ReadRune() (r rune, size int, err error)
//   UnreadByte() error
//   UnreadRune() error
//   Seek(offset int64, whence int) (int64, error)
//   Size() int64
//   Reset(b []byte)
//   Len() int                  - Unread portion
//   WriteTo(w io.Writer) (n int64, err error)
//
// USE CASES
//
// - Testing code that expects io.Reader
// - Passing []byte to APIs that need io.Reader
// - Re-reading same data multiple times
//
// =============================================================================

func DemonstrateReader() {
	data := []byte("Hello, World!")
	r := bytes.NewReader(data)

	// Basic reading
	buf := make([]byte, 5)
	n, _ := r.Read(buf)
	fmt.Printf("Read %d bytes: %s\n", n, buf)

	// ReadByte
	b, _ := r.ReadByte()
	fmt.Printf("Next byte: %c\n", b)

	// Seek
	r.Seek(0, io.SeekStart)
	buf2 := make([]byte, 5)
	r.Read(buf2)
	fmt.Printf("After seek: %s\n", buf2)

	// Size and Len
	fmt.Printf("Size: %d, Unread: %d\n", r.Size(), r.Len())

	// Reset for reuse
	r.Reset([]byte("New content"))
	buf3 := make([]byte, 3)
	r.Read(buf3)
	fmt.Printf("After reset: %s\n", buf3)

	// WriteTo - efficient copy to writer
	r.Reset([]byte("Copy this"))
	var buf4 bytes.Buffer
	r.WriteTo(&buf4)
	fmt.Printf("Copied: %s\n", buf4.String())
}

// =============================================================================
// SECTION 12: COMMON PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Efficient byte accumulation
	buildData := func() []byte {
		var buf bytes.Buffer
		buf.Grow(1024) // Pre-allocate if size known
		
		for i := 0; i < 10; i++ {
			buf.WriteString("item ")
			fmt.Fprintf(&buf, "%d\n", i)
		}
		
		return buf.Bytes()
	}
	data := buildData()
	fmt.Printf("Built %d bytes\n", len(data))

	// Pattern 2: Read-write buffer for testing
	testBuffer := func() {
		var buf bytes.Buffer
		
		// Write
		buf.WriteString("test data")
		
		// Read it back
		result, _ := io.ReadAll(&buf)
		fmt.Printf("Test read: %s\n", result)
	}
	testBuffer()

	// Pattern 3: Modify bytes in place (when safe)
	modifyInPlace := func(b []byte) {
		// Only safe if you own the slice!
		for i := range b {
			if b[i] >= 'a' && b[i] <= 'z' {
				b[i] -= 32 // ASCII uppercase
			}
		}
	}
	data2 := []byte("hello")
	modifyInPlace(data2)
	fmt.Printf("Modified: %s\n", data2)

	// Pattern 4: Zero-copy split (careful!)
	zeroCopySplit := func(b []byte, sep byte) [][]byte {
		// Result shares memory with b
		var parts [][]byte
		start := 0
		for i, c := range b {
			if c == sep {
				parts = append(parts, b[start:i])
				start = i + 1
			}
		}
		parts = append(parts, b[start:])
		return parts
	}
	parts := zeroCopySplit([]byte("a,b,c"), ',')
	fmt.Printf("Zero-copy split: %s\n", parts)

	// Pattern 5: Safe copy for long-lived data
	safeCopy := func(b []byte) []byte {
		// Use bytes.Clone or make+copy
		return bytes.Clone(b)
	}
	original := []byte("important data")
	safe := safeCopy(original)
	safe[0] = 'X'
	fmt.Printf("Original: %s, Copy: %s\n", original, safe)

	// Pattern 6: Build protocol message
	buildMessage := func(msgType byte, payload []byte) []byte {
		var buf bytes.Buffer
		buf.WriteByte(msgType)
		buf.Write(payload)
		return buf.Bytes()
	}
	msg := buildMessage(0x01, []byte("data"))
	fmt.Printf("Message: %v\n", msg)

	// Pattern 7: Parse fixed-width fields
	parseFixed := func(data []byte, widths []int) [][]byte {
		var fields [][]byte
		offset := 0
		for _, w := range widths {
			if offset+w > len(data) {
				break
			}
			fields = append(fields, data[offset:offset+w])
			offset += w
		}
		return fields
	}
	fixed := []byte("John      Smith     ")
	fields := parseFixed(fixed, []int{10, 10})
	for i, f := range fields {
		fmt.Printf("Field %d: '%s'\n", i, bytes.TrimSpace(f))
	}
}

// =============================================================================
// SECTION 13: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Using == to compare []byte
	_ = `
	// WRONG - compares references
	b1 := []byte("hello")
	b2 := []byte("hello")
	if b1 == b2 { ... }  // COMPILE ERROR

	// RIGHT
	if bytes.Equal(b1, b2) { ... }
	`
	fmt.Println("Mistake 1: Use bytes.Equal, not ==")

	// Mistake 2: Modifying buf.Bytes() result
	_ = `
	// WRONG - modifies internal buffer
	var buf bytes.Buffer
	buf.WriteString("hello")
	b := buf.Bytes()
	b[0] = 'X'  // Corrupts buffer!

	// RIGHT - copy if you need to modify
	b := bytes.Clone(buf.Bytes())
	b[0] = 'X'
	`
	fmt.Println("Mistake 2: Don't modify buf.Bytes() directly")

	// Mistake 3: Forgetting Split shares memory
	_ = `
	// DANGEROUS
	original := []byte("a,b,c")
	parts := bytes.Split(original, []byte(","))
	parts[0][0] = 'X'  // Modifies original!

	// SAFE
	parts := bytes.Split(original, []byte(","))
	safePart := bytes.Clone(parts[0])
	safePart[0] = 'X'
	`
	fmt.Println("Mistake 3: Split shares memory with original")

	// Mistake 4: Inefficient concatenation
	_ = `
	// WRONG - allocates many times
	var result []byte
	for i := 0; i < 1000; i++ {
		result = append(result, []byte("x")...)
	}

	// RIGHT - use Buffer
	var buf bytes.Buffer
	for i := 0; i < 1000; i++ {
		buf.WriteByte('x')
	}
	result := buf.Bytes()
	`
	fmt.Println("Mistake 4: Use Buffer for building, not append")

	// Mistake 5: Not checking buffer operations
	_ = `
	// WRONG - ignoring errors
	var buf bytes.Buffer
	buf.ReadString('\n')  // Could return error

	// RIGHT
	line, err := buf.ReadString('\n')
	if err != nil {
		// Handle
	}
	`
	fmt.Println("Mistake 5: Check errors even with Buffer")

	// Mistake 6: Unnecessary conversions
	_ = `
	// SLOW - converts back and forth
	b := []byte(strings.ToUpper(string(data)))

	// FAST - work in bytes
	b := bytes.ToUpper(data)
	`
	fmt.Println("Mistake 6: Avoid string/[]byte conversions")

	// Mistake 7: Slice expressions with wrong bounds
	b := []byte("hello")
	// b[0:100]  // PANIC - out of bounds
	// Use len(b) checks or TrimSpace/etc
	fmt.Printf("Mistake 7: Check bounds - len is %d\n", len(b))
}

// =============================================================================
// SECTION 14: BYTES VS STRINGS DECISION GUIDE
// =============================================================================
//
// USE STRING WHEN:
//
// - Text won't be modified
// - Function parameters (strings pass by value efficiently)
// - Map keys (strings comparable, []byte not)
// - Sharing across goroutines (immutable = safe)
// - Most APIs expect string
//
// USE []BYTE WHEN:
//
// - Need to modify contents
// - Building/accumulating data (bytes.Buffer)
// - I/O operations (io.Reader/Writer use []byte)
// - Binary data / protocol parsing
// - Performance-critical (avoid string conversion overhead)
// - Working with crypto, hashing (they use []byte)
//
// CONVERSION COST
//
//   string(b)  - Allocates and copies
//   []byte(s)  - Allocates and copies
//
// In hot paths, stay in one type to avoid conversions.
//
// =============================================================================

func DemonstrateDecisionGuide() {
	// Example: Process configuration file

	// Scenario 1: Just parsing and storing - use string
	config := map[string]string{
		"host": "localhost",
		"port": "8080",
	}
	fmt.Printf("Config: %v\n", config)

	// Scenario 2: Building output - use bytes
	var output bytes.Buffer
	for k, v := range config {
		fmt.Fprintf(&output, "%s=%s\n", k, v)
	}
	fmt.Printf("Output:\n%s", output.String())

	// Scenario 3: Protocol message - use bytes
	buildFrame := func(data []byte) []byte {
		var frame bytes.Buffer
		frame.WriteByte(0xFF) // Header
		frame.WriteByte(byte(len(data)))
		frame.Write(data)
		frame.WriteByte(0xFE) // Footer
		return frame.Bytes()
	}
	frame := buildFrame([]byte("payload"))
	fmt.Printf("Frame: %v\n", frame)
}

// =============================================================================
// SECTION 15: PERFORMANCE NOTES
// =============================================================================
//
// BENCHMARKING INSIGHTS
//
// 1. bytes.Buffer vs append
//    Buffer is faster for 100+ appends
//    Pre-allocating with Grow() helps significantly
//
// 2. bytes.Equal vs manual comparison
//    bytes.Equal is optimized, use it
//
// 3. String/[]byte conversions
//    Each conversion allocates - minimize them
//
// 4. Split vs manual parsing
//    Split allocates [][]byte
//    Manual parsing can be faster for simple cases
//
// 5. Clone vs copy
//    bytes.Clone is optimized for the common case
//
// 6. In-place modifications
//    Only worthwhile for very hot paths
//    Complexity often not worth it
//
// =============================================================================

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateComparison()
	DemonstrateSearching()
	DemonstratePrefixSuffix()
	DemonstrateSplitJoin()
	DemonstrateTrim()
	DemonstrateCase()
	DemonstrateReplace()
	DemonstrateRunes()
	DemonstrateBuffer()
	DemonstrateReader()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
	DemonstrateDecisionGuide()
}
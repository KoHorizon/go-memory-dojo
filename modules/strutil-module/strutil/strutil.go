// Package strutil provides comprehensive documentation and working examples
// for Go's strings package - essential string manipulation functions.
package strutil

import (
	"fmt"
	"strings"
	"unicode"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF THE STRINGS PACKAGE
// =============================================================================
//
// WHY DOES THE STRINGS PACKAGE EXIST?
//
// In Go, strings are immutable sequences of bytes (typically UTF-8).
// Once created, a string's content cannot be changed. This immutability
// is fundamental to Go's design but creates a challenge:
//
// How do you manipulate strings efficiently?
//
// The strings package is Go's answer. It provides functions that:
// 1. Search within strings
// 2. Compare strings
// 3. Transform strings (creating new ones)
// 4. Split and join strings
// 5. Build strings efficiently (strings.Builder)
//
// STRINGS ARE IMMUTABLE
//
// Every operation that "modifies" a string actually creates a NEW string:
//
//   s := "hello"
//   s = strings.ToUpper(s)  // Creates new string "HELLO"
//
// The original "hello" still exists somewhere in memory (until GC'd).
// This is why strings.Builder exists - to avoid creating many intermediate
// strings when building a final result.
//
// UTF-8 AWARENESS
//
// Most string functions work on bytes, not runes (Unicode code points).
// But many have rune-aware variants or work correctly with UTF-8:
//
//   strings.Contains("hello", "ll")      // Byte search
//   strings.ContainsRune("hello", 'l')   // Rune search
//
// =============================================================================

// =============================================================================
// SECTION 2: SEARCHING - CONTAINS FAMILY
// =============================================================================
//
// Contains(s, substr string) bool
//   Does s contain substr?
//   Byte-level search, works with UTF-8
//
// ContainsAny(s, chars string) bool
//   Does s contain any rune from chars?
//   Useful: ContainsAny(s, "aeiou") for vowels
//
// ContainsRune(s string, r rune) bool
//   Does s contain the rune r?
//   Same as strings.IndexRune(s, r) >= 0
//
// ContainsFunc(s string, f func(rune) bool) bool
//   Does s contain any rune satisfying f?
//   Powerful: ContainsFunc(s, unicode.IsDigit)
//
// Count(s, substr string) int
//   How many non-overlapping instances of substr in s?
//   Empty substr counts rune count + 1
//
// =============================================================================

func DemonstrateContains() {
	text := "The quick brown fox jumps over the lazy dog"

	// Basic contains
	fmt.Printf("Contains 'fox': %v\n", strings.Contains(text, "fox"))
	fmt.Printf("Contains 'cat': %v\n", strings.Contains(text, "cat"))

	// ContainsAny - any character match
	fmt.Printf("Contains vowel: %v\n", strings.ContainsAny(text, "aeiou"))
	fmt.Printf("Contains digit: %v\n", strings.ContainsAny(text, "0123456789"))

	// ContainsRune
	fmt.Printf("Contains 'x': %v\n", strings.ContainsRune(text, 'x'))
	fmt.Printf("Contains '世': %v\n", strings.ContainsRune(text, '世'))

	// ContainsFunc - custom predicate
	hasDigit := strings.ContainsFunc(text, unicode.IsDigit)
	fmt.Printf("Has digit: %v\n", hasDigit)

	hasUppercase := strings.ContainsFunc(text, unicode.IsUpper)
	fmt.Printf("Has uppercase: %v\n", hasUppercase)

	// Count occurrences
	fmt.Printf("Count of 'o': %d\n", strings.Count(text, "o"))
	fmt.Printf("Count of 'the': %d\n", strings.Count(text, "the")) // Case-sensitive!
	fmt.Printf("Count of 'The': %d\n", strings.Count(text, "The"))

	// Edge cases
	fmt.Printf("Count of '': %d\n", strings.Count("hello", "")) // Rune count + 1 = 6
}

// =============================================================================
// SECTION 3: SEARCHING - INDEX FAMILY
// =============================================================================
//
// FINDING POSITIONS
//
// Index(s, substr string) int
//   Position of first occurrence of substr in s, or -1
//
// LastIndex(s, substr string) int
//   Position of last occurrence of substr in s, or -1
//
// IndexByte(s string, c byte) int
//   Position of first occurrence of byte c, or -1
//   Faster than Index for single bytes
//
// IndexRune(s string, r rune) int
//   Position (byte offset!) of first occurrence of rune r
//
// IndexAny(s, chars string) int
//   Position of first occurrence of ANY rune from chars
//
// LastIndexAny(s, chars string) int
//   Position of last occurrence of ANY rune from chars
//
// IndexFunc(s string, f func(rune) bool) int
//   Position of first rune satisfying f
//
// LastIndexFunc(s string, f func(rune) bool) int
//   Position of last rune satisfying f
//
// CRITICAL: These return BYTE positions, not rune positions!
// In UTF-8, one rune can be multiple bytes.
//
// =============================================================================

func DemonstrateIndex() {
	text := "Hello, 世界! Welcome to Go!"

	// Basic index
	fmt.Printf("Index of 'o': %d\n", strings.Index(text, "o"))
	fmt.Printf("LastIndex of 'o': %d\n", strings.LastIndex(text, "o"))

	// Not found returns -1
	fmt.Printf("Index of 'xyz': %d\n", strings.Index(text, "xyz"))

	// IndexByte - for single bytes (faster)
	fmt.Printf("IndexByte of 'o': %d\n", strings.IndexByte(text, 'o'))

	// IndexRune - for Unicode characters
	fmt.Printf("IndexRune of '世': %d\n", strings.IndexRune(text, '世'))

	// IndexAny - first occurrence of any character
	fmt.Printf("First vowel at: %d\n", strings.IndexAny(text, "aeiou"))
	fmt.Printf("Last vowel at: %d\n", strings.LastIndexAny(text, "aeiou"))

	// IndexFunc - custom predicate
	firstDigit := strings.IndexFunc("abc123def", unicode.IsDigit)
	fmt.Printf("First digit at: %d\n", firstDigit)

	firstUpper := strings.IndexFunc(text, unicode.IsUpper)
	fmt.Printf("First uppercase at: %d\n", firstUpper)

	// Example: Find all occurrences
	substr := "o"
	s := "Hello World"
	positions := []int{}
	for i := 0; i < len(s); {
		pos := strings.Index(s[i:], substr)
		if pos == -1 {
			break
		}
		actualPos := i + pos
		positions = append(positions, actualPos)
		i = actualPos + len(substr)
	}
	fmt.Printf("All 'o' positions in '%s': %v\n", s, positions)
}

// =============================================================================
// SECTION 4: PREFIX AND SUFFIX
// =============================================================================
//
// HasPrefix(s, prefix string) bool
//   Does s start with prefix?
//   Common: checking file extensions, URL schemes
//
// HasSuffix(s, suffix string) bool
//   Does s end with suffix?
//   Common: filtering files, checking domains
//
// TrimPrefix(s, prefix string) string
//   Remove prefix if present (once)
//   If not present, returns s unchanged
//
// TrimSuffix(s, suffix string) string
//   Remove suffix if present (once)
//   If not present, returns s unchanged
//
// CutPrefix(s, prefix string) (after string, found bool)
//   Go 1.20+: Like TrimPrefix but tells you if it was found
//
// CutSuffix(s, suffix string) (before string, found bool)
//   Go 1.20+: Like TrimSuffix but tells you if it was found
//
// =============================================================================

func DemonstratePrefixSuffix() {
	filename := "document.txt"
	url := "https://example.com/api/v1/users"

	// HasPrefix/HasSuffix
	fmt.Printf("Is text file: %v\n", strings.HasSuffix(filename, ".txt"))
	fmt.Printf("Is HTTPS: %v\n", strings.HasPrefix(url, "https://"))
	fmt.Printf("Is API path: %v\n", strings.HasPrefix(url, "https://example.com/api/"))

	// TrimPrefix/TrimSuffix
	noExt := strings.TrimSuffix(filename, ".txt")
	fmt.Printf("Without extension: %s\n", noExt)

	noScheme := strings.TrimPrefix(url, "https://")
	fmt.Printf("Without scheme: %s\n", noScheme)

	// Only removes if present
	fmt.Printf("Trim absent prefix: %s\n", strings.TrimPrefix("hello", "world"))

	// Go 1.20+ CutPrefix/CutSuffix - tells you if found
	if name, ok := strings.CutSuffix(filename, ".txt"); ok {
		fmt.Printf("Filename without .txt: %s\n", name)
	}

	if path, ok := strings.CutPrefix(url, "https://example.com"); ok {
		fmt.Printf("Path only: %s\n", path)
	}

	// Pattern: Handle multiple possible extensions
	exts := []string{".txt", ".md", ".doc"}
	name := "readme.md"
	for _, ext := range exts {
		if strings.HasSuffix(name, ext) {
			base := strings.TrimSuffix(name, ext)
			fmt.Printf("Base name: %s, Extension: %s\n", base, ext)
			break
		}
	}
}

// =============================================================================
// SECTION 5: SPLITTING AND JOINING
// =============================================================================
//
// Split(s, sep string) []string
//   Split on every occurrence of sep
//   Empty sep splits into individual runes
//   Result ALWAYS has at least one element
//
// SplitN(s, sep string, n int) []string
//   Split, but at most n substrings
//   Last substring is the "rest" (unsplit)
//   n < 0 means no limit (same as Split)
//
// SplitAfter(s, sep string) []string
//   Like Split, but keeps separator at end of each piece
//
// SplitAfterN(s, sep string, n int) []string
//   SplitAfter with limit
//
// Fields(s string) []string
//   Split on whitespace, discard empty strings
//   Whitespace: space, tab, newline, etc.
//   Most common for parsing space-separated input
//
// FieldsFunc(s string, f func(rune) bool) []string
//   Split on runes satisfying f
//
// Join(elems []string, sep string) string
//   Join strings with separator
//   Inverse of Split
//
// =============================================================================

func DemonstrateSplitJoin() {
	// Basic split
	csv := "apple,banana,cherry"
	fruits := strings.Split(csv, ",")
	fmt.Printf("Fruits: %v\n", fruits)

	// Split with limit
	path := "a/b/c/d/e"
	parts := strings.SplitN(path, "/", 3)
	fmt.Printf("Split with limit 3: %v\n", parts) // [a b c/d/e]

	// SplitAfter - keeps separator
	lines := strings.SplitAfter("line1\nline2\nline3", "\n")
	fmt.Printf("Lines with \\n: %v\n", lines)

	// Empty separator splits into runes
	chars := strings.Split("hello", "")
	fmt.Printf("Characters: %v\n", chars)

	// Fields - split on whitespace
	text := "  one   two\tthree\nfour  "
	words := strings.Fields(text)
	fmt.Printf("Words: %v\n", words) // [one two three four]

	// FieldsFunc - custom splitter
	data := "apple,banana;cherry:date"
	parts2 := strings.FieldsFunc(data, func(r rune) bool {
		return r == ',' || r == ';' || r == ':'
	})
	fmt.Printf("Multi-delimiter split: %v\n", parts2)

	// Join - inverse of split
	joined := strings.Join(fruits, " | ")
	fmt.Printf("Joined: %s\n", joined)

	// Edge cases
	fmt.Printf("Split empty: %v\n", strings.Split("", ","))     // [""]
	fmt.Printf("Fields empty: %v\n", strings.Fields(""))        // []
	fmt.Printf("Fields spaces: %v\n", strings.Fields("   "))    // []
	fmt.Printf("Join empty: %q\n", strings.Join([]string{}, ",")) // ""
}

// =============================================================================
// SECTION 6: TRIMMING
// =============================================================================
//
// REMOVING CHARACTERS FROM ENDS
//
// Trim(s, cutset string) string
//   Remove any runes in cutset from both ends
//
// TrimLeft(s, cutset string) string
//   Remove from left/start only
//
// TrimRight(s, cutset string) string
//   Remove from right/end only
//
// TrimSpace(s string) string
//   Remove whitespace from both ends
//   Most commonly used trim function
//
// TrimFunc(s string, f func(rune) bool) string
//   Remove runes satisfying f from both ends
//
// TrimLeftFunc, TrimRightFunc
//   Trim from one end with custom function
//
// CRITICAL: These trim ANY runes in cutset, not a substring!
//   Trim("hello", "lo") removes 'l' and 'o' chars, not the substring "lo"
//
// =============================================================================

func DemonstrateTrim() {
	// TrimSpace - most common
	padded := "   hello world   \n\t"
	fmt.Printf("Trimmed: '%s'\n", strings.TrimSpace(padded))

	// Trim with cutset
	bracketed := "[[hello]]"
	fmt.Printf("Trim brackets: %s\n", strings.Trim(bracketed, "[]"))

	// Trim removes ANY of the runes
	messy := "!!!hello???"
	fmt.Printf("Trim punctuation: %s\n", strings.Trim(messy, "!?"))

	// TrimLeft and TrimRight
	fmt.Printf("TrimLeft: %s\n", strings.TrimLeft(messy, "!"))
	fmt.Printf("TrimRight: %s\n", strings.TrimRight(messy, "?"))

	// TrimFunc - custom predicate
	text := "123hello456"
	noDigits := strings.TrimFunc(text, unicode.IsDigit)
	fmt.Printf("Trim digits: %s\n", noDigits)

	// Common mistake: Trim is not Remove
	s := "hello"
	fmt.Printf("Trim 'lo': %s\n", strings.Trim(s, "lo")) // "he" (removes l, o chars)
	fmt.Printf("TrimSuffix 'lo': %s\n", strings.TrimSuffix(s, "lo")) // "hel" (removes substring)

	// Pattern: Normalize input
	userInput := "  \t user@example.com \n "
	email := strings.TrimSpace(strings.ToLower(userInput))
	fmt.Printf("Normalized email: %s\n", email)
}

// =============================================================================
// SECTION 7: CASE CONVERSION
// =============================================================================
//
// ToUpper(s string) string
//   Convert to uppercase
//   Unicode-aware: "hello" → "HELLO", "ß" → "SS"
//
// ToLower(s string) string
//   Convert to lowercase
//   Unicode-aware
//
// ToTitle(s string) string
//   Convert to title case (similar to ToUpper)
//   Difference matters for some Unicode scripts
//
// Title(s string) string (DEPRECATED in Go 1.18)
//   Convert to title case (capitalize first letter of each word)
//   Use cases.Title from golang.org/x/text/cases instead
//
// EqualFold(s, t string) bool
//   Case-insensitive comparison
//   More efficient than strings.ToLower(s) == strings.ToLower(t)
//
// ToUpperSpecial(c unicode.SpecialCase, s string) string
//   Case conversion with special rules (Turkish, etc.)
//
// ToLowerSpecial, ToTitleSpecial
//
// =============================================================================

func DemonstrateCase() {
	text := "Hello, World!"

	// Basic conversions
	fmt.Printf("Upper: %s\n", strings.ToUpper(text))
	fmt.Printf("Lower: %s\n", strings.ToLower(text))

	// Unicode awareness
	german := "Straße"
	fmt.Printf("Upper (ß→SS): %s\n", strings.ToUpper(german))

	greek := "Ελληνικά"
	fmt.Printf("Upper Greek: %s\n", strings.ToUpper(greek))

	// Case-insensitive comparison
	fmt.Printf("'hello' == 'HELLO': %v\n", "hello" == "HELLO")
	fmt.Printf("EqualFold: %v\n", strings.EqualFold("hello", "HELLO"))
	fmt.Printf("EqualFold: %v\n", strings.EqualFold("Straße", "STRASSE"))

	// Pattern: Case-insensitive search
	haystack := "The Quick Brown Fox"
	needle := "quick"
	found := strings.Contains(strings.ToLower(haystack), strings.ToLower(needle))
	fmt.Printf("Case-insensitive contains: %v\n", found)

	// Better: Use EqualFold for exact match
	words := strings.Fields(haystack)
	for _, word := range words {
		if strings.EqualFold(word, needle) {
			fmt.Printf("Found word: %s\n", word)
		}
	}
}

// =============================================================================
// SECTION 8: REPLACE
// =============================================================================
//
// Replace(s, old, new string, n int) string
//   Replace first n occurrences of old with new
//   n < 0 means replace all
//
// ReplaceAll(s, old, new string) string
//   Replace all occurrences
//   Same as Replace(s, old, new, -1)
//   More readable for "replace all" intent
//
// Map(mapping func(rune) rune, s string) string
//   Map each rune through function
//   Return negative value to drop the rune
//   Powerful for custom transformations
//
// =============================================================================

func DemonstrateReplace() {
	text := "foo bar foo baz foo"

	// Replace first N
	fmt.Printf("Replace 1: %s\n", strings.Replace(text, "foo", "XXX", 1))
	fmt.Printf("Replace 2: %s\n", strings.Replace(text, "foo", "XXX", 2))

	// Replace all
	fmt.Printf("Replace all: %s\n", strings.Replace(text, "foo", "XXX", -1))
	fmt.Printf("ReplaceAll: %s\n", strings.ReplaceAll(text, "foo", "XXX"))

	// Empty string replacement
	fmt.Printf("Replace 'o': %s\n", strings.ReplaceAll("hello", "o", ""))

	// Map function - custom transformation
	rot13 := func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return 'a' + (r-'a'+13)%26
		}
		if r >= 'A' && r <= 'Z' {
			return 'A' + (r-'A'+13)%26
		}
		return r
	}
	encoded := strings.Map(rot13, "Hello, World!")
	fmt.Printf("ROT13: %s\n", encoded)
	decoded := strings.Map(rot13, encoded)
	fmt.Printf("Decoded: %s\n", decoded)

	// Map to remove characters
	onlyDigits := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1 // Drop this rune
	}, "abc123def456")
	fmt.Printf("Only digits: %s\n", onlyDigits)

	// Pattern: Multiple replacements
	s := "Hello {name}, welcome to {place}!"
	s = strings.ReplaceAll(s, "{name}", "Alice")
	s = strings.ReplaceAll(s, "{place}", "Go")
	fmt.Printf("Template: %s\n", s)
}

// =============================================================================
// SECTION 9: strings.Replacer - MULTIPLE REPLACEMENTS
// =============================================================================
//
// WHY REPLACER?
//
// When you need to replace multiple different strings at once:
//
//   // Inefficient - multiple passes
//   s = strings.ReplaceAll(s, "foo", "bar")
//   s = strings.ReplaceAll(s, "baz", "qux")
//   s = strings.ReplaceAll(s, "old", "new")
//
//   // Efficient - single pass
//   r := strings.NewReplacer("foo", "bar", "baz", "qux", "old", "new")
//   s = r.Replace(s)
//
// CREATING A REPLACER
//
//   r := strings.NewReplacer(oldnew ...string)
//
// Arguments are old, new, old, new, ... (pairs)
// Must be even number of arguments
//
// METHODS
//
//   r.Replace(s string) string
//     Replace all occurrences in s
//
//   r.WriteString(w io.Writer, s string) (n int, err error)
//     Write replaced version to writer (avoids intermediate string)
//
// REPLACEMENT ORDER
//
// Replacer processes from longest to shortest old string.
// Prevents "recursive" replacements:
//
//   r := NewReplacer("a", "ab", "ab", "c")
//   r.Replace("a")  // Returns "ab", not "c"
//
// PERFORMANCE
//
// - For 1-2 replacements: ReplaceAll is fine
// - For 3+ replacements: Replacer is faster
// - Replacer can be reused (create once, use many times)
//
// =============================================================================

func DemonstrateReplacer() {
	// Basic replacer
	r := strings.NewReplacer("foo", "FOO", "bar", "BAR", "baz", "BAZ")
	text := "foo and bar and baz"
	result := r.Replace(text)
	fmt.Printf("Replaced: %s\n", result)

	// Pattern: HTML escaping
	htmlEscaper := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		`"`, "&quot;",
		"'", "&#39;",
	)
	html := `<script>alert("XSS")</script>`
	escaped := htmlEscaper.Replace(html)
	fmt.Printf("Escaped: %s\n", escaped)

	// Pattern: Normalize line endings
	normalizeLines := strings.NewReplacer("\r\n", "\n", "\r", "\n")
	mixed := "line1\r\nline2\rline3\n"
	normalized := normalizeLines.Replace(mixed)
	fmt.Printf("Normalized: %q\n", normalized)

	// Pattern: Remove multiple unwanted characters
	cleaner := strings.NewReplacer(
		"\t", "",
		"\n", "",
		"\r", "",
		"  ", " ",
	)
	dirty := "hello\t\nworld  foo"
	clean := cleaner.Replace(dirty)
	fmt.Printf("Cleaned: %q\n", clean)

	// WriteString - avoid intermediate string
	var buf strings.Builder
	r2 := strings.NewReplacer("a", "A", "e", "E", "i", "I", "o", "O", "u", "U")
	r2.WriteString(&buf, "hello world")
	fmt.Printf("Written to buffer: %s\n", buf.String())

	// Replacer is reusable
	sanitizer := strings.NewReplacer(
		"password", "[REDACTED]",
		"secret", "[REDACTED]",
		"key", "[REDACTED]",
	)
	logs := []string{
		"user logged in with password abc123",
		"API key xyz789 used",
		"secret token found",
	}
	for _, log := range logs {
		fmt.Printf("Sanitized log: %s\n", sanitizer.Replace(log))
	}

	// Order matters: longest first
	r3 := strings.NewReplacer("aa", "X", "a", "Y")
	fmt.Printf("'aaa' -> %s\n", r3.Replace("aaa")) // "XY" not "YYY"
}

// =============================================================================
// SECTION 10: strings.Clone
// =============================================================================
//
// Clone(s string) string
//   Returns a fresh copy of s
//
// WHY DOES THIS EXIST?
//
// Strings in Go are immutable, but they can share underlying memory.
// When you slice a large string, the slice keeps the entire original in memory:
//
//   large := strings.Repeat("x", 1000000) // 1MB string
//   small := large[0:10]                  // 10 bytes, but keeps 1MB!
//
// Clone creates a new string with its own memory:
//
//   small := strings.Clone(large[0:10])   // Only 10 bytes allocated
//   large = ""                             // Original can be GC'd
//
// WHEN TO USE
//
// - After slicing a large string, keeping only a small part
// - When sharing strings between goroutines with different lifetimes
// - Breaking reference to large parent string
//
// WHEN NOT TO USE
//
// - Regular string operations (unnecessary overhead)
// - Strings that aren't slices of larger strings
// - Short-lived strings
//
// NOTE: Added in Go 1.18
//
// =============================================================================

func DemonstrateClone() {
	// Problem: Slicing keeps original in memory
	largeFile := strings.Repeat("x", 1000000) + "\nfirst line\nsecond line"
	firstLine, _, _ := strings.Cut(largeFile, "\n")
	// firstLine still references the 1MB+ largeFile!

	// Solution: Clone breaks the reference
	firstLineCopy := strings.Clone(firstLine)
	largeFile = "" // Can now be garbage collected
	fmt.Printf("Cloned first line length: %d\n", len(firstLineCopy))

	// Pattern: Extract small data from large response
	processLogLine := func(hugelog string) string {
		// Extract just the timestamp (first 20 chars)
		if len(hugelog) < 20 {
			return hugelog
		}
		// Clone to release reference to huge log
		return strings.Clone(hugelog[0:20])
	}
	timestamp := processLogLine(strings.Repeat("x", 1000000) + "2024-01-01T12:00:00 message")
	fmt.Printf("Timestamp: %s\n", timestamp)

	// Example: Parse CSV with large fields, keep only IDs
	csvLine := "00001," + strings.Repeat("x", 100000) + "," + strings.Repeat("y", 100000)
	fields := strings.Split(csvLine, ",")
	// Only want the ID (first field), but Split keeps entire csvLine in memory
	id := strings.Clone(fields[0])
	fields = nil // Allow GC
	csvLine = ""
	fmt.Printf("ID: %s\n", id)

	// Edge case: Clone is a no-op for string literals
	literal := "hello"
	cloned := strings.Clone(literal)
	// Both point to same memory (compiler optimization)
	fmt.Printf("Clone of literal: %s\n", cloned)
}

// =============================================================================
// SECTION 11: REPEAT AND COMPARE
// =============================================================================
//
// Repeat(s string, count int) string
//   Repeat s count times
//   Efficient - doesn't create intermediate strings
//
// Compare(a, b string) int
//   Returns: < 0 if a < b, 0 if a == b, > 0 if a > b
//   Lexicographic byte comparison
//   Usually just use == or < directly
//
// =============================================================================

func DemonstrateRepeatCompare() {
	// Repeat
	fmt.Printf("Repeat: %s\n", strings.Repeat("Go", 5))
	fmt.Printf("Separator: %s\n", strings.Repeat("-", 40))

	// Use case: padding
	text := "Title"
	padding := strings.Repeat(" ", (40-len(text))/2)
	fmt.Printf("%s%s%s\n", padding, text, padding)

	// Compare (rarely needed - use == or < instead)
	fmt.Printf("Compare 'a' vs 'b': %d\n", strings.Compare("a", "b")) // -1
	fmt.Printf("Compare 'b' vs 'a': %d\n", strings.Compare("b", "a")) // 1
	fmt.Printf("Compare 'a' vs 'a': %d\n", strings.Compare("a", "a")) // 0

	// Usually better to use operators
	fmt.Printf("'a' == 'a': %v\n", "a" == "a")
	fmt.Printf("'a' < 'b': %v\n", "a" < "b")
}

// =============================================================================
// SECTION 12: CUT (Go 1.18+)
// =============================================================================
//
// Cut(s, sep string) (before, after string, found bool)
//   Split s around first occurrence of sep
//   Returns: text before sep, text after sep, whether sep was found
//
// Why Cut?
//   - More efficient than Index + slicing
//   - Cleaner than handling Index returning -1
//   - One allocation instead of multiple
//
// CRITICAL: Cut finds the FIRST occurrence, not last
//
// =============================================================================

func DemonstrateCut() {
	// Basic cut
	email := "user@example.com"
	user, domain, found := strings.Cut(email, "@")
	if found {
		fmt.Printf("User: %s, Domain: %s\n", user, domain)
	}

	// Not found
	_, _, found = strings.Cut("no-separator", "@")
	fmt.Printf("Found separator: %v\n", found)

	// Pattern: Parse key=value
	config := "timeout=30"
	key, value, ok := strings.Cut(config, "=")
	if ok {
		fmt.Printf("Config: %s = %s\n", key, value)
	}

	// Pattern: Split URL path and query
	url := "/api/users?page=2&limit=10"
	path, query, _ := strings.Cut(url, "?")
	fmt.Printf("Path: %s, Query: %s\n", path, query)

	// Before Go 1.18, had to do:
	// i := strings.Index(url, "?")
	// if i < 0 {
	//     path, query = url, ""
	// } else {
	//     path, query = url[:i], url[i+1:]
	// }

	// CutPrefix and CutSuffix shown in Section 4
}

// =============================================================================
// SECTION 13: strings.Builder - EFFICIENT STRING BUILDING
// =============================================================================
//
// WHY STRINGS.BUILDER?
//
// Strings are immutable. Concatenating with + creates many intermediate strings:
//
//   s := ""
//   for i := 0; i < 1000; i++ {
//       s += "x"  // Creates 1000 strings! Slow and wasteful
//   }
//
// strings.Builder solves this:
//
//   var b strings.Builder
//   for i := 0; i < 1000; i++ {
//       b.WriteString("x")  // One final string
//   }
//   s := b.String()
//
// METHODS
//
//   WriteString(s string) (int, error)  - Most common
//   WriteByte(c byte) error
//   WriteRune(r rune) (int, error)
//   Write(p []byte) (int, error)        - Implements io.Writer
//   Grow(n int)                          - Pre-allocate capacity
//   String() string                      - Get result
//   Reset()                              - Clear and reuse
//   Len() int                            - Current length
//   Cap() int                            - Current capacity
//
// PERFORMANCE TIPS
//
// - Call Grow() if you know approximate size
// - Reuse Builder with Reset()
// - Don't call String() until done (it's not free)
//
// =============================================================================

func DemonstrateBuilder() {
	// Basic usage
	var b strings.Builder
	b.WriteString("Hello, ")
	b.WriteString("World")
	b.WriteByte('!')
	result := b.String()
	fmt.Printf("Built: %s\n", result)

	// Grow for better performance
	var b2 strings.Builder
	b2.Grow(100) // Pre-allocate for ~100 bytes
	for i := 0; i < 10; i++ {
		fmt.Fprintf(&b2, "Line %d\n", i)
	}
	fmt.Printf("Built with Grow:\n%s", b2.String())

	// Reusing builder
	var b3 strings.Builder
	for i := 0; i < 3; i++ {
		b3.Reset()
		fmt.Fprintf(&b3, "Iteration %d", i)
		fmt.Printf("Reused: %s\n", b3.String())
	}

	// Pattern: Building CSV
	fields := []string{"Alice", "30", "Engineer"}
	var csv strings.Builder
	for i, field := range fields {
		if i > 0 {
			csv.WriteString(",")
		}
		csv.WriteString(field)
	}
	fmt.Printf("CSV: %s\n", csv.String())

	// Pattern: Building HTML
	var html strings.Builder
	html.WriteString("<html><body>\n")
	items := []string{"apple", "banana", "cherry"}
	html.WriteString("<ul>\n")
	for _, item := range items {
		html.WriteString("  <li>")
		html.WriteString(item)
		html.WriteString("</li>\n")
	}
	html.WriteString("</ul>\n</body></html>")
	fmt.Printf("HTML:\n%s\n", html.String())
}

// =============================================================================
// SECTION 14: strings.Reader - READING FROM STRINGS
// =============================================================================
//
// strings.Reader implements io.Reader, io.ReaderAt, io.Seeker, and more.
// Useful for passing strings to APIs that expect io.Reader.
//
// CREATION
//
//   r := strings.NewReader(s)
//
// METHODS (implements many interfaces)
//
//   Read(p []byte) (n int, err error)      - io.Reader
//   ReadAt(p []byte, off int64) (n, error) - io.ReaderAt
//   ReadByte() (byte, error)
//   ReadRune() (rune, int, error)
//   Seek(offset int64, whence int) (int64, error)
//   Size() int64                            - String length
//   Reset(s string)                         - Reuse with new string
//   Len() int                               - Unread portion
//
// COMMON USE CASES
//
// - Testing code that reads from io.Reader
// - Passing strings to HTTP request bodies
// - Feeding strings to JSON/XML decoders
//
// =============================================================================

func DemonstrateReader() {
	text := "Hello, World!"
	r := strings.NewReader(text)

	// Read like any io.Reader
	buf := make([]byte, 5)
	n, _ := r.Read(buf)
	fmt.Printf("Read %d bytes: %s\n", n, buf[:n])

	// ReadByte
	b, _ := r.ReadByte()
	fmt.Printf("Next byte: %c\n", b)

	// Seek back to start
	r.Seek(0, 0)
	buf2 := make([]byte, 5)
	r.Read(buf2)
	fmt.Printf("After seek: %s\n", buf2)

	// Size and Len
	r2 := strings.NewReader("test")
	fmt.Printf("Size: %d, Unread: %d\n", r2.Size(), r2.Len())
	r2.Read(make([]byte, 2))
	fmt.Printf("After read - Size: %d, Unread: %d\n", r2.Size(), r2.Len())

	// Reset for reuse
	r2.Reset("new content")
	fmt.Printf("After reset - Size: %d\n", r2.Size())
}

// =============================================================================
// SECTION 15: COMMON PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Normalize whitespace
	normalizeSpace := func(s string) string {
		return strings.Join(strings.Fields(s), " ")
	}
	messy := "  hello    world  \n  foo  "
	fmt.Printf("Normalized: '%s'\n", normalizeSpace(messy))

	// Pattern 2: Check if string is blank (empty or whitespace)
	isBlank := func(s string) bool {
		return strings.TrimSpace(s) == ""
	}
	fmt.Printf("'   ' is blank: %v\n", isBlank("   "))
	fmt.Printf("' x ' is blank: %v\n", isBlank(" x "))

	// Pattern 3: Truncate with ellipsis
	truncate := func(s string, maxLen int) string {
		if len(s) <= maxLen {
			return s
		}
		return s[:maxLen-3] + "..."
	}
	long := "This is a very long string"
	fmt.Printf("Truncated: %s\n", truncate(long, 15))

	// Pattern 4: Sentence case (capitalize first letter)
	sentenceCase := func(s string) string {
		if len(s) == 0 {
			return s
		}
		runes := []rune(s)
		runes[0] = unicode.ToUpper(runes[0])
		return string(runes)
	}
	fmt.Printf("Sentence case: %s\n", sentenceCase("hello world"))

	// Pattern 5: Count lines
	countLines := func(s string) int {
		if s == "" {
			return 0
		}
		return strings.Count(s, "\n") + 1
	}
	multiline := "line 1\nline 2\nline 3"
	fmt.Printf("Lines: %d\n", countLines(multiline))

	// Pattern 6: Split and keep empty strings
	// (Split includes empties, Fields doesn't)
	csv := "a,,c"
	parts := strings.Split(csv, ",")
	fmt.Printf("Split CSV: %v\n", parts) // [a  c]

	// Pattern 7: Join with "and" for last element
	items := []string{"apples", "oranges", "bananas"}
	var result string
	if len(items) == 0 {
		result = ""
	} else if len(items) == 1 {
		result = items[0]
	} else {
		result = strings.Join(items[:len(items)-1], ", ") + " and " + items[len(items)-1]
	}
	fmt.Printf("Natural join: %s\n", result)
}

// =============================================================================
// SECTION 16: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: String concatenation in loop
	_ = `
	// WRONG - creates many intermediate strings
	s := ""
	for i := 0; i < 1000; i++ {
		s += "x"  // Allocates 1000 times!
	}

	// RIGHT - use strings.Builder
	var b strings.Builder
	for i := 0; i < 1000; i++ {
		b.WriteString("x")
	}
	s := b.String()
	`
	fmt.Println("Mistake 1: Use Builder, not +=")

	// Mistake 2: Trim vs TrimSuffix
	_ = `
	// WRONG - Trim removes characters, not substring
	s := "hello"
	strings.Trim(s, "lo")  // Returns "he" (removes l and o)

	// RIGHT - TrimSuffix removes substring
	strings.TrimSuffix(s, "lo")  // Returns "hel"
	`
	fmt.Println("Mistake 2: Trim removes chars, not substring")

	// Mistake 3: Case-insensitive comparison
	_ = `
	// SLOW - allocates lowercase strings
	if strings.ToLower(a) == strings.ToLower(b) { ... }

	// FAST - no allocation
	if strings.EqualFold(a, b) { ... }
	`
	fmt.Println("Mistake 3: Use EqualFold for case-insensitive")

	// Mistake 4: Contains for prefix/suffix
	_ = `
	// LESS CLEAR
	if strings.Contains(s, "http://") && strings.Index(s, "http://") == 0 { ... }

	// CLEARER
	if strings.HasPrefix(s, "http://") { ... }
	`
	fmt.Println("Mistake 4: Use HasPrefix/HasSuffix for clarity")

	// Mistake 5: Fields vs Split
	_ = `
	s := "  a  b  c  "

	// Fields removes empty strings
	strings.Fields(s)  // ["a", "b", "c"]

	// Split keeps them
	strings.Split(s, " ")  // ["", "", "a", "", "b", "", "c", "", ""]
	`
	fmt.Println("Mistake 5: Fields vs Split behavior")

	// Mistake 6: Index returns -1, not error
	_ = `
	// WRONG - will panic if not found
	i := strings.Index(s, "x")
	before := s[:i]  // Panics if i == -1!

	// RIGHT - check first
	i := strings.Index(s, "x")
	if i >= 0 {
		before := s[:i]
	}
	`
	fmt.Println("Mistake 6: Check Index != -1")

	// Mistake 7: Split empty string
	parts := strings.Split("", ",")
	fmt.Printf("Mistake 7: Split('', ',') = %v (length %d)\n", parts, len(parts))
	// Returns [""], not []!
}

// =============================================================================
// SECTION 17: PERFORMANCE NOTES
// =============================================================================
//
// BENCHMARKING INSIGHTS
//
// 1. strings.Builder vs +=
//    Builder is 100x+ faster for 1000+ concatenations
//
// 2. EqualFold vs ToLower+==
//    EqualFold is ~2x faster (no allocations)
//
// 3. IndexByte vs Index for single byte
//    IndexByte is faster, use when possible
//
// 4. Contains vs Index
//    Contains is slightly clearer, same performance
//
// 5. Split allocates
//    Every call allocates a slice. Reuse when possible.
//
// 6. Builder.Grow()
//    Pre-allocating saves ~20-30% time for large builds
//
// When does it matter?
//   - Hot paths (called millions of times)
//   - Large strings (MB+)
//   - Tight loops
//
// For typical usage, clarity > micro-optimization.
//
// =============================================================================

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateContains()
	DemonstrateIndex()
	DemonstratePrefixSuffix()
	DemonstrateSplitJoin()
	DemonstrateTrim()
	DemonstrateCase()
	DemonstrateReplace()
	DemonstrateReplacer()
	DemonstrateClone()
	DemonstrateRepeatCompare()
	DemonstrateCut()
	DemonstrateBuilder()
	DemonstrateReader()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
}
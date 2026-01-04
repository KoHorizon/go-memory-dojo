// Package strutil_practice is your daily practice space for the strings package.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against strutil/strutil.go
// 3. Note what you missed - focus on those tomorrow
package strutil_practice

// =============================================================================
// EXERCISE 1: CONTAINS FAMILY
// =============================================================================

func PracticeContains() {
	// TODO: Check if string contains substring
	// found := strings.???(s, substr)

	// TODO: Check if string contains any character from a set
	// hasVowel := strings.???(s, "aeiou")

	// TODO: Check if string contains a specific rune
	// hasX := strings.???(s, 'x')

	// TODO: Check if string contains rune matching predicate
	// hasDigit := strings.???(s, unicode.IsDigit)

	// TODO: Count occurrences of substring
	// count := strings.???(s, substr)

	// TODO: What does Count return for empty substring?
	// Answer: ???
}

// =============================================================================
// EXERCISE 2: INDEX FAMILY
// =============================================================================

func PracticeIndex() {
	// TODO: Find first occurrence of substring
	// pos := strings.???(s, substr)

	// TODO: Find last occurrence of substring
	// pos := strings.???(s, substr)

	// TODO: Find first occurrence of byte (faster for single bytes)
	// pos := strings.???(s, 'x')

	// TODO: Find first occurrence of rune
	// pos := strings.???(s, 'x')

	// TODO: Find first occurrence of any character from set
	// pos := strings.???(s, "aeiou")

	// TODO: Find first rune matching predicate
	// pos := strings.???(s, unicode.IsDigit)

	// TODO: What do these return if not found?
	// Answer: ???

	// TODO: Do they return rune positions or byte positions?
	// Answer: ???
}

// =============================================================================
// EXERCISE 3: PREFIX AND SUFFIX
// =============================================================================

func PracticePrefixSuffix() {
	// TODO: Check if string starts with prefix
	// hasPrefix := strings.???(s, prefix)

	// TODO: Check if string ends with suffix
	// hasSuffix := strings.???(s, suffix)

	// TODO: Remove prefix if present
	// withoutPrefix := strings.???(s, prefix)

	// TODO: Remove suffix if present
	// withoutSuffix := strings.???(s, suffix)

	// TODO: Remove prefix and tell if it was found (Go 1.20+)
	// after, found := strings.???(s, prefix)

	// TODO: Remove suffix and tell if it was found (Go 1.20+)
	// before, found := strings.???(s, suffix)

	// TODO: What happens if prefix/suffix not present in TrimPrefix/TrimSuffix?
	// Answer: ???
}

// =============================================================================
// EXERCISE 4: SPLITTING
// =============================================================================

func PracticeSplitting() {
	// TODO: Split on every occurrence of separator
	// parts := strings.???(s, sep)

	// TODO: Split with maximum N parts
	// parts := strings.???(s, sep, n)

	// TODO: Split but keep separator at end of each part
	// parts := strings.???(s, sep)

	// TODO: Split on any whitespace
	// words := strings.???(s)

	// TODO: Split using custom predicate function
	// parts := strings.???(s, func(r rune) bool { ... })

	// TODO: What does Split return for empty string?
	// strings.Split("", ",") returns ???

	// TODO: What's the difference between Fields and Split?
	// Fields: ???
	// Split: ???
}

// =============================================================================
// EXERCISE 5: JOINING
// =============================================================================

func PracticeJoining() {
	// TODO: Join strings with separator
	// result := strings.???(parts, sep)

	// TODO: What's the relationship between Join and Split?
	// Answer: ???

	// TODO: What does Join return for empty slice?
	// strings.Join([]string{}, ",") returns ???
}

// =============================================================================
// EXERCISE 6: TRIMMING
// =============================================================================

func PracticeTrimming() {
	// TODO: Remove whitespace from both ends
	// trimmed := strings.???(s)

	// TODO: Remove any characters from cutset from both ends
	// trimmed := strings.???(s, cutset)

	// TODO: Trim from left side only
	// trimmed := strings.???(s, cutset)

	// TODO: Trim from right side only
	// trimmed := strings.???(s, cutset)

	// TODO: Trim using predicate function
	// trimmed := strings.???(s, func(r rune) bool { ... })

	// TODO: CRITICAL: What does Trim actually remove?
	// Trim("hello", "lo") returns ???
	// Because: ???

	// TODO: How is Trim different from TrimPrefix/TrimSuffix?
	// Trim: removes ??? from ends
	// TrimSuffix: removes ??? from end
}

// =============================================================================
// EXERCISE 7: CASE CONVERSION
// =============================================================================

func PracticeCase() {
	// TODO: Convert to uppercase
	// upper := strings.???(s)

	// TODO: Convert to lowercase
	// lower := strings.???(s)

	// TODO: Case-insensitive comparison
	// equal := strings.???(s1, s2)

	// TODO: Why use EqualFold instead of ToLower comparison?
	// Answer: ???

	// TODO: Are these functions Unicode-aware?
	// Answer: ???
}

// =============================================================================
// EXERCISE 8: REPLACE
// =============================================================================

func PracticeReplace() {
	// TODO: Replace first N occurrences
	// result := strings.???(s, old, new, n)

	// TODO: Replace all occurrences
	// result := strings.???(s, old, new)

	// TODO: What's the shorthand for "replace all"?
	// result := strings.???(s, old, new)

	// TODO: Transform each rune with a function
	// result := strings.???(mapping, s)

	// TODO: How do you drop a rune in Map function?
	// Return ??? from the mapping function
}

// =============================================================================
// EXERCISE 9: strings.Replacer
// =============================================================================

func PracticeReplacer() {
	// TODO: Create a Replacer for multiple replacements
	// r := strings.???(old1, new1, old2, new2, ...)

	// TODO: Replace all occurrences in string
	// result := r.???(s)

	// TODO: Write replaced string to writer
	// n, err := r.???(writer, s)

	// TODO: How many arguments does NewReplacer take?
	// Answer: ??? number (must be even - pairs of old/new)

	// TODO: What order does Replacer process strings?
	// Answer: From ??? to ??? old string

	// TODO: When use Replacer vs multiple ReplaceAll calls?
	// Answer: ???

	// TODO: Can Replacer be reused?
	// Answer: ???
}

// =============================================================================
// EXERCISE 10: strings.Clone
// =============================================================================

func PracticeClone() {
	// TODO: Clone a string
	// cloned := strings.???(s)

	// TODO: Why does Clone exist?
	// Answer: To ??? when slicing large strings

	// TODO: When should you use Clone?
	// 1. After ??? a large string
	// 2. When sharing strings between ??? with different lifetimes
	// 3. Breaking reference to ??? parent string

	// TODO: When NOT to use Clone?
	// Answer: ???

	// TODO: What happens when cloning a string literal?
	// Answer: ???

	// TODO: What Go version added Clone?
	// Answer: Go ???
}

// =============================================================================
// EXERCISE 11: REPEAT AND COMPARE
// =============================================================================

func PracticeRepeatCompare() {
	// TODO: Repeat string N times
	// result := strings.???(s, count)

	// TODO: Compare two strings lexicographically
	// cmp := strings.???(a, b)

	// TODO: What does Compare return?
	// < 0 if ???
	// 0 if ???
	// > 0 if ???

	// TODO: Why rarely use Compare?
	// Answer: ???
}

// =============================================================================
// EXERCISE 12: CUT (Go 1.18+)
// =============================================================================

func PracticeCut() {
	// TODO: Split around first occurrence of separator
	// before, after, found := strings.???(s, sep)

	// TODO: What does Cut return if separator not found?
	// before = ???
	// after = ???
	// found = ???

	// TODO: Why use Cut instead of Index + slicing?
	// 1. ???
	// 2. ???
	// 3. ???

	// TODO: Does Cut find first or last occurrence?
	// Answer: ???
}

// =============================================================================
// EXERCISE 13: strings.Builder
// =============================================================================

func PracticeBuilder() {
	// TODO: Create a Builder
	// var b strings.???

	// TODO: Write a string to Builder
	// b.???(s)

	// TODO: Write a byte to Builder
	// b.???(c)

	// TODO: Write a rune to Builder
	// b.???(r)

	// TODO: Get the final string
	// result := b.???()

	// TODO: Pre-allocate capacity
	// b.???(n)

	// TODO: Clear and reuse Builder
	// b.???()

	// TODO: Get current length
	// length := b.???()

	// TODO: Why use Builder instead of +=?
	// Answer: ???

	// TODO: When should you call Grow()?
	// Answer: ???
}

// =============================================================================
// EXERCISE 14: strings.Reader
// =============================================================================

func PracticeReader() {
	// TODO: Create a Reader from string
	// r := strings.???(s)

	// TODO: Read bytes from Reader
	// n, err := r.???(buf)

	// TODO: Read a single byte
	// b, err := r.???()

	// TODO: Read a single rune
	// r, size, err := r.???()

	// TODO: Seek to position
	// pos, err := r.???(offset, whence)

	// TODO: Get original string length
	// size := r.???()

	// TODO: Get unread portion length
	// remaining := r.???()

	// TODO: Reuse Reader with new string
	// r.???(newString)

	// TODO: Why use strings.Reader?
	// Answer: ???
}

// =============================================================================
// EXERCISE 15: COMMON PATTERNS
// =============================================================================

func PracticePatterns() {
	// TODO: Normalize whitespace (collapse multiple spaces)
	// How: ???

	// TODO: Check if string is blank (empty or whitespace only)
	// isBlank := strings.???(s) == ""

	// TODO: Truncate string with ellipsis
	// If len(s) > maxLen, return ???

	// TODO: Count lines in text
	// count := strings.???(s, "\n") + ???

	// TODO: Case-insensitive search
	// Use ??? to avoid allocations

	// TODO: Parse key=value
	// key, value, ok := strings.???(s, "=")
}

// =============================================================================
// EXERCISE 16: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// s := ""
	// for i := 0; i < 1000; i++ {
	//     s += "x"
	// }
	// Answer: ???
	// Fix: ???

	// Mistake 2: What's wrong?
	// s := "hello"
	// result := strings.Trim(s, "lo")  // Expecting "hel"
	// Answer: ???
	// Fix: ???

	// Mistake 3: What's wrong?
	// if strings.ToLower(a) == strings.ToLower(b) { ... }
	// Answer: ???
	// Fix: ???

	// Mistake 4: What's wrong?
	// i := strings.Index(s, "x")
	// before := s[:i]
	// Answer: ???
	// Fix: ???

	// Mistake 5: What happens?
	// parts := strings.Split("", ",")
	// len(parts) == ???
	// parts[0] == ???
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// 1. Check if contains substring?
	_ = "strings.???(s, substr)"

	// 2. Find first occurrence?
	_ = "strings.???(s, substr)"

	// 3. Check prefix?
	_ = "strings.???(s, prefix)"

	// 4. Remove suffix?
	_ = "strings.???(s, suffix)"

	// 5. Split on separator?
	_ = "strings.???(s, sep)"

	// 6. Split on whitespace?
	_ = "strings.???(s)"

	// 7. Join with separator?
	_ = "strings.???(parts, sep)"

	// 8. Remove whitespace?
	_ = "strings.???(s)"

	// 9. To uppercase?
	_ = "strings.???(s)"

	// 10. Case-insensitive equal?
	_ = "strings.???(s1, s2)"

	// 11. Replace all?
	_ = "strings.???(s, old, new)"

	// 12. Repeat N times?
	_ = "strings.???(s, count)"

	// 13. Split around separator (Go 1.18+)?
	_ = "strings.???(s, sep)"

	// 14. Efficient string building?
	_ = "var b strings.???; b.WriteString(...); b.String()"

	// 15. Reader from string?
	_ = "strings.???(s)"

	// 16. Count occurrences?
	_ = "strings.???(s, substr)"

	// 17. Find last occurrence?
	_ = "strings.???(s, substr)"

	// 18. Map runes?
	_ = "strings.???(mapping, s)"

	// 19. Multiple replacements?
	_ = "strings.???(old1, new1, old2, new2, ...)"

	// 20. Clone string?
	_ = "strings.???(s)"
}

// =============================================================================
// MINI PROJECT: TEXT PROCESSOR
// =============================================================================

func MiniProject() {
	// Build a text processing utility with these functions:
	//
	// 1. func NormalizeWhitespace(s string) string
	//    - Trim leading/trailing whitespace
	//    - Collapse multiple spaces into one
	//    - Replace tabs/newlines with spaces
	//
	// 2. func WordCount(s string) int
	//    - Count words (split on whitespace)
	//
	// 3. func SentenceCount(s string) int
	//    - Count sentences (split on . ! ?)
	//    - Handle abbreviations like "Dr." carefully
	//
	// 4. func ExtractEmails(s string) []string
	//    - Find all email addresses (simple: contains @ and .)
	//    - Return unique list
	//
	// 5. func Slugify(s string) string
	//    - Convert "Hello, World!" to "hello-world"
	//    - Lowercase, replace spaces with hyphens
	//    - Remove non-alphanumeric except hyphens
	//
	// 6. func TruncateWords(s string, n int) string
	//    - Truncate to first N words
	//    - Add "..." if truncated
	//
	// 7. func ReverseWords(s string) string
	//    - Reverse order of words
	//    - "hello world" → "world hello"
	//
	// 8. func CamelToSnake(s string) string
	//    - "getUserName" → "get_user_name"
	//
	// 9. func SnakeToCamel(s string) string
	//    - "get_user_name" → "getUserName"
	//
	// 10. func IndentLines(s string, indent int) string
	//     - Add N spaces before each line
	//
	// Scaffold:
	//
	// func NormalizeWhitespace(s string) string {
	//     // Replace tabs and newlines with spaces
	//     s = strings.ReplaceAll(s, "\t", " ")
	//     s = strings.ReplaceAll(s, "\n", " ")
	//     // Collapse multiple spaces
	//     words := strings.Fields(s)
	//     return strings.Join(words, " ")
	// }
	//
	// func WordCount(s string) int {
	//     words := strings.Fields(s)
	//     return len(words)
	// }
	//
	// func Slugify(s string) string {
	//     // Convert to lowercase
	//     s = strings.ToLower(s)
	//     
	//     // Replace spaces with hyphens
	//     s = strings.ReplaceAll(s, " ", "-")
	//     
	//     // Remove non-alphanumeric except hyphens
	//     s = strings.Map(func(r rune) rune {
	//         if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
	//             return r
	//         }
	//         return -1  // Drop character
	//     }, s)
	//     
	//     // Collapse multiple hyphens
	//     for strings.Contains(s, "--") {
	//         s = strings.ReplaceAll(s, "--", "-")
	//     }
	//     
	//     // Trim hyphens from ends
	//     s = strings.Trim(s, "-")
	//     
	//     return s
	// }
	//
	// func TruncateWords(s string, n int) string {
	//     words := strings.Fields(s)
	//     if len(words) <= n {
	//         return s
	//     }
	//     truncated := strings.Join(words[:n], " ")
	//     return truncated + "..."
	// }
	//
	// func ReverseWords(s string) string {
	//     words := strings.Fields(s)
	//     for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
	//         words[i], words[j] = words[j], words[i]
	//     }
	//     return strings.Join(words, " ")
	// }
	//
	// func CamelToSnake(s string) string {
	//     var b strings.Builder
	//     for i, r := range s {
	//         if unicode.IsUpper(r) && i > 0 {
	//             b.WriteByte('_')
	//         }
	//         b.WriteRune(unicode.ToLower(r))
	//     }
	//     return b.String()
	// }
	//
	// func SnakeToCamel(s string) string {
	//     parts := strings.Split(s, "_")
	//     var b strings.Builder
	//     for i, part := range parts {
	//         if i == 0 {
	//             b.WriteString(part)
	//         } else if len(part) > 0 {
	//             runes := []rune(part)
	//             runes[0] = unicode.ToUpper(runes[0])
	//             b.WriteString(string(runes))
	//         }
	//     }
	//     return b.String()
	// }
	//
	// func IndentLines(s string, indent int) string {
	//     if indent <= 0 {
	//         return s
	//     }
	//     
	//     prefix := strings.Repeat(" ", indent)
	//     lines := strings.Split(s, "\n")
	//     
	//     var b strings.Builder
	//     for i, line := range lines {
	//         if i > 0 {
	//             b.WriteByte('\n')
	//         }
	//         if line != "" {  // Don't indent empty lines
	//             b.WriteString(prefix)
	//         }
	//         b.WriteString(line)
	//     }
	//     
	//     return b.String()
	// }
	//
	// // Test your implementation:
	// func main() {
	//     text := "  Hello,   World!  This is   a test.  "
	//     fmt.Println("Original:", text)
	//     fmt.Println("Normalized:", NormalizeWhitespace(text))
	//     fmt.Println("Word count:", WordCount(text))
	//     fmt.Println("Slugified:", Slugify("Hello, World!"))
	//     fmt.Println("Truncated:", TruncateWords(text, 3))
	//     fmt.Println("Reversed:", ReverseWords(text))
	//     fmt.Println("CamelToSnake:", CamelToSnake("getUserName"))
	//     fmt.Println("SnakeToCamel:", SnakeToCamel("get_user_name"))
	// }
}

// AllPractice is not meant to be called - this file is for reading.
func AllPractice() {
	PracticeContains()
	PracticeIndex()
	PracticePrefixSuffix()
	PracticeSplitting()
	PracticeJoining()
	PracticeTrimming()
	PracticeCase()
	PracticeReplace()
	PracticeReplacer()
	PracticeClone()
	PracticeRepeatCompare()
	PracticeCut()
	PracticeBuilder()
	PracticeReader()
	PracticePatterns()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
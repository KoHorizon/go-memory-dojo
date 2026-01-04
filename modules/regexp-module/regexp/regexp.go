// Package regexpmod provides comprehensive documentation and working examples
// for Go's regexp package - regular expression matching and manipulation.
package regexpmod

import (
	"fmt"
	"regexp"
	"strings"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF REGEXP IN GO
// =============================================================================
//
// WHY REGULAR EXPRESSIONS?
//
// Regular expressions (regex) are patterns for matching text.
// They're powerful for:
//   - Validation (email, phone, URL)
//   - Extraction (dates, IDs, tags)
//   - Transformation (search and replace)
//   - Parsing (simple text formats)
//
// GO'S REGEX ENGINE
//
// - RE2 syntax (safe, guaranteed linear time)
// - No backtracking (prevents catastrophic backtracking)
// - No lookahead/lookbehind (RE2 limitation)
// - UTF-8 aware
//
// COMPILE ONCE, USE MANY
//
// Compiling a regex is expensive. Always compile once and reuse:
//
//   // BAD - compiles every iteration
//   for _, s := range data {
//       matched, _ := regexp.MatchString(`\d+`, s)
//   }
//
//   // GOOD - compile once
//   re := regexp.MustCompile(`\d+`)
//   for _, s := range data {
//       matched := re.MatchString(s)
//   }
//
// WHEN NOT TO USE REGEX
//
// - Simple substring searches → use strings.Contains
// - Prefix/suffix checks → use strings.HasPrefix/HasSuffix
// - Fixed string replacement → use strings.Replace
// - Parsing structured data → use encoding/json, encoding/xml
//
// Regex is powerful but slower than simple string operations.
//
// =============================================================================

// =============================================================================
// SECTION 2: COMPILING PATTERNS
// =============================================================================
//
// Compile(expr string) (*Regexp, error)
//   Compile and return regex, or error if invalid
//
// MustCompile(expr string) *Regexp
//   Compile and panic on error
//   Use for package-level variables with known-good patterns
//
// CompilePOSIX(expr string) (*Regexp, error)
//   Compile with POSIX semantics (leftmost-longest match)
//
// WHEN TO USE EACH
//
//   MustCompile: Package-level constants, known-good patterns
//   Compile: User input, runtime patterns, error handling needed
//
// PATTERN SYNTAX
//
//   .        - Any character (including newline in multiline mode)
//   [abc]    - Character class (a, b, or c)
//   [^abc]   - Negated class (not a, b, or c)
//   [a-z]    - Range
//   \d       - Digit [0-9]
//   \D       - Non-digit
//   \w       - Word character [A-Za-z0-9_]
//   \W       - Non-word character
//   \s       - Whitespace [ \t\n\r\f]
//   \S       - Non-whitespace
//   ^        - Start of text (or line in multiline)
//   $        - End of text (or line in multiline)
//   \b       - Word boundary
//   \B       - Not word boundary
//
// QUANTIFIERS
//
//   *        - 0 or more (greedy)
//   +        - 1 or more (greedy)
//   ?        - 0 or 1 (greedy)
//   {n}      - Exactly n
//   {n,}     - n or more
//   {n,m}    - Between n and m
//   *?       - 0 or more (non-greedy)
//   +?       - 1 or more (non-greedy)
//   ??       - 0 or 1 (non-greedy)
//
// GROUPING
//
//   (...)    - Capturing group
//   (?:...)  - Non-capturing group
//   (?P<name>...) - Named capturing group
//   |        - Alternation (or)
//
// =============================================================================

func DemonstrateCompiling() {
	// Compile - with error handling
	re1, err := regexp.Compile(`\d+`)
	if err != nil {
		fmt.Printf("Compile error: %v\n", err)
		return
	}
	fmt.Printf("Compiled pattern: %s\n", re1.String())

	// MustCompile - panics on error
	re2 := regexp.MustCompile(`[a-z]+`)
	fmt.Printf("MustCompile pattern: %s\n", re2.String())

	// Invalid pattern
	_, err = regexp.Compile(`[invalid`)
	fmt.Printf("Invalid pattern error: %v\n", err)

	// Package-level pattern (common idiom)
	var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	fmt.Printf("Email pattern: %s\n", emailPattern.String())
}

// =============================================================================
// SECTION 3: MATCHING - DOES IT MATCH?
// =============================================================================
//
// MATCHING FUNCTIONS
//
// MatchString(pattern, s string) (bool, error)
//   Package function - compiles every time (slow!)
//
// re.MatchString(s string) bool
//   Does s match the pattern?
//   Most common matching method
//
// re.Match(b []byte) bool
//   Match against []byte
//
// SIMPLE YES/NO QUESTIONS
//
// Use these when you only need to know IF something matches,
// not WHERE or WHAT matched.
//
// =============================================================================

func DemonstrateMatching() {
	// Package function (avoid in loops!)
	matched, _ := regexp.MatchString(`\d+`, "abc123")
	fmt.Printf("MatchString: %v\n", matched)

	// Compiled regex (reusable)
	digitPattern := regexp.MustCompile(`\d+`)
	fmt.Printf("Has digits 'abc123': %v\n", digitPattern.MatchString("abc123"))
	fmt.Printf("Has digits 'abc': %v\n", digitPattern.MatchString("abc"))

	// Common patterns
	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	fmt.Printf("Valid email 'user@example.com': %v\n", emailPattern.MatchString("user@example.com"))
	fmt.Printf("Valid email 'invalid': %v\n", emailPattern.MatchString("invalid"))

	phonePattern := regexp.MustCompile(`^\d{3}-\d{3}-\d{4}$`)
	fmt.Printf("Valid phone '123-456-7890': %v\n", phonePattern.MatchString("123-456-7890"))
	fmt.Printf("Valid phone '1234567890': %v\n", phonePattern.MatchString("1234567890"))

	// []byte version
	data := []byte("test 123")
	fmt.Printf("Match []byte: %v\n", digitPattern.Match(data))
}

// =============================================================================
// SECTION 4: FINDING - WHERE/WHAT MATCHES?
// =============================================================================
//
// FIND METHODS
//
// re.FindString(s string) string
//   Return first match, or "" if none
//
// re.FindStringIndex(s string) []int
//   Return [start, end] of first match, or nil
//
// re.FindAllString(s string, n int) []string
//   Return up to n matches (n < 0 for all)
//
// re.FindAllStringIndex(s string, n int) [][]int
//   Return indices of up to n matches
//
// []BYTE VERSIONS
//
// re.Find(b []byte) []byte
// re.FindIndex(b []byte) []int
// re.FindAll(b []byte, n int) [][]byte
// re.FindAllIndex(b []byte, n int) [][]int
//
// SUBMATCH VERSIONS (see Section 5)
//
// re.FindStringSubmatch(s string) []string
// re.FindAllStringSubmatch(s string, n int) [][]string
//
// =============================================================================

func DemonstrateFinding() {
	text := "The quick brown fox jumps over the lazy dog"

	// FindString - first match
	wordPattern := regexp.MustCompile(`\b\w{5}\b`) // 5-letter words
	first := wordPattern.FindString(text)
	fmt.Printf("First 5-letter word: %s\n", first)

	// FindStringIndex - position
	indices := wordPattern.FindStringIndex(text)
	if indices != nil {
		fmt.Printf("Found at: [%d:%d]\n", indices[0], indices[1])
		fmt.Printf("Text at position: %s\n", text[indices[0]:indices[1]])
	}

	// FindAllString - all matches
	all := wordPattern.FindAllString(text, -1)
	fmt.Printf("All 5-letter words: %v\n", all)

	// FindAllString with limit
	limited := wordPattern.FindAllString(text, 3)
	fmt.Printf("First 3 matches: %v\n", limited)

	// Extract numbers
	numberPattern := regexp.MustCompile(`\d+`)
	numbers := numberPattern.FindAllString("abc 123 def 456 ghi 789", -1)
	fmt.Printf("Numbers: %v\n", numbers)

	// Extract emails
	emailPattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	emails := emailPattern.FindAllString("Contact: alice@example.com or bob@test.org", -1)
	fmt.Printf("Emails: %v\n", emails)

	// []byte version
	data := []byte("test 123 hello 456")
	matches := numberPattern.FindAll(data, -1)
	fmt.Printf("Numbers ([]byte): %s\n", matches)
}

// =============================================================================
// SECTION 5: SUBMATCHES - CAPTURING GROUPS
// =============================================================================
//
// WHAT ARE CAPTURING GROUPS?
//
// Parentheses in patterns create groups:
//   `(\d+)-(\d+)` has 2 groups
//
// Submatches include the whole match + each group:
//   Pattern: `(\d+)-(\d+)`
//   Text: "123-456"
//   Submatch: ["123-456", "123", "456"]
//                ^whole    ^g1    ^g2
//
// SUBMATCH METHODS
//
// re.FindStringSubmatch(s string) []string
//   Return [whole_match, group1, group2, ...]
//   nil if no match
//
// re.FindAllStringSubmatch(s string, n int) [][]string
//   Return all matches with their submatches
//
// re.FindStringSubmatchIndex(s string) []int
//   Return [whole_start, whole_end, g1_start, g1_end, ...]
//
// NON-CAPTURING GROUPS
//
// Use (?:...) when you need grouping but not capturing:
//   `(?:\d+)-(\d+)` only captures the second part
//
// =============================================================================

func DemonstrateSubmatches() {
	// Basic capturing
	datePattern := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	match := datePattern.FindStringSubmatch("Date: 2024-01-15")
	if match != nil {
		fmt.Printf("Full match: %s\n", match[0])
		fmt.Printf("Year: %s\n", match[1])
		fmt.Printf("Month: %s\n", match[2])
		fmt.Printf("Day: %s\n", match[3])
	}

	// Multiple matches with groups
	timePattern := regexp.MustCompile(`(\d{2}):(\d{2}):(\d{2})`)
	times := timePattern.FindAllStringSubmatch("Times: 10:30:45 and 14:22:33", -1)
	for i, t := range times {
		fmt.Printf("Time %d: %s (h=%s, m=%s, s=%s)\n", i, t[0], t[1], t[2], t[3])
	}

	// Extract key-value pairs
	kvPattern := regexp.MustCompile(`(\w+)=(\w+)`)
	pairs := kvPattern.FindAllStringSubmatch("name=Alice age=30 city=NYC", -1)
	config := make(map[string]string)
	for _, pair := range pairs {
		config[pair[1]] = pair[2]
	}
	fmt.Printf("Config: %v\n", config)

	// Non-capturing group
	urlPattern := regexp.MustCompile(`https?://([a-z.]+)`)
	url := urlPattern.FindStringSubmatch("Visit https://example.com")
	if url != nil {
		fmt.Printf("Domain: %s\n", url[1]) // Only one group
	}

	// Indices for extraction
	emailPattern := regexp.MustCompile(`([a-z]+)@([a-z]+\.[a-z]+)`)
	indices := emailPattern.FindStringSubmatchIndex("alice@example.com")
	if indices != nil {
		text := "alice@example.com"
		fmt.Printf("Full: %s\n", text[indices[0]:indices[1]])
		fmt.Printf("User: %s\n", text[indices[2]:indices[3]])
		fmt.Printf("Domain: %s\n", text[indices[4]:indices[5]])
	}
}

// =============================================================================
// SECTION 6: NAMED GROUPS
// =============================================================================
//
// NAMED CAPTURING GROUPS
//
// Syntax: (?P<name>pattern)
//
// Benefits:
//   - Self-documenting patterns
//   - Access by name instead of index
//   - Easier to maintain
//
// METHODS
//
// re.SubexpNames() []string
//   Return names of all groups (first element is always "")
//
// Use with FindStringSubmatch to build map
//
// =============================================================================

func DemonstrateNamedGroups() {
	// Named groups
	pattern := regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`)

	// Get group names
	names := pattern.SubexpNames()
	fmt.Printf("Group names: %v\n", names)

	// Match and extract
	match := pattern.FindStringSubmatch("2024-01-15")
	if match != nil {
		// Build map of name → value
		result := make(map[string]string)
		for i, name := range names {
			if i != 0 && name != "" { // Skip whole match and unnamed
				result[name] = match[i]
			}
		}
		fmt.Printf("Date parts: %v\n", result)
	}

	// Practical example: Parse log line
	logPattern := regexp.MustCompile(
		`^\[(?P<timestamp>[^\]]+)\] ` +
			`(?P<level>\w+) ` +
			`(?P<message>.+)$`)

	logLine := "[2024-01-15 10:30:45] ERROR Connection timeout"
	match2 := logPattern.FindStringSubmatch(logLine)
	if match2 != nil {
		log := make(map[string]string)
		for i, name := range logPattern.SubexpNames() {
			if i != 0 && name != "" {
				log[name] = match2[i]
			}
		}
		fmt.Printf("Log: timestamp=%s, level=%s, message=%s\n",
			log["timestamp"], log["level"], log["message"])
	}

	// Helper function to convert to map
	matchToMap := func(re *regexp.Regexp, s string) map[string]string {
		match := re.FindStringSubmatch(s)
		if match == nil {
			return nil
		}
		result := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		return result
	}

	urlPattern := regexp.MustCompile(`(?P<scheme>https?)://(?P<host>[^/]+)(?P<path>/.*)?`)
	url := matchToMap(urlPattern, "https://example.com/api/v1/users")
	fmt.Printf("URL: %v\n", url)
}

// =============================================================================
// SECTION 7: REPLACING
// =============================================================================
//
// REPLACE METHODS
//
// re.ReplaceAllString(src, repl string) string
//   Replace all matches with repl
//
// re.ReplaceAllLiteralString(src, repl string) string
//   Replace with literal repl (no expansion)
//
// re.ReplaceAllStringFunc(src string, repl func(string) string) string
//   Replace using function
//
// EXPANSION IN REPLACEMENT
//
// In repl string:
//   $0 or $& - Whole match
//   $1, $2   - Captured groups
//   ${1}     - Named reference (clearer)
//   $$       - Literal $
//
// =============================================================================

func DemonstrateReplacing() {
	text := "Hello, World! Hello, Go!"

	// Simple replacement
	pattern := regexp.MustCompile(`Hello`)
	result := pattern.ReplaceAllString(text, "Hi")
	fmt.Printf("Replaced: %s\n", result)

	// Replace with captured groups
	datePattern := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	formatted := datePattern.ReplaceAllString(
		"Date: 2024-01-15",
		"$2/$3/$1", // MM/DD/YYYY
	)
	fmt.Printf("Reformatted date: %s\n", formatted)

	// Named groups in replacement
	namedPattern := regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`)
	formatted2 := namedPattern.ReplaceAllString(
		"2024-01-15",
		"${month}/${day}/${year}",
	)
	fmt.Printf("With named groups: %s\n", formatted2)

	// Literal replacement (no expansion)
	literal := pattern.ReplaceAllLiteralString(text, "$1")
	fmt.Printf("Literal $1: %s\n", literal)

	// Function-based replacement
	digitPattern := regexp.MustCompile(`\d+`)
	doubled := digitPattern.ReplaceAllStringFunc(
		"Numbers: 5, 10, 15",
		func(s string) string {
			// Parse, double, return
			var n int
			fmt.Sscanf(s, "%d", &n)
			return fmt.Sprintf("%d", n*2)
		},
	)
	fmt.Printf("Doubled: %s\n", doubled)

	// Remove matches
	spacePattern := regexp.MustCompile(`\s+`)
	noSpaces := spacePattern.ReplaceAllString("a  b    c", "")
	fmt.Printf("No spaces: %s\n", noSpaces)

	// Normalize whitespace
	normalized := spacePattern.ReplaceAllString("  hello   world  ", " ")
	normalized = strings.TrimSpace(normalized)
	fmt.Printf("Normalized: %s\n", normalized)

	// Redact sensitive data
	ssnPattern := regexp.MustCompile(`\d{3}-\d{2}-\d{4}`)
	redacted := ssnPattern.ReplaceAllString(
		"SSN: 123-45-6789",
		"XXX-XX-XXXX",
	)
	fmt.Printf("Redacted: %s\n", redacted)
}

// =============================================================================
// SECTION 8: SPLITTING
// =============================================================================
//
// re.Split(s string, n int) []string
//   Split string on matches
//   n < 0: return all substrings
//   n == 0: result is nil
//   n > 0: at most n substrings
//
// Unlike strings.Split, separator is a pattern, not literal.
//
// =============================================================================

func DemonstrateSplitting() {
	// Split on whitespace (including multiple spaces)
	spacePattern := regexp.MustCompile(`\s+`)
	words := spacePattern.Split("  hello   world   foo  ", -1)
	fmt.Printf("Words: %v\n", words)

	// Split on comma with optional spaces
	csvPattern := regexp.MustCompile(`\s*,\s*`)
	fields := csvPattern.Split("apple, banana ,cherry,  date", -1)
	fmt.Printf("Fields: %v\n", fields)

	// Split with limit
	limited := csvPattern.Split("a,b,c,d,e", 3)
	fmt.Printf("Limited: %v\n", limited) // [a b c,d,e]

	// Split on multiple delimiters
	delimPattern := regexp.MustCompile(`[,;:|]`)
	parts := delimPattern.Split("a,b;c:d|e", -1)
	fmt.Printf("Multi-delim: %v\n", parts)

	// Split sentences (simplified)
	sentencePattern := regexp.MustCompile(`[.!?]\s+`)
	sentences := sentencePattern.Split("Hello! How are you? I'm fine.", -1)
	fmt.Printf("Sentences: %v\n", sentences)
}

// =============================================================================
// SECTION 9: LONGEST MATCH
// =============================================================================
//
// re.Longest()
//   Make regex prefer longest match (POSIX behavior)
//   By default, Go regex is leftmost-first (like Perl)
//
// GREEDY VS NON-GREEDY
//
//   .*   - Greedy: matches as much as possible
//   .*?  - Non-greedy: matches as little as possible
//
// LEFTMOST-FIRST VS LEFTMOST-LONGEST
//
//   Leftmost-first: Return first match found
//   Leftmost-longest: Among earliest matches, return longest
//
// =============================================================================

func DemonstrateLongest() {
	text := "Hello World"

	// Default: leftmost-first
	pattern1 := regexp.MustCompile(`\w+`)
	fmt.Printf("Leftmost-first: %s\n", pattern1.FindString(text))

	// Longest: leftmost-longest
	pattern2 := regexp.MustCompile(`\w+`)
	pattern2.Longest()
	fmt.Printf("Leftmost-longest: %s\n", pattern2.FindString(text))

	// Greedy vs non-greedy
	greedyPattern := regexp.MustCompile(`<.*>`)
	nonGreedyPattern := regexp.MustCompile(`<.*?>`)

	html := "<b>bold</b> <i>italic</i>"
	fmt.Printf("Greedy: %s\n", greedyPattern.FindString(html))        // <b>bold</b> <i>italic</i>
	fmt.Printf("Non-greedy: %s\n", nonGreedyPattern.FindString(html)) // <b>
}

// =============================================================================
// SECTION 10: COMMON PATTERNS
// =============================================================================

func DemonstrateCommonPatterns() {
	// Email validation (simple)
	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	fmt.Printf("Email 'user@example.com': %v\n", emailPattern.MatchString("user@example.com"))

	// URL extraction
	urlPattern := regexp.MustCompile(`https?://[^\s]+`)
	urls := urlPattern.FindAllString("Visit https://example.com or http://test.org", -1)
	fmt.Printf("URLs: %v\n", urls)

	// Phone number (US format)
	phonePattern := regexp.MustCompile(`^\(?(\d{3})\)?[-.\s]?(\d{3})[-.\s]?(\d{4})$`)
	phones := []string{"123-456-7890", "(123) 456-7890", "123.456.7890", "1234567890"}
	for _, p := range phones {
		fmt.Printf("Valid phone '%s': %v\n", p, phonePattern.MatchString(p))
	}

	// Extract hashtags
	hashtagPattern := regexp.MustCompile(`#\w+`)
	hashtags := hashtagPattern.FindAllString("Love #golang and #programming!", -1)
	fmt.Printf("Hashtags: %v\n", hashtags)

	// Remove HTML tags
	htmlPattern := regexp.MustCompile(`<[^>]*>`)
	cleaned := htmlPattern.ReplaceAllString("<p>Hello <b>World</b></p>", "")
	fmt.Printf("No HTML: %s\n", cleaned)

	// Validate hex color
	colorPattern := regexp.MustCompile(`^#[0-9A-Fa-f]{6}$`)
	colors := []string{"#FF0000", "#abc", "#gggggg"}
	for _, c := range colors {
		fmt.Printf("Valid color '%s': %v\n", c, colorPattern.MatchString(c))
	}

	// Extract mentions
	mentionPattern := regexp.MustCompile(`@\w+`)
	mentions := mentionPattern.FindAllString("Hi @alice and @bob!", -1)
	fmt.Printf("Mentions: %v\n", mentions)

	// Validate IP address (simple)
	ipPattern := regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)
	ips := []string{"192.168.1.1", "999.999.999.999", "10.0.0.1"}
	for _, ip := range ips {
		fmt.Printf("IP format '%s': %v\n", ip, ipPattern.MatchString(ip))
	}
}

// =============================================================================
// SECTION 11: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Compiling in a loop
	_ = `
	// WRONG - compiles every iteration
	for _, s := range data {
		matched, _ := regexp.MatchString(pattern, s)
	}

	// RIGHT - compile once
	re := regexp.MustCompile(pattern)
	for _, s := range data {
		matched := re.MatchString(s)
	}
	`
	fmt.Println("Mistake 1: Compile once, use many times")

	// Mistake 2: Not escaping metacharacters
	_ = `
	// WRONG - . matches any character
	dotPattern := regexp.MustCompile(".")
	dotPattern.MatchString("example.com")  // Matches "e", not "."

	// RIGHT - escape the dot
	dotPattern := regexp.MustCompile("\.")
	// Or use backticks: regexp.MustCompile(` + "`" + `\.` + "`" + `)
	`
	fmt.Println("Mistake 2: Escape metacharacters with \\")

	// Mistake 3: Forgetting anchors
	_ = `
	// WRONG - matches anywhere in string
	pattern := regexp.MustCompile("\d{3}")
	pattern.MatchString("abc123def")  // true (probably not intended)

	// RIGHT - use anchors
	pattern := regexp.MustCompile("^\d{3}$")
	pattern.MatchString("123")  // true
	pattern.MatchString("abc123")  // false
	`
	fmt.Println("Mistake 3: Use ^ and $ for exact matches")

	// Mistake 4: Wrong group indexing
	_ = `
	// WRONG - groups start at 1, not 0
	match := re.FindStringSubmatch(s)
	group1 := match[0]  // This is the whole match!

	// RIGHT
	wholeMatch := match[0]
	group1 := match[1]
	group2 := match[2]
	`
	fmt.Println("Mistake 4: match[0] is whole match, groups start at [1]")

	// Mistake 5: Not checking for nil
	_ = `
	// WRONG - panics if no match
	match := re.FindStringSubmatch(s)
	value := match[1]  // PANIC if match is nil!

	// RIGHT
	match := re.FindStringSubmatch(s)
	if match != nil {
		value := match[1]
	}
	`
	fmt.Println("Mistake 5: Always check if match is nil")

	// Mistake 6: Overusing regex
	_ = `
	// WRONG - regex for simple substring
	re := regexp.MustCompile("hello")
	found := re.MatchString(text)

	// RIGHT - use strings package
	found := strings.Contains(text, "hello")
	`
	fmt.Println("Mistake 6: Use strings package for simple operations")

	// Mistake 7: Catastrophic backtracking concern (not in Go)
	fmt.Println("Mistake 7: Go uses RE2 (no backtracking), but still avoid complex patterns")
}

// =============================================================================
// SECTION 12: PERFORMANCE TIPS
// =============================================================================
//
// 1. COMPILE ONCE, REUSE
//    Most important optimization
//
// 2. USE SIMPLE PATTERNS
//    Simpler = faster
//
// 3. ANCHOR WHEN POSSIBLE
//    ^...$ helps regex engine skip non-matching positions
//
// 4. USE NON-CAPTURING GROUPS
//    (?:...) when you don't need the capture
//
// 5. PREFER strings PACKAGE
//    For simple operations (Contains, HasPrefix, etc.)
//
// 6. COMPILE ERRORS AT STARTUP
//    Use MustCompile for package-level patterns
//    Fail fast on bad patterns
//
// 7. BENCHMARK YOUR PATTERNS
//    Profile before optimizing
//
// =============================================================================

// =============================================================================
// SECTION 13: QUICK SYNTAX REFERENCE
// =============================================================================
//
// CHARACTER CLASSES
//   .       Any character
//   [abc]   a, b, or c
//   [^abc]  Not a, b, or c
//   [a-z]   a through z
//   \d      Digit [0-9]
//   \D      Not digit
//   \w      Word char [A-Za-z0-9_]
//   \W      Not word char
//   \s      Whitespace
//   \S      Not whitespace
//
// ANCHORS
//   ^       Start of text
//   $       End of text
//   \b      Word boundary
//   \B      Not word boundary
//
// QUANTIFIERS
//   *       0 or more
//   +       1 or more
//   ?       0 or 1
//   {n}     Exactly n
//   {n,}    n or more
//   {n,m}   Between n and m
//   *?      0 or more (non-greedy)
//   +?      1 or more (non-greedy)
//
// GROUPS
//   (...)           Capturing group
//   (?:...)         Non-capturing group
//   (?P<name>...)   Named group
//   |               Alternation
//
// ESCAPING
//   \               Escape next character
//   \\              Literal backslash
//   \.              Literal dot
//   \$              Literal dollar
//
// FLAGS (via (?flags))
//   (?i)    Case-insensitive
//   (?m)    Multi-line (^ and $ match line boundaries)
//   (?s)    Dot matches newline
//
// =============================================================================

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateCompiling()
	DemonstrateMatching()
	DemonstrateFinding()
	DemonstrateSubmatches()
	DemonstrateNamedGroups()
	DemonstrateReplacing()
	DemonstrateSplitting()
	DemonstrateLongest()
	DemonstrateCommonPatterns()
	DemonstrateCommonMistakes()
}
// Package regexp_practice is your daily practice space for the regexp package.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against regexp/regexp.go
// 3. Note what you missed - focus on those tomorrow
package regexp_practice

// =============================================================================
// EXERCISE 1: PHILOSOPHY
// =============================================================================

func PracticePhilosophy() {
	// TODO: What regex engine does Go use?
	// Answer: ???

	// TODO: Key properties of Go's regex (2 things)
	// 1. ???
	// 2. ???

	// TODO: When NOT to use regex?
	// - Simple substring → use strings.???
	// - Prefix check → use strings.???
	// - Fixed replacement → use strings.???

	// TODO: Why compile once?
	// Answer: ???
}

// =============================================================================
// EXERCISE 2: COMPILING
// =============================================================================

func PracticeCompiling() {
	// TODO: Compile with error handling
	// re, err := regexp.???(pattern)

	// TODO: Compile and panic on error
	// re := regexp.???(pattern)

	// TODO: When use Compile vs MustCompile?
	// Compile: ???
	// MustCompile: ???

	// TODO: POSIX semantics compilation
	// re, err := regexp.???(pattern)
}

// =============================================================================
// EXERCISE 3: PATTERN SYNTAX - CHARACTER CLASSES
// =============================================================================

func PracticeCharacterClasses() {
	// TODO: Match any character
	// ???

	// TODO: Match a, b, or c
	// ???

	// TODO: Match NOT a, b, or c
	// ???

	// TODO: Match a through z
	// ???

	// TODO: Match any digit
	// ???

	// TODO: Match any non-digit
	// ???

	// TODO: Match word character [A-Za-z0-9_]
	// ???

	// TODO: Match whitespace
	// ???
}

// =============================================================================
// EXERCISE 4: PATTERN SYNTAX - ANCHORS AND BOUNDARIES
// =============================================================================

func PracticeAnchors() {
	// TODO: Start of text
	// ???

	// TODO: End of text
	// ???

	// TODO: Word boundary
	// ???

	// TODO: Not word boundary
	// ???
}

// =============================================================================
// EXERCISE 5: PATTERN SYNTAX - QUANTIFIERS
// =============================================================================

func PracticeQuantifiers() {
	// TODO: 0 or more (greedy)
	// ???

	// TODO: 1 or more (greedy)
	// ???

	// TODO: 0 or 1 (greedy)
	// ???

	// TODO: Exactly n times
	// ???

	// TODO: n or more times
	// ???

	// TODO: Between n and m times
	// ???

	// TODO: 0 or more (non-greedy)
	// ???

	// TODO: 1 or more (non-greedy)
	// ???
}

// =============================================================================
// EXERCISE 6: PATTERN SYNTAX - GROUPS
// =============================================================================

func PracticeGroups() {
	// TODO: Capturing group
	// ???

	// TODO: Non-capturing group
	// ???

	// TODO: Named capturing group
	// ???

	// TODO: Alternation (or)
	// ???
}

// =============================================================================
// EXERCISE 7: MATCHING
// =============================================================================

func PracticeMatching() {
	// TODO: Package-level match (slow, avoid in loops!)
	// matched, err := regexp.???(pattern, s)

	// TODO: Match string against compiled regex
	// matched := re.???(s)

	// TODO: Match []byte against compiled regex
	// matched := re.???(b)
}

// =============================================================================
// EXERCISE 8: FINDING - SINGLE MATCH
// =============================================================================

func PracticeFindingSingle() {
	// TODO: Find first match as string
	// match := re.???(s)

	// TODO: Find first match indices [start, end]
	// indices := re.???(s)

	// TODO: What does FindString return if no match?
	// Answer: ???

	// TODO: What does FindStringIndex return if no match?
	// Answer: ???
}

// =============================================================================
// EXERCISE 9: FINDING - ALL MATCHES
// =============================================================================

func PracticeFindingAll() {
	// TODO: Find all matches
	// matches := re.???(s, n)

	// TODO: What does n mean in FindAllString?
	// n < 0: ???
	// n == 0: ???
	// n > 0: ???

	// TODO: Find all match indices
	// indices := re.???(s, n)
}

// =============================================================================
// EXERCISE 10: SUBMATCHES (CAPTURING GROUPS)
// =============================================================================

func PracticeSubmatches() {
	// TODO: Find first match with groups
	// match := re.???(s)

	// TODO: What does FindStringSubmatch return?
	// match[0] = ???
	// match[1] = ???
	// match[2] = ???

	// TODO: Find all matches with groups
	// matches := re.???(s, n)

	// TODO: Find submatch indices
	// indices := re.???(s)

	// TODO: What do indices contain?
	// indices[0:2] = ???
	// indices[2:4] = ???
	// indices[4:6] = ???
}

// =============================================================================
// EXERCISE 11: NAMED GROUPS
// =============================================================================

func PracticeNamedGroups() {
	// TODO: Named group syntax
	// (?P<???>pattern)

	// TODO: Get group names from compiled regex
	// names := re.???()

	// TODO: First element of SubexpNames is always:
	// Answer: ???

	// TODO: Convert match to map (pattern)
	// match := re.FindStringSubmatch(s)
	// for i, name := range re.???() {
	//     if i != 0 && name != "" {
	//         result[name] = match[i]
	//     }
	// }
}

// =============================================================================
// EXERCISE 12: REPLACING
// =============================================================================

func PracticeReplacing() {
	// TODO: Replace all matches
	// result := re.???(src, repl)

	// TODO: Replace with literal (no expansion)
	// result := re.???(src, repl)

	// TODO: Replace using function
	// result := re.???(src, func(s string) string { ... })

	// TODO: Expansion syntax in replacement
	// $0 or $& = ???
	// $1, $2 = ???
	// ${name} = ???
	// $$ = ???
}

// =============================================================================
// EXERCISE 13: SPLITTING
// =============================================================================

func PracticeSplitting() {
	// TODO: Split string on pattern
	// parts := re.???(s, n)

	// TODO: What does n mean in Split?
	// n < 0: ???
	// n == 0: ???
	// n > 0: ???

	// TODO: Difference from strings.Split?
	// Answer: ???
}

// =============================================================================
// EXERCISE 14: LONGEST MATCH
// =============================================================================

func PracticeLongest() {
	// TODO: Make regex prefer longest match
	// re.???()

	// TODO: Default Go behavior is:
	// Answer: ???

	// TODO: Greedy quantifier
	// .*

	// TODO: Non-greedy quantifier
	// ???
}

// =============================================================================
// EXERCISE 15: COMMON PATTERNS
// =============================================================================

func PracticeCommonPatterns() {
	// TODO: Match digits only
	// `???`

	// TODO: Match word characters only
	// `???`

	// TODO: Match email (simple)
	// `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// TODO: Match URL
	// `https?://???`

	// TODO: Match date YYYY-MM-DD
	// `(\d{4})-(\d{2})-(\d{2})`

	// TODO: Match hashtag
	// `#???`

	// TODO: Match @mention
	// `@???`
}

// =============================================================================
// EXERCISE 16: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// for _, s := range data {
	//     matched, _ := regexp.MatchString(pattern, s)
	// }
	// Answer: ???
	// Fix: ???

	// Mistake 2: What's wrong?
	// dotPattern := regexp.MustCompile(".")
	// Answer: ???
	// Fix: ???

	// Mistake 3: What's wrong?
	// pattern := regexp.MustCompile(`\d{3}`)
	// pattern.MatchString("abc123def")  // returns true
	// Answer: ???
	// Fix: ???

	// Mistake 4: What's wrong?
	// match := re.FindStringSubmatch(s)
	// firstGroup := match[0]
	// Answer: ???
	// Fix: ???

	// Mistake 5: What's wrong?
	// match := re.FindStringSubmatch(s)
	// value := match[1]
	// Answer: ???
	// Fix: ???

	// Mistake 6: What's wrong?
	// re := regexp.MustCompile("hello")
	// found := re.MatchString(text)
	// Answer: ???
	// Fix: ???
}

// =============================================================================
// EXERCISE 17: FLAGS
// =============================================================================

func PracticeFlags() {
	// TODO: Case-insensitive flag
	// (??)

	// TODO: Multi-line flag (^ and $ match line boundaries)
	// (??)

	// TODO: Dot matches newline flag
	// (??)

	// TODO: Example: case-insensitive match
	// pattern := `(??)hello`
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// COMPILING
	// 1. Compile with error?
	_ = "regexp.???(pattern)"

	// 2. Compile or panic?
	_ = "regexp.???(pattern)"

	// MATCHING
	// 3. Match string?
	_ = "re.???(s)"

	// FINDING
	// 4. Find first match?
	_ = "re.???(s)"

	// 5. Find all matches?
	_ = "re.???(s, n)"

	// 6. Find with groups?
	_ = "re.???(s)"

	// 7. Find all with groups?
	_ = "re.???(s, n)"

	// NAMED GROUPS
	// 8. Get group names?
	_ = "re.???()"

	// 9. Named group syntax?
	_ = "(?P<name>???)"

	// REPLACING
	// 10. Replace all?
	_ = "re.???(src, repl)"

	// 11. Replace with function?
	_ = "re.???(src, func)"

	// 12. Replace literal (no expansion)?
	_ = "re.???(src, repl)"

	// SPLITTING
	// 13. Split on pattern?
	_ = "re.???(s, n)"

	// SYNTAX
	// 14. Any digit?
	_ = `???`

	// 15. Word character?
	_ = `???`

	// 16. Whitespace?
	_ = `???`

	// 17. Start of string?
	_ = `???`

	// 18. End of string?
	_ = `???`

	// 19. Word boundary?
	_ = `???`

	// 20. Non-greedy quantifier?
	_ = `*? or +? or ???`
}

// =============================================================================
// MINI PROJECT: LOG PARSER
// =============================================================================

func MiniProject() {
	// Build a log file parser that:
	//
	// 1. Parses log lines with format:
	//    [2024-01-15 10:30:45] LEVEL: Message here
	//
	// 2. Extracts timestamp, level, and message
	//
	// 3. Filters by log level (ERROR, WARN, INFO, DEBUG)
	//
	// 4. Extracts IP addresses from messages
	//
	// 5. Redacts email addresses in messages
	//
	// 6. Counts occurrences of each log level
	//
	// Scaffold:
	//
	// type LogEntry struct {
	//     Timestamp string
	//     Level     string
	//     Message   string
	// }
	//
	// var (
	//     logPattern = regexp.MustCompile(
	//         `^\[(?P<timestamp>[^\]]+)\] (?P<level>\w+): (?P<message>.+)$`)
	//     ipPattern    = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	//     emailPattern = regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	// )
	//
	// func ParseLogLine(line string) (*LogEntry, error) {
	//     match := logPattern.FindStringSubmatch(line)
	//     if match == nil {
	//         return nil, errors.New("invalid log format")
	//     }
	//
	//     // Build map from named groups
	//     result := make(map[string]string)
	//     for i, name := range logPattern.SubexpNames() {
	//         if i != 0 && name != "" {
	//             result[name] = match[i]
	//         }
	//     }
	//
	//     return &LogEntry{
	//         Timestamp: result["timestamp"],
	//         Level:     result["level"],
	//         Message:   result["message"],
	//     }, nil
	// }
	//
	// func FilterByLevel(entries []*LogEntry, level string) []*LogEntry {
	//     var filtered []*LogEntry
	//     for _, e := range entries {
	//         if e.Level == level {
	//             filtered = append(filtered, e)
	//         }
	//     }
	//     return filtered
	// }
	//
	// func ExtractIPs(entries []*LogEntry) []string {
	//     seen := make(map[string]bool)
	//     var ips []string
	//
	//     for _, e := range entries {
	//         matches := ipPattern.FindAllString(e.Message, -1)
	//         for _, ip := range matches {
	//             if !seen[ip] {
	//                 seen[ip] = true
	//                 ips = append(ips, ip)
	//             }
	//         }
	//     }
	//     return ips
	// }
	//
	// func RedactEmails(entry *LogEntry) *LogEntry {
	//     redacted := emailPattern.ReplaceAllString(entry.Message, "[REDACTED]")
	//     return &LogEntry{
	//         Timestamp: entry.Timestamp,
	//         Level:     entry.Level,
	//         Message:   redacted,
	//     }
	// }
	//
	// func CountByLevel(entries []*LogEntry) map[string]int {
	//     counts := make(map[string]int)
	//     for _, e := range entries {
	//         counts[e.Level]++
	//     }
	//     return counts
	// }
	//
	// func ParseLogFile(content string) ([]*LogEntry, error) {
	//     lines := strings.Split(content, "\n")
	//     var entries []*LogEntry
	//
	//     for _, line := range lines {
	//         if strings.TrimSpace(line) == "" {
	//             continue
	//         }
	//         entry, err := ParseLogLine(line)
	//         if err != nil {
	//             continue // Skip invalid lines
	//         }
	//         entries = append(entries, entry)
	//     }
	//     return entries, nil
	// }
	//
	// // Test:
	// func main() {
	//     logs := `[2024-01-15 10:30:45] INFO: Server started on 192.168.1.1
	// [2024-01-15 10:30:46] DEBUG: Connection from 10.0.0.1
	// [2024-01-15 10:30:47] ERROR: Failed login for user@example.com
	// [2024-01-15 10:30:48] WARN: High memory usage
	// [2024-01-15 10:30:49] ERROR: Connection timeout from 192.168.1.100`
	//
	//     entries, _ := ParseLogFile(logs)
	//     fmt.Printf("Total entries: %d\n", len(entries))
	//
	//     // Count by level
	//     counts := CountByLevel(entries)
	//     fmt.Printf("Counts: %v\n", counts)
	//
	//     // Filter errors
	//     errors := FilterByLevel(entries, "ERROR")
	//     fmt.Printf("Errors: %d\n", len(errors))
	//
	//     // Extract IPs
	//     ips := ExtractIPs(entries)
	//     fmt.Printf("IPs: %v\n", ips)
	//
	//     // Redact emails
	//     for _, e := range entries {
	//         redacted := RedactEmails(e)
	//         if redacted.Message != e.Message {
	//             fmt.Printf("Redacted: %s\n", redacted.Message)
	//         }
	//     }
	// }
}

// AllPractice is not meant to be called - this file is for reading.
func AllPractice() {
	PracticePhilosophy()
	PracticeCompiling()
	PracticeCharacterClasses()
	PracticeAnchors()
	PracticeQuantifiers()
	PracticeGroups()
	PracticeMatching()
	PracticeFindingSingle()
	PracticeFindingAll()
	PracticeSubmatches()
	PracticeNamedGroups()
	PracticeReplacing()
	PracticeSplitting()
	PracticeLongest()
	PracticeCommonPatterns()
	PracticeMistakes()
	PracticeFlags()
	SelfTest()
	MiniProject()
}
// Package time_practice is your daily practice space for the time package.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against time/time.go
// 3. Note what you missed - focus on those tomorrow
package time_practice

// =============================================================================
// EXERCISE 1: CORE TYPES
// =============================================================================

func PracticeCoreTypes() {
	// TODO: What are the three core types in time package?
	// 1. time.??? - a moment in time
	// 2. time.??? - a length of time
	// 3. time.??? - a timezone

	// TODO: What two clocks does time.Time contain (Go 1.9+)?
	// 1. ??? clock - what time is it
	// 2. ??? clock - how long since (never jumps)
}

// =============================================================================
// EXERCISE 2: CREATING TIMES
// =============================================================================

func PracticeCreating() {
	// TODO: Get current time
	// now := time.???()

	// TODO: Create specific time
	// t := time.???(year, month, day, hour, min, sec, nsec, loc)

	// TODO: Create from Unix timestamp (seconds)
	// t := time.???(seconds, nanoseconds)

	// TODO: Create from milliseconds
	// t := time.???(milliseconds)

	// TODO: What are the two built-in locations?
	// time.??? and time.???

	// TODO: How do you check if time is zero?
	// t.???()
}

// =============================================================================
// EXERCISE 3: EXTRACTING COMPONENTS
// =============================================================================

func PracticeComponents() {
	// TODO: Get individual components
	// t.???()  - year
	// t.???()  - month (returns time.Month)
	// t.???()  - day of month
	// t.???()  - hour 0-23
	// t.???()  - minute 0-59
	// t.???()  - second 0-59
	// t.???()  - weekday (returns time.Weekday)

	// TODO: Get date as three values
	// year, month, day := t.???()

	// TODO: Get time as three values
	// hour, min, sec := t.???()

	// TODO: Get Unix timestamp
	// seconds := t.???()
	// millis := t.???()
	// nanos := t.???()
}

// =============================================================================
// EXERCISE 4: THE REFERENCE TIME
// =============================================================================

func PracticeReferenceTime() {
	// TODO: What is Go's reference time?
	// Mon Jan ??? ???:???:??? MST 2006

	// TODO: Fill in the reference values
	// Month = ??? (January)
	// Day = ???
	// Hour = ??? (24h) or ??? (12h)
	// Minute = ???
	// Second = ???
	// Year = ???
	// Zone offset = ???

	// TODO: What's the mnemonic?
	// 1 2 3 4 5 6 7 means: ???
}

// =============================================================================
// EXERCISE 5: FORMATTING
// =============================================================================

func PracticeFormatting() {
	// TODO: Format a time
	// s := t.???(layout)

	// TODO: Write layouts for:
	// ISO date:        "???-???-???"
	// 24-hour time:    "???:???:???"
	// 12-hour time:    "???:??? ???"
	// Date and time:   "???-???-??? ???:???:???"

	// TODO: Name 4 predefined layout constants
	// time.???  - "2006-01-02T15:04:05Z07:00"
	// time.???  - "2006-01-02 15:04:05"
	// time.???  - "2006-01-02"
	// time.???  - "15:04:05"
}

// =============================================================================
// EXERCISE 6: PARSING
// =============================================================================

func PracticeParsing() {
	// TODO: Parse a time string
	// t, err := time.???(layout, value)

	// TODO: Parse with explicit location
	// t, err := time.???(layout, value, loc)

	// TODO: What happens if layout has no timezone?
	// Answer: ???

	// TODO: Load a timezone
	// loc, err := time.???(zoneName)
}

// =============================================================================
// EXERCISE 7: DURATIONS
// =============================================================================

func PracticeDurations() {
	// TODO: Name the duration constants
	// time.??? = 1 nanosecond
	// time.??? = 1000 nanoseconds
	// time.??? = 1000 microseconds
	// time.??? = 1000 milliseconds
	// time.??? = 60 seconds
	// time.??? = 60 minutes

	// TODO: Why is there no time.Day constant?
	// Answer: ???

	// TODO: Create durations
	// 5 seconds: ??? * time.Second
	// 2.5 hours: ??? * time.Hour + ??? * time.Minute

	// TODO: Parse a duration string
	// d, err := time.???(str)

	// TODO: Get duration components
	// d.???()  - float64 hours
	// d.???()  - float64 seconds
	// d.???() - int64 milliseconds
}

// =============================================================================
// EXERCISE 8: TIME ARITHMETIC
// =============================================================================

func PracticeArithmetic() {
	// TODO: Add a duration to time
	// future := t.???(duration)

	// TODO: Get difference between two times
	// diff := t2.???(t1)  // Returns Duration

	// TODO: Add calendar units
	// next := t.???(years, months, days)

	// TODO: Comparisons
	// t.???(other)  - is t before other?
	// t.???(other)  - is t after other?
	// t.???(other)  - same instant?

	// TODO: Why not use == for time comparison?
	// Answer: ???
}

// =============================================================================
// EXERCISE 9: TIMEZONES
// =============================================================================

func PracticeTimezones() {
	// TODO: Convert time to a location
	// t2 := t.???(loc)

	// TODO: Shortcuts
	// t.???()  - convert to UTC
	// t.???()  - convert to local

	// TODO: Get zone info
	// name, offset := t.???()
}

// =============================================================================
// EXERCISE 10: TIMERS AND TICKERS
// =============================================================================

func PracticeTimersTickers() {
	// TODO: Sleep for a duration
	// time.???(duration)

	// TODO: Create a timer (fires once)
	// timer := time.???(duration)
	// <-timer.???  // Wait for it

	// TODO: Stop a timer
	// stopped := timer.???()

	// TODO: Shorthand for timer channel
	// <-time.???(duration)

	// TODO: Call function after duration
	// time.???(duration, func)

	// TODO: Create a ticker (fires repeatedly)
	// ticker := time.???(duration)
	// for t := range ticker.??? { ... }
	// ticker.???()  // Stop it

	// TODO: Why avoid time.Tick()?
	// Answer: ???
}

// =============================================================================
// EXERCISE 11: COMMON PATTERNS
// =============================================================================

func PracticePatterns() {
	// TODO: Measure execution time
	// start := time.???()
	// // ... work ...
	// elapsed := time.???(start)

	// TODO: Timeout pattern
	// select {
	// case result := <-work:
	//     // use result
	// case <-time.???(duration):
	//     // timeout
	// }

	// TODO: Start of day
	// startOfDay := time.???(
	//     now.Year(), now.Month(), now.Day(),
	//     0, 0, 0, 0, now.Location(),
	// )
}

// =============================================================================
// EXERCISE 12: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// if t1 == t2 { ... }
	// Answer: ???

	// Mistake 2: What's wrong?
	// t.Format("01-02-2006")  // For "2024-03-15"
	// Answer: ???

	// Mistake 3: What's wrong?
	// for t := range time.Tick(1 * time.Second) { ... }
	// Answer: ???

	// Mistake 4: What's wrong?
	// tomorrow := now.Add(24 * time.Hour)
	// Answer: ???

	// Mistake 5: What's wrong?
	// t, _ := time.Parse("2006-01-02 15:04", userInput)
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// 1. Current time?
	_ = "time.???()"

	// 2. Create specific time?
	_ = "time.???(y, m, d, h, min, s, ns, loc)"

	// 3. Format time?
	_ = "t.???(layout)"

	// 4. Parse time?
	_ = "time.???(layout, value)"

	// 5. Reference time?
	_ = "Mon Jan ??? ???:???:??? MST ???"

	// 6. Add duration?
	_ = "t.???(d)"

	// 7. Subtract times?
	_ = "t2.???(t1)"

	// 8. Add calendar units?
	_ = "t.???(y, m, d)"

	// 9. Compare times correctly?
	_ = "t.???(other)"

	// 10. Parse duration?
	_ = "time.???(str)"

	// 11. Create timer?
	_ = "time.???(d)"

	// 12. Create ticker?
	_ = "time.???(d)"

	// 13. Measure elapsed?
	_ = "time.???(start)"

	// 14. Load timezone?
	_ = "time.???(name)"
}

// =============================================================================
// MINI PROJECT: MEETING SCHEDULER
// =============================================================================

func MiniProject() {
	// Build a meeting scheduler that:
	//
	// 1. Parses meeting time from user input (multiple formats)
	// 2. Converts between timezones for participants
	// 3. Checks if meeting is in business hours (9am-5pm)
	// 4. Calculates duration until meeting
	// 5. Sends reminder 15 minutes before
	//
	// Scaffold:
	//
	// type Meeting struct {
	//     Title    string
	//     Time     time.Time
	//     Duration time.Duration
	// }
	//
	// func parseMeetingTime(input, timezone string) (time.Time, error) {
	//     formats := []string{
	//         "2006-01-02 15:04",
	//         "Jan 2, 2006 3:04 PM",
	//         time.RFC3339,
	//     }
	//     loc, err := time.LoadLocation(timezone)
	//     if err != nil {
	//         return time.Time{}, err
	//     }
	//     for _, f := range formats {
	//         if t, err := time.ParseInLocation(f, input, loc); err == nil {
	//             return t, nil
	//         }
	//     }
	//     return time.Time{}, fmt.Errorf("cannot parse: %s", input)
	// }
	//
	// func isBusinessHours(t time.Time) bool {
	//     hour := t.Hour()
	//     weekday := t.Weekday()
	//     return hour >= 9 && hour < 17 &&
	//            weekday != time.Saturday &&
	//            weekday != time.Sunday
	// }
	//
	// func formatForParticipant(t time.Time, timezone string) string {
	//     loc, _ := time.LoadLocation(timezone)
	//     return t.In(loc).Format("Mon, Jan 2 at 3:04 PM MST")
	// }
	//
	// func scheduleReminder(m Meeting) {
	//     reminderTime := m.Time.Add(-15 * time.Minute)
	//     until := time.Until(reminderTime)
	//     if until > 0 {
	//         time.AfterFunc(until, func() {
	//             fmt.Printf("Reminder: %s starts in 15 minutes!\n", m.Title)
	//         })
	//     }
	// }
}

// AllPractice is not meant to be called - this file is for reading.
func AllPractice() {
	PracticeCoreTypes()
	PracticeCreating()
	PracticeComponents()
	PracticeReferenceTime()
	PracticeFormatting()
	PracticeParsing()
	PracticeDurations()
	PracticeArithmetic()
	PracticeTimezones()
	PracticeTimersTickers()
	PracticePatterns()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
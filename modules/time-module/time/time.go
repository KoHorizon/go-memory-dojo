// Package timemod provides comprehensive documentation and working examples
// for Go's time package - dates, times, durations, and timers.
package timemod

import (
	"fmt"
	"time"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF TIME IN GO
// =============================================================================
//
// Go's time package represents moments and durations with clarity:
//
//   time.Time     - A moment in time (instant)
//   time.Duration - A length of time (nanoseconds)
//   time.Location - A timezone
//
// KEY INSIGHT: Time.Time IS timezone-aware. Every Time has a Location.
// Operations preserve the location. This avoids common timezone bugs.
//
// MONOTONIC CLOCKS
//
// Go 1.9+ Time values contain TWO clocks:
//   - Wall clock: What time is it? (can jump due to NTP, daylight saving)
//   - Monotonic clock: How long since? (always moves forward)
//
// time.Now() captures both. Comparisons and durations use monotonic.
// Formatting and timezone operations use wall clock.
//
// =============================================================================

// =============================================================================
// SECTION 2: CREATING TIME VALUES
// =============================================================================
//
// time.Now() - Current time with local timezone
//
// time.Date(year, month, day, hour, min, sec, nsec, loc)
//   Creates a specific time in a specific location
//
// time.Unix(sec, nsec) - From Unix timestamp
// time.UnixMilli(msec) - From milliseconds since epoch
// time.UnixMicro(usec) - From microseconds since epoch
//
// =============================================================================

func DemonstrateCreating() {
	// Current time
	now := time.Now()
	fmt.Printf("Now: %v\n", now)
	fmt.Printf("Location: %v\n", now.Location())

	// Specific date/time
	birthday := time.Date(1990, time.March, 15, 14, 30, 0, 0, time.UTC)
	fmt.Printf("Birthday: %v\n", birthday)

	// Using local timezone
	local := time.Date(2024, time.December, 25, 10, 0, 0, 0, time.Local)
	fmt.Printf("Local Christmas: %v\n", local)

	// From Unix timestamp
	epoch := time.Unix(0, 0)
	fmt.Printf("Unix epoch: %v\n", epoch)

	fromUnix := time.Unix(1700000000, 0)
	fmt.Printf("From Unix: %v\n", fromUnix)

	// From milliseconds
	fromMilli := time.UnixMilli(1700000000000)
	fmt.Printf("From Milli: %v\n", fromMilli)

	// Zero time
	var zero time.Time
	fmt.Printf("Zero time: %v, IsZero: %v\n", zero, zero.IsZero())
}

// =============================================================================
// SECTION 3: EXTRACTING COMPONENTS
// =============================================================================
//
// Time accessors:
//   t.Year()       - 2024
//   t.Month()      - time.December (1-12)
//   t.Day()        - Day of month (1-31)
//   t.Hour()       - 0-23
//   t.Minute()     - 0-59
//   t.Second()     - 0-59
//   t.Nanosecond() - 0-999999999
//
//   t.Weekday()    - time.Sunday through time.Saturday
//   t.YearDay()    - Day of year (1-366)
//   t.ISOWeek()    - ISO 8601 year and week number
//
//   t.Date()       - year, month, day
//   t.Clock()      - hour, min, sec
//
//   t.Unix()       - Seconds since epoch
//   t.UnixMilli()  - Milliseconds since epoch
//   t.UnixNano()   - Nanoseconds since epoch
//
// =============================================================================

func DemonstrateComponents() {
	t := time.Date(2024, time.March, 15, 14, 30, 45, 123456789, time.UTC)

	// Individual components
	fmt.Printf("Year: %d\n", t.Year())
	fmt.Printf("Month: %v (%d)\n", t.Month(), t.Month())
	fmt.Printf("Day: %d\n", t.Day())
	fmt.Printf("Hour: %d\n", t.Hour())
	fmt.Printf("Minute: %d\n", t.Minute())
	fmt.Printf("Second: %d\n", t.Second())
	fmt.Printf("Nanosecond: %d\n", t.Nanosecond())

	// Date info
	fmt.Printf("Weekday: %v\n", t.Weekday())
	fmt.Printf("YearDay: %d\n", t.YearDay())
	year, week := t.ISOWeek()
	fmt.Printf("ISO Week: %d-W%02d\n", year, week)

	// Convenience methods
	y, m, d := t.Date()
	fmt.Printf("Date(): %d-%02d-%02d\n", y, m, d)

	h, min, s := t.Clock()
	fmt.Printf("Clock(): %02d:%02d:%02d\n", h, min, s)

	// Unix timestamps
	fmt.Printf("Unix: %d\n", t.Unix())
	fmt.Printf("UnixMilli: %d\n", t.UnixMilli())
	fmt.Printf("UnixNano: %d\n", t.UnixNano())
}

// =============================================================================
// SECTION 4: FORMATTING - THE REFERENCE TIME
// =============================================================================
//
// THE MAGIC DATE: Mon Jan 2 15:04:05 MST 2006
//
// This is Go's reference time. Each component has a specific value:
//   Month:   January = 1
//   Day:     2
//   Hour:    15 (3 PM in 24h) or 3 (in 12h)
//   Minute:  04
//   Second:  05
//   Year:    2006
//   Zone:    MST (or -0700)
//
// MNEMONIC: 1 2 3 4 5 6 7 (Jan 2, 15:04:05, 2006, -0700)
//
// COMMON FORMATS
//
//   "2006-01-02"                    - ISO date
//   "2006-01-02 15:04:05"           - Date and time
//   "15:04:05"                      - 24-hour time
//   "3:04 PM"                       - 12-hour time
//   "Mon, 02 Jan 2006"              - RFC 822 style
//   "2006-01-02T15:04:05Z07:00"     - RFC 3339
//
// PREDEFINED LAYOUTS (constants in time package)
//
//   time.RFC3339    = "2006-01-02T15:04:05Z07:00"
//   time.RFC822     = "02 Jan 06 15:04 MST"
//   time.Kitchen    = "3:04PM"
//   time.DateTime   = "2006-01-02 15:04:05"
//   time.DateOnly   = "2006-01-02"
//   time.TimeOnly   = "15:04:05"
//
// =============================================================================

func DemonstrateFormatting() {
	t := time.Date(2024, time.March, 15, 14, 30, 45, 0, time.UTC)

	// Custom formats using reference time
	fmt.Printf("ISO Date: %s\n", t.Format("2006-01-02"))
	fmt.Printf("US Date: %s\n", t.Format("01/02/2006"))
	fmt.Printf("EU Date: %s\n", t.Format("02/01/2006"))
	fmt.Printf("Time 24h: %s\n", t.Format("15:04:05"))
	fmt.Printf("Time 12h: %s\n", t.Format("3:04 PM"))
	fmt.Printf("Full: %s\n", t.Format("Monday, January 2, 2006 at 3:04 PM"))

	// Predefined layouts
	fmt.Printf("RFC3339: %s\n", t.Format(time.RFC3339))
	fmt.Printf("RFC822: %s\n", t.Format(time.RFC822))
	fmt.Printf("Kitchen: %s\n", t.Format(time.Kitchen))
	fmt.Printf("DateTime: %s\n", t.Format(time.DateTime))
	fmt.Printf("DateOnly: %s\n", t.Format(time.DateOnly))
	fmt.Printf("TimeOnly: %s\n", t.Format(time.TimeOnly))

	// With timezone
	fmt.Printf("With zone: %s\n", t.Format("2006-01-02 15:04:05 MST"))
	fmt.Printf("With offset: %s\n", t.Format("2006-01-02T15:04:05-07:00"))
}

// =============================================================================
// SECTION 5: PARSING
// =============================================================================
//
// time.Parse(layout, value) (Time, error)
//   Parses using the same reference format as Format()
//   Assumes UTC if no timezone in layout/value
//
// time.ParseInLocation(layout, value, loc) (Time, error)
//   Parses with explicit default location
//
// =============================================================================

func DemonstrateParsing() {
	// Basic parsing
	t1, _ := time.Parse("2006-01-02", "2024-03-15")
	fmt.Printf("Parsed date: %v\n", t1)

	t2, _ := time.Parse("2006-01-02 15:04:05", "2024-03-15 14:30:00")
	fmt.Printf("Parsed datetime: %v\n", t2)

	// Using predefined layouts
	t3, _ := time.Parse(time.RFC3339, "2024-03-15T14:30:00Z")
	fmt.Printf("Parsed RFC3339: %v\n", t3)

	t4, _ := time.Parse(time.Kitchen, "2:30PM")
	fmt.Printf("Parsed Kitchen: %v\n", t4)

	// ParseInLocation - useful for user input without timezone
	loc, _ := time.LoadLocation("America/New_York")
	t5, _ := time.ParseInLocation("2006-01-02 15:04", "2024-03-15 14:30", loc)
	fmt.Printf("Parsed in NY: %v\n", t5)

	// Error handling
	_, err := time.Parse("2006-01-02", "not-a-date")
	fmt.Printf("Parse error: %v\n", err)

	// Common gotcha: wrong layout
	_, err2 := time.Parse("01-02-2006", "2024-03-15") // Wrong!
	fmt.Printf("Wrong layout error: %v\n", err2)
}

// =============================================================================
// SECTION 6: DURATIONS
// =============================================================================
//
// time.Duration is int64 nanoseconds
//
// CONSTANTS
//
//   time.Nanosecond  = 1
//   time.Microsecond = 1000 * Nanosecond
//   time.Millisecond = 1000 * Microsecond
//   time.Second      = 1000 * Millisecond
//   time.Minute      = 60 * Second
//   time.Hour        = 60 * Minute
//
// NOTE: No Day constant (days aren't always 24 hours - DST!)
//
// CREATING DURATIONS
//
//   5 * time.Second
//   time.Duration(n) * time.Millisecond
//   time.ParseDuration("1h30m")
//
// DURATION METHODS
//
//   d.Hours()        - Float64 hours
//   d.Minutes()      - Float64 minutes
//   d.Seconds()      - Float64 seconds
//   d.Milliseconds() - Int64 milliseconds
//   d.Nanoseconds()  - Int64 nanoseconds
//   d.String()       - "1h30m0s"
//   d.Truncate(m)    - Round down to multiple
//   d.Round(m)       - Round to nearest multiple
//
// =============================================================================

func DemonstrateDurations() {
	// Creating durations
	d1 := 5 * time.Second
	d2 := 2*time.Hour + 30*time.Minute
	d3 := time.Duration(1500) * time.Millisecond

	fmt.Printf("5 seconds: %v\n", d1)
	fmt.Printf("2h30m: %v\n", d2)
	fmt.Printf("1500ms: %v\n", d3)

	// Parsing durations
	d4, _ := time.ParseDuration("1h30m45s")
	fmt.Printf("Parsed: %v\n", d4)

	d5, _ := time.ParseDuration("300ms")
	fmt.Printf("Parsed ms: %v\n", d5)

	// Extracting components
	d := 2*time.Hour + 30*time.Minute + 45*time.Second
	fmt.Printf("Hours: %.2f\n", d.Hours())
	fmt.Printf("Minutes: %.2f\n", d.Minutes())
	fmt.Printf("Seconds: %.2f\n", d.Seconds())
	fmt.Printf("Milliseconds: %d\n", d.Milliseconds())

	// Truncate and Round
	d6 := 1*time.Hour + 23*time.Minute + 45*time.Second
	fmt.Printf("Truncate to hour: %v\n", d6.Truncate(time.Hour))
	fmt.Printf("Round to hour: %v\n", d6.Round(time.Hour))
}

// =============================================================================
// SECTION 7: TIME ARITHMETIC
// =============================================================================
//
// ADDING/SUBTRACTING DURATIONS
//
//   t.Add(d Duration) Time      - Add duration
//   t.Sub(u Time) Duration      - Difference between times
//
// ADDING CALENDAR UNITS
//
//   t.AddDate(years, months, days) Time
//   Handles month boundaries correctly
//
// COMPARISONS
//
//   t.Before(u)    - Is t before u?
//   t.After(u)     - Is t after u?
//   t.Equal(u)     - Same instant? (handles timezones correctly)
//   t.Compare(u)   - -1, 0, or 1 (Go 1.20+)
//
// NEVER USE == FOR TIME COMPARISON (different locations!)
//
// =============================================================================

func DemonstrateArithmetic() {
	now := time.Now()

	// Add duration
	future := now.Add(2 * time.Hour)
	fmt.Printf("2 hours from now: %v\n", future.Format(time.Kitchen))

	past := now.Add(-30 * time.Minute)
	fmt.Printf("30 min ago: %v\n", past.Format(time.Kitchen))

	// AddDate for calendar math
	nextYear := now.AddDate(1, 0, 0)
	fmt.Printf("Next year: %v\n", nextYear.Format(time.DateOnly))

	nextMonth := now.AddDate(0, 1, 0)
	fmt.Printf("Next month: %v\n", nextMonth.Format(time.DateOnly))

	// Sub - difference between times
	t1 := time.Date(2024, 3, 15, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 3, 15, 14, 30, 0, 0, time.UTC)
	diff := t2.Sub(t1)
	fmt.Printf("Difference: %v\n", diff)

	// Comparisons
	fmt.Printf("t1.Before(t2): %v\n", t1.Before(t2))
	fmt.Printf("t1.After(t2): %v\n", t1.After(t2))
	fmt.Printf("t1.Equal(t1): %v\n", t1.Equal(t1))

	// Compare (Go 1.20+)
	fmt.Printf("t1.Compare(t2): %d\n", t1.Compare(t2))

	// Why not ==? Different locations!
	utc := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	local := utc.In(time.Local)
	fmt.Printf("== says: %v\n", utc == local)       // May be false!
	fmt.Printf("Equal says: %v\n", utc.Equal(local)) // Always correct
}

// =============================================================================
// SECTION 8: TIMEZONES
// =============================================================================
//
// time.Location represents a timezone
//
// BUILT-IN LOCATIONS
//
//   time.UTC   - Universal Coordinated Time
//   time.Local - System's local timezone
//
// LOADING LOCATIONS
//
//   time.LoadLocation("America/New_York")
//   time.LoadLocation("Europe/London")
//   time.LoadLocation("Asia/Tokyo")
//
// Uses IANA timezone database (tzdata)
//
// CONVERTING
//
//   t.In(loc)   - Convert to location (same instant, different representation)
//   t.UTC()     - Shorthand for t.In(time.UTC)
//   t.Local()   - Shorthand for t.In(time.Local)
//
// =============================================================================

func DemonstrateTimezones() {
	// Create time in UTC
	t := time.Date(2024, 3, 15, 12, 0, 0, 0, time.UTC)
	fmt.Printf("UTC: %v\n", t)

	// Convert to other zones
	nyLoc, _ := time.LoadLocation("America/New_York")
	tokyoLoc, _ := time.LoadLocation("Asia/Tokyo")

	fmt.Printf("New York: %v\n", t.In(nyLoc))
	fmt.Printf("Tokyo: %v\n", t.In(tokyoLoc))
	fmt.Printf("Local: %v\n", t.Local())

	// Get timezone info
	name, offset := t.In(nyLoc).Zone()
	fmt.Printf("Zone: %s, Offset: %d seconds\n", name, offset)

	// Same instant, different representation
	utc := time.Now().UTC()
	local := utc.Local()
	fmt.Printf("UTC: %v\n", utc.Format(time.RFC3339))
	fmt.Printf("Local: %v\n", local.Format(time.RFC3339))
	fmt.Printf("Same instant: %v\n", utc.Equal(local))
}

// =============================================================================
// SECTION 9: TIMERS AND TICKERS
// =============================================================================
//
// time.Sleep(d) - Block for duration
//
// TIMERS - Fire once
//
//   time.NewTimer(d)       - Returns *Timer
//   timer.C                - Channel that receives time
//   timer.Stop()           - Cancel (returns true if stopped before fire)
//   timer.Reset(d)         - Reset to new duration
//
//   time.After(d)          - Shorthand: returns <-chan Time
//   time.AfterFunc(d, f)   - Call f after d
//
// TICKERS - Fire repeatedly
//
//   time.NewTicker(d)      - Returns *Ticker
//   ticker.C               - Channel that receives time
//   ticker.Stop()          - Stop the ticker
//   ticker.Reset(d)        - Change interval
//
//   time.Tick(d)           - Shorthand, but can't be stopped (leaks!)
//
// =============================================================================

func DemonstrateTimers() {
	// time.Sleep
	fmt.Println("Sleeping 100ms...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Awake!")

	// Timer
	timer := time.NewTimer(50 * time.Millisecond)
	<-timer.C
	fmt.Println("Timer fired")

	// Stop timer before it fires
	timer2 := time.NewTimer(1 * time.Second)
	go func() {
		time.Sleep(10 * time.Millisecond)
		stopped := timer2.Stop()
		fmt.Printf("Timer stopped before fire: %v\n", stopped)
	}()
	time.Sleep(50 * time.Millisecond)

	// time.After (shorthand)
	select {
	case <-time.After(50 * time.Millisecond):
		fmt.Println("time.After fired")
	}

	// time.AfterFunc
	done := make(chan bool)
	time.AfterFunc(50*time.Millisecond, func() {
		fmt.Println("AfterFunc called")
		done <- true
	})
	<-done

	// Ticker
	ticker := time.NewTicker(30 * time.Millisecond)
	count := 0
	for t := range ticker.C {
		count++
		fmt.Printf("Tick %d at %v\n", count, t.Format("15:04:05.000"))
		if count >= 3 {
			ticker.Stop()
			break
		}
	}
	fmt.Println("Ticker stopped")
}

// =============================================================================
// SECTION 10: COMMON PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Timeout with select
	done := make(chan bool)
	go func() {
		time.Sleep(50 * time.Millisecond)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Pattern 1: Work completed")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Pattern 1: Timeout")
	}

	// Pattern 2: Rate limiting
	rate := time.Tick(50 * time.Millisecond) // Note: can't stop this
	for i := 0; i < 3; i++ {
		<-rate
		fmt.Printf("Pattern 2: Request %d at %v\n", i, time.Now().Format("15:04:05.000"))
	}

	// Pattern 3: Measure execution time
	start := time.Now()
	time.Sleep(100 * time.Millisecond) // Simulate work
	elapsed := time.Since(start)
	fmt.Printf("Pattern 3: Elapsed: %v\n", elapsed)

	// Pattern 4: Deadline check
	deadline := time.Now().Add(200 * time.Millisecond)
	for time.Now().Before(deadline) {
		// Do work
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Pattern 4: Working...")
	}
	fmt.Println("Pattern 4: Deadline reached")

	// Pattern 5: Format for logging
	log := func(msg string) {
		fmt.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05.000"), msg)
	}
	log("Pattern 5: Log message")

	// Pattern 6: Parse user input date
	parseUserDate := func(input string) (time.Time, error) {
		formats := []string{
			"2006-01-02",
			"01/02/2006",
			"02-01-2006",
			"Jan 2, 2006",
		}
		for _, f := range formats {
			if t, err := time.Parse(f, input); err == nil {
				return t, nil
			}
		}
		return time.Time{}, fmt.Errorf("unknown format: %s", input)
	}
	t, _ := parseUserDate("Mar 15, 2024")
	fmt.Printf("Pattern 6: Parsed: %v\n", t)

	// Pattern 7: Start of day
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	fmt.Printf("Pattern 7: Start of today: %v\n", startOfDay)
}

// =============================================================================
// SECTION 11: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Using == for time comparison
	t1 := time.Now().UTC()
	t2 := t1.Local()
	fmt.Printf("Mistake 1: == gives %v, Equal gives %v\n", t1 == t2, t1.Equal(t2))
	fmt.Println("  Use Equal() or Before()/After()")

	// Mistake 2: Wrong reference time in format
	t := time.Date(2024, 3, 15, 14, 30, 0, 0, time.UTC)
	wrong := t.Format("01-02-2006")     // Wrong month/day order!
	right := t.Format("2006-01-02")     // Correct
	fmt.Printf("Mistake 2: Wrong=%s, Right=%s\n", wrong, right)
	fmt.Println("  Remember: 01=month, 02=day in reference time")

	// Mistake 3: time.Tick leaks
	_ = `
	// WRONG - ticker can never be stopped, leaks goroutine
	for t := range time.Tick(1 * time.Second) {
		process(t)
	}

	// RIGHT - can be stopped
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for t := range ticker.C {
		process(t)
	}
	`
	fmt.Println("Mistake 3: time.Tick can't be stopped, use NewTicker")

	// Mistake 4: Assuming 24 hours in a day
	_ = `
	// WRONG - DST days are 23 or 25 hours
	tomorrow := now.Add(24 * time.Hour)

	// RIGHT - handles DST correctly
	tomorrow := now.AddDate(0, 0, 1)
	`
	fmt.Println("Mistake 4: Use AddDate for calendar days, not Add(24*Hour)")

	// Mistake 5: Not handling parse errors
	_ = `
	// WRONG
	t, _ := time.Parse("2006-01-02", userInput)

	// RIGHT
	t, err := time.Parse("2006-01-02", userInput)
	if err != nil {
		return fmt.Errorf("invalid date: %w", err)
	}
	`
	fmt.Println("Mistake 5: Always check Parse errors")

	// Mistake 6: Forgetting timezone in Parse
	t1parsed, _ := time.Parse("2006-01-02 15:04", "2024-03-15 14:30")
	fmt.Printf("Mistake 6: Parsed without zone: %v (assumes UTC)\n", t1parsed.Location())
	fmt.Println("  Use ParseInLocation for local times")
}

// RunAllDemonstrations is not meant to be called - this file is for reading.
func RunAllDemonstrations() {
	DemonstrateCreating()
	DemonstrateComponents()
	DemonstrateFormatting()
	DemonstrateParsing()
	DemonstrateDurations()
	DemonstrateArithmetic()
	DemonstrateTimezones()
	DemonstrateTimers()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
}
// Package logmod provides comprehensive documentation and working examples
// for Go's log package (standard) and log/slog package (structured, Go 1.21+).
package logmod

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"time"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF GO LOGGING
// =============================================================================
//
// Go provides two logging packages:
//
//   log      - Simple, traditional logging (since Go 1.0)
//   log/slog - Structured logging (since Go 1.21)
//
// WHEN TO USE WHICH
//
//   log:  Simple scripts, quick debugging, legacy code
//   slog: Production applications, structured data, JSON output
//
// BOTH PACKAGES SHARE
//
// - Thread-safe (safe for concurrent use)
// - Configurable output destination
// - Prefix/timestamp support
//
// =============================================================================

// =============================================================================
// SECTION 2: LOG PACKAGE BASICS
// =============================================================================
//
// DEFAULT LOGGER
//
// The log package has a default logger that writes to os.Stderr.
//
// FUNCTIONS
//
//   log.Print(v...)      - Print (like fmt.Print)
//   log.Printf(fmt, v...)- Formatted print
//   log.Println(v...)    - Print with newline
//
//   log.Fatal(v...)      - Print + os.Exit(1)
//   log.Fatalf(fmt, v...)
//   log.Fatalln(v...)
//
//   log.Panic(v...)      - Print + panic()
//   log.Panicf(fmt, v...)
//   log.Panicln(v...)
//
// =============================================================================

func DemonstrateBasicLog() {
	// Basic logging
	log.Print("This is a basic log message")
	log.Println("This is a log message with newline")
	log.Printf("User %s logged in from %s", "alice", "192.168.1.1")

	// Fatal exits the program - use carefully!
	// log.Fatal("This would exit the program")

	// Panic panics - use for unrecoverable errors
	// log.Panic("This would panic")
}

// =============================================================================
// SECTION 3: LOG FLAGS AND PREFIX
// =============================================================================
//
// FLAGS (bitwise OR)
//
//   log.Ldate         - Date: 2009/01/23
//   log.Ltime         - Time: 01:23:23
//   log.Lmicroseconds - Microsecond time: 01:23:23.123123
//   log.Llongfile     - Full file path + line: /a/b/c/d.go:23
//   log.Lshortfile    - File + line: d.go:23
//   log.LUTC          - Use UTC instead of local time
//   log.Lmsgprefix    - Move prefix from start to before message
//   log.LstdFlags     - Ldate | Ltime (default)
//
// FUNCTIONS
//
//   log.SetFlags(flags)  - Set flags for default logger
//   log.Flags() int      - Get current flags
//   log.SetPrefix(s)     - Set prefix for default logger
//   log.Prefix() string  - Get current prefix
//   log.SetOutput(w)     - Set output writer
//   log.Writer() io.Writer
//
// =============================================================================

func DemonstrateFlags() {
	// Save original settings
	originalFlags := log.Flags()
	originalPrefix := log.Prefix()
	defer func() {
		log.SetFlags(originalFlags)
		log.SetPrefix(originalPrefix)
	}()

	// Default flags (date + time)
	log.SetFlags(log.LstdFlags)
	log.Println("Default flags (date + time)")

	// Add microseconds
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("With microseconds")

	// Add short file
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("With short file")

	// Date only
	log.SetFlags(log.Ldate)
	log.Println("Date only")

	// Time only with UTC
	log.SetFlags(log.Ltime | log.LUTC)
	log.Println("Time only (UTC)")

	// No flags
	log.SetFlags(0)
	log.Println("No flags - just message")

	// With prefix
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[APP] ")
	log.Println("Message with prefix")

	// Prefix after timestamp (Lmsgprefix)
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
	log.SetPrefix("[INFO] ")
	log.Println("Prefix before message")
}

// =============================================================================
// SECTION 4: CUSTOM LOGGERS
// =============================================================================
//
// log.New(out io.Writer, prefix string, flag int) *Logger
//   Create a custom logger
//
// Logger has same methods as package-level functions:
//   l.Print, l.Printf, l.Println
//   l.Fatal, l.Fatalf, l.Fatalln
//   l.Panic, l.Panicf, l.Panicln
//   l.SetFlags, l.Flags
//   l.SetPrefix, l.Prefix
//   l.SetOutput, l.Writer
//
// =============================================================================

func DemonstrateCustomLoggers() {
	// Logger to stdout
	stdoutLogger := log.New(os.Stdout, "[STDOUT] ", log.LstdFlags)
	stdoutLogger.Println("This goes to stdout")

	// Logger to stderr with short file
	stderrLogger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	stderrLogger.Println("This goes to stderr with file info")

	// Logger to a buffer (useful for testing)
	var buf bytes.Buffer
	bufLogger := log.New(&buf, "[BUF] ", log.LstdFlags)
	bufLogger.Println("This goes to buffer")
	// buf.String() contains the log output

	// Logger to discard (silent)
	silentLogger := log.New(io.Discard, "", 0)
	silentLogger.Println("This is discarded")

	// Multiple destinations with io.MultiWriter
	multi := io.MultiWriter(os.Stdout, &buf)
	multiLogger := log.New(multi, "[MULTI] ", log.LstdFlags)
	multiLogger.Println("This goes to both stdout and buffer")

	// Different log levels with multiple loggers
	infoLog := log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	errorLog := log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile)

	infoLog.Println("Application started")
	errorLog.Println("Something went wrong")
}

// =============================================================================
// SECTION 5: SLOG BASICS (Go 1.21+)
// =============================================================================
//
// slog provides structured logging with levels and key-value attributes.
//
// LEVELS
//
//   slog.LevelDebug = -4
//   slog.LevelInfo  = 0
//   slog.LevelWarn  = 4
//   slog.LevelError = 8
//
// BASIC FUNCTIONS
//
//   slog.Debug(msg, attrs...)
//   slog.Info(msg, attrs...)
//   slog.Warn(msg, attrs...)
//   slog.Error(msg, attrs...)
//
//   slog.Log(ctx, level, msg, attrs...)
//
// ATTRIBUTES
//
//   slog.String("key", "value")
//   slog.Int("count", 42)
//   slog.Bool("active", true)
//   slog.Duration("elapsed", dur)
//   slog.Time("timestamp", t)
//   slog.Any("data", anyValue)
//   slog.Group("request", attrs...)  // Nested attributes
//
// =============================================================================

func DemonstrateSlogBasics() {
	// Basic logging at different levels
	slog.Debug("Debug message", "detail", "extra info")
	slog.Info("User logged in", "user", "alice", "ip", "192.168.1.1")
	slog.Warn("Cache miss", "key", "user:123")
	slog.Error("Database error", "err", "connection refused", "retry", 3)

	// Using typed attributes
	slog.Info("Request completed",
		slog.String("method", "GET"),
		slog.String("path", "/api/users"),
		slog.Int("status", 200),
		slog.Duration("latency", 42*time.Millisecond),
	)

	// With context
	ctx := context.Background()
	slog.InfoContext(ctx, "With context", "key", "value")

	// Using slog.Log for dynamic level
	level := slog.LevelWarn
	slog.Log(ctx, level, "Dynamic level message")

	// Groups for nested attributes
	slog.Info("HTTP request",
		slog.Group("request",
			slog.String("method", "POST"),
			slog.String("path", "/api/data"),
		),
		slog.Group("response",
			slog.Int("status", 201),
			slog.Int("bytes", 1024),
		),
	)
}

// =============================================================================
// SECTION 6: SLOG HANDLERS
// =============================================================================
//
// HANDLER TYPES
//
//   slog.TextHandler - Human-readable output
//   slog.JSONHandler - JSON output for log aggregation
//
// CREATING HANDLERS
//
//   slog.NewTextHandler(w, opts)
//   slog.NewJSONHandler(w, opts)
//
// HandlerOptions:
//   Level        - Minimum level to log
//   AddSource    - Include file:line
//   ReplaceAttr  - Transform attributes
//
// SETTING DEFAULT HANDLER
//
//   slog.SetDefault(logger)
//
// =============================================================================

func DemonstrateSlogHandlers() {
	// TextHandler - human readable
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	textLogger := slog.New(textHandler)
	textLogger.Info("Text format", "key", "value")

	// JSONHandler - machine readable
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	jsonLogger := slog.New(jsonHandler)
	jsonLogger.Info("JSON format", "key", "value")

	// With source location
	sourceHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})
	sourceLogger := slog.New(sourceHandler)
	sourceLogger.Info("With source location")

	// Set minimum level
	warnHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	})
	warnLogger := slog.New(warnHandler)
	warnLogger.Debug("This won't show") // Below Warn level
	warnLogger.Warn("This will show")

	// Set as default logger
	slog.SetDefault(jsonLogger)
	slog.Info("Now using JSON handler globally")
}

// =============================================================================
// SECTION 7: SLOG LOGGER WITH ATTRIBUTES
// =============================================================================
//
// ADDING PERSISTENT ATTRIBUTES
//
//   logger.With(attrs...) *Logger
//     Returns new logger with additional attributes
//
//   logger.WithGroup(name) *Logger
//     Returns new logger with group prefix
//
// =============================================================================

func DemonstrateSlogWith() {
	handler := slog.NewTextHandler(os.Stdout, nil)
	baseLogger := slog.New(handler)

	// Logger with persistent attributes
	requestLogger := baseLogger.With(
		slog.String("request_id", "abc-123"),
		slog.String("user_id", "user-456"),
	)

	// All messages include request_id and user_id
	requestLogger.Info("Processing started")
	requestLogger.Info("Step 1 complete")
	requestLogger.Error("Step 2 failed", "err", "timeout")

	// Logger with group prefix
	dbLogger := baseLogger.WithGroup("database")
	dbLogger.Info("Connected", "host", "localhost", "port", 5432)
	// Output: database.host=localhost database.port=5432

	// Combine With and WithGroup
	serviceLogger := baseLogger.With("service", "api").WithGroup("metrics")
	serviceLogger.Info("Request", "latency_ms", 42, "status", 200)
}

// =============================================================================
// SECTION 8: SLOG CUSTOM ATTRIBUTES
// =============================================================================
//
// slog.LogValuer INTERFACE
//
//   type LogValuer interface {
//       LogValue() slog.Value
//   }
//
// Implement to control how your types are logged.
// Useful for:
//   - Hiding sensitive data
//   - Custom formatting
//   - Lazy evaluation
//
// =============================================================================

// User implements slog.LogValuer to hide sensitive data
type User struct {
	ID       int
	Name     string
	Email    string
	Password string // Sensitive!
}

func (u User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("id", u.ID),
		slog.String("name", u.Name),
		slog.String("email", u.Email),
		// Password intentionally omitted
	)
}

// Token hides its value in logs
type Token string

func (t Token) LogValue() slog.Value {
	if len(t) > 4 {
		return slog.StringValue(string(t)[:4] + "****")
	}
	return slog.StringValue("****")
}

func DemonstrateSlogCustom() {
	user := User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "secret123",
	}

	token := Token("abcd1234efgh5678")

	slog.Info("User action",
		slog.Any("user", user),   // Password hidden
		slog.Any("token", token), // Shows "abcd****"
	)
}

// =============================================================================
// SECTION 9: SLOG REPLACE ATTRIBUTES
// =============================================================================
//
// ReplaceAttr in HandlerOptions lets you transform attributes.
//
//   func(groups []string, a slog.Attr) slog.Attr
//
// Use cases:
//   - Remove attributes (return empty Attr)
//   - Rename keys
//   - Transform values
//   - Mask sensitive data
//
// =============================================================================

func DemonstrateSlogReplace() {
	// Remove time attribute
	noTimeHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{} // Remove
			}
			return a
		},
	})
	noTimeLogger := slog.New(noTimeHandler)
	noTimeLogger.Info("No timestamp")

	// Rename level to severity
	renameHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				a.Key = "severity"
			}
			return a
		},
	})
	renameLogger := slog.New(renameHandler)
	renameLogger.Info("Level renamed to severity")

	// Custom time format
	customTimeHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := a.Value.Time()
				a.Value = slog.StringValue(t.Format("15:04:05"))
			}
			return a
		},
	})
	customTimeLogger := slog.New(customTimeHandler)
	customTimeLogger.Info("Custom time format")
}

// =============================================================================
// SECTION 10: COMMON PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Package-level loggers
	_ = `
	// logger.go
	package myapp

	var (
		Info  = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
		Error = log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile)
		Debug = log.New(io.Discard, "[DEBUG] ", log.LstdFlags)
	)

	// Enable debug logging
	func EnableDebug() {
		Debug.SetOutput(os.Stdout)
	}
	`

	// Pattern 2: Log to file
	_ = `
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	`

	// Pattern 3: Log to both file and console
	_ = `
	f, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	multi := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multi)
	`

	// Pattern 4: Request ID middleware with slog
	_ = `
	func LoggingMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.New().String()
			logger := slog.With("request_id", requestID)

			// Add to context
			ctx := context.WithValue(r.Context(), "logger", logger)

			logger.Info("Request started",
				"method", r.Method,
				"path", r.URL.Path,
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	`

	// Pattern 5: Conditional logging (slog levels)
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // Change to LevelDebug to see debug logs
	})
	logger := slog.New(handler)
	logger.Debug("This won't show at Info level")
	logger.Info("This will show")

	// Pattern 6: Structured error logging
	err := fmt.Errorf("connection timeout")
	slog.Error("Database operation failed",
		"operation", "query",
		"table", "users",
		"err", err,
	)

	// Pattern 7: Performance logging
	start := time.Now()
	// ... operation ...
	time.Sleep(10 * time.Millisecond)
	slog.Info("Operation completed",
		slog.Duration("elapsed", time.Since(start)),
	)
}

// =============================================================================
// SECTION 11: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Using log.Fatal in library code
	_ = `
	// WRONG - exits the program, prevents caller from handling error
	func Connect() *DB {
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)  // Never in library code!
		}
		return db
	}

	// RIGHT - return error
	func Connect() (*DB, error) {
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			return nil, fmt.Errorf("connect: %w", err)
		}
		return db, nil
	}
	`
	fmt.Println("Mistake 1: Never use log.Fatal in library code")

	// Mistake 2: Forgetting slog attributes come in pairs
	_ = `
	// WRONG - odd number of arguments
	slog.Info("User", "name", "alice", "age")  // age has no value!

	// RIGHT - key-value pairs
	slog.Info("User", "name", "alice", "age", 30)
	`
	fmt.Println("Mistake 2: slog attributes must be key-value pairs")

	// Mistake 3: Not setting up handler for slog
	_ = `
	// Default slog uses text handler to stderr
	// For JSON output, must configure:
	handler := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(handler))
	`
	fmt.Println("Mistake 3: Configure slog handler for JSON output")

	// Mistake 4: Logging sensitive data
	_ = `
	// WRONG
	slog.Info("Login", "user", user.Email, "password", user.Password)

	// RIGHT - implement LogValuer or filter
	slog.Info("Login", "user", user)  // LogValuer hides password
	`
	fmt.Println("Mistake 4: Use LogValuer to hide sensitive data")

	// Mistake 5: Not checking if debug logging is enabled
	_ = `
	// WRONG - expensive even when debug disabled
	slog.Debug("Data dump", "data", expensiveSerialize(data))

	// RIGHT - check level first or use lazy evaluation
	if logger.Enabled(ctx, slog.LevelDebug) {
		slog.Debug("Data dump", "data", expensiveSerialize(data))
	}
	`
	fmt.Println("Mistake 5: Check level before expensive log operations")

	// Mistake 6: Inconsistent log formats
	_ = `
	// WRONG - mixing styles
	log.Printf("user=%s", name)
	slog.Info("User logged in", "user", name)

	// RIGHT - pick one style and stick with it
	`
	fmt.Println("Mistake 6: Be consistent - use log OR slog, not both")
}

// =============================================================================
// SECTION 12: LOG VS SLOG COMPARISON
// =============================================================================
//
// log package:
//   + Simple, no setup required
//   + Familiar printf-style
//   - No structured data
//   - No built-in levels
//   - Hard to parse
//
// slog package:
//   + Structured key-value logging
//   + Built-in levels (Debug/Info/Warn/Error)
//   + JSON output for log aggregation
//   + Type-safe attributes
//   + Context support
//   - More verbose
//   - Requires Go 1.21+
//
// RECOMMENDATION
//
//   Quick scripts/debugging → log
//   Production apps → slog
//
// =============================================================================

func DemonstrateComparison() {
	// log - simple, unstructured
	log.Printf("user %s logged in from %s", "alice", "192.168.1.1")
	// Output: 2024/01/15 10:30:00 user alice logged in from 192.168.1.1

	// slog - structured, parseable
	slog.Info("user logged in", "user", "alice", "ip", "192.168.1.1")
	// Text: time=2024-01-15T10:30:00 level=INFO msg="user logged in" user=alice ip=192.168.1.1
	// JSON: {"time":"2024-01-15T10:30:00","level":"INFO","msg":"user logged in","user":"alice","ip":"192.168.1.1"}
}

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateBasicLog()
	DemonstrateFlags()
	DemonstrateCustomLoggers()
	DemonstrateSlogBasics()
	DemonstrateSlogHandlers()
	DemonstrateSlogWith()
	DemonstrateSlogCustom()
	DemonstrateSlogReplace()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
	DemonstrateComparison()
}
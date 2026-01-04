// Package log_practice is your daily practice space for log and log/slog.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against log/log.go
// 3. Note what you missed - focus on those tomorrow
package log_practice

// =============================================================================
// EXERCISE 1: TWO PACKAGES
// =============================================================================

func PracticeTwoPackages() {
	// TODO: Which package for simple logging?
	// import "???"

	// TODO: Which package for structured logging (Go 1.21+)?
	// import "log/???"

	// TODO: When use log vs slog?
	// log: ???
	// slog: ???
}

// =============================================================================
// EXERCISE 2: LOG BASICS
// =============================================================================

func PracticeLogBasics() {
	// TODO: Basic print functions
	// log.???(v...)       - like fmt.Print
	// log.???(fmt, v...)  - formatted
	// log.???(v...)       - with newline

	// TODO: Print and exit
	// log.???(v...)       - print + os.Exit(1)

	// TODO: Print and panic
	// log.???(v...)       - print + panic()

	// TODO: Where does default logger write?
	// Answer: os.???
}

// =============================================================================
// EXERCISE 3: LOG FLAGS
// =============================================================================

func PracticeLogFlags() {
	// TODO: Name the log flags
	// log.???          - date: 2009/01/23
	// log.???          - time: 01:23:23
	// log.???          - microseconds
	// log.???          - full file path + line
	// log.???          - short file + line
	// log.???          - use UTC time
	// log.???          - prefix before message
	// log.???          - default (Ldate | Ltime)

	// TODO: Set flags
	// log.???(flags)

	// TODO: Get current flags
	// flags := log.???()

	// TODO: Set prefix
	// log.???("[APP] ")

	// TODO: Set output destination
	// log.???(writer)
}

// =============================================================================
// EXERCISE 4: CUSTOM LOGGERS
// =============================================================================

func PracticeCustomLoggers() {
	// TODO: Create a custom logger
	// logger := log.???(output, prefix, flags)

	// TODO: Logger to discard output
	// logger := log.New(io.???, "", 0)

	// TODO: Logger to multiple destinations
	// multi := io.???(os.Stdout, file)
	// logger := log.New(multi, "", log.LstdFlags)
}

// =============================================================================
// EXERCISE 5: SLOG LEVELS
// =============================================================================

func PracticeSlogLevels() {
	// TODO: Name the four slog levels (with values)
	// slog.Level??? = -4
	// slog.Level??? = 0
	// slog.Level??? = 4
	// slog.Level??? = 8

	// TODO: Log at each level
	// slog.???(msg, attrs...)
	// slog.???(msg, attrs...)
	// slog.???(msg, attrs...)
	// slog.???(msg, attrs...)

	// TODO: Log with context
	// slog.???(ctx, msg, attrs...)
}

// =============================================================================
// EXERCISE 6: SLOG ATTRIBUTES
// =============================================================================

func PracticeSlogAttributes() {
	// TODO: Create typed attributes
	// slog.???(key, value)      - string
	// slog.???(key, value)      - int
	// slog.???(key, value)      - bool
	// slog.???(key, value)      - time.Duration
	// slog.???(key, value)      - time.Time
	// slog.???(key, value)      - any type

	// TODO: Group attributes
	// slog.???(name, attrs...)

	// TODO: Simple key-value (inline)
	// slog.Info("msg", "key1", value1, "key2", value2)
}

// =============================================================================
// EXERCISE 7: SLOG HANDLERS
// =============================================================================

func PracticeSlogHandlers() {
	// TODO: Create text handler (human readable)
	// handler := slog.???(writer, opts)

	// TODO: Create JSON handler
	// handler := slog.???(writer, opts)

	// TODO: Create logger from handler
	// logger := slog.???(handler)

	// TODO: Set as default logger
	// slog.???(logger)

	// TODO: HandlerOptions fields
	// &slog.HandlerOptions{
	//     ???:  slog.LevelInfo,  // minimum level
	//     ???:  true,            // include file:line
	//     ???:  func(...),      // transform attrs
	// }
}

// =============================================================================
// EXERCISE 8: SLOG WITH
// =============================================================================

func PracticeSlogWith() {
	// TODO: Add persistent attributes
	// newLogger := logger.???(attrs...)

	// TODO: Add group prefix
	// newLogger := logger.???(name)

	// TODO: Example: request logger
	// reqLogger := baseLogger.???(
	//     slog.String("request_id", id),
	//     slog.String("user_id", uid),
	// )
}

// =============================================================================
// EXERCISE 9: SLOG LOGVALUER
// =============================================================================

func PracticeSlogLogValuer() {
	// TODO: Interface for custom log formatting
	// type ??? interface {
	//     ???() slog.Value
	// }

	// TODO: Create a group value
	// slog.???(attrs...)

	// TODO: Create a string value
	// slog.???(str)

	// TODO: Why implement LogValuer?
	// 1. ???
	// 2. ???
}

// =============================================================================
// EXERCISE 10: SLOG REPLACE ATTR
// =============================================================================

func PracticeSlogReplaceAttr() {
	// TODO: ReplaceAttr function signature
	// func(??? []string, ??? slog.Attr) slog.Attr

	// TODO: Built-in attribute keys
	// slog.???  - time key
	// slog.???  - level key
	// slog.???  - message key
	// slog.???  - source key

	// TODO: Remove an attribute
	// return slog.???{}

	// TODO: Rename level to severity
	// if a.Key == slog.LevelKey {
	//     a.??? = "severity"
	// }
}

// =============================================================================
// EXERCISE 11: COMMON PATTERNS
// =============================================================================

func PracticePatterns() {
	// TODO: Log to file
	// f, _ := os.???(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// log.???(f)

	// TODO: Log to file AND console
	// multi := io.???(os.Stdout, file)
	// log.SetOutput(multi)

	// TODO: Conditional debug logging
	// if logger.???(ctx, slog.LevelDebug) {
	//     slog.Debug(...)
	// }

	// TODO: Performance logging
	// start := time.Now()
	// // ... work ...
	// slog.Info("done", slog.???(key, time.Since(start)))
}

// =============================================================================
// EXERCISE 12: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// func LibraryFunc() {
	//     log.Fatal("error occurred")
	// }
	// Answer: ???

	// Mistake 2: What's wrong?
	// slog.Info("User", "name", "alice", "age")
	// Answer: ???

	// Mistake 3: What's wrong?
	// slog.Info("Login", "password", user.Password)
	// Answer: ???

	// Mistake 4: What's wrong?
	// slog.Debug("Data", "dump", veryExpensiveFunc())
	// Answer: ???

	// Mistake 5: What's wrong?
	// log.Printf("user=%s", name)
	// slog.Info("action", "user", name)
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// LOG PACKAGE
	// 1. Print + exit?
	_ = "log.???()"

	// 2. Print + panic?
	_ = "log.???()"

	// 3. Set flags?
	_ = "log.???(flags)"

	// 4. Set prefix?
	_ = "log.???(str)"

	// 5. Set output?
	_ = "log.???(writer)"

	// 6. Create custom logger?
	_ = "log.???(out, prefix, flags)"

	// 7. Default flags constant?
	_ = "log.???"

	// SLOG PACKAGE
	// 8. Four levels?
	_ = "slog.Level???, Level???, Level???, Level???"

	// 9. Create text handler?
	_ = "slog.???(w, opts)"

	// 10. Create JSON handler?
	_ = "slog.???(w, opts)"

	// 11. Create logger from handler?
	_ = "slog.???(handler)"

	// 12. Set default logger?
	_ = "slog.???(logger)"

	// 13. Add persistent attrs?
	_ = "logger.???(attrs...)"

	// 14. Add group prefix?
	_ = "logger.???(name)"

	// 15. String attribute?
	_ = "slog.???(key, value)"

	// 16. Duration attribute?
	_ = "slog.???(key, dur)"

	// 17. LogValuer interface method?
	_ = "???() slog.Value"
}

// =============================================================================
// MINI PROJECT: STRUCTURED LOGGER
// =============================================================================

func MiniProject() {
	// Build a structured logging system that:
	//
	// 1. Supports multiple output formats (text, JSON)
	// 2. Has configurable minimum level
	// 3. Adds common fields (service name, version)
	// 4. Provides request-scoped logging
	// 5. Masks sensitive fields
	//
	// Scaffold:
	//
	// type LogConfig struct {
	//     Service   string
	//     Version   string
	//     Level     slog.Level
	//     Format    string  // "text" or "json"
	//     AddSource bool
	// }
	//
	// func NewLogger(cfg LogConfig) *slog.Logger {
	//     opts := &slog.HandlerOptions{
	//         Level:     cfg.Level,
	//         AddSource: cfg.AddSource,
	//     }
	//
	//     var handler slog.Handler
	//     if cfg.Format == "json" {
	//         handler = slog.NewJSONHandler(os.Stdout, opts)
	//     } else {
	//         handler = slog.NewTextHandler(os.Stdout, opts)
	//     }
	//
	//     return slog.New(handler).With(
	//         slog.String("service", cfg.Service),
	//         slog.String("version", cfg.Version),
	//     )
	// }
	//
	// func RequestLogger(base *slog.Logger, requestID, userID string) *slog.Logger {
	//     return base.With(
	//         slog.String("request_id", requestID),
	//         slog.String("user_id", userID),
	//     )
	// }
	//
	// // Sensitive data masking
	// type Password string
	//
	// func (p Password) LogValue() slog.Value {
	//     return slog.StringValue("[REDACTED]")
	// }
	//
	// type CreditCard struct {
	//     Number string
	//     CVV    string
	// }
	//
	// func (cc CreditCard) LogValue() slog.Value {
	//     masked := "****"
	//     if len(cc.Number) > 4 {
	//         masked = "****" + cc.Number[len(cc.Number)-4:]
	//     }
	//     return slog.GroupValue(
	//         slog.String("number", masked),
	//         // CVV intentionally omitted
	//     )
	// }
	//
	// Usage:
	// logger := NewLogger(LogConfig{
	//     Service: "api",
	//     Version: "1.0.0",
	//     Level:   slog.LevelInfo,
	//     Format:  "json",
	// })
	//
	// reqLogger := RequestLogger(logger, "req-123", "user-456")
	// reqLogger.Info("Processing payment",
	//     slog.Any("card", card),  // Masked automatically
	// )
}

// AllPractice is not meant to be called - this file is for reading.
func AllPractice() {
	PracticeTwoPackages()
	PracticeLogBasics()
	PracticeLogFlags()
	PracticeCustomLoggers()
	PracticeSlogLevels()
	PracticeSlogAttributes()
	PracticeSlogHandlers()
	PracticeSlogWith()
	PracticeSlogLogValuer()
	PracticeSlogReplaceAttr()
	PracticePatterns()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
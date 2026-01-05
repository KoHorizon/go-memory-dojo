# Go Memory Trainer

A hands-on learning project designed to build muscle memory for Go programming through daily practice and deep understanding of core packages.

## Philosophy

1. **Understand deeply** - Don't just memorize syntax. Understand *why* things work the way they do.
2. **Practice daily** - Repetition builds neural pathways. Each morning, fill in the practice files from memory.
3. **Build incrementally** - Start with fundamentals, then layer on complexity.

## Project Structure

```
go-memory-trainer/
├── cmd/
│   └── main.go                     # CLI for project info
└── modules/
    ├── strings-module/             # fmt package (formatting)
    ├── strutil-module/             # strings package (manipulation)
    ├── bytes-module/               # bytes package
    ├── http-module/                # net/http package
    ├── errors-module/              # error handling
    ├── bufio-module/               # buffered I/O
    ├── context-module/             # context package
    ├── time-module/                # time package
    ├── concurrency-module/         # goroutines, channels, sync
    ├── json-module/                # encoding/json
    ├── io-module/                  # io package
    ├── testing-module/             # testing package
    ├── slices-module/              # slices/maps/cmp packages
    ├── filepath-module/            # path/filepath package
    ├── log-module/                 # log and log/slog packages
    ├── os-module/                  # os package
    ├── regexp-module/              # regexp package
    └── interfaces-module/          # interface concepts
```

## How to Use

### 1. Study the Theory File
Open the theory file (e.g., `modules/errors-module/errors/errors.go`) in your editor. Read the code and comments carefully. The examples show real patterns you'll use daily.

### 2. Practice from Memory
Open the practice file (e.g., `modules/errors-module/errors_practice/practice.go`). Fill in all the TODOs **without looking at the theory file**.

### 3. Check Your Answers
Compare your answers to the theory file. Note what you missed - focus on those tomorrow.

## CLI Commands

```bash
# Show usage information
go run cmd/main.go

# List all modules
go run cmd/main.go --list

# Show info for a specific module
go run cmd/main.go --info errors
```

## Modules

| Module | Package | Description | Status |
|--------|---------|-------------|--------|
| strings | `fmt` | Format verbs, flags, width/precision, print family | ✅ |
| strutil | `strings` | String manipulation, Builder, Reader, searching, splitting | ✅ |
| bytes | `bytes` | Byte slice manipulation, Buffer, Reader, comparison | ✅ |
| http | `net/http` | Client, server, handlers, routing, middleware | ✅ |
| errors | `errors` | Wrapping, Is/As, sentinel errors, custom types | ✅ |
| bufio | `bufio` | Buffered I/O, Scanner, Reader, Writer | ✅ |
| context | `context` | Cancellation, timeouts, deadlines, values | ✅ |
| time | `time` | Time, Duration, formatting, parsing, timers | ✅ |
| concurrency | `goroutines/channels/sync` | Goroutines, channels, select, WaitGroup, Mutex | ✅ |
| json | `encoding/json` | Marshal, unmarshal, struct tags, streaming | ✅ |
| io | `io` | Reader, Writer, Closer, utilities, Pipe | ✅ |
| testing | `testing` | Unit tests, table tests, benchmarks, examples, fuzzing | ✅ |
| slices | `slices/maps/cmp` | Generic slice/map operations, searching, sorting | ✅ |
| filepath | `path/filepath` | OS file paths, Join, Split, Walk, Glob, Clean | ✅ |
| log | `log/slog` | Traditional and structured logging, levels, handlers | ✅ |
| os | `os` | Files, directories, env vars, process, signals | ✅ |
| regexp | `regexp` | Regular expressions, matching, finding, replacing | ✅ |
| interfaces | `interface` | Interface concepts, type assertions, composition, design | ✅ |

## Module Categories

### Data & Text
- **strings** (fmt) - Formatting and printing
- **strutil** (strings) - String manipulation
- **bytes** - Byte slice operations
- **json** - JSON encoding/decoding
- **regexp** - Pattern matching

### I/O & Files
- **io** - Core I/O interfaces
- **bufio** - Buffered I/O
- **os** - OS operations, files, env
- **filepath** - Path manipulation

### Concurrency & Control
- **concurrency** - Goroutines, channels, sync
- **context** - Cancellation and timeouts
- **time** - Time operations

### Web & Network
- **http** - HTTP client and server

### Collections
- **slices** - Slice/map operations (Go 1.21+)

### Quality
- **errors** - Error handling patterns
- **testing** - Testing and benchmarking
- **log** - Logging (traditional and structured)

### Concepts
- **interfaces** - Go's interface system

## Tips for Effective Practice

1. **Morning sessions** - Your brain consolidates learning during sleep. Practice first thing.

2. **Don't cheat** - The struggle to remember is where learning happens. Make your best guess first.

3. **Type, don't copy-paste** - The physical act of typing engages motor memory.

4. **Understand before memorizing** - If you don't understand WHY, you'll forget. Go back to the theory.

5. **Track progress** - Keep a log of what you got right/wrong. You'll see improvement.

## Suggested Learning Order

If you're new to Go, follow this progression:

1. **interfaces** - Understand Go's core abstraction first
2. **errors** - Error handling is fundamental
3. **strings** (fmt) - Formatting basics
4. **strutil** (strings) - String manipulation
5. **bytes** - Working with byte slices
6. **io** - Core I/O patterns
7. **bufio** - Buffered I/O
8. **os** - File and OS operations
9. **filepath** - Path handling
10. **json** - Data serialization
11. **time** - Time operations
12. **context** - Request-scoped data and cancellation
13. **http** - Web programming
14. **concurrency** - Goroutines and channels
15. **slices** - Modern collection utilities
16. **regexp** - Pattern matching
17. **testing** - Testing your code
18. **log** - Observability

## License

MIT
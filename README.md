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
    ├── strings-module/             # fmt package
    │   ├── strings/strings.go      # Theory - read and study
    │   └── strings_practice/       # Practice - fill in TODOs
    ├── http-module/                # net/http package
    ├── errors-module/              # error handling
    ├── bufio-module/               # buffered I/O
    └── ...
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
| http | `net/http` | Client, server, handlers, routing, middleware | ✅ |
| errors | `errors` | Wrapping, Is/As, sentinel errors, custom types | ✅ |
| bufio | `bufio` | Buffered I/O, Scanner, Reader, Writer | ✅ |
| context | `context` | Cancellation, timeouts, deadlines, values | ✅ |
| time | `time` | Time, Duration, formatting, parsing, timers | ✅ |
| concurrency | `goroutines/channels/sync` | Goroutines, channels, select, WaitGroup, Mutex | ✅ |
| json | `encoding/json` | Marshal, unmarshal, struct tags, streaming | ✅ |
| io | `io` | Reader, Writer, Closer, utilities, Pipe | ✅ |
| testing | `testing` | Unit tests, table tests, benchmarks, examples, fuzzing | ✅ |

## Tips for Effective Practice

1. **Morning sessions** - Your brain consolidates learning during sleep. Practice first thing.

2. **Don't cheat** - The struggle to remember is where learning happens. Make your best guess first.

3. **Type, don't copy-paste** - The physical act of typing engages motor memory.

4. **Understand before memorizing** - If you don't understand WHY, you'll forget. Go back to the theory.

5. **Track progress** - Keep a log of what you got right/wrong. You'll see improvement.

## License

MIT
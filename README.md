# Go Memory Trainer

A hands-on learning project designed to build muscle memory for Go programming through
daily practice and deep understanding of core packages and concepts.

## Philosophy

This project follows a simple but effective learning methodology:

1. **Understand deeply** - Don't just memorize syntax. Understand *why* things work the way they do.
2. **Practice daily** - Repetition builds neural pathways. Each morning, fill in the practice files from memory.
3. **Build incrementally** - Start with fundamentals, then layer on complexity.

## Project Structure

```
go-memory-trainer/
├── go.mod
├── README.md
├── cmd/
│   └── main.go                    # Entry point to run examples
└── modules/
    ├── strings-module/            # fmt package - formatting and printing
    │   ├── strings/               # Theory + working examples
    │   │   └── strings.go
    │   └── strings_practice/      # Your daily practice space
    │       └── practice.go
    ├── http-module/               # (future) net/http package
    ├── bufio-module/              # (future) buffered I/O
    ├── concurrency-module/        # (future) goroutines, channels, sync
    └── ...
```

## How to Use This Project

### Learning Phase
1. Read through the theory file (`strings/strings.go`) carefully
2. Run the examples: `go run cmd/main.go`
3. Make sure you understand not just WHAT but WHY

### Daily Practice Phase
1. Open `strings_practice/practice.go`
2. Fill in all the TODOs from memory (no peeking!)
3. Run to verify: `go run cmd/main.go --practice`
4. Check your answers against the theory file
5. Note what you got wrong - focus on those tomorrow

### Expanding the Project
As you master each module, add new ones:
- Copy the module structure
- Write theory with deep explanations
- Create corresponding practice exercises

## Modules

| Module | Package Focus | Status |
|--------|---------------|--------|
| strings-module | `fmt` - Format verbs, flags, printing | ✅ Complete |
| http-module | `net/http` - HTTP client/server, routing, middleware | ✅ Complete |
| errors-module | `errors` - Error handling, wrapping, Is/As | 🔜 Planned |
| time-module | `time` - Parsing, formatting, durations | 🔜 Planned |
| bufio-module | `bufio` - Buffered I/O | 🔜 Planned |
| concurrency-module | Goroutines, channels, sync | 🔜 Planned |
| io-module | `io` - Reader/Writer interfaces | 🔜 Planned |
| json-module | `encoding/json` - JSON marshaling | 🔜 Planned |
| testing-module | `testing` - Unit tests, benchmarks | 🔜 Planned |
| context-module | `context` - Cancellation, timeouts | 🔜 Planned |
| slices-module | `slices` + `maps` - Generic collections (Go 1.21+) | 🔜 Planned |

## Running the Project

```bash
# Run all theory examples
go run cmd/main.go

# Run practice mode (your filled-in exercises)
go run cmd/main.go --practice

# Run a specific module
go run cmd/main.go --module strings
```

## Tips for Effective Practice

1. **Morning sessions work best** - Your brain consolidates learning during sleep.
   Practice first thing to reinforce those neural pathways.

2. **Don't cheat** - The struggle of trying to remember is where learning happens.
   If you can't remember, make your best guess, then check.

3. **Track your progress** - Keep a simple log of what you got right/wrong each day.
   You'll see improvement over time.

4. **Understand before memorizing** - If you don't understand WHY something works,
   you'll forget it. Go back to the theory.

5. **Type, don't copy-paste** - The physical act of typing engages motor memory.
   This is deliberate - it's called "muscle memory" for a reason.

## License

MIT - Use this for your own learning!
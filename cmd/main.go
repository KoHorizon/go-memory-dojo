// Command main provides information about the Go Memory Trainer project.
//
// This is a learning project - the modules are meant to be READ, not executed.
// Open the theory files to study patterns, then fill in the practice files.
package main

import (
	"flag"
	"fmt"
	"os"
)

var modules = []struct {
	name        string
	pkg         string
	description string
	theoryPath  string
	practicePath string
	status      string
}{
	{
		name:        "strings",
		pkg:         "fmt",
		description: "Format verbs, flags, width/precision, print family",
		theoryPath:  "modules/strings-module/strings/strings.go",
		practicePath: "modules/strings-module/strings_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "http",
		pkg:         "net/http",
		description: "Client, server, handlers, routing, middleware, context",
		theoryPath:  "modules/http-module/http/http.go",
		practicePath: "modules/http-module/http_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "errors",
		pkg:         "errors",
		description: "Wrapping, Is/As, sentinel errors, custom types, patterns",
		theoryPath:  "modules/errors-module/errors/errors.go",
		practicePath: "modules/errors-module/errors_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "bufio",
		pkg:         "bufio",
		description: "Buffered I/O, Scanner, Reader, Writer, split functions",
		theoryPath:  "modules/bufio-module/bufio/bufio.go",
		practicePath: "modules/bufio-module/bufio_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "context",
		pkg:         "context",
		description: "Cancellation, timeouts, deadlines, values, propagation",
		theoryPath:  "modules/context-module/context/context.go",
		practicePath: "modules/context-module/context_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "time",
		pkg:         "time",
		description: "Time, Duration, formatting, parsing, timers, tickers",
		theoryPath:  "modules/time-module/time/time.go",
		practicePath: "modules/time-module/time_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "concurrency",
		pkg:         "goroutines/channels/sync",
		description: "Goroutines, channels, select, WaitGroup, Mutex, atomic",
		theoryPath:  "modules/concurrency-module/concurrency/concurrency.go",
		practicePath: "modules/concurrency-module/concurrency_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "json",
		pkg:         "encoding/json",
		description: "Marshal, unmarshal, struct tags, streaming, RawMessage",
		theoryPath:  "modules/json-module/json/json.go",
		practicePath: "modules/json-module/json_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "io",
		pkg:         "io",
		description: "Reader, Writer, Closer, utilities, Pipe, seeking",
		theoryPath:  "modules/io-module/io/io.go",
		practicePath: "modules/io-module/io_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "testing",
		pkg:         "testing",
		description: "Unit tests, table tests, benchmarks, examples, fuzzing",
		theoryPath:  "modules/testing-module/testing/testing.go",
		practicePath: "modules/testing-module/testing_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "slices",
		pkg:         "slices/maps/cmp",
		description: "Generic slice/map operations, searching, sorting, comparison",
		theoryPath:  "modules/slices-module/slices/slices.go",
		practicePath: "modules/slices-module/slices_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "filepath",
		pkg:         "path/filepath",
		description: "OS file paths, Join, Split, Walk, Glob, Clean",
		theoryPath:  "modules/filepath-module/filepath/filepath.go",
		practicePath: "modules/filepath-module/filepath_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "log",
		pkg:         "log/slog",
		description: "Traditional and structured logging, levels, handlers",
		theoryPath:  "modules/log-module/log/log.go",
		practicePath: "modules/log-module/log_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "os",
		pkg:         "os",
		description: "Files, directories, env vars, process, signals",
		theoryPath:  "modules/os-module/os/os.go",
		practicePath: "modules/os-module/os_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "regexp",
		pkg:         "regexp",
		description: "Regular expressions, matching, finding, replacing",
		theoryPath:  "modules/regexp-module/regexp/regexp.go",
		practicePath: "modules/regexp-module/regexp_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "interfaces",
		pkg:         "interface",
		description: "Interface concepts, type assertions, composition, design",
		theoryPath:  "modules/interfaces-module/interfaces/interfaces.go",
		practicePath: "modules/interfaces-module/interfaces_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "strutil",
		pkg:         "strings",
		description: "String manipulation, Builder, Reader, searching, splitting",
		theoryPath:  "modules/strutil-module/strutil/strutil.go",
		practicePath: "modules/strutil-module/strutil_practice/practice.go",
		status:      "✅",
	},
	{
		name:        "bytes",
		pkg:         "bytes",
		description: "Byte slice manipulation, Buffer, Reader, comparison",
		theoryPath:  "modules/bytes-module/bytes/bytes.go",
		practicePath: "modules/bytes-module/bytes_practice/practice.go",
		status:      "✅",
	},
}

func main() {
	list := flag.Bool("list", false, "List all modules")
	info := flag.String("info", "", "Show info for a specific module")
	flag.Parse()

	if *list {
		listModules()
		return
	}

	if *info != "" {
		showModuleInfo(*info)
		return
	}

	showUsage()
}

func showUsage() {
	fmt.Println(`
╔═══════════════════════════════════════════════════════════════╗
║                    GO MEMORY TRAINER                          ║
╚═══════════════════════════════════════════════════════════════╝

A hands-on project for building Go muscle memory through daily practice.

HOW TO USE THIS PROJECT:

  1. STUDY the theory file - read the code and comments
     The examples show real patterns you'll use daily.

  2. PRACTICE each morning - open the practice file
     Fill in the TODOs from memory (no peeking!)

  3. CHECK your answers against the theory file
     Note what you missed, focus on those tomorrow.

COMMANDS:

  go run cmd/main.go --list           List all modules
  go run cmd/main.go --info strings   Show details for a module

FILES:

  Each module has two files:
  • theory file    - Documented examples to study
  • practice file  - TODOs to fill in from memory

The goal is to make these patterns automatic - muscle memory, not conscious recall.
`)
}

func listModules() {
	fmt.Println("\n AVAILABLE MODULES\n")
	fmt.Printf(" %-12s %-24s %s\n", "MODULE", "PACKAGE", "STATUS")
	fmt.Printf(" %-12s %-24s %s\n", "------", "-------", "------")
	
	for _, m := range modules {
		fmt.Printf(" %-12s %-24s %s\n", m.name, m.pkg, m.status)
	}
	
	fmt.Println("\n Use --info <module> for details")
	fmt.Println()
}

func showModuleInfo(name string) {
	for _, m := range modules {
		if m.name == name {
			fmt.Printf("\n MODULE: %s\n", m.name)
			fmt.Printf(" Package: %s\n", m.pkg)
			fmt.Printf(" Status: %s\n", m.status)
			fmt.Printf(" Description: %s\n", m.description)
			
			if m.theoryPath != "" {
				fmt.Println("\n Files:")
				fmt.Printf("   Theory:   %s\n", m.theoryPath)
				fmt.Printf("   Practice: %s\n", m.practicePath)
				fmt.Println("\n Workflow:")
				fmt.Println("   1. Open and study the theory file")
				fmt.Println("   2. Open the practice file, fill in TODOs from memory")
				fmt.Println("   3. Compare your answers to the theory file")
			} else {
				fmt.Println("\n Coming soon!")
			}
			fmt.Println()
			return
		}
	}
	
	fmt.Printf("\n Unknown module: %s\n", name)
	fmt.Println(" Use --list to see available modules\n")
	os.Exit(1)
}
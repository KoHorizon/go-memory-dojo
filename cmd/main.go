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
		name:        "time",
		pkg:         "time",
		description: "Parsing, formatting, durations, timers, tickers",
		status:      "🔜",
	},
	{
		name:        "context",
		pkg:         "context",
		description: "Cancellation, timeouts, values, propagation",
		status:      "✅",
	},
	{
		name:        "concurrency",
		pkg:         "goroutines/channels/sync",
		description: "Goroutines, channels, select, sync primitives",
		status:      "🔜",
	},
	{
		name:        "json",
		pkg:         "encoding/json",
		description: "Marshal, unmarshal, struct tags, streaming",
		status:      "🔜",
	},
	{
		name:        "io",
		pkg:         "io",
		description: "Reader, Writer, interfaces, composition",
		status:      "🔜",
	},
	{
		name:        "testing",
		pkg:         "testing",
		description: "Unit tests, table tests, benchmarks, examples",
		status:      "🔜",
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
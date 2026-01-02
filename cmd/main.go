// Command main provides the entry point for running the Go Memory Trainer.
//
// Usage:
//
//	go run cmd/main.go            # Run theory examples
//	go run cmd/main.go --practice # Run practice mode
//	go run cmd/main.go --module strings  # Run specific module
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KoHorizon/go-memory-dojo/modules/strings-module/strings"
	"github.com/KoHorizon/go-memory-dojo/modules/strings-module/strings_practice"
	"github.com/KoHorizon/go-memory-dojo/modules/http-module/http"
	"github.com/KoHorizon/go-memory-dojo/modules/http-module/http_practice"
)

/*
 * Available modules :
 * "strings"
 * "http"
 */
var moduleToRun = ""

func main() {
	// Define flags
	practice := flag.Bool("practice", false, "Run practice mode instead of theory examples")
	module := flag.String("module", "", "Run a specific module (e.g., 'strings')")
	list := flag.Bool("list", false, "List all available modules")

	flag.Parse()

	// List modules
	if *list {
		listModules()
		return
	}

	// Determine which module to run
	targetModule := *module
	if targetModule == "" {
		targetModule = moduleToRun // Default to strings module
	}

	// Run the appropriate mode
	switch targetModule {
	case "strings":
		runStringsModule(*practice)
	case "http":
		runHTTPModule(*practice)
	default:
		fmt.Printf("Unknown module: %s\n", targetModule)
		fmt.Println("Use --list to see available modules")
		os.Exit(1)
	}
}

func runHTTPModule(practiceMode bool) {
	if practiceMode {
		printHeader("PRACTICE MODE: net/http Package")
		fmt.Println("Fill in the TODOs in http_practice/practice.go")
		fmt.Println("Then run this to check your work!")
		fmt.Println()
		http_practice.RunAllPractice()
	} else {
		printHeader("THEORY MODE: net/http Package")
		fmt.Println("Study these examples and understand WHY they work.")
		fmt.Println("When ready, test yourself with: go run cmd/main.go --module http --practice")
		fmt.Println()
		http.RunAllDemonstrations()
	}
}

func runStringsModule(practiceMode bool) {
	if practiceMode {
		printHeader("PRACTICE MODE: fmt Package")
		fmt.Println("Fill in the TODOs in strings_practice/practice.go")
		fmt.Println("Then run this to check your work!")
		fmt.Println()
		strings_practice.RunAllPractice()
	} else {
		printHeader("THEORY MODE: fmt Package")
		fmt.Println("Study these examples and understand WHY they work.")
		fmt.Println("When ready, test yourself with: go run cmd/main.go --practice")
		fmt.Println()
		strings.RunAllDemonstrations()
	}
}

func listModules() {
	fmt.Println("Available Modules:")
	fmt.Println()
	fmt.Println("  strings    - fmt package: format verbs, flags, printing         ✅")
	fmt.Println("  http       - net/http: HTTP client/server, routing, middleware  ✅")
	fmt.Println("  bufio      - (coming soon) buffered I/O")
	fmt.Println("  concurrency - (coming soon) goroutines, channels, sync")
	fmt.Println("  errors     - (coming soon) error handling, wrapping, Is/As")
	fmt.Println("  time       - (coming soon) parsing, formatting, durations")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run cmd/main.go --module strings          # Theory mode")
	fmt.Println("  go run cmd/main.go --module strings --practice  # Practice mode")
	fmt.Println("  go run cmd/main.go --module http             # HTTP theory")
	fmt.Println("  go run cmd/main.go --module http --practice  # HTTP practice")
}

func printHeader(title string) {
	line := "========================================================"
	fmt.Println(line)
	fmt.Printf("  %s\n", title)
	fmt.Println(line)
	fmt.Println()
}
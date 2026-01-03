// Package bufio_practice is your daily practice space for the bufio package.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against bufio/bufio.go
// 3. Note what you missed - focus on those tomorrow
package bufio_practice

import (
	"bufio"
	"bytes"
	"strings"
)

// =============================================================================
// EXERCISE 1: WHY BUFFERED I/O
// =============================================================================

func PracticeWhyBuffering() {
	// TODO: What problem does buffered I/O solve?
	// Answer: ???

	// TODO: What does bufio.Reader do?
	// Answer: ???

	// TODO: What does bufio.Writer do?
	// Answer: ???

	// TODO: When should you NOT use bufio?
	// Answer: ???
}

// =============================================================================
// EXERCISE 2: CREATING READERS AND WRITERS
// =============================================================================

func PracticeCreation() {
	var input strings.Reader
	var output bytes.Buffer

	// TODO: Create a buffered reader with default buffer size
	// r := bufio.???(???)
	_ = &input

	// TODO: Create a buffered reader with 64KB buffer
	// r := bufio.???(???, ???)

	// TODO: Create a buffered writer with default buffer size
	// w := bufio.???(???)
	_ = &output

	// TODO: Create a buffered writer with 64KB buffer
	// w := bufio.???(???, ???)
}

// =============================================================================
// EXERCISE 3: READER METHODS
// =============================================================================

func PracticeReaderMethods() {
	r := bufio.NewReader(strings.NewReader("hello\nworld\n"))

	// TODO: Read until newline, returns string (includes delimiter!)
	// line, err := r.???(???)
	_ = r

	// TODO: Read until newline, returns []byte (includes delimiter!)
	// data, err := r.???(???)

	// TODO: Read a single byte
	// b, err := r.???()

	// TODO: Look at next 5 bytes without consuming them
	// peek, err := r.???(???)

	// TODO: How many bytes are in the buffer?
	// n := r.???()

	// QUESTION: Does ReadString('\n') include the newline?
	// Answer: ???
}

// =============================================================================
// EXERCISE 4: SCANNER BASICS
// =============================================================================

func PracticeScannerBasics() {
	data := "line one\nline two\nline three"

	// TODO: Create a scanner
	// scanner := bufio.???(strings.NewReader(data))
	_ = data

	// TODO: Write the standard scan loop
	// for scanner.???() {
	//     line := scanner.???()  // Get string
	//     // or
	//     data := scanner.???()  // Get []byte
	// }
	// if err := scanner.???(); err != nil {
	//     // handle error
	// }

	// QUESTION: Does scanner.Text() include the delimiter?
	// Answer: ???

	// QUESTION: What's the difference between Text() and Bytes()?
	// Answer: ???
}

// =============================================================================
// EXERCISE 5: SCANNER SPLIT FUNCTIONS
// =============================================================================

func PracticeSplitFunctions() {
	// TODO: Name the 4 built-in split functions
	// 1. bufio.Scan??? - splits on newlines (default)
	// 2. bufio.Scan??? - splits on whitespace
	// 3. bufio.Scan??? - one byte at a time
	// 4. bufio.Scan??? - one UTF-8 rune at a time

	// TODO: Set a scanner to split on words
	// scanner := bufio.NewScanner(reader)
	// scanner.???(bufio.???)

	// TODO: What is the default max token size?
	// Answer: ???

	// TODO: How do you handle lines larger than 64KB?
	// buf := make([]byte, ???)
	// scanner.???(buf, ???)
}

// =============================================================================
// EXERCISE 6: CUSTOM SPLIT FUNCTION
// =============================================================================

func PracticeCustomSplit() {
	// TODO: Write the split function signature
	// type SplitFunc func(??? []byte, ??? bool) (??? int, ??? []byte, ??? error)

	// TODO: What do the return values mean?
	// advance: ???
	// token: ???
	// err: ???

	// TODO: What do you return to request more data?
	// return ???, ???, ???

	// TODO: Write a split function for comma-separated values
	// splitOnComma := func(data []byte, atEOF bool) (int, []byte, error) {
	//     ???
	// }
}

// =============================================================================
// EXERCISE 7: WRITER METHODS
// =============================================================================

func PracticeWriterMethods() {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	// TODO: Write a string (efficient, no conversion)
	// w.???(???)

	// TODO: Write a single byte
	// w.???(???)

	// TODO: Write a byte slice
	// w.???(???)

	// TODO: CRITICAL - push buffer to underlying writer
	// w.???()

	// TODO: How many bytes are waiting in the buffer?
	// n := w.???()

	// TODO: How much space is left in the buffer?
	// n := w.???()

	// TODO: Reuse this writer with a different destination
	// var newBuf bytes.Buffer
	// w.???(???)

	_ = w
}

// =============================================================================
// EXERCISE 8: THE FLUSH REQUIREMENT
// =============================================================================

func PracticeFlush() {
	// TODO: What happens if you don't call Flush()?
	// Answer: ???

	// TODO: When does the buffer auto-flush?
	// Answer: ???

	// TODO: Write the common pattern with defer
	// w := bufio.NewWriter(file)
	// defer w.???()
	// ... write operations
}

// =============================================================================
// EXERCISE 9: ReadWriter
// =============================================================================

func PracticeReadWriter() {
	// TODO: What is bufio.ReadWriter?
	// Answer: ???

	// TODO: Create a ReadWriter
	// rw := bufio.???(reader, writer)

	// TODO: When would you use ReadWriter?
	// Answer: ???
}

// =============================================================================
// EXERCISE 10: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// var buf bytes.Buffer
	// w := bufio.NewWriter(&buf)
	// w.WriteString("data")
	// // program exits
	// Answer: ???

	// Mistake 2: What's wrong?
	// scanner := bufio.NewScanner(reader)
	// for scanner.Scan() {
	//     process(scanner.Text())
	// }
	// // done
	// Answer: ???

	// Mistake 3: What's wrong?
	// r := bufio.NewReader(input)
	// line, _ := r.ReadString('\n')
	// if line == "hello" { ... }  // Never matches!
	// Answer: ???

	// Mistake 4: What's wrong?
	// scanner := bufio.NewScanner(reader)
	// var saved []byte
	// scanner.Scan()
	// saved = scanner.Bytes()
	// scanner.Scan()
	// use(saved)  // Bug!
	// Answer: ???

	// Mistake 5: What's wrong?
	// scanner := bufio.NewScanner(reader)  // reader has 100KB lines
	// for scanner.Scan() { ... }
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// 1. Create a buffered reader?
	_ = "bufio.???(reader)"

	// 2. Create a buffered writer?
	_ = "bufio.???(writer)"

	// 3. Read until delimiter (string)?
	_ = "r.???(delim)"

	// 4. Read until delimiter (bytes)?
	_ = "r.???(delim)"

	// 5. Look ahead without consuming?
	_ = "r.???(n)"

	// 6. Scanner iteration pattern?
	_ = "for scanner.???() { scanner.???() }"

	// 7. Check scanner error?
	_ = "scanner.???()"

	// 8. Set scanner split function?
	_ = "scanner.???(bufio.???)"

	// 9. Write buffer to underlying writer?
	_ = "w.???()"

	// 10. Reuse writer with new destination?
	_ = "w.???(newWriter)"
}

// =============================================================================
// MINI PROJECT: LOG PROCESSOR
// =============================================================================

func MiniProject() {
	// Build a log processor that:
	//
	// 1. Reads log lines from input
	// 2. Filters lines containing "ERROR"
	// 3. Writes matching lines to output
	// 4. Counts total lines and error lines
	//
	// Requirements:
	// - Use bufio.Scanner for reading
	// - Use bufio.Writer for writing
	// - Don't forget to Flush!
	// - Check scanner.Err()
	//
	// Scaffold:
	//
	// func processLogs(input io.Reader, output io.Writer) (total, errors int, err error) {
	//     scanner := bufio.NewScanner(input)
	//     writer := bufio.NewWriter(output)
	//     defer writer.Flush()
	//
	//     for scanner.Scan() {
	//         total++
	//         line := scanner.Text()
	//         if strings.Contains(line, "ERROR") {
	//             errors++
	//             writer.WriteString(line + "\n")
	//         }
	//     }
	//     return total, errors, scanner.Err()
	// }
}

// RunAllPractice is not meant to be called - this file is for reading.
func RunAllPractice() {
	PracticeWhyBuffering()
	PracticeCreation()
	PracticeReaderMethods()
	PracticeScannerBasics()
	PracticeSplitFunctions()
	PracticeCustomSplit()
	PracticeWriterMethods()
	PracticeFlush()
	PracticeReadWriter()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
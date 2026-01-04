// Package bufio provides comprehensive documentation and working examples
// for Go's bufio package - buffered I/O for efficient reading and writing.
package bufio

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// =============================================================================
// SECTION 1: WHY BUFFERED I/O EXISTS
// =============================================================================
//
// THE PROBLEM WITH UNBUFFERED I/O
//
// Every Read() or Write() call to the OS is expensive - it involves:
//   - Context switch from user space to kernel space
//   - System call overhead
//   - Potential disk/network latency
//
// Reading a file byte-by-byte means thousands of system calls.
// Writing a log line-by-line means thousands more.
//
// THE SOLUTION: BUFFERING
//
// Instead of reading 1 byte at a time from disk:
//   Read 4KB into memory, then serve bytes from that buffer.
//
// Instead of writing 1 line at a time to disk:
//   Accumulate lines in memory, then write 4KB at once.
//
// bufio wraps any io.Reader or io.Writer with a buffer:
//   - bufio.Reader: Reads large chunks, serves small pieces
//   - bufio.Writer: Accumulates small writes, flushes large chunks
//   - bufio.Scanner: High-level line/word/token reading
//
// WHEN TO USE BUFIO
//
// Use it when:
//   - Reading files line by line
//   - Processing streams token by token
//   - Writing many small pieces (logs, output)
//   - Wrapping network connections
//
// Don't need it when:
//   - Reading entire file at once (os.ReadFile)
//   - Using already-buffered types (bytes.Buffer)
//   - Single large read/write operations
//
// =============================================================================

// =============================================================================
// SECTION 2: bufio.Reader BASICS
// =============================================================================
//
// CREATING A READER
//
//   r := bufio.NewReader(file)           // Default 4KB buffer
//   r := bufio.NewReaderSize(file, 64*1024)  // Custom 64KB buffer
//
// KEY METHODS
//
//   Read(p []byte) (n int, err error)  - Standard io.Reader
//   ReadByte() (byte, error)           - Single byte
//   ReadBytes(delim byte) ([]byte, error)  - Until delimiter (includes it)
//   ReadString(delim byte) (string, error) - Until delimiter (includes it)
//   ReadLine() ([]byte, bool, error)   - Single line (avoid, use Scanner)
//   Peek(n int) ([]byte, error)        - Look ahead without consuming
//   Buffered() int                      - Bytes available in buffer
//
// THE DELIMITER IS INCLUDED
//
// ReadBytes('\n') and ReadString('\n') include the newline.
// You often need to trim it: strings.TrimSuffix(line, "\n")
//
// =============================================================================

func DemonstrateReaderBasics() {
	data := "line one\nline two\nline three"
	r := bufio.NewReader(strings.NewReader(data))

	// ReadString - reads until delimiter (includes delimiter)
	line1, _ := r.ReadString('\n')
	fmt.Printf("ReadString: %q\n", line1) // "line one\n"

	// ReadBytes - same but returns []byte
	line2, _ := r.ReadBytes('\n')
	fmt.Printf("ReadBytes: %q\n", line2) // "line two\n"

	// Last line has no newline - ReadString returns io.EOF
	line3, err := r.ReadString('\n')
	fmt.Printf("Last line: %q, err: %v\n", line3, err) // "line three", EOF

	// Peek - look ahead without consuming
	r2 := bufio.NewReader(strings.NewReader("Hello, World"))
	peeked, _ := r2.Peek(5)
	fmt.Printf("Peeked: %q\n", peeked) // "Hello"

	// Data is still there
	full, _ := io.ReadAll(r2)
	fmt.Printf("Full: %q\n", full) // "Hello, World"

	// ReadByte - single byte at a time
	r3 := bufio.NewReader(strings.NewReader("ABC"))
	b1, _ := r3.ReadByte()
	b2, _ := r3.ReadByte()
	fmt.Printf("Bytes: %c, %c\n", b1, b2) // A, B

	// Buffered - how much is in the buffer
	r4 := bufio.NewReader(strings.NewReader("0123456789"))
	r4.Peek(1) // Force a read into buffer
	fmt.Printf("Buffered: %d bytes\n", r4.Buffered())
}

// =============================================================================
// SECTION 3: bufio.Scanner - THE HIGH-LEVEL API
// =============================================================================
//
// WHAT IS SCANNER?
//
// Scanner provides a convenient interface for reading delimited data.
// It handles buffering, delimiter detection, and iteration cleanly.
//
// THE SCAN PATTERN
//
//   scanner := bufio.NewScanner(reader)
//   for scanner.Scan() {
//       line := scanner.Text()  // or scanner.Bytes()
//       // process line
//   }
//   if err := scanner.Err(); err != nil {
//       // handle error
//   }
//
// BUILT-IN SPLIT FUNCTIONS
//
//   ScanLines  - Split on \n (default) - most common
//   ScanWords  - Split on whitespace
//   ScanBytes  - One byte at a time
//   ScanRunes  - One UTF-8 rune at a time
//
// Set with: scanner.Split(bufio.ScanWords)
//
// KEY DIFFERENCE FROM ReadString
//
// Scanner does NOT include the delimiter in the result.
// scanner.Text() returns "line one", not "line one\n"
//
// BUFFER SIZE
//
// Default max token size is 64KB. For larger lines:
//   buf := make([]byte, 1024*1024)  // 1MB
//   scanner.Buffer(buf, cap(buf))
//
// =============================================================================

func DemonstrateScanner() {
	// Basic line scanning (most common use)
	data := "line one\nline two\nline three"
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		fmt.Printf("Line: %q\n", scanner.Text()) // No newline included!
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Word scanning
	words := "  hello   world   foo   bar  "
	scanner2 := bufio.NewScanner(strings.NewReader(words))
	scanner2.Split(bufio.ScanWords)

	var wordList []string
	for scanner2.Scan() {
		wordList = append(wordList, scanner2.Text())
	}
	fmt.Printf("Words: %v\n", wordList) // [hello world foo bar]

	// Byte scanning
	scanner3 := bufio.NewScanner(strings.NewReader("ABC"))
	scanner3.Split(bufio.ScanBytes)

	for scanner3.Scan() {
		fmt.Printf("Byte: %q\n", scanner3.Text())
	}

	// Rune scanning (UTF-8 aware)
	scanner4 := bufio.NewScanner(strings.NewReader("Hello, 世界"))
	scanner4.Split(bufio.ScanRunes)

	var runes []string
	for scanner4.Scan() {
		runes = append(runes, scanner4.Text())
	}
	fmt.Printf("Runes: %v\n", runes)

	// scanner.Bytes() vs scanner.Text()
	// Bytes() returns the underlying slice (no allocation, but may change)
	// Text() returns a new string (safe to keep, allocates)
}

// =============================================================================
// SECTION 4: CUSTOM SPLIT FUNCTIONS
// =============================================================================
//
// SPLIT FUNCTION SIGNATURE
//
//   type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
//
// Parameters:
//   data  - The unprocessed data in the buffer
//   atEOF - True if no more data is coming
//
// Returns:
//   advance - How many bytes to consume from data
//   token   - The token to return (nil means need more data)
//   err     - Any error (usually nil or ErrFinalToken)
//
// COMMON PATTERNS
//
// Return (0, nil, nil) to request more data
// Return (n, token, nil) to return a token and advance n bytes
// Return (0, nil, ErrFinalToken) to stop scanning
//
// =============================================================================

func DemonstrateCustomSplit() {
	// Split on comma
	splitOnComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, ','); i >= 0 {
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil // Request more data
	}

	csv := "one,two,three,four"
	scanner := bufio.NewScanner(strings.NewReader(csv))
	scanner.Split(splitOnComma)

	for scanner.Scan() {
		fmt.Printf("CSV field: %q\n", scanner.Text())
	}

	// Split on double newline (paragraphs)
	splitParagraphs := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
			return i + 2, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}

	text := "First paragraph here.\n\nSecond paragraph.\n\nThird one."
	scanner2 := bufio.NewScanner(strings.NewReader(text))
	scanner2.Split(splitParagraphs)

	for scanner2.Scan() {
		fmt.Printf("Paragraph: %q\n", scanner2.Text())
	}
}

// =============================================================================
// SECTION 5: bufio.Writer BASICS
// =============================================================================
//
// CREATING A WRITER
//
//   w := bufio.NewWriter(file)              // Default 4KB buffer
//   w := bufio.NewWriterSize(file, 64*1024) // Custom 64KB buffer
//
// KEY METHODS
//
//   Write(p []byte) (n int, err error)   - Standard io.Writer
//   WriteByte(c byte) error              - Single byte
//   WriteString(s string) (int, error)   - String (more efficient)
//   WriteRune(r rune) (int, error)       - Single rune
//   Flush() error                        - CRITICAL: Write buffer to underlying writer
//   Buffered() int                       - Bytes waiting in buffer
//   Available() int                      - Space left in buffer
//   Reset(w io.Writer)                   - Reuse buffer with new writer
//
// CRITICAL: YOU MUST FLUSH
//
// Data stays in the buffer until:
//   - Buffer is full (auto-flush)
//   - You call Flush()
//   - You call Reset()
//
// If you don't flush, data is LOST when the program exits!
//
// COMMON PATTERN
//
//   w := bufio.NewWriter(file)
//   defer w.Flush()  // Don't forget!
//   // ... write operations
//
// =============================================================================

func DemonstrateWriterBasics() {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	// WriteString - most common for text
	w.WriteString("Hello, ")
	w.WriteString("World!\n")

	// WriteByte - single byte
	w.WriteByte('!')

	// Write - byte slice
	w.Write([]byte(" More data."))

	// Check buffer state
	fmt.Printf("Buffered: %d, Available: %d\n", w.Buffered(), w.Available())

	// Nothing in underlying buffer yet!
	fmt.Printf("Before flush: %q\n", buf.String()) // ""

	// FLUSH - critical!
	w.Flush()
	fmt.Printf("After flush: %q\n", buf.String()) // "Hello, World!\n! More data."

	// Reset to reuse buffer with different writer
	var buf2 bytes.Buffer
	w.Reset(&buf2)
	w.WriteString("New destination")
	w.Flush()
	fmt.Printf("New buffer: %q\n", buf2.String())
}

// =============================================================================
// SECTION 6: bufio.ReadWriter
// =============================================================================
//
// WHAT IS ReadWriter?
//
// Combines a Reader and Writer into a single struct.
// Useful for bidirectional streams (network connections, pipes).
//
//   type ReadWriter struct {
//       *Reader
//       *Writer
//   }
//
// Create with:
//   rw := bufio.NewReadWriter(reader, writer)
//
// =============================================================================

func DemonstrateReadWriter() {
	// Simulating a bidirectional connection
	var readBuf bytes.Buffer
	var writeBuf bytes.Buffer

	readBuf.WriteString("incoming data")

	rw := bufio.NewReadWriter(
		bufio.NewReader(&readBuf),
		bufio.NewWriter(&writeBuf),
	)

	// Read from input
	line, _ := rw.ReadString('\n')
	fmt.Printf("Read: %q\n", line)

	// Write to output
	rw.WriteString("response data\n")
	rw.Flush()
	fmt.Printf("Written: %q\n", writeBuf.String())
}

// =============================================================================
// SECTION 7: REAL-WORLD PATTERNS
// =============================================================================
//
// These patterns appear constantly in production Go code.
//
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Process file line by line
	processLines := func(r io.Reader) error {
		scanner := bufio.NewScanner(r)
		lineNum := 0
		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			_ = line // process line
		}
		return scanner.Err()
	}

	data := "line1\nline2\nline3"
	processLines(strings.NewReader(data))
	fmt.Println("Pattern 1: Line processing done")

	// Pattern 2: Buffered file writing
	writeLines := func(w io.Writer, lines []string) error {
		bw := bufio.NewWriter(w)
		defer bw.Flush() // Don't forget!

		for _, line := range lines {
			if _, err := bw.WriteString(line + "\n"); err != nil {
				return err
			}
		}
		return nil
	}

	var buf bytes.Buffer
	writeLines(&buf, []string{"a", "b", "c"})
	fmt.Printf("Pattern 2: Written %q\n", buf.String())

	// Pattern 3: Read until specific marker
	readUntilMarker := func(r *bufio.Reader, marker string) (string, error) {
		var result strings.Builder
		for {
			line, err := r.ReadString('\n')
			if strings.TrimSpace(line) == marker {
				break
			}
			result.WriteString(line)
			if err != nil {
				return result.String(), err
			}
		}
		return result.String(), nil
	}

	input := "header1\nheader2\n---\nbody\n"
	content, _ := readUntilMarker(bufio.NewReader(strings.NewReader(input)), "---")
	fmt.Printf("Pattern 3: Before marker: %q\n", content)

	// Pattern 4: Peek to detect format
	detectFormat := func(r *bufio.Reader) string {
		peek, _ := r.Peek(1)
		if len(peek) == 0 {
			return "empty"
		}
		switch peek[0] {
		case '{':
			return "json"
		case '<':
			return "xml"
		default:
			return "unknown"
		}
	}

	jsonReader := bufio.NewReader(strings.NewReader(`{"key": "value"}`))
	xmlReader := bufio.NewReader(strings.NewReader(`<root></root>`))
	fmt.Printf("Pattern 4: JSON=%s, XML=%s\n", detectFormat(jsonReader), detectFormat(xmlReader))

	// Pattern 5: Word counting
	countWords := func(r io.Reader) int {
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanWords)
		count := 0
		for scanner.Scan() {
			count++
		}
		return count
	}

	text := "  one two   three four five  "
	fmt.Printf("Pattern 5: Word count = %d\n", countWords(strings.NewReader(text)))

	// Pattern 6: Large line handling
	readLargeLines := func(r io.Reader) error {
		scanner := bufio.NewScanner(r)

		// Allow lines up to 1MB
		buf := make([]byte, 64*1024)
		scanner.Buffer(buf, 1024*1024)

		for scanner.Scan() {
			_ = scanner.Text()
		}
		return scanner.Err()
	}
	_ = readLargeLines
	fmt.Println("Pattern 6: Large line support configured")
}

// =============================================================================
// SECTION 8: COMMON MISTAKES
// =============================================================================
//
// 1. Forgetting to Flush Writer
// 2. Not checking Scanner.Err()
// 3. Assuming ReadString excludes delimiter
// 4. Keeping scanner.Bytes() reference (it changes!)
// 5. Buffer too small for large lines
//
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Forgetting to flush
	var buf1 bytes.Buffer
	w := bufio.NewWriter(&buf1)
	w.WriteString("data")
	// w.Flush() <- MISSING! Data is lost!
	fmt.Printf("Mistake 1 - No flush, data lost: %q\n", buf1.String())
	w.Flush() // Fix
	fmt.Printf("After flush: %q\n", buf1.String())

	// Mistake 2: Not checking scanner error
	scanner := bufio.NewScanner(strings.NewReader("line1\nline2"))
	for scanner.Scan() {
		_ = scanner.Text()
	}
	// if err := scanner.Err(); err != nil { } <- Don't forget!
	fmt.Printf("Mistake 2 - Always check: scanner.Err() = %v\n", scanner.Err())

	// Mistake 3: Delimiter IS included in ReadString/ReadBytes
	r := bufio.NewReader(strings.NewReader("hello\nworld"))
	line, _ := r.ReadString('\n')
	fmt.Printf("Mistake 3 - Delimiter included: %q (len=%d)\n", line, len(line))
	// Fix: strings.TrimSuffix(line, "\n")

	// Mistake 4: scanner.Bytes() slice is reused
	scanner2 := bufio.NewScanner(strings.NewReader("one\ntwo"))
	var saved []byte
	scanner2.Scan()
	saved = scanner2.Bytes() // Points to internal buffer!
	scanner2.Scan()          // Buffer contents changed!
	fmt.Printf("Mistake 4 - Bytes() changed: %q (expected 'one')\n", saved)
	// Fix: Use scanner.Text() or copy the slice

	// Mistake 5: Token too long
	longLine := strings.Repeat("x", 100000) // 100KB line
	scanner3 := bufio.NewScanner(strings.NewReader(longLine))
	// Default buffer is 64KB - this will fail!
	if !scanner3.Scan() {
		fmt.Printf("Mistake 5 - Token too long: %v\n", scanner3.Err())
	}
	// Fix: scanner.Buffer(buf, maxSize)
}

// =============================================================================
// SECTION 9: PERFORMANCE CONSIDERATIONS
// =============================================================================
//
// BUFFER SIZE MATTERS
//
// Default 4KB is good for most cases.
// Larger buffers help with high-throughput sequential I/O.
// Smaller buffers use less memory.
//
// SCANNER VS READER
//
// Scanner: Simpler API, copies data to new strings
// Reader: More control, can avoid allocations with Bytes()
//
// REUSING BUFFERS
//
// Use Reset() to reuse bufio.Writer with new destinations.
// Avoids allocation overhead in loops.
//
// =============================================================================

func DemonstratePerformance() {
	// Reusing writer buffer
	w := bufio.NewWriter(nil)

	for i := 0; i < 3; i++ {
		var buf bytes.Buffer
		w.Reset(&buf) // Reuse the same bufio.Writer
		fmt.Fprintf(w, "iteration %d", i)
		w.Flush()
		fmt.Printf("Reused buffer: %q\n", buf.String())
	}

	// Custom buffer size for large files
	// r := bufio.NewReaderSize(file, 256*1024) // 256KB buffer

	// Avoiding allocations with ReadSlice (advanced)
	r := bufio.NewReader(strings.NewReader("line1\nline2\n"))
	slice, _ := r.ReadSlice('\n') // Returns slice into buffer, no allocation
	fmt.Printf("ReadSlice (no alloc): %q\n", slice)
	// Warning: slice is only valid until next read!
}

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateReaderBasics()
	DemonstrateScanner()
	DemonstrateCustomSplit()
	DemonstrateWriterBasics()
	DemonstrateReadWriter()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
	DemonstratePerformance()
}
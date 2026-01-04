// Package io provides comprehensive documentation and working examples
// for Go's io package - the foundation of all I/O in Go.
package io

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF IO
// =============================================================================
//
// THE BRILLIANCE OF GO'S IO MODEL
//
// Go's io package defines tiny interfaces that compose beautifully:
//
//   io.Reader - Anything you can read from
//   io.Writer - Anything you can write to
//
// This abstraction is incredibly powerful because:
//   - Files, network connections, buffers, HTTP bodies - all the same interface
//   - Functions accept interfaces, not concrete types
//   - Easy to test (pass a buffer instead of a file)
//   - Easy to compose (chain readers/writers)
//
// THE CORE INTERFACES
//
//   type Reader interface {
//       Read(p []byte) (n int, err error)
//   }
//
//   type Writer interface {
//       Write(p []byte) (n int, err error)
//   }
//
//   type Closer interface {
//       Close() error
//   }
//
// COMPOSITION INTERFACES
//
//   ReadWriter    = Reader + Writer
//   ReadCloser    = Reader + Closer
//   WriteCloser   = Writer + Closer
//   ReadWriteCloser = Reader + Writer + Closer
//
// =============================================================================

// =============================================================================
// SECTION 2: io.Reader
// =============================================================================
//
// THE READ CONTRACT
//
//   Read(p []byte) (n int, err error)
//
// - Reads UP TO len(p) bytes into p
// - Returns number of bytes read (n) and any error
// - n may be less than len(p) even without error
// - Returns io.EOF when no more data
// - io.EOF is NOT an error - it's the normal end signal
//
// COMMON READERS
//
//   *os.File           - Files
//   *bytes.Buffer      - In-memory buffer
//   *bytes.Reader      - Read-only in-memory
//   *strings.Reader    - Read from string
//   http.Response.Body - HTTP response
//   net.Conn           - Network connection
//
// =============================================================================

func DemonstrateReader() {
	// strings.Reader - read from a string
	r := strings.NewReader("Hello, World!")

	buf := make([]byte, 5)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Printf("Read %d bytes: %q\n", n, buf[:n])
		}
		if err == io.EOF {
			fmt.Println("EOF reached")
			break
		}
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
	}

	// bytes.Buffer - read and write
	var buf2 bytes.Buffer
	buf2.WriteString("Buffer content")
	data := make([]byte, 6)
	n, _ := buf2.Read(data)
	fmt.Printf("From buffer: %q\n", data[:n])

	// bytes.Reader - read-only, supports Seek
	br := bytes.NewReader([]byte("Seekable"))
	br.Seek(4, io.SeekStart) // Jump to position 4
	rest, _ := io.ReadAll(br)
	fmt.Printf("After seek: %q\n", rest)
}

// =============================================================================
// SECTION 3: io.Writer
// =============================================================================
//
// THE WRITE CONTRACT
//
//   Write(p []byte) (n int, err error)
//
// - Writes len(p) bytes from p
// - Returns number of bytes written
// - If n < len(p), err must be non-nil
// - Must not modify the slice data
//
// COMMON WRITERS
//
//   *os.File         - Files
//   *bytes.Buffer    - In-memory buffer
//   os.Stdout        - Standard output
//   os.Stderr        - Standard error
//   http.ResponseWriter - HTTP response
//   net.Conn         - Network connection
//
// =============================================================================

func DemonstrateWriter() {
	// bytes.Buffer as writer
	var buf bytes.Buffer
	buf.Write([]byte("Hello, "))
	buf.WriteString("World!")
	fmt.Printf("Buffer: %s\n", buf.String())

	// os.Stdout is a Writer
	io.WriteString(os.Stdout, "Direct to stdout\n")

	// Multiple writes
	var buf2 bytes.Buffer
	fmt.Fprintf(&buf2, "Name: %s, ", "Alice")
	fmt.Fprintf(&buf2, "Age: %d", 30)
	fmt.Printf("Formatted: %s\n", buf2.String())
}

// =============================================================================
// SECTION 4: io.Closer
// =============================================================================
//
// THE CLOSE CONTRACT
//
//   Close() error
//
// - Releases resources (file handles, network connections)
// - Should be called when done, typically via defer
// - Can be called multiple times (should be safe)
// - Behavior after Close is undefined
//
// THE DEFER PATTERN
//
//   f, err := os.Open("file.txt")
//   if err != nil {
//       return err
//   }
//   defer f.Close()  // Always close!
//
// =============================================================================

func DemonstrateCloser() {
	// Standard pattern with defer
	r := io.NopCloser(strings.NewReader("data"))
	defer r.Close()

	data, _ := io.ReadAll(r)
	fmt.Printf("Read before close: %s\n", data)

	// io.NopCloser wraps a Reader as ReadCloser (Close does nothing)
	reader := strings.NewReader("no close needed")
	rc := io.NopCloser(reader)
	defer rc.Close() // Safe to call, does nothing
	fmt.Println("NopCloser: wraps Reader as ReadCloser")
}

// =============================================================================
// SECTION 5: COMPOSITE INTERFACES
// =============================================================================
//
// Go composes interfaces by embedding:
//
//   type ReadWriter interface {
//       Reader
//       Writer
//   }
//
//   type ReadCloser interface {
//       Reader
//       Closer
//   }
//
//   type WriteCloser interface {
//       Writer
//       Closer
//   }
//
//   type ReadWriteCloser interface {
//       Reader
//       Writer
//       Closer
//   }
//
// WHY COMPOSITION MATTERS
//
// - os.File implements ReadWriteCloser
// - http.Response.Body is ReadCloser
// - You can pass any of these where Reader/Writer is expected
//
// =============================================================================

func DemonstrateComposite() {
	// bytes.Buffer is a ReadWriter
	var buf bytes.Buffer
	var rw io.ReadWriter = &buf

	rw.Write([]byte("Hello"))
	data := make([]byte, 5)
	rw.Read(data)
	fmt.Printf("ReadWriter: %s\n", data)

	// Function accepting Reader works with ReadCloser
	readAll := func(r io.Reader) string {
		data, _ := io.ReadAll(r)
		return string(data)
	}

	// Pass a ReadCloser where Reader expected
	rc := io.NopCloser(strings.NewReader("content"))
	result := readAll(rc)
	fmt.Printf("ReadCloser as Reader: %s\n", result)
}

// =============================================================================
// SECTION 6: UTILITY FUNCTIONS
// =============================================================================
//
// io.Copy(dst Writer, src Reader) (int64, error)
//   Copies from src to dst until EOF
//
// io.CopyN(dst Writer, src Reader, n int64) (int64, error)
//   Copies exactly n bytes
//
// io.CopyBuffer(dst Writer, src Reader, buf []byte) (int64, error)
//   Copy with provided buffer (avoids allocation)
//
// io.ReadAll(r Reader) ([]byte, error)
//   Reads until EOF, returns all data
//
// io.ReadFull(r Reader, buf []byte) (int, error)
//   Reads exactly len(buf) bytes, errors if can't
//
// io.WriteString(w Writer, s string) (int, error)
//   Writes string to writer (optimized if possible)
//
// io.LimitReader(r Reader, n int64) Reader
//   Returns Reader that reads at most n bytes
//
// io.TeeReader(r Reader, w Writer) Reader
//   Reader that writes to w what it reads from r
//
// io.MultiReader(readers ...Reader) Reader
//   Concatenates multiple readers
//
// io.MultiWriter(writers ...Writer) Writer
//   Writes to multiple writers simultaneously
//
// =============================================================================

func DemonstrateUtilities() {
	// io.Copy - copy everything
	src := strings.NewReader("Copy this content")
	var dst bytes.Buffer
	n, _ := io.Copy(&dst, src)
	fmt.Printf("Copied %d bytes: %s\n", n, dst.String())

	// io.CopyN - copy exactly N bytes
	src2 := strings.NewReader("Only first five")
	var dst2 bytes.Buffer
	io.CopyN(&dst2, src2, 5)
	fmt.Printf("CopyN: %s\n", dst2.String())

	// io.ReadAll - slurp everything
	src3 := strings.NewReader("Read it all")
	data, _ := io.ReadAll(src3)
	fmt.Printf("ReadAll: %s\n", data)

	// io.ReadFull - must read exactly N bytes
	src4 := strings.NewReader("Hello, World!")
	buf := make([]byte, 5)
	io.ReadFull(src4, buf)
	fmt.Printf("ReadFull: %s\n", buf)

	// io.LimitReader - cap the bytes
	src5 := strings.NewReader("This is a very long string")
	limited := io.LimitReader(src5, 10)
	data2, _ := io.ReadAll(limited)
	fmt.Printf("Limited: %s\n", data2)

	// io.TeeReader - read and copy simultaneously
	src6 := strings.NewReader("Tee this")
	var teeDst bytes.Buffer
	tee := io.TeeReader(src6, &teeDst)
	io.ReadAll(tee) // Read from tee
	fmt.Printf("TeeReader copied: %s\n", teeDst.String())

	// io.MultiReader - concatenate readers
	r1 := strings.NewReader("Hello, ")
	r2 := strings.NewReader("World!")
	multi := io.MultiReader(r1, r2)
	all, _ := io.ReadAll(multi)
	fmt.Printf("MultiReader: %s\n", all)

	// io.MultiWriter - write to multiple destinations
	var buf1, buf2 bytes.Buffer
	mw := io.MultiWriter(&buf1, &buf2)
	mw.Write([]byte("To both"))
	fmt.Printf("MultiWriter: buf1=%s, buf2=%s\n", buf1.String(), buf2.String())
}

// =============================================================================
// SECTION 7: io.Pipe
// =============================================================================
//
// WHAT IS A PIPE?
//
// Creates a synchronous in-memory pipe.
// Writes to PipeWriter are available to read from PipeReader.
// Useful for connecting code that writes to code that reads.
//
//   r, w := io.Pipe()
//
//   go func() {
//       w.Write(data)
//       w.Close()
//   }()
//
//   io.Copy(dst, r)
//
// BEHAVIOR
//
// - Writes block until read (no internal buffer)
// - Closing writer sends EOF to reader
// - Closing reader makes writes return ErrClosedPipe
//
// =============================================================================

func DemonstratePipe() {
	r, w := io.Pipe()

	// Writer goroutine
	go func() {
		defer w.Close()
		w.Write([]byte("First chunk. "))
		w.Write([]byte("Second chunk."))
	}()

	// Reader
	data, _ := io.ReadAll(r)
	fmt.Printf("Pipe received: %s\n", data)

	// Pipe with error
	r2, w2 := io.Pipe()

	go func() {
		w2.CloseWithError(errors.New("custom error"))
	}()

	_, err := io.ReadAll(r2)
	fmt.Printf("Pipe error: %v\n", err)
}

// =============================================================================
// SECTION 8: SEEKING
// =============================================================================
//
// io.Seeker INTERFACE
//
//   type Seeker interface {
//       Seek(offset int64, whence int) (int64, error)
//   }
//
// WHENCE CONSTANTS
//
//   io.SeekStart   = 0 - Seek from beginning
//   io.SeekCurrent = 1 - Seek from current position
//   io.SeekEnd     = 2 - Seek from end
//
// COMPOSITE INTERFACES
//
//   io.ReadSeeker  = Reader + Seeker
//   io.WriteSeeker = Writer + Seeker
//   io.ReadWriteSeeker = Reader + Writer + Seeker
//
// NOT ALL READERS SUPPORT SEEKING
//
// - bytes.Reader: Yes
// - strings.Reader: Yes
// - os.File: Yes
// - Network streams: No
// - HTTP response body: No
//
// =============================================================================

func DemonstrateSeeking() {
	// bytes.Reader supports seeking
	data := []byte("0123456789ABCDEF")
	r := bytes.NewReader(data)

	// Seek from start
	r.Seek(5, io.SeekStart)
	b := make([]byte, 3)
	r.Read(b)
	fmt.Printf("SeekStart+5: %s\n", b)

	// Seek from current position
	r.Seek(2, io.SeekCurrent)
	r.Read(b)
	fmt.Printf("SeekCurrent+2: %s\n", b)

	// Seek from end
	r.Seek(-4, io.SeekEnd)
	r.Read(b)
	fmt.Printf("SeekEnd-4: %s\n", b)

	// Get current position
	pos, _ := r.Seek(0, io.SeekCurrent)
	fmt.Printf("Current position: %d\n", pos)
}

// =============================================================================
// SECTION 9: COMMON PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Read loop with EOF check
	readLoop := func(r io.Reader) ([]byte, error) {
		var result []byte
		buf := make([]byte, 32)
		for {
			n, err := r.Read(buf)
			if n > 0 {
				result = append(result, buf[:n]...)
			}
			if err == io.EOF {
				return result, nil // EOF is not an error
			}
			if err != nil {
				return result, err
			}
		}
	}
	data, _ := readLoop(strings.NewReader("Pattern 1"))
	fmt.Printf("Read loop: %s\n", data)

	// Pattern 2: Copy with progress
	copyWithProgress := func(dst io.Writer, src io.Reader, size int64) error {
		buf := make([]byte, 1024)
		var copied int64
		for {
			n, err := src.Read(buf)
			if n > 0 {
				dst.Write(buf[:n])
				copied += int64(n)
				pct := float64(copied) / float64(size) * 100
				fmt.Printf("  Progress: %.0f%%\n", pct)
			}
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
		}
	}
	var dst bytes.Buffer
	copyWithProgress(&dst, strings.NewReader("12345678901234567890"), 20)

	// Pattern 3: Limit request body
	_ = `
	body := io.LimitReader(r.Body, 1024*1024) // Max 1MB
	data, err := io.ReadAll(body)
	`
	fmt.Println("Pattern 3: LimitReader for request bodies")

	// Pattern 4: Tee for logging
	_ = `
	var logBuf bytes.Buffer
	tee := io.TeeReader(r.Body, &logBuf)
	json.NewDecoder(tee).Decode(&data)
	log.Printf("Request body: %s", logBuf.String())
	`
	fmt.Println("Pattern 4: TeeReader for logging")

	// Pattern 5: Pipe for streaming
	_ = `
	r, w := io.Pipe()
	go func() {
		json.NewEncoder(w).Encode(data)
		w.Close()
	}()
	http.Post(url, "application/json", r)
	`
	fmt.Println("Pattern 5: Pipe for streaming HTTP")

	// Pattern 6: Interface-based testing
	processData := func(r io.Reader) (int, error) {
		data, err := io.ReadAll(r)
		return len(data), err
	}
	// In tests: processData(strings.NewReader("test data"))
	// In prod:  processData(file)
	n, _ := processData(strings.NewReader("test data"))
	fmt.Printf("Pattern 6: Interface testing, read %d bytes\n", n)
}

// =============================================================================
// SECTION 10: IMPLEMENTING CUSTOM READERS/WRITERS
// =============================================================================

// CountingReader wraps a Reader and counts bytes read
type CountingReader struct {
	r     io.Reader
	count int64
}

func (cr *CountingReader) Read(p []byte) (int, error) {
	n, err := cr.r.Read(p)
	cr.count += int64(n)
	return n, err
}

func (cr *CountingReader) Count() int64 {
	return cr.count
}

// UpperWriter converts all bytes to uppercase
type UpperWriter struct {
	w io.Writer
}

func (uw *UpperWriter) Write(p []byte) (int, error) {
	upper := bytes.ToUpper(p)
	return uw.w.Write(upper)
}

// ZeroReader returns infinite zeros
type ZeroReader struct{}

func (ZeroReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 0
	}
	return len(p), nil
}

func DemonstrateCustom() {
	// CountingReader
	src := strings.NewReader("Count these bytes")
	cr := &CountingReader{r: src}
	io.ReadAll(cr)
	fmt.Printf("CountingReader: %d bytes\n", cr.Count())

	// UpperWriter
	var buf bytes.Buffer
	uw := &UpperWriter{w: &buf}
	uw.Write([]byte("hello world"))
	fmt.Printf("UpperWriter: %s\n", buf.String())

	// ZeroReader with LimitReader
	zr := ZeroReader{}
	limited := io.LimitReader(zr, 10)
	zeros, _ := io.ReadAll(limited)
	fmt.Printf("ZeroReader: %v\n", zeros)
}

// =============================================================================
// SECTION 11: SENTINEL ERRORS
// =============================================================================
//
// io.EOF
//   End of input, not an error. Normal termination.
//   Check: if err == io.EOF or errors.Is(err, io.EOF)
//
// io.ErrUnexpectedEOF
//   EOF when more data was expected (e.g., ReadFull didn't complete)
//
// io.ErrClosedPipe
//   Write to closed pipe
//
// io.ErrNoProgress
//   Reader repeatedly returned 0 bytes without error
//
// io.ErrShortBuffer
//   Buffer too small for operation
//
// io.ErrShortWrite
//   Write didn't write all bytes requested
//
// =============================================================================

func DemonstrateSentinelErrors() {
	// EOF is the normal end
	r := strings.NewReader("data")
	for {
		buf := make([]byte, 10)
		n, err := r.Read(buf)
		if err == io.EOF {
			fmt.Println("EOF: end of data (normal)")
			break
		}
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Printf("Read: %s\n", buf[:n])
	}

	// ErrUnexpectedEOF from ReadFull
	short := strings.NewReader("Hi")
	buf := make([]byte, 10)
	_, err := io.ReadFull(short, buf)
	if errors.Is(err, io.ErrUnexpectedEOF) {
		fmt.Println("ErrUnexpectedEOF: wanted more data")
	}

	// ErrClosedPipe
	r2, w := io.Pipe()
	r2.Close()
	_, err2 := w.Write([]byte("data"))
	if errors.Is(err2, io.ErrClosedPipe) {
		fmt.Println("ErrClosedPipe: pipe was closed")
	}
}

// =============================================================================
// SECTION 12: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Treating EOF as an error
	_ = `
	// WRONG
	n, err := r.Read(buf)
	if err != nil {  // EOF goes here!
		return err
	}

	// RIGHT
	n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		return err
	}
	`
	fmt.Println("Mistake 1: EOF is not an error, check separately")

	// Mistake 2: Ignoring partial reads
	_ = `
	// WRONG
	buf := make([]byte, 100)
	r.Read(buf)  // May read less than 100!
	process(buf)  // Processing uninitialized bytes

	// RIGHT
	n, _ := r.Read(buf)
	process(buf[:n])  // Only process what was read
	`
	fmt.Println("Mistake 2: Always use buf[:n], not buf")

	// Mistake 3: Not closing resources
	_ = `
	// WRONG
	f, _ := os.Open("file.txt")
	data, _ := io.ReadAll(f)
	// File never closed!

	// RIGHT
	f, err := os.Open("file.txt")
	if err != nil { return err }
	defer f.Close()
	data, _ := io.ReadAll(f)
	`
	fmt.Println("Mistake 3: Always defer Close() after opening")

	// Mistake 4: Ignoring Write return value
	_ = `
	// WRONG
	w.Write(data)  // Ignores errors and short writes

	// RIGHT
	n, err := w.Write(data)
	if err != nil { return err }
	if n < len(data) { return io.ErrShortWrite }
	`
	fmt.Println("Mistake 4: Check Write returns for errors")

	// Mistake 5: Reusing buffer incorrectly
	_ = `
	// WRONG
	buf := make([]byte, 4)
	r.Read(buf)  // Reads "Hell"
	r.Read(buf)  // Reads "o Wo", but first 4 bytes still "Hell"
	// buf is now "o Wo", previous content overwritten

	// RIGHT
	for {
		n, err := r.Read(buf)
		process(buf[:n])  // Fresh slice each time
	}
	`
	fmt.Println("Mistake 5: Buffer contains stale data, use [:n]")
}

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateReader()
	DemonstrateWriter()
	DemonstrateCloser()
	DemonstrateComposite()
	DemonstrateUtilities()
	DemonstratePipe()
	DemonstrateSeeking()
	DemonstratePatterns()
	DemonstrateCustom()
	DemonstrateSentinelErrors()
	DemonstrateCommonMistakes()
}
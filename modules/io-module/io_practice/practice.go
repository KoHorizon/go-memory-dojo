// Package io_practice is your daily practice space for the io package.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against io/io.go
// 3. Note what you missed - focus on those tomorrow
package io_practice

// =============================================================================
// EXERCISE 1: CORE INTERFACES
// =============================================================================

func PracticeCoreInterfaces() {
	// TODO: Write the Reader interface
	// type Reader interface {
	//     ???(p []byte) (??? int, ??? error)
	// }

	// TODO: Write the Writer interface
	// type Writer interface {
	//     ???(p []byte) (??? int, ??? error)
	// }

	// TODO: Write the Closer interface
	// type Closer interface {
	//     ???() error
	// }

	// TODO: What does the Read contract say about n?
	// Answer: ???

	// TODO: What error signals end of input?
	// Answer: io.???

	// TODO: Is io.EOF an error?
	// Answer: ???
}

// =============================================================================
// EXERCISE 2: COMPOSITE INTERFACES
// =============================================================================

func PracticeCompositeInterfaces() {
	// TODO: Name the 4 composite interfaces
	// 1. io.??? = Reader + Writer
	// 2. io.??? = Reader + Closer
	// 3. io.??? = Writer + Closer
	// 4. io.??? = Reader + Writer + Closer

	// TODO: What interface does os.File implement?
	// Answer: io.???

	// TODO: What interface is http.Response.Body?
	// Answer: io.???
}

// =============================================================================
// EXERCISE 3: COMMON READERS AND WRITERS
// =============================================================================

func PracticeCommonTypes() {
	// TODO: Create a Reader from a string
	// r := strings.???(str)

	// TODO: Create a read-only Reader from bytes
	// r := bytes.???(data)

	// TODO: Create a read-write Buffer
	// var buf bytes.???

	// TODO: Name 3 things that implement Writer
	// 1. ???
	// 2. ???
	// 3. ???
}

// =============================================================================
// EXERCISE 4: UTILITY FUNCTIONS
// =============================================================================

func PracticeUtilities() {
	// TODO: Copy all data from Reader to Writer
	// n, err := io.???(dst, src)

	// TODO: Copy exactly N bytes
	// n, err := io.???(dst, src, n)

	// TODO: Read everything into a slice
	// data, err := io.???(r)

	// TODO: Read exactly len(buf) bytes
	// n, err := io.???(r, buf)

	// TODO: Write a string to a Writer
	// n, err := io.???(w, str)

	// TODO: Limit a Reader to N bytes
	// limited := io.???(r, n)

	// TODO: Read from r, copy to w simultaneously
	// tee := io.???(r, w)

	// TODO: Concatenate multiple Readers
	// multi := io.???(r1, r2, r3)

	// TODO: Write to multiple Writers at once
	// multi := io.???(w1, w2, w3)
}

// =============================================================================
// EXERCISE 5: io.Pipe
// =============================================================================

func PracticePipe() {
	// TODO: Create a pipe
	// r, w := io.???()

	// TODO: What blocks until what?
	// Answer: ???

	// TODO: How do you signal EOF to the reader?
	// Answer: w.???()

	// TODO: How do you signal an error?
	// Answer: w.???(err)

	// TODO: What error does the writer get if reader closes?
	// Answer: io.???
}

// =============================================================================
// EXERCISE 6: SEEKING
// =============================================================================

func PracticeSeeking() {
	// TODO: Write the Seeker interface
	// type Seeker interface {
	//     ???(offset int64, whence int) (int64, error)
	// }

	// TODO: What are the three whence constants?
	// io.??? = 0 - from beginning
	// io.??? = 1 - from current
	// io.??? = 2 - from end

	// TODO: Seek to position 10 from start
	// r.Seek(10, io.???)

	// TODO: Seek back 5 bytes from current
	// r.Seek(-5, io.???)

	// TODO: Seek to 10 bytes before end
	// r.Seek(-10, io.???)

	// TODO: Get current position without moving
	// pos, _ := r.Seek(???, io.???)
}

// =============================================================================
// EXERCISE 7: io.NopCloser
// =============================================================================

func PracticeNopCloser() {
	// TODO: What does io.NopCloser do?
	// Answer: ???

	// TODO: Why would you use it?
	// Answer: ???

	// TODO: Use NopCloser
	// rc := io.???(reader)
}

// =============================================================================
// EXERCISE 8: SENTINEL ERRORS
// =============================================================================

func PracticeSentinelErrors() {
	// TODO: Name 6 sentinel errors in io package
	// 1. io.??? - end of input (normal)
	// 2. io.??? - EOF when more expected
	// 3. io.??? - write to closed pipe
	// 4. io.??? - repeated zero-byte reads
	// 5. io.??? - buffer too small
	// 6. io.??? - write didn't complete

	// TODO: How do you check for EOF?
	// if err == io.??? || errors.Is(err, io.???)
}

// =============================================================================
// EXERCISE 9: READ LOOP PATTERN
// =============================================================================

func PracticeReadLoop() {
	// TODO: Write the standard read loop
	// buf := make([]byte, 1024)
	// for {
	//     n, err := r.???(buf)
	//     if n > 0 {
	//         // Process buf[:???]
	//     }
	//     if err == io.??? {
	//         break  // Done
	//     }
	//     if err != nil {
	//         return err
	//     }
	// }

	// TODO: Why check n > 0 before checking err?
	// Answer: ???
}

// =============================================================================
// EXERCISE 10: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// n, err := r.Read(buf)
	// if err != nil {
	//     return err
	// }
	// Answer: ???

	// Mistake 2: What's wrong?
	// buf := make([]byte, 100)
	// r.Read(buf)
	// process(buf)
	// Answer: ???

	// Mistake 3: What's wrong?
	// f, _ := os.Open("file.txt")
	// data, _ := io.ReadAll(f)
	// // use data
	// Answer: ???

	// Mistake 4: What's wrong?
	// w.Write(data)  // Ignoring return
	// Answer: ???
}

// =============================================================================
// EXERCISE 11: IMPLEMENTING INTERFACES
// =============================================================================

func PracticeImplementing() {
	// TODO: Implement a CountingReader
	// type CountingReader struct {
	//     r     io.???
	//     count ???
	// }
	//
	// func (cr *CountingReader) Read(p []byte) (int, error) {
	//     n, err := cr.r.???(p)
	//     cr.count += int64(???)
	//     return n, err
	// }

	// TODO: What must a type have to be a Reader?
	// Answer: ???

	// TODO: What must a type have to be a Writer?
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// 1. Reader interface method?
	_ = "Read(p []byte) (???, ???)"

	// 2. Writer interface method?
	_ = "Write(p []byte) (???, ???)"

	// 3. Copy from Reader to Writer?
	_ = "io.???(dst, src)"

	// 4. Read all into slice?
	_ = "io.???(r)"

	// 5. Read exactly N bytes?
	_ = "io.???(r, buf)"

	// 6. Limit reader to N bytes?
	_ = "io.???(r, n)"

	// 7. Tee reader?
	_ = "io.???(r, w)"

	// 8. Concatenate readers?
	_ = "io.???(r1, r2)"

	// 9. Write to multiple?
	_ = "io.???(w1, w2)"

	// 10. Create pipe?
	_ = "io.???()"

	// 11. Seek from start constant?
	_ = "io.???"

	// 12. End of input error?
	_ = "io.???"
}

// =============================================================================
// MINI PROJECT: PROGRESS READER
// =============================================================================

func MiniProject() {
	// Build a ProgressReader that:
	//
	// 1. Wraps another Reader
	// 2. Tracks bytes read
	// 3. Calls a callback with progress percentage
	// 4. Requires total size to calculate percentage
	//
	// Scaffold:
	//
	// type ProgressReader struct {
	//     r        io.Reader
	//     total    int64
	//     read     int64
	//     callback func(pct float64)
	// }
	//
	// func NewProgressReader(r io.Reader, total int64, cb func(float64)) *ProgressReader {
	//     return &ProgressReader{r: r, total: total, callback: cb}
	// }
	//
	// func (pr *ProgressReader) Read(p []byte) (int, error) {
	//     n, err := pr.r.Read(p)
	//     if n > 0 {
	//         pr.read += int64(n)
	//         pct := float64(pr.read) / float64(pr.total) * 100
	//         pr.callback(pct)
	//     }
	//     return n, err
	// }
	//
	// Usage:
	// pr := NewProgressReader(file, fileSize, func(pct float64) {
	//     fmt.Printf("\rProgress: %.1f%%", pct)
	// })
	// io.Copy(dest, pr)
}

// RunAllPractice is not meant to be called - this file is for reading.
func RunAllPractice() {
	PracticeCoreInterfaces()
	PracticeCompositeInterfaces()
	PracticeCommonTypes()
	PracticeUtilities()
	PracticePipe()
	PracticeSeeking()
	PracticeNopCloser()
	PracticeSentinelErrors()
	PracticeReadLoop()
	PracticeMistakes()
	PracticeImplementing()
	SelfTest()
	MiniProject()
}
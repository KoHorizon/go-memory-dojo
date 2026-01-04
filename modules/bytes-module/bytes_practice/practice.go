// Package bytes_practice is your daily practice space for the bytes package.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against bytes/bytes.go
// 3. Note what you missed - focus on those tomorrow
package bytes_practice

// =============================================================================
// EXERCISE 1: PHILOSOPHY
// =============================================================================

func PracticePhilosophy() {
	// TODO: What's the fundamental difference between string and []byte?
	// string: ???
	// []byte: ???

	// TODO: Can you modify a string's bytes?
	// Answer: ???

	// TODO: Can you modify a []byte's bytes?
	// Answer: ???

	// TODO: When to use string vs []byte?
	// Use string when: ???
	// Use []byte when: ???

	// TODO: Do string ↔ []byte conversions copy data?
	// Answer: ???
}

// =============================================================================
// EXERCISE 2: COMPARISON
// =============================================================================

func PracticeComparison() {
	// TODO: Why can't you use == for []byte?
	// Answer: ???

	// TODO: Compare two []byte for equality
	// equal := bytes.???(b1, b2)

	// TODO: Case-insensitive comparison
	// equal := bytes.???(b1, b2)

	// TODO: Lexicographic comparison
	// cmp := bytes.???(b1, b2)

	// TODO: What does Compare return?
	// < 0 if ???
	// 0 if ???
	// > 0 if ???
}

// =============================================================================
// EXERCISE 3: SEARCHING
// =============================================================================

func PracticeSearching() {
	// TODO: Check if contains subslice
	// found := bytes.???(b, subslice)

	// TODO: Check if contains any character from string
	// found := bytes.???(b, "aeiou")

	// TODO: Check if contains rune
	// found := bytes.???(b, 'x')

	// TODO: Check if contains rune matching predicate
	// found := bytes.???(b, unicode.IsDigit)

	// TODO: Find first occurrence
	// pos := bytes.???(b, subslice)

	// TODO: Find last occurrence
	// pos := bytes.???(b, subslice)

	// TODO: Find first occurrence of byte
	// pos := bytes.???(b, 'x')

	// TODO: Count occurrences
	// count := bytes.???(b, subslice)
}

// =============================================================================
// EXERCISE 4: PREFIX, SUFFIX, CUT
// =============================================================================

func PracticePrefixSuffix() {
	// TODO: Check if starts with prefix
	// hasPrefix := bytes.???(b, prefix)

	// TODO: Check if ends with suffix
	// hasSuffix := bytes.???(b, suffix)

	// TODO: Remove prefix
	// result := bytes.???(b, prefix)

	// TODO: Remove suffix
	// result := bytes.???(b, suffix)

	// TODO: Split around separator
	// before, after, found := bytes.???(b, sep)
}

// =============================================================================
// EXERCISE 5: SPLITTING AND JOINING
// =============================================================================

func PracticeSplitJoin() {
	// TODO: Split on separator
	// parts := bytes.???(b, sep)

	// TODO: What type does Split return?
	// Answer: ???

	// TODO: Split with maximum N parts
	// parts := bytes.???(b, sep, n)

	// TODO: Split on whitespace
	// words := bytes.???(b)

	// TODO: Join slices with separator
	// result := bytes.???(parts, sep)

	// TODO: CRITICAL: Does Split copy data or share memory?
	// Answer: ???

	// TODO: What happens if you modify original after Split?
	// Answer: ???
}

// =============================================================================
// EXERCISE 6: TRIMMING
// =============================================================================

func PracticeTrimming() {
	// TODO: Remove whitespace from both ends
	// trimmed := bytes.???(b)

	// TODO: Remove characters from cutset
	// trimmed := bytes.???(b, cutset)

	// TODO: Trim using predicate function
	// trimmed := bytes.???(b, func(r rune) bool { ... })

	// TODO: Difference between Trim and TrimPrefix?
	// Trim: removes ??? from ends
	// TrimPrefix: removes ??? from start
}

// =============================================================================
// EXERCISE 7: CASE CONVERSION
// =============================================================================

func PracticeCase() {
	// TODO: Convert to uppercase
	// upper := bytes.???(b)

	// TODO: Convert to lowercase
	// lower := bytes.???(b)

	// TODO: Do these functions modify in place?
	// Answer: ???

	// TODO: Replace invalid UTF-8 sequences
	// valid := bytes.???(b, replacement)
}

// =============================================================================
// EXERCISE 8: REPLACE
// =============================================================================

func PracticeReplace() {
	// TODO: Replace first N occurrences
	// result := bytes.???(b, old, new, n)

	// TODO: Replace all occurrences
	// result := bytes.???(b, old, new)

	// TODO: Transform each rune
	// result := bytes.???(mapping, b)

	// TODO: How to drop a rune in Map?
	// Return ??? from mapping function

	// TODO: Repeat N times
	// result := bytes.???(b, count)
}

// =============================================================================
// EXERCISE 9: RUNES
// =============================================================================

func PracticeRunes() {
	// TODO: Convert []byte to []rune
	// runes := bytes.???(b)

	// TODO: Why convert to runes?
	// Answer: To handle ??? correctly

	// TODO: Does len([]byte) equal len([]rune)?
	// Answer: ???
}

// =============================================================================
// EXERCISE 10: bytes.Buffer - CREATION
// =============================================================================

func PracticeBufferCreation() {
	// TODO: Create an empty buffer (3 ways)
	// Way 1: var buf bytes.???
	// Way 2: buf := ???(bytes.Buffer)
	// Way 3: buf := bytes.???(initial)

	// TODO: Create buffer from string
	// buf := bytes.???(s)

	// TODO: Is zero value of Buffer ready to use?
	// Answer: ???
}

// =============================================================================
// EXERCISE 11: bytes.Buffer - WRITING
// =============================================================================

func PracticeBufferWriting() {
	// TODO: Write bytes to buffer
	// n, err := buf.???(p)

	// TODO: Write single byte
	// err := buf.???(c)

	// TODO: Write rune
	// n, err := buf.???(r)

	// TODO: Write string
	// n, err := buf.???(s)

	// TODO: Does Buffer implement io.Writer?
	// Answer: ???
}

// =============================================================================
// EXERCISE 12: bytes.Buffer - READING
// =============================================================================

func PracticeBufferReading() {
	// TODO: Read bytes from buffer
	// n, err := buf.???(p)

	// TODO: Read single byte
	// b, err := buf.???()

	// TODO: Read until delimiter
	// data, err := buf.???(delim)

	// TODO: Read next N bytes
	// data := buf.???(n)

	// TODO: Does Buffer implement io.Reader?
	// Answer: ???
}

// =============================================================================
// EXERCISE 13: bytes.Buffer - ACCESSING
// =============================================================================

func PracticeBufferAccessing() {
	// TODO: Get buffer contents as []byte
	// b := buf.???()

	// TODO: Get buffer contents as string
	// s := buf.???()

	// TODO: Can you safely modify the []byte from Bytes()?
	// Answer: ???

	// TODO: Get current length
	// length := buf.???()

	// TODO: Get current capacity
	// capacity := buf.???()

	// TODO: Pre-allocate space
	// buf.???(n)

	// TODO: Clear and reuse buffer
	// buf.???()

	// TODO: Keep only first N bytes
	// buf.???(n)
}

// =============================================================================
// EXERCISE 14: bytes.Reader
// =============================================================================

func PracticeReader() {
	// TODO: Create Reader from []byte
	// r := bytes.???(b)

	// TODO: Read bytes
	// n, err := r.???(p)

	// TODO: Read single byte
	// b, err := r.???()

	// TODO: Read single rune
	// r, size, err := r.???()

	// TODO: Seek to position
	// pos, err := r.???(offset, whence)

	// TODO: Get original size
	// size := r.???()

	// TODO: Get unread portion length
	// remaining := r.???()

	// TODO: Reuse Reader with new data
	// r.???(newData)

	// TODO: Why use bytes.Reader?
	// Answer: To pass []byte to APIs expecting ???
}

// =============================================================================
// EXERCISE 15: COMMON PATTERNS
// =============================================================================

func PracticePatterns() {
	// TODO: Efficient byte accumulation pattern
	// var buf bytes.???
	// buf.???(expectedSize)
	// for ... {
	//     buf.WriteString(...)
	// }
	// result := buf.???()

	// TODO: How to safely copy []byte for long-lived use?
	// copy := bytes.???(original)

	// TODO: Build protocol message pattern
	// var buf bytes.Buffer
	// buf.???(headerByte)
	// buf.???(payload)
	// return buf.???()
}

// =============================================================================
// EXERCISE 16: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// b1 := []byte("hello")
	// b2 := []byte("hello")
	// if b1 == b2 { ... }
	// Answer: ???
	// Fix: ???

	// Mistake 2: What's wrong?
	// var buf bytes.Buffer
	// buf.WriteString("hello")
	// b := buf.Bytes()
	// b[0] = 'X'
	// Answer: ???
	// Fix: ???

	// Mistake 3: What's wrong?
	// original := []byte("a,b,c")
	// parts := bytes.Split(original, []byte(","))
	// parts[0][0] = 'X'
	// Answer: ???
	// Fix: ???

	// Mistake 4: What's wrong?
	// var result []byte
	// for i := 0; i < 1000; i++ {
	//     result = append(result, []byte("x")...)
	// }
	// Answer: ???
	// Fix: ???

	// Mistake 5: What's wrong?
	// b := []byte(strings.ToUpper(string(data)))
	// Answer: ???
	// Fix: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// 1. Compare []byte for equality?
	_ = "bytes.???(b1, b2)"

	// 2. Check if contains?
	_ = "bytes.???(b, subslice)"

	// 3. Find first occurrence?
	_ = "bytes.???(b, subslice)"

	// 4. Check prefix?
	_ = "bytes.???(b, prefix)"

	// 5. Split on separator?
	_ = "bytes.???(b, sep)"

	// 6. Split on whitespace?
	_ = "bytes.???(b)"

	// 7. Join slices?
	_ = "bytes.???(parts, sep)"

	// 8. Remove whitespace?
	_ = "bytes.???(b)"

	// 9. To uppercase?
	_ = "bytes.???(b)"

	// 10. Replace all?
	_ = "bytes.???(b, old, new)"

	// 11. Convert to runes?
	_ = "bytes.???(b)"

	// 12. Create empty buffer?
	_ = "var buf bytes.???"

	// 13. Write string to buffer?
	_ = "buf.???(s)"

	// 14. Get buffer as string?
	_ = "buf.???()"

	// 15. Clear buffer?
	_ = "buf.???()"

	// 16. Pre-allocate buffer space?
	_ = "buf.???(n)"

	// 17. Create Reader?
	_ = "bytes.???(b)"

	// 18. Clone []byte?
	_ = "bytes.???(b)"
}

// =============================================================================
// MINI PROJECT: PROTOCOL PARSER
// =============================================================================

func MiniProject() {
	// Build a simple binary protocol parser/builder that:
	//
	// 1. Message format:
	//    [1 byte: type] [2 bytes: length] [N bytes: payload] [1 byte: checksum]
	//
	// 2. func BuildMessage(msgType byte, payload []byte) []byte
	//    - Build a protocol message
	//    - Calculate checksum (XOR of all payload bytes)
	//
	// 3. func ParseMessage(data []byte) (msgType byte, payload []byte, err error)
	//    - Parse and validate message
	//    - Check length field matches payload
	//    - Verify checksum
	//
	// 4. func ReadMessages(r io.Reader) ([]Message, error)
	//    - Read multiple messages from reader
	//    - Stop at EOF
	//
	// 5. func WriteMessages(w io.Writer, msgs []Message) error
	//    - Write multiple messages
	//
	// 6. Message types:
	//    const (
	//        MsgTypeData = 0x01
	//        MsgTypeAck  = 0x02
	//        MsgTypePing = 0x03
	//    )
	//
	// Scaffold:
	//
	// type Message struct {
	//     Type    byte
	//     Payload []byte
	// }
	//
	// func BuildMessage(msgType byte, payload []byte) []byte {
	//     var buf bytes.Buffer
	//     
	//     // Type
	//     buf.WriteByte(msgType)
	//     
	//     // Length (2 bytes, big-endian)
	//     length := len(payload)
	//     buf.WriteByte(byte(length >> 8))
	//     buf.WriteByte(byte(length))
	//     
	//     // Payload
	//     buf.Write(payload)
	//     
	//     // Checksum (XOR of all payload bytes)
	//     var checksum byte
	//     for _, b := range payload {
	//         checksum ^= b
	//     }
	//     buf.WriteByte(checksum)
	//     
	//     return buf.Bytes()
	// }
	//
	// func ParseMessage(data []byte) (*Message, error) {
	//     if len(data) < 4 {
	//         return nil, errors.New("message too short")
	//     }
	//     
	//     // Parse header
	//     msgType := data[0]
	//     length := int(data[1])<<8 | int(data[2])
	//     
	//     // Check length
	//     if len(data) < 3+length+1 {
	//         return nil, errors.New("incomplete message")
	//     }
	//     
	//     // Extract payload
	//     payload := data[3 : 3+length]
	//     
	//     // Verify checksum
	//     var checksum byte
	//     for _, b := range payload {
	//         checksum ^= b
	//     }
	//     if checksum != data[3+length] {
	//         return nil, errors.New("checksum mismatch")
	//     }
	//     
	//     return &Message{
	//         Type:    msgType,
	//         Payload: bytes.Clone(payload),
	//     }, nil
	// }
	//
	// func ReadMessages(r io.Reader) ([]*Message, error) {
	//     var messages []*Message
	//     buf := make([]byte, 4096)
	//     
	//     for {
	//         n, err := r.Read(buf)
	//         if err == io.EOF {
	//             break
	//         }
	//         if err != nil {
	//             return nil, err
	//         }
	//         
	//         data := buf[:n]
	//         for len(data) >= 4 {
	//             msg, err := ParseMessage(data)
	//             if err != nil {
	//                 return nil, err
	//             }
	//             messages = append(messages, msg)
	//             
	//             // Advance past this message
	//             msgLen := 3 + len(msg.Payload) + 1
	//             data = data[msgLen:]
	//         }
	//     }
	//     
	//     return messages, nil
	// }
	//
	// func WriteMessages(w io.Writer, msgs []*Message) error {
	//     var buf bytes.Buffer
	//     
	//     for _, msg := range msgs {
	//         data := BuildMessage(msg.Type, msg.Payload)
	//         buf.Write(data)
	//     }
	//     
	//     _, err := w.Write(buf.Bytes())
	//     return err
	// }
	//
	// // Test:
	// func main() {
	//     // Build message
	//     msg := BuildMessage(MsgTypeData, []byte("Hello"))
	//     fmt.Printf("Built: %v\n", msg)
	//     
	//     // Parse it back
	//     parsed, err := ParseMessage(msg)
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	//     fmt.Printf("Parsed: Type=%d, Payload=%s\n", parsed.Type, parsed.Payload)
	//     
	//     // Round-trip test
	//     var buf bytes.Buffer
	//     msgs := []*Message{
	//         {Type: MsgTypeData, Payload: []byte("first")},
	//         {Type: MsgTypePing, Payload: []byte{}},
	//         {Type: MsgTypeAck, Payload: []byte("ok")},
	//     }
	//     
	//     WriteMessages(&buf, msgs)
	//     read, err := ReadMessages(&buf)
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	//     
	//     fmt.Printf("Round-trip: %d messages\n", len(read))
	//     for i, m := range read {
	//         fmt.Printf("  %d: Type=%d, Payload=%s\n", i, m.Type, m.Payload)
	//     }
	// }
}

// AllPractice is not meant to be called - this file is for reading.
func AllPractice() {
	PracticePhilosophy()
	PracticeComparison()
	PracticeSearching()
	PracticePrefixSuffix()
	PracticeSplitJoin()
	PracticeTrimming()
	PracticeCase()
	PracticeReplace()
	PracticeRunes()
	PracticeBufferCreation()
	PracticeBufferWriting()
	PracticeBufferReading()
	PracticeBufferAccessing()
	PracticeReader()
	PracticePatterns()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
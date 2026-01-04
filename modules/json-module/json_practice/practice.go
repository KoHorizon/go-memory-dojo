// Package json_practice is your daily practice space for encoding/json.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against json/json.go
// 3. Note what you missed - focus on those tomorrow
package json_practice

// =============================================================================
// EXERCISE 1: TYPE MAPPING
// =============================================================================

func PracticeTypeMapping() {
	// TODO: What Go types map to what JSON types?
	// Go              JSON
	// ──────────────────────────
	// bool         → ???
	// int, float   → ???
	// string       → ???
	// []T          → ???
	// map[string]T → ???
	// struct       → ???
	// nil pointer  → ???
	// nil slice    → ???

	// TODO: Which struct fields are marshaled/unmarshaled?
	// Answer: ???
}

// =============================================================================
// EXERCISE 2: BASIC MARSHALING
// =============================================================================

func PracticeMarshal() {
	// TODO: Marshal a value to JSON bytes
	// data, err := json.???(value)

	// TODO: Marshal with pretty printing
	// data, err := json.???(value, "", "  ")

	// TODO: What types CAN'T be marshaled?
	// Answer: ???
}

// =============================================================================
// EXERCISE 3: BASIC UNMARSHALING
// =============================================================================

func PracticeUnmarshal() {
	// TODO: Unmarshal JSON bytes into a value
	// err := json.???(data, ???v)  // What must v be?

	// TODO: What happens with missing JSON fields?
	// Answer: ???

	// TODO: What happens with extra JSON fields?
	// Answer: ???

	// TODO: What happens with type mismatch?
	// Answer: ???
}

// =============================================================================
// EXERCISE 4: STRUCT TAGS
// =============================================================================

func PracticeStructTags() {
	// TODO: Complete the struct tag syntax
	// `json:"???,???"`

	// TODO: What do these tags mean?
	// `json:"name"`         → ???
	// `json:"name,omitempty"` → ???
	// `json:"-"`            → ???
	// `json:",string"`      → ???

	// TODO: What values count as "empty" for omitempty?
	// Answer: ???

	// TODO: Write a struct with custom JSON field name
	// type User struct {
	//     ID int `json:"???"`
	// }

	// TODO: Write a field that's excluded from JSON
	// Password string `json:"???"`

	// TODO: Write a field omitted when zero
	// Email string `json:"???,???"`
}

// =============================================================================
// EXERCISE 5: EMBEDDED STRUCTS
// =============================================================================

func PracticeEmbedding() {
	// TODO: What happens to embedded struct fields in JSON?
	// Answer: ???

	// TODO: How do you get nested JSON instead of flattened?
	// Answer: ???

	// Example:
	// type Address struct { City string }
	// type Person struct {
	//     Name    string
	//     Address        // Embedded → ???
	// }
	// type Person2 struct {
	//     Name    string
	//     Address Address  // Named → ???
	// }
}

// =============================================================================
// EXERCISE 6: POINTERS AND NIL
// =============================================================================

func PracticePointersAndNil() {
	// TODO: What does a nil pointer marshal to?
	// Answer: ???

	// TODO: Why use pointer + omitempty?
	// Answer: ???

	// TODO: How do you distinguish "not set" from "zero value"?
	// type Config struct {
	//     Count ???int `json:"count,omitempty"`
	// }
}

// =============================================================================
// EXERCISE 7: DYNAMIC JSON
// =============================================================================

func PracticeDynamicJSON() {
	// TODO: What type do you use for unknown JSON structure?
	// var data ???[string]???

	// TODO: What Go types do JSON values unmarshal to?
	// object → ???
	// array  → ???
	// string → ???
	// number → ??? (important!)
	// bool   → ???
	// null   → ???

	// TODO: Why is the number type surprising?
	// Answer: ???

	// TODO: How do you preserve precise numbers?
	// dec := json.NewDecoder(reader)
	// dec.???()
}

// =============================================================================
// EXERCISE 8: json.RawMessage
// =============================================================================

func PracticeRawMessage() {
	// TODO: What is json.RawMessage?
	// Answer: ???

	// TODO: Name two use cases
	// 1. ???
	// 2. ???

	// TODO: How do you delay parsing part of JSON?
	// type Envelope struct {
	//     Type    string          `json:"type"`
	//     Payload json.??? `json:"payload"`
	// }
}

// =============================================================================
// EXERCISE 9: CUSTOM MARSHALING
// =============================================================================

func PracticeCustomMarshal() {
	// TODO: What interface for custom marshaling?
	// type ??? interface {
	//     ???() ([]byte, error)
	// }

	// TODO: What interface for custom unmarshaling?
	// type ??? interface {
	//     ???([]byte) error
	// }

	// TODO: Name two use cases for custom marshaling
	// 1. ???
	// 2. ???
}

// =============================================================================
// EXERCISE 10: STREAMING
// =============================================================================

func PracticeStreaming() {
	// TODO: Create an encoder
	// enc := json.???(writer)

	// TODO: Create a decoder
	// dec := json.???(reader)

	// TODO: Encode a value
	// enc.???(value)

	// TODO: Decode a value
	// dec.???(&value)

	// TODO: How do you error on unknown fields?
	// dec.???()

	// TODO: When use Encoder/Decoder vs Marshal/Unmarshal?
	// Encoder/Decoder: ???
	// Marshal/Unmarshal: ???
}

// =============================================================================
// EXERCISE 11: COMMON PATTERNS
// =============================================================================

func PracticePatterns() {
	// TODO: HTTP request body pattern
	// var req RequestBody
	// if err := json.???(r.Body).???(&req); err != nil {
	//     ...
	// }

	// TODO: HTTP response pattern
	// w.Header().Set("Content-Type", "???")
	// json.???(w).???(response)

	// TODO: Default values pattern - why does this work?
	// config := Config{Host: "localhost", Port: 8080}
	// json.Unmarshal(data, &config)
	// Answer: ???

	// TODO: Patch/update pattern - why use pointer?
	// type Patch struct {
	//     Name *string `json:"name,omitempty"`
	// }
	// Answer: ???
}

// =============================================================================
// EXERCISE 12: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// var d Data
	// json.Unmarshal(data, d)
	// Answer: ???

	// Mistake 2: What's wrong?
	// type data struct {
	//     value int
	// }
	// Answer: ???

	// Mistake 3: What's wrong?
	// var m map[string]interface{}
	// json.Unmarshal([]byte(`{"n": 42}`), &m)
	// count := m["n"].(int)
	// Answer: ???

	// Mistake 4: What's wrong?
	// json.Unmarshal(data, &v)
	// use(v)
	// Answer: ???

	// Mistake 5: What's wrong?
	// type T struct {
	//     Name string `json: "name"`
	// }
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// 1. Marshal to bytes?
	_ = "json.???(v)"

	// 2. Unmarshal from bytes?
	_ = "json.???(data, &v)"

	// 3. Pretty print?
	_ = "json.???(v, prefix, indent)"

	// 4. Custom field name tag?
	_ = "`json:\"???\"`"

	// 5. Omit when empty tag?
	_ = "`json:\"name,???\"`"

	// 6. Exclude field tag?
	_ = "`json:\"???\"`"

	// 7. Create encoder?
	_ = "json.???(w)"

	// 8. Create decoder?
	_ = "json.???(r)"

	// 9. Raw JSON for delayed parsing?
	_ = "json.???"

	// 10. JSON numbers unmarshal to?
	_ = "??? (not int!)"

	// 11. nil pointer marshals to?
	_ = "???"

	// 12. Embedded struct in JSON?
	_ = "??? (flat/nested?)"
}

// =============================================================================
// MINI PROJECT: API RESPONSE HANDLER
// =============================================================================

func MiniProject() {
	// Build an API response handler that:
	//
	// 1. Defines a response envelope with status, message, and raw data
	// 2. Uses RawMessage for the data field
	// 3. Has different concrete types for success vs error
	// 4. Implements custom time format
	// 5. Uses omitempty appropriately
	//
	// Scaffold:
	//
	// type APIResponse struct {
	//     Success   bool            `json:"success"`
	//     Message   string          `json:"message,omitempty"`
	//     Data      json.RawMessage `json:"data,omitempty"`
	//     Timestamp CustomTime      `json:"timestamp"`
	// }
	//
	// type CustomTime struct {
	//     time.Time
	// }
	//
	// func (ct CustomTime) MarshalJSON() ([]byte, error) {
	//     return json.Marshal(ct.Format(time.RFC3339))
	// }
	//
	// type UserData struct {
	//     ID    int    `json:"id"`
	//     Name  string `json:"name"`
	//     Email string `json:"email,omitempty"`
	// }
	//
	// type ErrorData struct {
	//     Code    string `json:"code"`
	//     Details string `json:"details,omitempty"`
	// }
	//
	// func parseResponse(jsonStr string) {
	//     var resp APIResponse
	//     json.Unmarshal([]byte(jsonStr), &resp)
	//
	//     if resp.Success {
	//         var user UserData
	//         json.Unmarshal(resp.Data, &user)
	//         // Use user
	//     } else {
	//         var errData ErrorData
	//         json.Unmarshal(resp.Data, &errData)
	//         // Handle error
	//     }
	// }
}

// AllPractice is not meant to be called - this file is for reading.
func AllPractice() {
	PracticeTypeMapping()
	PracticeMarshal()
	PracticeUnmarshal()
	PracticeStructTags()
	PracticeEmbedding()
	PracticePointersAndNil()
	PracticeDynamicJSON()
	PracticeRawMessage()
	PracticeCustomMarshal()
	PracticeStreaming()
	PracticePatterns()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
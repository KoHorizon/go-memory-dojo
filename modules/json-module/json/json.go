// Package json provides comprehensive documentation and working examples
// for Go's encoding/json package - marshaling and unmarshaling JSON data.
package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF ENCODING/JSON
// =============================================================================
//
// JSON (JavaScript Object Notation) is the lingua franca of web APIs.
// Go's encoding/json package provides:
//
//   Marshal   - Go value → JSON bytes
//   Unmarshal - JSON bytes → Go value
//
// THE TYPE MAPPING
//
//   Go              JSON
//   ──────────────────────────
//   bool            true/false
//   int, float      number
//   string          string
//   []T             array
//   map[string]T    object
//   struct          object
//   nil pointer     null
//   nil slice/map   null
//
// EXPORTED FIELDS ONLY
//
// Only exported fields (Capitalized) are marshaled/unmarshaled.
// Unexported fields are invisible to the json package.
//
// =============================================================================

// =============================================================================
// SECTION 2: BASIC MARSHALING
// =============================================================================
//
// json.Marshal(v) → ([]byte, error)
//
// Converts a Go value to JSON bytes.
// Returns error if the value can't be marshaled (e.g., channels, functions).
//
// json.MarshalIndent(v, prefix, indent) → ([]byte, error)
//
// Same as Marshal but with pretty-printing.
//
// =============================================================================

func DemonstrateMarshal() {
	// Simple types
	b, _ := json.Marshal(true)
	fmt.Printf("bool: %s\n", b)

	n, _ := json.Marshal(42)
	fmt.Printf("int: %s\n", n)

	s, _ := json.Marshal("hello")
	fmt.Printf("string: %s\n", s)

	// Slices become arrays
	slice, _ := json.Marshal([]int{1, 2, 3})
	fmt.Printf("slice: %s\n", slice)

	// Maps become objects
	m, _ := json.Marshal(map[string]int{"a": 1, "b": 2})
	fmt.Printf("map: %s\n", m)

	// Structs become objects
	type Person struct {
		Name string
		Age  int
	}
	p, _ := json.Marshal(Person{Name: "Alice", Age: 30})
	fmt.Printf("struct: %s\n", p)

	// Pretty printing
	data := map[string]interface{}{
		"name": "Bob",
		"tags": []string{"go", "json"},
	}
	pretty, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("pretty:\n%s\n", pretty)

	// nil becomes null
	var nilSlice []int
	null, _ := json.Marshal(nilSlice)
	fmt.Printf("nil slice: %s\n", null)
}

// =============================================================================
// SECTION 3: BASIC UNMARSHALING
// =============================================================================
//
// json.Unmarshal(data, &v) → error
//
// Parses JSON bytes into a Go value.
// CRITICAL: v must be a pointer!
//
// BEHAVIOR
//
// - Missing JSON fields → Go field keeps zero value
// - Extra JSON fields → ignored (unless using DisallowUnknownFields)
// - Type mismatch → error
//
// =============================================================================

func DemonstrateUnmarshal() {
	// Into a struct
	type Person struct {
		Name string
		Age  int
	}

	jsonData := []byte(`{"Name": "Alice", "Age": 30}`)
	var p Person
	json.Unmarshal(jsonData, &p) // Must pass pointer!
	fmt.Printf("Struct: %+v\n", p)

	// Missing fields keep zero value
	partial := []byte(`{"Name": "Bob"}`)
	var p2 Person
	json.Unmarshal(partial, &p2)
	fmt.Printf("Partial: %+v (Age is zero)\n", p2)

	// Extra fields are ignored
	extra := []byte(`{"Name": "Carol", "Age": 25, "City": "NYC"}`)
	var p3 Person
	json.Unmarshal(extra, &p3)
	fmt.Printf("Extra ignored: %+v\n", p3)

	// Into a map (dynamic JSON)
	var m map[string]interface{}
	json.Unmarshal([]byte(`{"key": "value", "num": 42}`), &m)
	fmt.Printf("Map: %v\n", m)

	// Into a slice
	var nums []int
	json.Unmarshal([]byte(`[1, 2, 3, 4, 5]`), &nums)
	fmt.Printf("Slice: %v\n", nums)

	// Type mismatch error
	var wrongType int
	err := json.Unmarshal([]byte(`"not a number"`), &wrongType)
	fmt.Printf("Type error: %v\n", err)
}

// =============================================================================
// SECTION 4: STRUCT TAGS
// =============================================================================
//
// Struct tags control JSON field names and behavior.
//
// SYNTAX: `json:"fieldname,options"`
//
// OPTIONS
//
//   json:"name"        - Use "name" instead of field name
//   json:"name,omitempty" - Omit if value is empty/zero
//   json:"-"           - Never marshal/unmarshal this field
//   json:"-,"          - Use "-" as the actual field name
//   json:",string"     - Marshal number/bool as string (unmarshal too)
//
// EMPTY VALUES (for omitempty)
//
//   false, 0, "", nil, empty array, empty map
//
// =============================================================================

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"-"` // Never in JSON
	Age       int    `json:"age,omitempty"`
	IsActive  bool   `json:"is_active"`
	Score     int    `json:"score,string"` // As JSON string
	CreatedAt string `json:"created_at"`
}

func DemonstrateStructTags() {
	// Custom field names
	u := User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "secret123", // Won't appear in JSON
		Age:      30,
		IsActive: true,
		Score:    100,
	}
	data, _ := json.MarshalIndent(u, "", "  ")
	fmt.Printf("With tags:\n%s\n", data)

	// omitempty in action
	u2 := User{
		ID:   2,
		Name: "Bob",
		// Email omitted - will be excluded
		// Age is 0 - will be excluded
	}
	data2, _ := json.Marshal(u2)
	fmt.Printf("With omitempty: %s\n", data2)

	// Unmarshal respects tags too
	jsonStr := `{"id": 3, "name": "Carol", "score": "200"}`
	var u3 User
	json.Unmarshal([]byte(jsonStr), &u3)
	fmt.Printf("Unmarshaled: ID=%d, Name=%s, Score=%d\n", u3.ID, u3.Name, u3.Score)
}

// =============================================================================
// SECTION 5: EMBEDDED STRUCTS
// =============================================================================
//
// Embedded structs are "flattened" into the parent's JSON.
//
//   type Address struct {
//       City string `json:"city"`
//   }
//   type Person struct {
//       Name string `json:"name"`
//       Address     // Embedded - fields appear at top level
//   }
//
// JSON: {"name": "Alice", "city": "NYC"}
//
// NOT: {"name": "Alice", "Address": {"city": "NYC"}}
//
// =============================================================================

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Employee struct {
	Name    string `json:"name"`
	Address        // Embedded - flattened
}

type EmployeeNested struct {
	Name    string  `json:"name"`
	Address Address `json:"address"` // Named field - nested
}

func DemonstrateEmbedding() {
	// Embedded (flattened)
	e := Employee{
		Name: "Alice",
		Address: Address{
			Street: "123 Main St",
			City:   "NYC",
		},
	}
	data, _ := json.Marshal(e)
	fmt.Printf("Embedded (flat): %s\n", data)

	// Named field (nested)
	en := EmployeeNested{
		Name: "Bob",
		Address: Address{
			Street: "456 Oak Ave",
			City:   "LA",
		},
	}
	dataN, _ := json.Marshal(en)
	fmt.Printf("Named (nested): %s\n", dataN)

	// Unmarshal embedded
	jsonFlat := `{"name": "Carol", "street": "789 Pine Rd", "city": "Chicago"}`
	var e2 Employee
	json.Unmarshal([]byte(jsonFlat), &e2)
	fmt.Printf("Unmarshaled: %+v\n", e2)
}

// =============================================================================
// SECTION 6: POINTERS AND NIL
// =============================================================================
//
// POINTERS
//
// - nil pointer → null in JSON
// - Pointer to value → marshals the value itself
// - Use pointers + omitempty to distinguish "not set" from "zero value"
//
// THE NULL PROBLEM
//
//   type Person struct {
//       Age int `json:"age,omitempty"`
//   }
//
//   Age=0 → field omitted (but what if 0 is a valid age?)
//
//   type Person struct {
//       Age *int `json:"age,omitempty"`
//   }
//
//   Age=nil → omitted
//   Age=&zero → "age": 0
//
// =============================================================================

func DemonstratePointersAndNil() {
	// Pointer to value
	name := "Alice"
	type PtrExample struct {
		Name *string `json:"name"`
	}
	data, _ := json.Marshal(PtrExample{Name: &name})
	fmt.Printf("Pointer to value: %s\n", data)

	// nil pointer becomes null
	data2, _ := json.Marshal(PtrExample{Name: nil})
	fmt.Printf("nil pointer: %s\n", data2)

	// Distinguishing "not set" from "zero value"
	type Config struct {
		Enabled *bool `json:"enabled,omitempty"`
		Count   *int  `json:"count,omitempty"`
	}

	// Not set
	c1 := Config{}
	d1, _ := json.Marshal(c1)
	fmt.Printf("Not set: %s\n", d1)

	// Explicitly false/zero
	f := false
	zero := 0
	c2 := Config{Enabled: &f, Count: &zero}
	d2, _ := json.Marshal(c2)
	fmt.Printf("Explicit zero: %s\n", d2)

	// Unmarshal null
	var c3 Config
	json.Unmarshal([]byte(`{"enabled": null}`), &c3)
	fmt.Printf("Unmarshal null: Enabled=%v\n", c3.Enabled)
}

// =============================================================================
// SECTION 7: INTERFACES AND DYNAMIC JSON
// =============================================================================
//
// UNKNOWN STRUCTURE: map[string]interface{}
//
// When you don't know the JSON structure at compile time.
//
// TYPE ASSERTIONS
//
// JSON types unmarshal to:
//   object  → map[string]interface{}
//   array   → []interface{}
//   string  → string
//   number  → float64 (!)
//   bool    → bool
//   null    → nil
//
// NUMBER GOTCHA: All JSON numbers become float64!
//
// =============================================================================

func DemonstrateDynamicJSON() {
	// Unknown structure
	jsonStr := `{
		"name": "Alice",
		"age": 30,
		"active": true,
		"tags": ["go", "json"],
		"meta": {"key": "value"}
	}`

	var data map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &data)

	// Type assertions
	name := data["name"].(string)
	age := data["age"].(float64) // Note: float64, not int!
	active := data["active"].(bool)
	tags := data["tags"].([]interface{})
	meta := data["meta"].(map[string]interface{})

	fmt.Printf("name: %s (type: %T)\n", name, name)
	fmt.Printf("age: %.0f (type: %T)\n", age, age)
	fmt.Printf("active: %v (type: %T)\n", active, active)
	fmt.Printf("tags: %v (type: %T)\n", tags, tags)
	fmt.Printf("meta: %v (type: %T)\n", meta, meta)

	// Safe type assertion
	if score, ok := data["score"].(float64); ok {
		fmt.Printf("score: %f\n", score)
	} else {
		fmt.Println("score not found or wrong type")
	}

	// json.Number for precise numbers
	dec := json.NewDecoder(strings.NewReader(`{"big": 9223372036854775807}`))
	dec.UseNumber()
	var data2 map[string]interface{}
	dec.Decode(&data2)
	num := data2["big"].(json.Number)
	bigInt, _ := num.Int64()
	fmt.Printf("Big int: %d\n", bigInt)
}

// =============================================================================
// SECTION 8: json.RawMessage
// =============================================================================
//
// WHAT IS RawMessage?
//
// A raw encoded JSON value. Delays parsing or preserves original encoding.
//
// USE CASES
//
// 1. Delay parsing until you know the type
// 2. Re-marshal without changes
// 3. Partial unmarshaling (parse some fields, keep others raw)
//
// =============================================================================

type Envelope struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type TextMessage struct {
	Text string `json:"text"`
}

type ImageMessage struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func DemonstrateRawMessage() {
	// Delayed parsing based on type
	messages := []string{
		`{"type": "text", "payload": {"text": "Hello!"}}`,
		`{"type": "image", "payload": {"url": "pic.jpg", "width": 100, "height": 200}}`,
	}

	for _, msg := range messages {
		var env Envelope
		json.Unmarshal([]byte(msg), &env)

		switch env.Type {
		case "text":
			var txt TextMessage
			json.Unmarshal(env.Payload, &txt)
			fmt.Printf("Text: %s\n", txt.Text)
		case "image":
			var img ImageMessage
			json.Unmarshal(env.Payload, &img)
			fmt.Printf("Image: %s (%dx%d)\n", img.URL, img.Width, img.Height)
		}
	}

	// Preserve raw JSON through marshal/unmarshal cycle
	original := `{"type":"data","payload":{"complex":{"nested":"value"}}}`
	var env Envelope
	json.Unmarshal([]byte(original), &env)
	reencoded, _ := json.Marshal(env)
	fmt.Printf("Preserved: %s\n", reencoded)
}

// =============================================================================
// SECTION 9: CUSTOM MARSHALING
// =============================================================================
//
// INTERFACES
//
//   json.Marshaler   - type has MarshalJSON() ([]byte, error)
//   json.Unmarshaler - type has UnmarshalJSON([]byte) error
//
// USE CASES
//
// - Custom time formats
// - Enums as strings
// - Complex transformations
// - Validation during unmarshal
//
// =============================================================================

type Status int

const (
	StatusPending Status = iota
	StatusActive
	StatusComplete
)

func (s Status) MarshalJSON() ([]byte, error) {
	var str string
	switch s {
	case StatusPending:
		str = "pending"
	case StatusActive:
		str = "active"
	case StatusComplete:
		str = "complete"
	default:
		str = "unknown"
	}
	return json.Marshal(str)
}

func (s *Status) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	switch str {
	case "pending":
		*s = StatusPending
	case "active":
		*s = StatusActive
	case "complete":
		*s = StatusComplete
	default:
		return fmt.Errorf("unknown status: %s", str)
	}
	return nil
}

type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Format("2006-01-02"))
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func DemonstrateCustomMarshal() {
	// Enum as string
	type Task struct {
		Name   string `json:"name"`
		Status Status `json:"status"`
	}

	t := Task{Name: "Do stuff", Status: StatusActive}
	data, _ := json.Marshal(t)
	fmt.Printf("Enum: %s\n", data)

	// Unmarshal enum
	var t2 Task
	json.Unmarshal([]byte(`{"name": "Other", "status": "complete"}`), &t2)
	fmt.Printf("Unmarshaled: %+v\n", t2)

	// Custom time format
	type Event struct {
		Name string     `json:"name"`
		Date CustomTime `json:"date"`
	}

	e := Event{Name: "Meeting", Date: CustomTime{time.Now()}}
	data2, _ := json.Marshal(e)
	fmt.Printf("Custom time: %s\n", data2)
}

// =============================================================================
// SECTION 10: STREAMING WITH ENCODER/DECODER
// =============================================================================
//
// json.NewEncoder(w io.Writer) → *Encoder
// json.NewDecoder(r io.Reader) → *Decoder
//
// USE CASES
//
// - Writing JSON to HTTP response
// - Reading JSON from HTTP request
// - Processing large JSON files line by line (JSON Lines format)
// - Memory efficiency (no intermediate []byte)
//
// DECODER OPTIONS
//
//   dec.DisallowUnknownFields() - Error on unknown fields
//   dec.UseNumber()             - Keep numbers as json.Number
//
// =============================================================================

func DemonstrateStreaming() {
	// Encoder - write to buffer
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")

	enc.Encode(map[string]string{"hello": "world"})
	enc.Encode(map[string]int{"count": 42})
	fmt.Printf("Encoded:\n%s", buf.String())

	// Decoder - read from reader
	jsonStream := `{"name": "Alice"}
{"name": "Bob"}
{"name": "Carol"}`

	dec := json.NewDecoder(strings.NewReader(jsonStream))

	for {
		var obj map[string]string
		if err := dec.Decode(&obj); err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Printf("Decoded: %v\n", obj)
	}

	// DisallowUnknownFields
	strict := json.NewDecoder(strings.NewReader(`{"name": "Test", "unknown": 123}`))
	strict.DisallowUnknownFields()

	type Strict struct {
		Name string `json:"name"`
	}
	var s Strict
	if err := strict.Decode(&s); err != nil {
		fmt.Printf("Strict error: %v\n", err)
	}

	// UseNumber for precision
	numDec := json.NewDecoder(strings.NewReader(`{"id": 9223372036854775807}`))
	numDec.UseNumber()
	var numData map[string]interface{}
	numDec.Decode(&numData)
	fmt.Printf("Number type: %T\n", numData["id"])
}

// =============================================================================
// SECTION 11: COMMON PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: HTTP request/response
	_ = `
	// Reading JSON request body
	var req RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Writing JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	`

	// Pattern 2: Config file
	_ = `
	data, _ := os.ReadFile("config.json")
	var config Config
	json.Unmarshal(data, &config)
	`

	// Pattern 3: Deep copy via JSON
	type Data struct {
		Values []int
	}
	original := Data{Values: []int{1, 2, 3}}
	bytes, _ := json.Marshal(original)
	var copied Data
	json.Unmarshal(bytes, &copied)
	copied.Values[0] = 999
	fmt.Printf("Original: %v, Copy: %v\n", original.Values, copied.Values)

	// Pattern 4: Default values
	type ConfigWithDefaults struct {
		Host    string `json:"host"`
		Port    int    `json:"port"`
		Timeout int    `json:"timeout"`
	}
	defaults := ConfigWithDefaults{Host: "localhost", Port: 8080, Timeout: 30}
	json.Unmarshal([]byte(`{"port": 9090}`), &defaults) // Only overrides port
	fmt.Printf("With defaults: %+v\n", defaults)

	// Pattern 5: Optional fields with pointer
	type Patch struct {
		Name  *string `json:"name,omitempty"`
		Email *string `json:"email,omitempty"`
	}
	var patch Patch
	json.Unmarshal([]byte(`{"name": "NewName"}`), &patch)
	if patch.Name != nil {
		fmt.Printf("Update name to: %s\n", *patch.Name)
	}
	if patch.Email != nil {
		fmt.Printf("Update email to: %s\n", *patch.Email)
	} else {
		fmt.Println("Email not in patch, don't update")
	}
}

// =============================================================================
// SECTION 12: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Forgetting to pass pointer
	type Data struct {
		Value int
	}
	var d Data
	_ = json.Unmarshal([]byte(`{"Value": 42}`), d) // WRONG: not a pointer
	_ = json.Unmarshal([]byte(`{"Value": 42}`), &d) // RIGHT
	fmt.Println("Mistake 1: Pass pointer to Unmarshal")

	// Mistake 2: Unexported fields
	type private struct {
		Public  int // Will be marshaled
		private int // Will NOT be marshaled
	}
	p := private{Public: 1, private: 2}
	data, _ := json.Marshal(p)
	fmt.Printf("Mistake 2: Unexported ignored: %s\n", data)

	// Mistake 3: Expecting int, getting float64
	var m map[string]interface{}
	json.Unmarshal([]byte(`{"count": 42}`), &m)
	// count := m["count"].(int)  // PANIC! It's float64
	count := m["count"].(float64)
	fmt.Printf("Mistake 3: Numbers are float64: %T\n", count)

	// Mistake 4: Not handling errors
	_ = `
	// BAD
	json.Unmarshal(data, &v)

	// GOOD
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	`
	fmt.Println("Mistake 4: Always check Unmarshal errors")

	// Mistake 5: Tag typo
	type Typo struct {
		Name string `json: "name"` // WRONG: space after colon
	}
	fmt.Println("Mistake 5: No space in struct tags")
}

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateMarshal()
	DemonstrateUnmarshal()
	DemonstrateStructTags()
	DemonstrateEmbedding()
	DemonstratePointersAndNil()
	DemonstrateDynamicJSON()
	DemonstrateRawMessage()
	DemonstrateCustomMarshal()
	DemonstrateStreaming()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
}
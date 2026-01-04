// Package interfacesmod provides comprehensive documentation and working examples
// for understanding Go interfaces - the foundation of Go's polymorphism.
package interfacesmod

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// =============================================================================
// SECTION 1: WHAT IS AN INTERFACE?
// =============================================================================
//
// AN INTERFACE IS A CONTRACT
//
// An interface defines a set of method signatures. Any type that implements
// all those methods satisfies the interface - automatically, implicitly.
//
//   type Reader interface {
//       Read(p []byte) (n int, err error)
//   }
//
// Any type with a Read method matching this signature IS a Reader.
// No "implements" keyword. No explicit declaration. Just methods.
//
// WHY INTERFACES MATTER
//
// 1. Decoupling: Code depends on behavior, not concrete types
// 2. Testing: Easy to mock dependencies
// 3. Flexibility: Swap implementations without changing callers
// 4. Composition: Build complex behaviors from simple ones
//
// THE GOLDEN RULE
//
//   "Accept interfaces, return structs"
//
// Functions should accept interfaces (flexible) but return concrete types
// (clear, no hidden implementations).
//
// =============================================================================

// =============================================================================
// SECTION 2: DEFINING AND IMPLEMENTING INTERFACES
// =============================================================================
//
// DEFINING AN INTERFACE
//
//   type InterfaceName interface {
//       Method1(params) returnType
//       Method2(params) (returnType, error)
//   }
//
// IMPLEMENTING AN INTERFACE
//
// Just implement the methods. No explicit declaration needed.
//
//   type MyType struct { ... }
//
//   func (m MyType) Method1(params) returnType { ... }
//   func (m MyType) Method2(params) (returnType, error) { ... }
//
// Now MyType satisfies InterfaceName automatically.
//
// =============================================================================

// Speaker is a simple interface
type Speaker interface {
	Speak() string
}

// Dog implements Speaker
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says woof!"
}

// Cat implements Speaker
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + " says meow!"
}

// Robot also implements Speaker
type Robot struct {
	ID string
}

func (r Robot) Speak() string {
	return "Robot " + r.ID + " says beep boop"
}

func DemonstrateBasicInterface() {
	// All three types satisfy Speaker
	var speakers []Speaker = []Speaker{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Robot{ID: "R2D2"},
	}

	// Polymorphism: same method call, different behavior
	for _, s := range speakers {
		fmt.Println(s.Speak())
	}

	// Function accepting interface
	greet := func(s Speaker) {
		fmt.Println("Greeting:", s.Speak())
	}

	greet(Dog{Name: "Max"})
	greet(Cat{Name: "Luna"})
}

// =============================================================================
// SECTION 3: IMPLICIT SATISFACTION
// =============================================================================
//
// GO'S KEY INSIGHT: IMPLICIT INTERFACES
//
// In Go, interface satisfaction is implicit and checked at compile time.
//
//   // Java/C# way (explicit)
//   class Dog implements Speaker { ... }
//
//   // Go way (implicit)
//   type Dog struct{}
//   func (d Dog) Speak() string { ... }
//   // Dog automatically satisfies Speaker!
//
// BENEFITS OF IMPLICIT INTERFACES
//
// 1. Types can satisfy interfaces they don't know about
// 2. Interfaces can be defined after types exist
// 3. No import dependencies between interface and implementation
// 4. Easier to create small, focused interfaces
//
// COMPILE-TIME CHECK
//
// Use this pattern to verify interface satisfaction at compile time:
//
//   var _ Speaker = Dog{}        // Compile error if Dog doesn't satisfy Speaker
//   var _ Speaker = (*Dog)(nil)  // For pointer receivers
//
// =============================================================================

// Compile-time interface satisfaction checks
var _ Speaker = Dog{}
var _ Speaker = Cat{}
var _ Speaker = Robot{}

// Stringer is from fmt package - anything with String() string
// Our types can satisfy it without importing fmt in their definition

func (d Dog) String() string {
	return fmt.Sprintf("Dog(%s)", d.Name)
}

func DemonstrateImplicitSatisfaction() {
	// Dog now satisfies both Speaker AND fmt.Stringer
	dog := Dog{Name: "Buddy"}

	// As Speaker
	var s Speaker = dog
	fmt.Println(s.Speak())

	// As Stringer (fmt.Println uses this automatically)
	fmt.Println(dog) // Calls String() method

	// A type can satisfy many interfaces
	// without knowing about any of them
}

// =============================================================================
// SECTION 4: THE EMPTY INTERFACE
// =============================================================================
//
// interface{} AND any
//
// The empty interface has no methods, so EVERY type satisfies it.
//
//   interface{}  // Old syntax
//   any          // Go 1.18+ alias (preferred)
//
// USE CASES
//
// - Generic containers (before generics)
// - JSON unmarshaling to unknown structure
// - Printf-style variadic functions
// - When you truly don't know the type
//
// DOWNSIDES
//
// - No compile-time type safety
// - Requires type assertions to use values
// - Hides bugs until runtime
//
// RULE: Avoid interface{}/any when possible
//       Prefer specific interfaces or generics (Go 1.18+)
//
// =============================================================================

func DemonstrateEmptyInterface() {
	// any can hold anything
	var anything any

	anything = 42
	fmt.Printf("Int: %v (type: %T)\n", anything, anything)

	anything = "hello"
	fmt.Printf("String: %v (type: %T)\n", anything, anything)

	anything = Dog{Name: "Buddy"}
	fmt.Printf("Dog: %v (type: %T)\n", anything, anything)

	// Slice of any
	mixed := []any{1, "two", 3.0, true, Dog{Name: "Max"}}
	for i, v := range mixed {
		fmt.Printf("  [%d] %v (type: %T)\n", i, v, v)
	}

	// Map with any values
	data := map[string]any{
		"name":   "Alice",
		"age":    30,
		"active": true,
	}
	fmt.Printf("Data: %v\n", data)
}

// =============================================================================
// SECTION 5: TYPE ASSERTIONS
// =============================================================================
//
// TYPE ASSERTION: Extract concrete type from interface
//
//   value := interfaceVar.(ConcreteType)      // Panics if wrong type
//   value, ok := interfaceVar.(ConcreteType)  // Safe - ok is false if wrong
//
// ALWAYS USE THE COMMA-OK FORM unless you're certain of the type.
//
// =============================================================================

func DemonstrateTypeAssertions() {
	var s Speaker = Dog{Name: "Buddy"}

	// Unsafe assertion (panics if wrong)
	dog := s.(Dog)
	fmt.Printf("Dog name: %s\n", dog.Name)

	// Safe assertion with comma-ok
	if dog, ok := s.(Dog); ok {
		fmt.Printf("It's a dog: %s\n", dog.Name)
	} else {
		fmt.Println("Not a dog")
	}

	// Check for Cat (will be false)
	if cat, ok := s.(Cat); ok {
		fmt.Printf("It's a cat: %s\n", cat.Name)
	} else {
		fmt.Println("Not a cat")
	}

	// Type assertion to interface
	// Check if Speaker also implements fmt.Stringer
	if stringer, ok := s.(fmt.Stringer); ok {
		fmt.Printf("String representation: %s\n", stringer.String())
	}

	// Common pattern: process based on type
	process := func(v any) {
		if s, ok := v.(string); ok {
			fmt.Printf("String of length %d: %s\n", len(s), s)
		} else if n, ok := v.(int); ok {
			fmt.Printf("Integer doubled: %d\n", n*2)
		} else {
			fmt.Printf("Unknown type: %T\n", v)
		}
	}

	process("hello")
	process(42)
	process(3.14)
}

// =============================================================================
// SECTION 6: TYPE SWITCHES
// =============================================================================
//
// TYPE SWITCH: Multi-way type assertion
//
//   switch v := interfaceVar.(type) {
//   case Type1:
//       // v is Type1
//   case Type2:
//       // v is Type2
//   case Type3, Type4:
//       // v is Type3 or Type4
//   default:
//       // v is original interface type
//   }
//
// More readable than chained type assertions.
//
// =============================================================================

func DemonstrateTypeSwitch() {
	describe := func(v any) string {
		switch x := v.(type) {
		case nil:
			return "nil"
		case bool:
			if x {
				return "boolean true"
			}
			return "boolean false"
		case int:
			return fmt.Sprintf("integer %d", x)
		case float64:
			return fmt.Sprintf("float %.2f", x)
		case string:
			return fmt.Sprintf("string %q (len=%d)", x, len(x))
		case []int:
			return fmt.Sprintf("int slice with %d elements", len(x))
		case Dog:
			return fmt.Sprintf("Dog named %s", x.Name)
		case Cat:
			return fmt.Sprintf("Cat named %s", x.Name)
		case Speaker:
			return fmt.Sprintf("Some speaker: %s", x.Speak())
		default:
			return fmt.Sprintf("unknown type %T", x)
		}
	}

	values := []any{
		nil,
		true,
		42,
		3.14159,
		"hello",
		[]int{1, 2, 3},
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Robot{ID: "R2"},
	}

	for _, v := range values {
		fmt.Printf("  %v → %s\n", v, describe(v))
	}
}

// =============================================================================
// SECTION 7: INTERFACE COMPOSITION
// =============================================================================
//
// EMBEDDING INTERFACES
//
// Interfaces can embed other interfaces to compose larger contracts.
//
//   type Reader interface {
//       Read(p []byte) (n int, err error)
//   }
//
//   type Writer interface {
//       Write(p []byte) (n int, err error)
//   }
//
//   type ReadWriter interface {
//       Reader  // Embedded
//       Writer  // Embedded
//   }
//
// This is how io.ReadWriter, io.ReadCloser, etc. are defined.
//
// =============================================================================

// Mover can move
type Mover interface {
	Move() string
}

// SpeakingMover combines Speaker and Mover
type SpeakingMover interface {
	Speaker
	Mover
}

// Horse implements both Speaker and Mover
type Horse struct {
	Name string
}

func (h Horse) Speak() string {
	return h.Name + " says neigh!"
}

func (h Horse) Move() string {
	return h.Name + " gallops"
}

// Verify Horse satisfies SpeakingMover
var _ SpeakingMover = Horse{}

func DemonstrateInterfaceComposition() {
	horse := Horse{Name: "Spirit"}

	// Horse satisfies Speaker
	var speaker Speaker = horse
	fmt.Println(speaker.Speak())

	// Horse satisfies Mover
	var mover Mover = horse
	fmt.Println(mover.Move())

	// Horse satisfies SpeakingMover (both)
	var sm SpeakingMover = horse
	fmt.Println(sm.Speak())
	fmt.Println(sm.Move())

	// Real example: io.ReadWriter
	var buf bytes.Buffer
	var rw io.ReadWriter = &buf

	rw.Write([]byte("hello"))
	data := make([]byte, 5)
	rw.Read(data)
	fmt.Printf("Read from buffer: %s\n", data)
}

// =============================================================================
// SECTION 8: COMMON STANDARD LIBRARY INTERFACES
// =============================================================================
//
// io.Reader
//   Read(p []byte) (n int, err error)
//   Files, network connections, buffers, strings.Reader
//
// io.Writer
//   Write(p []byte) (n int, err error)
//   Files, network connections, buffers, os.Stdout
//
// io.Closer
//   Close() error
//   Files, network connections
//
// io.ReadWriter, io.ReadCloser, io.WriteCloser, io.ReadWriteCloser
//   Combinations of above
//
// fmt.Stringer
//   String() string
//   Custom string representation for fmt.Print*
//
// error
//   Error() string
//   All error types
//
// sort.Interface
//   Len() int
//   Less(i, j int) bool
//   Swap(i, j int)
//   For custom sorting
//
// =============================================================================

// Person implements fmt.Stringer
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

// ValidationError implements error
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

// People implements sort.Interface
type People []Person

func (p People) Len() int           { return len(p) }
func (p People) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func DemonstrateStandardInterfaces() {
	// fmt.Stringer
	person := Person{Name: "Alice", Age: 30}
	fmt.Println(person) // Uses String() method

	// error interface
	err := ValidationError{Field: "email", Message: "invalid format"}
	fmt.Printf("Error: %v\n", err)

	// Also works with errors.Is pattern if you implement it
	var e error = err
	fmt.Printf("As error: %v\n", e)

	// sort.Interface
	people := People{
		{Name: "Charlie", Age: 25},
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 20},
	}
	sort.Sort(people)
	fmt.Println("Sorted by age:", people)

	// io.Reader - strings.Reader satisfies it
	reader := strings.NewReader("Hello, World!")
	data := make([]byte, 5)
	reader.Read(data)
	fmt.Printf("Read: %s\n", data)

	// io.Writer - bytes.Buffer satisfies it
	var buf bytes.Buffer
	buf.Write([]byte("Hello"))
	buf.WriteString(" World")
	fmt.Printf("Buffer: %s\n", buf.String())
}

// =============================================================================
// SECTION 9: POINTER VS VALUE RECEIVERS
// =============================================================================
//
// METHOD RECEIVERS AND INTERFACE SATISFACTION
//
// Value receiver: Both T and *T satisfy the interface
// Pointer receiver: Only *T satisfies the interface
//
//   func (t T) Method()   // T and *T both work
//   func (t *T) Method()  // Only *T works
//
// WHY?
//
// A value receiver can be called on both values and pointers (Go auto-dereferences).
// A pointer receiver needs addressable values, which interface values might not be.
//
// =============================================================================

// Counter with pointer receiver (needs to modify state)
type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}

// Incrementer interface
type Incrementer interface {
	Increment()
}

func DemonstrateReceivers() {
	// Pointer receiver: only *Counter satisfies Incrementer
	var inc Incrementer = &Counter{} // Must use pointer
	// var inc Incrementer = Counter{} // COMPILE ERROR!

	inc.Increment()
	inc.Increment()
	fmt.Printf("Count: %d\n", inc.(*Counter).Value())

	// Value receiver example
	dog := Dog{Name: "Buddy"}
	var s1 Speaker = dog  // Value works
	var s2 Speaker = &dog // Pointer also works
	fmt.Println(s1.Speak())
	fmt.Println(s2.Speak())

	// Rule of thumb:
	// - Use pointer receiver if method modifies receiver
	// - Use pointer receiver for large structs (avoid copying)
	// - Be consistent within a type
}

// =============================================================================
// SECTION 10: THE NIL INTERFACE GOTCHA
// =============================================================================
//
// CRITICAL: Interface values have two components
//
//   interface value = (type, value)
//
// An interface is nil ONLY if both type and value are nil.
//
//   var s Speaker           // (nil, nil) - this IS nil
//   var d *Dog = nil
//   var s Speaker = d       // (*Dog, nil) - this is NOT nil!
//
// This is a common source of bugs, especially with error returns.
//
// =============================================================================

func DemonstrateNilInterface() {
	// Truly nil interface
	var s1 Speaker
	fmt.Printf("s1 == nil: %v\n", s1 == nil) // true

	// Interface holding nil pointer - NOT nil!
	var d *Dog = nil
	var s2 Speaker = d
	fmt.Printf("s2 == nil: %v\n", s2 == nil) // false!
	fmt.Printf("s2 type: %T, value: %v\n", s2, s2)

	// This is dangerous with errors
	returnsError := func(fail bool) error {
		var err *ValidationError = nil
		if fail {
			err = &ValidationError{Field: "test", Message: "failed"}
		}
		return err // BUG: Always returns non-nil error!
	}

	err := returnsError(false)
	if err != nil {
		fmt.Println("Got 'error' even though nothing failed!")
		fmt.Printf("  Type: %T, Value: %v\n", err, err)
	}

	// Correct way
	returnsErrorCorrect := func(fail bool) error {
		if fail {
			return &ValidationError{Field: "test", Message: "failed"}
		}
		return nil // Return literal nil
	}

	err2 := returnsErrorCorrect(false)
	if err2 == nil {
		fmt.Println("Correctly returned nil")
	}
}

// =============================================================================
// SECTION 11: INTERFACE DESIGN PRINCIPLES
// =============================================================================
//
// 1. KEEP INTERFACES SMALL
//
//    Prefer many small interfaces over few large ones.
//    io.Reader has 1 method. That's intentional.
//
// 2. ACCEPT INTERFACES, RETURN STRUCTS
//
//    Functions should accept interfaces for flexibility
//    but return concrete types for clarity.
//
// 3. DEFINE INTERFACES AT POINT OF USE
//
//    Define interfaces in the package that uses them,
//    not the package that implements them.
//
// 4. DON'T EXPORT INTERFACES FOR JUST ONE IMPLEMENTATION
//
//    If there's only one implementation, you don't need an interface.
//    Add interfaces when you need polymorphism or testing.
//
// 5. NAME INTERFACES BY BEHAVIOR
//
//    Reader, Writer, Closer, Stringer, Handler
//    Not IReader, ReaderInterface
//
// =============================================================================

// Good: Small, focused interface
type Processor interface {
	Process(data []byte) ([]byte, error)
}

// Good: Accept interface, return struct
func ProcessData(p Processor, input []byte) ([]byte, error) {
	return p.Process(input)
}

// UppercaseProcessor is a concrete implementation
type UppercaseProcessor struct{}

func (u UppercaseProcessor) Process(data []byte) ([]byte, error) {
	return bytes.ToUpper(data), nil
}

// Good: Interface defined where it's used (not with implementation)
// This allows the caller to define exactly what they need

func DemonstrateDesignPrinciples() {
	// User of Processor doesn't need to know about UppercaseProcessor
	var p Processor = UppercaseProcessor{}
	result, _ := ProcessData(p, []byte("hello"))
	fmt.Printf("Processed: %s\n", result)

	// Easy to test with mock
	// mockProcessor := MockProcessor{...}
	// ProcessData(mockProcessor, input)
}

// =============================================================================
// SECTION 12: TESTING WITH INTERFACES
// =============================================================================
//
// Interfaces make testing easy by allowing mock implementations.
//
// PATTERN:
//
// 1. Define interface for dependency
// 2. Production code uses real implementation
// 3. Tests use mock implementation
//
// =============================================================================

// DataStore interface for dependency injection
type DataStore interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

// Service depends on DataStore interface, not concrete type
type Service struct {
	store DataStore
}

func (s *Service) GetValue(key string) (string, error) {
	return s.store.Get(key)
}

// MockDataStore for testing
type MockDataStore struct {
	data map[string]string
}

func (m *MockDataStore) Get(key string) (string, error) {
	if v, ok := m.data[key]; ok {
		return v, nil
	}
	return "", fmt.Errorf("key not found: %s", key)
}

func (m *MockDataStore) Set(key string, value string) error {
	if m.data == nil {
		m.data = make(map[string]string)
	}
	m.data[key] = value
	return nil
}

func DemonstrateTesting() {
	// In production: use real database
	// realStore := &PostgresDataStore{...}
	// service := &Service{store: realStore}

	// In tests: use mock
	mockStore := &MockDataStore{
		data: map[string]string{
			"name": "Alice",
			"city": "NYC",
		},
	}
	service := &Service{store: mockStore}

	// Test the service
	name, _ := service.GetValue("name")
	fmt.Printf("Got name: %s\n", name)

	city, _ := service.GetValue("city")
	fmt.Printf("Got city: %s\n", city)

	_, err := service.GetValue("missing")
	fmt.Printf("Missing key error: %v\n", err)
}

// =============================================================================
// SECTION 13: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Returning interface when concrete type works
	_ = `
	// WRONG - hides implementation details unnecessarily
	func NewService() ServiceInterface {
		return &ConcreteService{}
	}

	// RIGHT - return concrete type
	func NewService() *ConcreteService {
		return &ConcreteService{}
	}
	`
	fmt.Println("Mistake 1: Return concrete types, accept interfaces")

	// Mistake 2: Interface pollution - too many interfaces
	_ = `
	// WRONG - interface for everything
	type UserInterface interface { ... }
	type OrderInterface interface { ... }
	
	// RIGHT - interfaces only when needed for:
	// - Multiple implementations
	// - Testing/mocking
	// - Decoupling packages
	`
	fmt.Println("Mistake 2: Don't create interfaces unless needed")

	// Mistake 3: Forgetting nil interface gotcha
	_ = `
	// WRONG
	func GetError() error {
		var err *MyError = nil
		return err  // NOT nil!
	}

	// RIGHT
	func GetError() error {
		return nil  // Return literal nil
	}
	`
	fmt.Println("Mistake 3: Return literal nil for nil errors")

	// Mistake 4: Using interface{} instead of generics
	_ = `
	// OLD WAY (pre Go 1.18)
	func Contains(slice []interface{}, item interface{}) bool

	// NEW WAY (Go 1.18+)
	func Contains[T comparable](slice []T, item T) bool
	`
	fmt.Println("Mistake 4: Use generics instead of interface{} when possible")

	// Mistake 5: Large interfaces
	_ = `
	// WRONG - too many methods
	type Repository interface {
		Get() ...
		GetAll() ...
		Create() ...
		Update() ...
		Delete() ...
		Search() ...
		// ... 20 more methods
	}

	// RIGHT - small, focused interfaces
	type Getter interface { Get() ... }
	type Creator interface { Create() ... }
	type Deleter interface { Delete() ... }
	`
	fmt.Println("Mistake 5: Keep interfaces small and focused")

	// Mistake 6: Unsafe type assertion
	_ = `
	// WRONG - panics if wrong type
	dog := speaker.(Dog)

	// RIGHT - safe with comma-ok
	dog, ok := speaker.(Dog)
	if !ok {
		// Handle wrong type
	}
	`
	fmt.Println("Mistake 6: Always use comma-ok for type assertions")
}

// =============================================================================
// SECTION 14: SUMMARY - INTERFACE MENTAL MODEL
// =============================================================================
//
// WHAT IS AN INTERFACE?
//   A contract defining behavior (methods), not data.
//
// HOW DOES A TYPE SATISFY AN INTERFACE?
//   By implementing all methods. Implicit, automatic.
//
// WHAT IS interface{}/any?
//   Empty interface - satisfied by all types. Use sparingly.
//
// HOW TO EXTRACT CONCRETE TYPE?
//   Type assertion: value, ok := iface.(Type)
//   Type switch: switch v := iface.(type) { ... }
//
// INTERFACE VALUE = (TYPE, VALUE)
//   nil only if both are nil!
//
// DESIGN PRINCIPLES:
//   - Small interfaces (1-3 methods)
//   - Accept interfaces, return structs
//   - Define at point of use
//   - Name by behavior (-er suffix)
//
// =============================================================================

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateBasicInterface()
	DemonstrateImplicitSatisfaction()
	DemonstrateEmptyInterface()
	DemonstrateTypeAssertions()
	DemonstrateTypeSwitch()
	DemonstrateInterfaceComposition()
	DemonstrateStandardInterfaces()
	DemonstrateReceivers()
	DemonstrateNilInterface()
	DemonstrateDesignPrinciples()
	DemonstrateTesting()
	DemonstrateCommonMistakes()
}
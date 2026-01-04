// Package interfaces_practice is your daily practice space for Go interfaces.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against interfaces/interfaces.go
// 3. Note what you missed - focus on those tomorrow
package interfaces_practice

// =============================================================================
// EXERCISE 1: WHAT IS AN INTERFACE?
// =============================================================================

func PracticeWhatIsInterface() {
	// TODO: What is an interface in Go?
	// Answer: A ??? defining a set of ???

	// TODO: How does a type "implement" an interface in Go?
	// Answer: By implementing all its ???. No ??? keyword needed.

	// TODO: What is the golden rule?
	// Answer: "Accept ???, return ???"

	// TODO: Why use interfaces?
	// 1. ???
	// 2. ???
	// 3. ???
}

// =============================================================================
// EXERCISE 2: DEFINING INTERFACES
// =============================================================================

func PracticeDefiningInterfaces() {
	// TODO: Define an interface with one method
	// type ??? interface {
	//     MethodName(params) returnType
	// }

	// TODO: Define an interface with multiple methods
	// type ??? interface {
	//     Method1() string
	//     Method2(int) error
	// }

	// TODO: How to compose interfaces (embedding)?
	// type Combined interface {
	//     Interface1
	//     ???
	// }
}

// =============================================================================
// EXERCISE 3: IMPLEMENTING INTERFACES
// =============================================================================

func PracticeImplementingInterfaces() {
	// TODO: Implement an interface (no explicit declaration needed)
	// type MyType struct{}
	// func (m MyType) ???() string {
	//     return "implemented"
	// }

	// TODO: Compile-time check that type satisfies interface
	// var _ InterfaceName = ???{}
	// var _ InterfaceName = (*???)??? // For pointer receiver
}

// =============================================================================
// EXERCISE 4: EMPTY INTERFACE
// =============================================================================

func PracticeEmptyInterface() {
	// TODO: Two ways to write empty interface
	// Old: ???{}
	// New (Go 1.18+): ???

	// TODO: What types satisfy empty interface?
	// Answer: ???

	// TODO: When to use empty interface?
	// 1. ???
	// 2. ???

	// TODO: Why avoid empty interface when possible?
	// Answer: No compile-time ???, requires ??? assertions
}

// =============================================================================
// EXERCISE 5: TYPE ASSERTIONS
// =============================================================================

func PracticeTypeAssertions() {
	// TODO: Unsafe type assertion (panics if wrong)
	// value := interfaceVar.(???)

	// TODO: Safe type assertion with comma-ok
	// value, ok := interfaceVar.(???)

	// TODO: What does ok mean?
	// true: ???
	// false: ???

	// TODO: Type assert to another interface
	// stringer, ok := speaker.(fmt.???)
}

// =============================================================================
// EXERCISE 6: TYPE SWITCHES
// =============================================================================

func PracticeTypeSwitches() {
	// TODO: Write a type switch
	// switch v := interfaceVar.(???) {
	// case Type1:
	//     // v is ???
	// case Type2, Type3:
	//     // v is ??? or ???
	// default:
	//     // v is ???
	// }

	// TODO: What keyword is used in type switch?
	// .(???)
}

// =============================================================================
// EXERCISE 7: COMMON STANDARD INTERFACES
// =============================================================================

func PracticeStandardInterfaces() {
	// TODO: io.Reader method signature
	// ???(p []byte) (n int, err error)

	// TODO: io.Writer method signature
	// ???(p []byte) (n int, err error)

	// TODO: io.Closer method signature
	// ???() error

	// TODO: fmt.Stringer method signature
	// ???() string

	// TODO: error interface method signature
	// ???() string

	// TODO: sort.Interface methods (3)
	// ???() int
	// ???(i, j int) bool
	// ???(i, j int)
}

// =============================================================================
// EXERCISE 8: POINTER VS VALUE RECEIVERS
// =============================================================================

func PracticeReceivers() {
	// TODO: Value receiver - which types satisfy interface?
	// func (t T) Method()
	// Answer: Both ??? and ??? satisfy

	// TODO: Pointer receiver - which types satisfy interface?
	// func (t *T) Method()
	// Answer: Only ??? satisfies

	// TODO: Why this difference?
	// Answer: Pointer receiver needs ??? values

	// TODO: When to use pointer receiver? (2 reasons)
	// 1. When method ??? the receiver
	// 2. For ??? structs (avoid copying)
}

// =============================================================================
// EXERCISE 9: NIL INTERFACE GOTCHA
// =============================================================================

func PracticeNilInterface() {
	// TODO: What are the two components of an interface value?
	// Answer: (???, ???)

	// TODO: When is an interface value nil?
	// Answer: Only when ??? are nil

	// TODO: What's the bug here?
	// var err *MyError = nil
	// return err  // As error interface
	// Answer: Returns ???-nil because type is ???

	// TODO: How to fix it?
	// Answer: Return literal ???

	// TODO: Why is this dangerous with errors?
	// if err != nil { ... }
	// Answer: Condition is ??? even for nil pointer
}

// =============================================================================
// EXERCISE 10: INTERFACE DESIGN PRINCIPLES
// =============================================================================

func PracticeDesignPrinciples() {
	// TODO: Principle 1 - Interface size
	// Answer: Keep interfaces ??? (1-3 methods)

	// TODO: Principle 2 - Parameters and returns
	// Answer: Accept ???, return ???

	// TODO: Principle 3 - Where to define interfaces
	// Answer: At point of ???, not with ???

	// TODO: Principle 4 - When to create interfaces
	// Answer: Only when needed for ??? or ???

	// TODO: Principle 5 - Naming convention
	// Answer: Name by ??? (-er suffix: Reader, Writer, ???)
}

// =============================================================================
// EXERCISE 11: TESTING WITH INTERFACES
// =============================================================================

func PracticeTesting() {
	// TODO: Pattern for testable code
	// 1. Define ??? for dependency
	// 2. Production uses ??? implementation
	// 3. Tests use ??? implementation

	// TODO: Why interfaces help testing
	// Answer: Easy to create ??? implementations
}

// =============================================================================
// EXERCISE 12: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// func NewService() ServiceInterface {
	//     return &ConcreteService{}
	// }
	// Answer: Should return ???, not interface

	// Mistake 2: What's wrong?
	// type UserInterface interface { ... }  // For every type
	// Answer: Interface ???, only create when needed

	// Mistake 3: What's wrong?
	// func GetError() error {
	//     var err *MyError = nil
	//     return err
	// }
	// Answer: Returns ??? error, should return ???

	// Mistake 4: What's wrong?
	// dog := speaker.(Dog)
	// Answer: No ???-ok, will ??? if wrong type

	// Mistake 5: What's wrong?
	// type Repository interface {
	//     // 20 methods...
	// }
	// Answer: Interface too ???, should be ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// BASICS
	// 1. How does Go know if type implements interface?
	_ = "By having all ???"

	// 2. Empty interface syntax (new)?
	_ = "???"

	// 3. Empty interface syntax (old)?
	_ = "???{}"

	// TYPE ASSERTIONS
	// 4. Safe type assertion?
	_ = "value, ok := iface.(???)"

	// 5. Type switch keyword?
	_ = "switch v := iface.(???) { ... }"

	// STANDARD INTERFACES
	// 6. io.Reader method?
	_ = "???([]byte) (int, error)"

	// 7. io.Writer method?
	_ = "???([]byte) (int, error)"

	// 8. fmt.Stringer method?
	_ = "???() string"

	// 9. error method?
	_ = "???() string"

	// RECEIVERS
	// 10. Value receiver satisfies?
	_ = "Both T and ???"

	// 11. Pointer receiver satisfies?
	_ = "Only ???"

	// NIL INTERFACE
	// 12. Interface value components?
	_ = "(???, ???)"

	// 13. Interface is nil when?
	_ = "Both ??? and ??? are nil"

	// DESIGN
	// 14. Golden rule?
	_ = "Accept ???, return ???"

	// 15. Interface naming?
	_ = "-??? suffix (Reader, Writer)"

	// COMPILE CHECK
	// 16. Verify interface satisfaction?
	_ = "var _ Interface = ???{}"
}

// =============================================================================
// MINI PROJECT: PLUGIN SYSTEM
// =============================================================================

func MiniProject() {
	// Build a simple plugin system that:
	//
	// 1. Defines a Plugin interface
	// 2. Has multiple plugin implementations
	// 3. Loads and runs plugins
	// 4. Handles plugin errors gracefully
	//
	// Scaffold:
	//
	// // Plugin is the interface all plugins must implement
	// type Plugin interface {
	//     Name() string
	//     Init() error
	//     Execute(data []byte) ([]byte, error)
	//     Close() error
	// }
	//
	// // PluginManager manages plugins
	// type PluginManager struct {
	//     plugins map[string]Plugin
	// }
	//
	// func NewPluginManager() *PluginManager {
	//     return &PluginManager{
	//         plugins: make(map[string]Plugin),
	//     }
	// }
	//
	// func (pm *PluginManager) Register(p Plugin) error {
	//     name := p.Name()
	//     if _, exists := pm.plugins[name]; exists {
	//         return fmt.Errorf("plugin %s already registered", name)
	//     }
	//     if err := p.Init(); err != nil {
	//         return fmt.Errorf("failed to init plugin %s: %w", name, err)
	//     }
	//     pm.plugins[name] = p
	//     return nil
	// }
	//
	// func (pm *PluginManager) Execute(name string, data []byte) ([]byte, error) {
	//     p, ok := pm.plugins[name]
	//     if !ok {
	//         return nil, fmt.Errorf("plugin %s not found", name)
	//     }
	//     return p.Execute(data)
	// }
	//
	// func (pm *PluginManager) CloseAll() error {
	//     var errs []error
	//     for name, p := range pm.plugins {
	//         if err := p.Close(); err != nil {
	//             errs = append(errs, fmt.Errorf("%s: %w", name, err))
	//         }
	//     }
	//     if len(errs) > 0 {
	//         return fmt.Errorf("errors closing plugins: %v", errs)
	//     }
	//     return nil
	// }
	//
	// // UppercasePlugin converts data to uppercase
	// type UppercasePlugin struct{}
	//
	// func (p *UppercasePlugin) Name() string { return "uppercase" }
	// func (p *UppercasePlugin) Init() error { return nil }
	// func (p *UppercasePlugin) Execute(data []byte) ([]byte, error) {
	//     return bytes.ToUpper(data), nil
	// }
	// func (p *UppercasePlugin) Close() error { return nil }
	//
	// // ReversePlugin reverses the data
	// type ReversePlugin struct{}
	//
	// func (p *ReversePlugin) Name() string { return "reverse" }
	// func (p *ReversePlugin) Init() error { return nil }
	// func (p *ReversePlugin) Execute(data []byte) ([]byte, error) {
	//     result := make([]byte, len(data))
	//     for i, b := range data {
	//         result[len(data)-1-i] = b
	//     }
	//     return result, nil
	// }
	// func (p *ReversePlugin) Close() error { return nil }
	//
	// // Base64Plugin encodes to base64
	// type Base64Plugin struct{}
	//
	// func (p *Base64Plugin) Name() string { return "base64" }
	// func (p *Base64Plugin) Init() error { return nil }
	// func (p *Base64Plugin) Execute(data []byte) ([]byte, error) {
	//     encoded := base64.StdEncoding.EncodeToString(data)
	//     return []byte(encoded), nil
	// }
	// func (p *Base64Plugin) Close() error { return nil }
	//
	// // Compile-time checks
	// var _ Plugin = (*UppercasePlugin)(nil)
	// var _ Plugin = (*ReversePlugin)(nil)
	// var _ Plugin = (*Base64Plugin)(nil)
	//
	// // Test:
	// func main() {
	//     pm := NewPluginManager()
	//
	//     // Register plugins
	//     pm.Register(&UppercasePlugin{})
	//     pm.Register(&ReversePlugin{})
	//     pm.Register(&Base64Plugin{})
	//
	//     input := []byte("hello world")
	//
	//     // Execute each plugin
	//     result1, _ := pm.Execute("uppercase", input)
	//     fmt.Printf("Uppercase: %s\n", result1)
	//
	//     result2, _ := pm.Execute("reverse", input)
	//     fmt.Printf("Reverse: %s\n", result2)
	//
	//     result3, _ := pm.Execute("base64", input)
	//     fmt.Printf("Base64: %s\n", result3)
	//
	//     // Chain plugins
	//     chained := input
	//     for _, name := range []string{"uppercase", "reverse"} {
	//         chained, _ = pm.Execute(name, chained)
	//     }
	//     fmt.Printf("Chained: %s\n", chained)
	//
	//     pm.CloseAll()
	// }
}

// AllPractice is not meant to be called - this file is for reading.
func AllPractice() {
	PracticeWhatIsInterface()
	PracticeDefiningInterfaces()
	PracticeImplementingInterfaces()
	PracticeEmptyInterface()
	PracticeTypeAssertions()
	PracticeTypeSwitches()
	PracticeStandardInterfaces()
	PracticeReceivers()
	PracticeNilInterface()
	PracticeDesignPrinciples()
	PracticeTesting()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
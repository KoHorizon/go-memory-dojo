// Package concurrency_practice is your daily practice space for Go concurrency.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against concurrency/concurrency.go
// 3. Note what you missed - focus on those tomorrow
package concurrency_practice

// =============================================================================
// EXERCISE 1: PHILOSOPHY
// =============================================================================

func PracticePhilosophy() {
	// TODO: Complete Go's concurrency mantra
	// "Do not communicate by ??? ???; ??? ??? by ???."

	// TODO: What are Go's three concurrency building blocks?
	// 1. ??? - lightweight threads
	// 2. ??? - typed conduits for communication
	// 3. ??? - multiplex channel operations

	// TODO: When to use channels vs mutex vs atomic?
	// Channels: ???
	// Mutex: ???
	// Atomic: ???
}

// =============================================================================
// EXERCISE 2: GOROUTINES
// =============================================================================

func PracticeGoroutines() {
	// TODO: Start a goroutine with a named function
	// ??? functionName()

	// TODO: Start a goroutine with an anonymous function
	// ??? func() { ... }()

	// TODO: What happens when main() returns?
	// Answer: ???

	// TODO: What's wrong with this code?
	// for i := 0; i < 3; i++ {
	//     go func() { fmt.Println(i) }()
	// }
	// Answer: ???

	// TODO: How do you fix it?
	// for i := 0; i < 3; i++ {
	//     go func(???) { fmt.Println(???) }(???)
	// }
}

// =============================================================================
// EXERCISE 3: CHANNELS - BASICS
// =============================================================================

func PracticeChannelBasics() {
	// TODO: Create an unbuffered channel of ints
	// ch := ???(chan int)

	// TODO: Create a buffered channel of strings with capacity 10
	// ch := ???(chan string, ???)

	// TODO: Send a value to a channel
	// ch ??? value

	// TODO: Receive a value from a channel
	// value := ???ch

	// TODO: Close a channel
	// ???(ch)

	// TODO: Complete the channel axioms
	// 1. Send to nil channel: ???
	// 2. Receive from nil channel: ???
	// 3. Send to closed channel: ???
	// 4. Receive from closed channel: ???
	// 5. Close nil channel: ???
	// 6. Close already-closed channel: ???
}

// =============================================================================
// EXERCISE 4: CHANNELS - DIRECTION
// =============================================================================

func PracticeChannelDirection() {
	// TODO: Function signature for send-only channel
	// func produce(out chan??? int)

	// TODO: Function signature for receive-only channel
	// func consume(in ???chan int)

	// TODO: Why use directional channels?
	// Answer: ???
}

// =============================================================================
// EXERCISE 5: CLOSING AND RANGING
// =============================================================================

func PracticeClosingChannels() {
	// TODO: Detect if channel is closed
	// value, ??? := <-ch
	// if !??? {
	//     // Channel is closed
	// }

	// TODO: Range over channel until closed
	// for value := ??? ch {
	//     // Process value
	// }

	// TODO: Who should close a channel - sender or receiver?
	// Answer: ???

	// TODO: What does closing signal?
	// Answer: ???
}

// =============================================================================
// EXERCISE 6: SELECT
// =============================================================================

func PracticeSelect() {
	// TODO: Basic select structure
	// ??? {
	// case v := <-ch1:
	//     // ...
	// case ch2 <- x:
	//     // ...
	// ???:
	//     // Non-blocking
	// }

	// TODO: Timeout pattern
	// select {
	// case v := <-ch:
	//     // Got value
	// case <-time.???(1 * time.Second):
	//     // Timeout
	// }

	// TODO: Cancellation pattern
	// select {
	// case v := <-ch:
	//     // Got value
	// case <-ctx.???():
	//     // Cancelled
	// }

	// TODO: If multiple cases are ready, which is chosen?
	// Answer: ???

	// TODO: What does default do?
	// Answer: ???
}

// =============================================================================
// EXERCISE 7: sync.WaitGroup
// =============================================================================

func PracticeWaitGroup() {
	// TODO: WaitGroup methods
	// wg.???(n)    // Increment counter
	// wg.???()     // Decrement counter
	// wg.???()     // Block until zero

	// TODO: Write the standard WaitGroup pattern
	// var wg sync.WaitGroup
	// for i := 0; i < n; i++ {
	//     wg.???(1)   // WHERE does this go?
	//     go func() {
	//         ??? wg.???()
	//         // Do work
	//     }()
	// }
	// wg.???()

	// TODO: Why must Add() be called before starting the goroutine?
	// Answer: ???
}

// =============================================================================
// EXERCISE 8: sync.Mutex
// =============================================================================

func PracticeMutex() {
	// TODO: Mutex methods
	// mu.???()    // Acquire lock
	// mu.???()    // Release lock

	// TODO: Write the safe pattern with defer
	// mu.???()
	// ??? mu.???()
	// // Access shared state

	// TODO: RWMutex methods
	// mu.???() / mu.???()  // For reading (shared)
	// mu.???() / mu.???()  // For writing (exclusive)

	// TODO: When use RWMutex vs Mutex?
	// Answer: ???
}

// =============================================================================
// EXERCISE 9: sync.Once
// =============================================================================

func PracticeOnce() {
	// TODO: How do you use sync.Once?
	// var once sync.Once
	// once.???(func() {
	//     // Runs exactly once
	// })

	// TODO: What's the use case for sync.Once?
	// Answer: ???
}

// =============================================================================
// EXERCISE 10: sync/atomic
// =============================================================================

func PracticeAtomics() {
	// TODO: Old-style atomic operations
	// atomic.???(addr, delta)  // Add
	// atomic.???(addr)         // Read
	// atomic.???(addr, val)    // Write

	// TODO: New-style (Go 1.19+)
	// var counter atomic.???
	// counter.???(1)
	// counter.???()
	// counter.???(value)

	// TODO: When use atomics vs mutex?
	// Answer: ???
}

// =============================================================================
// EXERCISE 11: COMMON PATTERNS
// =============================================================================

func PracticePatterns() {
	// TODO: Worker pool pattern
	// 1. Create ??? channel for jobs
	// 2. Create ??? channel for results
	// 3. Start N worker goroutines that ??? jobs
	// 4. Send jobs, ??? jobs channel
	// 5. Collect results

	// TODO: Semaphore pattern (bounded concurrency)
	// sem := make(chan struct{}, maxConcurrent)
	// sem ??? struct{}{}   // Acquire
	// ???sem               // Release

	// TODO: Done channel for cancellation
	// done := make(chan struct{})
	// // In goroutine:
	// select {
	// case <-???:
	//     return
	// default:
	//     // Work
	// }
	// // To cancel:
	// ???(done)
}

// =============================================================================
// EXERCISE 12: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// for i := 0; i < 3; i++ {
	//     go func() { fmt.Println(i) }()
	// }
	// Answer: ???

	// Mistake 2: What's wrong?
	// ch := make(chan int)
	// close(ch)
	// ch <- 1
	// Answer: ???

	// Mistake 3: What's wrong?
	// go func() {
	//     wg.Add(1)
	//     defer wg.Done()
	// }()
	// Answer: ???

	// Mistake 4: What's wrong?
	// mu.Lock()
	// if err != nil {
	//     return err
	// }
	// mu.Unlock()
	// Answer: ???

	// Mistake 5: What's wrong?
	// var mu sync.Mutex
	// mu2 := mu
	// Answer: ???

	// Mistake 6: What's wrong?
	// go func() {
	//     <-ch  // ch is never closed or sent to
	// }()
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// 1. Start a goroutine?
	_ = "??? func() {...}()"

	// 2. Create unbuffered channel?
	_ = "???(chan int)"

	// 3. Create buffered channel?
	_ = "???(chan int, size)"

	// 4. Send to channel?
	_ = "ch ??? value"

	// 5. Receive from channel?
	_ = "value := ???ch"

	// 6. Check if channel closed?
	_ = "value, ??? := <-ch"

	// 7. Range over channel?
	_ = "for v := ??? ch {...}"

	// 8. WaitGroup pattern?
	_ = "wg.Add(1); defer wg.???(); wg.???()"

	// 9. Mutex pattern?
	_ = "mu.???(); defer mu.???()"

	// 10. Select with timeout?
	_ = "case <-time.???(duration):"

	// 11. Channel axiom: send to closed?
	_ = "???"

	// 12. Channel axiom: receive from closed?
	_ = "???"
}

// =============================================================================
// MINI PROJECT: PARALLEL DOWNLOADER
// =============================================================================

func MiniProject() {
	// Build a parallel downloader that:
	//
	// 1. Takes a list of URLs
	// 2. Downloads max 3 concurrently (semaphore)
	// 3. Collects results through a channel
	// 4. Handles timeouts with context
	// 5. Waits for all to complete with WaitGroup
	//
	// Scaffold:
	//
	// func download(ctx context.Context, urls []string) []Result {
	//     results := make(chan Result, len(urls))
	//     sem := make(chan struct{}, 3)  // Max 3 concurrent
	//     var wg sync.WaitGroup
	//
	//     for _, url := range urls {
	//         wg.Add(1)
	//         go func(u string) {
	//             defer wg.Done()
	//
	//             sem <- struct{}{}        // Acquire
	//             defer func() { <-sem }() // Release
	//
	//             select {
	//             case <-ctx.Done():
	//                 results <- Result{URL: u, Err: ctx.Err()}
	//                 return
	//             default:
	//             }
	//
	//             // Simulate download
	//             time.Sleep(100 * time.Millisecond)
	//             results <- Result{URL: u, Data: "content"}
	//         }(url)
	//     }
	//
	//     go func() {
	//         wg.Wait()
	//         close(results)
	//     }()
	//
	//     var out []Result
	//     for r := range results {
	//         out = append(out, r)
	//     }
	//     return out
	// }
	//
	// type Result struct {
	//     URL  string
	//     Data string
	//     Err  error
	// }
}

// RunAllPractice is not meant to be called - this file is for reading.
func RunAllPractice() {
	PracticePhilosophy()
	PracticeGoroutines()
	PracticeChannelBasics()
	PracticeChannelDirection()
	PracticeClosingChannels()
	PracticeSelect()
	PracticeWaitGroup()
	PracticeMutex()
	PracticeOnce()
	PracticeAtomics()
	PracticePatterns()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
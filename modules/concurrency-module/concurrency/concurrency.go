// Package concurrency provides comprehensive documentation and working examples
// for Go's concurrency primitives - goroutines, channels, and sync package.
package concurrency

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF GO CONCURRENCY
// =============================================================================
//
// "DO NOT COMMUNICATE BY SHARING MEMORY; SHARE MEMORY BY COMMUNICATING."
//
// This is Go's concurrency mantra. Instead of:
//   - Threads fighting over shared variables with locks
//
// Go encourages:
//   - Goroutines passing data through channels
//
// THE BUILDING BLOCKS
//
//   Goroutines - Lightweight threads (start with 2KB stack, grows as needed)
//   Channels   - Typed conduits for communication between goroutines
//   Select     - Multiplex operations on multiple channels
//   sync pkg   - Low-level primitives (Mutex, WaitGroup, Once, etc.)
//
// WHEN TO USE WHAT
//
//   Channels: Passing ownership of data, signaling, coordinating goroutines
//   Mutex:    Protecting shared state that can't be passed around
//   Atomic:   Simple counters, flags (faster than mutex for simple ops)
//
// =============================================================================

// =============================================================================
// SECTION 2: GOROUTINES
// =============================================================================
//
// WHAT IS A GOROUTINE?
//
// A goroutine is a lightweight thread managed by Go runtime.
// Starts with ~2KB stack (vs ~1MB for OS thread).
// Can have thousands running concurrently.
//
// STARTING A GOROUTINE
//
//   go functionName()           // Named function
//   go func() { ... }()         // Anonymous function (closure)
//   go obj.Method()             // Method call
//
// CRITICAL: GOROUTINES DON'T WAIT
//
// When main() returns, all goroutines are killed immediately.
// You must coordinate completion (WaitGroup, channels, etc.)
//
// CLOSURE GOTCHA
//
//   for i := 0; i < 3; i++ {
//       go func() { fmt.Println(i) }()  // BUG: captures variable, not value
//   }
//   // Likely prints: 3 3 3
//
//   for i := 0; i < 3; i++ {
//       go func(n int) { fmt.Println(n) }(i)  // FIX: pass as parameter
//   }
//   // Prints: 0 1 2 (in some order)
//
// =============================================================================

func DemonstrateGoroutines() {
	// Basic goroutine
	go func() {
		fmt.Println("Hello from goroutine")
	}()

	// Closure gotcha - WRONG way
	fmt.Println("Closure gotcha (wrong):")
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("  wrong: %d\n", i) // Captures variable, not value
		}()
	}
	time.Sleep(10 * time.Millisecond)

	// Closure gotcha - CORRECT way
	fmt.Println("Closure gotcha (correct):")
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Printf("  correct: %d\n", n) // Value passed as parameter
		}(i)
	}
	time.Sleep(10 * time.Millisecond)

	// Waiting for goroutines with WaitGroup
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Worker %d done\n", n)
		}(i)
	}
	wg.Wait()
	fmt.Println("All workers finished")
}

// =============================================================================
// SECTION 3: CHANNELS - BASICS
// =============================================================================
//
// WHAT IS A CHANNEL?
//
// A typed conduit for passing values between goroutines.
// Channels synchronize execution - sends and receives block.
//
// CREATING CHANNELS
//
//   ch := make(chan int)      // Unbuffered - send blocks until receive
//   ch := make(chan int, 10)  // Buffered - send blocks when buffer full
//
// OPERATIONS
//
//   ch <- value    // Send
//   value := <-ch  // Receive
//   close(ch)      // Close - no more sends allowed
//
// CHANNEL AXIOMS
//
//   1. Send to nil channel blocks forever
//   2. Receive from nil channel blocks forever
//   3. Send to closed channel panics
//   4. Receive from closed channel returns zero value immediately
//   5. Close nil channel panics
//   6. Close already-closed channel panics
//
// =============================================================================

func DemonstrateChannelBasics() {
	// Unbuffered channel - synchronous
	ch := make(chan string)

	go func() {
		ch <- "hello" // Blocks until received
	}()

	msg := <-ch // Blocks until sent
	fmt.Printf("Received: %s\n", msg)

	// Buffered channel - async up to capacity
	buffered := make(chan int, 3)
	buffered <- 1 // Doesn't block
	buffered <- 2 // Doesn't block
	buffered <- 3 // Doesn't block
	// buffered <- 4 // Would block - buffer full

	fmt.Printf("Buffered: %d, %d, %d\n", <-buffered, <-buffered, <-buffered)

	// Channel direction in function signatures
	produce := func(out chan<- int) { // Send-only
		out <- 42
	}
	consume := func(in <-chan int) int { // Receive-only
		return <-in
	}

	ch2 := make(chan int, 1)
	produce(ch2)
	result := consume(ch2)
	fmt.Printf("Directional channels: %d\n", result)
}

// =============================================================================
// SECTION 4: CHANNELS - CLOSING AND RANGING
// =============================================================================
//
// WHY CLOSE CHANNELS?
//
// Closing signals "no more values will be sent."
// Receivers can detect this and stop waiting.
//
// DETECTING CLOSED CHANNELS
//
//   value, ok := <-ch
//   if !ok {
//       // Channel is closed
//   }
//
// RANGE OVER CHANNEL
//
//   for value := range ch {
//       // Receives until channel is closed
//   }
//
// WHO CLOSES?
//
// Only the SENDER should close. Never the receiver.
// Closing is about signaling "done sending", not "done receiving."
//
// =============================================================================

func DemonstrateClosingChannels() {
	ch := make(chan int)

	// Producer closes when done
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Receiver detects close with ok
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Printf("Received: %d\n", value)
	}

	// Range - cleaner way to receive until close
	ch2 := make(chan string)
	go func() {
		ch2 <- "one"
		ch2 <- "two"
		ch2 <- "three"
		close(ch2)
	}()

	for msg := range ch2 {
		fmt.Printf("Range received: %s\n", msg)
	}
}

// =============================================================================
// SECTION 5: SELECT
// =============================================================================
//
// WHAT IS SELECT?
//
// Multiplexes operations on multiple channels.
// Like switch, but for channel operations.
//
// HOW IT WORKS
//
//   select {
//   case v := <-ch1:
//       // Received from ch1
//   case ch2 <- x:
//       // Sent to ch2
//   default:
//       // No channel ready (non-blocking)
//   }
//
// BEHAVIOR
//
// - If multiple cases ready: one chosen at random
// - If none ready: blocks (unless default exists)
// - default makes it non-blocking
//
// COMMON PATTERNS
//
// Timeout:
//   select {
//   case v := <-ch:
//       // Got value
//   case <-time.After(1 * time.Second):
//       // Timeout
//   }
//
// Cancellation:
//   select {
//   case v := <-ch:
//       // Got value
//   case <-ctx.Done():
//       // Cancelled
//   }
//
// =============================================================================

func DemonstrateSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	// Receive from whichever is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Printf("Select: %s\n", msg)
		case msg := <-ch2:
			fmt.Printf("Select: %s\n", msg)
		}
	}

	// Timeout pattern
	ch3 := make(chan int)
	select {
	case v := <-ch3:
		fmt.Printf("Got: %d\n", v)
	case <-time.After(10 * time.Millisecond):
		fmt.Println("Timeout!")
	}

	// Non-blocking with default
	ch4 := make(chan int)
	select {
	case v := <-ch4:
		fmt.Printf("Got: %d\n", v)
	default:
		fmt.Println("Nothing ready, continuing...")
	}

	// Cancellation pattern
	ctx, cancel := context.WithCancel(context.Background())
	ch5 := make(chan int)

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	select {
	case v := <-ch5:
		fmt.Printf("Got: %d\n", v)
	case <-ctx.Done():
		fmt.Println("Cancelled")
	}
}

// =============================================================================
// SECTION 6: sync.WaitGroup
// =============================================================================
//
// WHAT IS WAITGROUP?
//
// A counter for waiting on a collection of goroutines.
//
// METHODS
//
//   wg.Add(n)   // Increment counter by n
//   wg.Done()   // Decrement counter by 1 (same as Add(-1))
//   wg.Wait()   // Block until counter reaches 0
//
// THE PATTERN
//
//   var wg sync.WaitGroup
//   for i := 0; i < n; i++ {
//       wg.Add(1)
//       go func() {
//           defer wg.Done()
//           // Do work
//       }()
//   }
//   wg.Wait()
//
// CRITICAL: Add BEFORE starting goroutine, not inside it.
//
// =============================================================================

func DemonstrateWaitGroup() {
	var wg sync.WaitGroup

	// Standard pattern
	for i := 0; i < 3; i++ {
		wg.Add(1) // Add BEFORE goroutine
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n*10) * time.Millisecond)
			fmt.Printf("Worker %d complete\n", n)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers done")

	// Common mistake: Add inside goroutine
	// DON'T DO THIS:
	// go func() {
	//     wg.Add(1)  // Race condition! Wait() might return early
	//     defer wg.Done()
	// }()
}

// =============================================================================
// SECTION 7: sync.Mutex
// =============================================================================
//
// WHAT IS MUTEX?
//
// Mutual exclusion lock. Only one goroutine can hold it at a time.
// Used to protect shared state from concurrent access.
//
// METHODS
//
//   mu.Lock()     // Acquire lock (blocks if held)
//   mu.Unlock()   // Release lock
//   mu.TryLock()  // Try to acquire, returns bool (Go 1.18+)
//
// THE PATTERN
//
//   var mu sync.Mutex
//   var count int
//
//   mu.Lock()
//   count++
//   mu.Unlock()
//
// Or with defer:
//
//   mu.Lock()
//   defer mu.Unlock()
//   count++
//
// RWMUTEX
//
// sync.RWMutex allows multiple readers OR one writer:
//   mu.RLock() / mu.RUnlock()  - For reading (shared)
//   mu.Lock() / mu.Unlock()    - For writing (exclusive)
//
// =============================================================================

func DemonstrateMutex() {
	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("Counter with mutex: %d\n", counter)

	// Without mutex - data race!
	unsafeCounter := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			unsafeCounter++ // DATA RACE!
		}()
	}
	wg.Wait()
	fmt.Printf("Counter without mutex (race): %d\n", unsafeCounter)

	// RWMutex - multiple readers allowed
	var rwmu sync.RWMutex
	data := make(map[string]int)
	data["key"] = 0

	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		rwmu.Lock()
		defer rwmu.Unlock()
		data["key"]++
		fmt.Println("Writer updated")
	}()

	// Multiple readers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			rwmu.RLock()
			defer rwmu.RUnlock()
			fmt.Printf("Reader %d sees: %d\n", n, data["key"])
		}(i)
	}
	wg.Wait()
}

// =============================================================================
// SECTION 8: sync.Once
// =============================================================================
//
// WHAT IS ONCE?
//
// Ensures a function is called exactly once, even from multiple goroutines.
// Perfect for lazy initialization.
//
// THE PATTERN
//
//   var once sync.Once
//   var instance *Database
//
//   func GetDB() *Database {
//       once.Do(func() {
//           instance = connectToDatabase()
//       })
//       return instance
//   }
//
// =============================================================================

func DemonstrateOnce() {
	var once sync.Once
	var initialized bool

	initialize := func() {
		fmt.Println("Initializing... (should only see once)")
		initialized = true
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			once.Do(initialize) // Only first call executes
			fmt.Printf("Goroutine %d sees initialized: %v\n", n, initialized)
		}(i)
	}
	wg.Wait()
}

// =============================================================================
// SECTION 9: sync/atomic
// =============================================================================
//
// WHAT ARE ATOMICS?
//
// Lock-free operations for simple types.
// Faster than mutex for simple counters/flags.
//
// COMMON FUNCTIONS
//
//   atomic.AddInt64(&n, 1)        // Add
//   atomic.LoadInt64(&n)          // Read
//   atomic.StoreInt64(&n, value)  // Write
//   atomic.SwapInt64(&n, value)   // Swap, return old
//   atomic.CompareAndSwapInt64(&n, old, new)  // CAS
//
// Go 1.19+ atomic types:
//
//   var counter atomic.Int64
//   counter.Add(1)
//   counter.Load()
//   counter.Store(value)
//
// =============================================================================

func DemonstrateAtomics() {
	// Old style (Go 1.18 and earlier)
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("Atomic counter (old): %d\n", atomic.LoadInt64(&counter))

	// New style (Go 1.19+)
	var counter2 atomic.Int64

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter2.Add(1)
		}()
	}
	wg.Wait()
	fmt.Printf("Atomic counter (new): %d\n", counter2.Load())

	// Atomic bool for flags
	var done atomic.Bool
	done.Store(false)

	go func() {
		time.Sleep(10 * time.Millisecond)
		done.Store(true)
	}()

	for !done.Load() {
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Done flag set")
}

// =============================================================================
// SECTION 10: COMMON PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Worker pool
	fmt.Println("Pattern 1: Worker pool")
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start workers
	for w := 0; w < 3; w++ {
		go func(id int) {
			for job := range jobs {
				results <- job * 2
			}
		}(w)
	}

	// Send jobs
	for j := 0; j < 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 0; r < 5; r++ {
		fmt.Printf("  Result: %d\n", <-results)
	}

	// Pattern 2: Fan-out, Fan-in
	fmt.Println("Pattern 2: Fan-out, Fan-in")
	fanOut := func(in <-chan int, n int) []<-chan int {
		outs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			out := make(chan int)
			outs[i] = out
			go func() {
				defer close(out)
				for v := range in {
					out <- v * 2
				}
			}()
		}
		return outs
	}

	fanIn := func(ins ...<-chan int) <-chan int {
		out := make(chan int)
		var wg sync.WaitGroup
		for _, in := range ins {
			wg.Add(1)
			go func(ch <-chan int) {
				defer wg.Done()
				for v := range ch {
					out <- v
				}
			}(in)
		}
		go func() {
			wg.Wait()
			close(out)
		}()
		return out
	}

	input := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			input <- i
		}
		close(input)
	}()

	workers := fanOut(input, 2)
	for v := range fanIn(workers...) {
		fmt.Printf("  Fan result: %d\n", v)
	}

	// Pattern 3: Semaphore (bounded concurrency)
	fmt.Println("Pattern 3: Semaphore")
	sem := make(chan struct{}, 2) // Max 2 concurrent

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sem <- struct{}{}        // Acquire
			defer func() { <-sem }() // Release
			fmt.Printf("  Worker %d running (max 2 concurrent)\n", n)
			time.Sleep(10 * time.Millisecond)
		}(i)
	}
	wg.Wait()

	// Pattern 4: Done channel for cancellation
	fmt.Println("Pattern 4: Done channel")
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("  Worker stopped")
				return
			default:
				// Work
			}
		}
	}()
	time.Sleep(10 * time.Millisecond)
	close(done) // Signal all workers to stop
	time.Sleep(10 * time.Millisecond)

	// Pattern 5: Or-done channel
	fmt.Println("Pattern 5: Pipeline with done")
	orDone := func(done <-chan struct{}, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for {
				select {
				case <-done:
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					select {
					case out <- v:
					case <-done:
						return
					}
				}
			}
		}()
		return out
	}
	_ = orDone
	fmt.Println("  (or-done pattern defined)")
}

// =============================================================================
// SECTION 11: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Loop variable capture
	fmt.Println("Mistake 1: Loop variable capture")
	// WRONG:
	// for i := 0; i < 3; i++ {
	//     go func() { fmt.Println(i) }()  // All print 3!
	// }
	// FIX: Pass as parameter
	fmt.Println("  Pass loop variable as parameter")

	// Mistake 2: Send on closed channel
	fmt.Println("Mistake 2: Send on closed channel panics")
	// ch := make(chan int)
	// close(ch)
	// ch <- 1  // PANIC!
	fmt.Println("  Only sender should close, never receiver")

	// Mistake 3: Closing nil channel
	fmt.Println("Mistake 3: Close nil channel panics")
	// var ch chan int  // nil
	// close(ch)  // PANIC!
	fmt.Println("  Check for nil before closing")

	// Mistake 4: WaitGroup Add inside goroutine
	fmt.Println("Mistake 4: WaitGroup.Add inside goroutine")
	// WRONG:
	// go func() {
	//     wg.Add(1)  // Race! Wait() might return early
	//     defer wg.Done()
	// }()
	fmt.Println("  Add before starting goroutine")

	// Mistake 5: Forgetting to unlock mutex
	fmt.Println("Mistake 5: Forgetting to unlock")
	// mu.Lock()
	// if err != nil {
	//     return err  // DEADLOCK! mu never unlocked
	// }
	// mu.Unlock()
	fmt.Println("  Use defer mu.Unlock() right after Lock()")

	// Mistake 6: Copying sync types
	fmt.Println("Mistake 6: Copying sync types")
	// var mu sync.Mutex
	// mu2 := mu  // DON'T! Copies lock state
	fmt.Println("  Never copy Mutex, WaitGroup, Cond, etc.")

	// Mistake 7: Goroutine leak
	fmt.Println("Mistake 7: Goroutine leak")
	// go func() {
	//     <-ch  // Blocks forever if ch never closed/sent
	// }()
	fmt.Println("  Ensure goroutines can exit (context, done channel)")
}

// =============================================================================
// SECTION 12: sync.Cond (Advanced)
// =============================================================================
//
// WHAT IS COND?
//
// Condition variable - allows goroutines to wait for a condition.
// Used when you need to wait for something more complex than a simple lock.
//
//   cond := sync.NewCond(&mu)
//   cond.Wait()      // Atomically unlock and wait, relock on wakeup
//   cond.Signal()    // Wake one waiting goroutine
//   cond.Broadcast() // Wake all waiting goroutines
//
// =============================================================================

func DemonstrateCond() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false

	// Waiter
	go func() {
		mu.Lock()
		for !ready {
			cond.Wait() // Releases lock, waits, reacquires lock
		}
		fmt.Println("Cond: Waiter woke up, ready is true")
		mu.Unlock()
	}()

	// Signaler
	time.Sleep(10 * time.Millisecond)
	mu.Lock()
	ready = true
	cond.Signal() // Wake one waiter
	mu.Unlock()

	time.Sleep(10 * time.Millisecond)

	// Broadcast example
	var wg sync.WaitGroup
	start := false
	cond2 := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			cond2.L.Lock()
			for !start {
				cond2.Wait()
			}
			cond2.L.Unlock()
			fmt.Printf("Cond: Runner %d started\n", n)
		}(i)
	}

	time.Sleep(10 * time.Millisecond)
	cond2.L.Lock()
	start = true
	cond2.Broadcast() // Wake ALL waiters
	cond2.L.Unlock()
	wg.Wait()
}

// RunAllDemonstrations is not meant to be called - this file is for reading.
func RunAllDemonstrations() {
	DemonstrateGoroutines()
	DemonstrateChannelBasics()
	DemonstrateClosingChannels()
	DemonstrateSelect()
	DemonstrateWaitGroup()
	DemonstrateMutex()
	DemonstrateOnce()
	DemonstrateAtomics()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
	DemonstrateCond()
}
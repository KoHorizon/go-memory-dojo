// Package slicesmod provides comprehensive documentation and working examples
// for Go's slices package (Go 1.21+) and maps package - generic collection utilities.
package slicesmod

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strings"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF SLICES AND MAPS PACKAGES
// =============================================================================
//
// Before Go 1.21, you had to write the same loops repeatedly:
//   - Find element in slice? Write a loop.
//   - Sort slice? Use sort.Slice with less func.
//   - Copy a map? Write a loop.
//
// The slices and maps packages provide generic functions for common operations.
// They work with any slice/map type thanks to Go generics.
//
// PACKAGES
//
//   slices - Functions for slice operations
//   maps   - Functions for map operations
//   cmp    - Comparison utilities (Ordered constraint, Compare, Less)
//
// IMPORT
//
//   import "slices"
//   import "maps"
//   import "cmp"
//
// =============================================================================

// =============================================================================
// SECTION 2: SLICES - SEARCHING
// =============================================================================
//
// slices.Contains(s, v) bool
//   Does s contain v?
//
// slices.Index(s, v) int
//   Index of first v, or -1 if not found
//
// slices.BinarySearch(s, v) (int, bool)
//   Search in SORTED slice, returns index and found
//
// FUNC VARIANTS (custom comparison)
//
// slices.ContainsFunc(s, func(E) bool) bool
// slices.IndexFunc(s, func(E) bool) int
// slices.BinarySearchFunc(s, target, func(E, T) int) (int, bool)
//
// =============================================================================

func DemonstrateSearching() {
	nums := []int{1, 2, 3, 4, 5}
	words := []string{"apple", "banana", "cherry"}

	// Contains
	fmt.Printf("Contains 3: %v\n", slices.Contains(nums, 3))
	fmt.Printf("Contains 9: %v\n", slices.Contains(nums, 9))
	fmt.Printf("Contains 'banana': %v\n", slices.Contains(words, "banana"))

	// Index
	fmt.Printf("Index of 3: %d\n", slices.Index(nums, 3))
	fmt.Printf("Index of 9: %d\n", slices.Index(nums, 9)) // -1

	// BinarySearch (slice must be sorted!)
	sorted := []int{1, 3, 5, 7, 9}
	idx, found := slices.BinarySearch(sorted, 5)
	fmt.Printf("BinarySearch 5: idx=%d, found=%v\n", idx, found)

	idx2, found2 := slices.BinarySearch(sorted, 4)
	fmt.Printf("BinarySearch 4: idx=%d, found=%v\n", idx2, found2) // idx where it would be inserted

	// ContainsFunc - custom predicate
	hasEven := slices.ContainsFunc(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Has even: %v\n", hasEven)

	// IndexFunc - find by predicate
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{{"Alice", 30}, {"Bob", 25}, {"Carol", 35}}
	idx3 := slices.IndexFunc(people, func(p Person) bool {
		return p.Name == "Bob"
	})
	fmt.Printf("Bob is at index: %d\n", idx3)
}

// =============================================================================
// SECTION 3: SLICES - SORTING
// =============================================================================
//
// slices.Sort(s)
//   Sort in ascending order (requires cmp.Ordered elements)
//
// slices.SortFunc(s, cmp func(a, b E) int)
//   Sort with custom comparison (-1, 0, +1)
//
// slices.SortStableFunc(s, cmp func(a, b E) int)
//   Stable sort (preserves order of equal elements)
//
// slices.IsSorted(s) bool
//   Check if sorted
//
// slices.IsSortedFunc(s, cmp func(a, b E) int) bool
//   Check if sorted by custom comparison
//
// cmp.Compare(a, b) int
//   Returns -1, 0, or +1 for ordered types
//
// cmp.Less(a, b) bool
//   Returns a < b for ordered types
//
// =============================================================================

func DemonstrateSorting() {
	// Basic sort
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	slices.Sort(nums)
	fmt.Printf("Sorted: %v\n", nums)

	// Sort strings
	words := []string{"banana", "apple", "cherry"}
	slices.Sort(words)
	fmt.Printf("Sorted words: %v\n", words)

	// SortFunc - custom order
	nums2 := []int{3, 1, 4, 1, 5, 9, 2, 6}
	slices.SortFunc(nums2, func(a, b int) int {
		return cmp.Compare(b, a) // Descending
	})
	fmt.Printf("Descending: %v\n", nums2)

	// Sort structs
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{{"Carol", 35}, {"Alice", 30}, {"Bob", 25}}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.Age, b.Age) // By age
	})
	fmt.Printf("By age: %v\n", people)

	slices.SortFunc(people, func(a, b Person) int {
		return strings.Compare(a.Name, b.Name) // By name
	})
	fmt.Printf("By name: %v\n", people)

	// IsSorted
	sorted := []int{1, 2, 3, 4, 5}
	unsorted := []int{1, 3, 2, 4, 5}
	fmt.Printf("Is sorted: %v\n", slices.IsSorted(sorted))
	fmt.Printf("Is sorted: %v\n", slices.IsSorted(unsorted))

	// cmp.Compare
	fmt.Printf("Compare(1, 2): %d\n", cmp.Compare(1, 2))   // -1
	fmt.Printf("Compare(2, 2): %d\n", cmp.Compare(2, 2))   // 0
	fmt.Printf("Compare(3, 2): %d\n", cmp.Compare(3, 2))   // +1
}

// =============================================================================
// SECTION 4: SLICES - MIN, MAX, COMPARISON
// =============================================================================
//
// slices.Min(s) E
//   Minimum element (panics if empty)
//
// slices.Max(s) E
//   Maximum element (panics if empty)
//
// slices.MinFunc(s, cmp) E
// slices.MaxFunc(s, cmp) E
//   With custom comparison
//
// slices.Equal(s1, s2) bool
//   Are slices equal? (same length and elements)
//
// slices.EqualFunc(s1, s2, eq func(E1, E2) bool) bool
//   Custom equality
//
// slices.Compare(s1, s2) int
//   Lexicographic comparison (-1, 0, +1)
//
// slices.CompareFunc(s1, s2, cmp func(E1, E2) int) int
//   Custom comparison
//
// =============================================================================

func DemonstrateMinMaxCompare() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// Min and Max
	fmt.Printf("Min: %d\n", slices.Min(nums))
	fmt.Printf("Max: %d\n", slices.Max(nums))

	// MinFunc/MaxFunc for structs
	type Item struct {
		Name  string
		Price float64
	}
	items := []Item{{"A", 10.0}, {"B", 5.0}, {"C", 15.0}}

	cheapest := slices.MinFunc(items, func(a, b Item) int {
		return cmp.Compare(a.Price, b.Price)
	})
	fmt.Printf("Cheapest: %v\n", cheapest)

	// Equal
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}
	fmt.Printf("a == b: %v\n", slices.Equal(a, b))
	fmt.Printf("a == c: %v\n", slices.Equal(a, c))

	// Compare (lexicographic)
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	fmt.Printf("Compare [1,2,3] vs [1,2,4]: %d\n", slices.Compare(s1, s2)) // -1
}

// =============================================================================
// SECTION 5: SLICES - MODIFICATION
// =============================================================================
//
// slices.Clone(s) []E
//   Shallow copy
//
// slices.Compact(s) []E
//   Remove consecutive duplicates (in-place, returns new length)
//
// slices.CompactFunc(s, eq func(E, E) bool) []E
//   Custom equality
//
// slices.Grow(s, n) []E
//   Grow capacity by at least n
//
// slices.Clip(s) []E
//   Remove unused capacity
//
// slices.Reverse(s)
//   Reverse in place
//
// slices.Replace(s, i, j, v...) []E
//   Replace s[i:j] with v...
//
// slices.Insert(s, i, v...) []E
//   Insert v... at index i
//
// slices.Delete(s, i, j) []E
//   Delete s[i:j]
//
// slices.DeleteFunc(s, del func(E) bool) []E
//   Delete elements matching predicate
//
// =============================================================================

func DemonstrateModification() {
	// Clone
	original := []int{1, 2, 3}
	cloned := slices.Clone(original)
	cloned[0] = 999
	fmt.Printf("Original: %v, Clone: %v\n", original, cloned)

	// Compact - remove consecutive duplicates
	dups := []int{1, 1, 2, 2, 2, 3, 3}
	compacted := slices.Compact(dups)
	fmt.Printf("Compacted: %v\n", compacted)

	// Reverse
	nums := []int{1, 2, 3, 4, 5}
	slices.Reverse(nums)
	fmt.Printf("Reversed: %v\n", nums)

	// Insert
	s := []int{1, 2, 5}
	s = slices.Insert(s, 2, 3, 4)
	fmt.Printf("After insert: %v\n", s)

	// Delete
	s = slices.Delete(s, 1, 3) // Delete index 1 and 2
	fmt.Printf("After delete: %v\n", s)

	// DeleteFunc
	nums2 := []int{1, 2, 3, 4, 5, 6}
	odds := slices.DeleteFunc(nums2, func(n int) bool {
		return n%2 == 0 // Delete evens
	})
	fmt.Printf("After delete evens: %v\n", odds)

	// Replace
	s2 := []int{1, 2, 3, 4, 5}
	s2 = slices.Replace(s2, 1, 3, 20, 30, 40)
	fmt.Printf("After replace: %v\n", s2)

	// Grow and Clip
	small := make([]int, 3, 3)
	grown := slices.Grow(small, 100)
	fmt.Printf("Grew capacity: %d → %d\n", cap(small), cap(grown))

	big := make([]int, 3, 1000)
	clipped := slices.Clip(big)
	fmt.Printf("Clipped capacity: %d → %d\n", cap(big), cap(clipped))
}

// =============================================================================
// SECTION 6: SLICES - CONCATENATION
// =============================================================================
//
// slices.Concat(slices ...[]E) []E
//   Concatenate multiple slices into new slice (Go 1.22+)
//
// =============================================================================

func DemonstrateConcatenation() {
	a := []int{1, 2}
	b := []int{3, 4}
	c := []int{5, 6}

	// Concat (Go 1.22+)
	combined := slices.Concat(a, b, c)
	fmt.Printf("Concat: %v\n", combined)

	// Empty slices are fine
	empty := []int{}
	result := slices.Concat(a, empty, c)
	fmt.Printf("With empty: %v\n", result)
}

// =============================================================================
// SECTION 7: MAPS PACKAGE
// =============================================================================
//
// maps.Clone(m) M
//   Shallow copy of map
//
// maps.Copy(dst, src)
//   Copy src entries to dst (overwrites existing keys)
//
// maps.DeleteFunc(m, del func(K, V) bool)
//   Delete entries matching predicate
//
// maps.Equal(m1, m2) bool
//   Are maps equal? (same keys and values)
//
// maps.EqualFunc(m1, m2, eq func(V1, V2) bool) bool
//   Custom value equality
//
// maps.Keys(m) []K (Go 1.23+, or use iter)
//   Return slice of keys
//
// maps.Values(m) []V (Go 1.23+, or use iter)
//   Return slice of values
//
// =============================================================================

func DemonstrateMaps() {
	// Clone
	original := map[string]int{"a": 1, "b": 2}
	cloned := maps.Clone(original)
	cloned["a"] = 999
	fmt.Printf("Original: %v, Clone: %v\n", original, cloned)

	// Copy
	dst := map[string]int{"a": 0, "c": 3}
	src := map[string]int{"a": 1, "b": 2}
	maps.Copy(dst, src)
	fmt.Printf("After copy: %v\n", dst)

	// DeleteFunc
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	maps.DeleteFunc(m, func(k string, v int) bool {
		return v%2 == 0 // Delete even values
	})
	fmt.Printf("After delete evens: %v\n", m)

	// Equal
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 2}
	m3 := map[string]int{"a": 1, "b": 3}
	fmt.Printf("m1 == m2: %v\n", maps.Equal(m1, m2))
	fmt.Printf("m1 == m3: %v\n", maps.Equal(m1, m3))

	// Keys and Values iteration (works in Go 1.21+ with range-over-func)
	// For Go 1.21-1.22, collect manually:
	m4 := map[string]int{"x": 1, "y": 2, "z": 3}
	keys := make([]string, 0, len(m4))
	values := make([]int, 0, len(m4))
	for k, v := range m4 {
		keys = append(keys, k)
		values = append(values, v)
	}
	fmt.Printf("Keys: %v, Values: %v\n", keys, values)
}

// =============================================================================
// SECTION 8: CMP PACKAGE
// =============================================================================
//
// cmp.Ordered
//   Constraint for types that support < > <= >=
//   (integers, floats, strings)
//
// cmp.Compare(x, y) int
//   -1 if x < y, 0 if x == y, +1 if x > y
//
// cmp.Less(x, y) bool
//   x < y
//
// cmp.Or(vals ...T) T
//   Returns first non-zero value (like || for values)
//
// =============================================================================

func DemonstrateCmp() {
	// Compare
	fmt.Printf("Compare(1, 2): %d\n", cmp.Compare(1, 2))     // -1
	fmt.Printf("Compare(2, 2): %d\n", cmp.Compare(2, 2))     // 0
	fmt.Printf("Compare(3, 2): %d\n", cmp.Compare(3, 2))     // +1
	fmt.Printf("Compare(\"a\", \"b\"): %d\n", cmp.Compare("a", "b")) // -1

	// Less
	fmt.Printf("Less(1, 2): %v\n", cmp.Less(1, 2)) // true
	fmt.Printf("Less(2, 1): %v\n", cmp.Less(2, 1)) // false

	// Or - first non-zero value (Go 1.22+)
	fmt.Printf("Or(\"\", \"\", \"default\"): %q\n", cmp.Or("", "", "default"))
	fmt.Printf("Or(0, 0, 42): %d\n", cmp.Or(0, 0, 42))
	fmt.Printf("Or(\"first\", \"second\"): %q\n", cmp.Or("first", "second"))
}

// =============================================================================
// SECTION 9: COMMON PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Find or default
	findOrDefault := func(s []int, pred func(int) bool, def int) int {
		idx := slices.IndexFunc(s, pred)
		if idx == -1 {
			return def
		}
		return s[idx]
	}
	nums := []int{1, 3, 5, 7}
	even := findOrDefault(nums, func(n int) bool { return n%2 == 0 }, -1)
	fmt.Printf("Pattern 1 - Find even or -1: %d\n", even)

	// Pattern 2: Filter slice
	filter := func(s []int, keep func(int) bool) []int {
		result := slices.Clone(s)
		return slices.DeleteFunc(result, func(n int) bool {
			return !keep(n)
		})
	}
	evens := filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Pattern 2 - Filter evens: %v\n", evens)

	// Pattern 3: Map/transform (not in slices pkg, but common)
	transform := func(s []int, f func(int) int) []int {
		result := make([]int, len(s))
		for i, v := range s {
			result[i] = f(v)
		}
		return result
	}
	doubled := transform([]int{1, 2, 3}, func(n int) int { return n * 2 })
	fmt.Printf("Pattern 3 - Double: %v\n", doubled)

	// Pattern 4: Unique (sort + compact)
	unique := func(s []int) []int {
		result := slices.Clone(s)
		slices.Sort(result)
		return slices.Compact(result)
	}
	uniq := unique([]int{3, 1, 2, 1, 3, 2})
	fmt.Printf("Pattern 4 - Unique: %v\n", uniq)

	// Pattern 5: Top N
	topN := func(s []int, n int) []int {
		result := slices.Clone(s)
		slices.SortFunc(result, func(a, b int) int {
			return cmp.Compare(b, a) // Descending
		})
		if n > len(result) {
			n = len(result)
		}
		return result[:n]
	}
	top3 := topN([]int{5, 2, 8, 1, 9, 3}, 3)
	fmt.Printf("Pattern 5 - Top 3: %v\n", top3)

	// Pattern 6: Group by (with maps)
	type Person struct {
		Name string
		City string
	}
	people := []Person{
		{"Alice", "NYC"}, {"Bob", "LA"}, {"Carol", "NYC"}, {"Dave", "LA"},
	}
	grouped := make(map[string][]Person)
	for _, p := range people {
		grouped[p.City] = append(grouped[p.City], p)
	}
	fmt.Printf("Pattern 6 - Grouped: %v\n", grouped)
}

// =============================================================================
// SECTION 10: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: BinarySearch on unsorted slice
	_ = `
	// WRONG
	unsorted := []int{3, 1, 4}
	idx, found := slices.BinarySearch(unsorted, 1) // Wrong result!

	// RIGHT
	sorted := []int{1, 3, 4}
	idx, found := slices.BinarySearch(sorted, 1)
	`
	fmt.Println("Mistake 1: BinarySearch requires sorted slice")

	// Mistake 2: Modifying slice during iteration
	_ = `
	// WRONG
	for i, v := range s {
		if v == target {
			s = slices.Delete(s, i, i+1) // Invalidates iteration!
		}
	}

	// RIGHT
	s = slices.DeleteFunc(s, func(v int) bool { return v == target })
	`
	fmt.Println("Mistake 2: Use DeleteFunc instead of Delete in loop")

	// Mistake 3: Forgetting Clone modifies original
	_ = `
	// Clone is shallow!
	type Item struct { Data *string }
	items := []Item{{Data: &str}}
	cloned := slices.Clone(items)
	*cloned[0].Data = "modified" // Also modifies original!
	`
	fmt.Println("Mistake 3: Clone is shallow - pointers are shared")

	// Mistake 4: Min/Max on empty slice panics
	_ = `
	// WRONG
	empty := []int{}
	min := slices.Min(empty) // PANIC!

	// RIGHT
	if len(s) > 0 {
		min := slices.Min(s)
	}
	`
	fmt.Println("Mistake 4: Min/Max panic on empty slice")

	// Mistake 5: Expecting Compact to remove all duplicates
	_ = `
	// Compact only removes CONSECUTIVE duplicates
	s := []int{1, 2, 1, 2}
	slices.Compact(s) // Still [1, 2, 1, 2]!

	// For all duplicates: sort first
	slices.Sort(s)    // [1, 1, 2, 2]
	slices.Compact(s) // [1, 2]
	`
	fmt.Println("Mistake 5: Compact only removes consecutive duplicates")
}

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateSearching()
	DemonstrateSorting()
	DemonstrateMinMaxCompare()
	DemonstrateModification()
	DemonstrateConcatenation()
	DemonstrateMaps()
	DemonstrateCmp()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
}
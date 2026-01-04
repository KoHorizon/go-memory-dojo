// Package slices_practice is your daily practice space for slices/maps/cmp packages.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against slices/slices.go
// 3. Note what you missed - focus on those tomorrow
package slices_practice

// =============================================================================
// EXERCISE 1: IMPORTS
// =============================================================================

func PracticeImports() {
	// TODO: What three packages for generic collections?
	// import "???"  - slice operations
	// import "???"  - map operations
	// import "???"  - comparison utilities

	// TODO: What Go version introduced slices package?
	// Answer: ???
}

// =============================================================================
// EXERCISE 2: SEARCHING
// =============================================================================

func PracticeSearching() {
	// TODO: Check if slice contains value
	// slices.???(s, value)

	// TODO: Find index of value (-1 if not found)
	// slices.???(s, value)

	// TODO: Binary search in sorted slice
	// idx, found := slices.???(s, value)

	// TODO: Contains with custom predicate
	// slices.???(s, func(E) bool)

	// TODO: Index with custom predicate
	// slices.???(s, func(E) bool)

	// TODO: What's required for BinarySearch?
	// Answer: ???
}

// =============================================================================
// EXERCISE 3: SORTING
// =============================================================================

func PracticeSorting() {
	// TODO: Sort slice in ascending order
	// slices.???(s)

	// TODO: Sort with custom comparison
	// slices.???(s, func(a, b E) int)

	// TODO: Stable sort (preserves equal order)
	// slices.???(s, func(a, b E) int)

	// TODO: Check if sorted
	// slices.???(s)

	// TODO: What does comparison func return?
	// Answer: ???, ???, or ???

	// TODO: Sort descending
	// slices.SortFunc(s, func(a, b int) int {
	//     return cmp.Compare(???, ???)
	// })
}

// =============================================================================
// EXERCISE 4: CMP PACKAGE
// =============================================================================

func PracticeCmp() {
	// TODO: Compare two ordered values
	// result := cmp.???(a, b)  // -1, 0, or +1

	// TODO: Check if a < b
	// result := cmp.???(a, b)  // bool

	// TODO: First non-zero value
	// result := cmp.???(val1, val2, val3)

	// TODO: What types satisfy cmp.Ordered?
	// Answer: ???
}

// =============================================================================
// EXERCISE 5: MIN, MAX, COMPARISON
// =============================================================================

func PracticeMinMaxCompare() {
	// TODO: Get minimum element
	// min := slices.???(s)

	// TODO: Get maximum element
	// max := slices.???(s)

	// TODO: Min/Max with custom comparison
	// min := slices.???(s, func(a, b E) int)
	// max := slices.???(s, func(a, b E) int)

	// TODO: Check if two slices are equal
	// equal := slices.???(s1, s2)

	// TODO: Lexicographic comparison
	// cmp := slices.???(s1, s2)  // -1, 0, or +1

	// TODO: What happens if Min/Max on empty slice?
	// Answer: ???
}

// =============================================================================
// EXERCISE 6: MODIFICATION
// =============================================================================

func PracticeModification() {
	// TODO: Shallow copy of slice
	// copy := slices.???(s)

	// TODO: Remove consecutive duplicates
	// s = slices.???(s)

	// TODO: Reverse slice in place
	// slices.???(s)

	// TODO: Insert values at index i
	// s = slices.???(s, i, values...)

	// TODO: Delete range s[i:j]
	// s = slices.???(s, i, j)

	// TODO: Delete by predicate
	// s = slices.???(s, func(E) bool)

	// TODO: Replace s[i:j] with values
	// s = slices.???(s, i, j, values...)

	// TODO: Grow capacity by at least n
	// s = slices.???(s, n)

	// TODO: Remove unused capacity
	// s = slices.???(s)
}

// =============================================================================
// EXERCISE 7: CONCATENATION
// =============================================================================

func PracticeConcatenation() {
	// TODO: Concatenate multiple slices (Go 1.22+)
	// result := slices.???(s1, s2, s3)
}

// =============================================================================
// EXERCISE 8: MAPS PACKAGE
// =============================================================================

func PracticeMaps() {
	// TODO: Shallow copy of map
	// copy := maps.???(m)

	// TODO: Copy src entries to dst
	// maps.???(dst, src)

	// TODO: Delete entries by predicate
	// maps.???(m, func(K, V) bool)

	// TODO: Check if maps are equal
	// equal := maps.???(m1, m2)

	// TODO: How do you get all keys? (manual way)
	// keys := make([]K, 0, len(m))
	// for k := range m {
	//     keys = ???(keys, k)
	// }
}

// =============================================================================
// EXERCISE 9: COMMON PATTERNS
// =============================================================================

func PracticePatterns() {
	// TODO: Unique elements (sort + compact)
	// result := slices.Clone(s)
	// slices.???(result)
	// result = slices.???(result)

	// TODO: Filter slice (keep matching)
	// Use DeleteFunc with inverted predicate

	// TODO: Top N elements
	// 1. Clone
	// 2. SortFunc descending
	// 3. Slice [:n]

	// TODO: Check if sorted descending
	// slices.IsSortedFunc(s, func(a, b int) int {
	//     return cmp.Compare(???, ???)
	// })
}

// =============================================================================
// EXERCISE 10: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// unsorted := []int{3, 1, 4}
	// idx, _ := slices.BinarySearch(unsorted, 1)
	// Answer: ???

	// Mistake 2: What's wrong?
	// for i, v := range s {
	//     if v == target {
	//         s = slices.Delete(s, i, i+1)
	//     }
	// }
	// Answer: ???

	// Mistake 3: What's wrong?
	// original := []Item{{ptr: &data}}
	// cloned := slices.Clone(original)
	// *cloned[0].ptr = "new"  // Affects original?
	// Answer: ???

	// Mistake 4: What's wrong?
	// empty := []int{}
	// min := slices.Min(empty)
	// Answer: ???

	// Mistake 5: What's wrong?
	// s := []int{1, 2, 1, 2}
	// slices.Compact(s)  // Expected [1, 2]?
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// 1. Contains element?
	_ = "slices.???(s, v)"

	// 2. Index of element?
	_ = "slices.???(s, v)"

	// 3. Binary search?
	_ = "slices.???(s, v)"

	// 4. Sort ascending?
	_ = "slices.???(s)"

	// 5. Sort with func?
	_ = "slices.???(s, cmp)"

	// 6. Is sorted?
	_ = "slices.???(s)"

	// 7. Minimum?
	_ = "slices.???(s)"

	// 8. Maximum?
	_ = "slices.???(s)"

	// 9. Slices equal?
	_ = "slices.???(s1, s2)"

	// 10. Clone slice?
	_ = "slices.???(s)"

	// 11. Remove consecutive dups?
	_ = "slices.???(s)"

	// 12. Reverse?
	_ = "slices.???(s)"

	// 13. Insert?
	_ = "slices.???(s, i, vals...)"

	// 14. Delete range?
	_ = "slices.???(s, i, j)"

	// 15. Clone map?
	_ = "maps.???(m)"

	// 16. Compare values?
	_ = "cmp.???(a, b)"
}

// =============================================================================
// MINI PROJECT: LEADERBOARD
// =============================================================================

func MiniProject() {
	// Build a leaderboard system that:
	//
	// 1. Stores players with scores
	// 2. Can add/update player scores
	// 3. Returns top N players
	// 4. Checks if player exists
	// 5. Removes players below threshold
	//
	// Scaffold:
	//
	// type Player struct {
	//     Name  string
	//     Score int
	// }
	//
	// type Leaderboard struct {
	//     players []Player
	// }
	//
	// func (lb *Leaderboard) AddOrUpdate(name string, score int) {
	//     idx := slices.IndexFunc(lb.players, func(p Player) bool {
	//         return p.Name == name
	//     })
	//     if idx != -1 {
	//         lb.players[idx].Score = score
	//     } else {
	//         lb.players = append(lb.players, Player{name, score})
	//     }
	// }
	//
	// func (lb *Leaderboard) TopN(n int) []Player {
	//     sorted := slices.Clone(lb.players)
	//     slices.SortFunc(sorted, func(a, b Player) int {
	//         return cmp.Compare(b.Score, a.Score)  // Descending
	//     })
	//     if n > len(sorted) {
	//         n = len(sorted)
	//     }
	//     return sorted[:n]
	// }
	//
	// func (lb *Leaderboard) HasPlayer(name string) bool {
	//     return slices.ContainsFunc(lb.players, func(p Player) bool {
	//         return p.Name == name
	//     })
	// }
	//
	// func (lb *Leaderboard) RemoveBelow(threshold int) {
	//     lb.players = slices.DeleteFunc(lb.players, func(p Player) bool {
	//         return p.Score < threshold
	//     })
	// }
	//
	// func (lb *Leaderboard) HighestScore() (Player, bool) {
	//     if len(lb.players) == 0 {
	//         return Player{}, false
	//     }
	//     return slices.MaxFunc(lb.players, func(a, b Player) int {
	//         return cmp.Compare(a.Score, b.Score)
	//     }), true
	// }
}

// AllPractice is not meant to be called - this file is for reading.
func AllPractice() {
	PracticeImports()
	PracticeSearching()
	PracticeSorting()
	PracticeCmp()
	PracticeMinMaxCompare()
	PracticeModification()
	PracticeConcatenation()
	PracticeMaps()
	PracticePatterns()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
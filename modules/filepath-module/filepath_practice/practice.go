// Package filepath_practice is your daily practice space for path/filepath.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against filepath/filepath.go
// 3. Note what you missed - focus on those tomorrow
package filepath_practice

// =============================================================================
// EXERCISE 1: PATH VS FILEPATH
// =============================================================================

func PracticePathVsFilepath() {
	// TODO: Which package for file system paths?
	// import "path/???"

	// TODO: Which package for URL paths?
	// import "???"

	// TODO: What's the rule?
	// Answer: If it touches the ???, use filepath

	// TODO: OS path separator constant?
	// filepath.???

	// TODO: PATH list separator constant (: or ;)?
	// filepath.???
}

// =============================================================================
// EXERCISE 2: JOINING AND SPLITTING
// =============================================================================

func PracticeJoinSplit() {
	// TODO: Join path elements
	// p := filepath.???("home", "user", "file.txt")

	// TODO: Split into directory and filename
	// dir, file := filepath.???(path)

	// TODO: Get just the directory (no trailing sep)
	// dir := filepath.???(path)

	// TODO: Get just the filename
	// name := filepath.???(path)

	// TODO: Get file extension (with dot)
	// ext := filepath.???(path)

	// TODO: What does Split return for "/a/b/c"?
	// dir = ???, file = ???

	// TODO: What does Dir return for "/a/b/c"?
	// dir = ???

	// TODO: What does Ext return for "file.tar.gz"?
	// ext = ???
}

// =============================================================================
// EXERCISE 3: CLEANING PATHS
// =============================================================================

func PracticeClean() {
	// TODO: Clean/normalize a path
	// clean := filepath.???(path)

	// TODO: What does Clean do? (list 4 things)
	// 1. Replace multiple ??? with single
	// 2. Eliminate ??? (current dir)
	// 3. Eliminate ??? where possible
	// 4. Remove trailing ???

	// TODO: Convert forward slashes to OS separator
	// osPath := filepath.???(path)

	// TODO: Convert OS separator to forward slashes
	// slashPath := filepath.???(path)
}

// =============================================================================
// EXERCISE 4: ABSOLUTE AND RELATIVE
// =============================================================================

func PracticeAbsRel() {
	// TODO: Check if path is absolute
	// isAbs := filepath.???(path)

	// TODO: Convert to absolute path
	// abs, err := filepath.???(path)

	// TODO: Get relative path from base to target
	// rel, err := filepath.???(base, target)

	// TODO: What does Abs do if path is relative?
	// Answer: Prepends ???
}

// =============================================================================
// EXERCISE 5: MATCHING AND GLOBBING
// =============================================================================

func PracticeMatching() {
	// TODO: Match against shell pattern
	// match, err := filepath.???(pattern, name)

	// TODO: Find all matching files
	// files, err := filepath.???(pattern)

	// TODO: Pattern syntax
	// ??? - match any sequence (not separator)
	// ??? - match single char (not separator)
	// ??? - match character class

	// TODO: Does Glob support **?
	// Answer: ???
}

// =============================================================================
// EXERCISE 6: WALKING DIRECTORIES
// =============================================================================

func PracticeWalking() {
	// TODO: Walk directory tree (efficient)
	// err := filepath.???(root, func(path string, d fs.DirEntry, err error) error {
	//     ...
	// })

	// TODO: WalkDirFunc signature
	// func(??? string, ??? fs.DirEntry, ??? error) error

	// TODO: Skip current directory
	// return filepath.???

	// TODO: Stop walking entirely (Go 1.20+)
	// return filepath.???

	// TODO: Legacy walk function (less efficient)?
	// filepath.???

	// TODO: Why is WalkDir more efficient than Walk?
	// Answer: Doesn't call ??? for every file
}

// =============================================================================
// EXERCISE 7: VOLUME AND ENVIRONMENT
// =============================================================================

func PracticeVolume() {
	// TODO: Get volume name (Windows: "C:")
	// vol := filepath.???(path)

	// TODO: Split PATH-style list
	// dirs := filepath.???(pathEnv)

	// TODO: Resolve symlinks
	// resolved, err := filepath.???(path)
}

// =============================================================================
// EXERCISE 8: COMMON PATTERNS
// =============================================================================

func PracticePatterns() {
	// TODO: Get filename without extension
	// ext := filepath.Ext(path)
	// name := path[:len(path)-len(???)]

	// TODO: Change file extension
	// ext := filepath.???(path)
	// newPath := path[:len(path)-len(ext)] + ".newext"

	// TODO: Ensure parent directory exists
	// dir := filepath.???(path)
	// os.???(dir, 0755)

	// TODO: Safe path join (prevent traversal)
	// full := filepath.Join(base, filepath.???(???+userInput))
	// Check that full still starts with ???

	// TODO: Find all files with extension
	// filepath.???(root, func(path string, d fs.DirEntry, err error) error {
	//     if filepath.???(path) == ".go" {
	//         files = append(files, path)
	//     }
	//     return nil
	// })

	// TODO: Expand ~ to home directory
	// home, _ := os.???()
	// path = filepath.???(home, path[1:])
}

// =============================================================================
// EXERCISE 9: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// path := "home" + "/" + "user"
	// Answer: ???

	// Mistake 2: What's wrong?
	// import "path"
	// configPath := path.Join("config", "app.json")
	// os.ReadFile(configPath)
	// Answer: ???

	// Mistake 3: What's wrong?
	// filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
	//     name := d.Name()  // Using d without checking err
	// })
	// Answer: ???

	// Mistake 4: What's wrong?
	// files, _ := filepath.Glob("**/*.go")
	// Answer: ???

	// Mistake 5: What's wrong?
	// path := filepath.Join(baseDir, userInput)
	// Answer: ???

	// Mistake 6: What's the difference?
	// dir, _ := filepath.Split("/a/b/c")  // dir = ???
	// dir := filepath.Dir("/a/b/c")       // dir = ???
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// 1. Join path elements?
	_ = "filepath.???(elems...)"

	// 2. Split into dir and file?
	_ = "filepath.???(path)"

	// 3. Get directory only?
	_ = "filepath.???(path)"

	// 4. Get filename only?
	_ = "filepath.???(path)"

	// 5. Get extension?
	_ = "filepath.???(path)"

	// 6. Clean/normalize path?
	_ = "filepath.???(path)"

	// 7. Is absolute?
	_ = "filepath.???(path)"

	// 8. Make absolute?
	_ = "filepath.???(path)"

	// 9. Make relative?
	_ = "filepath.???(base, target)"

	// 10. Match pattern?
	_ = "filepath.???(pattern, name)"

	// 11. Find matching files?
	_ = "filepath.???(pattern)"

	// 12. Walk directory?
	_ = "filepath.???(root, fn)"

	// 13. Skip directory in walk?
	_ = "return filepath.???"

	// 14. Forward to OS separator?
	_ = "filepath.???(path)"

	// 15. OS to forward separator?
	_ = "filepath.???(path)"

	// 16. Split PATH list?
	_ = "filepath.???(pathEnv)"
}

// =============================================================================
// MINI PROJECT: FILE ORGANIZER
// =============================================================================

func MiniProject() {
	// Build a file organizer that:
	//
	// 1. Walks a directory tree
	// 2. Groups files by extension
	// 3. Calculates total size per extension
	// 4. Lists largest files
	// 5. Optionally moves files to extension-named folders
	//
	// Scaffold:
	//
	// type FileInfo struct {
	//     Path string
	//     Size int64
	//     Ext  string
	// }
	//
	// type Stats struct {
	//     ByExtension map[string][]FileInfo
	//     TotalSize   map[string]int64
	// }
	//
	// func Analyze(root string) (*Stats, error) {
	//     stats := &Stats{
	//         ByExtension: make(map[string][]FileInfo),
	//         TotalSize:   make(map[string]int64),
	//     }
	//
	//     err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
	//         if err != nil {
	//             return nil
	//         }
	//         if d.IsDir() {
	//             return nil
	//         }
	//
	//         info, err := d.Info()
	//         if err != nil {
	//             return nil
	//         }
	//
	//         ext := filepath.Ext(path)
	//         if ext == "" {
	//             ext = "(no extension)"
	//         }
	//
	//         fi := FileInfo{
	//             Path: path,
	//             Size: info.Size(),
	//             Ext:  ext,
	//         }
	//
	//         stats.ByExtension[ext] = append(stats.ByExtension[ext], fi)
	//         stats.TotalSize[ext] += fi.Size
	//
	//         return nil
	//     })
	//
	//     return stats, err
	// }
	//
	// func LargestFiles(stats *Stats, n int) []FileInfo {
	//     var all []FileInfo
	//     for _, files := range stats.ByExtension {
	//         all = append(all, files...)
	//     }
	//     slices.SortFunc(all, func(a, b FileInfo) int {
	//         return cmp.Compare(b.Size, a.Size)  // Descending
	//     })
	//     if n > len(all) {
	//         n = len(all)
	//     }
	//     return all[:n]
	// }
	//
	// func Organize(root string) error {
	//     return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
	//         if err != nil || d.IsDir() {
	//             return nil
	//         }
	//
	//         ext := filepath.Ext(path)
	//         if ext == "" {
	//             return nil
	//         }
	//
	//         // Create extension folder
	//         extDir := filepath.Join(root, ext[1:])  // Remove dot
	//         os.MkdirAll(extDir, 0755)
	//
	//         // Move file
	//         newPath := filepath.Join(extDir, filepath.Base(path))
	//         return os.Rename(path, newPath)
	//     })
	// }
}

// AllPractice is not meant to be called - this file is for reading.
func AllPractice() {
	PracticePathVsFilepath()
	PracticeJoinSplit()
	PracticeClean()
	PracticeAbsRel()
	PracticeMatching()
	PracticeWalking()
	PracticeVolume()
	PracticePatterns()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
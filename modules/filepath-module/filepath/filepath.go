// Package filepathmod provides comprehensive documentation and working examples
// for Go's path/filepath package - OS-native file path manipulation.
package filepathmod

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF PATH/FILEPATH
// =============================================================================
//
// TWO PATH PACKAGES
//
//   path          - For slash-separated paths (URLs, always forward slash)
//   path/filepath - For OS file paths (uses OS separator)
//
// USE path/filepath FOR FILE SYSTEM OPERATIONS
//
// On Windows: C:\Users\name\file.txt (backslash)
// On Unix:    /home/name/file.txt (forward slash)
//
// filepath handles this automatically. Never hardcode separators!
//
// KEY CONSTANTS
//
//   filepath.Separator     - OS path separator ('/' or '\')
//   filepath.ListSeparator - PATH list separator (':' or ';')
//
// =============================================================================

func DemonstrateBasics() {
	// OS-specific separator
	fmt.Printf("Separator: %q\n", string(filepath.Separator))
	fmt.Printf("ListSeparator: %q\n", string(filepath.ListSeparator))

	// Join uses correct separator
	path := filepath.Join("home", "user", "documents", "file.txt")
	fmt.Printf("Joined: %s\n", path)

	// Works the same on all platforms
	path2 := filepath.Join("a", "b", "c")
	fmt.Printf("Joined a/b/c: %s\n", path2)
}

// =============================================================================
// SECTION 2: JOINING AND SPLITTING PATHS
// =============================================================================
//
// filepath.Join(elem ...string) string
//   Join path elements with OS separator
//   Cleans the result (removes redundant separators, dots)
//
// filepath.Split(path string) (dir, file string)
//   Split into directory and file
//   dir includes trailing separator
//
// filepath.Dir(path string) string
//   Return directory (without trailing separator)
//
// filepath.Base(path string) string
//   Return last element (file name)
//
// filepath.Ext(path string) string
//   Return extension (including dot)
//
// =============================================================================

func DemonstrateJoinSplit() {
	// Join - combines path elements
	p1 := filepath.Join("home", "user", "file.txt")
	fmt.Printf("Join: %s\n", p1)

	// Join cleans as it goes
	p2 := filepath.Join("home", "", "user", ".", "file.txt")
	fmt.Printf("Join (with cleanup): %s\n", p2)

	// Join with absolute path resets
	p3 := filepath.Join("relative", "/absolute", "file.txt")
	fmt.Printf("Join with abs: %s\n", p3)

	// Split - separate dir and file
	dir, file := filepath.Split("/home/user/document.txt")
	fmt.Printf("Split: dir=%q, file=%q\n", dir, file)

	// Dir - just directory (no trailing separator)
	d := filepath.Dir("/home/user/document.txt")
	fmt.Printf("Dir: %s\n", d)

	// Base - just filename
	b := filepath.Base("/home/user/document.txt")
	fmt.Printf("Base: %s\n", b)

	// Ext - extension with dot
	e := filepath.Ext("/home/user/document.txt")
	fmt.Printf("Ext: %s\n", e)

	// No extension
	e2 := filepath.Ext("/home/user/Makefile")
	fmt.Printf("Ext (none): %q\n", e2)

	// Multiple dots
	e3 := filepath.Ext("archive.tar.gz")
	fmt.Printf("Ext (tar.gz): %s\n", e3) // Just .gz
}

// =============================================================================
// SECTION 3: CLEANING AND NORMALIZING
// =============================================================================
//
// filepath.Clean(path string) string
//   Lexically clean path:
//   - Replace multiple separators with single
//   - Eliminate . (current dir)
//   - Eliminate .. where possible
//   - Remove trailing separator
//
// filepath.FromSlash(path string) string
//   Convert forward slashes to OS separator
//
// filepath.ToSlash(path string) string
//   Convert OS separator to forward slashes
//
// =============================================================================

func DemonstrateClean() {
	// Clean normalizes paths
	paths := []string{
		"a//b//c",
		"a/./b/./c",
		"a/b/../c",
		"/a/b/../c/./d",
		"./a/b",
		"a/b/",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, filepath.Clean(p))
	}

	// FromSlash - convert to OS separator
	unixPath := "home/user/file.txt"
	osPath := filepath.FromSlash(unixPath)
	fmt.Printf("FromSlash: %s\n", osPath)

	// ToSlash - convert to forward slashes
	back := filepath.ToSlash(osPath)
	fmt.Printf("ToSlash: %s\n", back)
}

// =============================================================================
// SECTION 4: ABSOLUTE AND RELATIVE PATHS
// =============================================================================
//
// filepath.Abs(path string) (string, error)
//   Return absolute path
//   Prepends current working directory if relative
//
// filepath.Rel(basepath, targpath string) (string, error)
//   Return relative path from base to target
//   Error if can't be made relative
//
// filepath.IsAbs(path string) bool
//   Is path absolute?
//
// =============================================================================

func DemonstrateAbsRel() {
	// IsAbs
	fmt.Printf("IsAbs(/home/user): %v\n", filepath.IsAbs("/home/user"))
	fmt.Printf("IsAbs(relative/path): %v\n", filepath.IsAbs("relative/path"))

	// Abs - make absolute
	abs, _ := filepath.Abs("relative/path")
	fmt.Printf("Abs: %s\n", abs)

	abs2, _ := filepath.Abs(".")
	fmt.Printf("Abs(.): %s\n", abs2)

	// Rel - make relative
	rel, _ := filepath.Rel("/home/user", "/home/user/documents/file.txt")
	fmt.Printf("Rel: %s\n", rel)

	rel2, _ := filepath.Rel("/home/user/documents", "/home/user/images/photo.jpg")
	fmt.Printf("Rel (sibling): %s\n", rel2)

	// Can't make relative between different roots (Windows)
	// rel3, err := filepath.Rel("C:\\Users", "D:\\Data")
	// Error: can't make relative
}

// =============================================================================
// SECTION 5: MATCHING AND GLOBBING
// =============================================================================
//
// filepath.Match(pattern, name string) (bool, error)
//   Match name against shell pattern
//   Patterns: * ? [...] (no **)
//
// filepath.Glob(pattern string) ([]string, error)
//   Return all matching file paths
//   Supports same patterns as Match
//
// PATTERN SYNTAX
//
//   *      - Match any sequence (not separator)
//   ?      - Match single char (not separator)
//   [abc]  - Match one of these chars
//   [a-z]  - Match range
//   [^abc] - Match any except these
//
// NOTE: ** (recursive) is NOT supported. Use WalkDir instead.
//
// =============================================================================

func DemonstrateMatching() {
	// Match patterns
	patterns := []struct {
		pattern string
		name    string
	}{
		{"*.txt", "file.txt"},
		{"*.txt", "file.go"},
		{"file.*", "file.txt"},
		{"file.???", "file.txt"},
		{"file.??", "file.txt"},
		{"[fF]ile.txt", "File.txt"},
		{"[!f]ile.txt", "bile.txt"},
	}

	for _, p := range patterns {
		match, _ := filepath.Match(p.pattern, p.name)
		fmt.Printf("Match(%q, %q): %v\n", p.pattern, p.name, match)
	}

	// Glob - find matching files
	// Note: This requires actual files to exist
	goFiles, _ := filepath.Glob("*.go")
	fmt.Printf("Glob(*.go): %v\n", goFiles)

	// Glob with directory
	allTxt, _ := filepath.Glob(filepath.Join("testdata", "*.txt"))
	fmt.Printf("Glob(testdata/*.txt): %v\n", allTxt)
}

// =============================================================================
// SECTION 6: WALKING DIRECTORIES
// =============================================================================
//
// filepath.WalkDir(root string, fn WalkDirFunc) error
//   Walk directory tree, calling fn for each entry
//   More efficient than Walk (doesn't call Lstat for every file)
//
// filepath.Walk(root string, fn WalkFunc) error
//   Legacy version, calls Lstat for each file
//   Prefer WalkDir in new code
//
// WalkDirFunc signature:
//   func(path string, d DirEntry, err error) error
//
// RETURN VALUES
//
//   nil               - Continue walking
//   filepath.SkipDir  - Skip this directory
//   filepath.SkipAll  - Stop walking entirely (Go 1.20+)
//   any error         - Stop walking, return error
//
// =============================================================================

func DemonstrateWalking() {
	// WalkDir - efficient directory traversal
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", path, err)
			return nil // Continue despite errors
		}

		// Skip hidden directories
		if d.IsDir() && len(d.Name()) > 1 && d.Name()[0] == '.' {
			fmt.Printf("Skipping hidden: %s\n", path)
			return filepath.SkipDir
		}

		// Print info
		if d.IsDir() {
			fmt.Printf("[DIR]  %s\n", path)
		} else {
			fmt.Printf("[FILE] %s\n", path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Walk error: %v\n", err)
	}

	// Example: Find all Go files
	var goFiles []string
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() && filepath.Ext(path) == ".go" {
			goFiles = append(goFiles, path)
		}
		return nil
	})
	fmt.Printf("Go files found: %v\n", goFiles)

	// Example: Count files and directories
	var files, dirs int
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			dirs++
		} else {
			files++
		}
		return nil
	})
	fmt.Printf("Found %d files, %d directories\n", files, dirs)
}

// =============================================================================
// SECTION 7: VOLUME AND ROOT (WINDOWS-SPECIFIC)
// =============================================================================
//
// filepath.VolumeName(path string) string
//   Return volume name (Windows: "C:", Unix: "")
//
// filepath.SplitList(path string) []string
//   Split PATH-style list (colon or semicolon separated)
//
// filepath.EvalSymlinks(path string) (string, error)
//   Evaluate any symlinks in path
//
// =============================================================================

func DemonstrateVolumeAndRoot() {
	// VolumeName (mainly for Windows)
	paths := []string{
		"/home/user/file.txt",
		"C:\\Users\\name\\file.txt",
		"relative/path",
	}

	for _, p := range paths {
		vol := filepath.VolumeName(p)
		fmt.Printf("VolumeName(%q) = %q\n", p, vol)
	}

	// SplitList - split PATH environment variable
	pathEnv := os.Getenv("PATH")
	pathDirs := filepath.SplitList(pathEnv)
	fmt.Printf("PATH has %d directories\n", len(pathDirs))
	if len(pathDirs) > 0 {
		fmt.Printf("First: %s\n", pathDirs[0])
	}

	// EvalSymlinks - resolve symlinks
	resolved, err := filepath.EvalSymlinks(".")
	if err != nil {
		fmt.Printf("EvalSymlinks error: %v\n", err)
	} else {
		fmt.Printf("EvalSymlinks(.): %s\n", resolved)
	}
}

// =============================================================================
// SECTION 8: COMMON PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Get filename without extension
	removeExt := func(path string) string {
		ext := filepath.Ext(path)
		return path[:len(path)-len(ext)]
	}
	fmt.Printf("Pattern 1 - Without ext: %s\n", removeExt("document.txt"))

	// Pattern 2: Change extension
	changeExt := func(path, newExt string) string {
		ext := filepath.Ext(path)
		return path[:len(path)-len(ext)] + newExt
	}
	fmt.Printf("Pattern 2 - Change ext: %s\n", changeExt("file.txt", ".md"))

	// Pattern 3: Ensure directory exists
	ensureDir := func(path string) error {
		dir := filepath.Dir(path)
		return os.MkdirAll(dir, 0755)
	}
	_ = ensureDir
	fmt.Println("Pattern 3 - ensureDir: os.MkdirAll(filepath.Dir(path))")

	// Pattern 4: Safe path join (prevent directory traversal)
	safePath := func(base, requested string) (string, error) {
		// Clean and join
		full := filepath.Join(base, filepath.Clean("/"+requested))
		// Verify it's still under base
		if !hasPrefix(full, base) {
			return "", fmt.Errorf("path escapes base directory")
		}
		return full, nil
	}
	safe, _ := safePath("/var/www", "images/photo.jpg")
	fmt.Printf("Pattern 4 - Safe path: %s\n", safe)

	_, err := safePath("/var/www", "../etc/passwd")
	fmt.Printf("Pattern 4 - Blocked: %v\n", err)

	// Pattern 5: Find files by extension
	findByExt := func(root, ext string) ([]string, error) {
		var files []string
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if !d.IsDir() && filepath.Ext(path) == ext {
				files = append(files, path)
			}
			return nil
		})
		return files, err
	}
	_ = findByExt
	fmt.Println("Pattern 5 - findByExt: WalkDir + Ext check")

	// Pattern 6: Relative path from current directory
	relFromCwd := func(path string) string {
		cwd, _ := os.Getwd()
		rel, err := filepath.Rel(cwd, path)
		if err != nil {
			return path
		}
		return rel
	}
	_ = relFromCwd
	fmt.Println("Pattern 6 - relFromCwd: Rel(Getwd(), path)")

	// Pattern 7: Expand ~ to home directory
	expandHome := func(path string) string {
		if len(path) > 0 && path[0] == '~' {
			home, _ := os.UserHomeDir()
			return filepath.Join(home, path[1:])
		}
		return path
	}
	fmt.Printf("Pattern 7 - Expand ~: %s\n", expandHome("~/documents"))
}

// hasPrefix checks if path starts with prefix (after cleaning)
func hasPrefix(path, prefix string) bool {
	path = filepath.Clean(path)
	prefix = filepath.Clean(prefix)
	return len(path) >= len(prefix) && path[:len(prefix)] == prefix
}

// =============================================================================
// SECTION 9: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Hardcoding separators
	_ = `
	// WRONG
	path := "home" + "/" + "user"  // Fails on Windows

	// RIGHT
	path := filepath.Join("home", "user")
	`
	fmt.Println("Mistake 1: Never hardcode / or \\, use filepath.Join")

	// Mistake 2: Using path instead of path/filepath
	_ = `
	// WRONG - path uses forward slash only
	import "path"
	p := path.Join("a", "b")  // Always a/b, even on Windows

	// RIGHT - filepath uses OS separator
	import "path/filepath"
	p := filepath.Join("a", "b")  // a/b on Unix, a\b on Windows
	`
	fmt.Println("Mistake 2: Use path/filepath for files, path for URLs")

	// Mistake 3: Not handling WalkDir errors
	_ = `
	// WRONG - ignoring error parameter
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		// err might be non-nil!
		name := d.Name()  // Panic if err != nil
		...
	})

	// RIGHT
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil  // or return err to stop
		}
		name := d.Name()
		...
	})
	`
	fmt.Println("Mistake 3: Always check err in WalkDirFunc")

	// Mistake 4: Assuming Glob supports **
	_ = `
	// WRONG - ** not supported
	files, _ := filepath.Glob("**/*.go")  // Returns nil!

	// RIGHT - use WalkDir for recursive search
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".go" {
			files = append(files, path)
		}
		return nil
	})
	`
	fmt.Println("Mistake 4: Glob doesn't support **, use WalkDir")

	// Mistake 5: Not cleaning user input paths
	_ = `
	// WRONG - allows directory traversal
	path := filepath.Join(baseDir, userInput)

	// RIGHT - clean and validate
	path := filepath.Join(baseDir, filepath.Clean("/"+userInput))
	if !strings.HasPrefix(path, baseDir) {
		return errors.New("invalid path")
	}
	`
	fmt.Println("Mistake 5: Clean and validate user-provided paths")

	// Mistake 6: Confusing Split and Dir
	_ = `
	// Split returns dir WITH trailing separator
	dir, file := filepath.Split("/a/b/c")  // dir="/a/b/", file="c"

	// Dir returns WITHOUT trailing separator
	dir := filepath.Dir("/a/b/c")  // dir="/a/b"
	`
	fmt.Println("Mistake 6: Split includes trailing sep, Dir doesn't")
}

// =============================================================================
// SECTION 10: PATH VS FILEPATH COMPARISON
// =============================================================================
//
// path package (import "path"):
//   - Always uses forward slash
//   - For URLs, virtual paths, non-filesystem paths
//   - path.Join, path.Split, path.Dir, etc.
//
// path/filepath package (import "path/filepath"):
//   - Uses OS-specific separator
//   - For actual file system paths
//   - filepath.Join, filepath.Split, filepath.Dir, etc.
//
// RULE: If it touches the file system, use filepath
//
// =============================================================================

func DemonstratePathVsFilepath() {
	// path (always forward slash)
	// import "path"
	// urlPath := path.Join("api", "v1", "users")  // "api/v1/users" everywhere

	// filepath (OS separator)
	filePath := filepath.Join("home", "user", "file")
	fmt.Printf("filepath.Join: %s\n", filePath)

	// Use path for URLs
	_ = `
	import "path"
	url := "https://example.com/" + path.Join("api", "users", id)
	`

	// Use filepath for files
	_ = `
	import "path/filepath"
	configPath := filepath.Join(homeDir, ".config", "app", "config.json")
	data, err := os.ReadFile(configPath)
	`
}

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateBasics()
	DemonstrateJoinSplit()
	DemonstrateClean()
	DemonstrateAbsRel()
	DemonstrateMatching()
	DemonstrateWalking()
	DemonstrateVolumeAndRoot()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
	DemonstratePathVsFilepath()
}
// Package osmod provides comprehensive documentation and working examples
// for Go's os package - file operations, environment, process management.
package osmod

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

// =============================================================================
// SECTION 1: THE PHILOSOPHY OF OS PACKAGE
// =============================================================================
//
// The os package provides a platform-independent interface to operating system
// functionality. It's your gateway to:
//
//   - File operations (create, read, write, delete)
//   - Environment variables
//   - Process information
//   - Signals
//   - Working directory
//
// KEY TYPES
//
//   os.File      - An open file descriptor
//   os.FileInfo  - File metadata (via fs.FileInfo)
//   os.DirEntry  - Directory entry (via fs.DirEntry)
//   os.FileMode  - File permissions (via fs.FileMode)
//   os.PathError - Error with path context
//
// STANDARD FILE HANDLES
//
//   os.Stdin   - Standard input (*os.File)
//   os.Stdout  - Standard output (*os.File)
//   os.Stderr  - Standard error (*os.File)
//
// =============================================================================

// =============================================================================
// SECTION 2: READING FILES
// =============================================================================
//
// SIMPLE: Read entire file into memory
//
//   os.ReadFile(name) ([]byte, error)
//
// STREAM: Open file for reading
//
//   os.Open(name) (*File, error)     - Read-only
//   file.Read(b []byte) (n int, err error)
//   file.Close() error
//
// =============================================================================

func DemonstrateReadingFiles() {
	// Create a test file first
	os.WriteFile("test.txt", []byte("Hello, World!\nLine 2\nLine 3"), 0644)
	defer os.Remove("test.txt")

	// ReadFile - slurp entire file (simple, for small files)
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("ReadFile error: %v\n", err)
		return
	}
	fmt.Printf("ReadFile: %s\n", data)

	// Open + Read - for streaming/large files
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("Open error: %v\n", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 5)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			fmt.Printf("Read %d bytes: %q\n", n, buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Read error: %v\n", err)
			break
		}
	}

	// Read entire file via io.ReadAll
	file2, _ := os.Open("test.txt")
	defer file2.Close()
	all, _ := io.ReadAll(file2)
	fmt.Printf("io.ReadAll: %s\n", all)
}

// =============================================================================
// SECTION 3: WRITING FILES
// =============================================================================
//
// SIMPLE: Write entire file at once
//
//   os.WriteFile(name, data, perm) error
//
// STREAM: Open file for writing
//
//   os.Create(name) (*File, error)           - Create/truncate
//   os.OpenFile(name, flag, perm) (*File, error) - Full control
//   file.Write(b []byte) (n int, err error)
//   file.WriteString(s string) (n int, err error)
//
// FLAGS (combine with |)
//
//   os.O_RDONLY   - Read-only
//   os.O_WRONLY   - Write-only
//   os.O_RDWR     - Read-write
//   os.O_APPEND   - Append to file
//   os.O_CREATE   - Create if not exists
//   os.O_EXCL     - Error if exists (with O_CREATE)
//   os.O_TRUNC    - Truncate if exists
//   os.O_SYNC     - Synchronous I/O
//
// =============================================================================

func DemonstrateWritingFiles() {
	// WriteFile - write entire file at once (creates or truncates)
	err := os.WriteFile("output.txt", []byte("Hello, File!"), 0644)
	if err != nil {
		fmt.Printf("WriteFile error: %v\n", err)
		return
	}
	fmt.Println("WriteFile: created output.txt")
	defer os.Remove("output.txt")

	// Create - create new file (truncates if exists)
	file, _ := os.Create("created.txt")
	file.WriteString("Line 1\n")
	file.WriteString("Line 2\n")
	file.Close()
	fmt.Println("Create: created created.txt")
	defer os.Remove("created.txt")

	// OpenFile - full control
	// Append mode
	appendFile, _ := os.OpenFile("created.txt", os.O_APPEND|os.O_WRONLY, 0644)
	appendFile.WriteString("Appended line\n")
	appendFile.Close()
	fmt.Println("OpenFile O_APPEND: appended to created.txt")

	// Create only if not exists
	_, err = os.OpenFile("created.txt", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("O_EXCL: file exists, got error: %v\n", err)
	}

	// Read-write mode
	rwFile, _ := os.OpenFile("created.txt", os.O_RDWR, 0644)
	buf := make([]byte, 10)
	rwFile.Read(buf)
	fmt.Printf("Read from RW file: %q\n", buf)
	rwFile.Close()
}

// =============================================================================
// SECTION 4: FILE PERMISSIONS
// =============================================================================
//
// FILE MODE (os.FileMode / fs.FileMode)
//
// Octal notation: 0755, 0644, etc.
//
//   Owner  Group  Other
//   rwx    rwx    rwx
//   421    421    421
//
//   0755 = rwxr-xr-x (owner all, others read+execute)
//   0644 = rw-r--r-- (owner read+write, others read)
//   0700 = rwx------ (owner all, others nothing)
//   0600 = rw------- (owner read+write only)
//
// SPECIAL BITS
//
//   os.ModeDir       - Is a directory
//   os.ModeSymlink   - Is a symlink
//   os.ModePerm      - Permission bits mask (0777)
//
// =============================================================================

func DemonstratePermissions() {
	// Create with specific permissions
	os.WriteFile("private.txt", []byte("secret"), 0600) // Owner read+write only
	defer os.Remove("private.txt")

	os.WriteFile("public.txt", []byte("public"), 0644) // Owner rw, others r
	defer os.Remove("public.txt")

	// Check permissions
	info, _ := os.Stat("private.txt")
	mode := info.Mode()
	fmt.Printf("private.txt mode: %v (%o)\n", mode, mode.Perm())

	// Change permissions
	os.Chmod("private.txt", 0644)
	info2, _ := os.Stat("private.txt")
	fmt.Printf("After chmod: %v\n", info2.Mode())

	// Check permission bits
	fmt.Printf("Is directory: %v\n", mode.IsDir())
	fmt.Printf("Is regular: %v\n", mode.IsRegular())
	fmt.Printf("Permission bits: %o\n", mode.Perm())
}

// =============================================================================
// SECTION 5: FILE INFO AND STAT
// =============================================================================
//
// os.Stat(name) (FileInfo, error)   - Get file info (follows symlinks)
// os.Lstat(name) (FileInfo, error)  - Get file info (doesn't follow symlinks)
// file.Stat() (FileInfo, error)     - Stat an open file
//
// FileInfo interface (fs.FileInfo):
//
//   Name() string       - Base name
//   Size() int64        - Size in bytes
//   Mode() FileMode     - Permission bits
//   ModTime() time.Time - Modification time
//   IsDir() bool        - Is directory?
//   Sys() interface{}   - System-specific info
//
// =============================================================================

func DemonstrateFileInfo() {
	// Create test file
	os.WriteFile("info.txt", []byte("Some content here"), 0644)
	defer os.Remove("info.txt")

	// Get file info
	info, err := os.Stat("info.txt")
	if err != nil {
		fmt.Printf("Stat error: %v\n", err)
		return
	}

	fmt.Printf("Name: %s\n", info.Name())
	fmt.Printf("Size: %d bytes\n", info.Size())
	fmt.Printf("Mode: %v\n", info.Mode())
	fmt.Printf("ModTime: %v\n", info.ModTime())
	fmt.Printf("IsDir: %v\n", info.IsDir())

	// Check if file exists
	_, err = os.Stat("nonexistent.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	}

	// Stat vs Lstat (for symlinks)
	// Lstat doesn't follow symlinks
	// info, _ := os.Lstat("symlink")
	// if info.Mode()&os.ModeSymlink != 0 {
	//     fmt.Println("It's a symlink")
	// }
}

// =============================================================================
// SECTION 6: DIRECTORIES
// =============================================================================
//
// os.Mkdir(name, perm) error        - Create single directory
// os.MkdirAll(path, perm) error     - Create directory tree (like mkdir -p)
// os.Remove(name) error             - Remove file or empty directory
// os.RemoveAll(path) error          - Remove directory tree (like rm -rf)
// os.ReadDir(name) ([]DirEntry, error) - List directory contents
// os.Rename(old, new) error         - Rename/move file or directory
//
// =============================================================================

func DemonstrateDirectories() {
	// Create single directory
	err := os.Mkdir("testdir", 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		fmt.Printf("Mkdir error: %v\n", err)
	}
	fmt.Println("Created testdir/")

	// Create directory tree
	os.MkdirAll("testdir/sub1/sub2", 0755)
	fmt.Println("Created testdir/sub1/sub2/")

	// Create files in directories
	os.WriteFile("testdir/file1.txt", []byte("1"), 0644)
	os.WriteFile("testdir/file2.txt", []byte("2"), 0644)
	os.WriteFile("testdir/sub1/file3.txt", []byte("3"), 0644)

	// List directory contents
	entries, _ := os.ReadDir("testdir")
	fmt.Println("Contents of testdir/:")
	for _, entry := range entries {
		info, _ := entry.Info()
		if entry.IsDir() {
			fmt.Printf("  [DIR]  %s/\n", entry.Name())
		} else {
			fmt.Printf("  [FILE] %s (%d bytes)\n", entry.Name(), info.Size())
		}
	}

	// Rename/move
	os.Rename("testdir/file1.txt", "testdir/renamed.txt")
	fmt.Println("Renamed file1.txt to renamed.txt")

	// Remove single file
	os.Remove("testdir/renamed.txt")
	fmt.Println("Removed renamed.txt")

	// Remove directory tree
	os.RemoveAll("testdir")
	fmt.Println("Removed testdir/ and all contents")
}

// =============================================================================
// SECTION 7: WORKING DIRECTORY AND PATHS
// =============================================================================
//
// os.Getwd() (string, error)       - Get current working directory
// os.Chdir(dir string) error       - Change working directory
// os.TempDir() string              - Get temp directory path
// os.UserHomeDir() (string, error) - Get user home directory
// os.UserCacheDir() (string, error)- Get user cache directory
// os.UserConfigDir() (string, error)- Get user config directory
// os.Executable() (string, error)  - Get path to current executable
//
// =============================================================================

func DemonstrateWorkingDirectory() {
	// Get current directory
	cwd, _ := os.Getwd()
	fmt.Printf("Current directory: %s\n", cwd)

	// Change directory
	os.Mkdir("tempwork", 0755)
	defer os.RemoveAll("tempwork")

	os.Chdir("tempwork")
	newCwd, _ := os.Getwd()
	fmt.Printf("After Chdir: %s\n", newCwd)

	os.Chdir(cwd) // Go back

	// Temp directory
	fmt.Printf("Temp dir: %s\n", os.TempDir())

	// User directories
	home, _ := os.UserHomeDir()
	fmt.Printf("Home dir: %s\n", home)

	cache, _ := os.UserCacheDir()
	fmt.Printf("Cache dir: %s\n", cache)

	config, _ := os.UserConfigDir()
	fmt.Printf("Config dir: %s\n", config)

	// Current executable
	exe, _ := os.Executable()
	fmt.Printf("Executable: %s\n", exe)
}

// =============================================================================
// SECTION 8: TEMPORARY FILES AND DIRECTORIES
// =============================================================================
//
// os.CreateTemp(dir, pattern) (*File, error)
//   Create temp file. Pattern can include * for random part.
//   dir="" uses os.TempDir()
//
// os.MkdirTemp(dir, pattern) (string, error)
//   Create temp directory.
//
// =============================================================================

func DemonstrateTempFiles() {
	// Create temp file
	tmpFile, err := os.CreateTemp("", "myapp-*.txt")
	if err != nil {
		fmt.Printf("CreateTemp error: %v\n", err)
		return
	}
	fmt.Printf("Temp file: %s\n", tmpFile.Name())

	// Write to temp file
	tmpFile.WriteString("Temporary data")
	tmpFile.Close()

	// Clean up
	os.Remove(tmpFile.Name())

	// Create temp file in specific directory
	os.Mkdir("mytmp", 0755)
	defer os.RemoveAll("mytmp")

	tmpFile2, _ := os.CreateTemp("mytmp", "data-*.json")
	fmt.Printf("Temp file in mytmp: %s\n", tmpFile2.Name())
	tmpFile2.Close()

	// Create temp directory
	tmpDir, _ := os.MkdirTemp("", "myapp-*")
	fmt.Printf("Temp dir: %s\n", tmpDir)

	// Use temp directory
	os.WriteFile(filepath.Join(tmpDir, "data.txt"), []byte("temp"), 0644)

	// Clean up
	os.RemoveAll(tmpDir)
}

// =============================================================================
// SECTION 9: ENVIRONMENT VARIABLES
// =============================================================================
//
// os.Getenv(key) string              - Get env var (empty if not set)
// os.LookupEnv(key) (string, bool)   - Get env var with existence check
// os.Setenv(key, value) error        - Set env var
// os.Unsetenv(key) error             - Remove env var
// os.Clearenv()                      - Remove all env vars
// os.Environ() []string              - Get all env vars as "KEY=value"
// os.ExpandEnv(s) string             - Expand $VAR or ${VAR} in string
//
// =============================================================================

func DemonstrateEnvironment() {
	// Get environment variable
	path := os.Getenv("PATH")
	fmt.Printf("PATH length: %d\n", len(path))

	// Getenv returns empty string if not set
	missing := os.Getenv("DEFINITELY_NOT_SET")
	fmt.Printf("Missing var: %q\n", missing)

	// LookupEnv distinguishes empty from missing
	val, exists := os.LookupEnv("DEFINITELY_NOT_SET")
	fmt.Printf("LookupEnv: val=%q, exists=%v\n", val, exists)

	// Set environment variable
	os.Setenv("MY_APP_DEBUG", "true")
	fmt.Printf("MY_APP_DEBUG: %s\n", os.Getenv("MY_APP_DEBUG"))

	// Expand environment variables in string
	os.Setenv("NAME", "Alice")
	expanded := os.ExpandEnv("Hello, $NAME! Your home is ${HOME}.")
	fmt.Printf("Expanded: %s\n", expanded)

	// Unset
	os.Unsetenv("MY_APP_DEBUG")
	os.Unsetenv("NAME")

	// List all environment variables
	// for _, env := range os.Environ() {
	//     fmt.Println(env)
	// }
}

// =============================================================================
// SECTION 10: PROCESS AND ARGS
// =============================================================================
//
// os.Args []string           - Command line arguments (os.Args[0] is program)
// os.Getpid() int            - Current process ID
// os.Getppid() int           - Parent process ID
// os.Getuid() int            - User ID
// os.Getgid() int            - Group ID
// os.Hostname() (string, error) - System hostname
// os.Exit(code int)          - Exit immediately with code
//
// =============================================================================

func DemonstrateProcess() {
	// Command line arguments
	fmt.Printf("Program: %s\n", os.Args[0])
	fmt.Printf("Args: %v\n", os.Args[1:])

	// Process IDs
	fmt.Printf("PID: %d\n", os.Getpid())
	fmt.Printf("Parent PID: %d\n", os.Getppid())

	// User/Group IDs (Unix)
	fmt.Printf("UID: %d\n", os.Getuid())
	fmt.Printf("GID: %d\n", os.Getgid())

	// Hostname
	hostname, _ := os.Hostname()
	fmt.Printf("Hostname: %s\n", hostname)

	// Exit - use sparingly, skips defers!
	// os.Exit(0)  // Success
	// os.Exit(1)  // Failure
}

// =============================================================================
// SECTION 11: SIGNALS
// =============================================================================
//
// os/signal package for handling OS signals.
//
// signal.Notify(c chan os.Signal, sig ...os.Signal)
//   Register channel to receive signals
//
// signal.Stop(c chan os.Signal)
//   Stop receiving signals on channel
//
// COMMON SIGNALS
//
//   syscall.SIGINT   - Interrupt (Ctrl+C)
//   syscall.SIGTERM  - Termination request
//   syscall.SIGHUP   - Hangup
//   syscall.SIGUSR1  - User-defined signal 1
//
// =============================================================================

func DemonstrateSignals() {
	// Create channel for signals
	sigChan := make(chan os.Signal, 1)

	// Register for SIGINT and SIGTERM
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Signal handling registered (not waiting in this demo)")

	// Typical pattern for graceful shutdown:
	_ = `
	go func() {
		sig := <-sigChan
		fmt.Printf("Received signal: %v\n", sig)
		// Cleanup...
		os.Exit(0)
	}()
	`

	// Or with context:
	_ = `
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		fmt.Println("Shutting down...")
	}
	`

	// Stop receiving signals
	signal.Stop(sigChan)
}

// =============================================================================
// SECTION 12: ERROR HANDLING
// =============================================================================
//
// SENTINEL ERRORS
//
//   os.ErrNotExist    - File/directory doesn't exist
//   os.ErrExist       - File/directory already exists
//   os.ErrPermission  - Permission denied
//   os.ErrClosed      - File already closed
//   os.ErrInvalid     - Invalid argument
//
// Use errors.Is() to check:
//
//   if errors.Is(err, os.ErrNotExist) { ... }
//
// PATH ERRORS
//
// Many os functions return *os.PathError with additional context:
//
//   type PathError struct {
//       Op   string  // Operation ("open", "read", etc.)
//       Path string  // File path
//       Err  error   // Underlying error
//   }
//
// =============================================================================

func DemonstrateErrors() {
	// Check for specific errors
	_, err := os.Open("nonexistent.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	}

	// PathError contains context
	if pathErr, ok := err.(*fs.PathError); ok {
		fmt.Printf("Operation: %s\n", pathErr.Op)
		fmt.Printf("Path: %s\n", pathErr.Path)
		fmt.Printf("Error: %v\n", pathErr.Err)
	}

	// Permission error
	os.WriteFile("readonly.txt", []byte("test"), 0000)
	defer os.Remove("readonly.txt")

	_, err = os.OpenFile("readonly.txt", os.O_WRONLY, 0)
	if errors.Is(err, os.ErrPermission) {
		fmt.Println("Permission denied")
	}

	// Cleanup - reset permissions so we can delete
	os.Chmod("readonly.txt", 0644)

	// Check if path is directory
	os.Mkdir("testdir2", 0755)
	defer os.RemoveAll("testdir2")

	_, err = os.ReadFile("testdir2")
	if err != nil {
		fmt.Printf("ReadFile on dir: %v\n", err)
	}
}

// =============================================================================
// SECTION 13: COMMON PATTERNS
// =============================================================================

func DemonstratePatterns() {
	// Pattern 1: Check if file exists
	fileExists := func(path string) bool {
		_, err := os.Stat(path)
		return !errors.Is(err, os.ErrNotExist)
	}
	fmt.Printf("Pattern 1 - exists: %v\n", fileExists("os.go"))

	// Pattern 2: Ensure directory exists
	ensureDir := func(path string) error {
		return os.MkdirAll(path, 0755)
	}
	_ = ensureDir
	fmt.Println("Pattern 2 - ensureDir: os.MkdirAll(path, 0755)")

	// Pattern 3: Safe file write (atomic)
	_ = `
	func writeFileAtomic(path string, data []byte) error {
		// Write to temp file first
		tmp, err := os.CreateTemp(filepath.Dir(path), ".tmp-*")
		if err != nil {
			return err
		}
		tmpPath := tmp.Name()

		// Cleanup on error
		success := false
		defer func() {
			if !success {
				os.Remove(tmpPath)
			}
		}()

		if _, err := tmp.Write(data); err != nil {
			tmp.Close()
			return err
		}
		if err := tmp.Close(); err != nil {
			return err
		}

		// Atomic rename
		if err := os.Rename(tmpPath, path); err != nil {
			return err
		}

		success = true
		return nil
	}
	`
	fmt.Println("Pattern 3 - Atomic write: temp file + rename")

	// Pattern 4: Read config with fallback
	getConfig := func(key, fallback string) string {
		if val, ok := os.LookupEnv(key); ok {
			return val
		}
		return fallback
	}
	port := getConfig("PORT", "8080")
	fmt.Printf("Pattern 4 - Config: PORT=%s\n", port)

	// Pattern 5: Graceful shutdown
	_ = `
	ctx, stop := signal.NotifyContext(context.Background(), 
		syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := &http.Server{...}
	go server.ListenAndServe()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(shutdownCtx)
	`
	fmt.Println("Pattern 5 - Graceful shutdown with signal.NotifyContext")

	// Pattern 6: Walk and process files
	_ = `
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil  // Skip errors
		}
		if !d.IsDir() && filepath.Ext(path) == ".go" {
			processFile(path)
		}
		return nil
	})
	`
	fmt.Println("Pattern 6 - Walk + filter by extension")
}

// =============================================================================
// SECTION 14: COMMON MISTAKES
// =============================================================================

func DemonstrateCommonMistakes() {
	// Mistake 1: Not closing files
	_ = `
	// WRONG - file never closed, leaks file descriptor
	file, _ := os.Open("data.txt")
	data, _ := io.ReadAll(file)

	// RIGHT - always defer close
	file, err := os.Open("data.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	`
	fmt.Println("Mistake 1: Always defer file.Close()")

	// Mistake 2: Checking error after defer
	_ = `
	// WRONG - error check after defer is useless
	file, _ := os.Open("data.txt")
	defer file.Close()
	if file == nil {
		// Too late!
	}

	// RIGHT - check error before defer
	file, err := os.Open("data.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	`
	fmt.Println("Mistake 2: Check error before defer Close")

	// Mistake 3: Using os.Exit in library code
	_ = `
	// WRONG - prevents cleanup, skips defers
	func LibraryFunc() {
		os.Exit(1)  // Never do this in libraries!
	}

	// RIGHT - return error
	func LibraryFunc() error {
		return errors.New("something failed")
	}
	`
	fmt.Println("Mistake 3: Never os.Exit() in library code")

	// Mistake 4: Ignoring MkdirAll return
	_ = `
	// WRONG - might silently fail
	os.MkdirAll("path/to/dir", 0755)
	os.WriteFile("path/to/dir/file.txt", data, 0644)

	// RIGHT - check error
	if err := os.MkdirAll("path/to/dir", 0755); err != nil {
		return err
	}
	`
	fmt.Println("Mistake 4: Always check MkdirAll error")

	// Mistake 5: Getenv vs LookupEnv
	_ = `
	// WRONG - can't tell if empty or missing
	val := os.Getenv("KEY")
	if val == "" {
		// Is it empty or not set?
	}

	// RIGHT - distinguish empty from missing
	val, exists := os.LookupEnv("KEY")
	if !exists {
		// Not set
	} else if val == "" {
		// Set but empty
	}
	`
	fmt.Println("Mistake 5: Use LookupEnv to distinguish empty from missing")

	// Mistake 6: Race condition in file exists check
	_ = `
	// WRONG - TOCTOU race condition
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Create(path)  // Another process might create it!
	}

	// RIGHT - use O_EXCL for atomic create-if-not-exists
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if errors.Is(err, os.ErrExist) {
		// File already exists
	}
	`
	fmt.Println("Mistake 6: Use O_EXCL for atomic create-if-not-exists")
}

// Ensure we use time package to avoid unused import error
var _ = time.Second

// AllDemonstrations is not meant to be called - this file is for reading.
func AllDemonstrations() {
	DemonstrateReadingFiles()
	DemonstrateWritingFiles()
	DemonstratePermissions()
	DemonstrateFileInfo()
	DemonstrateDirectories()
	DemonstrateWorkingDirectory()
	DemonstrateTempFiles()
	DemonstrateEnvironment()
	DemonstrateProcess()
	DemonstrateSignals()
	DemonstrateErrors()
	DemonstratePatterns()
	DemonstrateCommonMistakes()
}
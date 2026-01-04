// Package os_practice is your daily practice space for the os package.
//
// INSTRUCTIONS:
// 1. Fill in all the TODOs from memory
// 2. Check your answers against os/os.go
// 3. Note what you missed - focus on those tomorrow
package os_practice

// =============================================================================
// EXERCISE 1: STANDARD HANDLES
// =============================================================================

func PracticeStandardHandles() {
	// TODO: Name the three standard file handles
	// os.???   - standard input
	// os.???   - standard output
	// os.???   - standard error

	// TODO: What type are they?
	// Answer: *os.???
}

// =============================================================================
// EXERCISE 2: READING FILES
// =============================================================================

func PracticeReadingFiles() {
	// TODO: Read entire file into []byte
	// data, err := os.???(filename)

	// TODO: Open file for reading (read-only)
	// file, err := os.???(filename)

	// TODO: Always close the file
	// defer file.???()

	// TODO: Read from file into buffer
	// n, err := file.???(buf)
}

// =============================================================================
// EXERCISE 3: WRITING FILES
// =============================================================================

func PracticeWritingFiles() {
	// TODO: Write entire file at once
	// err := os.???(filename, data, perm)

	// TODO: Create new file (truncates if exists)
	// file, err := os.???(filename)

	// TODO: Open with full control
	// file, err := os.???(filename, flags, perm)

	// TODO: Write bytes to file
	// n, err := file.???(data)

	// TODO: Write string to file
	// n, err := file.???(str)
}

// =============================================================================
// EXERCISE 4: OPEN FILE FLAGS
// =============================================================================

func PracticeFlags() {
	// TODO: Name the common flags
	// os.O_???    - read only
	// os.O_???    - write only
	// os.O_???    - read and write
	// os.O_???    - append to file
	// os.O_???    - create if not exists
	// os.O_???    - error if exists (with O_CREATE)
	// os.O_???    - truncate if exists

	// TODO: Open for appending
	// file, _ := os.OpenFile(name, os.O_???|os.O_???, 0644)

	// TODO: Create only if not exists
	// file, _ := os.OpenFile(name, os.O_???|os.O_???|os.O_???, 0644)
}

// =============================================================================
// EXERCISE 5: FILE PERMISSIONS
// =============================================================================

func PracticePermissions() {
	// TODO: Common permission values
	// 0??? = rwxr-xr-x (755)
	// 0??? = rw-r--r-- (644)
	// 0??? = rwx------ (700)
	// 0??? = rw------- (600)

	// TODO: Change file permissions
	// err := os.???(filename, 0644)

	// TODO: Get permission bits from FileMode
	// mode := info.Mode()
	// perm := mode.???()
}

// =============================================================================
// EXERCISE 6: FILE INFO
// =============================================================================

func PracticeFileInfo() {
	// TODO: Get file info (follows symlinks)
	// info, err := os.???(filename)

	// TODO: Get file info (doesn't follow symlinks)
	// info, err := os.???(filename)

	// TODO: FileInfo methods
	// info.???()   - base name
	// info.???()   - size in bytes
	// info.???()   - file mode
	// info.???()   - modification time
	// info.???()   - is directory?

	// TODO: Check if file exists
	// _, err := os.Stat(filename)
	// if errors.Is(err, os.???) {
	//     // doesn't exist
	// }
}

// =============================================================================
// EXERCISE 7: DIRECTORIES
// =============================================================================

func PracticeDirectories() {
	// TODO: Create single directory
	// err := os.???(name, perm)

	// TODO: Create directory tree (mkdir -p)
	// err := os.???(path, perm)

	// TODO: Remove file or empty directory
	// err := os.???(name)

	// TODO: Remove directory tree (rm -rf)
	// err := os.???(path)

	// TODO: List directory contents
	// entries, err := os.???(name)

	// TODO: Rename/move file
	// err := os.???(oldpath, newpath)
}

// =============================================================================
// EXERCISE 8: WORKING DIRECTORY
// =============================================================================

func PracticeWorkingDirectory() {
	// TODO: Get current working directory
	// cwd, err := os.???()

	// TODO: Change working directory
	// err := os.???(dir)

	// TODO: Get temp directory
	// tmp := os.???()

	// TODO: Get user home directory
	// home, err := os.???()

	// TODO: Get user cache directory
	// cache, err := os.???()

	// TODO: Get user config directory
	// config, err := os.???()

	// TODO: Get current executable path
	// exe, err := os.???()
}

// =============================================================================
// EXERCISE 9: TEMP FILES
// =============================================================================

func PracticeTempFiles() {
	// TODO: Create temp file
	// file, err := os.???(dir, pattern)

	// TODO: Create temp directory
	// dir, err := os.???(parentDir, pattern)

	// TODO: Pattern with wildcard
	// "myapp-*.txt" → "myapp-abc123.txt"
}

// =============================================================================
// EXERCISE 10: ENVIRONMENT
// =============================================================================

func PracticeEnvironment() {
	// TODO: Get env var (empty if not set)
	// val := os.???(key)

	// TODO: Get env var with existence check
	// val, exists := os.???(key)

	// TODO: Set env var
	// err := os.???(key, value)

	// TODO: Remove env var
	// err := os.???(key)

	// TODO: Get all env vars
	// envs := os.???()

	// TODO: Expand $VAR in string
	// expanded := os.???(str)
}

// =============================================================================
// EXERCISE 11: PROCESS
// =============================================================================

func PracticeProcess() {
	// TODO: Command line arguments
	// args := os.???

	// TODO: Program name
	// program := os.???[0]

	// TODO: Process ID
	// pid := os.???()

	// TODO: Parent process ID
	// ppid := os.???()

	// TODO: User ID
	// uid := os.???()

	// TODO: Hostname
	// host, err := os.???()

	// TODO: Exit program
	// os.???(code)
}

// =============================================================================
// EXERCISE 12: SIGNALS
// =============================================================================

func PracticeSignals() {
	// TODO: Register for signals
	// signal.???(channel, signals...)

	// TODO: Stop receiving signals
	// signal.???(channel)

	// TODO: Common signals
	// syscall.???    - Ctrl+C
	// syscall.???    - termination

	// TODO: Context-based signal handling
	// ctx, stop := signal.???(parentCtx, signals...)
	// defer stop()
}

// =============================================================================
// EXERCISE 13: ERRORS
// =============================================================================

func PracticeErrors() {
	// TODO: Sentinel errors
	// os.???       - file doesn't exist
	// os.???       - file already exists
	// os.???       - permission denied
	// os.???       - file already closed

	// TODO: Check for specific error
	// if errors.???(err, os.ErrNotExist) { ... }

	// TODO: PathError fields
	// pathErr.???   - operation name
	// pathErr.???   - file path
	// pathErr.???   - underlying error
}

// =============================================================================
// EXERCISE 14: COMMON PATTERNS
// =============================================================================

func PracticePatterns() {
	// TODO: Check if file exists
	// _, err := os.???(path)
	// exists := !errors.Is(err, os.???)

	// TODO: Ensure directory exists
	// os.???(path, 0755)

	// TODO: Config with fallback
	// if val, ok := os.???(key); ok {
	//     return val
	// }
	// return fallback

	// TODO: Atomic create (fail if exists)
	// f, err := os.OpenFile(path, os.O_CREATE|os.O_???|os.O_WRONLY, 0644)
}

// =============================================================================
// EXERCISE 15: COMMON MISTAKES
// =============================================================================

func PracticeMistakes() {
	// Mistake 1: What's wrong?
	// file, _ := os.Open("data.txt")
	// data, _ := io.ReadAll(file)
	// Answer: ???

	// Mistake 2: What's wrong?
	// file, _ := os.Open("data.txt")
	// defer file.Close()
	// if file == nil { ... }
	// Answer: ???

	// Mistake 3: What's wrong?
	// func LibFunc() {
	//     os.Exit(1)
	// }
	// Answer: ???

	// Mistake 4: What's wrong?
	// val := os.Getenv("KEY")
	// if val == "" { /* not set? */ }
	// Answer: ???

	// Mistake 5: What's wrong?
	// if _, err := os.Stat(path); os.IsNotExist(err) {
	//     os.Create(path)
	// }
	// Answer: ???
}

// =============================================================================
// SELF-TEST: QUICK RECALL
// =============================================================================

func SelfTest() {
	// FILE OPERATIONS
	// 1. Read entire file?
	_ = "os.???(name)"

	// 2. Write entire file?
	_ = "os.???(name, data, perm)"

	// 3. Open read-only?
	_ = "os.???(name)"

	// 4. Create/truncate?
	_ = "os.???(name)"

	// 5. Open with flags?
	_ = "os.???(name, flags, perm)"

	// DIRECTORIES
	// 6. Make single dir?
	_ = "os.???(name, perm)"

	// 7. Make dir tree?
	_ = "os.???(path, perm)"

	// 8. Remove file?
	_ = "os.???(name)"

	// 9. Remove tree?
	_ = "os.???(path)"

	// 10. List dir?
	_ = "os.???(name)"

	// FILE INFO
	// 11. Stat file?
	_ = "os.???(name)"

	// 12. Stat without symlink follow?
	_ = "os.???(name)"

	// ENVIRONMENT
	// 13. Get env?
	_ = "os.???(key)"

	// 14. Get env with exists?
	_ = "os.???(key)"

	// 15. Set env?
	_ = "os.???(key, val)"

	// TEMP
	// 16. Temp file?
	_ = "os.???(dir, pattern)"

	// 17. Temp dir?
	_ = "os.???(dir, pattern)"

	// PROCESS
	// 18. Current dir?
	_ = "os.???()"

	// 19. Home dir?
	_ = "os.???()"

	// 20. Exit?
	_ = "os.???(code)"
}

// =============================================================================
// MINI PROJECT: FILE MANAGER
// =============================================================================

func MiniProject() {
	// Build a simple file manager that:
	//
	// 1. Lists directory contents with size and permissions
	// 2. Creates directories recursively
	// 3. Copies files safely
	// 4. Moves/renames files
	// 5. Handles errors gracefully
	//
	// Scaffold:
	//
	// type FileManager struct {
	//     root string
	// }
	//
	// func NewFileManager(root string) (*FileManager, error) {
	//     abs, err := filepath.Abs(root)
	//     if err != nil {
	//         return nil, err
	//     }
	//     return &FileManager{root: abs}, nil
	// }
	//
	// func (fm *FileManager) List(subdir string) ([]FileInfo, error) {
	//     path := filepath.Join(fm.root, subdir)
	//     entries, err := os.ReadDir(path)
	//     if err != nil {
	//         return nil, err
	//     }
	//
	//     var files []FileInfo
	//     for _, e := range entries {
	//         info, _ := e.Info()
	//         files = append(files, FileInfo{
	//             Name:  e.Name(),
	//             IsDir: e.IsDir(),
	//             Size:  info.Size(),
	//             Mode:  info.Mode(),
	//         })
	//     }
	//     return files, nil
	// }
	//
	// func (fm *FileManager) MakeDir(subdir string) error {
	//     path := filepath.Join(fm.root, subdir)
	//     return os.MkdirAll(path, 0755)
	// }
	//
	// func (fm *FileManager) Copy(src, dst string) error {
	//     srcPath := filepath.Join(fm.root, src)
	//     dstPath := filepath.Join(fm.root, dst)
	//
	//     // Ensure destination directory exists
	//     if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
	//         return err
	//     }
	//
	//     srcFile, err := os.Open(srcPath)
	//     if err != nil {
	//         return err
	//     }
	//     defer srcFile.Close()
	//
	//     dstFile, err := os.Create(dstPath)
	//     if err != nil {
	//         return err
	//     }
	//     defer dstFile.Close()
	//
	//     _, err = io.Copy(dstFile, srcFile)
	//     return err
	// }
	//
	// func (fm *FileManager) Move(src, dst string) error {
	//     srcPath := filepath.Join(fm.root, src)
	//     dstPath := filepath.Join(fm.root, dst)
	//     return os.Rename(srcPath, dstPath)
	// }
	//
	// func (fm *FileManager) Delete(path string) error {
	//     fullPath := filepath.Join(fm.root, path)
	//     info, err := os.Stat(fullPath)
	//     if err != nil {
	//         return err
	//     }
	//     if info.IsDir() {
	//         return os.RemoveAll(fullPath)
	//     }
	//     return os.Remove(fullPath)
	// }
}

// AllPractice is not meant to be called - this file is for reading.
func AllPractice() {
	PracticeStandardHandles()
	PracticeReadingFiles()
	PracticeWritingFiles()
	PracticeFlags()
	PracticePermissions()
	PracticeFileInfo()
	PracticeDirectories()
	PracticeWorkingDirectory()
	PracticeTempFiles()
	PracticeEnvironment()
	PracticeProcess()
	PracticeSignals()
	PracticeErrors()
	PracticePatterns()
	PracticeMistakes()
	SelfTest()
	MiniProject()
}
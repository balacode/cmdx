// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-25 01:19:16 6FAAD7                         [cmdx/env_provider.go]
// -----------------------------------------------------------------------------

package main

import "fmt"       // standard
import "os"        // standard
import "io/ioutil" // standard

import "github.com/balacode/zr"    // Zirconium
import "github.com/balacode/zr_fs" // Zirconium

var env EnvProvider = Env{}

// PL is env.Println() but is used only for debugging.
var PL = env.Println

// -----------------------------------------------------------------------------
// # Interface

// EnvProvider __
type EnvProvider interface {

	// -------------------------------------------------------------------------
	// # Console Output

	// Print __
	Print(a ...interface{}) (n int, err error)

	// Printf __
	Printf(format string, a ...interface{}) (n int, err error)

	// Println __
	Println(a ...interface{}) (n int, err error)

	// -------------------------------------------------------------------------
	// # Error Logging

	// Error outputs an error message to the standard output and to a
	// log file named 'run.log' saved in the program's current directory,
	// It also outputs the call stack (names and line numbers of callers.)
	// Returns an error value initialized with the message.
	Error(args ...interface{}) error

	// -------------------------------------------------------------------------
	// # File Operations

	// DeleteFile deletes 'filename' and returns true if it no longer exists.
	DeleteFile(filename string) bool

	// ReadFile reads 'filename' from the disk and returns the entire
	// file contents in 'data' and true in 'done' if successful.
	// Otherwise returns an empty byte array and false.
	ReadFile(filename string) (data []byte, done bool)

	// ReadFileChunks reads a file in chunks and repeatedly calls the
	// supplied 'reader' function with each read chunk.
	// This function is useful for processing large files to
	// avoid  having to load the entire file into memory.
	//
	// filename:  Name of the file to read from.
	//
	// chunkSize: Size of chunks to read from the file, in bytes.
	//
	// reader:    Function to call with each read chunk, which is passed
	//            in the 'chunk' parameter. The function should return
	//            int64(len(chunk)) to continue reading from the next
	//            position, or 0 if further reading should stop.
	//            It also accepts a negative value to skip
	//            the reading position back.
	//
	// Returns an error if the file can't be opened or a read fails.
	ReadFileChunks(
		filename string,
		chunkSize int64,
		reader func(chunk []byte) int64,
	) error

	// ReadFileLines __
	ReadFileLines(filename string) []string

	// RenameFile __
	RenameFile(oldpath, newpath string) bool

	// WriteFile __
	WriteFile(filename string, data []byte) bool

	// WriteFileLines __
	WriteFileLines(filename string, lines []string) bool

	// # File System Info

	// FileExists returns true if the file given by 'path' exists.
	FileExists(path string) bool

	// GetFilePaths returns a list of file names (with full path) contained
	// in folder 'dir' that match the given file extensions.
	// Extensions should be specified as: "ext", or ".ext", not "*.ext"
	GetFilePaths(dir string, exts ...string) []string

	// Getwd returns the current working directory.
	Getwd() string

	// PathSeparator __
	PathSeparator() string

	// TextFileExts returns a list of all text file extensions.
	TextFileExts() []string
} //                                                                 EnvProvider

// Env __
type Env struct{}

// -----------------------------------------------------------------------------
// # Console Output

// Print __
func (ob Env) Print(a ...interface{}) (n int, err error) {
	return fmt.Print(a...)
}

// Printf __
func (ob Env) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

// Println __
func (ob Env) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

// -----------------------------------------------------------------------------
// # Error Logging

// Error outputs an error message to the standard output and to a
// log file named 'run.log' saved in the program's current directory,
// It also outputs the call stack (names and line numbers of callers.)
// Returns an error value initialized with the message.
func (ob Env) Error(args ...interface{}) error {
	return zr.Error(args...)
}

// -----------------------------------------------------------------------------
// # File Operations

// DeleteFile deletes 'filename' and returns true if it no longer exists.
func (ob Env) DeleteFile(filename string) bool {
	var err = os.Remove(filename)
	if err != nil {
		ob.Println("Failed deleting", filename, "due to:", err)
		return false
	}
	return true
}

// ReadFile reads 'filename' from the disk and returns the entire
// file contents in 'data' and true in 'done' if successful.
// Otherwise returns an empty byte array and false.
func (ob Env) ReadFile(filename string) (data []byte, done bool) {
	var err error
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		ob.Println("Failed reading", filename, "due to:", err)
		return []byte{}, false
	}
	return data, true
}

// ReadFileChunks reads a file in chunks and repeatedly calls the
// supplied 'reader' function with each read chunk.
// This function is useful for processing large files to
// avoid  having to load the entire file into memory.
//
// filename:  Name of the file to read from.
//
// chunkSize: Size of chunks to read from the file, in bytes.
//
// reader:    Function to call with each read chunk, which is passed
//            in the 'chunk' parameter. The function should return
//            int64(len(chunk)) to continue reading from the next
//            position, or 0 if further reading should stop.
//            It also accepts a negative value to skip
//            the reading position back.
//
// Returns an error if the file can't be opened or a read fails.
func (ob Env) ReadFileChunks(
	filename string,
	chunkSize int64,
	reader func(chunk []byte) int64,
) error {
	return fs.ReadFileChunks(filename, chunkSize, reader)
}

// ReadFileLines __
func (ob Env) ReadFileLines(filename string) []string {
	//TODO: add error return value to fs.ReadFileLines, add bool return here
	return fs.ReadFileLines(filename)
}

// RenameFile __
func (ob Env) RenameFile(oldpath, newpath string) bool {
	var err = os.Rename(oldpath, newpath)
	if err != nil {
		ob.Println("Failed renaming", oldpath, " to ", newpath, "due to:", err)
		return false
	}
	return true
}

// WriteFile __
func (ob Env) WriteFile(filename string, data []byte) bool {
	var err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		ob.Println("Failed writing", filename, "due to:", err)
		return false
	}
	return true
}

// WriteFileLines __
func (ob Env) WriteFileLines(filename string, lines []string) bool {
	var err = fs.WriteFileLines(filename, lines)
	if err != nil {
		ob.Println("Failed writing", filename, "due to:", err)
		return false
	}
	return true
}

// -----------------------------------------------------------------------------
// # File System Info

// FileExists returns true if the file given by 'path' exists.
func (ob Env) FileExists(path string) bool {
	return fs.FileExists(path)
}

// GetFilePaths returns a list of file names (with full path) contained
// in folder 'dir' that match the given file extensions.
// Extensions should be specified as: "ext", or ".ext", not "*.ext"
func (ob Env) GetFilePaths(dir string, exts ...string) []string {
	return fs.GetFilePaths(dir, exts...)
}

// Getwd returns the current working directory.
func (ob Env) Getwd() string {
	var dir, err = os.Getwd()
	if err != nil {
		ob.Println("Failed to determine working directory due to:", err)
		return ""
	}
	return dir
}

// PathSeparator __
func (ob Env) PathSeparator() string {
	return string(os.PathSeparator)
}

// NewDirWatcher __
func (ob Env) NewDirWatcher(dir string) *fs.DirWatcher {
	return fs.NewDirWatcher(dir)
}

// TextFileExts returns a list of all text file extensions.
func (ob Env) TextFileExts() []string {
	return fs.TextFileExts
}

//TODO: env_provider.go: add the following:
// fs.IsTextFile
// fs.WalkPath
// os.Args
// os.Exit
// os.File
// os.FileInfo, err error) error
// os.Open(filename)
// os.Stat(filename)

//end

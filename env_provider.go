// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                   cmdx/[env_provider.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"os"

	"github.com/balacode/zr"
	fs "github.com/balacode/zr-fs"
)

// Env defines functions that interact with the app's environment.
// All IO is channelled through calls on the 'env' variable.
type Env struct{}

var (
	// env is the one instance of Env
	env EnvProvider = Env{}

	// PL is env.Println() but is used only for debugging.
	PL = env.Println
)

// -----------------------------------------------------------------------------
// # Interface

// EnvProvider defines the app's environment provider interface,
// i.e. all the input and output functions
type EnvProvider interface {

	// -------------------------------------------------------------------------
	// # Console Output

	// Print _ _
	Print(a ...interface{}) (n int, err error)

	// Printf _ _
	Printf(format string, a ...interface{}) (n int, err error)

	// Println _ _
	Println(a ...interface{}) (n int, err error)

	// -------------------------------------------------------------------------
	// # Error Logging

	// Error outputs an error message to the standard output and to a
	// log file named "<process>.log" in the program's current directory,
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

	// ReadFileLines _ _
	ReadFileLines(filename string) []string

	// RenameFile _ _
	RenameFile(oldPath, newPath string) bool

	// WriteFile _ _
	WriteFile(filename string, data []byte) bool

	// WriteFileLines _ _
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

	// PathSeparator _ _
	PathSeparator() string

	// TextFileExts returns a list of all text file extensions.
	TextFileExts() []string
}

// -----------------------------------------------------------------------------
// # Console Output

// Print _ _
func (Env) Print(a ...interface{}) (n int, err error) {
	return fmt.Print(a...)
}

// Printf _ _
func (Env) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

// Println _ _
func (Env) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

// -----------------------------------------------------------------------------
// # Error Logging

// Error outputs an error message to the standard output and to a
// log file named "<process>.log" in the program's current directory,
// It also outputs the call stack (names and line numbers of callers.)
// Returns an error value initialized with the message.
func (Env) Error(args ...interface{}) error {
	return zr.Error(args...)
}

// -----------------------------------------------------------------------------
// # File Operations

// DeleteFile deletes 'filename' and returns true if it no longer exists.
func (Env) DeleteFile(filename string) bool {
	err := os.Remove(filename)
	if err != nil {
		env.Println("Failed deleting", filename, "due to:", err)
		return false
	}
	return true
}

// ReadFile reads 'filename' from the disk and returns the entire
// file contents in 'data' and true in 'done' if successful.
// Otherwise returns an empty byte array and false.
func (Env) ReadFile(filename string) (data []byte, done bool) {
	var err error
	data, err = os.ReadFile(filename)
	if err != nil {
		env.Println("Failed reading", filename, "due to:", err)
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
func (Env) ReadFileChunks(
	filename string,
	chunkSize int64,
	reader func(chunk []byte) int64,
) error {
	return fs.ReadFileChunks(filename, chunkSize, reader)
}

// ReadFileLines _ _
func (Env) ReadFileLines(filename string) []string {
	// TODO: add error return value to fs.ReadFileLines, add bool return here
	return fs.ReadFileLines(filename)
}

// RenameFile _ _
func (Env) RenameFile(oldPath, newPath string) bool {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		env.Println("Failed renaming", oldPath, " to ", newPath, "due to:", err)
		return false
	}
	return true
}

// WriteFile _ _
func (Env) WriteFile(filename string, data []byte) bool {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		env.Println("Failed writing", filename, "due to:", err)
		return false
	}
	return true
}

// WriteFileLines _ _
func (Env) WriteFileLines(filename string, lines []string) bool {
	err := fs.WriteFileLines(filename, lines)
	if err != nil {
		env.Println("Failed writing", filename, "due to:", err)
		return false
	}
	return true
}

// -----------------------------------------------------------------------------
// # File System Info

// FileExists returns true if the file given by 'path' exists.
func (Env) FileExists(path string) bool {
	return fs.FileExists(path)
}

// GetFilePaths returns a list of file names (with full path) contained
// in folder 'dir' that match the given file extensions.
// Extensions should be specified as: "ext", or ".ext", not "*.ext"
func (Env) GetFilePaths(dir string, exts ...string) []string {
	return fs.GetFilePaths(dir, exts...)
}

// Getwd returns the current working directory.
func (Env) Getwd() string {
	dir, err := os.Getwd()
	if err != nil {
		env.Println("Failed to determine working directory due to:", err)
		return ""
	}
	return dir
}

// PathSeparator _ _
func (Env) PathSeparator() string {
	return string(os.PathSeparator)
}

// TextFileExts returns a list of all text file extensions.
func (Env) TextFileExts() []string {
	return fs.TextFileExts
}

// TODO: env_provider.go: add the following:
// fs.FileExists
// fs.GetFilePaths
// fs.IsFileExt
// fs.IsTextFile
// fs.ReadFileChunks
// fs.ReadFileLines
// fs.TextFileExts
// fs.WalkPath
// fs.WriteFileLines
// os.Args
// os.Exit
// os.File
// os.FileInfo, err error) error
// os.Open(filename)
// os.Stat(filename)

// end

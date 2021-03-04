// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                       cmdx/[settings.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

// -----------------------------------------------------------------------------
// # Paths and Files

// DefaultExts specifies file types that CMDX processes.
// By default this compreises of all known text file extensions.
var DefaultExts = env.TextFileExts()

// DefaultLibPath specifies the default path of the Zircon-Go library.
var DefaultLibPath = func() string {
	if env.PathSeparator() == "/" {
		return hardcodedDefaultLibPathOnLinux
	}
	return hardcodedDefaultLibPathOnWindows
}()

// DefaultPath _ _
var DefaultPath = "."

// RootPath _ _
var RootPath = hardcodedRootPath

// TimeLogPaths _ _
var TimeLogPaths = hardcodedTimeLogPaths

// IgnoreFilenamesWith specifies file names ignored
// by such functions as markTimeInFiles().
// I.e. if any part of a file's path contains one of these
// substrings, the file will not be processed.
var IgnoreFilenamesWith = hardcodedIgnoreFilenamesWith

// WordListFile _ _
var WordListFile = hardcodedWordListFile

// end

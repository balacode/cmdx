// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-26 14:45:21 28B927                             cmdx/[settings.go]
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

// DefaultPath __
var DefaultPath = "."

// RootPath __
var RootPath = hardcodedRootPath

// TimeLogPaths __
var TimeLogPaths = hardcodedTimeLogPaths

// IgnoreFilenamesWith specifies file names ignored by such
// functions as markTimeInFiles() and runInteractive().
// I.e. if any part of a file's path contains one of these
// substrings, the file will not be processed.
var IgnoreFilenamesWith = hardcodedIgnoreFilenamesWith

// WordListFile __
var WordListFile = hardcodedWordListFile

//end

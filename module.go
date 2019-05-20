// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-20 11:10:53 446DB9                               cmdx/[module.go]
// -----------------------------------------------------------------------------

// See TODO at the end

// Package cmdx contains the cmdx (cx) command which amalgamates
// various useful command-line utilities in one executable.
package main

import (
	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # Constants: Source Markers

const (
	// CommandMark __
	CommandMark = "/" + "/" + "/"

	// DefaultMark __
	DefaultMark = "~" + "~"

	// ErrorMark __
	ErrorMark = "//" + "^ "

	// ErrorEndMark __
	ErrorEndMark = " " + "`" + "`" + "`" + "`"

	// OldMark __
	OldMark = "OLD"
)

// -----------------------------------------------------------------------------
// # Limiting Constants: these constants don't need to be changed normally.

const (
	// ColumnsNorm is the expected number of columns in source files.
	// The 'copy long lines' command (CLL) lists all lines longer than this.
	ColumnsNorm = 80

	// FileChunkSize specifies the size of __
	FileChunkSize = 2 * 1024 * 1024 // 2 MB chunk

	// LongestLine ignore lines exceeding this # of columns.
	LongestLine = 2048

	// LongestWord __
	LongestWord = 256

	// RecentFiles __
	RecentFiles = 10

	// ShownResultsLimit __
	ShownResultsLimit = 50
)

// -----------------------------------------------------------------------------
// # Debugging and Related Constants:

const (
	// DebugReplaceLines global flag specifies if arguments and return
	// value of ReplaceLines() should be sent to the console.
	DebugReplaceLines = false

	// ShowProgressIndicator __
	ShowProgressIndicator = false
)

// VL is zr.VerboseLog() but is used only for debugging.
var VL = zr.VerboseLog

// -----------------------------------------------------------------------------
// # Constants: Other

const (
	// HeaderSignatureRX __
	HeaderSignatureRX = `:v: \d{4}-\d{2}-\d{2}` +
		` \d{2}:\d{2}:\d{2} [0-9A-Fa-f]{6}`

	// HeaderTimePos __
	HeaderTimePos = 4
)

// TODO: Create command to convert all literal strings to $string constants.
//       Requires a list of input JS files to modify, and a
//       constants.js file where to write the constants.
//
// TODO: JS Unconstifier: replaces all $string constants with normal strings.
//
// TODO: in replaceLinesInFiles():
//       - Add number of files changed
//       - Check if FIND and REPL already exists, then no need to add it.
//       - Remember the indentation of each line.
//       - Store line number in _REPL_LINES.TXT and use it as an ID.
//       - Write time, file and numbers of replacements made to logfile.
//       - Add an option to tag changed lines.
//
// TODO: Add a command to harvest lines to change from source.
//
//       F - find
//       R - replace
//       I - ignore case
//       W - whole word
//       B - block mode
//       M - multiple line replacement

//end

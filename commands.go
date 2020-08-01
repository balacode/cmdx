// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2020-08-01 22:31:18 A840C7                             cmdx/[commands.go]
// -----------------------------------------------------------------------------

package main

// Command _ _
type Command struct {
	ShortName string
	FullName  string
	ShortInfo string
	FullInfo  string
	Category  int
	Handler   func(cmd Command, args []string)
} //                                                                     Command

// Interactive Commands:

// blocks
const (
	// BB is the Begin Block command.
	BB = "BB"

	// BC is the Blocks Copy command. (same as CB)
	BC = "BC"

	// BE is the Block End command (same as EB)
	BE = "BE"

	// CB is the Copy Blocks command (same as BC)
	CB = "CB"

	// CLL copies all extra-long lines in source
	CLL = "CLL"

	// EB is the End Block (same as BE) command
	EB = "EB"

	// FF is the Find in Files command
	FF = "FF"

	// ID is the ID (insert) command
	ID = "ID"

	// UUID is the UUID (insert) command
	UUID = "UUID"

	// LT is the Less-Than sign (command applies to same line)
	LT = "<"

	// T timestamp (insert)
	T = "T"

	// XE is the Remove Error markers
	XE = "XE"
)

// AllCategories _ _
var AllCategories = map[int]string{
	1: "File Manipulation",
	2: "Text Manipulation",
	3: "Other",
}

// AllCommands _ _
// The commands are grouped by category
var AllCommands = []Command{
	// File Manipulation:
	{
		ShortName: "dd",
		FullName:  "del-dup",
		ShortInfo: "Delete identical files: read <source>, delete in <target>",
		Handler:   deleteIdenticalFiles,
		Category:  1,
	},
	{
		ShortName: "ld",
		FullName:  "list-dup",
		ShortInfo: "List identical files: read <source>, list in <target>",
		Handler:   listIdenticalFiles,
		Category:  1,
	},
	{
		ShortName: "mg",
		FullName:  "merge",
		ShortInfo: "Merge files in 'source=' into 'target=' (use merge.rgon)",
		Handler:   mergeFiles,
		Category:  1,
	},
	{
		ShortName: "rd",
		FullName:  "ren-dup",
		ShortInfo: "Rename identical files: read <source>, rename in <target>",
		Handler:   renameIdenticalFiles,
		Category:  1,
	},
	{
		ShortName: "rh",
		FullName:  "ren-hash",
		ShortInfo: "Renames files by prefixing their name with a hash",
		Handler:   renameFilesByHash,
		Category:  1,
	},
	// Text Manipulation:
	{
		ShortName: "fw",
		FullName:  "file-words",
		ShortInfo: "Lists all words with alphanumeric characters from <file>",
		Handler:   wordsInFile,
		Category:  2,
	},
	{
		ShortName: "me",
		FullName:  "mark-errors",
		ShortInfo: "Insert build errors as comments in source code files",
		Handler:   markErrorsInSource, // uses 'flag' module to read arguments
		Category:  2,
	},
	{
		ShortName: "mt",
		FullName:  "mark-time",
		ShortInfo: "Change timestamps in source files. Requires path",
		Handler:   markTimeInFiles,
		Category:  2,
	},
	{
		ShortName: "pb",
		FullName:  "part-backup",
		ShortInfo: "Continously back-up .part files streamed in current folder",
		Handler:   partBackup,
		Category:  2,
	},
	{
		ShortName: "rl",
		FullName:  "rep-lines",
		ShortInfo: "Replace lines in file(s). Requires <command-file>",
		Handler:   replaceLinesInFiles,
		Category:  2,
	},
	{
		ShortName: "rs",
		FullName:  "replace-strings",
		ShortInfo: "Replace strings in file(s). Requires <command-file>",
		Handler:   replaceStringsInFiles,
		Category:  2,
	},
	{
		ShortName: "rt",
		FullName:  "rep-time",
		ShortInfo: "Replace time entries. requires <source-file> <target-file>",
		Handler:   replaceTime,
		Category:  2,
	},
	{
		ShortName: "sf",
		FullName:  "sort-file",
		ShortInfo: "Sorts all the lines in a file",
		Handler:   sortFileLines,
		Category:  2,
	},
	// Other:
	{
		ShortName: "ri",
		FullName:  "run",
		ShortInfo: "Runs in interactive mode",
		Handler:   runInteractive,
		Category:  3,
	},
	{
		ShortName: "mw",
		FullName:  "match-words",
		ShortInfo: "Lists all words that use all the specified letters.",
		Handler:   matchWords,
		Category:  3,
	},
	{
		ShortName: "tr",
		FullName:  "time-report",
		ShortInfo: "Summarizes time from log files",
		Handler:   timeReport,
		Category:  3,
	},
	{
		ShortName: "now",
		FullName:  "time-now",
		ShortInfo: "Displays current date/time in yyyy-mm-dd hh:nn:ss format",
		Handler:   timeNow,
		Category:  3,
	},
} //                                                                 AllCommands

//end

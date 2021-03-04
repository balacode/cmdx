// -----------------------------------------------------------------------------
// CMDX Utilities Suite                            cmdx/[sort_lines_in_files.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	"sort"
	"strings"
)

// sortFileLines sorts all the lines in the specified file, removing
// non-unique lines. This command can be used to keep log files sorted.
func sortFileLines(cmd Command, args []string) {
	if len(args) != 1 {
		env.Println("requires <file-name> parameter")
		return
	}
	// read the file
	var (
		filename   = args[0]
		lines      = env.ReadFileLines(filename)
		oldContent = strings.Join(lines, "\n")
	)
	// remove non-unique lines
	if true {
		unique := make(map[string]bool, len(lines))
		for _, line := range lines {
			unique[line] = true
		}
		lines = make([]string, 0, len(unique))
		for key := range unique {
			lines = append(lines, key)
		}
	}
	// sort the lines
	sort.Strings(lines)
	//
	// don't save if nothing changed
	if strings.Join(lines, "\n") == oldContent {
		return
	}
	// save the file
	env.WriteFileLines(filename, lines)
} //                                                               sortFileLines

// end

// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-03-18 01:07:59 173DB1                      cmdx/[sort_file_lines.go]
// -----------------------------------------------------------------------------

package main

import (
	"sort"
	str "strings"
)

// sortFileLines sorts all the lines in the specified file, removing
// non-unique lines. This command can be used to keep log files sorted.
func sortFileLines(cmd Command, args []string) {
	if len(args) != 1 {
		env.Println("requires <file-name> parameter")
		return
	}
	// read the file
	filename := args[0]
	lines := env.ReadFileLines(filename)
	oldContent := str.Join(lines, "\n")
	//
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
	if str.Join(lines, "\n") == oldContent {
		return
	}
	// save the file
	env.WriteFileLines(filename, lines)
} //                                                               sortFileLines

//end

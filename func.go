// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-09 18:06:19 50D97F                                 cmdx/[func.go]
// -----------------------------------------------------------------------------

package main

//   filterLongLines(
//       lines []string,
//       maxLineLength int,
//   ) (ret []string)
//   getFilesMap(dir, filter string) FilesMap
//   sortUniqueStrings(a []string) []string
//   splitArgsFilter(args []string) (retArgs []string, filter string)
//   trim(s string) string

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// filterLongLines __
func filterLongLines(
	lines []string,
	longerThan int,
) (
	ret []string,
) {
	for i, s := range lines {
		if strings.Contains(s, "\t") {
			s = strings.Replace(s, "\t", "    ", -1)
		}
		n := len(s)
		if n > longerThan && n < LongestLine {
			ret = append(ret, lines[i])
		}
	}
	return ret
} //                                                             filterLongLines

// getFilesMap __
func getFilesMap(dir, filter string) FilesMap {
	filter = strings.ToLower(filter)
	ret := make(FilesMap, 1000)
	// TODO: use fs.WalkPath() instead of this; then remove "os" dependency
	filepath.Walk(
		dir, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, "$RECYCLE.BIN") {
				return nil
			}
			if err != nil {
				env.Printf("in path %s: %s\n", path, err)
				return nil
			}
			if info.IsDir() {
				return nil
			}
			if strings.Index(strings.ToLower(path), filter) == -1 {
				return nil
			}
			size := info.Size()
			ret[size] = append(ret[size], &PathAndSize{Path: path, Size: size})
			return nil
		},
	)
	return ret
} //                                                                 getFilesMap

// sortUniqueStrings sorts string array 'a' and removes any repeated values
func sortUniqueStrings(a []string) []string {
	unique := make(map[string]bool, len(a))
	for _, s := range a {
		unique[s] = true
	}
	ret := make([]string, 0, len(unique))
	for s := range unique {
		ret = append(ret, s)
	}
	// sort the lines
	sort.Strings(ret)
	return ret
} //                                                           sortUniqueStrings

// splitArgsFilter extracts '-filter expr' or '--filter expr' from args,
// and returns args with the option removed, and the extracted filter value.
func splitArgsFilter(args []string) (retArgs []string, filter string) {
	//
	endArg := len(args) - 1
	for i := 0; i <= endArg; i++ {
		arg := strings.ToLower(args[i])
		if arg == "-filter" || arg == "--filter" {
			if i == endArg {
				env.Println(arg + " is missing its value")
				return args, ""
			}
			filter = args[i+1]
			args = args[:i+copy(args[i:], args[i+2:])]
			endArg = len(args) - 1
			i--
		}
	}
	return args, filter
} //                                                             splitArgsFilter

// trim removes all leading and trailing white-spaces from a string
func trim(s string) string {
	return strings.TrimSpace(s)
} //                                                                        trim

//end

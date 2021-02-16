// -----------------------------------------------------------------------------
// CMDX Utility                                                   cmdx/[func.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

//   checksum(s string) string
//   filterLongLines(
//       lines []string,
//       maxLineLength int,
//   ) (ret []string)
//   getFilesMap(dir, filter string) FilesMap
//   isHelpRequested(args []string)
//   parseTime(value interface{}) time.Time
//   sortUniqueStrings(a []string) []string
//   splitArgsFilter(args []string) (retArgs []string, filter string)
//   trim(s string) string

import (
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/balacode/zr"
)

// checksum returns a shortened CRC32 checksum of the given string.
// The returned checksum is a string made up of 6 hexadecimal digits,
// shorter than the 8 hex digits required for a normal CRC32.
func checksum(s string) string {
	chk := crc32.ChecksumIEEE([]byte(s))
	chk = (chk / 0x00FFFFFF) ^ (chk & 0x00FFFFFF) // <- from 4 to 3 bytes
	return fmt.Sprintf("%06X", chk)
} //                                                                    checksum

// filterLongLines _ _
func filterLongLines(
	lines []string,
	longerThan int,
) (
	ret []string,
) {
	for i, s := range lines {
		if strings.Contains(s, "\t") {
			s = strings.ReplaceAll(s, "\t", "    ")
		}
		n := len(s)
		if n > longerThan && n < LongestLine {
			ret = append(ret, lines[i])
		}
	}
	return ret
} //                                                             filterLongLines

// getFilesMap _ _
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

// isHelpRequested returns true if args contains a help request.
// That is '?', 'h', 'hlp', or 'help' (or '/help', '-help', etc.)
func isHelpRequested(args []string) bool {
	for _, arg := range args {
		arg = strings.ToLower(strings.Trim(arg, "-/\\"))
		if arg == "?" || arg == "h" || arg == "hlp" || arg == "help" {
			return true
		}
	}
	return false
} //                                                             isHelpRequested

// parseTime converts any string-like value to time.Time without returning
// an error if the conversion failed, in which case it logs an error
// and returns a zero-value time.Time.
//
// If val is a zero-length string, returns a zero-value time.Time
// but does not log a warning.
//
// It also accepts a time.Time as input.
//
// In both cases the returned Time type will contain only the date
// part without the time or time zone components.
//
// Note: fmt.Stringer (or fmt.GoStringer) interfaces are not treated as
// strings to avoid bugs from implicit conversion. Use the String method.
//
func parseTime(value interface{}) time.Time {
	switch v := value.(type) {
	case time.Time:
		{
			return v
		}
	case string:
		{
			if v == "" {
				return time.Time{}
			}
			var tm time.Time
			var err error
			if len(v) == 10 {
				tm, err = time.Parse("2006-01-02", v)
				if err == nil && !tm.IsZero() {
					return parseTime(tm)
				}
			}
			if len(v) == 19 {
				tm, err = time.Parse("2006-01-02 15:04:05", v)
				if err == nil && !tm.IsZero() {
					return parseTime(tm)
				}
			}
			if err != nil {
				zr.Error(err)
			}
			return time.Time{}
		}
	case *string:
		if v != nil {
			return parseTime(*v)
		}
	}
	zr.Error("Can not convert", reflect.TypeOf(value), "to int:", value)
	return time.Time{}
} //                                                                   parseTime

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

// end

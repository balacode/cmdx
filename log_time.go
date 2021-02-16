// -----------------------------------------------------------------------------
// CMDX Utility                                               cmdx/[log_time.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

// # Command Function
//   logTime(cmd Command, args []string)

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/balacode/cmdx/cxfunc"
	"github.com/balacode/zr"
	fs "github.com/balacode/zr-fs"
)

// AutotimeFilename _ _
const AutotimeFilename = "autotime.log"

// -----------------------------------------------------------------------------
// # Command Function

// logTime _ _
func logTime(cmd Command, args []string) {
	var (
		backlogArg string
		verboseArg = false
		now        = time.Now()
		today      = now.Format("2006-01-02")
		logFiles   = ltListAutotimeFiles()
		logEntries = ltGetLogEntries(logFiles)
		changes    = map[string]string{} // key:path value:modTime
		backlog    time.Duration
	)
	// read arguments
	{
		f := flag.NewFlagSet("", flag.ExitOnError)
		backlog := f.String("backlog", "today", "")
		f.Parse(args)
		backlogArg = *backlog
	}
	if backlogArg != "today" {
		var err error
		backlog, err = cxfunc.ParseDuration(backlogArg)
		if err != nil {
			zr.Error(zr.EInvalidArg, "^backlog", ":^", backlog)
			return
		}
	}
	// detect modified files in the current folder and its subfolders
	ltProcessTextFilesInCurrentFolder(func(path, modTime string) {
		if backlogArg == "today" && modTime < today {
			if verboseArg {
				text := modTime + " SKIP: " + path
				fmt.Println(text)
			}
			return
		}
		if backlog != 0 {
			tm := parseTime(modTime)
			diff := now.Sub(tm)
			if diff > backlog {
				return
			}
		}
		hasEntry := logEntries[path][modTime]
		if hasEntry {
			return
		}
		changes[path] = modTime
	})
	// write changed timestamps and paths to the nearest ancestor log file
	var prev string
	for path, modTime := range changes {
		text := modTime + " " + path
		logFile := ltGetAutotimeFile(path)
		zr.AppendToTextFile(logFile, text+"\n")
		if prev != logFile {
			fmt.Println("\nLOGFILE ----------> " + logFile + ":")
			prev = logFile
		}
		fmt.Println(text)
	}
	fmt.Println()
} //                                                                     logTime

// -----------------------------------------------------------------------------
// # Internal Functions

// ltGetAutotimeFile _ _
func ltGetAutotimeFile(path string) string {
	var (
		ret   = ""
		dir   = filepath.Dir(path)
		sep   = string(os.PathSeparator)
		parts = strings.Split(dir, sep)
	)
	for len(parts) > 1 {
		logFile := strings.Join(parts, sep) + sep + AutotimeFilename
		if fs.FileExists(logFile) {
			ret = logFile
			break
		}
		parts = parts[:len(parts)-1]
	}
	if ret == "" {
		ret = dir + sep + AutotimeFilename
	}
	return ret
} //                                                           ltGetAutotimeFile

// ltGetLogEntries _ _
func ltGetLogEntries(logFiles []string) map[string]map[string]bool {
	ret := map[string]map[string]bool{}
	for _, path := range logFiles {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("ERROR: ", err)
			continue
		}
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			var date, filename string
			if len(line) >= 19 && zr.IsDate(line[:19]) {
				date, filename = line[:19], line[20:] // YYYY-MM-DD hh:mm:ss
			} else if len(line) >= 16 && zr.IsDate(line[:16]) {
				date, filename = line[:16], line[17:] // YYYY-MM-DD hh:mm
			} else {
				continue
			}
			filename = strings.TrimSpace(filename)
			if ret[filename] == nil {
				ret[filename] = map[string]bool{}
			}
			ret[filename][date] = true
		}
	}
	return ret
} //                                                             ltGetLogEntries

// ltListAutotimeFiles _ _
func ltListAutotimeFiles() []string {
	currentDir, err := os.Getwd()
	if err != nil {
		zr.Error(err)
		return nil
	}
	ret := []string{}
	sep := string(os.PathSeparator)
	//
	// list all time log files in parent folders of the current folder
	parentPaths := strings.Split(currentDir, sep)
	for len(parentPaths) > 1 {
		parentPaths = parentPaths[:len(parentPaths)-1]
		path := strings.Join(parentPaths, sep) + sep + AutotimeFilename
		if fs.FileExists(path) {
			ret = append(ret, path)
		}
	}
	// list all time log files in or under the current folder
	for _, path := range fs.GetFilePaths(currentDir, "log") {
		name := path
		if strings.Contains(name, sep) {
			ar := strings.Split(name, sep)
			name = ar[len(ar)-1]
		}
		if name == AutotimeFilename {
			ret = append(ret, path)
		}
	}
	// sort the list in descending order, with longest paths first
	sort.Slice(ret, func(i, j int) bool {
		a, b := ret[i], ret[j]
		return len(a) > len(b) || (len(a) == len(b) && a > b)
	})
	return ret
} //                                                         ltListAutotimeFiles

// ltProcessTextFilesInCurrentFolder _ _
func ltProcessTextFilesInCurrentFolder(
	processFile func(path string, modTime string),
) {
	currentFolder, err := os.Getwd()
	if err != nil {
		fmt.Println(zr.Error(err))
		return
	}
	sep := string(os.PathSeparator)
	err = filepath.Walk(
		currentFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return zr.Error("Error reading path", path, "due to:", err)
			}
			if info.IsDir() {
				return nil
			}
			// TODO: create a way for the user to specify files/paths to ignore
			if strings.HasSuffix(path, ".log") {
				return nil
			}
			if strings.Contains(path, sep+"node_modules"+sep) ||
				strings.Contains(path, "build"+sep+"intermediates"+sep) {
				return nil
			}
			if !fs.IsTextFile(path) {
				return nil
			}
			modTime := info.ModTime().Format("2006-01-02 15:04:05")
			processFile(path, modTime)
			return nil
		},
	)
	if err != nil {
		fmt.Println(zr.Error(err))
	}
} //                                           ltProcessTextFilesInCurrentFolder

// end

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
	"hash/crc32"
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
		backlogArg time.Duration
		repeatArg  time.Duration
		verboseArg bool
	)
	// read arguments
	{
		var (
			fl      = flag.NewFlagSet("", flag.ExitOnError)
			backlog = fl.String("backlog", "today", "")
			repeat  = fl.String("repeat", "disabled", "")
			verbose = fl.Bool("verbose", false, "")
		)
		fl.Parse(args)
		var err error
		if *backlog != "today" {
			backlogArg, err = cxfunc.ParseDuration(*backlog)
			if err != nil {
				zr.Error(zr.EInvalidArg, "^backlog", ":^", backlogArg)
				return
			}
		}
		if *repeat != "disabled" {
			repeatArg, err = cxfunc.ParseDuration(*repeat)
			if err != nil {
				zr.Error(zr.EInvalidArg, "^repeat", ":^", repeatArg)
				return
			}
			fmt.Println("\n" + "log-time repeat --> " + repeatArg.String())
		}
		verboseArg = *verbose
	}
	type Change struct {
		modTime  string
		checksum string
	}
	for {
		var (
			now        = time.Now()
			today      = now.Format("2006-01-02")
			logFiles   = ltListAutotimeFiles()
			logEntries = ltGetLogEntries(logFiles)
			changes    = map[string]Change{}
		)
		if repeatArg > 0 {
			timestamp := time.Now().Format("2006-01-02 15:04:05")
			fmt.Println("\n" + "log-time scan ----> " + timestamp)
		}
		// detect modified files in the current folder and its subfolders
		ltProcessTextFilesInCurrentFolder(func(path, modTime string) {
			if backlogArg == 0 && modTime < today {
				if verboseArg {
					text := modTime + " SKIP: " + path
					fmt.Println(text)
				}
				return
			}
			if backlogArg != 0 {
				tm := parseTime(modTime)
				diff := now.Sub(tm)
				if diff > backlogArg {
					return
				}
			}
			hasModTime := logEntries[path][modTime]
			if hasModTime {
				return
			}
			// calculate file's checksum
			content, err := ioutil.ReadFile(path)
			if err != nil {
				zr.Error(err)
				return
			}
			checksum := fmt.Sprintf("%08X", crc32.ChecksumIEEE(content))
			hasChecksum := logEntries[path][checksum]
			if hasChecksum {
				return
			}
			// if the change is unique, add the file to
			changes[path] = Change{modTime: modTime, checksum: checksum}
		})
		// write changed timestamps and paths to the nearest ancestor log file
		var prev string
		for path, it := range changes {
			text := it.modTime + " " + it.checksum + " " + path
			logFile := ltGetAutotimeFile(path)
			zr.AppendToTextFile(logFile, text+"\n")
			if prev != logFile {
				fmt.Println("\n" + "log-time file ----> " + logFile + ":")
				prev = logFile
			}
			fmt.Println(text)
		}
		// continue looping if '--repeat' has been specified, exit otherwise
		if repeatArg > 0 {
			timestamp := time.Now().Format("2006-01-02 15:04:05")
			fmt.Println("\n" + "log-time done ----> " + timestamp)
			time.Sleep(repeatArg)
			continue
		}
		break
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
			var date, checksum, filename string
			if len(line) >= 29 && zr.IsDate(line[:19]) {
				// YYYY-MM-DD hh:mm:ss 00000000 /PATH
				date, checksum, filename = line[:19], line[20:28], line[29:]
			} else if len(line) >= 26 && zr.IsDate(line[:16]) {
				// YYYY-MM-DD hh:mm 00000000 /PATH
				date, checksum, filename = line[:16], line[17:25], line[26:]
			} else {
				continue
			}
			filename = strings.TrimSpace(filename)
			if ret[filename] == nil {
				ret[filename] = map[string]bool{}
			}
			ret[filename][date] = true
			ret[filename][checksum] = true
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

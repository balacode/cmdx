// -----------------------------------------------------------------------------
// CMDX Utilities Suite                             cmdx/[mark_time_in_files.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

// # Command Handler
//   markTimeInFiles(cmd Command, args []string)
//
// # Support (File Scope)
//   autoTimeLog(path string, timestamp string)
//   getTimeLogPath(path string) string
//   processDir(dir string, changeTime bool)
//   processFile(path, name string, modTime time.Time) error
//   replaceVersion(s, path, filename string, modTime time.Time) string

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # Command Handler

// markTimeInFiles _ _
// The 'cmd' argument is not used.
func markTimeInFiles(cmd Command, args []string) {
	var (
		repeat     = false
		changeTime = true
	)
	for _, arg := range args {
		switch strings.ToLower(strings.Trim(arg, SPACES+"-/")) {
		case "repeat", "r":
			{
				repeat = true
			}
		case "hash-only":
			{
				changeTime = false
			}
		default:
			env.Printf("unknown argument '%s'\n", arg)
			return
		}
	}
	dir := env.Getwd()
	if dir == "" {
		fmt.Println("Failed to determine current folder")
		return
	}
	if repeat {
		fmt.Println("running 'mt' command in repeat mode (every 29 sec.)")
		ticker := time.NewTicker(time.Second * 29)
		defer ticker.Stop()
		stop := make(chan bool)
		go func() {
			// time.Sleep(24 * time.Hour)
			// stop <- true
		}()
		for {
			select {
			case <-ticker.C: // sends time which is not used
				{
					processDir(dir, changeTime)
				}
			case <-stop:
				return
			}
		}
	} else {
		processDir(dir, changeTime)
	}
} //                                                             markTimeInFiles

// -----------------------------------------------------------------------------
// # Support (File Scope)

// logs the file modification in autotime.log
// path: the fully-qualified file name
func autoTimeLog(path string, timestamp string) {
	//
	// show timestamp and file
	filename := filepath.Base(path)
	env.Printf("on %s %s\n", timestamp, filename)
	//
	//
	logPath := strings.ToLower(getTimeLogPath(path))
	//
	// the entry written in the log file:
	entry := strings.ToLower(timestamp + " " + path + "\n")
	if strings.Contains(entry, logPath) {
		entry = strings.ReplaceAll(entry, logPath, "")
	}
	entry = strings.ReplaceAll(entry, "\\", "/")
	//
	// append to autotime.log
	logPath += env.PathSeparator() + "autotime.log"
	zr.AppendToTextFile(logPath, entry) // write to logName
} //                                                                 autoTimeLog

// getTimeLogPath _ _
func getTimeLogPath(path string) string {
	path = strings.ToLower(path)
	max := -1
	var ret string
	for _, s := range TimeLogPaths {
		s = strings.ToLower(s)
		if !strings.Contains(path, s) || len(s) < max {
			continue
		}
		ret = s
		max = len(s)
	}
	return ret
} //                                                              getTimeLogPath

// processDir calls processFile() for each file in 'dir' and its subfolders.
func processDir(dir string, changeTime bool) {
	log.SetFlags(log.Lshortfile)
	// TODO: use fs.WalkPath() instead of this; then remove "os" dependency
	err := filepath.Walk(
		dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return zr.Error("Error reading path", path, "due to:", err)
			}
			if info.IsDir() {
				return nil
			}
			var tm time.Time
			if changeTime {
				tm = info.ModTime()
			}
			return processFile(path, info.Name(), tm)
		},
	)
	if err != nil {
		zr.Error(err.Error())
	}
} //                                                                  processDir

// processFile is called by processDir() to update timestamps
// in the named file. It ignores files that are not text files.
//
// path is the full path and file name
func processFile(path, name string, modTime time.Time) error {
	if !isTextFile(path) {
		return nil
	}
	data, done := env.ReadFile(path)
	if !done {
		return zr.Error(zr.EFailedReading, "file", path)
	}
	oldContent := string(data)
	content := replaceVersion(oldContent, path, name, modTime)
	if content == oldContent || content == "" {
		return nil
	}
	if !env.WriteFile(path, []byte(content)) {
		return zr.Error(zr.EFailedWriting, "file", path)
	}
	return nil
} //                                                                 processFile

// replaceVersion replaces the timestamp within a string
//
// s:        String to replace
// path:     Path of the file, excluding the file's name.
// filename: File's name without the path.
// modTime:  File's last modification time.
//           If it is zero-valued, does not change the timestamp.
func replaceVersion(s, path, filename string, modTime time.Time) string {
	// TODO: in replaceVersion(), 'path' is already including the filename.
	//      Then remove 'filename' argument.
	var loc []int
	{
		re := regexp.MustCompile(HeaderSignatureRX)
		loc = re.FindStringIndex(s)
	}
	// if there is no version string, return a blank to indicate so
	if loc == nil {
		return ""
	}
	var now string
	if modTime.IsZero() {
		i := loc[0] + HeaderTimePos
		now = s[i : i+19] // len('YYYY-MM-DD hh:mmm:ss')+1
	} else {
		now = zr.Timestamp()
	}
	body := s[loc[1]:]
	chk := checksum(body)
	if strings.Contains(s, chk) { // exit if there are no changes
		return ""
	}
	isTimed := true
	for _, s := range IgnoreFilenamesWith {
		if strings.Contains(strings.ToLower(path), s) {
			isTimed = false
			break
		}
	}
	if isTimed {
		autoTimeLog(path, now)
	}
	// modify and return file content
	// (writing the file is done by the caller)
	s = s[:loc[0]] + ":v: " + now + " " + chk + s[loc[1]:]
	return s
} //                                                              replaceVersion

// end

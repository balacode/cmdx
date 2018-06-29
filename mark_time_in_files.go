// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-28 14:11:34 190B85                   cmdx/[mark_time_in_files.go]
// -----------------------------------------------------------------------------

package main

// # Command Handler
//   markTimeInFiles(cmd Command, args []string)
//
// # Support (File Scope)
//   autoTimeLog(path string, timestamp string)
//   checksum(s string) string
//   getTimeLogPath(path string) string
//   processDir(dir string, changeTime bool)
//   processFile(path, name string, modTime time.Time) error
//   replaceVersion(s, path, filename string, modTime time.Time) string

import (
	"fmt"
	"hash/crc32"
	"log"
	"os"
	"path/filepath"
	"regexp"
	str "strings"
	"time"

	"github.com/balacode/zr"
	"github.com/balacode/zr-fs"
)

// -----------------------------------------------------------------------------
// # Command Handler

// markTimeInFiles __
// The 'cmd' argument is not used.
func markTimeInFiles(cmd Command, args []string) {
	var changeTime = true
	for _, arg := range args {
		switch str.ToLower(str.Trim(arg, SPACES+"-/")) {
		case "hash-only":
			changeTime = false
		default:
			env.Printf("unknown argument '%s'"+LF, arg)
			return
		}
	}
	var dir = env.Getwd()
	if dir != "" {
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
	var filename = filepath.Base(path)
	env.Printf("on %s %s"+LF, timestamp, filename)
	//
	//
	var logPath = str.ToLower(getTimeLogPath(path))
	//
	// the entry written in the log file:
	var entry = str.ToLower(timestamp + " " + path + LF)
	if str.Contains(entry, logPath) {
		entry = str.Replace(entry, logPath, "", -1)
	}
	entry = str.Replace(entry, "\\", "/", -1)
	//
	// append to autotime.log
	logPath += env.PathSeparator() + "autotime.log"
	zr.AppendToTextFile(logPath, entry) // write to logName
} //                                                                 autoTimeLog

// checksum returns a shortened CRC32 checksum of the given string.
// The returned checksum is a string made up of 6 hexadecimal digits,
// shorter than the 8 hex digits required for a normal CRC32.
func checksum(s string) string {
	var chk = crc32.ChecksumIEEE([]byte(s))
	chk = (chk / 0x00FFFFFF) ^ (chk & 0x00FFFFFF) // <- from 4 to 3 bytes
	return fmt.Sprintf("%06X", chk)
} //                                                                    checksum

// getTimeLogPath __
func getTimeLogPath(path string) string {
	path = str.ToLower(path)
	var max = -1
	var ret string
	for _, s := range TimeLogPaths {
		s = str.ToLower(s)
		if !str.Contains(path, s) || len(s) < max {
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
	//TODO: use fs.WalkPath() instead of this; then remove "os" dependency
	var err = filepath.Walk(
		dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				zr.Error("Error reading path", path, "due to:", err)
				return err
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
	if !fs.IsTextFile(path) {
		return nil
	}
	var data, done = env.ReadFile(path)
	if !done {
		return zr.Error(zr.EFailedReading, "file", path)
	}
	var oldContent = string(data)
	var content = replaceVersion(oldContent, path, name, modTime)
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
	//TODO: in replaceVersion(), 'path' is already including the filename.
	//      Then remove 'filename' argument.
	var loc []int
	{
		var re = regexp.MustCompile(HeaderSignatureRX)
		loc = re.FindStringIndex(s)
	}
	// if there is no version string, return a blank to indicate so
	if loc == nil {
		return ""
	}
	var now string
	if modTime.IsZero() {
		var i = loc[0] + HeaderTimePos
		now = s[i : i+19] // len('YYYY-MM-DD hh:mmm:ss')+1
	} else {
		now = zr.Timestamp()
	}
	var body = s[loc[1]:]
	var chk = checksum(body)
	if str.Contains(s, chk) { // exit if there are no changes
		return ""
	}
	var isTimed = true
	for _, s := range IgnoreFilenamesWith {
		if str.Contains(str.ToLower(path), s) {
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

//end

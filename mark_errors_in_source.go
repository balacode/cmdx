// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 01:58:41 6FBFA2                [cmdx/mark_errors_in_source.go]
// -----------------------------------------------------------------------------

package main

import "flag"        // standard
import "path"        // standard
import str "strings" // standard

import "github.com/balacode/zr" // Zirconium

/*
   to test this manually:
       remove any errors, and compile cx.exe:
           build.bat
       add some errors in source, then:
           go build -gcflags="-e" -o tmp.exe 2> build.log
           cx mark-errors -buildlog=.\build.log
           del tmp.exe
*/

// # Command Handler
//   markErrorsInSource(cmd Command, args []string)
//
// # Support (File Scope)
//   isErrorComment(line string) bool
//   makePath(absPath, relPath string) string
//   readBuildIssues(buildLog string) (ret []BuildIssue)
//   removeOldErrorComments(lines []string) []string
//   saveFile(buildPath, filename string, lines []string)

// -----------------------------------------------------------------------------
// # Command Handler

// markErrorsInSource __
func markErrorsInSource(cmd Command, args []string) {
	//TODO: change to use 'args' instead of reading flags
	//
	// read command-line arguments
	var buildPath, libPath, buildLog string
	flag.StringVar(&buildPath, "buildpath", "", "project build path")
	flag.StringVar(&libPath, "libpath", "", "library path")
	flag.StringVar(&buildLog, "buildlog", "", "build log file")
	flag.Parse()
	//
	// prepare paths
	if buildPath == "" {
		var path = env.Getwd()
		if path == "" {
			return
		}
		buildPath = path
	}
	if libPath == "" {
		libPath = DefaultLibPath
	}
	if buildLog == "" {
		buildLog = buildPath + env.PathSeparator() + "build.log"
	}
	// get array with issues/errors
	var issues = readBuildIssues(buildLog)
	if len(issues) == 0 {
		return
	}
	// print all issues/errors on the screen
	for i, issue := range issues {
		env.Printf("ISSUE %d file:%s line:%d msg:%s"+zr.LF,
			i, issue.File, issue.Line, issue.Msg)
	}
	var lines []string
	var prevFile string
	// iterate over issues array and insert error comments in the source file
	// important: must be done in reverse order to keep existing line numbers)
	for i := len(issues) - 1; i >= 0; i-- {
		var issue = issues[i]
		if prevFile != issue.File {
			// save previous file, and load next file
			saveFile(buildPath, prevFile, lines)
			prevFile = issue.File
			// read file into 'lines' array
			var path = makePath(buildPath, issue.File)
			var data, done = env.ReadFile(path)
			if !done {
				return
			}
			// mark existing error comments (for later removal)
			lines = str.Split(string(data), zr.LF)
			for i, line := range lines {
				if isErrorComment(line) {
					lines[i] = str.Replace(lines[i],
						ErrorEndMark, ErrorEndMark+OldMark, -1)
				}
			}
		}
		var msg = ErrorMark + issue.Msg + ErrorEndMark
		if issue.Line >= len(lines) {
			lines = append(lines, msg)
			continue
		}
		if trim(lines[issue.Line]) == trim(msg) {
			continue
		}
		{ // indent the error comment to match the error's column
			var gap = ""
			var line = lines[issue.Line-1]
			var max = len(line)
			for i := 0; i < (issue.Col-2) && i < max; i++ {
				if line[i] == '\t' {
					gap += "\t"
				} else {
					gap += " "
				}
			}
			msg = gap + msg
		}
		// shift lines to insert the comment
		lines = append(lines, "")
		copy(lines[issue.Line+1:], lines[issue.Line:])
		lines[issue.Line] = msg
	}
	saveFile(buildPath, prevFile, lines)
} //                                                          markErrorsInSource

// -----------------------------------------------------------------------------
// # Support (File Scope)

// isErrorComment returns true if the given line is an error comment
func isErrorComment(line string) bool {
	var find = trim(line)
	return str.HasPrefix(find, ErrorMark) &&
		str.Contains(find, ErrorEndMark)
} //                                                              isErrorComment

// makePath combines an absolute and relative path,
// returning an absolute path.
func makePath(absPath, relPath string) string {
	var ret string
	switch {
	case relPath == "":
		ret = absPath
	case path.IsAbs(relPath):
		ret = relPath
	default:
		var sep = env.PathSeparator()
		var abs = str.Split(absPath, sep)
		var rel = str.Split(relPath, sep)
		for len(rel) > 0 && rel[0] == ".." {
			abs = abs[:len(abs)-1]
			rel = rel[1:]
		}
		abs = append(abs, rel...)
		ret = str.Join(abs, sep)
	}
	return ret
} //                                                                    makePath

// readBuildIssues reads file given by 'buildLog'
// and returns an array of issues.
func readBuildIssues(buildLog string) (ret []BuildIssue) {
	// load the build log file:
	var data, done = env.ReadFile(buildLog)
	if !done {
		return nil
	}
	// fill issues array:
	for _, s := range str.Split(string(data), zr.LF) {
		var ar = str.Split(s, ":")
		if len(ar) >= 4 {
			ret = append(ret, BuildIssue{
				File: trim(ar[0]),
				Line: zr.Int(ar[1]),
				Col:  zr.Int(ar[2]),
				Msg:  trim(str.Join(ar[3:], ":")),
			})
		}
	}
	return ret
} //                                                             readBuildIssues

// removeOldErrorComments returns an array of lines,
// removing all error comments
func removeOldErrorComments(lines []string) []string {
	var ret []string
	for _, line := range lines {
		if isErrorComment(line) && str.HasSuffix(line, OldMark) {
			continue
		}
		ret = append(ret, line)
	}
	return ret
} //                                                      removeOldErrorComments

// saveFile saves 'filename' in 'buildPath' using 'lines' for its content
func saveFile(buildPath, filename string, lines []string) {
	if filename == "" {
		return
	}
	lines = removeOldErrorComments(lines)
	env.WriteFile(
		makePath(buildPath, filename),
		[]byte(str.Join(lines, zr.LF)),
	)
} //                                                                    saveFile

//TODO: use saveFile() in ase/zr_fs

//end
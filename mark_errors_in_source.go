// -----------------------------------------------------------------------------
// CMDX Utilities Suite                          cmdx/[mark_errors_in_source.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	"flag"
	"path"
	"strings"

	"github.com/balacode/zr"
)

//  to test this manually:
//      remove any errors, and compile cx.exe:
//          build.bat
//      add some errors in source, then:
//          go build -gcflags="-e" -o tmp.exe 2> build.log
//          cx mark-errors -buildlog=.\build.log
//          del tmp.exe

// # Command Handler
//   markErrorsInSource(cmd Command, args []string)
//
// # Support (File Scope)
//   isErrorComment(line string) bool
//   isRepeatComment(msg string, lineNo int, lines []string) bool
//   makePath(absPath, relPath string) string
//   readBuildIssues(buildLog string) (ret []BuildIssue)
//   removeOldErrorComments(lines []string) []string
//   saveFile(buildPath, filename string, lines []string)

// -----------------------------------------------------------------------------
// # Command Handler

// markErrorsInSource _ _
func markErrorsInSource(cmd Command, args []string) {
	// TODO: change to use 'args' instead of reading flags
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
		path := env.Getwd()
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
	issues := readBuildIssues(buildLog)
	if len(issues) == 0 {
		return
	}
	// print all issues/errors on the screen
	for i, issue := range issues {
		env.Printf("ISSUE %d file:%s line:%d msg:%s\n",
			i, issue.File, issue.Line, issue.Msg)
	}
	var lines []string
	var prevFile string
	//
	// iterate over issues array and insert error comments in the source file
	// important: must be done in reverse order to keep existing line numbers)
	for i := len(issues) - 1; i >= 0; i-- {
		issue := issues[i]
		if prevFile != issue.File {
			//
			// save previous file, and load next file
			saveFile(buildPath, prevFile, lines)
			prevFile = issue.File
			//
			// read file into 'lines' array
			path := makePath(buildPath, issue.File)
			data, done := env.ReadFile(path)
			if !done {
				return
			}
			// mark existing error comments (for later removal)
			lines = strings.Split(string(data), "\n")
			for i, line := range lines {
				if isErrorComment(line) {
					lines[i] = strings.ReplaceAll(lines[i],
						ErrorEndMark, ErrorEndMark+OldMark)
				}
			}
		}
		msg := ErrorMark + issue.Msg + ErrorEndMark
		if issue.Line >= len(lines) {
			lines = append(lines, msg)
			continue
		}
		// avoid repeating the same comment
		if isRepeatComment(msg, issue.Line, lines) {
			continue
		}
		{ // align the comment to the error's column
			gap := ""
			line := lines[issue.Line-1]
			max := len(line)
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
}

// -----------------------------------------------------------------------------
// # Support (File Scope)

// isErrorComment returns true if the given line is an error comment
func isErrorComment(line string) bool {
	find := trim(line)
	return strings.HasPrefix(find, ErrorMark) &&
		strings.Contains(find, ErrorEndMark)
}

// isRepeatComment returns true if msg already
// exists at or adjacent to the line at lineNo.
func isRepeatComment(msg string, lineNo int, lines []string) bool {
	msg = trim(msg)
	//
	// check on exact line
	if trim(lines[lineNo]) == msg {
		return true
	}
	// check preceding line
	if (lineNo-1) >= 0 && trim(lines[lineNo-1]) == msg {
		return true
	}
	// check following line
	if (lineNo+1) < len(lines) && trim(lines[lineNo+1]) == msg {
		return true
	}
	return false
}

// makePath combines an absolute and relative path,
// returning an absolute path.
func makePath(absPath, relPath string) string {
	var ret string
	switch {
	case relPath == "":
		{
			ret = absPath
		}
	case path.IsAbs(relPath):
		{
			ret = relPath
		}
	default:
		var (
			sep = env.PathSeparator()
			abs = strings.Split(absPath, sep)
			rel = strings.Split(relPath, sep)
		)
		for len(rel) > 0 && rel[0] == ".." {
			abs = abs[:len(abs)-1]
			rel = rel[1:]
		}
		abs = append(abs, rel...)
		ret = strings.Join(abs, sep)
	}
	return ret
}

// readBuildIssues reads file given by 'buildLog'
// and returns an array of issues.
func readBuildIssues(buildLog string) (ret []BuildIssue) {
	// load the build log file:
	data, done := env.ReadFile(buildLog)
	if !done {
		return nil
	}
	// fill issues array:
	m := map[string]bool{}
	lines := strings.Split(string(data), "\n")
	for _, s := range lines {
		ar := strings.Split(s, ":")
		if len(ar) < 4 { //                                     skip short lines
			continue
		}
		if _, exist := m[s]; exist { //                     skip repeated errors
			continue
		}
		m[s] = true
		ret = append(ret, BuildIssue{
			File: trim(ar[0]),
			Line: zr.Int(ar[1]),
			Col:  zr.Int(ar[2]),
			Msg:  trim(strings.Join(ar[3:], ":")),
		})
	}
	return ret
}

// removeOldErrorComments returns an array of lines,
// removing all error comments
func removeOldErrorComments(lines []string) []string {
	var ret []string
	for _, line := range lines {
		if isErrorComment(line) && strings.HasSuffix(line, OldMark) {
			continue
		}
		ret = append(ret, line)
	}
	return ret
}

// saveFile saves 'filename' in 'buildPath' using 'lines' for its content
func saveFile(buildPath, filename string, lines []string) {
	if filename == "" {
		return
	}
	lines = removeOldErrorComments(lines)
	env.WriteFile(
		makePath(buildPath, filename),
		[]byte(strings.Join(lines, "\n")),
	)
}

// TODO: use saveFile() in github.com/balacode/zr-fs

// end

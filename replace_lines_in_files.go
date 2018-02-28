// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-28 14:06:54 A20374               [cmdx/replace_lines_in_files.go]
// -----------------------------------------------------------------------------

package main

import "path/filepath" // standard
import "sync"          // standard
import "sync/atomic"   // standard
import str "strings"   // standard

import "github.com/balacode/zr" // Zircon-Go

//   DebugReplaceLinesInFiles
//
// # Command Handler
//   replaceLinesInFiles(cmd Command, args []string)
//
// # Subfunctions
//   replaceLinesInFilesM struct
//   (M replaceLinesInFilesM) getFindRepls(
//                                configLines []string,
//                            ) (ret []FindReplLines)
//   (M replaceLinesInFilesM) replaceFileAsync(
//       task *sync.WaitGroup,
//       changesAtomic *int32,
//       filename string,
//       lines []string,
//       findRepls []FindReplLines,
//   )
//   (M replaceLinesInFilesM) trimBlankLines(lines []string) []string
//   (M replaceLinesInFilesM) trimStrings(strs []string) []string

// DebugReplaceLinesInFiles displays arguments and the
// return value  to the console, when set to true
const DebugReplaceLinesInFiles = false

// replaceLinesInFilesM joins all subfunctions used by
// replaceLinesInFiles(), so that their names don't
// clutter the project's namespace.
type replaceLinesInFilesM struct{}

// -----------------------------------------------------------------------------
// # Command Handler

// replaceLinesInFiles __
func replaceLinesInFiles(cmd Command, args []string) {
	if len(args) != 1 {
		env.Println("requires <command-file> parameter")
		return
	}
	var M replaceLinesInFilesM
	var divider = str.Repeat("-", 80)
	var configFile = args[0]
	var pathExtsMap = map[string][]FindReplLines{}
	var findRepls = M.getFindRepls(env.ReadFileLines(configFile))
	var err error
	//
	configFile, err = filepath.Abs(configFile)
	env.Println(divider)
	if err != nil {
		env.Println("Failed getting path of", configFile, "due to:", err)
		return
	}
	// group batches of items by their path and extensions (using a map)
	for _, it := range findRepls {
		// join path and extensions list to give a map key
		var key = it.Path + LF + str.Join(it.Exts, LF)
		pathExtsMap[key] = append(pathExtsMap[key], it)
	}
	env.Println(divider)
	//
	// for each file in group, call replaceFileAsync() with applicable items
	var task sync.WaitGroup
	var changesAtomic int32
	for key, items := range pathExtsMap {
		var ar = str.Split(key, LF) // read details back from key
		var path = ar[0]
		var exts = ar[1:]
		var fileList = env.GetFilePaths(path, exts...)
		for _, filename := range fileList {
			if filename == configFile {
				continue // must not overwrite the config file itself
			}
			var data, done = env.ReadFile(filename)
			if !done {
				continue
			}
			var lines = str.Split(string(data), "\n")
			task.Add(1)
			go M.replaceFileAsync(&task, &changesAtomic,
				filename, lines, items)
		}
	}
	task.Wait()
	//
	// report total number of changes
	var n = atomic.LoadInt32(&changesAtomic)
	if n == 0 {
		env.Println("NO CHANGES")
	} else {
		env.Println(n, "TOTAL")
	}
} //                                                         replaceLinesInFiles

// -----------------------------------------------------------------------------
// # Subfunctions

// getFindRepls __
func (M replaceLinesInFilesM) getFindRepls(
	configLines []string,
) (ret []FindReplLines) {
	const FreeMode = 0
	const FindMode = 1
	const ReplMode = 2
	//
	var mark = DefaultMark
	var path = DefaultPath
	var exts = DefaultExts
	var undo = false
	var mode = FreeMode
	var caseMode = zr.MatchCase
	var findGroup []string
	var replGroup []string
	for _, line := range configLines {
		// lines that begin with the marker are configuration or comments:
		if str.HasPrefix(line, mark) {
			line = str.Trim(line[len(mark):], SPACES)
			switch {
			case str.HasPrefix(line, "path"):
				path = str.Trim(line[5:], SPACES)
				env.Println("SET PATH:", path)
			case str.HasPrefix(line, "exts"):
				exts = str.Fields(line[5:])
				env.Println("SET EXTS:", exts)
			case str.HasPrefix(line, "mark"):
				mark = str.Trim(line[5:], SPACES)
				if mark == "" {
					mark = DefaultMark
				}
				env.Println("SET MARK:", mark)
			// booleans:
			case hasConfigBool(line, "case"):
				var match, _ = getConfigBool(line, "case")
				env.Println("SET CASE:", match)
				if match {
					caseMode = zr.MatchCase
				} else {
					caseMode = zr.IgnoreCase
				}
			case hasConfigBool(line, "undo"):
				undo, _ = getConfigBool(line, "undo")
				env.Println("SET UNDO:", undo)
			case str.HasPrefix(line, "<"): // find lines
				mode = FindMode
			case str.HasPrefix(line, ">"): // replace with lines
				mode = ReplMode
			case str.HasPrefix(line, "."):
				mode = FreeMode
				var it = FindReplLines{
					Path:      path,
					Exts:      exts,
					FindLines: M.trimBlankLines(M.trimStrings(findGroup)),
					ReplLines: M.trimBlankLines(replGroup),
					CaseMode:  caseMode,
				}
				if undo {
					it.FindLines, it.ReplLines = it.ReplLines, it.FindLines
				}
				ret = append(ret, it)
				findGroup = []string{}
				replGroup = []string{}
			}
			continue
		}
		switch {
		case mode == FindMode:
			findGroup = append(findGroup, line)
		case mode == ReplMode:
			replGroup = append(replGroup, line)
		}
	}
	if DebugReplaceLinesInFiles {
		var pl = env.Println
		pl("getFindRepls() RETURNS", len(ret), "ITEM(S):")
		for i, it := range ret {
			pl("ITEM", i)
			zr.DV("Path:", it.Path)
			zr.DV("Exts:", it.Exts)
			zr.DV("CaseMode", it.CaseMode)
			zr.DV("FindLines:", it.FindLines)
			zr.DV("ReplLines:", it.ReplLines)
		}
	}
	return ret
} //                                                                getFindRepls

// replaceFileAsync __
func (M replaceLinesInFilesM) replaceFileAsync(
	task *sync.WaitGroup,
	changesAtomic *int32,
	filename string,
	lines []string,
	findRepls []FindReplLines,
) {
	if task != nil {
		defer task.Done()
	}
	var changes = 0
	for _, caseMode := range []zr.CaseMode{zr.MatchCase, zr.IgnoreCase} {
		var finds, repls [][]string
		for _, it := range findRepls {
			if it.CaseMode != caseMode {
				continue
			}
			finds = append(finds, it.FindLines)
			repls = append(repls, it.ReplLines)
		}
		var n int
		lines, n = replaceLines(lines, finds, repls, caseMode)
		changes += n
	}
	if changes == 0 {
		return
	}
	if !env.WriteFileLines(filename, lines) {
		return // auto-logs error
	}
	env.Println(changes, "replacement in", filename)
	atomic.AddInt32(changesAtomic, int32(changes))
} //                                                            replaceFileAsync

// trimBlankLines removes leading and trailing blank lines in a
// string slice. Does not remove blank lines between non-blank lines.
// Lines that only contain white spaces are treated as blank lines.
func (M replaceLinesInFilesM) trimBlankLines(lines []string) []string {
	// trim leading blank lines
	for len(lines) > 0 && str.Trim(lines[0], SPACES) == "" {
		lines = lines[1:]
	}
	// trim trailing blank lines
	for len(lines) > 0 &&
		str.Trim(lines[len(lines)-1], SPACES) == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
} //                                                              trimBlankLines

// trimStrings removes leading and trailing spaces from each line in strs.
func (M replaceLinesInFilesM) trimStrings(strs []string) []string {
	for i, s := range strs {
		strs[i] = str.Trim(s, SPACES)
	}
	return strs
} //                                                                 trimStrings

//end

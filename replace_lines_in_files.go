// -----------------------------------------------------------------------------
// CMDX Utilities Suite                         cmdx/[replace_lines_in_files.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/balacode/zr"
)

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
//   (M replaceLinesInFilesM) trimStrings(ar []string) []string

// DebugReplaceLinesInFiles displays arguments and the
// return value  to the console, when set to true
const DebugReplaceLinesInFiles = false

// replaceLinesInFilesM joins all subfunctions used by
// replaceLinesInFiles(), so that their names don't
// clutter the project's namespace.
type replaceLinesInFilesM struct{}

// -----------------------------------------------------------------------------
// # Command Handler

// replaceLinesInFiles _ _
func replaceLinesInFiles(cmd Command, args []string) {
	if len(args) != 1 {
		env.Println("requires <command-file> parameter")
		return
	}
	var (
		divider     = strings.Repeat("-", 80)
		configFile  = args[0]
		pathExtsMap = map[string][]FindReplLines{}
		M           replaceLinesInFilesM
		findRepls   = M.getFindRepls(env.ReadFileLines(configFile))
		err         error
	)
	configFile, err = filepath.Abs(configFile)
	env.Println(divider)
	if err != nil {
		env.Println("Failed getting path of", configFile, "due to:", err)
		return
	}
	// group batches of items by their path and extensions (using a map)
	for _, it := range findRepls {
		// join path and extensions list to give a map key
		key := it.Path + "\n" + strings.Join(it.Exts, "\n")
		pathExtsMap[key] = append(pathExtsMap[key], it)
	}
	env.Println(divider)
	//
	// for each file in group, call replaceFileAsync() with applicable items
	var task sync.WaitGroup
	var changesAtomic int32
	for key, items := range pathExtsMap {
		var (
			ar       = strings.Split(key, "\n") // read details back from key
			path     = ar[0]
			exts     = ar[1:]
			fileList = env.GetFilePaths(path, exts...)
		)
		for _, filename := range fileList {
			if filename == configFile {
				continue // must not overwrite the config file itself
			}
			data, done := env.ReadFile(filename)
			if !done {
				continue
			}
			lines := strings.Split(string(data), "\n")
			task.Add(1)
			go M.replaceFileAsync(&task, &changesAtomic,
				filename, lines, items)
		}
	}
	task.Wait()
	//
	// report total number of changes
	n := atomic.LoadInt32(&changesAtomic)
	if n == 0 {
		env.Println("NO CHANGES")
	} else {
		env.Println(n, "TOTAL")
	}
}

// -----------------------------------------------------------------------------
// # Subfunctions

// getFindRepls _ _
func (M replaceLinesInFilesM) getFindRepls(
	configLines []string,
) (ret []FindReplLines) {
	const (
		FreeMode = 0
		FindMode = 1
		ReplMode = 2
	)
	var (
		mark      = DefaultMark
		path      = DefaultPath
		exts      = DefaultExts
		undo      = false
		mode      = FreeMode
		caseMode  = zr.MatchCase
		findGroup []string
		replGroup []string
	)
	for _, line := range configLines {
		// lines that begin with the marker are configuration or comments:
		if strings.HasPrefix(line, mark) {
			line = strings.TrimSpace(line[len(mark):])
			switch {
			case strings.HasPrefix(line, "path"):
				path = strings.TrimSpace(line[5:])
				env.Println("SET PATH:", path)
			case strings.HasPrefix(line, "exts"):
				exts = strings.Fields(line[5:])
				env.Println("SET EXTS:", exts)
			case strings.HasPrefix(line, "mark"):
				mark = strings.TrimSpace(line[5:])
				if mark == "" {
					mark = DefaultMark
				}
				env.Println("SET MARK:", mark)
			// booleans:
			case hasConfigBool(line, "case"):
				match, _ := getConfigBool(line, "case")
				env.Println("SET CASE:", match)
				if match {
					caseMode = zr.MatchCase
				} else {
					caseMode = zr.IgnoreCase
				}
			case hasConfigBool(line, "undo"):
				undo, _ = getConfigBool(line, "undo")
				env.Println("SET UNDO:", undo)
			case strings.HasPrefix(line, "<"): // find lines
				mode = FindMode
			case strings.HasPrefix(line, ">"): // replace with lines
				mode = ReplMode
			case strings.HasPrefix(line, "."):
				mode = FreeMode
				it := FindReplLines{
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
		pl := env.Println
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
}

// replaceFileAsync _ _
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
	changes := 0
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
}

// trimBlankLines removes leading and trailing blank lines in a
// string slice. Does not remove blank lines between non-blank lines.
// Lines that only contain white spaces are treated as blank lines.
func (M replaceLinesInFilesM) trimBlankLines(lines []string) []string {
	// trim leading blank lines
	for len(lines) > 0 && strings.TrimSpace(lines[0]) == "" {
		lines = lines[1:]
	}
	// trim trailing blank lines
	for len(lines) > 0 &&
		strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

// trimStrings removes leading and trailing spaces from each line in strs.
func (M replaceLinesInFilesM) trimStrings(ar []string) []string {
	for i, s := range ar {
		ar[i] = strings.TrimSpace(s)
	}
	return ar
}

// end

// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 01:58:41 5F2111                      [cmdx/run_interactive.go]
// -----------------------------------------------------------------------------

package main

// # Main Command Handler
//   runInteractive(cmd Command, args []string)
//
// # Subfunctions
//   (ob Runner) getMarkedBlocks(lines []string) (ret []string)
//   (ob Runner) processFile(file *TextFile) (retAltered bool)
//   (ob Runner) sortByModTime(
//       files []string,
//       filesMap *map[string]*TextFile,
//   )
//   (ob Runner) stripErrorMarks(
//       lines []string,
//   ) (
//       modLines []string, altered bool,
//   )
//
// # Command Handlers
//   (ob Runner) insertID(ln, col int, modLines []string) (altered bool)
//   (ob Runner) insertTimestamp(
//       ln, col int,
//       modLines []string,
//   ) (altered bool)
//   (ob Runner) insertUUID(ln, col int, modLines []string) (altered bool)

import "fmt"         // standard
import "sort"        // standard
import "sync"        // standard
import str "strings" // standard

import "github.com/balacode/zr" // Zirconium
import "github.com/balacode/zr_fs"              // Zirconium

// Runner joins all subfunctions used by runInteractive(),
// so that their names don't clutter the project's namespace.
type Runner struct {
	memFiles LoadedTextFiles
}

// -----------------------------------------------------------------------------
// # Main Command Handler

// runInteractive runs the interactive text processor.
// which continuously watches all text files in RootPath
// for changes. When it detects text processing commands,
// it applies these commands to the applicable source files.
func runInteractive(cmd Command, args []string) {
	var refreshChan = fs.NewDirWatcher(RootPath).C
	var quitChan = make(chan bool)
	var fsMx sync.RWMutex // mutex for file system operations
	var ob Runner
	//
	// load all text files into memory
	ob.memFiles.LoadAll(RootPath, &fsMx)
	env.Println("Interactive mode...")
	//
	// begin interactive loop
loop:
	for {
		select {
		case <-refreshChan:
			var changedFilenames = ob.memFiles.LoadAll(RootPath, &fsMx)
			for i := range changedFilenames {
				var file = ob.memFiles.GetFile(changedFilenames[i])
				for _, find := range IgnoreFilenamesWith {
					if str.Contains(str.ToLower(file.Path), find) {
						continue loop
					}
				}
				autoTimeLog(file.Path, zr.Timestamp())
				//
				// call processFile() to edit the file and save if altered
				if !ob.processFile(file) {
					continue
				}
				fsMx.Lock()
				env.WriteFileLines(file.Path, file.Lines)
				fsMx.Unlock()
				ob.memFiles.LoadFile(file.Path, &fsMx)
				env.Println("changed", file.Path)
			}
		case <-quitChan:
			env.Println("QUIT")
			return
		}
	}
} //                                                              runInteractive

// -----------------------------------------------------------------------------
// # Subfunctions

// getMarkedBlocks filters lines, returning only these lines
// contained between block beginning and ending markers
func (ob Runner) getMarkedBlocks(lines []string) (ret []string) {
	var b = 0 // <- remaining lines in block
	for i, s := range lines {
		var ts = str.ToUpper(str.Trim(s, zr.SPACES))
		switch {
		case str.HasPrefix(ts, CommandMark+BB): // begin block
			b = -1
		case str.HasPrefix(ts, CommandMark+BE), // block end
			str.HasPrefix(ts, CommandMark+EB): // end block
			ret = append(ret, []string{"", "", ""}...) // three blank lines
			b = 0
		case str.HasSuffix(ts, CommandMark+LT):
			//TODO: IMPLEMENT
		case str.HasPrefix(ts, CommandMark) &&
			zr.IsNumber(ts[len(CommandMark):]):
			b = zr.Int(ts[len(CommandMark):])
		case b != 0:
			if b > 0 {
				b--
			}
			ret = append(ret, lines[i])
		}
	}
	return ret
} //                                                             getMarkedBlocks

// processFile __
// The lines slice is modified in-place.
func (ob Runner) processFile(file *TextFile) (retAltered bool) {
	var ln int   // current line number
	var s string // current line
	var lines = make([]string, len(file.Lines))
	copy(lines, file.Lines)
	//
	for ln < len(lines) {
		var altered bool
		s = lines[ln]
		if !str.Contains(s, CommandMark) {
			ln++
			continue
		}
		switch {
		case zr.ContainsI(s, CommandMark+CLL): // copy long lines
			lines[ln] = zr.ReplaceI(s,
				CommandMark+CLL, CommandMark+" DONE "+CLL)
			env.Println(file.Path + ":")
			for _, s := range filterLongLines(file.Lines, ColumnsNorm) {
				env.Println(s)
			}
			env.Println()
			//
		case zr.ContainsI(s, CommandMark+BC), // blocks (marked) copy
			zr.ContainsI(s, CommandMark+CB): // copy (marked) blocks
			var s = lines[ln]
			s = zr.ReplaceI(s, CommandMark+CB, CommandMark+" DONE "+CB)
			s = zr.ReplaceI(s, CommandMark+BC, CommandMark+" DONE "+BC)
			lines[ln] = s
			env.Println(str.Repeat("~", len(file.Path+":")))
			env.Println(file.Path + ":" + zr.LF + zr.LF + zr.LF)
			for _, s := range ob.getMarkedBlocks(file.Lines) {
				env.Println(s)
			}
			env.Println()
			//
		case zr.ContainsI(s, CommandMark+FF): // Find in Files
			var col = str.Index(str.ToUpper(s), CommandMark+FF)
			var dlg = str.Replace(FindInFilesDialog, "MRK", CommandMark, -1)
			lines[ln] = s[:col] + dlg + s[col+len(CommandMark+FF):]
			altered = true
			//
		case zr.ContainsI(s, CommandMark+XE): // remove error markers
			lines[ln] = zr.ReplaceI(s, CommandMark+XE, CommandMark+" DONE "+XE)
			if ar, some := ob.stripErrorMarks(lines); some {
				lines = ar
				altered = true
			}
		}
		for _, cmd := range []struct {
			name    string
			handler func(ln, col int, modLines []string) (altered bool)
		}{
			{ID, ob.insertID},
			{T, ob.insertTimestamp},
			{UUID, ob.insertUUID},
		} {
			var col = str.Index(str.ToUpper(s), CommandMark+cmd.name)
			if col != -1 {
				// modify lines in-place with handler
				altered = cmd.handler(ln, col, lines)
			}
		}
		if altered {
			file.Lines = lines
			ln = 0 // start over, since some lines may have been removed, etc.
			retAltered = true
			continue
		}
		ln++
	}
	return retAltered
} //                                                                 processFile

// sortByModTime sorts a list of file names by time modified.
//
// files:    the slice of file names being sorted
// filesMap: the map from which to read modification times (not modified)
func (ob Runner) sortByModTime(
	files []string,
	filesMap *map[string]*TextFile,
) {
	sort.Slice(files, func(i, j int) bool {
		var a = (*filesMap)[files[i]]
		var b = (*filesMap)[files[j]]
		//
		// this should never occur
		if a == nil || b == nil {
			zr.Error(zr.ENil)
			return true // push front
		}
		return a.ModTime.After(b.ModTime)
	})
} //                                                               sortByModTime

// stripErrorMarks removes all error markers from lines.
// Error markers are separate lines inserted under lines with build errors
// Each marked line begins with ErrorMark (can be preceded by white space)
func (ob Runner) stripErrorMarks(
	lines []string,
) (
	modLines []string, altered bool,
) {
	modLines = make([]string, 0, len(lines))
	for _, s := range lines {
		if str.HasPrefix(str.Trim(s, zr.SPACES), ErrorMark) {
			altered = true
			continue
		}
		modLines = append(modLines, s)
	}
	return modLines, altered
} //                                                             stripErrorMarks

// -----------------------------------------------------------------------------
// # Command Handlers

// insertID inserts a unique 6-digit hexadecimal ID
// at the specified position in lines, replaces the
// command marker and modifies lines in-place.
// The ID is unique between all text files in RootPath.
// Always returns true in altered.
func (ob Runner) insertID(ln, col int, modLines []string) (altered bool) {
	var s = modLines[ln]
	var id string
loop:
	for {
		id = str.ToUpper(zr.UUID()[:6])
		if id[0] != 'E' {
			continue
		}
		// check if the ID is unique
		var lineCount = 0
		for _, filename := range ob.memFiles.GetAllFilenames() {
			var file = ob.memFiles.GetFile(filename)
			for _, line := range file.Lines {
				line = str.ToUpper(line)
				if str.Contains(line, id) {
					fmt.Print("regenerating non-unique ID:", id, zr.LF)
					continue loop
				}
				lineCount++
			}
		}
		fmt.Print(
			"Created ID:", id, " (checked ", lineCount, " LOC)", zr.LF,
		)
		modLines[ln] = s[:col] + id + s[col+len(CommandMark+ID):]
		break
	}
	return true
} //                                                                    insertID

// insertTimestamp inserts a timestamp at the specified position in lines,
// replaces the command marker, and modifies lines in-place.
// Always returns true in altered.
func (ob Runner) insertTimestamp(ln, col int, lines []string) (altered bool) {
	var s = lines[ln]
	var aft = s[col+len(CommandMark+T):]
	var now = zr.Timestamp()
	if !str.HasPrefix(aft, " ") {
		now += " "
	}
	lines[ln] = s[:col] + "//" + now + aft
	return true
} //                                                             insertTimestamp

// insertUUID inserts a UUID at the specified position in lines,
// replaces the command marker, and modifies lines in-place.
// Always returns true in altered.
func (ob Runner) insertUUID(ln, col int, lines []string) (altered bool) {
	var s = lines[ln]
	lines[ln] = s[:col] + zr.UUID() + s[col+len(CommandMark+UUID):]
	return true
} //                                                                  insertUUID

//end
// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-26 23:09:04 D3E8DE             [cmdx/replace_strings_in_files.go]
// -----------------------------------------------------------------------------

package main

import "fmt"

import "path/filepath" // standard
import "sync"          // standard
import str "strings"   // standard

import "github.com/balacode/zr" // Zircon-Go

// # Command Handler
//   replaceStringsInFiles(cmd Command, args []string)
//
// # Support (File Scope)
//   replaceAsync(task *sync.WaitGroup, cmd ReplCmd)
//   replaceFileAsync(
//       task *sync.WaitGroup,
//       filename string,
//       content string,
//       items []ReplItem,
//   )

// -----------------------------------------------------------------------------
// # Command Handler

// replaceStringsInFiles __
func replaceStringsInFiles(cmd Command, args []string) {
	if len(args) != 1 {
		env.Println("requires <command-file> parameter")
		return
	}
	var configFile = args[0]
	var configLines []string
	var err error
	configFile, err = filepath.Abs(configFile)
	if err != nil {
		env.Println("command file path error: ", configFile)
		return
	}
	env.Println("FILE:", configFile)
	{
		var data, done = env.ReadFile(configFile)
		if !done {
			return
		}
		var s = string(data)
		s = str.Trim(s, zr.SPACES)
		s = str.Replace(s, "\r"+zr.LF, zr.LF, -1)
		for str.Contains(s, zr.LF+zr.LF) {
			s = str.Replace(s, zr.LF+zr.LF, zr.LF, -1)
		}
		configLines = str.Split(s, zr.LF)
		configLines = append(configLines, "") // initiates replacement
	}
	var path = DefaultPath
	var exts = DefaultExts
	var mark = DefaultMark
	var undo = false
	// each item:
	var caseMode = zr.MatchCase
	var wordMode = zr.IgnoreWord
	var items = []ReplItem{}
	// helper functions:
	var getBool = func(s, keyword string) (value, exists bool) {
		s = str.ToUpper(s)
		keyword = str.ToUpper(keyword)
		for i, ar := range [][]string{
			{"0", "FALSE", "OFF", "IGNORE"},
			{"1", "TRUE", "ON", "MATCH"},
		} {
			for _, match := range ar {
				if str.HasPrefix(s, keyword+" "+match) {
					return i == 1, true
				}
			}
		}
		return false, false
	}
	var hasBool = func(s, keyword string) (ret bool) {
		_, ret = getBool(s, keyword)
		return ret
	}
	env.Println(str.Repeat("-", 80))
	for lineNo, s := range configLines {
		s = str.Trim(s, zr.SPACES)
		// blank lines initiate replacement:
		if s == "" && len(items) > 0 {
			env.Println(str.Repeat("-", 80))
			var task sync.WaitGroup
			task.Add(1)
			var cmd = ReplCmd{
				Path:  path,
				Exts:  exts,
				Mark:  mark,
				Items: items,
			}
			go replaceAsync(&task, configFile, cmd)
			task.Wait()
			items = []ReplItem{}
			continue
		}
		// skip lines that don't contain marker:
		if !str.Contains(s, mark) {
			continue
		}
		// lines that begin with the marker are configuration or comments:
		if str.HasPrefix(s, mark) {
			s = str.Trim(s[len(mark):], zr.SPACES)
			switch {
			case str.HasPrefix(s, "path"):
				path = str.Trim(s[5:], zr.SPACES)
				env.Println("SET PATH:", path)
			case str.HasPrefix(s, "exts"):
				exts = str.Fields(s[5:])
				env.Println("SET EXTS:", exts)
			case str.HasPrefix(s, "mark"):
				mark = str.Trim(s[5:], zr.SPACES)
				if mark == "" {
					mark = DefaultMark
				}
				env.Println("SET MARK:", mark)
			// booleans:
			case hasBool(s, "case"):
				var match, _ = getBool(s, "case")
				env.Println("SET CASE:", match)
				if match {
					caseMode = zr.MatchCase
				} else {
					caseMode = zr.IgnoreCase
				}
			case hasBool(s, "undo"):
				undo, _ = getBool(s, "undo")
				env.Println("SET UNDO:", undo)
			case hasBool(s, "word"):
				var match, _ = getBool(s, "word")
				env.Println("SET WORD:", match)
				if match {
					wordMode = zr.MatchWord
				} else {
					wordMode = zr.IgnoreWord
				}
			}
			continue
		}
		// lines that contain but don't begin with the marker are replacements
		var i = str.Index(s, mark)
		if i > 0 {
			var item = ReplItem{
				Find:     str.Trim(s[:i], zr.SPACES),
				Repl:     str.Trim(s[i+len(mark):], zr.SPACES),
				CaseMode: caseMode,
				WordMode: wordMode,
			}
			if undo {
				item.Find, item.Repl = item.Repl, item.Find
			}
			if lineNo < ShownResultsLimit {
				env.Println(
					"FIND:", item.Find,
					"REPL:", item.Repl,
					"CASE:", item.CaseMode,
					"WORD:", item.WordMode,
				)
			} else if lineNo == ShownResultsLimit {
				env.Println("+", len(configLines)-lineNo, "more")
			}
			items = append(items, item)
		}
	}
} //                                                       replaceStringsInFiles

// -----------------------------------------------------------------------------
// # Support (File Scope)

// replaceAsync __
func replaceAsync(task *sync.WaitGroup, configFile string, cmd ReplCmd) {
	//TODO: you can remove configFile arg, and add an if condition in caller
	if task == nil {
		zr.Error("") //TODO: add error message (replaceAsync())
	}
	if task != nil {
		defer task.Done()
	}
	for _, filename := range env.GetFilePaths(cmd.Path, cmd.Exts...) {
		var data, done = env.ReadFile(filename)
		if !done {
			continue
		}
		if task != nil {
			task.Add(1)
		}
		go replaceFileAsync(task, configFile, filename, string(data), cmd.Items)
	}
} //                                                                replaceAsync

// replaceFileAsync __
func replaceFileAsync(
	task *sync.WaitGroup,
	configFile string,
	filename string,
	content string,
	items []ReplItem,
) {
	if task != nil {
		defer task.Done()
	}
	if filename == configFile {
		return
	}
	var oldContent = content
	var max = 100 / float64(len(items))
	var percent = ""
	if UseOldVersion {
		for i, it := range items {
			switch {
			case it.WordMode == zr.IgnoreWord && it.CaseMode == zr.MatchCase:
				content = str.Replace(content, it.Find, it.Repl, -1)
				//
			case it.WordMode == zr.IgnoreWord && it.CaseMode == zr.IgnoreCase:
				content = zr.ReplaceI(content, it.Find, it.Repl)
				//
			case it.WordMode == zr.MatchWord && it.CaseMode == zr.MatchCase:
				content = zr.ReplaceWord(content, it.Find, it.Repl,
					zr.MatchCase)
				//
			case it.WordMode == zr.MatchWord && it.CaseMode == zr.IgnoreCase:
				content = zr.ReplaceWord(content, it.Find, it.Repl,
					zr.IgnoreCase)
			}
			if ShowProgressIndicator {
				var pc = fmt.Sprintf("%1.1f%%",
					float64(int(max*float64(i)*10))/10)
				if percent != pc {
					env.Print(str.Repeat("\b", len(percent)), pc)
					percent = pc
				}
			}
		}
	}
	if UseNewVersion {
		var newContent = content
		for _, cm := range []zr.CaseMode{zr.MatchCase, zr.IgnoreCase} {
			for _, wm := range []zr.WordMode{zr.MatchWord, zr.IgnoreWord} {
				var finds, repls []string
				for i, it := range items {
					if it.CaseMode != cm || it.WordMode != wm {
						continue
					}
					finds = append(finds, it.Find)
					repls = append(repls, it.Repl)
					if ShowProgressIndicator {
						var pc = fmt.Sprintf("c:%v w:%v %1.1f%%",
							cm, wm, float64(int(max*float64(i)*10))/10)
						if percent != pc {
							env.Print(str.Repeat("\b", len(percent)), pc)
							percent = pc
						}
					}
				}
				newContent = zr.ReplaceMany(newContent, finds, repls,
					-1, cm, wm)
			}
		}
		if !UseOldVersion && UseNewVersion {
			content = newContent
		} else if newContent != content {
			env.Println("!!!! ReplaceMany() FAILED REPLACING PROPERLY !!!!")
			var name = hardcodedReplaceManyPath
			zr.AppendToTextFile(name+"_expected.txt", content)
			zr.AppendToTextFile(name+"_produced.txt", newContent)
		} else {
			env.Println("**** ReplaceMany() OK ****")
		}
	}
	if ShowProgressIndicator {
		env.Print(str.Repeat("\b", len(percent)))
	}
	if content == oldContent {
		return
	}
	if !env.WriteFile(filename, []byte(content)) {
		return
	}
	env.Println("changed ", filename)
} //                                                            replaceFileAsync

//end

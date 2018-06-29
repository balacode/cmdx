// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-28 14:11:34 E26B02             cmdx/[replace_strings_in_files.go]
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"path/filepath"
	str "strings"
	"sync"

	"github.com/balacode/zr"
)

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

// replConfig __
type replConfig struct {
	configFile string
	path       string
	exts       []string
	mark       string
	undo       bool
	caseMode   zr.CaseMode
	wordMode   zr.WordMode
} //                                                                  replConfig

// replaceStringsInFiles __
func replaceStringsInFiles(cmd Command, args []string) {
	if len(args) != 1 {
		env.Println("requires <command-file> parameter")
		return
	}
	var cfg = replConfig{
		configFile: args[0],
		path:       DefaultPath,
		exts:       DefaultExts,
		mark:       DefaultMark,
		undo:       false,
		caseMode:   zr.MatchCase,
		wordMode:   zr.IgnoreWord,
	}
	var configLines []string
	var err error
	cfg.configFile, err = filepath.Abs(cfg.configFile)
	if err != nil {
		env.Println("command file path error: ", cfg.configFile)
		return
	}
	env.Println("FILE:", cfg.configFile)
	{
		var data, done = env.ReadFile(cfg.configFile)
		if !done {
			return
		}
		var s = string(data)
		s = str.Trim(s, SPACES)
		s = str.Replace(s, "\r"+LF, LF, -1)
		for str.Contains(s, LF+LF) {
			s = str.Replace(s, LF+LF, LF, -1)
		}
		configLines = str.Split(s, LF)
		//
		// add a blank line to initiate replacement
		configLines = append(configLines, "")
	}
	//
	// each item:
	var items = []ReplItem{}
	env.Println(str.Repeat("-", 80))
	for lineNo, s := range configLines {
		s = str.Trim(s, SPACES)
		//
		// blank lines initiate replacement:
		if s == "" && len(items) > 0 {
			env.Println(str.Repeat("-", 80))
			var task sync.WaitGroup
			task.Add(1)
			var cmd = ReplCmd{
				Path:  cfg.path,
				Exts:  cfg.exts,
				Mark:  cfg.mark,
				Items: items,
			}
			go replaceAsync(&task, cfg.configFile, cmd)
			task.Wait()
			items = []ReplItem{}
			continue
		}
		// skip lines that don't contain marker:
		if !str.Contains(s, cfg.mark) {
			continue
		}
		// lines that begin with the marker are configuration or comments:
		if str.HasPrefix(s, cfg.mark) {
			setReplConfig(s, &cfg)
			continue
		}
		// lines that contain but don't begin with the marker are replacements
		var i = str.Index(s, cfg.mark)
		if i > 0 {
			var item = ReplItem{
				Find:     str.Trim(s[:i], SPACES),
				Repl:     str.Trim(s[i+len(cfg.mark):], SPACES),
				CaseMode: cfg.caseMode,
				WordMode: cfg.wordMode,
			}
			if cfg.undo {
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

// getBool __
func getBool(s, keyword string) (value, exists bool) {
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
} //                                                                     getBool

// hasBool __
func hasBool(s, keyword string) (ret bool) {
	_, ret = getBool(s, keyword)
	return ret
} //                                                                     hasBool

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
	var newContent = content
	//
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
			newContent = zr.ReplaceMany(newContent, finds, repls, -1, cm, wm)
		}
		content = newContent
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

// setReplConfig __
func setReplConfig(s string, cfg *replConfig) {
	s = str.Trim(s[len(cfg.mark):], SPACES)
	switch {
	case str.HasPrefix(s, "path"):
		cfg.path = str.Trim(s[5:], SPACES)
		env.Println("SET PATH:", cfg.path)
	//
	case str.HasPrefix(s, "exts"):
		cfg.exts = str.Fields(s[5:])
		env.Println("SET EXTS:", cfg.exts)
	//
	case str.HasPrefix(s, "mark"):
		cfg.mark = str.Trim(s[5:], SPACES)
		if cfg.mark == "" {
			cfg.mark = DefaultMark
		}
		env.Println("SET MARK:", cfg.mark)
	//
	case hasBool(s, "case"):
		var match, _ = getBool(s, "case")
		env.Println("SET CASE:", match)
		if match {
			cfg.caseMode = zr.MatchCase
		} else {
			cfg.caseMode = zr.IgnoreCase
		}
	//
	case hasBool(s, "undo"):
		cfg.undo, _ = getBool(s, "undo")
		env.Println("SET UNDO:", cfg.undo)
	//
	case hasBool(s, "word"):
		var match, _ = getBool(s, "word")
		env.Println("SET WORD:", match)
		if match {
			cfg.wordMode = zr.MatchWord
		} else {
			cfg.wordMode = zr.IgnoreWord
		}
	}
} //                                                               setReplConfig

//end

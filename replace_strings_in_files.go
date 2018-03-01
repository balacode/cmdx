// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-03-01 16:51:36 AFF406             [cmdx/replace_strings_in_files.go]
// -----------------------------------------------------------------------------

package main

import "fmt"

import "path/filepath" // standard
import "strings"       // standard
import "sync"          // standard

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
		s = strings.Trim(s, SPACES)
		s = strings.Replace(s, "\r"+LF, LF, -1)
		for strings.Contains(s, LF+LF) {
			s = strings.Replace(s, LF+LF, LF, -1)
		}
		configLines = strings.Split(s, LF)
		configLines = append(configLines, "") // initiates replacement
	}
	//
	// each item:
	var items = []ReplItem{}
	env.Println(strings.Repeat("-", 80))
	for lineNo, s := range configLines {
		s = strings.Trim(s, SPACES)
		//
		// blank lines initiate replacement:
		if s == "" && len(items) > 0 {
			env.Println(strings.Repeat("-", 80))
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
		if !strings.Contains(s, cfg.mark) {
			continue
		}
		// lines that begin with the marker are configuration or comments:
		if strings.HasPrefix(s, cfg.mark) {
			setReplConfig(s, &cfg)
			continue
		}
		// lines that contain but don't begin with the marker are replacements
		var i = strings.Index(s, cfg.mark)
		if i > 0 {
			var item = ReplItem{
				Find:     strings.Trim(s[:i], SPACES),
				Repl:     strings.Trim(s[i+len(cfg.mark):], SPACES),
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
	s = strings.ToUpper(s)
	keyword = strings.ToUpper(keyword)
	for i, ar := range [][]string{
		{"0", "FALSE", "OFF", "IGNORE"},
		{"1", "TRUE", "ON", "MATCH"},
	} {
		for _, match := range ar {
			if strings.HasPrefix(s, keyword+" "+match) {
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
						env.Print(strings.Repeat("\b", len(percent)), pc)
						percent = pc
					}
				}
			}
			newContent = zr.ReplaceMany(newContent, finds, repls, -1, cm, wm)
		}
		content = newContent
	}
	if ShowProgressIndicator {
		env.Print(strings.Repeat("\b", len(percent)))
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
	s = strings.Trim(s[len(cfg.mark):], SPACES)
	switch {
	case strings.HasPrefix(s, "path"):
		cfg.path = strings.Trim(s[5:], SPACES)
		env.Println("SET PATH:", cfg.path)
	//
	case strings.HasPrefix(s, "exts"):
		cfg.exts = strings.Fields(s[5:])
		env.Println("SET EXTS:", cfg.exts)
	//
	case strings.HasPrefix(s, "mark"):
		cfg.mark = strings.Trim(s[5:], SPACES)
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

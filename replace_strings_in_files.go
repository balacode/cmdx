// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-09 18:06:19 29EB0E             cmdx/[replace_strings_in_files.go]
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"path/filepath"
	"strings"
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
	cfg := replConfig{
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
		data, done := env.ReadFile(cfg.configFile)
		if !done {
			return
		}
		s := string(data)
		s = strings.TrimSpace(s)
		s = strings.Replace(s, "\r\n", "\n", -1)
		for strings.Contains(s, "\n\n") {
			s = strings.Replace(s, "\n\n", "\n", -1)
		}
		configLines = strings.Split(s, "\n")
		//
		// add a blank line to initiate replacement
		configLines = append(configLines, "")
	}
	//
	// each item:
	items := []ReplItem{}
	env.Println(strings.Repeat("-", 80))
	for lineNo, s := range configLines {
		s = strings.TrimSpace(s)
		//
		// blank lines initiate replacement:
		if s == "" && len(items) > 0 {
			env.Println(strings.Repeat("-", 80))
			var task sync.WaitGroup
			task.Add(1)
			cmd := ReplCmd{
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
		i := strings.Index(s, cfg.mark)
		if i > 0 {
			item := ReplItem{
				Find:     strings.TrimSpace(s[:i]),
				Repl:     strings.TrimSpace(s[i+len(cfg.mark):]),
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
	// TODO: you can remove configFile arg, and add an if condition in caller
	if task == nil {
		zr.Error("") // TODO: add error message (replaceAsync())
	}
	if task != nil {
		defer task.Done()
	}
	for _, filename := range env.GetFilePaths(cmd.Path, cmd.Exts...) {
		data, done := env.ReadFile(filename)
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
	oldContent := content
	max := 100 / float64(len(items))
	percent := ""
	newContent := content
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
					pc := fmt.Sprintf("c:%v w:%v %1.1f%%",
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
	s = strings.TrimSpace(s[len(cfg.mark):])
	switch {
	case strings.HasPrefix(s, "path"):
		cfg.path = strings.TrimSpace(s[5:])
		env.Println("SET PATH:", cfg.path)
	//
	case strings.HasPrefix(s, "exts"):
		cfg.exts = strings.Fields(s[5:])
		env.Println("SET EXTS:", cfg.exts)
	//
	case strings.HasPrefix(s, "mark"):
		cfg.mark = strings.TrimSpace(s[5:])
		if cfg.mark == "" {
			cfg.mark = DefaultMark
		}
		env.Println("SET MARK:", cfg.mark)
	//
	case hasBool(s, "case"):
		match, _ := getBool(s, "case")
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
		match, _ := getBool(s, "word")
		env.Println("SET WORD:", match)
		if match {
			cfg.wordMode = zr.MatchWord
		} else {
			cfg.wordMode = zr.IgnoreWord
		}
	}
} //                                                               setReplConfig

//end

// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2021-02-10 09:29:15 9A325C               cmdx/[replace_lines_in_files.go]
// -----------------------------------------------------------------------------

package main

import (
	"path/filepath"
	"strings"
)

// getConfigBool _ _
func getConfigBool(s, keyword string) (value, exists bool) {
	// TODO: apply this function in replaceStringsInFiles
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
} //                                                               getConfigBool

// hasConfigBool _ _
func hasConfigBool(s, keyword string) (ret bool) {
	// TODO: apply this function in replaceStringsInFiles
	_, ret = getConfigBool(s, keyword)
	return ret
} //                                                               hasConfigBool

// readConfigFileLines _ _
func readConfigFileLines(configFile string) (configLines []string) {
	// TODO: use this function in replaceStringsInFiles()
	// TODO: do away with this function, use env.ReadFileLines() instead
	var err error
	configFile, err = filepath.Abs(configFile)
	if err != nil {
		env.Println("command file path error: ", configFile)
		return
	}
	env.Println("FILE:", configFile)
	data, done := env.ReadFile(configFile)
	if !done {
		return []string{}
	}
	s := string(data)
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "\r\n", "\n", -1)
	for strings.Contains(s, "\n\n") {
		s = strings.Replace(s, "\n\n", "\n", -1)
	}
	configLines = strings.Split(s, "\n")
	configLines = append(configLines, "") // initiates replacement
	return configLines
} //                                                         readConfigFileLines

// end

// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                 cmdx/[replace_shared.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
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
	s = strings.ReplaceAll(s, "\r\n", "\n")
	for strings.Contains(s, "\n\n") {
		s = strings.ReplaceAll(s, "\n\n", "\n")
	}
	configLines = strings.Split(s, "\n")
	configLines = append(configLines, "") // initiates replacement
	return configLines
} //                                                         readConfigFileLines

// end

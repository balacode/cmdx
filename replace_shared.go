// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-28 13:59:12 2291C2               cmdx/[replace_lines_in_files.go]
// -----------------------------------------------------------------------------

package main

import (
	"path/filepath"
	str "strings"
)

// getConfigBool __
func getConfigBool(s, keyword string) (value, exists bool) {
	//TODO: apply this function in replaceStringsInFiles
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
} //                                                               getConfigBool

// hasConfigBool __
func hasConfigBool(s, keyword string) (ret bool) {
	//TODO: apply this function in replaceStringsInFiles
	_, ret = getConfigBool(s, keyword)
	return ret
} //                                                               hasConfigBool

// readConfigFileLines __
func readConfigFileLines(configFile string) (configLines []string) {
	//TODO: use this function in replaceStringsInFiles()
	//TODO: do away with this function, use env.ReadFileLines() instead
	var err error
	configFile, err = filepath.Abs(configFile)
	if err != nil {
		env.Println("command file path error: ", configFile)
		return
	}
	env.Println("FILE:", configFile)
	var data, done = env.ReadFile(configFile)
	if !done {
		return []string{}
	}
	var s = string(data)
	s = str.Trim(s, SPACES)
	s = str.Replace(s, CR+LF, LF, -1)
	for str.Contains(s, LF+LF) {
		s = str.Replace(s, LF+LF, LF, -1)
	}
	configLines = str.Split(s, LF)
	configLines = append(configLines, "") // initiates replacement
	return configLines
} //                                                         readConfigFileLines

//end

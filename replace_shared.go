// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-03-01 16:51:36 0F87F7               [cmdx/replace_lines_in_files.go]
// -----------------------------------------------------------------------------

package main

import "path/filepath" // standard
import "strings"       // standard

// getConfigBool __
func getConfigBool(s, keyword string) (value, exists bool) {
	//TODO: apply this function in replaceStringsInFiles
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
	s = strings.Trim(s, SPACES)
	s = strings.Replace(s, CR+LF, LF, -1)
	for strings.Contains(s, LF+LF) {
		s = strings.Replace(s, LF+LF, LF, -1)
	}
	configLines = strings.Split(s, LF)
	configLines = append(configLines, "") // initiates replacement
	return configLines
} //                                                         readConfigFileLines

//end

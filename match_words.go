// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 00:37:46 0135F9                          [cmdx/match_words.go]
// -----------------------------------------------------------------------------

package main

import str "strings" // standard

import "github.com/balacode/zr" // Zirconium

// matchWords __
func matchWords(cmd Command, args []string) {
	if len(args) != 2 {
		env.Println("requires <word-length> and <letter-set> parameters")
		return
	}
	var length = zr.Int(args[0])
	var letterSet = args[1]
	// load word list
	var words []string
	{
		var data, done = env.ReadFile(WordListFile)
		if !done {
			return
		}
		var s = string(data)
		s = str.Replace(s, "\r"+zr.LF, zr.LF, -1)
		for str.Contains(s, zr.LF+zr.LF) {
			s = str.Replace(s, zr.LF+zr.LF, zr.LF, -1)
		}
		words = str.Split(s, zr.LF)
	}
	for _, word := range words {
		if len(word) != length {
			continue
		}
		word = str.ToLower(word)
		var used = 0
		var letters = str.Split(letterSet, "")
		for _, wordLetter := range str.Split(word, "") {
			for i, letter := range letters {
				if wordLetter == letter {
					used++
					letters[i] = "0"
					break
				}
			}
			if used == length {
				env.Printf("    %s    ", word)
			}
		}
	}
} //                                                              matchWords

//end

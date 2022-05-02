// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                  cmdx/[words_in_file.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	"strings"
	"unicode"

	fs "github.com/balacode/zr-fs"
)

// wordsInFile _ _
// run cmdx with 'fw' or 'file-words'
// cx fw test_file.txt
//
// Either accepts 1 or 2 arguments.
// The first argument is the name of the input file.
//
// The second argument is the output file, where
// the list of unique words will be written.
func wordsInFile(cmd Command, args []string) {
	if len(args) < 1 || len(args) > 2 {
		env.Println(
			"requires <input-file> and optional <output-file> parameters",
		)
		return
	}
	var (
		filename = args[0]
		fragNo   = 0
		word     = [LongestWord]rune{}
		words    = make(map[string]int)
	)
	_ = fs.ReadFileChunks(filename, FileChunkSize+LongestWord,
		func(chunk []byte) int64 {
			fragNo++
			env.Print(" ", fragNo)
			// store words in map
			var (
				wordLen int
				hasA    bool
				hasD    bool
			)
			for _, ch := range string(chunk) {
				isA, isD := unicode.IsLetter(ch), unicode.IsDigit(ch)
				if isA {
					hasA = true
				}
				if isD {
					hasD = true
				}
				if wordLen < LongestWord && (ch == '_' || isA || isD) {
					word[wordLen] = ch
					wordLen++
					continue
				}
				if wordLen > 0 {
					if hasA && !hasD && wordLen < LongestWord {
						s := string(word[:wordLen])
						if n, exist := words[s]; exist {
							words[s] = n + 1
						} else {
							words[s] = 1
						}
					}
					hasA = false
					hasD = false
					wordLen = 0
				}
			}
			return int64(len(chunk))
		},
	)
	// read fragments from file, store words in map
	gap := strings.Repeat(" ", 10)
	for word, count := range words {
		env.Println(word, gap, count)
	}
}

// TODO: create Words() function in Zircon-Go lib

// TODO: create a text module or 'tstr'.

// end

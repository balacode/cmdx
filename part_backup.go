// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2021-02-10 09:29:15 F5DFA1                          cmdx/[part_backup.go]
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/balacode/zr"
)

// partBackup _ _
func partBackup(cmd Command, args []string) {
	fileSize := func(filename string) int64 {
		if !env.FileExists(filename) {
			return 0
		}
		info, err := os.Stat(filename)
		if err == nil {
			return info.Size()
		}
		return -1
	}
	const ChunkSize = 1 * 1024 * 1024 // 1MB
	chunk := make([]byte, ChunkSize)
	list := map[string]bool{}
	for {
		for _, filename := range env.GetFilePaths(".", "*.part") {
			_, exist := list[filename]
			if !exist {
				fmt.Println(filename, "ADDED")
				list[filename] = true
			}
		}
		for filename := range list {
			name := filename
			if !env.FileExists(name) {
				if !strings.HasSuffix(name, ".part") {
					continue
				}
				name = name[:len(name)-len(".part")]
			}
			if !env.FileExists(name) {
				delete(list, filename)
				continue
			}
			var (
				srcSize = fileSize(name)
				bakFile = name + ".bk"
				bakSize = fileSize(bakFile)
			)
			if bakSize >= srcSize {
				continue
			}
			// open the source file
			src, err := os.Open(name)
			if src == nil || err != nil {
				zr.Error("Failed opening", name, "due to:", err)
				continue
			}
			// open the target file
			bak, err := os.OpenFile(bakFile,
				os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0)
			if bak == nil || err != nil {
				zr.Error("Failed opening", bakFile, "due to:", err)
				continue
			}
			// repeatedly read chunks from the file
			hasErr := false
			pos := int64(bakSize)
			for {
				// advance the reading position
				_, err = src.Seek(pos, 0)
				if err != nil {
					zr.Error("Failed seeking", name, "due to:", err)
					hasErr = true
					break
				}
				// read the next chunk (nread = number of bytes read)
				nread, err := src.Read(chunk[:])
				if nread == 0 {
					break
				}
				if err != nil {
					zr.Error("Failed reading", name, "due to:", err)
					hasErr = true
					break
				}
				// advance the writing position
				_, err = bak.Seek(pos, 0)
				if err != nil {
					zr.Error("Failed seeking", bakFile, "due to:", err)
					hasErr = true
					break
				}
				// write the file
				nwrit, err := bak.Write(chunk[:nread])
				if nwrit != nread {
					zr.Error("Failed writing", bakFile,
						"read:", nread, "written:", nwrit)
					hasErr = true
					break
				}
				if err != nil {
					zr.Error("Failed writing", bakFile, "due to:", err)
					hasErr = true
					break
				}
				pos += int64(nwrit)
			}
			src.Close()
			bak.Sync()
			bak.Close()
			s := zr.ByteSizeString(pos, false)
			if hasErr {
				fmt.Println(filename, s, "FAILED")
			} else {
				fmt.Println(filename, s)
			}
		}
		time.Sleep(10 * time.Second)
	}
} //                                                                  partBackup

// end

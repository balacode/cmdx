// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                    cmdx/[part_backup.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
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
				// read the next chunk (nRead = number of bytes read)
				nRead, err := src.Read(chunk[:])
				if nRead == 0 {
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
				nWrit, err := bak.Write(chunk[:nRead])
				if nWrit != nRead {
					zr.Error("Failed writing", bakFile,
						"read:", nRead, "written:", nWrit)
					hasErr = true
					break
				}
				if err != nil {
					zr.Error("Failed writing", bakFile, "due to:", err)
					hasErr = true
					break
				}
				pos += int64(nWrit)
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
}

// end

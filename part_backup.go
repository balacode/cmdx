// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-12-26 05:37:09 BF2C5E                          cmdx/[part_backup.go]
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/balacode/zr"
)

// partBackup __
func partBackup(cmd Command, args []string) {
	var fileSize = func(filename string) int64 {
		info, err := os.Stat(filename)
		if err == nil {
			return info.Size()
		}
		return -1
	}
	for {
		for _, filename := range env.GetFilePaths(".", "*.part") {
			var size = fileSize(filename)
			var bak1 = filename + ".bak1"
			if fileSize(bak1) > fileSize(filename) {
				continue
			}
			fmt.Println(filename + " " + zr.ByteSizeString(size, false))
			var bak2 = filename + ".bak2"
			if env.FileExists(bak2) {
				env.DeleteFile(bak2)
			}
			if env.FileExists(bak1) {
				env.RenameFile(bak1, bak2)
			}
			copyFile(filename, bak1)
		}
		time.Sleep(10 * time.Second)
	}
} //                                                                  partBackup

// copyFile copies srcFileName to destFileName
func copyFile(srcFileName, destFileName string) {
	ar, err := ioutil.ReadFile(srcFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(destFileName, ar, 0644)
	if err != nil {
		fmt.Println("Error creating", destFileName)
		fmt.Println(err)
		return
	}
} //                                                                    copyFile

//end

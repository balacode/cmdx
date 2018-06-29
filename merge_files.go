// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-28 13:59:12 4E7CBD                          cmdx/[merge_files.go]
// -----------------------------------------------------------------------------

package main

import (
	"bytes"
	"path/filepath"
	str "strings"

	"github.com/balacode/zr-rgon"
)

// mergeFiles __
func mergeFiles(cmd Command, args []string) {
	// read parameters
	var from, to, mode string
	{
		var vars = []struct {
			name string
			val  *string
		}{
			{"source=", &from},
			{"target=", &to},
			{"mode=", &mode},
		}
		for _, arg := range args {
			for _, v := range vars {
				if str.HasPrefix(arg, v.name) {
					if *v.val != "" {
						env.Println("duplicate parameter: '" + v.name + "'")
						return
					}
					*v.val = arg[len(v.name):]
				}
			}
		}
		if from == "" {
			env.Println("'source=' not specified")
			return
		}
		if to == "" {
			env.Println("'target=' not specified")
			return
		}
	}
	// read the list of files in the mergefile
	var path = from
	if str.HasSuffix(path, ".rgon") {
		from = filepath.Dir(from)
	} else {
		if !str.HasSuffix(path, env.PathSeparator()) {
			path += env.PathSeparator()
		}
		path += "merge.rgon"
	}
	var data, done = env.ReadFile(path)
	if !done {
		return
	}
	var def, err = rgon.Parse(string(data))
	if err != nil {
		env.Println("Failed parsing", path, "due to:", err)
		return
	}
	var files = def.Objs("files")
	//
	// store the files in a memory buffer
	var buf bytes.Buffer
	for _, iter := range files {
		var filename = iter.Str("file")
		var fileMode = iter.Str("mode")
		if fileMode != "" && fileMode != mode {
			continue
		}
		var path = from + env.PathSeparator() + filename
		path = filepath.Clean(path)
		var data, done = env.ReadFile(path)
		if !done {
			return
		}
		buf.Write(data)
	}
	// write the buffer to the 'to' file
	env.WriteFile(to, buf.Bytes())
} //                                                                  mergeFiles

//end

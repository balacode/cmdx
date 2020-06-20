// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2020-06-20 09:58:17 405968                          cmdx/[merge_files.go]
// -----------------------------------------------------------------------------

package main

import (
	"bytes"
	"path/filepath"
	"strings"

	rgon "github.com/balacode/go-rgon"
)

// mergeFiles _ _
func mergeFiles(cmd Command, args []string) {
	// read parameters
	var from, to, mode string
	{
		vars := []struct {
			name string
			val  *string
		}{
			{"source=", &from},
			{"target=", &to},
			{"mode=", &mode},
		}
		for _, arg := range args {
			for _, v := range vars {
				if strings.HasPrefix(arg, v.name) {
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
	path := from
	if strings.HasSuffix(path, ".rgon") {
		from = filepath.Dir(from)
	} else {
		if !strings.HasSuffix(path, env.PathSeparator()) {
			path += env.PathSeparator()
		}
		path += "merge.rgon"
	}
	data, done := env.ReadFile(path)
	if !done {
		return
	}
	def, err := rgon.Parse(string(data))
	if err != nil {
		env.Println("Failed parsing", path, "due to:", err)
		return
	}
	files := def.Objs("files")
	//
	// store the files in a memory buffer
	var buf bytes.Buffer
	for _, iter := range files {
		var (
			filename = iter.Str("file")
			fileMode = iter.Str("mode")
		)
		if fileMode != "" && fileMode != mode {
			continue
		}
		path := from + env.PathSeparator() + filename
		path = filepath.Clean(path)
		data, done := env.ReadFile(path)
		if !done {
			return
		}
		buf.Write(data)
	}
	// write the buffer to the 'to' file
	env.WriteFile(to, buf.Bytes())
} //                                                                  mergeFiles

//end

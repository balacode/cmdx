// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-22 15:26:42 782246                          [cmdx/merge_files.go]
// -----------------------------------------------------------------------------

package main

import "bytes"         // standard
import "path/filepath" // standard
import str "strings"   // standard

import "ase/zr/ters" // Zirconium

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
	if str.HasSuffix(path, ".ters") {
		from = filepath.Dir(from)
	} else {
		if !str.HasSuffix(path, env.PathSeparator()) {
			path += env.PathSeparator()
		}
		path += "merge.ters"
	}
	var data, done = env.ReadFile(path)
	if !done {
		return
	}
	var def, err = ters.Parse(string(data))
	if err != nil {
		env.Println("Failed parsing", path, "due to:", err)
		return
	}
	var files = def.GetItems("files")
	//
	// store the files in a memory buffer
	var buf bytes.Buffer
	for _, iter := range files {
		var filename = iter.Get("file")
		var fileMode = iter.Get("mode")
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

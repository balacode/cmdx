// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 00:37:46 EE7B2B                    [cmdx/loaded_text_files.go]
// -----------------------------------------------------------------------------

package main

import "os"   // standard
import "sort" // standard
import "sync" // standard

import "github.com/balacode/zr" // Zirconium

// LoadedTextFiles __
type LoadedTextFiles struct {
	m map[string]*TextFile
} //                                                             LoadedTextFiles

// GetFile returns the named file stored in LoadedTextFiles
func (ob *LoadedTextFiles) GetFile(filename string) *TextFile {
	if ob.m == nil {
		env.Error("No loaded files.")
		return &TextFile{}
	}
	return ob.m[filename]
} //                                                                     GetFile

// GetAllFilenames returns a list of all loaded file names
func (ob *LoadedTextFiles) GetAllFilenames() []string {
	if ob.m == nil {
		env.Error("No loaded files.")
		return []string{}
	}
	var ret []string
	for key := range ob.m {
		ret = append(ret, key)
	}
	return ret
} //                                                             GetAllFilenames

// LoadAll loads all text files in 'path' into a map in memory.
// On repeated calls it reloads only files that changed after the last call.
// Returns a list of added or changed files (but not deleted files).
func (ob *LoadedTextFiles) LoadAll(
	path string,
	fsMx *sync.RWMutex,
) (
	changedFiles []string,
) {
	var paths = env.GetFilePaths(path, env.TextFileExts()...)
	for _, filename := range paths {
		var _, changed = ob.LoadFile(filename, fsMx)
		if changed {
			changedFiles = append(changedFiles, filename)
		}
	}
	ob.SortListByModTime(changedFiles)
	return changedFiles
} //                                                                     LoadAll

// LoadFile reads a file into the map of loaded files.
// The contents of the file are retained in memory.
func (ob *LoadedTextFiles) LoadFile(
	filename string,
	fsMx *sync.RWMutex,
) (
	file *TextFile,
	changed bool,
) {
	fsMx.Lock()
	defer fsMx.Unlock()
	if ob.m == nil {
		ob.m = make(map[string]*TextFile, 1000)
	}
	var inMap bool
	file, inMap = ob.m[filename]
	if !inMap {
		file = &TextFile{Path: filename}
		ob.m[filename] = file
	}
	if !env.FileExists(filename) {
		delete(ob.m, filename)
		return nil, false
	}
	var info, err = os.Stat(filename)
	if err != nil {
		zr.Error("Stat failed on file", filename)
		delete(ob.m, filename)
		return nil, false
	}
	if file.Size == info.Size() && file.ModTime == info.ModTime() {
		return file, false
	}
	file.Size = info.Size()
	file.ModTime = info.ModTime()
	file.Lines = env.ReadFileLines(file.Path)
	return file, true
} //                                                                    LoadFile

// SortListByModTime sorts a list of file names by time modified.
// 'modFiles' is the slice of file names being sorted.
func (ob *LoadedTextFiles) SortListByModTime(modFiles []string) {
	sort.Slice(modFiles, func(i, j int) bool {
		var a = ob.m[modFiles[i]]
		var b = ob.m[modFiles[j]]
		//
		// this should never occur
		if a == nil || b == nil {
			zr.Error(zr.ENil)
			return true // push front
		}
		return a.ModTime.After(b.ModTime)
	})
} //                                                          SortListByModTime

//TODO: GLOBAL: All doc. comments should not exceed 76 columns

//end
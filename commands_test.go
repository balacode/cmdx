// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 00:37:46 F72496                        [cmdx/commands_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in commands.go use:
    go test --run Test_cmds_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "reflect"     // standard
import "testing"     // standard
import str "strings" // standard

import "github.com/balacode/zr" // Zirconium

// -----------------------------------------------------------------------------

// go test --run Test_cmds_consts_
func Test_cmds_consts_(t *testing.T) {
	zr.TEqual(t, BB, ("BB"))
	zr.TEqual(t, BC, ("BC"))
	zr.TEqual(t, BE, ("BE"))
	zr.TEqual(t, CB, ("CB"))
	zr.TEqual(t, CLL, ("CLL"))
	zr.TEqual(t, EB, ("EB"))
	zr.TEqual(t, LT, ("<"))
	zr.TEqual(t, T, ("T"))
	zr.TEqual(t, XE, ("XE"))
} //                                                           Test_cmds_consts_

// -----------------------------------------------------------------------------

// go test --run Test_cmds_AllCategories_
func Test_cmds_AllCategories_(t *testing.T) {
	var categories = make(map[string]bool, len(AllCategories))
	for key, cat := range AllCategories {
		if key < 1 || key > 3 {
			t.Error("Invalid category key:", key)
		}
		var exist bool
		//
		// category must be consistent and not zero-length
		if len(str.Trim(cat, zr.SPACES)) == 0 {
			t.Error("Invalid category: '" + cat + "'")
		}
		// category must be unique
		_, exist = categories[cat]
		if exist {
			t.Error("Non-unique category '" + cat + "'")
		}
		categories[cat] = true
	}
} //                                                    Test_cmds_AllCategories_

// -----------------------------------------------------------------------------

// go test --run Test_cmds_AllCommands_
func Test_cmds_AllCommands_(t *testing.T) {
	var shortNames = make(map[string]bool, len(AllCommands))
	var fullNames = make(map[string]bool, len(AllCommands))
	var handlers = make(map[uintptr]bool, len(AllCommands))
	var isValidName = func(s string) bool {
		for _, ch := range s {
			if (ch < 'a' || ch > 'z') && ch != '-' {
				return false
			}
		}
		return true
	}
	var addressOf = func(value interface{}) uintptr {
		var ret uintptr
		var refVal = reflect.ValueOf(value)
		switch refVal.Kind() {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr,
			reflect.Slice, reflect.UnsafePointer:
			ret = refVal.Pointer()
		}
		return ret
	}
	for _, cmd := range AllCommands {
		// ShortName must be consistent and not zero-length
		if !isValidName(cmd.ShortName) {
			t.Error("Invalid 'ShortName':", zr.DescribeStruct(&cmd))
		}
		var exist bool
		// ShortName must be unique
		_, exist = shortNames[cmd.ShortName]
		if exist {
			t.Error(
				"Non-unique ShortName '"+cmd.ShortName+"' in",
				zr.DescribeStruct(&cmd),
			)
		}
		shortNames[cmd.ShortName] = true
		//
		// FullName must be consistent and not zero-length
		if !isValidName(cmd.FullName) {
			t.Error("Invalid 'FullName':", zr.DescribeStruct(&cmd))
		}
		// FullName must be unique
		_, exist = fullNames[cmd.FullName]
		if exist {
			t.Error(
				"Non-unique FullName '"+cmd.FullName+"' in",
				zr.DescribeStruct(&cmd),
			)
		}
		fullNames[cmd.FullName] = true
		//
		// Handler must not be nil
		if cmd.Handler == nil {
			t.Error("Handler must not be <nil> in", zr.DescribeStruct(&cmd))
		}
		// Handler must be unique
		var addr = addressOf(cmd.Handler)
		_, exist = handlers[addr]
		if exist {
			t.Error("Non-unique Handler in", zr.DescribeStruct(&cmd))
		}
		handlers[addr] = true
	}
} //                                                      Test_cmds_AllCommands_

//end

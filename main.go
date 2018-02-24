// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 00:37:46 520080                                 [cmdx/main.go]
// -----------------------------------------------------------------------------

package main

import "fmt"

import "os"          // standard
import str "strings" // standard

import "github.com/balacode/zr" // Zirconium

// -----------------------------------------------------------------------------
// # Main Function

// main function: execution begins here
func main() {
	var args = os.Args
	if len(args) == 1 { //        show help if no command-line options specified
		const Format = "%-3s  %-11s  %-65s" //               layout and headings
		var Div = str.Repeat("-", 80)
		env.Print(
			Div, zr.LF,
			fmt.Sprintf(Format, "AB.", "FULL", "DESCRIPTION OF COMMAND"), zr.LF,
			Div, zr.LF,
		)
		for cat := 1; cat <= 3; cat++ { //            group commands by category
			env.Print(zr.LF, str.ToUpper(AllCategories[cat])+":", zr.LF)
			for _, cmd := range AllCommands {
				if cmd.Category == cat {
					env.Println(fmt.Sprintf(Format,
						cmd.ShortName, cmd.FullName, cmd.ShortInfo))
				}
			}
		}
		env.Println(Div)
		return
	}
	var cmdName = str.ToLower(args[1]) // locate command to run in AllCommands
	args = args[2:]                    // remove program [0] & command name [1]
	for _, cmd := range AllCommands {
		if cmdName == cmd.ShortName || cmdName == cmd.FullName {
			cmd.Handler(cmd, args)
			return
		}
	}
	env.Printf("unknown command: %s", cmdName)
} //                                                                        main

//TODO: create a core package

//end

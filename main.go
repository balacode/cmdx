// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                           cmdx/[main.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"os"
	"strings"
)

const Version = "v1.1.2"

// -----------------------------------------------------------------------------
// # Main Function

// main function: execution begins here
func main() {
	args := os.Args
	if len(args) == 1 { //        show help if no command-line options specified
		const Format = "%-3s  %-11s  %-65s" //               layout and headings
		Div := strings.Repeat("-", 80)
		env.Print(
			Div, "\n",
			"CMDX Utility "+Version, "\n",
			"\n",
			fmt.Sprintf(Format, "AB.", "FULL", "DESCRIPTION OF COMMAND"), "\n",
			Div, "\n",
		)
		for cat := 1; cat <= 3; cat++ { //            group commands by category
			env.Print("\n", strings.ToUpper(AllCategories[cat])+":\n")
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
	// locate command to run in AllCommands
	cmdName := strings.ToLower(args[1])
	//
	// remove program [0] & command name [1]
	args = args[2:]
	for _, cmd := range AllCommands {
		if cmdName == cmd.ShortName || cmdName == cmd.FullName {
			cmd.Handler(cmd, args)
			return
		}
	}
	env.Printf("unknown command: %s", cmdName)
} //                                                                        main

// TODO: create a core package

// end

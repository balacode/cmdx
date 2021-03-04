// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                      cmdx/[hardcoded.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

// TODO: all the paths in this file will be written in config files later

const (
	hardcodedDefaultLibPathOnLinux = `/x/user/projects/code/go/src/base`

	hardcodedDefaultLibPathOnWindows = `X:\user\projects\code\go\src\base`

	hardcodedRootPath = `X:\user\projects\code\go\src\base`

	hardcodedWordListFile = `c:\__DEBDESK\words_alpha.txt`
)

var hardcodedAutologFiles = []string{
	`X:\user\projects\code\android\autotime.log`,
	`X:\user\projects\code\go\src\base\autotime.log`,
	`X:\user\projects\code\go\src\base\autotime_2017.txt`,
	`X:\user\projects\code\go\src\base\autotime_2018.txt`,
	`X:\user\projects\code\go\src\base\client\maqua\autotime.log`,
	`X:\user\projects\code\go\src\base\client\maqua\autotime_2017.txt`,
	`X:\user\projects\code\go\src\base\client\maqua\autotime_2018.txt`,
	`X:\user\projects\code\go\src\base\client\mnlhq\autotime.log`,
	`X:\user\projects\code\go\src\base\client\mnlhq\autotime_2017.txt`,
	`X:\user\projects\code\go\src\base\client\mnlhq\autotime_2018.txt`,
}

var hardcodedManualLogFiles = []string{
	`X:\user\projects\code\go\src\base\client\maqua\timelog.txt`,
	`X:\user\projects\code\go\src\base\client\mnlhq\timelog.txt`,
	`X:\user\projects\code\go\src\base\timelog.txt`,
}

var hardcodedTimeLogPaths = []string{
	`X:\user\projects\code\android\Shila`,
	`X:\user\projects\code\android\TheApp`,
	`X:\user\projects\code\go\src\base\client\dmd_app`,
	`X:\user\projects\code\go\src\base\client\dmo_app`,
	`X:\user\projects\code\go\src\base\client\gsmmkt`,
	`X:\user\projects\code\go\src\base\client\maqua`,
	`X:\user\projects\code\go\src\base\client\mnlhq`,
	`X:\user\projects\code\go\src\base\client\s11_app`,
	`X:\user\projects\code\go\src\base\client\sul_app`,
	`X:\user\projects\code\go\src\base\client\tlg`,
	`X:\user\projects\code\go\src\base_defer\gosql`,
	`X:\user\projects\code\go\src\base`,
	`X:\user\projects\code\go\src`,
}

var hardcodedIgnoreFilenamesWith = []string{
	".css",
	".idea",
	".log",
	".tmp",
	"__",
	"_app.js",
	"_repl_lines.txt",
	"_repl_strs.txt",
	"_zr.js",
	"`",
	"tmp.",
}

// end

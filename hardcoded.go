// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-06-15 14:55:18 A20927                            cmdx/[hardcoded.go]
// -----------------------------------------------------------------------------

package main

// TODO: all the paths in this file will be written in config files later

var hardcodedAutologFiles = []string{
	`X:\user\projects\code\android\autotime.log`,
	`X:\user\projects\code\go\src\base\autotime.log`,
	`X:\user\projects\code\go\src\base\autotime_2017.txt`,
	`X:\user\projects\code\go\src\base\autotime_2018.txt`,
	`X:\user\projects\code\go\src\base\client\mnl\autotime.log`,
	`X:\user\projects\code\go\src\base\client\mnl\autotime_2017.txt`,
	`X:\user\projects\code\go\src\base\client\mnl\autotime_2018.txt`,
}

var hardcodedManualLogFiles = []string{
	`X:\user\projects\code\go\src\base\timelog.txt`,
	`X:\user\projects\code\go\src\base\client\mnl\timelog.txt`,
}

const hardcodedDefaultLibPathOnLinux = `/x/user/projects/code/go/src/base`

const hardcodedDefaultLibPathOnWindows = `X:\user\projects\code\go\src\base`

const hardcodedRootPath = `X:\user\projects\code\go\src\base`

var hardcodedTimeLogPaths = []string{
	`X:\user\projects\code\android\autotime.log`,
	`X:\user\projects\code\go\src\base\client\dmd_app`,
	`X:\user\projects\code\go\src\base\client\dmd_label`,
	`X:\user\projects\code\go\src\base\client\dmd_spart`,
	`X:\user\projects\code\go\src\base\client\dmo_app`,
	`X:\user\projects\code\go\src\base\client\mnl`,
	`X:\user\projects\code\go\src\base\client\tlg`,
	`X:\user\projects\code\go\src\base_defer\gosql`,
	`X:\user\projects\code\go\src\base`,
	`X:\user\projects\code\go\src`,
}

var hardcodedIgnoreFilenamesWith = []string{
	".log",
	".tmp",
	"/_",
	"`" + "`",
	"tmp.",
	`/.idea/`,
	`\_`,
}

const hardcodedWordListFile = `c:\__DEBDESK\words_alpha.txt`

var hardcodedReplaceManyPath = `X:\user\projects\code\go\src\base\zr\ReplaceMany`

//end

# CMDX Tool
A command line tool to manage files and process source code.

This tool is a sort of Swiss-Army-Knife for managing files and text processing.
I wanted to avoid having too many little command line utilities,
so I created CMDX which means Command-line Extensions.

## Installation:
Use `go get` to install the utility and its dependencies. (If you use `git clone`, you'll have to also clone `zr` and `zr_fs` manually and make sure they're in the github.com/balacode branch.)

    go get github.com/balacode/cmdx

This will also install *github.com/balacode/zr* and *github.com/balacode/zr_fs* libraries (Zircon-Go) which CMDX uses. Once installed, change to `{go src path}/github.com/balacode/cmdx` and type `go build`.

## Usage Guide:

*I am in the process of adding a detailed explanation for the usage of each command. In the meantime, you can check the source code or ask me directly. Your questions will help make the documentation more concise.*

Below is a brief summary of the available commands.

## File Manipulation Commands:

**DD (del-dup): Delete Duplicates**

    cmdx dd /source /target

Deletes all files in /target that already exist in /source.
The command does not care about file names.
Instead it compares the contents of files with identical sizes.

**Use this command with care: files are deleted without going to trash can** :exclamation:

**LD (list-dup): List Duplicate Files**
Lists duplicate files in the specified folder (or source and target folders).
Does not delete or change any files.

To list all identical files in /source folder:

    cmdx ld /source

To list list all files in */target* that have identical files in */source*:

    cmdx ld /source /target

**RD (ren-dup): Rename Duplicate Files**

Given two folder paths: */source* and */target*,
this command finds files in /target that are identical to files
in /source and renames them to their file names in /source:

    cmdx rd /source /target

**RH (ren-hash): Rename-Hash**
Renames files by prefixing their name with a hash.

## Text Manipulation Commands:

**FW (file-words):**
Lists all words with alphanumeric characters from {file}:

    cmdx fw {file}

**ME (mark-errors):**
Inserts build errors as comments at the source of the error,
so you don't need to manually look-up the line numbers and
file names to edit. Just use your editor to find all error
markers, fix the error, and delete the comment when the error
is fixed.

To make this command work, send the output of the
`go build` or `go install` command to a build log file:

    go build -gcflags="-e" 2> build.log

The `=gcflags="-e"` option instructs the Go compiler to
output all build errors. Without this option, the compiler
stops reporting errors after about 10 errors.

Next, run the mark-errors command with the name of the build log:

    cmdx mark-errors -buildlog=.\build.log

After you run the command, you will see comments such as the following:

```go
    func main() {
        var args = os.Args
        if len(args) = 1 {
                   //^ syntax error: len(args) = 1 used as value
        }
    ...
```

**MT (mark-time):**
Changes timestamps in source files.
Requires paths (in hardcoded.go) to be set up.

**RL (rep-lines):**
Replaces lines in file(s). Requires {command-file}.
This command allows you to replace several blocks of code at once.

**RS (replace-strings):**
Makes multiple (different) replacements simultaneously in multiple files.
You can make thousands of simultaneous replacements as the command
uses goroutines to search and replace multiple files concurrently
once they are loaded in RAM.

    cmdx rs replacements.repl

The path, the types of files and the replacements are
specified in a replacements file. Example replacements file:

    mark ~~
    path X:\path
    case on
    word on

    ~~ comment

    find1 ~~ replace1
    find2 ~~ replace2
    for (var i = 0; i < 10; i++) {  ~~  for i := 0; i < 10; i++ {

- mark: the delimiter to denote comments and separate seach and replacement text.
- path: the replacement path (also replaces subfolders). You can only specify one path, for now.
- case: set 'on' to match case, or 'off' to ignore case.
- word: set 'on' to match whole words, or 'off' to replace substrings.
- comments start with the marker.
- text to find is on the left of '~~' and the replacement text on the right. You can list as many replacements as needed.

**SF (sort-file):**

    cmdx sf {filename}

Sorts all lines in a file and deletes duplicate lines.
This command is useful for sorting log files, dictionary lists, etc.

## Other Commands:
More specifics on these commands will be provided later.

**RT (rep-time):**
Replaces time entries in log files.

**RI (run):**
Runs the tool in source-code interactive mode.

**TR (time-report):**
Summarizes time from log files and presents it in a calendar format.
For example:

    2018 FEBRUARY
    *--------------------------------------------------------------*
    |  Mon   |  Tue   |  Wed   |  Thu   |  Fri   |  Sat   |  Sun   |
    |--------|--------|--------|--------|--------|--------|--------|
    |        |        |        | 1      | 2      | 3      | 4      |
    |        |        |        |   8.44 |   7.55 |   6.66 |   5.77 |
    |--------|--------|--------|--------|--------|--------|--------|
    | 5      | 6      | 7      | 8      | 9      | 10     | 11     |
    |   4.88 |   3.99 |   2.15 |   1.54 |      0 |      1 |      2 |
    |--------|--------|--------|--------|--------|--------|--------|
    | 12     | 13     | 14     | 15     | 16     | 17     | 18     |
    |      3 |      4 |      5 |      6 |      7 |      8 |      9 |
    |--------|--------|--------|--------|--------|--------|--------|
    | 19     | 20     | 21     | 22     | 23     | 24     | 25     |
    |     10 |        |        |        |        |        |        |
    |--------|--------|--------|--------|--------|--------|--------|
    | 26     | 27     | 28     |        |        |        |        |
    |        |        |        |        |        |        |        |
    |--------|--------|--------|--------|--------|--------|--------|
    |        |        |        |        |        |        |        |
    |        |        |        |        |        |        |        |
    *--------------------------------------------------------------*
    95.98

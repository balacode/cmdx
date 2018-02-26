# cmdx :nut_and_bolt:
A command line tool to manage files and process source code.

This tool is a sort of Swiss-Army-Knife for managing files and text processing.
I wanted to avoid having too many little command line utilities,
so I created CMDX which means Command-line Extensions.

Below is a brief summary of the available commands. 

I am in the process of adding a detailed explanation for the usage of each command. In the meantime, you can check the source code or ask me directly. Your questions will help make the documentation more concise.

## File Manipulation Commands:

**DD (del-dup): Delete Duplicates**  

    cmdx dd /source /target  

Deletes all files in /target that already exist in /source.  
The command does not care about file names.  
Instead it compares the contents of files with identical sizes.  

**Use this command with care: files are deleted at once, without going to trash can** :exclamation:

**LD (list-dup): List Duplicate Files**  
Lists duplicate files in the specified folder (or source and target folders).  
Does not delete or change any files.  

    cmdx ld /source

Lists all identical files in /source folder.  

cmdx ld /source /target  
Lists all files in /target that have identical files in /source.  

**RD (ren-dup): Rename Duplicate Files**  

    cmdx rd /source /target

Given two folders: /source and /target, this command  
finds files in /target that are identical to files in  
/source and renames them to their file names in /source.  

**RH (ren-hash): Rename-Hash**  
Renames files by prefixing their name with a hash.  

## Text Manipulation Commands:

**FW (file-words):**   
Lists all words with alphanumeric characters from {file}.

**ME (mark-errors):**   
Inserts build errors as comments at the source of the error,
so you don't need to manually look-up the line numbers and
file names to edit. Just use your editor to find all error
markers, fix the error, and delete the comment when the error
is fixed.

To make this command work, send the output of the
`go build` or `go install` command to a build log file:

    go build -gcflags="-e" 2> build.log

Then, run the mark-errors command with the name of the log:

    cmdx mark-errors -buildlog=.\build.log

**MT (mark-time):**  
Changes timestamps in source files.  
Requires paths (in hardcoded.go) to be set up.  

**RL (rep-lines):**  
Replaces lines in file(s). Requires {command-file}.  
This command allows you to replace several blocks of code at once.  

**RS (replace-strings):**  
Replaces strings in file(s). Requires {command-file}.

**SF (sort-file):**  

    cmdx sf {filename}

Sorts all the lines in a file.  
And makes sure each line is unique.  

## Other Commands:  
More specifics on these commands will be provided later.  

**RT (rep-time):**  
Replaces time entries in log files.  

**RI (run):**  
Runs the tool in source-code interactive mode.  

**TR (time-report):**  
Summarizes time from log files and presents it in a calendar format.  

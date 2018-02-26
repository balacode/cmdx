# cmdx
Multi-purpose command line tool to manage files and process source code

This tool is a sort of Swiss-Army-Knife for managing files and text processing.
I wanted to avoid having too many little command line utilities,
so I created CMDX which means Command-line Extensions.

## File Manipulation Commands:

**DD (del-dup): Delete Duplicates**  
cmdx dd /source /target  

**LD (list-dup): List Duplicate Files**  
Helps you to find duplicate files.  

cmdx dd /source  
Lists all identical files in /source folder.  

cmdx dd /source /target  
Lists all files in /target that have identical files in /source.  

**RD: Rename Duplicate Files**  
cmdx rd /source /target  
Given two folders: /source and /target, this command  
finds files in /target that are identical to files in  
/source and renames them to their file names in /source.  

**RH: Rename-Hash**  
Renames files by prefixing their name with a hash.  

## Text Manipulation Commands:

**FW (file-words):**   
Lists all words with alphanumeric characters from <file>.

**ME (mark-errors):**   
Inserts build errors as comments at the source of the error,
so you don't need to manually look-up the line numbers and
file names to edit. Just use your editor to find all error
markers, fix the error, and delete the comment when the error
is fixed.

**MT (mark-time):**  
Changes timestamps in source files. Requires path

**RL (rep-lines):**  
Replaces lines in file(s). Requires <command-file>

**RS (replace-strings):**  
Replaces strings in file(s). Requires <command-file>

**SF (sort-file):**  
cmdx sf <filename>
Sorts all the lines in a file.
And makes sure each line is unique.

## Other Commands:  
More specifics on these commands will be provided later.  

**RT (rep-time)**  
Replaces time entries in log files.  

**RI (run)**  
Runs the tool in source-code interactive mode.  

**TR (time-report)**  
Summarizes time from log files and presents it in a calendar format.  

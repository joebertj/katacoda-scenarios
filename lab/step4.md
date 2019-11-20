# Text Processing and Scripting Commands

## info

The awk general-purpose text processing language

Example 

To count the lines in a file.

`awk 'END { print NR }' meh`{{execute}}

### basename

Strips extensions and directories from filenames

Example

`awk 'END { print NR }' meh`{{execute}}

#### cat

Displays and concatenates files

##### cmp

Compares binary files

###### cut

Extracts columns of lines

Example

`ps -ef | cut -f8`{{execute}}

####### diff

Compares text files

######## dirname

Extracts the directory from a filename

######### echo

Prints text

########## ed



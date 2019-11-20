# Text Processing and Scripting Commands

## awk

The awk general-purpose text processing language

Example 

`ps -ef | awk '{print $8}'`{{execute}}

## basename

Strips extensions and directories from filenames

Example


## cat

Displays and concatenates files

## cmp

Compares binary files

## cut

Extracts columns of lines

Example

`ps -ef | cut -f8`{{execute}}

## diff

Compares text files



## dirname

Extracts the directory from a filename

Example 

`ls $DIR`{{execute}}

## echo

Prints text

Example 

`echo "1,2,3" | cut -d ',' -f2 `{{execute}}

## ed



##bunzip2

A command is a block-sorting file compressor that can be used to both compress and decompress data, depending on how you use it.

`bunzip2 /katacoda-scenarios/lab/sample.txt.bz2`{{execute}}


##bzip2

Is a command line based file compressor in Linux that uses the Burrows-Wheeler block sorting text compression algorithm and Huffman coding to carry out the compression process.

`bzip2 sample.txt`{{execute}}


##cpio

Stands for /"copy in, copy out". It is used for processing the archive files like *.cpio or *.tar. This command can copy files to and from archicves.

`cpio -o < name-list > archive`{{execute}}
`cpio -i < archive`{{execute}}


##gunzip

Is used to compress or expand a file or a list of file.

`gunzip sample.txt`{{execute}}


##gzip

This command compresses files. Each single files is compressed into a single file. The compressed file consists of a GNU zip header and deflated data.

`gzip sample.txt`{{execute}}
`gzip -k sample.txt`{{execute}}
`gzip -L sample.txt`{{execute}}
`gzip -f sample.txt`{{execute}}


##tar

Stands for tape archive, is used to create Achive and extract the Archive files. tar command in linux is one of the important command which provides archiving functionality in Linux.

`tar cvf sample.tar *.c`{{execute}}
`tar xvf sample.tar`{{execute}}


uncompress

A decompression program

unshar

A de-archival program

uudecode

A decoding program (the counterpart of uuencode)

uuencode

Encodes binary file into a text file

zcat

Decompresses into a file stream


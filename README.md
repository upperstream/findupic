# Image Duplicate Finder

This Go program scans a specified directory and its subdirectories to
find image files that have identical decoded bitmaps, and writes
the output to the standard output stream.

## Usage

```
findupic [ -csv ] [ -error-log=<logfile> ] <directory>...
```

where _directory_ is the path to the directory to scan for image files.

If `-csv` flag is added, this program prints the report in CSV format
to the standard output stream.  If `-error-log` flag is added, this
program writes the errors into the error log file named _logfile_.

This program does not stop when it finds an erroneous image file, but
writes its file name into the error log file and continues finding
duplicates.

This program returns `1` as an error exit status when it detects any
erroneous files, while returns `0` as the success exit status.

## Example

```
findupic /path/to/images
```

This command will scan the `/path/to/images` directory and descendants
for image files, and write the output to the standard output stream.

Two or more directories can be specified:

```
findupic dir1 dir2
```

This command will scan two directories `dir1` and `dir2`, as well as
their descendants for images files.  Duplicates will be reported
regardless of which directory tree the image file resides in.

## Output

The program writes the paths of duplicate image files to the standard
output stream following the SHA-256 hash of the decoded bitmap:

```
$ ./findupic.exe .
Duplicate images with hash c22c...:
dir1/image_B2.jpg
image_B1.jpg

Duplicate images with hash 92e4...:
dir1/image_A3.jpg
image_A1.jpg
image_A2.jpg

```

In CSV output mode, the same report will look like:

```
SHA256,Path
c22c...,dir1\image_A3.jpg
c22c...,image_A1.jpg
c22c...,image_A2.jpg
92e4...,dir1\image_B2.jpg
92e4...,dir2\image_B1.jpg
```

Note that in either mode SHA256 will be printed in full length.

## Build

In this project directory, execute the following command to compile
this program and generate an executable file:

```
go build
```

You need [Go](https://go.dev/) tool chain
[installed](https://go.dev/doc/install).

## Disclaimer

This program is generated using an AI language model and is intended
for educational purposes only. Since the majority of the work, even
including this document, is written by the AI language model and has
been modified and adapted for this specific task; the author of this
program does not claim copyright ownership for any part of the program.
The program is provided as-is, without any warranty or guarantee of its
fitness for any particular purpose. Use of the program is at your own
risk.

## Licensing

This program is released under the Unlicense. You are free to use,
modify, and distribute the program without restriction. The program is
provided as-is, without any warranty or guarantee of its fitness for
any particular purpose. Please see the [`UNLICENSE`](UNLICENSE.txt)
file for the full text of the Unlicense.

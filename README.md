# Image Duplicate Finder

This Go program scans a specified directory and its subdirectories to
find image files that have identical decoded bitmaps, and writes
the output to the standard output stream.

## Usage

```
findupic <directory>...
```

where _directory_ is the path to the directory to scan for image files.

## Example

```
findupic /path/to/images
```

This command will scan the `/path/to/images` directory and descendents
for image files, and write the output to the standard output stream.

Two or more directories can be specified:

```
findupic dir1 dir2
```

This command will scan two directories `dir1` and `dir2`, as well as
their descendents for images files.  Duplicates will be reported
regardless of which directory tree the image file resides in.

## Output

The program writes the paths of duplicate image files to the standard
output stream follwing the SHA-256 hash of the decoded bitmap:

```
$ ./findupic.exe .
Duplicate images with hash c22c....:
dir1/image_B2.jpg
image_B1.jpg

Duplicate images with hash 92e4...:
dir1/image_A3.jpg
image_A1.jpg
image_A2.jpg

```

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

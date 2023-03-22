/*
Package findupic implements a command-line tool that searches for
duplicate images within one or more directories. The tool finds images
that have identical SHA256 hashes, indicating that their decoded
bitmaps are identical.

# Usage

	findupic [OPTIONS] <directory>...

The findupic command accepts one or more directories as arguments. For
example:

	$ findupic ~/Pictures /mnt/media/photos

The tool searches each directory and its subdirectories for image files
with the extensions .jpg, .jpeg, .png, and .gif. For each image file,
the tool computes SHA256 hash of its decoded bitmap and adds the path
to a list of images with that hash. If the tool finds multiple images
with the same hash, it reports those images as duplicates.

# Options

	-csv             Enable CSV output
	-error-log FILE  Write error messages to FILE instead of stderr
	-h, --help       Show this help message and exit

Example usage:

	# Find duplicate images in the current directory and write the
	# results to a CSV file
	findupic -csv . > duplicates.csv

	# Find duplicate images in the specified directories and write
	# the results to stderr
	findupic -error-log=errors.log /path/to/dir1 /path/to/dir2

# Output

By default, findupic outputs the SHA256 hash and file path of duplicate
images in a human-readable format to the console. For example:

	Duplicate images with hash 34c7d0a9a0083c8613a3cd0f7c704d438ee88b3e3cd141f599ce694f4979ac9:
	/home/user/Pictures/IMG_1234.jpg
	/mnt/media/photos/2019/IMG_5678.jpg

If CSV output is enabled, the tool outputs the results in
comma-separated values format. The output is written to stdout by
default, but can be written to a file using shell redirection. For
example:

	SHA256,Path
	34c7d0a9a0083c8613a3cd0f7c704d438ee88b3e3cd141f599ce694f4979ac9,/home/user/Pictures/IMG_1234.jpg
	34c7d0a9a0083c8613a3cd0f7c704d438ee88b3e3cd141f599ce694f4979ac9,/mnt/media/photos/2019/IMG_5678.jpg

# Exit Status

If the tool encounters any errors while searching for images, it
returns a non-zero exit status to indicate the error. Otherwise, it
returns zero.

# License

This program is released under the Unlicense. You are free to use,
modify, and distribute the program without restriction. The program is
provided as-is, without any warranty or guarantee of its fitness for
any particular purpose. Please see the UNLICENSE.txt file for the full
text of the Unlicense.
*/
package main

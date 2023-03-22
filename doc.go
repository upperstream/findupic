/*
Package findupic implements a command-line tool that searches for
duplicate images within one or more directories. The tool finds images
that have identical SHA256 hashes, indicating that their decoded
bitmaps are identical.

By default, the tool outputs the paths of duplicate images to the
console. The tool also supports writing the output to a CSV file using
the -csv command-line flag. The CSV file consists of two columns:
SHA256 hash and image path.

# Usage

The findupic command accepts one or more directories as arguments. For
example:

	$ findupic ~/Pictures /mnt/media/photos

The tool searches each directory and its subdirectories for image files
with the extensions .jpg, .jpeg, .png, and .gif. For each image file,
the tool computes its SHA256 hash and adds the path to a list of images
with that hash. If the tool finds multiple images with the same hash,
it reports those images as duplicates.

# Output

By default, the tool outputs the paths of duplicate images to the
console. For example:

	Duplicate images with hash 34c7d0a9a0083c8613a3cd0f7c704d438ee88b3e3cd141f599ce694f4979ac9:
	/home/user/Pictures/IMG_1234.jpg
	/mnt/media/photos/2019/IMG_5678.jpg

When the -csv command-line flag is used, the tool writes the output to
a CSV file instead of the console. The CSV file consists of two
columns: SHA256 hash and image path. The header row of the CSV file is
"SHA256", "Path". For example:

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

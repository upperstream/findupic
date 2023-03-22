/*
 * This program is released under the Unlicense. You are free to use,
 * modify, and distribute the program without restriction. The program is
 * provided as-is, without any warranty or guarantee of its fitness for
 * any particular purpose. Please see the UNLICENSE.txt file for the full
 * text of the Unlicense.
 */

package main

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

type Config struct {
	directories []string
	csv         bool
	errorLog    string
}

func main() {
	// Parse command line arguments
	config, err := parseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Open the error log file or use stderr as default
	var errorLog *os.File
	if config.errorLog != "" {
		var err error
		errorLog, err = os.Create(config.errorLog)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating error log file: %s\n", err)
			os.Exit(1)
		}
		defer errorLog.Close()
	}

	// Find duplicate images
	duplicates, err := findDuplicateImages(config.directories, errorLog)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Print results
	if !config.csv {
		printResults(duplicates, os.Stdout)
	} else {
		printResults(duplicates, os.Stdout)
	}
}

// findDuplicateImages recursively walks through the specified directories to find all image files,
// calculates the SHA256 hash for each image file, and returns a map of hash values to slices of
// file paths for duplicate images.
func findDuplicateImages(dirs []string, errorLog *os.File) (map[string][]string, error) {
	// Count the number of erroneous files
	numErrors := 0

	duplicates := make(map[string][]string)
	for _, dir := range dirs {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			if isImageFile(path) {
				hash, err := getImageHash(path)
				if err != nil {
					fmt.Fprintf(errorLog, "%s\n", err.Error())
					numErrors++
					return nil
				}
				if len(duplicates[hash]) > 0 {
					duplicates[hash] = append(duplicates[hash], path)
				} else {
					duplicates[hash] = []string{path}
				}
			}
			return nil
		})
		if err != nil {
			fmt.Fprintf(errorLog, "%s\n", err)
			numErrors++
		}
	}
	if numErrors > 0 {
		fmt.Fprintf(os.Stderr, "Encountered %d error(s). Check the error log for details.\n", numErrors)
	}
	return duplicates, nil
}

// parseArgs parses command-line arguments and returns a slice of directory paths, a boolean flag
// indicating whether CSV output is enabled, and an error. The CSV flag is set to false by default
// and can be enabled by passing the -csv command-line option. If there are no directory arguments,
// the function returns an error indicating that at least one directory argument is required.
func parseArgs() (config Config, err error) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <directory>...\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	flag.BoolVar(&config.csv, "csv", false, "enable CSV output")
	flag.StringVar(&config.errorLog, "error-log", "", "file to log errors (default: stderr)")
	flag.BoolVar(&config.csv, "csv", false, "enable CSV output")
	flag.Parse()

	if len(flag.Args()) == 0 {
		return config, errors.New("at least one directory argument is required")
	}
	config.directories = flag.Args()
	return config, nil
}

// isImageFile returns true if the given file has an image file extension
func isImageFile(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}

// getImageHash returns the SHA256 hash of the given image file
func getImageHash(path string) (string, error) {
	//fmt.Printf("path: %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("error decoding image %s: %s", path, err.Error())
	}

	hasher := sha256.New()
	hasher.Write(convertToRGBA(img).Pix)
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// convertToRGBA converts the given image to an RGBA image
func convertToRGBA(img image.Image) *image.RGBA {
	// If the image is already an RGBA image, return it
	if image, ok := img.(*image.RGBA); ok {
		return image
	}

	// Convert the image to an RGBA image
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	return rgba
}

// printResults writes the image duplicates to a CSV file or standard output,
// depending on the output type specified.
//
// The images parameter is a map containing the SHA256 hash of each image
// and a list of paths to images with the same hash. The out parameter is
// the output file handle for writing the results.
func printResults(images map[string][]string, out *os.File) {
	writer := csv.NewWriter(out)
	defer writer.Flush()

	writer.Write([]string{"SHA256", "Path"})

	for hash, paths := range images {
		if len(paths) > 1 {
			for _, path := range paths {
				writer.Write([]string{hash, path})
			}
		}
	}
}

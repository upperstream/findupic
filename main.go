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

func main() {
	// Define command line options
	var errorLogPtr *string = flag.String("error-log", "", "file to log errors (default: stderr)")
	var csvPtr *bool = flag.Bool("csv", false, "print output in CSV format (default: false)")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: findupic [options] <directory>...")
		flag.PrintDefaults()
		return
	}

	// Open the error log file or use stderr as default
	var errorLog *os.File
	if *errorLogPtr != "" {
		var err error
		errorLog, err = os.Create(*errorLogPtr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating error log file: %s\n", err)
			os.Exit(1)
		}
		defer errorLog.Close()
	} else {
		errorLog = os.Stderr
	}

	// Open the CSV output file or use stdout as default
	var outputWriter *csv.Writer
	if *csvPtr {
		outputWriter = csv.NewWriter(os.Stdout)
		defer outputWriter.Flush()

		outputWriter.Write([]string{"SHA256", "Path"})
	}

	// Count the number of erroneous files
	numErrors := 0

	images := make(map[string][]string)
	for _, dir := range flag.Args() {
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
				images[hash] = append(images[hash], path)
			}
			return nil
		})
		if err != nil {
			fmt.Fprintf(errorLog, "%s\n", err)
			numErrors++
		}
	}

	for hash, paths := range images {
		if len(paths) > 1 {
			if *csvPtr {
				for _, path := range paths {
					outputWriter.Write([]string{hash, path})
				}
			} else {
				fmt.Printf("Duplicate images with hash %s:\n", hash)
				for _, path := range paths {
					fmt.Println(path)
				}
				fmt.Println()
			}
		}
	}

	if numErrors > 0 {
		fmt.Fprintf(os.Stderr, "Encountered %d error(s). Check the error log for details.\n", numErrors)
		os.Exit(1)
	}
}

func isImageFile(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}

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

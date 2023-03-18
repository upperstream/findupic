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
	"encoding/hex"
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
	if len(os.Args) < 2 {
		fmt.Println("Usage: findupic <directory>...")
		return
	}
	images := make(map[string][]string)
	for _, dir := range os.Args[1:] {
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
					return err
				}
				images[hash] = append(images[hash], path)
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	for hash, paths := range images {
		if len(paths) > 1 {
			fmt.Printf("Duplicate images with hash %s:\n", hash)
			for _, path := range paths {
				fmt.Println(path)
			}
			fmt.Println()
		}
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
		return "", err
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

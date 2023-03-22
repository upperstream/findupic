package main

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestIsImageFile(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected bool
	}{
		{"jpg file", "test.jpg", true},
		{"jpeg file", "test.jpeg", true},
		{"png file", "test.png", true},
		{"gif file", "test.gif", true},
		{"non-image file", "test.txt", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isImageFile(tc.path)
			if actual != tc.expected {
				t.Errorf("isImageFile(%s): expected %v, but got %v", tc.path, tc.expected, actual)
			}
		})
	}
}

func TestConvertToRGBA(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})
	img.Set(1, 1, color.RGBA{0, 0, 255, 255})

	actual := convertToRGBA(img)

	// Ensure that the converted image has the same dimensions as the original image
	if actual.Bounds() != img.Bounds() {
		t.Errorf("convertToRGBA: expected bounds %v, but got %v", img.Bounds(), actual.Bounds())
	}

	// Ensure that the converted image has the same colors as the original image
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			if actual.At(x, y) != img.At(x, y) {
				t.Errorf("convertToRGBA: expected color %v, but got %v", img.At(x, y), actual.At(x, y))
			}
		}
	}
}

func TestGetImageHash(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})
	img.Set(1, 1, color.RGBA{0, 0, 255, 255})

	// Write the image to a temporary file
	f, err := os.CreateTemp("", "image-*.png")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	if err := png.Encode(f, img); err != nil {
		t.Fatal(err)
	}

	actual, err := getImageHash(f.Name())
	if err != nil {
		t.Fatalf("getImageHash(%s): %v", f.Name(), err)
	}

	// Expected SHA256 hash of the image
	expected := "ec7c496d26600526d222d6ca19829609dbe79e6f28d0bc979bbb029f9692328f"

	if actual != expected {
		t.Errorf("getImageHash(%s): expected %s, but got %s", f.Name(), expected, actual)
	}
}

func TestPrintResults(t *testing.T) {
	// Define a test case
	tc := struct {
		results map[string][]string
		output  string
	}{
		results: map[string][]string{
			"hash1": {"/path/to/image1.png", "/path/to/image2.png"},
			"hash2": {"/path/to/image3.png", "/path/to/image4.png"},
		},
		output: "output.csv",
	}

	// Clean up the test file after the test completes
	defer os.Remove(tc.output)

	// Open the file for writing
	file, err := os.Create(tc.output)
	if err != nil {
		t.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	printResults(tc.results, file)

	// Verify that the file contents match the expected result
	file, err = os.Open(tc.output)
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var got string
	for scanner.Scan() {
		got += scanner.Text() + "\n"
	}

	expected := "SHA256,Path\n" +
		"hash1,/path/to/image1.png\n" +
		"hash1,/path/to/image2.png\n" +
		"hash2,/path/to/image3.png\n" +
		"hash2,/path/to/image4.png\n"

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}

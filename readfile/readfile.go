package readfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/johneliud/ASCII-Art-Fs/checksum"
)

/*
ReadFile reads the specified banner file, validates its integrity, and returns its content as a slice of strings. It handles different file extensions, performs checksum validation, and splits the content based on the file type.
*/
func ReadFile(bannerFile string) []string {
	if filepath.Ext(bannerFile) != ".txt" {
		fmt.Printf("Incorrect file extension associated with file %v", bannerFile)
		os.Exit(1)
	}

	bannerFileData, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Printf("Error reading file %v: %v", bannerFile, err)
		os.Exit(1)
	}

	fileHash := checksum.CheckFileValidity(bannerFileData)
	switch bannerFile {
	case "standard.txt":
		if fileHash != "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf" {
			fmt.Println("Error: Possible file corruption with \"standard.txt\"")
			os.Exit(1)
		}
	case "shadow.txt":
		if fileHash != "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73" {
			fmt.Println("Error: Possible file corruption with \"shadow.txt\"")
			os.Exit(1)
		}
	case "thinkertoy.txt":
		if fileHash != "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3" {
			fmt.Println("Error: Possible file corruption with \"thinkertoy.txt\"")
			os.Exit(1)
		}
	}

	var splitBannerFileData []string
	if bannerFile == "thinkertoy.txt" {
		splitBannerFileData = strings.Split(string(bannerFileData), "\r\n")
	} else {
		splitBannerFileData = strings.Split(string(bannerFileData), "\n")
	}
	return splitBannerFileData
}

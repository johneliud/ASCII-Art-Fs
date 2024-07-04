package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/johneliud/ASCII-Art-Fs/printart"
	"github.com/johneliud/ASCII-Art-Fs/readfile"
)

func main() {
	var bannerFont string
	var colorFlag string

	// Define and parse command line arguments
	flag.StringVar(&bannerFont, "font", "standard", "Flag to specify type of banner file to be used")
	flag.StringVar(&colorFlag, "color", "", "Flag to specify font color of printed output")
	flag.Parse()

	nonFlagArgs := flag.Args()

	if len(nonFlagArgs) == 0 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	var inputString, substringToColor string

	switch len(os.Args) {
	case 2:
		inputString = os.Args[1]
	case 3:
		if strings.HasPrefix(os.Args[1], "--color=") {
			inputString = nonFlagArgs[0]
		} else if strings.HasPrefix(os.Args[2], "--color=") {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			return
		} else if !strings.Contains(os.Args[1], "--color") && (os.Args[2] != "shadow" &&
			os.Args[2] != "standard" &&
			os.Args[2] != "thinkertoy") {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			return
		} else {
			inputString = os.Args[1]
			bannerFont = os.Args[2]
		}
	case 4:
		if !strings.HasPrefix(os.Args[1], "--color=") {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		} else if os.Args[3] == "shadow" ||
			os.Args[3] == "standard" ||
			os.Args[3] == "thinkertoy" {
			inputString = nonFlagArgs[0]
			bannerFont = nonFlagArgs[1]
		} else {
			substringToColor = nonFlagArgs[0]
			inputString = nonFlagArgs[1]
		}
	case 5:
		if !strings.HasPrefix(os.Args[1], "--color=") {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			return
		} else if os.Args[4] != "shadow" && os.Args[4] != "standard" && os.Args[4] != "thinkertoy" {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		} else {
			substringToColor = nonFlagArgs[0]
			inputString = nonFlagArgs[1]
			bannerFont = nonFlagArgs[2]
		}
	default:
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	switch bannerFont {
	case "standard":
		bannerFont = "standard.txt"
	case "shadow":
		bannerFont = "shadow.txt"
	case "thinkertoy":
		bannerFont = "thinkertoy.txt"
	default:
		bannerFont = "standard.txt"
	}

	bannerFile := readfile.ReadFile(bannerFont)
	printart.PrintArt(bannerFile, inputString, colorFlag, substringToColor)
}

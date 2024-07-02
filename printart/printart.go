package printart

import (
	"fmt"
	"strings"

	"github.com/johneliud/ASCII-Art-Fs/colorlibrary"
)

/*
PrintArt prints ASCII art based on the provided input string, color flag, and substring to color. It reads ASCII art from a bannerFileSlice and handles various scenarios, such as handling unprintable sequences, foreign input, and coloring substrings, printing newline characters, tab characters, and coloring the entire input string if no substring is specified.
*/
func PrintArt(bannerFileSlice []string, inputString, colorFlag, substringToColor string) {
	var (
		ansiCode      string
		ansiResetCode = "\033[0m"
	)

	if colorFlag != "" {
		colorToAccess := strings.ToLower(colorFlag)
		if strings.HasPrefix(colorToAccess, "rgb(") {
			ansiCode = colorlibrary.ConvertRgbToAnsi(colorToAccess)
		} else {
			ansiCode = colorlibrary.ColorsLibrary(colorToAccess)
		}
	} else {
		ansiCode = ""
		ansiResetCode = ""
	}

	if inputString == "\\n" {
		fmt.Println()
		return
	} else if inputString == "" {
		return
	} else if inputString == "\\t" {
		fmt.Println("    ")
		return
	}

	// Handle unprintable sequences
	unprintableSequences := []string{"\\a", "\\b", "\\v", "\\f", "\\r"}

	for _, unprintable := range unprintableSequences {
		if strings.Contains(inputString, unprintable) {
			fmt.Println("Input string contains an unprintable sequence.")
			return
		}
	}

	tabCharText := strings.Replace(inputString, "\\t", "    ", -1)
	newlineCharText := strings.ReplaceAll(tabCharText, "\\n", "\n")
	splitArguments := strings.Split(newlineCharText, "\n")

	// Handle foreign input
	for _, splitArg := range splitArguments {
		for _, char := range splitArg {
			if char < 32 || char > 126 {
				fmt.Println("Input string contains a foreign character absent in the ASCII manual.")
				return
			}
		}
	}

	for _, text := range splitArguments {
		if text == "" {
			fmt.Println()
			continue
		}

		const asciiHeight = 8
		for j := 0; j < asciiHeight; j++ {
			i := 0
			for i < len(text) {
				startingIndex := int(text[i]-32)*9 + 1
				if substringToColor != "" && i+len(substringToColor) <= len(text) && text[i:i+len(substringToColor)] == substringToColor {
					for k := 0; k < len(substringToColor); k++ {
						startingIndex := int(text[i+k]-32)*9 + 1
						fmt.Printf(ansiCode + bannerFileSlice[startingIndex+j] + ansiResetCode)
					}
					i += len(substringToColor)
				} else if substringToColor == "" && colorFlag != "" {
					fmt.Printf(ansiCode + bannerFileSlice[startingIndex+j] + ansiResetCode)
					i++
				} else {
					fmt.Printf(bannerFileSlice[startingIndex+j])
					i++
				}
			}
			fmt.Println()
		}
	}
}

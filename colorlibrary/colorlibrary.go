package colorlibrary

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
CheckValueRange checks if the given integer is within the valid range for color values (0-255). If the value is outside this range, it prints an error message and exits the program.
*/
func CheckValueRange(n int) {
	if n < 0 || n > 255 {
		fmt.Println("Invalid color value. Must be between 0 and 255.")
		os.Exit(1)
	}
}

/*
ConvertRgbToAnsi converts an RGB color value to its ANSI escape sequence. The RGB color value should be provided as a command-line argument in the format "rgb(r,g,b)". The function returns the ANSI escape sequence for the given RGB color.
*/
func ConvertRgbToAnsi(colorToAccess string) string {
	var ansiCode string

	rgbString := strings.ToLower(os.Args[1])
	values := strings.Split(strings.TrimSuffix(strings.Split(rgbString, "(")[1], ")"), ",")
	red := strings.TrimSpace(values[0])
	green := strings.TrimSpace(values[1])
	blue := strings.TrimSpace(values[2])

	r, err1 := strconv.Atoi(red)
	if err1!= nil {
    fmt.Println("Failed converting value to an integer", err1)
    os.Exit(1)
  }

	g, err2 := strconv.Atoi(green)
	if err2!= nil {
    fmt.Println("Failed converting value to an integer", err2)
    os.Exit(1)
  }

	b, err3 := strconv.Atoi(blue)
	if err3!= nil {
    fmt.Println("Failed converting value to an integer", err3)
    os.Exit(1)
  }

  CheckValueRange(r)
	CheckValueRange(g)
	CheckValueRange(b)
	ansiCode = fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	return ansiCode
}

/*
ColorsLibrary is a function that returns the ANSI escape sequence for a given color. It supports a predefined list of colors.
*/
func ColorsLibrary(colorToAccess string) string {
	colorsMap := map[string]string {
		"red":    "\033[31m",
		"green":  "\033[32m",
		"blue":   "\033[34m",
		"orange": "\033[38;5;208m",
		"yellow": "\033[33m",
		"black":  "\033[30m",
		"white":  "\033[37m",
		"pink":   "\033[95m",
		"teal":   "\033[36m",
		"purple": "\033[35m",
		"brown":  "\033[33;2m",
		"beige":  "\033[33;2m",
		"indigo": "\033[94m",
		"violet": "\033[94m",
		"maroon": "\033[31;2m",
		"cream":  "\033[97m",
	}

	ansiCode, ok := colorsMap[colorToAccess]
	if!ok {
    fmt.Printf("Color '%v' not found in the library.\n", colorToAccess)
    os.Exit(1)
  }
	return ansiCode
}

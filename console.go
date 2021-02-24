package console

import (
	"fmt"
	"os"
)

// reference
// https://en.wikipedia.org/wiki/ANSI_escape_code

// Style -> Define type color
type Style uint8

// Reset const
const (
	Reset Style = 0
)

// String constant
const (
	prefix       string = "\u001b["
	rgbFgPrefix  string = "\u001b[38;5;"
	rgbBgPrefix  string = "\u001b[48;5;"
	normalSuffix string = "m"
	brightSuffix string = ";1m"
)

// Mask constant
const (
	brightnessMask Style = 0b1000000
	decorationMask Style = 0b10000000
)

// foreground colors: 30-37 -> 30-37
const (
	FgBlack Style = iota + 30 // 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// foreground bright colors 30-37 -> ???
const (
	FgBrightBlack Style = iota + 30 + brightnessMask
	FgBrightRed
	FgBrightGreen
	FgBrightYellow
	FgBrightBlue
	FgBrightMagenta
	FgBrightCyan
	FgBrightWBright
)

// background colors 40-47 -> 40-47
const (
	BgBlack Style = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// background bright colors 40-47 -> 104-11
const (
	BgBrightBlack Style = iota + 40 + brightnessMask
	BgBrightRed
	BgBrightGreen
	BgBrightYellow
	BgBrightBlue
	BgBrightMagenta
	BgBrightCyan
	BgBrightWBright
)

// decorations 1-9 -> 129-137
const (
	Bold Style = iota + 1 + decorationMask
	Faint
	Italic
	Underline
	SlowBlink
	RapidBlink
	Invert
	Hide
	Strike
)

func getValue(c Style) (s string) {
	// bit masking to get value
	if c&decorationMask != 0 {
		return fmt.Sprint(c - decorationMask)
	} else if c&brightnessMask != 0 {
		return fmt.Sprint(c - brightnessMask)
	}

	return fmt.Sprint(c)
}

func getSuffix(c Style) (s string) {
	// bit masking to get suffix
	if c&brightnessMask == 0 || c&decorationMask == 0 {
		return normalSuffix
	}
	return brightSuffix
}

func createStyleString(c Style) (s string) {
	s = ""

	s += prefix
	s += getValue(c)
	s += getSuffix(c)

	return s
}

func createRGBString(r, g, b uint8) (s string) {
	s = fmt.Sprint(16 + r/51*36 + g/51*6 + b/51)
	return
}

// SetStyle -> set text and background colors
func SetStyle(colors ...Style) {
	s := ""
	for _, c := range colors {
		s += createStyleString(c)
	}

	fmt.Fprint(os.Stdout, s)
}

// SetFgRGB -> set text color via rgb
func SetFgRGB(r, g, b uint8) {
	s := ""
	s += rgbFgPrefix
	s += createRGBString(r, g, b)
	s += normalSuffix

	fmt.Fprint(os.Stdout, s)
}

// SetBgRgb -> set background color via rgb

// ResetStyle -> reset color and decoration to default
func ResetStyle() {
	fmt.Fprint(os.Stdout, createStyleString(Reset))
}

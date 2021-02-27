package colorize

import (
	"fmt"
)

// reference
// https://en.wikipedia.org/wiki/ANSI_escape_code

// Style -> Define type color
type Style uint8

// Reset const
const (
	Reset Style = 0
)

// String constants - prefixes
const (
	prefix            string = "\u001b["
	rgbFgPrefix       string = prefix + "38;5;"
	rgbBgPrefix       string = prefix + "48;5;"
	truecolorFgPrefix string = prefix + "38;2;"
	truecolorBgPrefix string = prefix + "48;2;"
)

// String constants - suffixes
const (
	suffix       string = "m"
	brightSuffix string = ";1m"
)

// String constants - cursor
const (
	moveHome         string = "H"
	moveUp           string = "1A"
	moveDown         string = "1B"
	moveLeft         string = "1D"
	moveRight        string = "1C"
	moveNewLine      string = "1E"
	movePreviousLine string = "1F"
)

// String constants - erase
const (
	clearScreen string = prefix + "2J"
	clearLine   string = prefix + "2K"
)

// Mask constants
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

// foreground bright colors 30-37 -> 94-101
const (
	FgBrightBlack Style = iota + 30 + brightnessMask
	FgBrightRed
	FgBrightGreen
	FgBrightYellow
	FgBrightBlue
	FgBrightMagenta
	FgBrightCyan
	FgBrightWhite
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

// get constant value by masking
func getValue(c Style) (s string) {
	// bit masking to get value
	if c&decorationMask != 0 {
		return fmt.Sprint(c - decorationMask)
	} else if c&brightnessMask != 0 {
		return fmt.Sprint(c - brightnessMask)
	}

	return fmt.Sprint(c)
}

// get suffix to end style string
func getSuffix(c Style) (s string) {
	// bit masking to get suffix
	if c&brightnessMask == 0 || c&decorationMask == 0 {
		return suffix
	}
	return brightSuffix
}

// create string from constant
func createStyleString(c Style) (s string) {
	s = ""
	s += prefix
	s += getValue(c)
	s += getSuffix(c)

	return s
}

// create string from rgb
func createRGBString(r, g, b uint8) (s string) {
	s = fmt.Sprint(16 + r/51*36 + g/51*6 + b/51)
	return
}

// create string from truecolor
func createTruecolorString(r, g, b uint8) (s string) {
	s = fmt.Sprintf("%d;%d;%d", r, g, b)
	return
}

// finally apply the style
func applyStyle(style string) (e error) {
	_, e = fmt.Print(style)
	return e
}

// create string to move cursor to xy
func createCursorXYString(x, y uint8) (s string) {
	s = fmt.Sprintf("%d;%d;H", y, x)
	return
}

// SetStyle -> set text and background colors
func SetStyle(colors ...Style) {
	s := ""
	for _, c := range colors {
		s += createStyleString(c)
	}

	applyStyle(s)
}

// SetFgRGB -> set text color via rgb
func SetFgRGB(r, g, b uint8) {
	s := ""
	s += rgbFgPrefix
	s += createRGBString(r, g, b)
	s += suffix

	applyStyle(s)
}

// SetBgRGB -> set background color via rgb
func SetBgRGB(r, g, b uint8) {
	s := ""
	s += rgbBgPrefix
	s += createRGBString(r, g, b)
	s += suffix

	applyStyle(s)
}

// SetFgTruecolor -> set text color via rgb (true color)
func SetFgTruecolor(r, g, b uint8) {
	s := ""
	s += truecolorFgPrefix
	s += createTruecolorString(r, g, b)
	s += suffix

	applyStyle(s)
}

// SetBgTruecolor -> set background color via rgb (true color)
func SetBgTruecolor(r, g, b uint8) {
	s := ""
	s += truecolorBgPrefix
	s += createTruecolorString(r, g, b)
	s += suffix

	applyStyle(s)
}

// ResetStyle -> reset color and decoration to default
func ResetStyle() {
	s := createStyleString(Reset)
	applyStyle(s)
}

// MoveCursorToXY -> Move cursor to a x,y position
func MoveCursorToXY(x, y uint8) {
	s := ""
	s += prefix
	s += createCursorXYString(x, y)
	applyStyle(s)
}

// Clear -> clears the console (everything) using the current style
func Clear() {
	applyStyle(clearScreen)
}

// ClearLine -> clears the current console line using the current style
func ClearLine() {
	applyStyle(clearLine)
}

// MoveCursorBy -> moves the cursor by x, y relative to current position
func MoveCursorBy(x, y int8) {
	var i, j int8
	var s string
	s = ""

	for i = 0; i < x; i++ {
		s += prefix
		if x > 0 {
			s += moveRight
		} else {
			s += moveLeft
		}
	}

	for j = 0; j < y; j++ {
		s += prefix
		if y > 0 {
			s += moveDown
		} else {
			s += moveUp
		}
	}

	applyStyle(s)
}

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
	if c&brightnessMask != 0 {
		return brightSuffix
	}
	return suffix
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

// List of functions for fast styling
// Normal colors

// Red text
func Red(s string) {
	SetStyle(FgRed)
	fmt.Print(s)
	ResetStyle()
}

// Green text
func Green(s string) {
	SetStyle(FgGreen)
	fmt.Print(s)
	ResetStyle()
}

// Yellow text
func Yellow(s string) {
	SetStyle(FgYellow)
	fmt.Print(s)
	ResetStyle()
}

// Blue text
func Blue(s string) {
	SetStyle(FgBlue)
	fmt.Print(s)
	ResetStyle()
}

// Magenta text
func Magenta(s string) {
	SetStyle(FgMagenta)
	fmt.Print(s)
	ResetStyle()
}

// Cyan text
func Cyan(s string) {
	SetStyle(FgCyan)
	fmt.Print(s)
	ResetStyle()
}

// White text
func White(s string) {
	SetStyle(FgWhite)
	fmt.Print(s)
	ResetStyle()
}

// Bright colors

// BrightRed -> Red bright text
func BrightRed(s string) {
	SetStyle(FgBrightRed)
	fmt.Print(s)
	ResetStyle()
}

// BrightGreen -> Green bright text
func BrightGreen(s string) {
	SetStyle(FgBrightGreen)
	fmt.Print(s)
	ResetStyle()
}

// BrightYellow -> Yellow bright text
func BrightYellow(s string) {
	SetStyle(FgBrightYellow)
	fmt.Print(s)
	ResetStyle()
}

// BrightBlue -> Blue bright text
func BrightBlue(s string) {
	SetStyle(FgBrightBlue)
	fmt.Print(s)
	ResetStyle()
}

// BrightMagenta -> Magenta bright text
func BrightMagenta(s string) {
	SetStyle(FgBrightMagenta)
	fmt.Print(s)
	ResetStyle()
}

// BrightCyan -> Cyan bright text
func BrightCyan(s string) {
	SetStyle(FgBrightCyan)
	fmt.Print(s)
	ResetStyle()
}

// BrightWhite -> White bright text
func BrightWhite(s string) {
	SetStyle(FgBrightWhite)
	fmt.Print(s)
	ResetStyle()
}

// Text styling

// BoldText -> print text in bold
func BoldText(s string) {
	SetStyle(Bold)
	fmt.Print(s)
	ResetStyle()
}

// FaintText -> print text in faint mode
func FaintText(s string) {
	SetStyle(Faint)
	fmt.Print(s)
	ResetStyle()
}

// ItalicText -> print text in italic
func ItalicText(s string) {
	SetStyle(Italic)
	fmt.Print(s)
	ResetStyle()
}

// UnderlineText -> print text with underline
func UnderlineText(s string) {
	SetStyle(Bold)
	fmt.Print(s)
	ResetStyle()
}

// SlowBlinkText -> print text with slow blink effect
func SlowBlinkText(s string) {
	SetStyle(SlowBlink)
	fmt.Print(s)
	ResetStyle()
}

// RapidBlinkText -> print text with rapid blink effect
func RapidBlinkText(s string) {
	SetStyle(RapidBlink)
	fmt.Print(s)
	ResetStyle()
}

// InvertText -> print text with inverted effect
func InvertText(s string) {
	SetStyle(Bold)
	fmt.Print(s)
	ResetStyle()
}

// HideText -> print text with hide effect
func HideText(s string) {
	SetStyle(Hide)
	fmt.Print(s)
	ResetStyle()
}

// StrikeText -> print text with Strike effect
func StrikeText(s string) {
	SetStyle(Strike)
	fmt.Print(s)
	ResetStyle()
}

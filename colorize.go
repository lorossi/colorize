// Package colorize is a simple Go package to have colored and formatted text inside your terminal
package colorize

// reference
// https://en.wikipedia.org/wiki/ANSI_escape_code

// Style -> Define type used in all the styling options
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
	BgBrightWhite
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
	Framed    Style = 51 + decorationMask // Not widely supported
	Encircled Style = 52 + decorationMask // Not widely supported
)

// SetStyle -> Set text and background colors.
func SetStyle(colors ...Style) {
	s := ""
	for _, c := range colors {
		s += createStyleString(c)
	}

	applyStyle(s)
}

// ResetStyle -> Reset color, background and decoration to default.
func ResetStyle() {
	s := createStyleString(Reset)
	applyStyle(s)
}

// SetFgRGB -> Set text color via RGB. RGB in range 0-255, for a total output of 256 colors.
func SetFgRGB(r, g, b uint8) {
	s := ""
	s += rgbFgPrefix
	s += createRGBString(r, g, b)
	s += suffix

	applyStyle(s)
}

// SetBgRGB -> Set background color via RGB. RGB in range 0-255, for a total output of 256 colors.
func SetBgRGB(r, g, b uint8) {
	s := ""
	s += rgbBgPrefix
	s += createRGBString(r, g, b)
	s += suffix

	applyStyle(s)
}

// SetFgTruecolor -> Set text color via RGB (true color). RGB in range 0-255, for a total output of 16777216 colors.
func SetFgTruecolor(r, g, b uint8) {
	s := ""
	s += truecolorFgPrefix
	s += createTruecolorString(r, g, b)
	s += suffix

	applyStyle(s)
}

// SetBgTruecolor -> Set background color via RGB (true color). RGB in range 0-255, for a total output of 16777216 colors.
func SetBgTruecolor(r, g, b uint8) {
	s := ""
	s += truecolorBgPrefix
	s += createTruecolorString(r, g, b)
	s += suffix

	applyStyle(s)
}

// SetFgTruecolorHSL -> Set text color via HSL (true color). HSL in range 0-255, for a total output of 16777216 colors.
func SetFgTruecolorHSL(h, s, l uint8) {
	r, g, b := hslTOrgb(h, s, l)
	style := ""
	style += truecolorFgPrefix
	style += createTruecolorString(r, g, b)
	style += suffix

	applyStyle(style)
}

// SetBgTruecolorHSL -> Set background color via HSL (true color). HSL in range 0-255, for a total output of 16777216 colors.
func SetBgTruecolorHSL(h, s, l uint8) {
	r, g, b := hslTOrgb(h, s, l)
	style := ""
	style += truecolorBgPrefix
	style += createTruecolorString(r, g, b)
	style += suffix

	applyStyle(style)
}

// MoveCursorToXY -> Move cursor to a x,y  (zero-indexed) position inside the terminal.
func MoveCursorToXY(x, y uint8) {
	s := ""
	s += prefix
	s += createCursorXYString(x-1, y-1)
	applyStyle(s)
}

// Clear -> Clear the console (deleting verything) using the current style.
func Clear() {
	applyStyle(clearScreen)
}

// ClearLine -> Clear the current console line using the current style.
func ClearLine() {
	applyStyle(clearLine)
}

// MoveCursorBy -> Move the cursor by x, y relative to current position.
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

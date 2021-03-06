// Package colorize is a simple Go package to have colored and formatted text inside your terminal
package colorize

import "fmt"

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

// StyleText -> Returns string formatted according to styles
func StyleText(text string, colors ...interface{}) (formatted string) {
	formatted = ""
	for _, c := range colors {
		formatted += createStyleString(c.(Style))
	}
	formatted += text
	formatted += createStyleString(Reset)
	return
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

// All these functions allow quick text styling
// Normal text colors, bright text colors, normal background colors, bright background clors, styled text

// Normal colors

// Red -> Red text
func Red(text ...interface{}) (formatted string) {
	return styledText(FgRed, text)
}

// Green -> Green text
func Green(text ...interface{}) (formatted string) {
	return styledText(FgGreen, text)
}

// Yellow -> Yellow text
func Yellow(text ...interface{}) (formatted string) {
	return styledText(FgYellow, text)
}

// Blue -> Blue text
func Blue(text ...interface{}) (formatted string) {
	return styledText(FgBlue, text)
}

// Magenta -> Magenta text
func Magenta(text ...interface{}) (formatted string) {
	return styledText(FgMagenta, text)
}

// Cyan -> Cyan text
func Cyan(text ...interface{}) (formatted string) {
	return styledText(FgCyan, text)
}

// White -> White text
func White(text ...interface{}) (formatted string) {
	return styledText(FgWhite, text)
}

// Bright colors

// BrightRed -> Red bright text
func BrightRed(text ...interface{}) (formatted string) {
	return styledText(FgBrightRed, text)
}

// BrightGreen -> Green bright text
func BrightGreen(text ...interface{}) (formatted string) {
	return styledText(FgBrightGreen, text)
}

// BrightYellow -> Yellow bright text
func BrightYellow(text ...interface{}) (formatted string) {
	return styledText(FgBrightYellow, text)
}

// BrightBlue -> Blue bright text
func BrightBlue(text ...interface{}) (formatted string) {
	return styledText(FgBrightBlue, text)
}

// BrightMagenta -> Magenta bright text
func BrightMagenta(text ...interface{}) (formatted string) {
	return styledText(FgBrightMagenta, text)
}

// BrightCyan -> Cyan bright text
func BrightCyan(text ...interface{}) (formatted string) {
	return styledText(FgBrightCyan, text)
}

// BrightWhite -> White bright text
func BrightWhite(text ...interface{}) (formatted string) {
	return styledText(FgBrightWhite, text)
}

// Background colors

// RedBg -> Red background
func RedBg(text ...interface{}) (formatted string) {
	return styledText(BgRed, text)
}

// GreenBg -> Green background
func GreenBg(text ...interface{}) (formatted string) {
	return styledText(BgGreen, text)
}

// YellowBg -> Yellow background
func YellowBg(text ...interface{}) (formatted string) {
	return styledText(BgYellow, text)
}

// BlueBg -> Blue background
func BlueBg(text ...interface{}) (formatted string) {
	return styledText(BgBlue, text)
}

// MagentaBg -> Magenta background
func MagentaBg(text ...interface{}) (formatted string) {
	return styledText(BgMagenta, text)
}

// CyanBg -> Cyan background
func CyanBg(text ...interface{}) (formatted string) {
	return styledText(BgCyan, text)
}

// WhiteBg -> White background
func WhiteBg(text ...interface{}) (formatted string) {
	return styledText(BgWhite, text)
}

// Bright colors

// BrightRedBg -> Red bright background
func BrightRedBg(text ...interface{}) (formatted string) {
	return styledText(BgBrightRed, text)
}

// BrightGreenBg -> Green bright background
func BrightGreenBg(text ...interface{}) (formatted string) {
	return styledText(BgBrightGreen, text)
}

// BrightYellowBg -> Yellow bright background
func BrightYellowBg(text ...interface{}) (formatted string) {
	return styledText(BgBrightYellow, text)
}

// BrightBlueBg -> Blue bright background
func BrightBlueBg(text ...interface{}) (formatted string) {
	return styledText(BgBrightBlue, text)
}

// BrightMagentaBg -> Magenta bright background
func BrightMagentaBg(text ...interface{}) (formatted string) {
	return styledText(BgBrightMagenta, text)
}

// BrightCyanBg -> Cyan bright background
func BrightCyanBg(text ...interface{}) (formatted string) {
	return styledText(BgBrightCyan, text)
}

// BrightWhiteBg -> White bright background
func BrightWhiteBg(text ...interface{}) (formatted string) {
	return styledText(BgBrightWhite, text)
}

// Text styling

// BoldStyle -> Bold text
func BoldStyle(text ...interface{}) (formatted string) {
	return styledText(Bold, text)
}

// FaintStyle -> Faint text
func FaintStyle(text ...interface{}) (formatted string) {
	return styledText(Faint, text)
}

// ItalicStyle -> Italic text
func ItalicStyle(text ...interface{}) (formatted string) {
	return styledText(Italic, text)
}

// UnderlineStyle -> Underlined text
func UnderlineStyle(text ...interface{}) (formatted string) {
	return styledText(Underline, text)
}

// SlowBlinkStyle -> Slow blinking text
func SlowBlinkStyle(text ...interface{}) (formatted string) {
	return styledText(SlowBlink, text)
}

// RapidBlinkStyle -> Rapid blink text
func RapidBlinkStyle(text ...interface{}) (formatted string) {
	return styledText(RapidBlink, text)
}

// InvertStyle -> Inverted text
func InvertStyle(text ...interface{}) (formatted string) {
	return styledText(Invert, text)
}

// HideStyle -> Hidden text
func HideStyle(text ...interface{}) (formatted string) {
	return styledText(Hide, text)
}

// StrikeStyle -> Striked text
func StrikeStyle(text ...interface{}) (formatted string) {
	return styledText(Strike, text)
}

// FramedStyle -> Framed text - NOT WIDELY SUPPORTED
func FramedStyle(text ...interface{}) (formatted string) {
	return styledText(Framed, text)
}

// EncircledStyle -> Encircled text - NOT WIDELY SUPPORTED
func EncircledStyle(text ...interface{}) (formatted string) {
	return styledText(Encircled, text)
}

// collection of functions used in colorize

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

// create styled text. This function both takes the style and the string.
func styledText(color Style, text ...interface{}) (formatted string) {
	// create colored string
	formatted = createStyleString(color)
	// iterate throught each interface item
	for _, t := range text {
		// try to convert interface to string
		s := fmt.Sprint(t)
		// if so, remove starting and leading square brackets
		if len(s) > 1 {
			formatted += s[1 : len(s)-1]
		}
	}
	// add reset character
	formatted += createStyleString(Reset)
	return
}

// converts HSL (range 0-255) to RGB (range 0-255)
func hslTOrgb(h, s, l uint8) (r, g, b uint8) {
	var R, G, B float64

	H := float64(h) / 255
	S := float64(s) / 255
	L := float64(l) / 255

	if S == 0 {
		R = L
		G = L
		B = L
	} else {
		var c1, c2 float64
		if L < 0.5 {
			c2 = L * (1 + S)
		} else {
			c2 = (L + S) - (L * S)
		}
		c1 = 2*L - c2

		hueToRgb := func(v1, v2, v3 float64) (v float64) {
			if v3 < 0 {
				v3++
			} else if v3 > 1 {
				v3--
			}

			if 6.0*v3 < 1 {
				return v1 + (v2-v1)*6.0*v3
			}
			if 2.0*v3 < 1 {
				return v2
			}
			if 3.0*v3 < 1 {
				return v1 + (v2-v1)*(2.0/3.0-v3)*6.0
			}
			return v1
		}

		R = hueToRgb(c1, c2, H+(1.0/3.0))
		G = hueToRgb(c1, c2, H)
		B = hueToRgb(c1, c2, H-(1.0/3.0))
	}

	r = uint8(R * 255)
	g = uint8(G * 255)
	b = uint8(B * 255)

	return r, g, b
}

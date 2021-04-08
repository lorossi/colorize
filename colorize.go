/*
	Package colorize is a simple Go package to have colored and formatted text inside your terminal
	By Lorenzo Rossi - www.lorenzoros.si
*/

package colorize

import (
	"errors"
)

// Style -> Define type used in all the styling options
type Style uint8

// SetStyle -> Set text and background colors.
func SetStyle(colors ...Style) {
	style := ""
	for _, c := range colors {
		style += createStyleString(c)
	}

	applyStyle(style)
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
	style := createStyleString(Reset)
	applyStyle(style)
}

// SetFgRGB -> Set text color via RGB. RGB in range 0-255, for a total output of 256 colors.
func SetFgRGB(r, g, b uint8) {
	style := ""
	style += rgbFgPrefix
	style += createRGBString(r, g, b)
	style += suffix

	applyStyle(style)
}

// SetBgRGB -> Set background color via RGB. RGB in range 0-255, for a total output of 256 colors.
func SetBgRGB(r, g, b uint8) {
	style := ""
	style += rgbBgPrefix
	style += createRGBString(r, g, b)
	style += suffix

	applyStyle(style)
}

// SetFgTruecolor -> Set text color via RGB (true color). RGB in range 0-255, for a total output of 16777216 colors.
func SetFgTruecolor(r, g, b uint8) {
	style := ""
	style += truecolorFgPrefix
	style += createTruecolorString(r, g, b)
	style += suffix

	applyStyle(style)
}

// SetBgTruecolor -> Set background color via RGB (true color). RGB in range 0-255, for a total output of 16777216 colors.
func SetBgTruecolor(r, g, b uint8) {
	style := ""
	style += truecolorBgPrefix
	style += createTruecolorString(r, g, b)
	style += suffix

	applyStyle(style)
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

// SetFgTruecolorHex -> Set text color via hex (true color). Hex in format #FFFFFF, #FFF, FFFFFF, FFF (case insensitive, with or without the leading #).
// It defaults to white in case an invalid sequence is passed
func SetFgTruecolorHex(hex string) {
	// string must be 3, 4, 6, 7 characters long
	// if it's 4 or 7, strip the leading #
	length := len(hex)

	switch length {
	case 4:
		length = 3
		hex = hex[1:4]
	case 7:
		length = 6
		hex = hex[1:7]
	case 3:
	case 6:
		break
	default:
		// default to white
		panic(errors.New("Invalid HEX color"))
	}

	var r, g, b uint8
	switch length {
	case 3:
		r = hexToUint8(hex[0:1])
		g = hexToUint8(hex[1:2])
		b = hexToUint8(hex[2:3])
	case 6:
		r = hexToUint8(hex[0:2])
		g = hexToUint8(hex[2:4])
		b = hexToUint8(hex[4:6])
	}

	style := ""
	style += truecolorFgPrefix
	style += createTruecolorString(r, g, b)
	style += suffix

	applyStyle(style)
}

// SetBgTruecolorHex -> Set background color via hex (true color). Hex in format #FFFFFF, #FFF, FFFFFF, FFF (case insensitive, with or without the leading #).
// It defaults to white in case an invalid sequence is passed
func SetBgTruecolorHex(hex string) {
	// string must be 3, 4, 6, 7 characters long
	// if it's 4 or 7, strip the leading #
	length := len(hex)

	switch length {
	case 4:
		length = 3
		hex = hex[1:4]
	case 7:
		length = 6
		hex = hex[1:7]
	case 3:
	case 6:
		break
	default:
		panic(errors.New("Invalid HEX color"))
	}

	var r, g, b uint8
	switch length {
	case 3:
		r = hexToUint8(hex[0:1])
		g = hexToUint8(hex[1:2])
		b = hexToUint8(hex[2:3])
	case 6:
		r = hexToUint8(hex[0:2])
		g = hexToUint8(hex[2:4])
		b = hexToUint8(hex[4:6])
	}

	style := ""
	style += truecolorBgPrefix
	style += createTruecolorString(r, g, b)
	style += suffix

	applyStyle(style)
}

// Clear -> Clear the console (deleting verything) using the current style.
func Clear() {
	applyStyle(clearScreen)
}

// ClearLine -> Clear the current console line using the current style.
func ClearLine() {
	applyStyle(clearLine)
}

// ClearLineUntilStart -> Clear the current console line from the current position until line start
func ClearLineUntilStart() {
	applyStyle(clearLineUntilStart)
}

// ClearLineUntilEnd -> Clear the current console line from the current position until line end
func ClearLineUntilEnd() {
	applyStyle(clearLineUntilEnd)
}

// ClearXY -> Clear the character in position x, y using the current style
func ClearXY(x, y uint8) {
	style := ""
	style += prefix
	style += createCursorXYString(x, y)
	style += "\b "
	applyStyle(style)
}

// MoveCursorBy -> Move the cursor by x, y relative to current position.
func MoveCursorBy(x, y int8) {
	var i, j int8
	var style string
	style = ""

	for i = 0; i < x; i++ {
		style += prefix
		if x > 0 {
			style += moveRight
		} else {
			style += moveLeft
		}
	}

	for j = 0; j < y; j++ {
		style += prefix
		if y > 0 {
			style += moveDown
		} else {
			style += moveUp
		}
	}

	applyStyle(style)
}

// MoveCursorToXY -> Move cursor to a x,y  (zero-indexed) position inside the terminal.
func MoveCursorToXY(x, y uint8) {
	style := ""
	style += prefix
	style += createCursorXYString(x, y)
	applyStyle(style)
}

// MoveCursorLine -> Move cursor by an amount of lines (positive is down, negative is up)
func MoveCursorLine(lines int8) {
	var dir int8

	if lines < 0 {
		dir = -1
	} else {
		dir = 1
	}

	style := ""
	for i := int8(0); i < lines*dir; i++ {
		style += prefix
		if dir == -1 {
			style += moveUp
		} else if dir == 1 {
			style += moveDown
		}
	}

	applyStyle(style)
}

// SaveCursor -> Save cursor position
func SaveCursor() {
	style := ""
	style += prefix
	style += savePos
	applyStyle(style)
}

// RestoreCursor -> Restore previously saved cursor position
func RestoreCursor() {
	style := ""
	style += prefix
	style += restorePos
	applyStyle(style)
}

// HideCursor -> Hide cursor from terminal (use ShowCursor to re-enable)
func HideCursor() {
	style := ""
	style += prefix
	style += hideCursor
	applyStyle(style)
}

// ShowCursor -> Show cursor in terminal
func ShowCursor() {
	style := ""
	style += prefix
	style += showCursor
	applyStyle(style)
}

package colorize

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// collection of functions used in colorize

// get constant value by masking
func getValue(c Style) (style string) {
	// bit masking to get value
	if c&decorationMask != 0 {
		style = fmt.Sprint(c - decorationMask)
	} else if c&brightnessMask != 0 {
		style = fmt.Sprint(c - brightnessMask)
	} else {
		style = fmt.Sprint(c)
	}
	return
}

// get suffix to end style string
func getSuffix(c Style) (style string) {
	if c&brightnessMask != 0 {
		style = brightSuffix
	} else {
		style = suffix
	}
	return
}

// create string from constant
func createStyleString(c Style) (style string) {
	style = ""
	style += prefix
	style += getValue(c)
	style += getSuffix(c)
	return
}

// create string from rgb
func createRGBString(r, g, b uint8) (style string) {
	style = fmt.Sprint(16 + r/51*36 + g/51*6 + b/51)
	return
}

// create string from truecolor
func createTruecolorString(r, g, b uint8) (style string) {
	style = fmt.Sprintf("%d;%d;%d", r, g, b)
	return
}

// finally apply the style
func applyStyle(style string) (e error) {
	_, e = fmt.Print(style)
	return
}

// create string to move cursor to xy
func createCursorXYString(x, y uint8) (style string) {
	style = fmt.Sprintf("%d;%d;H", y, x)
	return
}

// create styled text. This function both takes the style and the string.
func styledText(color Style, text ...interface{}) (formatted string) {
	// create colored string
	formatted = createStyleString(color)
	// iterate throught each interface item
	for _, t := range text {
		// try to convert interface to string
		style := fmt.Sprint(t)
		// if so, remove starting and leading square brackets
		if len(style) > 1 {
			formatted += style[1 : len(style)-1]
		}
	}
	// add reset character
	formatted += createStyleString(Reset)
	return
}

// converts the hue to rgb
func hueToRgb(v1, v2, h float64) (v float64) {
	if h < 0 {
		h++
	} else if h > 1 {
		h--
	}

	if 6*h < 1 {
		v = v1 + (v2-v1)*6*h
	} else if 2*h < 1 {
		v = v2
	} else if 3*h < 2 {
		v = v1 + (v2-v1)*(2/3-h)*6
	} else {
		v = v1
	}

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

		R = hueToRgb(c1, c2, H+(1/3))
		G = hueToRgb(c1, c2, H)
		B = hueToRgb(c1, c2, H-(1/3))
	}

	r = uint8(R * 255)
	g = uint8(G * 255)
	b = uint8(B * 255)

	return
}

// converts hex string to uint8
func hexToUint8(hex string) (num uint8) {
	hex = strings.ToLower(hex)
	bigNum, e := strconv.ParseUint(hex, 16, 8)

	if e == nil {
		num = uint8(bigNum)
	} else {
		panic(errors.New("Invalid hex value"))
	}

	return
}

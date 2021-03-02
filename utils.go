package colorize

import "fmt"

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

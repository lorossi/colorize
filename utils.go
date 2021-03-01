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

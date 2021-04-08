package colorize

// constants reference
// https://en.wikipedia.org/wiki/ANSI_escape_code

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
	savePos          string = "s"
	restorePos       string = "u"
	hideCursor       string = "?25l"
	showCursor       string = "?25h"
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
	clearScreen         string = prefix + "2J"
	clearLine           string = prefix + "2K"
	clearLineUntilEnd   string = prefix + "0K"
	clearLineUntilStart string = prefix + "1K"
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

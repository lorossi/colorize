package colorize

// All these functions allow quick text styling
// Normal text colors, bright text colors, normal background colors, bright background clors, styled text

// Normal colors

// Red text
func Red(text ...interface{}) (formatted string) {
	return styledText(FgRed, text)
}

// Green text
func Green(text ...interface{}) (formatted string) {
	return styledText(FgGreen, text)
}

// Yellow text
func Yellow(text ...interface{}) (formatted string) {
	return styledText(FgYellow, text)
}

// Blue text
func Blue(text ...interface{}) (formatted string) {
	return styledText(FgBlue, text)
}

// Magenta text
func Magenta(text ...interface{}) (formatted string) {
	return styledText(FgMagenta, text)
}

// Cyan text
func Cyan(text ...interface{}) (formatted string) {
	return styledText(FgCyan, text)
}

// White text
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

// RedBg -> red background
func RedBg(text ...interface{}) (formatted string) {
	return styledText(BgRed, text)
}

// GreenBg -> green background
func GreenBg(text ...interface{}) (formatted string) {
	return styledText(BgGreen, text)
}

// YellowBg -> yellow background
func YellowBg(text ...interface{}) (formatted string) {
	return styledText(BgYellow, text)
}

// BlueBg -> blue background
func BlueBg(text ...interface{}) (formatted string) {
	return styledText(BgBlue, text)
}

// MagentaBg -> magenta background
func MagentaBg(text ...interface{}) (formatted string) {
	return styledText(BgMagenta, text)
}

// CyanBg -> cyan background
func CyanBg(text ...interface{}) (formatted string) {
	return styledText(BgCyan, text)
}

// WhiteBg -> white background
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

// BoldStyle -> print text in bold
func BoldStyle(text ...interface{}) (formatted string) {
	return styledText(Bold, text)
}

// FaintStyle -> print text in faint mode
func FaintStyle(text ...interface{}) (formatted string) {
	return styledText(Faint, text)
}

// ItalicStyle -> print text in italic
func ItalicStyle(text ...interface{}) (formatted string) {
	return styledText(Italic, text)
}

// UnderlineStyle -> print text with underline
func UnderlineStyle(text ...interface{}) (formatted string) {
	return styledText(Underline, text)
}

// SlowBlinkStyle -> print text with slow blink effect
func SlowBlinkStyle(text ...interface{}) (formatted string) {
	return styledText(SlowBlink, text)
}

// RapidBlinkStyle -> print text with rapid blink effect
func RapidBlinkStyle(text ...interface{}) (formatted string) {
	return styledText(RapidBlink, text)
}

// InvertStyle -> print text with inverted effect
func InvertStyle(text ...interface{}) (formatted string) {
	return styledText(Invert, text)
}

// HideStyle -> print text with hide effect
func HideStyle(text ...interface{}) (formatted string) {
	return styledText(Hide, text)
}

// StrikeStyle -> print text with strike effect
func StrikeStyle(text ...interface{}) (formatted string) {
	return styledText(Strike, text)
}

// FramedStyle -> print text with framed effect - NOT WIDELY SUPPORTED
func FramedStyle(text ...interface{}) (formatted string) {
	return styledText(Framed, text)
}

// EncircledStyle -> print text with encircled effect - NOT WIDELY SUPPORTED
func EncircledStyle(text ...interface{}) (formatted string) {
	return styledText(Encircled, text)
}

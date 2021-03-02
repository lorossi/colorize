package colorize

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

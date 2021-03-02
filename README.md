# Colorize

<p align="center">
  <img src="/examples/readmedemo.png">  
</p>

<p align="center">
   <img src="https://img.shields.io/github/repo-size/lorossi/colorize?style=flat-square">
   <img src="https://img.shields.io/maintenance/yes/2021?style=flat-square">
   <img src="https://img.shields.io/github/last-commit/lorossi/colorize/main?style=flat-square">
   <img src="https://img.shields.io/github/v/release/lorossi/colorize?style=flat-square">
</p>

<p align="center">
  <span style="font-size:larger;">Colorize is a simple and handy Go package that lets you use colors and styling in your console!</span>
</p>

## Examples

Set the style manually:

```go
colorize.SetStyle(colorize.Bold, colorize.FgBrightBlue, colorize.BgBrightYellow)
Fmt.Println("WOW")
colorize.ResetStyle()
```

Set an rgb value for the text:

```go
colorize.SetFgRGB(255, 0, 0)
Fmt.Println("Colored!")
colorize.ResetStyle()
```

Set an rgb value for the background:

```go
colorize.SetBgRGB(0, 255, 0)
Fmt.Println("Now is green!")
colorize.ResetStyle()
```

Set truecolors text and background colors! This gives the users more colors than its rgb counterpart but it's less supported (Win10 powershell and linux terminal support this):

```go
colorize.SetFgTruecolor(255, 255, 0)
colorize.SetBgTruecolor(0, 0, 255)
Fmt.Println("Everything is so colorful!")
colorize.ResetStyle()
```

Set color by HSL values:

```go
colorize.SetFgTruecolorHSL(92, 255, 127)
colorize.SetBgTruecolorHSL(112, 255, 127)
fmt.Println("RED on GREEN!")
colorize.ResetStyle()
```

**Never foget to reset the style via the `ResetStyle()` function!**

Set a text color, background color or style with the quick functions:

```go
fmt.Println(colorize.Green("Green text!", 123, "also green numbers!"))
fmt.Println(colorize.BrightMagentaBg("So magenta and so bright!"))
fmt.Println(colorize.Bold("This is so bold!"))
```

**See a few more examples [here.](/examples/main.go)**

## List of constants

### Text color constants

```go
FgBlack
FgRed
FgGreen
FgYellow
FgBlue
FgMagenta
FgCyan
FgWhite
FgBrightBlack
FgBrightRed
FgBrightGreen
FgBrightYellow
FgBrightBlue
FgBrightMagenta
FgBrightCyan
FgBrightWhite
```

### Backgroud color constants

```go
BgBlack
BgRed
BgGreen
BgYellow
BgBlue
BgMagenta
BgCyan
BgWhite
BgBrightBlack
BgBrightRed
BgBrightGreen
BgBrightYellow
BgBrightBlue
BgBrightMagenta
BgBrightCyan
BgBrightWhite
```

### Text decoration constants

```go
Bold
Faint
Italic
Underline
SlowBlink
RapidBlink
Invert
Hide
Strike
Framed  // Not widely supported
Encircled // Not widely supported
```

## Quick functions

### Text color functions

```go
Red()
Green()
Yellow()
Blue()
Magenta()
Cyan()
White()
BrightRed()
BrightGreen()
BrightYellow()
BrightBlue()
BrightMagenta()
BrightCyan()
BrightWhite()
```

### Background color functions

```go
RedBg()
GreenBg()
YellowBg()
BlueBg()
MagentaBg()
CyanBg()
WhiteBg()
BrightRedBg()
BrightGreenBg()
BrightYellowBg()
BrightBlueBg()
BrightMagentaBg()
BrightCyanBg()
BrightWhiteBg()
```

### Text decoration functions

```go
BoldStyle()
FaintStyle()
ItalicStyle()
UnderlineStyle()
SlowBlinkStyle()
RapidBlinkStyle()
InvertStyle()
HideStyle()
StrikeStyle()
FramedStyle() // Not widely supported
EncircledStyle()  // Not widely supported
```

## Docs

[Read the documentation here](/DOCS.md).

This package uses no dependancies exluding the built-in functions.

There are a lot of packages similar (and even better!) than this one. But since I wanted to use colored terminal for my project [Journal](https://www.github.com/lorossi/journal), I decided to make my own package, as a learning experience. I don't regret it at all. It's been fun so far.

### File structure

There are currently **3** go files:

1. `colorize.go` is the main file, containing the public (exported) functions.
2. `quickfunctions.go` is the file containing the so called *(by me, at least)* "quick functions" to rapidly create colored or styled text.
3. `utils.go` is the file containing all the private (unexported) functions.

## License

This project is distributed under CC 4.0 License.

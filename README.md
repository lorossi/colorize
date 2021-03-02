# Colorize

Simple Go package to have colored and formatted text inside your terminal

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
fmt.Println(colorize.Green("Green text!"))
fmt.Println(colorize.BrightMagentaBg("So magenta and so bright!"))
fmt.Println(colorize.Bold("This is so bold!"))
```

## List of constants

### Text color

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

### Backgroud color

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

### Text decorations

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

## Docs

[Read the documentation here](/DOCS.md)

[Generated using Gomarkdoc](https://github.com/princjef/gomarkdoc)

## License

This project is distributed under CC 4.0 License.

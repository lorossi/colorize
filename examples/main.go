package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/lorossi/colorize"
)

func main() {
	fmt.Println("Now printing starry night")
	starryNight()
	/*fmt.Println("Now printing rainbow stripes demo")
	rainbowStripes()
	fmt.Println("Now printing advertising")
	advertising()
	fmt.Println("Now printing Christmas tree demo")
	christmasTree()
	fmt.Println("Now printing README demo")
	readmeDemo()
	fmt.Println("Now printing rainbow points demo")
	rainbowPoints()
	*/
}

func starryNight() {
	const width int = 100
	const height int = 25
	const stars int = 50

	// initialize the random number generator
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// save cursor location to restore it later
	colorize.SaveCursor()

	// set background and text color via Hex
	colorize.SetBgTruecolorHex("#000072")
	colorize.SetFgTruecolorHex("#8b8b00")
	// set cursor to top left corner
	colorize.MoveCursorToXY(0, 0)
	// clear screen (will also set everything to background color)
	colorize.Clear()
	// draw stars
	for i := 0; i < stars; i++ {
		x := uint8(r1.Intn(width))
		y := uint8(r1.Intn(height))
		colorize.MoveCursorToXY(x, y)
		fmt.Print("★")
	}
	colorize.ResetStyle()

	// reset cursor position and skip lines
	colorize.RestoreCursor()
	colorize.MoveCursorLine(int8(height))
}

func rainbowStripes() {
	var stripe string
	var stripesColors [7][3]uint8

	// initialize stripe width
	stripe = ""
	for i := 0; i < 32; i++ {
		stripe += " "
	}
	// initialize stripe color
	stripesColors = [7][3]uint8{{255, 0, 0}, {255, 127, 0}, {255, 255, 0}, {0, 255, 0}, {0, 0, 255}, {75, 0, 130}, {148, 0, 211}}

	fmt.Println()
	for _, color := range stripesColors {
		colorize.SetBgTruecolor(color[0], color[1], color[2])
		fmt.Println(stripe)
		colorize.ResetStyle()
	}
	fmt.Println()
}

func advertising() {
	fmt.Println()
	// set the colors in the verbose way
	colorize.SetStyle(colorize.FgBrightRed, colorize.BgWhite, colorize.Bold, colorize.Underline, colorize.RapidBlink)
	fmt.Println("NEW PACKAGE!")
	colorize.ResetStyle()
	// or use the faster way
	fmt.Println(colorize.StyleText("Offering:", colorize.FgBlue, colorize.BgWhite, colorize.Bold))
	// no need to reset!

	fmt.Println(colorize.BoldStyle("Bold text!"), colorize.Red("Colored text!"), colorize.UnderlineStyle("Underlined style!"), colorize.Yellow("Crazy stuff!"))
	fmt.Println()

	colorize.SetStyle(colorize.FgBrightGreen, colorize.RapidBlink)
	fmt.Println("BUY NOW! SUPER SALE!")
	colorize.ResetStyle()

	colorize.SetStyle(colorize.FgRed, colorize.Strike)
	fmt.Print("100$")
	colorize.ResetStyle()
	fmt.Println(colorize.BrightWhite(" NOW FREE! GO GET IT!"))
	fmt.Print(colorize.BrightRed("VISIT: "))
	colorize.SetStyle(colorize.FgBrightBlue, colorize.Underline)
	fmt.Println("github.com/lorossi/colorize")
	colorize.ResetStyle()
	colorize.SetStyle(colorize.FgBrightGreen, colorize.RapidBlink)
	fmt.Println("GO FAST! DON'T LOSE YOUR CHANCHE!")
	colorize.ResetStyle()

	fmt.Println()
}

func christmasTree() {
	// initialize the random number generator
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println()
	// initialize tree width
	const width int = 11
	// initialize light colors
	colors := [4]colorize.Style{colorize.FgBrightRed, colorize.FgBrightBlue, colorize.FgBrightGreen, colorize.FgBrightYellow}
	// print star on top
	colorize.MoveCursorBy(int8(width), 0)
	fmt.Println(colorize.Yellow("★"))
	for i := 0; i < width; i++ {
		// leave space on the left
		colorize.MoveCursorBy(int8(width-i), 0)
		// prepare the character
		colorize.SetStyle(colorize.RapidBlink, colorize.BgBrightGreen)
		for k := 0; k < 2*i+1; k++ {
			if r1.Float64() > 0.66 {
				// select current color
				lightColor := colors[rand.Intn(4)]
				colorize.SetStyle(lightColor)
				// blinking light!
				fmt.Print("*")
			} else {
				// tree leae
				fmt.Print(" ")
			}
		}
		colorize.ResetStyle()
		fmt.Println()
	}

	// draw trunk
	for i := 0; i < 2; i++ {
		colorize.MoveCursorBy(int8(width), 0)
		colorize.SetBgTruecolorHSL(11, 255, 60)
		fmt.Println(" ")
		colorize.ResetStyle()
	}
	fmt.Println()
}

func readmeDemo() {
	fmt.Println()
	colorize.MoveCursorBy(20, 0)
	colorize.SetStyle(colorize.FgRed, colorize.Bold)
	fmt.Println("Colorize")
	colorize.ResetStyle()
	colorize.MoveCursorBy(7, 0)
	colorize.SetStyle(colorize.FgGreen, colorize.Underline)
	fmt.Println("add colored output to your console")
	colorize.MoveCursorBy(9, 0)
	fmt.Println("100% compatible with every os!")
	colorize.ResetStyle()

	// print line of colors
	fmt.Println()
	for i := 0; i < 48; i++ {
		colorize.SetBgTruecolorHSL(uint8(i*255/48), 255, 127)
		fmt.Print(" ")
		colorize.ResetStyle()
	}
	fmt.Println()

	// print colored squares
	fmt.Println()
	for i := 0; i < 48; i++ {
		colorize.SetFgTruecolorHSL(uint8(255-i*255/48), 255, 127)
		fmt.Print("■")
		colorize.ResetStyle()
	}
	fmt.Println()

	fmt.Println()
	colorize.MoveCursorBy(6, 0)
	colorize.SetStyle(colorize.FgRed, colorize.Strike)
	fmt.Println("Boring, complex, convoluted packages")
	colorize.ResetStyle()
	colorize.MoveCursorBy(6, 0)
	colorize.SetStyle(colorize.FgBrightGreen, colorize.Bold)
	fmt.Println("Easy to use and documented colorize!")
	colorize.ResetStyle()
	fmt.Println()

	fmt.Println()
}

func rainbowPoints() {
	const size float64 = 40
	// initialize the random number generator
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// get a random seed
	seed := uint8(r1.Intn(255))

	fmt.Println()
	for x := 0.0; x < size; x++ {
		for y := 0.0; y < size; y++ {
			dist := math.Sqrt(x*x + y*y)
			hue := dist / (size * math.Sqrt(2)) * 255
			colorize.SetFgTruecolorHSL(uint8(hue)+seed, 255, 127)
			fmt.Print("#")
			colorize.ResetStyle()
		}
		fmt.Println()
	}
	fmt.Println()
}

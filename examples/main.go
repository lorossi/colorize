package main

import (
	"fmt"
	"math/rand"

	"github.com/lorossi/colorize"
)

func main() {
	rainbow()
	advertising()
	christmasTree()
}

func rainbow() {
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

	colorize.SetStyle(colorize.FgBrightRed, colorize.BgWhite, colorize.Bold, colorize.Underline, colorize.RapidBlink)
	fmt.Println("NEW PACKAGE!")
	colorize.ResetStyle()

	colorize.SetStyle(colorize.FgBlue, colorize.BgWhite, colorize.Bold)
	fmt.Println("Offering:")
	colorize.ResetStyle()

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
	fmt.Println()
	// initialize tree width
	width := 11
	// initialize light colors
	colors := [4]colorize.Style{colorize.FgBrightRed, colorize.FgBrightBlue, colorize.FgBrightGreen, colorize.FgBrightYellow}
	// print star on top
	colorize.MoveCursorBy(int8(width), 0)
	fmt.Println(colorize.Yellow("X"))
	for i := 0; i < width; i++ {
		// leave space on the left
		colorize.MoveCursorBy(int8(width-i), 0)
		// prepare the character
		colorize.SetStyle(colorize.RapidBlink, colorize.BgBrightGreen)
		for k := 0; k < 2*i+1; k++ {
			if rand.Float64() > 0.66 {
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

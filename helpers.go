package main

import "github.com/gdamore/tcell/v2"

// Preset styles
var defStyle tcell.Style = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
var blueForeground tcell.Style = tcell.StyleDefault.Foreground(tcell.ColorBlue)
var darkBlueForeground tcell.Style = tcell.StyleDefault.Foreground(tcell.ColorDarkBlue)
var yellowForeground tcell.Style = tcell.StyleDefault.Foreground(tcell.ColorYellow)
var greenForeground tcell.Style = tcell.StyleDefault.Foreground(tcell.ColorGreen)
var redForeground tcell.Style = tcell.StyleDefault.Foreground(tcell.ColorRed)
var darkRedForeground tcell.Style = tcell.StyleDefault.Foreground(tcell.ColorDarkRed)

// Draws text on single line
func drawString(screen tcell.Screen, x, y int, msg string) {
	for i, char := range []rune(msg) {
		screen.SetContent(x+i, y, char, nil, tcell.StyleDefault)
	}
}

// Draws test on single line with given style
func drawStyledString(screen tcell.Screen, x, y int, style tcell.Style, msg string) {
	for i, char := range []rune(msg) {
		screen.SetContent(x+i, y, char, nil, style)
	}
}

// Draws multi-line text
func drawText(screen tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		screen.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 || r == '\n' {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

// Draws a simple box with the given message. Message will be put on multiple
// lines if too long.
func drawBoxWithMessage(screen tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			screen.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		screen.SetContent(col, y1, tcell.RuneHLine, nil, style)
		screen.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		screen.SetContent(x1, row, tcell.RuneVLine, nil, style)
		screen.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		screen.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		screen.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		screen.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		screen.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	drawText(screen, x1+1, y1+1, x2-1, y2-1, style, text)
}

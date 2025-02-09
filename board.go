package main

import "github.com/gdamore/tcell/v2"

type Mark int

const (
	Empty Mark = iota
	X
	O
	Blocked
)

const (
	empty   = "   "
	xMark   = " X "
	oMark   = " O "
	blocked = "▓▓▓"
)

type Board struct {
	x             int
	y             int
	state         []Mark
	selectedIndex int
	showBoard     bool
}

func NewBoard(x, y int) Board {
	board := Board{
		x:             x,
		y:             y,
		state:         make([]Mark, 25),
		selectedIndex: 0,
		showBoard:     false,
	}

	for index := range 25 {
		board.state[index] = Empty
	}

	board.state[12] = Blocked

	return board
}

func (b *Board) Render(screen tcell.Screen) {
	drawString(screen, b.x, b.y, "┌───┬───┬───┬───┬───┐")
	drawString(screen, b.x, b.y+1, "│   │   │   │   │   │")
	drawString(screen, b.x, b.y+2, "├───┼───┼───┼───┼───┤")
	drawString(screen, b.x, b.y+3, "│   │   │   │   │   │")
	drawString(screen, b.x, b.y+4, "├───┼───┼───┼───┼───┤")
	drawString(screen, b.x, b.y+5, "│   │   │   │   │   │")
	drawString(screen, b.x, b.y+6, "├───┼───┼───┼───┼───┤")
	drawString(screen, b.x, b.y+7, "│   │   │   │   │   │")
	drawString(screen, b.x, b.y+8, "├───┼───┼───┼───┼───┤")
	drawString(screen, b.x, b.y+9, "│   │   │   │   │   │")
	drawString(screen, b.x, b.y+10, "└───┴───┴───┴───┴───┘")

	for index, mark := range b.state {
		x := index % 5
		y := index / 5
		xPos := x*4 + 1 + b.x
		yPos := y*2 + 1 + b.y

		switch mark {
		case Empty:
			if index == b.selectedIndex {
				drawStyledString(screen, xPos, yPos, tcell.StyleDefault.Background(tcell.ColorYellowGreen), empty)
			} else {
				drawString(screen, xPos, yPos, empty)
			}
		case X:
			if index == b.selectedIndex {
				drawStyledString(screen, xPos, yPos, tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorYellow), xMark)
			} else {
				drawString(screen, xPos, yPos, xMark)
			}
		case O:
			if index == b.selectedIndex {
				drawStyledString(screen, xPos, yPos, tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorYellow), oMark)
			} else {
				drawString(screen, xPos, yPos, oMark)
			}
		case Blocked:
			if index == b.selectedIndex {
				drawStyledString(screen, xPos, yPos, tcell.StyleDefault.Foreground(tcell.ColorDarkRed), blocked)
			} else {
				drawString(screen, xPos, yPos, blocked)
			}
		}
	}
}

func (b *Board) MoveIndexUp() {
	if !b.showBoard {
		return
	}
	newIndex := b.selectedIndex - 5
	if newIndex < 0 {
		return
	}
	b.selectedIndex = newIndex
}

func (b *Board) MoveIndexDown() {
	if !b.showBoard {
		return
	}
	newIndex := b.selectedIndex + 5
	if newIndex > 24 {
		return
	}
	b.selectedIndex = newIndex
}

func (b *Board) MoveIndexLeft() {
	if !b.showBoard || b.selectedIndex%5 == 0 {
		return
	}
	b.selectedIndex--
}

func (b *Board) MoveIndexRight() {
	if !b.showBoard || b.selectedIndex%5 == 4 {
		return
	}
	b.selectedIndex++
}

func (b *Board) GetAvailableIndexes() []int {
	available := make([]int, 0)
	for index, mark := range b.state {
		if mark == Empty {
			available = append(available, index)
		}
	}
	return available
}

func (b *Board) ShowBoard() {
	b.showBoard = true
}

func (b *Board) HideBoard() {
	b.showBoard = false
}

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
	computerScore int
	playerScore   int
}

func NewBoard(x, y int) Board {
	board := Board{
		x:             x,
		y:             y,
		state:         make([]Mark, 25),
		selectedIndex: 0,
		showBoard:     false,
		computerScore: 0,
		playerScore:   0,
	}

	for index := range 25 {
		board.state[index] = Empty
	}

	board.state[12] = Blocked

	return board
}

func (b *Board) Render(screen tcell.Screen) {
	if !b.showBoard {
		return
	}
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
				drawStyledString(screen, xPos, yPos, darkRedForeground, blocked)
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

func (b *Board) SelectedIndexValid() bool {
	return b.state[b.selectedIndex] == Empty
}

func (b *Board) Show() {
	b.showBoard = true
}

func (b *Board) Hide() {
	b.showBoard = false
}

func (b *Board) ProcessMoves(playerMove, computerMove int) {
	if playerMove == computerMove {
		b.state[playerMove] = Blocked
		return
	}

	b.state[playerMove] = X
	b.state[computerMove] = O
}

func (b *Board) CalculateScores() {
	playerScore := 0
	computerScore := 0
	for i := 0; i < len(b.state); i++ {
		// Get current mark
		mark := b.state[i]

		// We only need to process X's and O's
		if mark != X && mark != O {
			continue
		}

		// Look right for 3 in a row
		x := i % 5

		// Make sure there is enough room for 3
		if x+2 < 5 {
			// Check if next two match current mark
			if b.state[i+1] == mark && b.state[i+2] == mark {
				// Give the point to the proper player
				if mark == X {
					playerScore++
				} else {
					computerScore++
				}
			}
		}

		// Look down for 3 in a row
		y := i / 5

		// Make sure there is enough room for 3
		if y+2 < 5 {
			// Check if next two match current mark
			if b.state[i+5] == mark && b.state[i+10] == mark {
				// Give the point to the proper player
				if mark == X {
					playerScore++
				} else {
					computerScore++
				}
			}
		}

	}

	b.playerScore = playerScore
	b.computerScore = computerScore
}

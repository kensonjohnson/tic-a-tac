package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type game struct {
	screen    tcell.Screen
	running   bool
	board     Board
	gameState gameState
}

type gameState int

const (
	WelcomeScreen gameState = iota
	ShowBoard
	PlayerInput
)

var defStyle tcell.Style = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

func main() {

	// Initialize screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	screen.SetStyle(defStyle)
	screen.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Game initialization
	g := game{
		screen:    screen,
		running:   true,
		board:     NewBoard(1, 1),
		gameState: WelcomeScreen,
	}

	// Event loop
	for g.running {
		g.update()
		g.render()
	}
}

func (g *game) update() {
	event := g.screen.PollEvent()
	ev, isKeyEvent := event.(*tcell.EventKey)
	if isKeyEvent && (ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC) {
		g.running = false
		return
	}

	switch g.gameState {
	case WelcomeScreen:
		if isKeyEvent && ev.Key() == tcell.KeyEnter {
			g.gameState = ShowBoard
		}

	case ShowBoard:
		g.board.ShowBoard()

	case PlayerInput:
		if isKeyEvent {
			switch {
			case ev.Rune() == 'w' || ev.Key() == tcell.KeyUp:
				g.board.MoveIndexUp()

			case ev.Rune() == 's' || ev.Key() == tcell.KeyDown:
				g.board.MoveIndexDown()

			case ev.Rune() == 'a' || ev.Key() == tcell.KeyLeft:
				g.board.MoveIndexLeft()

			case ev.Rune() == 'd' || ev.Key() == tcell.KeyRight:
				g.board.MoveIndexRight()
			}
		}
	}
}

func (g *game) render() {
	g.screen.Clear()
	g.board.Render(g.screen)
	g.screen.Show()
}

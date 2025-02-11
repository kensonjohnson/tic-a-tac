package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

type game struct {
	screen     tcell.Screen
	running    bool
	welcome    Welcome
	board      Board
	scoreboard Scoreboard
	computer   Computer
	player     Player
	gameState  GameState
}

type GameState int

const (
	WelcomeScreen GameState = iota
	ShowBoard
	ComputerInput
	PlayerInput
	TakeTurn
	RoundOver
	GameOver
)

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
		screen:     screen,
		running:    true,
		welcome:    NewWelcome(),
		board:      NewBoard(1, 1),
		scoreboard: NewScoreboard(40, 0),
		computer:   NewComputer(),
		player:     NewPlayer(),
		gameState:  WelcomeScreen,
	}

	// Event loop
	for g.running {
		g.update()
		g.render()
	}
}

func (g *game) update() {
	switch g.gameState {
	case WelcomeScreen:
		event := g.screen.PollEvent()
		ev, isKeyEvent := event.(*tcell.EventKey)

		if !isKeyEvent {
			return
		}

		switch {
		case ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC:
			g.running = false
			return

		case ev.Key() == tcell.KeyEnter:
			g.gameState = ShowBoard
		}

	case ShowBoard:
		g.welcome.Hide()
		g.board.Show()
		g.scoreboard.Show()
		g.gameState = ComputerInput

	case ComputerInput:
		g.computer.makeNextMove(g.board)
		g.gameState = PlayerInput

	case PlayerInput:
		event := g.screen.PollEvent()
		ev, isKeyEvent := event.(*tcell.EventKey)

		if !isKeyEvent {
			return
		}

		switch {
		case ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC:
			g.running = false
		case ev.Rune() == 'w' || ev.Key() == tcell.KeyUp:
			g.board.MoveIndexUp()

		case ev.Rune() == 's' || ev.Key() == tcell.KeyDown:
			g.board.MoveIndexDown()

		case ev.Rune() == 'a' || ev.Key() == tcell.KeyLeft:
			g.board.MoveIndexLeft()

		case ev.Rune() == 'd' || ev.Key() == tcell.KeyRight:
			g.board.MoveIndexRight()

		case ev.Key() == tcell.KeyEnter:
			if g.player.attemptMove(g.board) {
				g.gameState = TakeTurn
			}

		}

	case TakeTurn:
		g.board.ProcessMoves(g.player.nextMove, g.computer.nextMove)
		g.board.CalculateScores()
		g.scoreboard.UpdateScores(g.board.playerScore, g.board.computerScore)
		g.gameState = ComputerInput
		if len(g.board.GetAvailableIndexes()) == 0 {
			g.gameState = RoundOver
		}

	case RoundOver:
		// Display exit screen
		g.board.Hide()
		g.gameState = GameOver

	case GameOver:
		event := g.screen.PollEvent()
		ev, isKeyEvent := event.(*tcell.EventKey)
		if isKeyEvent && (ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC) {
			g.running = false
		}
	}
}

func (g *game) render() {
	g.screen.Clear()
	g.welcome.Render(g.screen)
	g.board.Render(g.screen)
	g.scoreboard.Render(g.screen)
	g.drawGameOver()
	g.screen.Show()
}

func (g *game) drawGameOver() {
	if g.gameState != GameOver {
		return
	}
	drawString(g.screen, 2, 2, "Game Over!")

	drawString(g.screen, 2, 4, fmt.Sprintf("Player score: %v", g.board.playerScore))
	drawString(g.screen, 2, 5, fmt.Sprintf("Computer score: %v", g.board.computerScore))

	if g.board.playerScore > g.board.computerScore {
		drawString(g.screen, 2, 7, "Player Wins!")
	} else if g.board.playerScore < g.board.computerScore {
		drawString(g.screen, 2, 7, "Player Wins!")
	} else {
		drawString(g.screen, 2, 7, "It's a tie!")
	}
}

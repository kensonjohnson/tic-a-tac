package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Scoreboard struct {
	X, Y          int
	playerScore   int
	computerScore int
	show          bool
}

func NewScoreboard(x, y int) Scoreboard {
	return Scoreboard{
		X: x,
		Y: y,
	}
}

func (s *Scoreboard) Render(screen tcell.Screen) {
	if !s.show {
		return
	}
	drawBoxWithMessage(screen, s.X, s.Y, s.X+14, s.Y+3, blueForeground, fmt.Sprintf("Player:   %v\nComputer: %v", s.playerScore, s.computerScore))
}

func (s *Scoreboard) Show() {
	s.show = true
}

func (s *Scoreboard) Hide() {
	s.show = false
}

func (s *Scoreboard) UpdateScores(player, computer int) {
	s.playerScore = player
	s.computerScore = computer
}

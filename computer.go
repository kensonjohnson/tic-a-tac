package main

import "math/rand"

type Computer struct {
	nextMove int
}

func NewComputer() Computer {
	return Computer{}
}

func (c *Computer) makeNextMove(board Board) {
	available := board.GetAvailableIndexes()
	c.nextMove = available[rand.Intn(len(available))]
}

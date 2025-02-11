package main

type Player struct {
	nextMove int
}

func NewPlayer() Player {
	return Player{}
}

func (p *Player) attemptMove(board Board) bool {
	if board.SelectedIndexValid() {
		p.nextMove = board.selectedIndex
		return true
	}
	return false
}

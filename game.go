package main

// Won indicates when the game is won by the player
const Won = 1

// Board struct.
type Board struct {
}

// NewBoard creates a new Board.
func NewBoard(i int) Board {
	b := Board{}
	return b
}

// GameOver function.
func GameOver(b Board) int {
	return Won
}

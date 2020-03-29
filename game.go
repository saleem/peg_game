package main

// Won indicates when the game has been won
const Won = 1

// Lost indicates when the game has been lost
const Lost = 2

// Board struct.
type Board struct {
	a [15]bool
}

// NewBoard creates a new Board.
func NewBoard(pegs ...int) Board {
	b := Board{}
	for _, i := range pegs {
		b.a[i] = true
	}
	return b
}

// GameOver function.
func GameOver(b Board) int {
	if pegCount(b) == 1 {
		return Won
	} else {
		return Lost
	}
}

func pegCount(b Board) int {
	count := 0
	for i := 0; i < len(b.a); i++ {
		if b.a[i] == true {
			count = count + 1
		}
	}
	return count
}

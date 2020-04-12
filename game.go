package main

// Result indicates the outcome of the game
type Result int

const (
	Won Result = 1 + iota
	Lost
	InProgress
)

// Direction type.
type Direction pair

// Directions.
// The value of each Direction is a pair of numbers. 
// The first value in the pair represents how row-values change
//   as we move in that direction.
// The second value in the pair represents how column-values change
//   as we move in that direction.
// Because of the topology of the board, there are six Directions:
//   SE - Southeast (row- and column-values increase)
//   NW - Northwest (row- and column-values decrease)
//   SW - Southwest (row-values increase)
//   NE - Northeast (row-values decrease)
//   E  - East (column-values increase)
//   W  - West (column-values decrease)
var (
	SE Direction = Direction{1, 1}
	NW Direction = Direction{-1, -1}
	SW Direction = Direction{1, 0}
	NE Direction = Direction{-1, 0}
	E Direction = Direction{0, 1}
	W Direction = Direction{0, -1}
)

// rows is the number of rows on board.
// rows is hard-coded to 5 since our board-size is fixed.
const rows = 5

// slots is the number of slots.
// It can be defined using the formula for triangular numbers.
const slots = rows * (rows + 1) / 2

// Board struct.
type Board struct {
	// This is a wasteful array: unused entries = rows * (rows - 1) / 2
	a [rows][rows]bool
}

func (b *Board) fillSlot(loc int, peg bool) {
	row, col := locToIndexes(loc)
	b.a[row][col] = peg
}

func (b Board) getSlot(loc int) bool {
	row, col := locToIndexes(loc)
	return b.a[row][col]
}

func locToIndexes(loc int) (int, int) {
	row := -1
	for i := 1; i <= rows && row == -1; i++ {
		if i*(i+1)/2 >= loc {
			row = i - 1
		}
	}
	col := loc - (row * (row + 1) / 2) - 1
	return row, col
}

// NewBoard creates a new Board.
func NewBoard(pegs ...int) Board {
	b := Board{}
	for _, i := range pegs {
		b.fillSlot(i, true)
	}
	return b
}

// GameOver function.
func GameOver(b Board) Result {
	if pegCount(b) == 1 {
		return Won
	} else if moveExists(b) {
		return InProgress
	} else {
		return Lost
	}
}

func pegCount(b Board) int {
	count := 0
	for i := 1; i <= slots; i++ {
		if b.getSlot(i) {
			count = count + 1
		}
	}
	return count
}

func moveExists(b Board) bool {
	return false
}


func adjacentSlots(b Board, d Direction) (int, int) {
	return 2, 4
}

type pair struct {
	row, col int
}

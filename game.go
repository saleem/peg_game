package main

// Slot type
type Slot int

const (
	// Undefined means the Slot is undefined, e.g. "Which slot is above the topmost slot?"
	Undefined Slot = iota
	// Empty means the Slot does not have a Peg
	Empty
	// Full means the Slot has a Peg
	Full
)

// Result indicates the outcome of the game
type Result int

const (
	Won Result = 1 + iota
	Lost
	InProgress
)

// Direction type.
type Direction pair

// Directions map.
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
var Directions map[string]Direction = map[string]Direction{
	"SE": {1, 1},
	"NW": {-1, -1},
	"SW": {1, 0},
	"NE": {-1, 0},
	"E":  {0, 1},
	"W":  {0, -1},
}

// rows is the number of rows on board.
// rows is hard-coded to 5 since our board-size is fixed.
const rows = 5

// slots is the number of slots.
// It can be defined using the formula for triangular numbers.
const slots = rows * (rows + 1) / 2

// Board struct.
type Board struct {
	// This is a wasteful array: unused entries = rows * (rows - 1) / 2
	a [rows][rows]Slot
}

func (b *Board) fillSlot(loc int, peg Slot) {
	row, col := locToIndexes(loc)
	b.a[row][col] = peg
}

func (b Board) getSlot(loc int) Slot {
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

func indexesToLoc(row int, col int) int {
	return row*(row+1)/2 + col + 1
}

// NewBoard creates a new Board.
func NewBoard(pegs ...int) Board {
	b := Board{}
	for _, i := range pegs {
		b.fillSlot(i, Full)
	}
	for i := 1; i <= slots; i++ {
		if b.getSlot(i) != Full {
			b.fillSlot(i, Empty)
		}
	}
	return b
}

// GameStatus function.
func GameStatus(b Board) Result {
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
		if b.getSlot(i) == Full {
			count = count + 1
		}
	}
	return count
}

func moveExists(b Board) bool {
	// for each empty peg
	//   check if there are two pegs in any consecutive direction
	//   if yes, return true
	//   if no for all empty pegs, return false
	retVal := false
	for i := 1; i <= slots && !retVal; i++ {
		if b.getSlot(i) == Empty {
			for _, dir := range Directions {
				one, two := getTwoAdjacentSlots(b, i, dir)
				if one == Full && two == Full {
					retVal = true
				}
			}
		}
	}
	return retVal
}

// getTwoAdjacentSlots returns the two adjacent pegs in the given direction
func getTwoAdjacentSlots(b Board, loc int, dir Direction) (Slot, Slot) {
	i, j := locToIndexes(loc)
	nOneX, nOneY := i+dir.row, j+dir.col
	if nOneX < 0 || nOneY < 0 || nOneX >= rows || nOneY >= rows {
		return Undefined, Undefined
	}
	neighborOne := b.getSlot(indexesToLoc(nOneX, nOneY))
	nTwoX, nTwoY := nOneX+dir.row, nOneY+dir.col
	if nTwoX < 0 || nTwoY < 0 || nTwoX >= rows || nTwoY >= rows {
		return neighborOne, Undefined
	}
	neighborTwo := b.getSlot(indexesToLoc(nTwoX, nTwoY))
	return neighborOne, neighborTwo
}

type pair struct {
	row, col int
}

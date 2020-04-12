package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnitPlayerWinsWhenThereIsOnePegLeft(t *testing.T) {
	for i := 1; i <= 15; i++ {
		b := NewBoard(i)
		assert.Equal(t, Won, GameOver(b))
	}
}

func TestUnitPlayerLosesWhenThereAreTwoDistantPegsLeft(t *testing.T) {
	b := NewBoard(1, 4)
	assert.Equal(t, Lost, GameOver(b))
}

func TestUnitGameIsInProgressWhenThereAreTwoAdjacentPegsLeft(t *testing.T) {
	b := NewBoard(1, 2)
	assert.Equal(t, InProgress, GameOver(b))
}

func TestLocToRowCol(t *testing.T) {
	m := make(map[int]pair)
	m[1] = pair{0, 0}
	m[2] = pair{1, 0}
	m[3] = pair{1, 1}
	m[4] = pair{2, 0}
	m[5] = pair{2, 1}
	m[6] = pair{2, 2}
	m[7] = pair{3, 0}
	m[8] = pair{3, 1}
	m[9] = pair{3, 2}
	m[10] = pair{3, 3}
	m[11] = pair{4, 0}
	m[12] = pair{4, 1}
	m[13] = pair{4, 2}
	m[14] = pair{4, 3}
	m[15] = pair{4, 4}

	row, col := 0, 0
	for k, v := range m {
		row, col = locToIndexes(k)
		assert.Equal(t, v.row, row)
		assert.Equal(t, v.col, col)
	}
}

func TestUnitAdjacentSlots(t *testing.T) {
	b := NewBoard()
	i, j := adjacentSlots(b, SW)
	assert.Equal(t, 2, i)
	assert.Equal(t, 4, j)
}

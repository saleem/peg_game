package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnitPlayerWinsWhenThereIsOnePegLeft(t *testing.T) {
	for i := 1; i <= 15; i++ {
		b := NewBoard(i)
		assert.Equal(t, Won, GameStatus(b))
	}
}

func TestUnitPlayerLosesWhenThereAreTwoDistantPegsLeft(t *testing.T) {
	b := NewBoard(1, 4)
	assert.Equal(t, Lost, GameStatus(b))
}

func TestUnitGameIsInProgressWhenThereAreTwoAdjacentPegsLeftInAnyOrientation(t *testing.T) {
	b := NewBoard(1, 2) // Move exists: NE from slot 4
	assert.Equal(t, InProgress, GameStatus(b))

	b = NewBoard(1, 3) // Move exists: NW from slot 6
	assert.Equal(t, InProgress, GameStatus(b))

	b = NewBoard(7, 11) // Move exists: SW from slot 4
	assert.Equal(t, InProgress, GameStatus(b))

	b = NewBoard(10, 15) // Move exists: SE from slot 6
	assert.Equal(t, InProgress, GameStatus(b))

	b = NewBoard(4, 5) // Move exists: W from slot 6
	assert.Equal(t, InProgress, GameStatus(b))

	b = NewBoard(5, 6) // Move exists: E from slot 4
	assert.Equal(t, InProgress, GameStatus(b))
}

func TestUnitPlayerLosesWhenAnyOneRowIsFullOfPegs(t *testing.T) {
	b := NewBoard(1, 2, 4, 7, 11)
	assert.Equal(t, Lost, GameStatus(b))

	b = NewBoard(1, 3, 6, 10, 15)
	assert.Equal(t, Lost, GameStatus(b))

	b = NewBoard(11, 12, 13, 14, 15)
	assert.Equal(t, Lost, GameStatus(b))
}

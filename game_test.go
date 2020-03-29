package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitPlayerWinsWhenThereIsOnePegLeft(t *testing.T) {
	b := NewBoard(1)
	assert.Equal(t, Won, GameOver(b))
}

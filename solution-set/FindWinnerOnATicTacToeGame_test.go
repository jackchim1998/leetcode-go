package solutionSet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAWin(t *testing.T) {
	result := tictactoe([][]int{
		{0, 0},
		{2, 0},
		{1, 1},
		{2, 1},
		{2, 2},
	})
	assert.Equal(t, "A", result)
}

func TestBWin(t *testing.T) {
	result := tictactoe([][]int{
		{0, 0},
		{1, 1},
		{0, 1},
		{0, 2},
		{1, 0},
		{2, 0},
	})
	assert.Equal(t, "B", result)
}

func TestDraw(t *testing.T) {
	result := tictactoe([][]int{
		{0, 0},
		{1, 1},
		{2, 0},
		{1, 0},
		{1, 2},
		{2, 1},
		{0, 1},
		{0, 2},
		{2, 2},
	})
	assert.Equal(t, "Draw", result)
}

func TestPending(t *testing.T) {
	result := tictactoe([][]int{
		{0, 0},
		{1, 1},
	})
	assert.Equal(t, "Pending", result)
}

package dfs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoOfIsLandDFS(t *testing.T) {
	grid1 := [][]byte{
		{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','0','0'},
	}

	grid2 := [][]byte{
		{'1','1','0','0','0'},
		{'1','1','0','0','0'},
		{'0','0','1','0','0'},
		{'0','0','0','1','1'},
	}
	assert.Equal(t, 1, numIslands(grid1))
	assert.Equal(t, 3, numIslands(grid2))
}

func TestSurroundedRegions(t *testing.T) {
	board1 := [][]byte{
		{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'},
	}

	board2 := [][]byte{
		{'X','X','X','X'},{'X','X','X','X'},{'X','X','X','X'},{'X','O','X','X'},
	}

	surroundedRegion(board1)
	assert.Equal(t, board2, board1)
}
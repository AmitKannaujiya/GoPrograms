package bfs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoOfIsLandBFS(t *testing.T) {
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